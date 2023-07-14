package di

import "testing"

type A struct {
	B *B
}

func NewA(b *B) *A {
	return &A{B: b}
}

type B struct {
	C *C
}

func NewB(c *C) *B {
	return &B{C: c}
}

type C struct {
	Num int
}

func NewC() *C {
	return &C{Num: 1}
}

func TestDI(t *testing.T) {
	container := New()

	if err := container.Provide(NewA); err != nil {
		t.Fatal(err)
	}

	if err := container.Provide(NewB); err != nil {
		t.Fatal(err)
	}

	if err := container.Provide(NewC); err != nil {
		t.Fatal(err)
	}

	if err := container.Invoke(func(a *A) {
		if a.B.C.Num != 1 {
			t.Fatal("error")
		}
	}); err != nil {
		t.Fatal(err)
	}

}

//draw a picture of the the container and the dependencies providers and invocke
// container
// 	|_ providers
//		|_ NewA
//			|_ NewB
//				|_ NewC
//	|_ invocke
//		|_ func(a *A) {
//			|_ a.B.C.Num
//		}
//	|_ result
//		|_ 1
