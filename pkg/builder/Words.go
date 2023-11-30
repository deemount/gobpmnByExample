package factory

var (

	// type
	typeBuilder = "Builder"

	// anonymous classified (plural)
	fieldEventElements = "EventElements"
	fieldEvents        = "Events"
	fieldProcesses     = "Processes"
	fieldTasks         = "Tasks"

	// field
	fieldID      = "ID"
	fieldProcess = "Process"

	// field event elements
	fieldEndEvent   = "EndEvent"
	fieldStartEvent = "StartEvent"

	// tasks
	fieldTask = "Task"

	// unknown
	fieldUnknown = "Unknown"

	// bool
	boolIsExecutable = "IsExecutable"

	// methods (actually used in Builder.build only)
	MethodSetProcess = "SetProcess"
	MethodSetDiagram = "SetDiagram"
)
