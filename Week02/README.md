学习笔记

问题：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

回答：在dao层进行数据库操作，获取数据或者error，error不建议直接进行默认值或者降级处理，直接将error上抛上一层。上一层service层调用dao层的函数时，如果下层抛上来一个error，可以通过wrapf对之前的error增加新的信息，最终层层上抛，最后在最上层打印log或者将stack信息完全打印出来用以追踪问题。
