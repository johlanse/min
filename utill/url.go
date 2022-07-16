package utill

var (
	data = map[int]string{
		// 慕课平台
		0: "https://mooc.cdcas.com",
		// 实训平台
		1: "https://shixun.cdcas.com",
		// 创能实训
		2: "https://cdcas.chuangnengkeji.com",
		// 频蓝实训
		3: "https://cdcas.pinlankeji.com",
		// 戴希实训
		4: "https://cdcas.daixiwangluo.com/",
	}
)

func GetBaseUrl(index int) string {
	return data[index]
}
