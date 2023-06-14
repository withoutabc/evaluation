package ssh

import "golang.org/x/crypto/ssh"

func InitSsh(user, password, addr string) *ssh.Client {
	// SSH配置
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			// 使用密码进行认证
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 禁用SSH主机密钥检查
	}
	// 连接SSH服务器
	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		panic(err)
	}
	return client
}
