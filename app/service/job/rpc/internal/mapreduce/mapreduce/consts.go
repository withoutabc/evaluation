package mapreduce

const (
	url = "http://hadoop103:8088/ws/v1/cluster/apps"

	streamingJarPath     = " /opt/module/hadoop-3.1.3/share/hadoop/tools/lib/hadoop-streaming-3.1.3.jar "
	WordCountMapper      = " -mapper /opt/exe/wordcountmapper "
	WordCountReducer     = " -reducer /opt/exe/wordcountreducer "
	WordCountMapperFile  = " -file /opt/exe/wordcountmapper "
	WordCountReducerFile = " -file /opt/exe/wordcountreducer "

	CollectMapper      = " -mapper /opt/exe/collectmapper "
	CollectReducer     = " -reducer /opt/exe/collectreducer "
	CollectMapperFile  = " -file /opt/exe/collectmapper "
	CollectReducerFile = " -file /opt/exe/collectreducer "

	TestWordCountMapper      = " -mapper /opt/exe/mapper "
	TestWordCountReducer     = " -reducer /opt/exe/reducer "
	TestWordCountMapperFile  = " -file /opt/exe/mapper "
	TestWordCountReducerFile = " -file /opt/exe/reducer "

	Reduce  = " -D mapreduce.job.reduces=2 "
	JobName = " -D mapreduce.job.name="
)
