# etcd 和 Raft 实现

## 目录

#### Raft 介绍

节点的三种状态:
- Follower: 追随者
- Candidate: 候选者
- Leader: 领导者

流程:
- 所有节点一开始的状态都是追随者状态
- 如果这些处于追随者状态的节点没有收到领导者的信息. 这些节点将改变状态成为候选者
