package utils

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"fmt"
	"time"
	"regexp"
	"strconv"
	"organ-go-api/models"
	"github.com/blevesearch/bleve"
	"encoding/json"
)

var batchSize = flag.Int("batchSize", 100, "batch size for indexing") // 批量大小
var dataDir = flag.String("dataDir", "data", "data directory")        // data文件夹
var IsIndexCreating = false                                           // 创建构建标识，用于阻塞重复标识
var IsStopIndex = false                                               // 用于停止索引标识
var CurrentIndex, _ = InitIndex()                                     // 初始化索引

// 创建索引
func CreateIndex() {
	// 判断是否正在构建, 防止重复构建
	if (IsIndexCreating) {
		fmt.Println("有一个索引正在构建，请等待！")
		return
	}
	IsIndexCreating = true // 正在构建标识

	// 创建本地索引文件夹
	timeStr := time.Now().Format("2006-01-02 15:04:05") // 获取当前时间戳字符串形式
	reg := regexp.MustCompile("[' ' \\- :]")            // 2017-08-25 08:47:28 => 20170825084728
	timeRes := reg.ReplaceAllString(timeStr, "")
	path := *dataDir + string(os.PathSeparator) + timeRes
	fmt.Println("当前索引文件夹名称为：", path)

	mapping := bleve.NewIndexMapping()     // 创建一个索引 Mapping
	index, err := bleve.New(path, mapping) // 创建一个索引引擎
	if err != nil {
		log.Fatal("创建索引失败: " + err.Error())
		return
	}

	forFlag := true            // 循环标识
	total := models.GetTotal() // 查出所有数据总量
	var currentPage int64 = 0  // 当前页码为
	var currentIndex int64 = 0 // 当前纪录

	batch := index.NewBatch() // 生成batch对象

	// 将索引信息写入文件 	// 字符类型的文档总数
	WriteInfo(timeRes, path, total, 0, 0, false, true, false) // 正在创建索引

	startTime := time.Now() // 获取当前时间
	// 用 for 实现一个 while 循环
	for forFlag == true {
		organs, err := models.GetOrganByPage(currentPage, 5000) // 以5000条数据为单位分批读取
		// 处理异常错误
		if err != nil {
			fmt.Println(err)
			forFlag = false
			break
		}
		organs_len := int64(len(organs)) // 获取读取到的纪录的int型的长度
		currentIndex += organs_len       // 获取当前纪录数
		var docList []models.Organ       // 用于存储当前批中document数据的集合
		batchCount := 0                  // 用于计算以100为单位内的数量
		count := 0                       // 用于计算总构建量

		// 判断读取完成,或者中途停止 => 退出for循环
		if organs_len <= 0 || IsStopIndex {
			forFlag = false // 退出循环标识
			IsIndexCreating = false
			// 添加及时停止的判断
			if IsStopIndex {
				IsStopIndex = false // 回归
				// 写入文件,将索引信息写入文件
				WriteIndexInfo(startTime, timeRes, path, total, false, false, false)
				fmt.Println("已停止对索引的构建！")
				break
			}

			// 所有构建完成之后将索引信息写入文件
			WriteIndexInfo(startTime, timeRes, path, total, true, false, false)
			fmt.Println("已完成当前索引的构建！")
			break
		}

		fmt.Println("---- 已批量读取了", organs_len, "条数据,进度为: ", currentIndex, "/", total, ",正在构建中,请稍后： ----")

		// 处理organs数据, 将别名信息挂载在上面, 并循环处理，得到当前docList集合
		for _, v := range organs {
			OrganAliasNames, err := models.GetOrganAliasNameByUuid(v.Uuid) // 通过uuid来查库
			// 处理异常错误
			if err != nil {
				fmt.Println(err)
				forFlag = false
				break
			}
			var oa_string string = "" // 初始化机构别名
			// 将机构别名提取出来整理成字符串
			for _, oaVal := range OrganAliasNames {
				name := oaVal["Name"] // 读取出数据
				if _, ok := name.(string); ok {
					oa_string += name.(string) + ","
				}
			}
			v.Oa = oa_string // 挂载

			// 构建单条document对象
			data := models.Organ{
				Id:             v.Id,
				Uuid:           v.Uuid,
				Oid:            v.Oid,
				NameEn:         v.NameEn,
				NameCn:         v.NameCn,
				NameOrigin:     v.NameOrigin,
				ShortName:      v.ShortName,
				FirstLetter:    v.FirstLetter,
				Location:       v.Location,
				CountryCode:    v.CountryCode,
				CountryCn:      v.CountryCn,
				CountryEn:      v.CountryEn,
				Province:       v.Province,
				ProvinceCode:   v.ProvinceCode,
				ProvincePinyin: v.ProvincePinyin,
				City:           v.City,
				AbsEn:          v.AbsEn,
				AbsCn:          v.AbsCn,
				Oa:             v.Oa,
			}
			docList = append(docList, data) // 放入docList集合中
		}

		// 对当前的docList进行构建索引处理
		for _, v := range docList {
			// 添加及时停止的判断
			if IsStopIndex {
				break
			}
			batch.Index(strconv.Itoa(v.Id), v) // 将document添加进batch中
			batchCount++
			// 判断构建的数量是否大于批量设置的大小
			if batchCount >= *batchSize {
				err = index.Batch(batch) // 100为单位结束，着手进行构建
				if err != nil {
					fmt.Println("构建时出现错误：", err.Error())
					return
				}
				batch = index.NewBatch() // 重新new出一个batch
				batchCount = 0           // 重新归零
			}

			count++ // 总量自增

			// 以1000个document为单位显示构建效率
			if count%1000 == 0 {
				indexDuration := time.Since(startTime)
				indexDurationSeconds := float64(indexDuration) / float64(time.Second) // 动态获取构建时间，单位为秒
				timePerDoc := float64(indexDuration) / float64(count)                 // 动态获取构建每个文档用时
				log.Printf("构建了 %d／%d, 用时 %.2fs (平均效率 %.2fms/doc)", count, organs_len, indexDurationSeconds, timePerDoc/float64(time.Millisecond))
			}
		}

		// 构建最后一次不满足100条的索引
		if batchCount > 0 {
			err = index.Batch(batch)
			if err != nil {
				log.Fatal(err)
			}

			// 最终的5000条数据构建完成后的重新动态打印构建的数量,并计算构建每个文档耗时
			indexDuration := time.Since(startTime)
			indexDurationSeconds := float64(indexDuration) / float64(time.Second) // 动态获取构建时间，单位为秒
			timePerDoc := float64(indexDuration) / float64(count)                 // 动态获取构建每个文档用时
			log.Printf("构建了 %d／%d, 用时 %.2fs (平均效率 %.2fms/doc)", count, organs_len, indexDurationSeconds, timePerDoc/float64(time.Millisecond))
		}

		currentPage ++ // 页码自增
	}
}

