package single_responsibility

import (
	"testing"
	"designer_go/principle/single_responsibility/bad"
	"designer_go/principle/single_responsibility/good"
)

func Test_SimpleResponsibility_Bad(t *testing.T) {
	// 调用函数
	fnTestBadCourse := func(bc *bad.BadCourse) {
		bc.Play()
		bc.Pause()
		bc.Forward(30)
		bc.Backward(30)
	}

	fnTestBadCourse(bad.NewBadCourse(1, "直播课"))
	fnTestBadCourse(bad.NewBadCourse(2, "录播课"))
}

func Test_SimpleResponsibility_Good(t *testing.T) {
	fnTestGoodCourse := func(gc good.IGoodCourse) {
		pc := gc.Controller()
		pc.Play()
		pc.Pause()
		// 判定是否IReplayControl类型
		if rc, ok := pc.(good.IReplayControl); ok{
			rc.Forward(30)
			rc.Backward(30)
		}
	}

	fnTestGoodCourse(good.NewLiveCourse(11, "直播课"))
	fnTestGoodCourse(good.NewReplayCourse(12, "录播课"))
}
