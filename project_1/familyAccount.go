package ulits

import(
	"fmt"
	"time"
)

type FamilyAccount struct {
	key string // 接受用户输入
	loop bool // 判断是否退出
	// 定义账户余额
	now_money float64
	// 收支的金钱
	get_money float64
	// 收支说明
	note string
	// 记录是否有收支行为
	flag bool
	// 收支详情 使用字符串来记录
	details string  // 有收支时拼接字符串
}

func NewFamilyAccount() *FamilyAccount { // 工厂模式
	return &FamilyAccount{
		key:"",
		loop:true,
		now_money:10000.0,
		get_money:0.0,
		note:"",
		flag:false,
		details:"收支\t账户金额\t收支金额\t说明\n",
	}
}
// 1.收支明细改成一个功能
func (this *FamilyAccount)showDetails() {
	fmt.Println("----------------收支明细-----------------")
	if this.flag{fmt.Println(this.details)}else{fmt.Println("当前没有收支明细，来一笔吧")}
}
// 2.登记收入改成一个功能
func (this *FamilyAccount)income(){
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.get_money)
	this.now_money += this.get_money
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.note)
	// 将收入情况拼接到details变量
	this.details += fmt.Sprintf("收入\t%v\t%v\t%v\n", this.now_money,this.get_money,this.note)
	this.flag = true
}

// 3.登记支出改成一个功能
func (this *FamilyAccount)pay() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.get_money)
	if this.now_money < this.get_money{
		fmt.Println("余额不足")
		time.Sleep(1000*time.Millisecond)
		// break
	}else{
		this.now_money -= this.get_money
		fmt.Println("本次支出说明:")
		fmt.Scanln(&this.note)
		// 将收入情况拼接到details变量
		this.details += fmt.Sprintf("支出\t%v\t%v\t%v\n", this.now_money,this.get_money,this.note)
		this.flag = true}
}
// 4. 退出功能
func (this *FamilyAccount)exit(){
	fmt.Println("你确定退出吗？y/n")
	choose:=""
	fmt.Scanln(&choose)
	for{
		if choose=="y" || choose=="n"{
					break
				}else{
					fmt.Println("重新输入")
				}
			}
	if choose =="y"{
		this.loop = false
	}
}

// 显示主菜单
func (this *FamilyAccount) MainMenu(){
	for{
		fmt.Println("\n----------------家庭收支记账软件-----------------")
			fmt.Println("                1 收支明细")
			fmt.Println("                2 登记收入")
			fmt.Println("                3 登记支出")
			fmt.Println("                4 退出软件")
			fmt.Printf("请选择(1-4):")
	
			fmt.Scanln(&this.key)
	
	
			switch this.key{
			case "1":
				this.showDetails()
			case "2":
				this.income()
			case "3":
				this.pay()
			case "4":
				this.exit()
			default :
				fmt.Println("请输入正确的选项")
			}
			if !this.loop{
				break
			}
		}
}