// 停止索引
func StopIndex() {
	if IsIndexCreating == false {
		fmt.Println("未构建，无效的指令！")
		return
	}
	IsStopIndex = true // 发出停止的信号
}

// 激活索引
func ActiveIndex(id string) (err error) {
	path := *dataDir + string(os.PathSeparator) + id // 获取当前索引路径
	// 关闭之前索引
	if CurrentIndex.Index != nil {
		fmt.Println("正在关闭之前的索引...")
		err = CurrentIndex.Index.Close() // 关闭之前索引
		if err != nil {
			fmt.Println("关闭之前索引出现错误：", err.Error())
		} else {
			fmt.Println("关闭成功！")
			fmt.Println("--------------")
			fmt.Println("正在切换新的索引，文件较大请稍后！")
			// 切换索引慢的issue : https://github.com/blevesearch/bleve/issues/187
			startTime := time.Now()                    // 获取当前时间
			CurrentIndex.Index, err = bleve.Open(path) // 打开当前索引, 重新赋值
			if err != nil {
				fmt.Println("切换失败：", err.Error())
			} else {
				CurrentIndex.Id = id
				indexDuration := time.Since(startTime)
				indexDurationSeconds := float64(indexDuration) / float64(time.Second) // 动态获取构建时间，单位为秒
				t := strconv.FormatFloat(indexDurationSeconds, 'E', -1, 32)           // 统计时间
				fmt.Println("切换成功! 用时：" + t)
			}
		}
	}
	return err
}

// 切换索引
func SwapIndex(id string) (err error) {
	path := *dataDir + string(os.PathSeparator) + id // 获取当前索引路径

	// 步骤1：创建一份新的索引
	mapping := bleve.NewIndexMapping()     // 创建一个索引 Mapping
	index, err := bleve.New(*dataDir, mapping) // 创建一个索引引擎
	if err != nil {
		log.Fatal("创建索引失败: " + err.Error())
		return
	}

	// 2. Create an IndexAlias pointing to the Index
	// 3. Write your application to query the IndexAlias
	// 4. …At some later time, create a new Index
	// 5. Call Swap() on the IndexAlias passing in the new Index

	if false {
		fmt.Println(path)
		fmt.Println(index)
	}

	return err
}

