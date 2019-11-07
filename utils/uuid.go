package utils

const (
	MAXUINT32              = 4294967295
	DEFAULT_UUID_CNT_CACHE = 512
)

type UUIDGenerator struct {
	idGen        uint32
	internalChan chan uint32
}

func NewUUIDGenerator() *UUIDGenerator {
	gen := &UUIDGenerator{
		idGen:        0,
		internalChan: make(chan uint32, DEFAULT_UUID_CNT_CACHE),
	}
	gen.startGen()
	return gen
}

//开启 goroutine, 把生成的数字形式的UUID放入缓冲管道
func (this *UUIDGenerator) startGen() {
	go func() {
		for {
			if this.idGen == MAXUINT32 {
				this.idGen = 1
			} else {
				this.idGen += 1
			}
			this.internalChan <- this.idGen
		}
	}()
}

//获取uint32形式的UUID
func (this *UUIDGenerator) GetUint32() uint32 {
	return <-this.internalChan
}
