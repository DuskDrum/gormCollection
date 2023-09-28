-[x] gorm/gen 自动生成代码
-[x] 勾子函数&&自定义勾子函数
-[x] 自定义Value和Parse， BigDecimal
-[x] gorm/go-sqlmock
-[x] gorm.io/plugin/dbresolver  // 读写分离
-[x] gorm.io/plugin/soft_delete
-[ ] gorm.io/datatypes
-[ ] 新增、Upsert的源码
-[ ] 怎么自定义tag，自定义notBlank、notNull
-[ ] 事务，BeginTransaction
-[ ] Session与分页插件。会话模式：特殊的配置
-[ ] gorm.io/hints，限制索引






```go
type DB struct {
	// 等待连接的总时间
	waitDuration atomic.Int64
    // 连接器，用于创建和管理底层的数据库连接
	connector driver.Connector
	// 已关闭的连接数量
	numClosed atomic.Uint64

	mu           sync.Mutex    
    // 表示连接池中的空闲连接列表。当连接不再使用时，会被放入空闲连接列表中，以便重用。
	freeConn     []*driverConn 
	// 表示连接请求的映射，用于存储等待连接的请求。
	connRequests map[uint64]chan connRequest
	// 表示下一个连接请求的键值。用于生成连接请求的唯一标识。
	nextRequest  uint64
	// 表示当前打开的连接数，即连接池中当前正在使用的连接数量
	numOpen      int  
	// 表示连接打开器的通道。当连接池需要打开新的连接时，会从该通道获取连接打开器，并使用打开器创建新的连接。
	openerCh          chan struct{}
	// 表示连接池是否已关闭的标志。如果该属性为true，则连接池已关闭，不再接受新的连接请求。
	closed            bool
	// 表示连接池的依赖关系图，用于跟踪连接的依赖关系。当连接关闭时，将根据依赖关系图关闭相关的依赖连接。
	dep               map[finalCloser]depSet
	// 表示最近放入连接池的连接
	lastPut           map[*driverConn]string 
	// 定义0或者负数的话默认为2。表示连接池中最大的空闲连接数。连接池会维护一定数量的空闲连接，以便在需要时能够快速提供可用的连接。
	maxIdleCount      int   
	// 表示连接池中最大的打开连接数。该属性指定连接池允许同时打开的最大连接数。当达到最大连接数时，进一步的连接请求将被阻塞等待。
	maxOpen           int
	// 表示连接的最大复用时间。连接在使用一段时间后可能会被关闭并重新创建，以确保连接的可靠性和性能。
	maxLifetime       time.Duration         
	// 表示连接的最大空闲时间。当连接处于空闲状态并且空闲时间超过 maxIdleTime 时，连接可能会被关闭并从连接池中移除。
	maxIdleTime       time.Duration
	// 表示连接池的清理器通道。当连接池需要清理空闲连接时，会向该通道发送清理信号，触发空闲连接的清理操作。
	cleanerCh         chan struct{}
	// 表示当前等待连接的请求数量。当连接池中的连接已达到最大限制时，新的连接请求将被放入等待队列，并递增 waitCount。
	waitCount         int64 
	// 因为连接空闲时间过长而被关闭的连接数量
	maxIdleClosed     int64 
	// 因为连接空闲时间超过最大空闲时间而被关闭的连接数量
	maxIdleTimeClosed int64 
	// 因为连接的最大复用时间到达而被关闭的连接数量
	maxLifetimeClosed int64

	stop func() // stop cancels the connection opener.
}
```

gorm的结构体tag放在了"gorm-->schema下面的field.go。关于外键约束或者多对多约束的tag放在的是gorm--->schema下面的relationship.go


预加载，有关联的struct可以用Preload实现预加载，Joins也可以预加载