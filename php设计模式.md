# phpNote
1. ## 多态
```

//定义老虎抽象类
//老虎有多种，一种是XTiger,另一种是MTiger,都是climb方法

abstract class Tiger{
	public abstract function climb();
}

class XTiger extends Tiger{
	public function climb(){
		echo "爬到树顶</br>";
	}
}

class Cat{
	public function climb(){
		echo "在树上来去自如";
	}
}

class MTiger extends Tiger{
	public function climb(){
		echo "从树上摔下来</br>";
	}
}

class Client{
	public static function call(Tiger $animal){ //参数为Tiger类型，也可以不定义类型
		echo $animal->climb();
	}
}
Client::call(new XTiger());
Client::call(new MTiger());
//如果参数不加类型可以用Client::call(new Cat())

```

2. ## 简单工厂(面向接口开发)

```

//共同接口
interface db{
	function conn();
}



//服务端开发(不知道将会被谁调用)
class dbmysql implements db{
	public function conn(){
		echo "连接成功";
	}
}

class dbsqlite implements db{
	public function conn(){
		echo "连接上了sqlite";
	}
}


class Factory{
	public static function createDB($type){
		if($type == 'mysql'){
			return new dbmysql();
		}else if($type == 'sqlite'){
			return new dbsqlite();
		}else{
			throw new Exception('err db type',1);
		}
	}
}

//客户端，看不到dbmysql,dbsqlite内部及细节的
//只知道，上两个类实现了db接口
//简单工厂，然使用者知道的越少越好
//方法允许传递数据库名称

$mysql = Factory::createDB('mysql');
$mysql->conn();

$sqlite = Factory::createDB('sqlite');
$sqlite->conn();

//如果新增oracle类型，该怎么做
//服务端要修改Factory的内容
//在面向对象法则中，有重要的开闭原则---对于修改是封闭，对于扩展是开放的

```

3. ## 工厂方法

```

interface db{
	function conn();
}

interface Factory{
	function createDB();
}


//服务端开发(不知道将会被谁调用)
class dbmysql implements db {
	public function conn(){
		echo "连接成功</br>";
	}
}

class mysqlFactory implements Factory {
	public function createDB(){
		return new dbmysql();
	}
}

class dbsqlite implements db{
	public function conn(){
		echo "连接上了sqlite</br>";
	}
}


class sqliteFactory implements Factory{
	public function createDB(){
		return new dbsqlite();
	}
}

//++++++++++服务器端实现添加oracle类+++++++++++++

class oracle implements db{
	public function conn(){
		echo "连接上了oracle</br>";
	}
}

class oracleFactory implements Factory{
	public function createDB(){
		return new oracle();
	}
}

//+++++++++++客户端开发调用+++++++++++++++
$fact = new mysqlFactory();
$db = $fact->createDB();
$db->conn();

$fact = new sqliteFactory();
$db = $fact->createDB();
$db->conn();

$fac = new oracleFactory();
$db = $fac->createDB();
$db->conn();

//这样就比避免的修改源代码，只需扩展子类就ok

```

4. ##单例模式

```

单例模式

第一步普通类
class single{

}

$s1 = new single();
$s2 = new single();

if($s1 === $s2){
	echo "是同一个对象";
}else{
	echo "不是同一个对象"; //输出是一个对象
}

第二步 封锁new操作
class single{
	protected function __construct(){

	}
}
$s1 = new sigle();//这时构造函数已经被保护起来了，不能new


第三步 留个接口来new对象
class single{
	public static function getIns(){
		return new self();
	}
	protected function __construct(){

	}
}

$s1 = single::getIns();
$s2 = single::getIns();

if($s1 === $s2){
	echo "是同一个对象";
}else{
	echo "不是同一个对象";//输出不是一个对象
}

第四步,getIns先判断实例
+++++++++++++++++++继承后就既可以new多个++++++++++++++++++++
class single{

	protected static $ins = null;

	public static function getIns(){
		if(self::$ins === null){
			self::$ins == new self();
		}

		return self::$ins;
	}

	protected function __construct(){

	}
}

$s1 = single::getIns();
$s2 = single::getIns();

if($s1 === $s2){
	echo "是同一个对象";
}else{
	echo "不是同一个对象"; //输出是一个对象
}

class multi extends single{
	public function __construct(){

	}
}
$s1 = new multi();
$s2 = new multi();

if($s1 === $s2){
	echo "是同一个对象";
}else{
	echo "不是同一个对象"; //输出不是一个对象
}

第五步，用final，防止继承，被修改权限

class single{
	protected static $ins = null;

	public static function getIns(){
		if(slef::$ins === null){
			self::$ins = new self();
		}

		return self::$ins;
	}

	//方法前加final，则方法不能被覆盖，类前加final，则类不能被继承
	final protected function __construct(){

	}
}

$s1 = single::getIns();
$s2 = clone $s1;//被克隆了


第六步，禁止克隆

class single{
	protected static $ins = null;

	public static function getIns(){
		if(slef::$ins === null){
			self::$ins = new self();
		}

		return self::$ins;
	}

	//方法前加final，则方法不能被覆盖，类前加final，则类不能被继承
	final protected function __construct(){

	}

	//封锁克隆
	final protected function __clone(){

	}
}

$s1 = single::getIns();
$s2 = clone $s1;//会报错


```
