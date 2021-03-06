package entities

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
	others "valuebetsmining/data/Others"
)

const (
	DefaultDirTestGoals  = "../test/Goals"
	DefaultNameTestGoals = "Testing"
)

func TestGoalsFunc(t *testing.T) {
	goals, err := NewGoal(1, 2)
	if err != nil {
		t.Error(err)
	}
	rands, err := others.RandomArr(0, 10, 10)
	for key, value := range rands {
		goals.Update(value, key)
	}
	if len(goals.GoalsReceived) <= 10 {
		t.Errorf("Error parsing goals length: %d", len(goals.GoalsReceived))
	}
}

//TestPreviousNGoalsOfAMatch ... Testing function PreviousNGoalsOfAMatch to visualize if it returns correct values
func TestPreviousNGoalsOfAMatch(t *testing.T) {
	goals, err := NewGoal(1, 2)
	if err != nil {
		t.Error(err)
	}
	randsLocal, err := others.RandomArr(0, 10, 10)
	if err != nil {
		t.Error(err)
	}
	randsAway, err := others.RandomArr(0, 10, 10)
	if err != nil {
		t.Error(err)
	}
	goalsArr := []Goal{goals}
	for i := 0; i < len(randsLocal); i++ {
		err = goals.Update(randsLocal[i], randsAway[i])
		if err != nil {
			t.Error(err)
			break
		}
		goalsArr = append(goalsArr, goals)
	}
	for i := 1; i < len(goals.GoalsReceived); i++ {
		goalsPrevious, err := goals.PreviousNGoalsOfAMatch(i)
		if err != nil {
			t.Error(err)
			break
		}
		if !goalsPrevious.CompareOfGoals(goalsArr[len(goalsArr)-i]) {
			t.Errorf("Time:%d\nWant->\n%s\nGot->\n%s\n", i, goalsArr[len(goalsArr)-i].String(), goalsPrevious.String())
			break
		}
	}
}

func TestNewGoals(t *testing.T) {
	g, err := NewGoal(0, 1)
	if err != nil {
		t.Error(err)
	}
	randGoalsLocal, err := others.RandomArr(0, 5, 100)
	if err != nil {
		t.Error(err)
	}
	randsGoalsAway, err := others.RandomArr(0, 5, 100)
	if err != nil {
		t.Error(err)
	}
	data := make([][]string, 0, 101)
	data = append(data, strings.Split(fmt.Sprintf("%d,%d,%d,%d,%d,%s", -1, 0, 1, g.GoalsTuckedAmount, g.GoalsReceivedAmount, g.StringCSV()), ","))
	for i := 0; i < len(randGoalsLocal); i++ {
		g.Update(randGoalsLocal[i], randsGoalsAway[i])
		data = append(data, strings.Split(fmt.Sprintf("%d,%d,%d,%d,%d,%s", i, randGoalsLocal[i], randsGoalsAway[i], g.GoalsTuckedAmount, g.GoalsReceivedAmount, g.StringCSV()), ","))
	}
	file, err := os.Create(fmt.Sprintf("%s/%s%s_%d_%d%s", DefaultDirTestGoals, DefaultNameTestGoals, others.FuncName(), time.Now().Hour(), time.Now().Minute(), DefaultExtensionCSV))
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Index", "goals local", "goals away", "GoalsTuckedAmount", "GoalsReceivedAmount", "Average tucked goals local", "average received goals local"})

	writer.WriteAll(data)
}

func TestNewPreviousGoals(t *testing.T) {
	g, err := NewGoal(0, 1)
	if err != nil {
		t.Error(err)
	}
	randGoalsLocal, err := others.RandomArr(0, 5, 100)
	if err != nil {
		t.Error(err)
	}
	randsGoalsAway, err := others.RandomArr(0, 5, 100)
	if err != nil {
		t.Error(err)
	}
	data := make([][]string, 0, 101)
	data = append(data, strings.Split(fmt.Sprintf("%d,%d,%d,%s,%s,%s", -1, 0, 1, "nil", "nil", "nil"), ","))
	for i := 0; i < len(randGoalsLocal); i++ {
		err = g.Update(randGoalsLocal[i], randsGoalsAway[i])
		if err != nil {
			t.Error(err)
		}
		goal, err := g.PreviousNGoalsOfAMatch(1)
		if err != nil {
			t.Error(err)
		}
		data = append(data, strings.Split(fmt.Sprintf("%d,%d,%d,%d,%d,%s", i, randGoalsLocal[i], randsGoalsAway[i], goal.GoalsTuckedAmount, goal.GoalsReceivedAmount, goal.StringCSV()), ","))
	}
	file, err := os.Create(fmt.Sprintf("%s/%s%s_%d_%d%s", DefaultDirTestGoals, DefaultNameTestGoals, others.FuncName(), time.Now().Hour(), time.Now().Minute(), DefaultExtensionCSV))
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Index", "goals local", "goals away", "GoalsTuckedAmount", "GoalsReceivedAmount", "Average tucked goals local", "average received goals local"})

	writer.WriteAll(data)
}
