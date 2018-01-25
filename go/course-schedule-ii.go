package main

type CourseDependencies struct {
	DependenciesNumber int
	RequiredBy         []int
	Visited            bool
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	order := make([]int, 0)
	if numCourses == 0 {
		return order
	}
	courses := createGraph(numCourses, prerequisites)

	readyCources := make([]int, 0)

	for len(order) < numCourses {
		courseNumber := prepareNextCourse(&readyCources, courses, numCourses)
		if courseNumber < 0 {
			break
		}

		addCourseToOrder(courseNumber, &readyCources, &courses, &order)
	}

	if len(order) < numCourses {
		return []int{}
	}

	return order
}

func addCourseToOrder(courseNumber int, readyCources *[]int, courses *[]CourseDependencies, order *[]int) {
	(*courses)[courseNumber].Visited = true
	*order = append(*order, courseNumber)
	for i := range (*courses)[courseNumber].RequiredBy {
		candidate := (*courses)[courseNumber].RequiredBy[i]
		(*courses)[candidate].DependenciesNumber--
		if courseIsReady(candidate, *courses) {
			*readyCources = append(*readyCources, candidate)
		}
	}
}

func prepareNextCourse(readyCourses *[]int, courses []CourseDependencies, numCourses int) int {
	if len(*readyCourses) == 0 {
		for i := 0; i < numCourses; i++ {
			if courseIsReady(i, courses) {
				*readyCourses = append(*readyCourses, i)
				break
			}
		}
	}
	if len(*readyCourses) == 0 {
		return -1
	}

	courseNumber := (*readyCourses)[0]
	*readyCourses = (*readyCourses)[1:]
	return courseNumber
}

func courseIsReady(courseNumber int, courses []CourseDependencies) bool {
	if courses[courseNumber].DependenciesNumber == 0 && !courses[courseNumber].Visited {
		return true
	}
	return false
}

func createGraph(n int, prerequisites [][]int) []CourseDependencies {
	courses := make([]CourseDependencies, n)
	for i := 0; i < len(prerequisites); i++ {
		previousCourse := prerequisites[i][1]
		currentCourse := prerequisites[i][0]
		courses[previousCourse].RequiredBy = append(courses[previousCourse].RequiredBy, currentCourse)
		courses[currentCourse].DependenciesNumber++
	}
	return courses
}
