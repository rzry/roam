static ND:i32 = 5;
const SD:u32 = 111;
fn main() {
    //num();
    //tuple_array();
    //funcx();
    //str();
    //operator();
    //borrows();
    //reference();
    structs();
}
//常量
fn num(){
    let s = 4;
    print!("没有mut 不可变--> {} \n",s);

    let mut d = 13;
    //定义了一个变量后 一定要使用了 才能在修改后 读取有用
    println!("有mnt 读取第一次 --> {}",d);
    d = 11;
    println!("有mnt 修改后再读取 --> {}",d);

    println!("static 常量 --> {}",ND);

    println!("const 常亮--> {}",SD);
    //const 和 static 的区别是 const 可能会被内联优化不一定会分配空间

    let a = 1;
    let a = a*2;
    println!("直接定义两遍a --> {}",a);

    let str = "best wishes emacser ";
    println!("str len --> {}, str value --> {} ",str.len(),str)
}
//元祖和切片
fn tuple_array(){
    //tuple
    let xs:(i32,f64,u8) = (100,1.1,1);
    println!("xs array value --> {:?}",xs);

    let xd = (1,2.4,1);
    println!("xs value --> {:?}",xd);

    let xds=(1.3f32,2,3);
    println!("xds value --> {:?}",xds);

    //array
    let ars = [1,2,3,4];
    println!("ars value --> {:?}",ars);

    let arsd= [1.2f64,2.2,3.1];
    println!("arsd value --> {:?}",arsd);

    let ds:[i32;0] = [];
    println!("ds -> {:?}",ds);

    let pair:(char,i32,&str) = ('q',123,"hello");
    println!("pair value --> {:?}",pair);
    //元祖() 多类型 可以单独赋值
    let (hello,world) = (1,2);
    println!("hello value -> {} world value ->{}",hello,world);

}
//函数
fn funcx(){
    let (add_nud ,result) = add_num(10);
    println!("add_nud result value -> {}",result);
}
fn add_num(a:i32) ->(i32,i32){
    (a , a+1)
}
//字符串
fn str(){
    let stra = "hello";
    println!("stra value = {}",stra);
    let strb = "world".to_string();
    println!("strb value = {} ,--> {} ",&strb,strb);
}
//运算符
fn operator(){
    let a = 11;
    let c = (a as f32)/2.1;
    //保留小数位输出
    println!("c value --> {:.2}",c)
}
//引用,解引用
fn reference(){
    let str1 = String::from("best wishes emacser");
    println!("str1 --> {} ,str1 len --> {}\n",&str1,str1.len());

    println!("str1 len value --> {}",lenstr(&str1));
    fn lenstr(s:&String)->usize{
        s.len()
    }
    //str 是一个不可变的固定长度的字符串
    //String 是一个可变 堆上分配的字符串缓冲区

    let str2 = "emacser";
    //转成 String
    let str3 = str2.to_string();
    println!("str 转成 String 后 len --> {}",lenstr(&str3));
}
fn borrows(){
    let v = &mut[1,2,3,4];//可变
    //进入别的作用域了
    {
        let zero = v.get_mut(0).unwrap();
        *zero +=10;
        //这样修改是不会改变外部v的值
        v[0] = 11;
    }
    println!("array first value --> {:?}",v);
    v[0] = 2;
    println!("array value --> {:?}",v);
}

//struct
struct Vec2 {
    x: f64, // 64-bit floating point, aka "double precision"
    y: f64,
}
fn structs(){

    let v1 = Vec2{x:1.2,y:2.1};
    println!("v1 struct value x --> {:?} y -->{:?}", v1.x,v1.y);
    let v2 = Vec2{
      x:22.1,
        ..v1
    };
    println!(" v2 struct value x --> {:?} y -->{:?}", v2.x,v2.y);
    //值保留了 x 的值
    let Vec2{x,..} = v1;
    println!(" v2 struct value x --> {:?}", x);
}
