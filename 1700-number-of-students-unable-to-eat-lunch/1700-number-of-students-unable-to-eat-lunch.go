func countStudents(students []int, sandwiches []int) int {
    n := len(students)
    pref := [2]int{0,0}

    for i:=0;i<n;i++{
        pref[students[i]]++
    }

    for len(sandwiches) > 0 && pref[sandwiches[0]] > 0 {
        for len(students) > 0 && students[0] == sandwiches[0] {
            pref[students[0]]--
            students = students[1:]
            sandwiches = sandwiches[1:]
        }

        for len(students) > 0 && students[0] != sandwiches[0] && pref[sandwiches[0]] > 0 {
            topStudent := students[0]
            students = students[1:]
            students = append(students, topStudent)
        }

    }

    return len(students)
}