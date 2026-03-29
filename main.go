package main

import "fmt"

func main() {
	// 制作流程非常简单灵活：只要实现 Coffee 接口，就可以无限叠加"配料"

	// 1. 制作一杯：燕麦奶 + 标准浓缩
	var order1 Coffee = &SimpleCoffee{}
	order1 = &StandardEspresso{Coffee: order1}  // 第1步：加标准浓缩
	order1 = &OatMilk{Coffee: order1}           // 第2步：加燕麦奶

	fmt.Println("订单 1 (燕麦拿铁)：")
	fmt.Printf("描述: %s\n", order1.Description())
	fmt.Printf("价格: $%.2f\n\n", order1.Cost())

	// 2. 制作一杯：全脂奶 + 双份精华浓缩 (Double Ristretto)
	var order2 Coffee = &SimpleCoffee{}
	order2 = &Ristretto{Coffee: order2}         // 第1步：加精华浓缩
	order2 = &Ristretto{Coffee: order2}         // 第2步：再加一份精华浓缩
	order2 = &WholeMilk{Coffee: order2}         // 第3步：加全脂奶

	fmt.Println("订单 2 (双倍精华浓缩拿铁)：")
	fmt.Printf("描述: %s\n", order2.Description())
	fmt.Printf("价格: $%.2f\n", order2.Cost())
}
