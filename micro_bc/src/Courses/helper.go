package Courses

import "context"

func NewCourseModel(id int32, name string) *CourseModel {
	return &CourseModel{CourseId: id, CourseName: name}
}

type CourseServiceImpl struct {
}

func NewCourseServiceImpl() *CourseServiceImpl {
	return &CourseServiceImpl{}
}

func (c CourseServiceImpl) ListForTop(ctx context.Context, request *CourseLstRequest, response *CourseLstResponse) error {
	ret := make([]*CourseModel, 0)
	ret = append(ret, NewCourseModel(101, "golang课程"), NewCourseModel(102, "java课程"))
	response.Result = ret
	return nil
}

func (c CourseServiceImpl) GetDetail(ctx context.Context, request *DetailRequest, response *DetailResponse) error {
	ret := NewCourseModel(101, "golang课程")
	response.Course = ret
	return nil
}
