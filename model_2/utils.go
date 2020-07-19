package model_2

type Student struct {
	Name string
	Score float64   // 小写怎么办 定义一个方法来访问
}
type Student_2 struct {
	Name string
	score float64   // 小写怎么办 定义一个方法来访问
}
func (s Student_2) Getscore() float64{
	return s.score
}
func (s *Student_2) SetCore(num float64){ // 方法来设值
	(*s).score = num
}

func NewStudent(n string, s float64) *Student { // 返回一个指针，指向了Student{}的地址
	return &Student{
		Name : n,
		Score : s,
	}
}