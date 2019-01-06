
## 1.多态
```
举例子
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

## 2.简单工厂(面向接口开发)

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

## 3.工厂方法

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

## 4.单例模式

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

## 观察者模式

php提供的两个接口，一个被观察者接口SplSubject，一个或多个观察者接口SPLObserver,和一个可以储存对象的类SplObjectStorage。被观察者有三个方法，需要实现这三个方法，一个attach可以理解为添加一个观察者，detach可以理解为删除掉一个观察者，一个notify里面做循环执行被观察者的update方法(被观察者被存储在splobjectstorage类里面)，update方法把本类作为参数传进去。

//被观察者
```
class user implements SplSubject{
	public $lognum;
	public $work;

	public $observes = null;

	public function __construct($lognum,$work){

		$this->lognum = rand(1,10);
		$this->work = $work;

		$this->observes = new SplObjectStorage();

	}

	public function login(){

		$this->notify();
	}

	public function attach(SPLObserver $observer){

		$this->observes->attach($observer);

	}

	public function detach(SPLObserver $observer){
		$this->observes->detach($observer);
	}

	public function notify(){
		$this->observes->rewind();

		while ($this->observes->valid()) {
			
			$observer = $this->observes->current();

			$observer->update($this);

			$this->observes->next();
		}
	}

}


//安全观察者
class secrity implements SPLObserver{

	public function update(SplSubject $subject){

		if($subject->lognum<3){
			echo "安全登陆观察者：这是第".$subject->lognum."安全登陆</br>";
		}else{
			echo "安全登陆观察者：这是第".$subject->lognum."登陆异常</br>";;
		}
	}
}

class leader implements SPLObserver{

	protected $name;

	public function __construct($name){

		$this->name = $name;

	}

	public function update(SplSubject $subject){

		if($subject->work == '1' ){

			echo "老板观察者：".$this->name."叫你工作了";
		}else{

			echo "老板观察者：你可以走了";
		}
	}

}

$user = new user(1,'0');

$user->attach(new secrity());

$user->attach(new leader('老板'));

$user->login();

```