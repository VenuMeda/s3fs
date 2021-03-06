package command

// Mock Outputter

type mockOutputter struct {
	output []string
}

func (m *mockOutputter) Write(str string) {
	m.output = append(m.output, str)
}

// Mock S3Client

type mockS3Client struct {
	lsBucketsCallback func() ([]string, error)
	lsObjectsCallback func(string, string) ([]string, error)

	bucketExistsCallback func(string) (bool, error)
	objectExistsCallback func(string, string) (bool, error)
}

func (m mockS3Client) LsBuckets() ([]string, error) {
	return m.lsBucketsCallback()
}

func (m mockS3Client) LsObjects(bucket, prefix string) ([]string, error) {
	return m.lsObjectsCallback(bucket, prefix)
}

func (m mockS3Client) BucketExists(bucket string) (bool, error) {
	return m.bucketExistsCallback(bucket)
}

func (m mockS3Client) ObjectExists(bucket, key string) (bool, error) {
	return m.objectExistsCallback(bucket, key)
}
