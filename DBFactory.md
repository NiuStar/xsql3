DBFactory函数列表



| 函数名              | 传参                           | 说明                            |
| ------------------- | ------------------------------ | ------------------------------- |
| NewDBFactory        |                                | 创建一个数据库操作对象          |
| Limit               | start,end                      | 结果集分页                      |
| Count               |                                | 根据条件判断条数                |
| SetTable            | tables... Type.DBOperation     | 设置查询表对象                  |
| AddTable            | tables... Type.DBOperation     | 添加查询表对象                  |
| SetFields           | fields... Type.TableType       | 设置查询结果对象                |
| AddFields           | fields... Type.TableType       | 添加查询结果对象                |
| SetConditions       | conditions... *Type.Condition  | 设置sql语句的where的AND条件     |
| AddCondition        | conditions... *Type.Condition  | 添加sql语句的where的AND条件     |
| SetORConditions     | conditions... *Type.Condition  | 设置sql语句的where的OR条件      |
| AddORCondition      | conditions... *Type.Condition  | 添加sql语句的where的OR条件      |
| DESC                |                                | 设置排序方式为倒序              |
| SetOrderBy          | orderbys... Type.TableType     | 设置排序                        |
| SetGroupBy          | orderbys... Type.TableType     | 设置分组                        |
| SetJoins            | joins... *Join                 | 设置链表                        |
| AddJoin             | joins... *Join                 | 添加链表                        |
| GetResults          |                                | 获取查询结果                    |
| GetResultsOperation |                                | 获取转换后的查询结果            |
| InsertDB            | list... Type.DBOperation       | 把list内容插入数据库            |
| DeleteDB            | obj Type.DBOperation           | 根据where条件删除对象表内的内容 |
| DeleteDBALL         | tableName string               | 删除表对象所有内容              |
| UpdateDB            | list... Type.DBOperation       | 根据where条件内容更新list内容   |
| ParseResults        | datas []map[string]interface{} | 把数组对象转换为表对象          |

