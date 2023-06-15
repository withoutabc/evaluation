package mapreduce

const (
	url = "http://hadoop103:8088/ws/v1/cluster/apps"

	streamingJarPath     = " /opt/module/hadoop-3.1.3/share/hadoop/tools/lib/hadoop-streaming-3.1.3.jar "
	WordCountMapper      = " -mapper /opt/exe/wordcountmapper "
	WordCountReducer     = " -reducer /opt/exe/wordcountreducer "
	WordCountMapperFile  = " -file /opt/exe/wordcountmapper "
	WordCountReducerFile = " -file /opt/exe/wordcountreducer "

	Reduce  = " -D mapreduce.job.reduces=2 "
	JobName = " -D mapreduce.job.name="
)