// 构建通用索引结构
type CurrentIndexStruct struct {
	Index bleve.Index
	Id    string
}

// 初始化索引 将有效最新的索引, 作为默认索引
func InitIndex() (ci CurrentIndexStruct, err error) {
	// 读文件，遍历：找最新的，找index_info.json文件中 isEffective 为true的
	dirEntries, err := ioutil.ReadDir(*dataDir)
	if err != nil {
		log.Fatalf("error reading data dir: %v", err)
	}
	dirLenth := len(dirEntries) // 列出索引文件夹中的文件数量
	// 没有索引
	if (dirLenth == 0) {
		fmt.Println("当前没有索引文件")
		return ci, err
	}
	var ii IndexInfo // 用于序列化数据使用的
	var IndexArr []int64 // 创建一个数组,用于存储 IsEffective 为 true 的索引名称
	// 循环里面的每一个索引文件夹
	for _, dirInfo := range dirEntries {
		// 文件类型跳过
		if !dirInfo.IsDir() {
			continue
		}
		// 循环每个文件夹中的文件读取出相应的索引配置文件信息
		currentDir, err := ioutil.ReadDir(*dataDir + string(os.PathSeparator) + dirInfo.Name())
		if err != nil {
			fmt.Println("发生错误：", err.Error())
			break
		}

		// 先解析出有用的
		for _, v := range currentDir {
			// 对成功写入配置信息文件的，解析出对象，然后进行
			if v.Name() == "index_info.json" {
				// 读取出当前index_info.json文件
				by, err := ioutil.ReadFile(*dataDir + string(os.PathSeparator) + dirInfo.Name() + string(os.PathSeparator) + v.Name())
				if err != nil {
					fmt.Println(err.Error())
				}
				err = json.Unmarshal(by, &ii) // 解析
				if err != nil {
					fmt.Println(err.Error())
				}
				// 针对当前索引是否有效进行判断
				if ii.IsEffective {
					dirNameInt64, err := strconv.ParseInt(dirInfo.Name(), 10, 64)
					if err != nil {
						fmt.Println(err.Error())
						break
					}
					IndexArr = append(IndexArr, dirNameInt64) // 添加到数组中
				}
			}
		}
	}
	// 对当前有用的索引进行排序
	var Max int64 = IndexArr[0]
	// 对当前有用的索引进行取出最大的
	for _, v := range IndexArr {
		if Max < v {
			Max = v
		}
	}
	// 拿到当前最新且有效的索引路径
	var MaxStr = strconv.FormatInt(Max, 10) // 处理成字符串类型
	var ValidIndexPath string = *dataDir + string(os.PathSeparator) + MaxStr // 拼成全路径
	index, err := bleve.Open(ValidIndexPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	ci = CurrentIndexStruct {
		index,
		MaxStr,
	}
	return ci, err
}

// 构建索引信息的字段
type IndexInfo struct {
	Id           string   `json:"id"`
	Path         string   `json:"path"`
	DocCount     int64    `json:"docCount"`
	TotalTime    int64    `json:"totalTime"`
	IndexSize    int64    `json:"indexSize"`
	IsEffective  bool     `json:"isEffective"`
	IsCreating   bool     `json:"isCreating"`
	CurrentIndex bool     `json:"currentIndex"`
}

// 列出所有索引文件夹
func ListIndex() (list []IndexInfo, err error) {
	list = make([]IndexInfo, 0) // 初始化list
	dirEntries, err := ioutil.ReadDir(*dataDir)
	if err != nil {
		log.Fatalf("error reading data dir: %v", err)
	}
	dirLenth := len(dirEntries) // 列出索引文件夹中的文件数量
	if (dirLenth == 0) {
		return list, nil
	}
	var ii IndexInfo // 用于序列化数据使用的
	for _, dirInfo := range dirEntries {
		indexPath := *dataDir + string(os.PathSeparator) + dirInfo.Name()
		// 跳过文件类型
		if !dirInfo.IsDir() {
			log.Printf("not registering %s, skipping", indexPath)
			continue
		}
		// 循环每个文件夹中的文件读取出相应的索引配置文件信息
		currentDir, _ := ioutil.ReadDir(*dataDir + string(os.PathSeparator) + dirInfo.Name()) // 此处不会报错，省略
		// 找到本地文件夹
		for _, v := range currentDir {
			// 对成功写入配置信息文件的，解析出对象，然后进行
			if v.Name() == "index_info.json" {
				// 读取出当前index_info.json文件
				by, err := ioutil.ReadFile(*dataDir + string(os.PathSeparator) + dirInfo.Name() + string(os.PathSeparator) + v.Name())
				if err != nil {
					fmt.Println(err.Error())
				}
				err = json.Unmarshal(by, &ii) // 解析
				if err != nil {
					fmt.Println(err.Error())
				}
				// 针对当前索引的判断和赋值操作：
				if ii.Id == CurrentIndex.Id {
					ii.CurrentIndex = true
				}
				list = append(list, ii)
			}
		}
	}
	// 最后对list中的索引进行时间上的排序 冒泡排序
	listLen := len(list)
	for i := 0; i < listLen; i++ {
		for j := i + 1; j < listLen; j++ {
			i_Id, _ := strconv.Atoi(list[i].Id)
			j_Id, _ := strconv.Atoi(list[j].Id)
			if i_Id < j_Id {
				temp := list[i]
				list[i] = list[j]
				list[j] = temp
			}
		}
	}
	return list, err
}

// 列出一条索引的详情
func ListOneIndex(id string) (ii IndexInfo, err error) {
	// 查找纪录索引详情的文件
	InfoFileBy, err := ioutil.ReadFile(*dataDir + string(os.PathSeparator) + id + string(os.PathSeparator) + "index_info.json")
	if err != nil {
		log.Fatalf("读取所查找的索引纪录文件index_info.json失败：", err)
	}
	err = json.Unmarshal(InfoFileBy, &ii) // 解析
	if err != nil {
		fmt.Println(err.Error())
	}
	return ii, err
}

// 删除索引，目前按名称删除
func DeleteIndex(id string) (info string) {
	dirEntries, err := ioutil.ReadDir(*dataDir) // 读出所有文件夹

	if err != nil {
		log.Fatalf("error reading data dir: %v", err)
		info = err.Error()
		return
	}

	dirLenth := len(dirEntries) // 列出索引文件夹中的文件数量

	// 如果没有文件夹，那么返回空
	if (dirLenth == 0) {
		info = "0"
		fmt.Println("暂无索引")
		return info
	}

	suc := false // 定义一个标识

	for _, dirInfo := range dirEntries {
		indexPath := *dataDir + string(os.PathSeparator) + dirInfo.Name()

		// 跳过非文件夹类型
		if !dirInfo.IsDir() {
			continue
		}

		if id == dirInfo.Name() {
			err = nil
			err = os.RemoveAll(indexPath) // 删除该文件
			if err != nil {
				fmt.Println("删除失败：", err.Error())
			} else {
				suc = true
				fmt.Println("删除成功！")
			}
			break
		}
	}

	if suc == true {
		info = "1"
	} else {
		info = "2"
	}
	return info
}

/* ------------------------ 封装函数 -------------------------- */

// 用于写入文件的方法
func WriteInfo(id string, path string, docCount int64, totalTime int64, indexSize int64, isEffective bool, isCreating bool, currentIndex bool) (err error) {
	// 构建info对象
	info := IndexInfo{
		id,
		path,
		docCount,
		totalTime,
		indexSize,
		isEffective,
		isCreating,
		currentIndex,
	}
	InfoText, err := json.Marshal(info) // 将json对象解析成json文本
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	// 将文件写入
	writePath := path + string(os.PathSeparator) + "index_info.json" // 写入的路径
	err = ioutil.WriteFile(writePath, InfoText, 0644)                // 0644 具有读写权限
	if err != nil {
		fmt.Println("信息写入索引目录失败：", err.Error())
		return err
	}
	return err
}

// 用于统计文件夹目录大小的方法
func GetDirSize(indexId string) (size int64, err error) {
	dirEntries, err := ioutil.ReadDir(*dataDir + string(os.PathSeparator) + indexId) // 读出所有文件夹
	if err != nil {
		return size, err
	}
	for _, v := range dirEntries {
		size += v.Size() / 1000 / 1000
	}
	return size, nil
}

// 用于停止或者结束后将索引信息写入本地的方法 0 表示发
func WriteIndexInfo(startTime time.Time, timeRes string, path string, total int64, isEffective bool, isCreating bool, isCurrentIndex bool) {
	indexDuration := time.Since(startTime)
	indexDurationSeconds := int64(indexDuration) / int64(time.Second) // 动态获取构建时间，单位为秒

	// 统计现有索引大小
	size, err := GetDirSize(timeRes)
	if err != nil {
		fmt.Println(err.Error())
	}

	WriteInfo(timeRes, path, total, indexDurationSeconds, size, isEffective, isCreating, isCurrentIndex) // 重新写一次
}
