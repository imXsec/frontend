package main

import (
	"time"
	"nano/framework"
	"log"
)

type logCommandType int

const (
	cmdQueryLog = iota
	cmdAddLog
	cmdRemoveLog
)

const (
	TimeFormatLayout = "2006-01-02 15:04:05"
)

type LogManager struct {
	agent    *LogAgent
	commands chan logCommand
	runner   *framework.SimpleRunner
}

type LogQueryCondition struct {
	Limit     int
	Start     int
	BeginTime time.Time
	EndTime   time.Time
}

type LogEntry struct {
	ID      string
	Time    time.Time
	Content string
}

type LogResult struct {
	Error error
	Logs  []LogEntry
}

type logCommand struct {
	Type       logCommandType
	Condition  LogQueryCondition
	Content    string
	IDList     []string
	ResultChan chan LogResult
	ErrorChan  chan error
}

func CreateLogManager(dataPath string) (manager *LogManager, err error) {
	const (
		DefaultQueueLength = 1 << 10
	)
	manager = &LogManager{}
	manager.commands = make(chan logCommand, DefaultQueueLength)
	manager.runner = framework.CreateSimpleRunner(manager.Routine)
	manager.agent, err = CreateLogAgent(dataPath)
	//if err != nil{
	//	return
	//}
	return
}

func (manager *LogManager) Start() error{
	return manager.runner.Start()
}

func (manager *LogManager) Stop() error{
	return manager.runner.Stop()
}

func (manager *LogManager) Routine(c framework.RoutineController){
	log.Println("<log> started")
	for !c.IsStopping() {
		select {
		case <-c.GetNotifyChannel():
			c.SetStopping()
		case cmd := <-manager.commands:
			manager.handleCommand(cmd)
		}
	}
	c.NotifyExit()
	log.Println("<log> stopped")
}

func (manager *LogManager) QueryLog(condition LogQueryCondition, respChan chan LogResult)  {
	manager.commands <- logCommand{Type:cmdQueryLog, Condition:condition, ResultChan:respChan}
}

func (manager *LogManager) AddLog(content string, respChan chan error)  {
	manager.commands <- logCommand{Type:cmdAddLog, Content:content, ErrorChan:respChan}
}

func (manager *LogManager) RemoveLog(entries []string, respChan chan error)  {
	manager.commands <- logCommand{Type:cmdRemoveLog, IDList:entries, ErrorChan:respChan}
}

func (manager *LogManager) handleCommand(cmd logCommand){
	var err error
	switch cmd.Type {
	case cmdQueryLog:
		err = manager.handleQueryLog(cmd.Condition, cmd.ResultChan)
	case cmdAddLog:
		err = manager.handleAddLog(cmd.Content, cmd.ErrorChan)
	case cmdRemoveLog:
		err = manager.handleRemoveLog(cmd.IDList, cmd.ErrorChan)
	default:
		log.Printf("<log> unsupport command type %d", cmd.Type)
	}
	if err != nil{
		log.Printf("<log> handle command %d fail: %s", cmd.Type, err.Error())
	}
}

func (manager *LogManager) handleQueryLog(condition LogQueryCondition, respChan chan LogResult) (err error) {
	logs, err := manager.agent.Query(condition)
	if err != nil{
		respChan <- LogResult{Error:err}
		return
	}
	respChan <- LogResult{Logs:logs}
	return nil
}

func (manager *LogManager) handleAddLog(content string, respChan chan error) (err error) {
	err = manager.agent.Write(content)
	respChan <- err
	return
}

func (manager *LogManager) handleRemoveLog(entries []string, respChan chan error) (err error) {
	err = manager.agent.Remove(entries)
	respChan <- err
	return
}
