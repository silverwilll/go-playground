package main

// SimpleCoffee 是基础的咖啡实现 (例如基础水底)
type SimpleCoffee struct{}

func (c *SimpleCoffee) Cost() float64 {
	return 2.0
}

func (c *SimpleCoffee) Description() string {
	return "Base Coffee"
}

// --- 以下为各种配料的实现（这里使用了经典的设计模式：装饰器模式） ---

// OatMilk 燕麦奶实现
type OatMilk struct {
	Coffee Coffee
}

func (m *OatMilk) Cost() float64 {
	return m.Coffee.Cost() + 0.8 // 燕麦奶相对贵一点
}

func (m *OatMilk) Description() string {
	return m.Coffee.Description() + ", Oat Milk"
}

// WholeMilk 全脂牛奶实现
type WholeMilk struct {
	Coffee Coffee
}

func (m *WholeMilk) Cost() float64 {
	return m.Coffee.Cost() + 0.5
}

func (m *WholeMilk) Description() string {
	return m.Coffee.Description() + ", Whole Milk"
}

// StandardEspresso 标准浓缩咖啡实现
type StandardEspresso struct {
	Coffee Coffee
}

func (e *StandardEspresso) Cost() float64 {
	return e.Coffee.Cost() + 1.0
}

func (e *StandardEspresso) Description() string {
	return e.Coffee.Description() + ", Standard Espresso"
}

// Ristretto 精华浓缩咖啡（另一种 Espresso 分类，水量少口感更醇厚）
type Ristretto struct {
	Coffee Coffee
}

func (r *Ristretto) Cost() float64 {
	return r.Coffee.Cost() + 1.2
}

func (r *Ristretto) Description() string {
	return r.Coffee.Description() + ", Ristretto Shot"
}
