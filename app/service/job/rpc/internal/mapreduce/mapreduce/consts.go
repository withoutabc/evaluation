package mapreduce

const (
	url = "http://hadoop103:8088/ws/v1/cluster/apps"

	streamingJarPath     = " /opt/module/hadoop-3.1.3/share/hadoop/tools/lib/hadoop-streaming-3.1.3.jar "
	WordCountMapper      = " -mapper /opt/exe/mapper "
	WordCountReducer     = " -reducer /opt/exe/reducer "
	WordCountMapperFile  = " -file /opt/exe/mapper "
	WordCountReducerFile = " -file /opt/exe/reducer "
)
