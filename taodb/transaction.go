package taodb

/*
	事务
	一个事务的生命周期：
	创建 -> 写入 -> 提交
			-> 回滚 
	提交、回滚即死亡
*/
type Transaction struct {
	id uint32 // 事务id
	db *DB
	index TempIndex // 临时哈希表，当提交时，哈希表并入
	commited bool

}

func NewTransaction() *Transaction {
	TransactionMutex.Lock()
	defer TransactionMutex.Unlock()
	transaction := &Transaction{
		id: TransactionId,
	}
	TransactionId++
	return transaction
}

// 以下未完成

func (transaction *Transaction) Put(key, value string) error {
	newEntry = NewEntry(key, value, transaction.id, transaction.id, PUT)
	
	return nil
}

func (transaction *Transaction) Get(key string) (string, error) {
	
    return "", nil
}

func (transaction *Transaction) Delete(key string) error {
	return nil
}

func (transaction *Transaction) Exist(key string) (bool, error) {
    return false, nil
}

func (transaction *Transaction) Commit() error {

	return nil
}

func (transaction *Transaction) Rollback() error {
	return nil
}

func (transaction *Transaction) Close() error {
	if transaction.commited {
		return nil
	} else {
		// 尝试回滚
		return transaction.Rollback()
	}
}