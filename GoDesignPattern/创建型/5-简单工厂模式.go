/*
简单工厂模式：
优点：
1.实现了对象创建和使用的分离
2.不需要记住具体类名，记住参数即可，减少使用者记忆量
缺点：
1.对工厂类职责过重，一旦不能工作，系统受到影响
2.增加系统中类的个数，复杂度和理解度增加
3.违反"开闭原则"，添加新产品需要修改工厂逻辑，工厂越来越复杂

适用场景：
1.工厂类负责创建的对象比较少，由于创建的对象较少，不会造成工厂方法
中的业务逻辑太过复杂
2. 客户端只知道传入工厂类的参数，对于如何创建对象并不关心
*/
package pkg1

import "fmt"

// 抽象层
type Fruit interface {
	Show()
}

// 实现层
type Apple struct{}

func (apple *Apple) Show() {
	fmt.Println("apple!")
}

type Banana struct{}

func (banana *Banana) Show() {
	fmt.Println("banana!")
}

type Pear struct{}

func (pear *Pear) Show() {
	fmt.Println("pear!")
}

// 工厂模块
type Factory struct{}

func (factory *Factory) CreateFruit(name string) Fruit {
	var fruit Fruit
	if name == "apple" {
		fruit = new(Apple)
	} else if name == "banana" {
		fruit = new(Banana)
	} else if name == "pear" {
		fruit = new(Pear)
	}

	return fruit
}

// 业务逻辑层
func main5() {
	factory := new(Factory)
	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()
}
