package main

import "fmt"

type Next struct {
    ID          int
    Plan        string
    Description string
    Status      bool
}

type NextList struct {
    Plans []Next
}

func (t1 *NextList) addTask(newPlan Next) {
    t1.Plans = append(t1.Plans, newPlan)
}

func (t1 *NextList) getTasks() []Next {
    return t1.Plans
}

func (t1 *NextList) markPlanAsDone(planID int) error {
    for i, plan := range t1.Plans {
        if plan.ID == planID {
            t1.Plans[i].Status = true
            return nil
        }
    }
    return fmt.Errorf("task with ID %d not found", planID)
}

func (t1 *NextList) deletePlan(planID int) error {
    for i, plan := range t1.Plans {
        if plan.ID == planID {
            t1.Plans = append(t1.Plans[:i], t1.Plans[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("task with ID %d not found", planID)
}

func main() {
    NextList := NextList{}

    NextList.addTask(Next{ID: 1, Plan: "Go to the gym", Description: "I want to start going to the gym every day for 30 minutes", Status: true})
    NextList.addTask(Next{ID: 2, Plan: "Go to the market", Description: "Every day from 4 pm I will be going to the market for new foodstuffs", Status: false})

    plans := NextList.getTasks()
    fmt.Println("Tasks:", plans)

    err := NextList.markPlanAsDone(1)
    if err != nil {
        fmt.Println("Error:", err)
    }

    fmt.Println("Tasks after marking task 2 as done:", NextList.getTasks())
}
