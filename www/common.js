;;
if (typeof BootstrapDialog != "undefined") {
	BootstrapDialog.DEFAULT_TEXTS[BootstrapDialog.TYPE_DEFAULT] = '信息';
	BootstrapDialog.DEFAULT_TEXTS[BootstrapDialog.TYPE_INFO] = '信息';
	BootstrapDialog.DEFAULT_TEXTS[BootstrapDialog.TYPE_PRIMARY] = '重要';
	BootstrapDialog.DEFAULT_TEXTS[BootstrapDialog.TYPE_SUCCESS] = '成功';
	BootstrapDialog.DEFAULT_TEXTS[BootstrapDialog.TYPE_WARNING] = '警告';
	BootstrapDialog.DEFAULT_TEXTS[BootstrapDialog.TYPE_DANGER] = '危险';
	BootstrapDialog.DEFAULT_TEXTS['OK'] = '确定';
	BootstrapDialog.DEFAULT_TEXTS['CANCEL'] = '取消';
	BootstrapDialog.DEFAULT_TEXTS['CONFIRM'] = '确认';
}

function parseDateString(s ){
	let year = parseInt(s.substr(0,4));
	let month = parseInt(s.substr(4,2)) - 1;
	let date = parseInt(s.substr(6,2));
	return new Date(year,month,date);
}

Date.prototype.toDateInputValue = (function () {
	let local = new Date(this);
	if (!(local instanceof Date) || isNaN(local.getTime())) {
		return;
	}

	local.setTime(this.getTime() - this.getTimezoneOffset()*60000);
	return local.toJSON().slice(0, 10);
});

Date.prototype.addDay = (function (n) {
	let local = new Date(this);
	if (!(local instanceof Date) || isNaN(local.getTime())) {
		return;
	}
	
	 let ms = local.getTime()
	 ms += n * 24 * 60 * 60 * 1000
	 local.setTime(ms)
	 return local;
});

function toDBDate(d ){
	let year = d.getFullYear(); //获取完整的年份(4位,1970-????)
	let month = d.getMonth() + 1; //获取当前月份(0-11,0代表1月)
	let day = d.getDate(); //获取当前日(1-31)
	return ""+year+month+day;
}

// 解决浮点数精度问题
function _roundFloat(number, precision) {
	return Math.round(+number + 'e' + precision) / Math.pow(10, precision);
}

function floatRound(number) {
	return _roundFloat(number, 4)
}

$('a[href$=html]').each(function () {
	$(this).attr('val', $(this).attr('href')).attr('href','javascript:;').on('click', function () {
		var url = $(this).attr('val');
		window.location.href = url;
	}) ;
});

var multiline = function (fn) {
	var str = fn.toString().split('\n');
	return str.slice(1, str.length - 1 ).join('\n');
};



// h5+ custom js
// 设置顶层对象
(typeof window !== 'undefined'
  ? window : (typeof process === 'object' &&
             typeof require === 'function' &&
             typeof global === 'object')
             ? global
            : this);

function searchTableSql(table, condition, count = false){
	let columns = '*';
	if (count){
		columns = 'count(*)';
	}
	let ssql = 'SELECT '+columns+' FROM '+table+' t WHERE t.deleted=0';
	let keys = [];
	if (condition.searchKey != ''){
		if (condition.searchField == 0){
			ssql += ' and t.date LIKE "%'+condition.searchKey+'%"';
		} else {
			ssql += ' and t.provider LIKE "%'+condition.searchKey+'%"';
		}
	}
	
	if (count){
		console.log(ssql);
		return ssql;
	}
	
	ssql += ' LIMIT '+condition.pageSize+' OFFSET '+(condition.pageNumber-1)*condition.pageSize;
	
	return ssql;
}

console.log(loginReady);
// if (false){
if (!window.loginReady){
	var loginReady = () => {};
	var menuUI = () => {};
	var tableUI = () => {};
	console.log('no loginReady start');

window.searchPurchaseRecords = condition => new Promise(resolve => {
	console.log("searchPurchaseRecords");
	gDB.transaction(function(tx){
		console.log("before searchPurchaseRecords t_purchase_record");
	        tx.executeSql(searchTableSql('t_purchase_record', condition, true), [], (tx,okresult)=>{
				console.log("after searchPurchaseRecords t_purchase_record");
				// console.log(okresult.rows);
				let r = {};
				r.total = okresult.rows.item(0)['count(*)'];
				r.rows= [];
				if (r.total > 0 ){
					tx.executeSql(searchTableSql('t_purchase_record', condition, false), [], (tx, okresult) => {					
						for (let i = 0; i < r.total; i++){
							let rr = okresult.rows.item(i);
							for (p in rr){
								rr[ p.substr(0,1).toUpperCase()+p.substr(1) ] = rr[p];
							}
							r.rows.push(rr);
						}
						resolve(r);
					},(tx, errresult) => {
						console.log(errresult.message);
					});
				}else{
					resolve({total:0, rows:[]});
				} 
			}, (tx, errresult)=>{
				console.log(errresult.message);
				resolve({total:0});
			});
	});
});

window.newOrEditPurchaseRecord = (record) => new Promise(resolve => {
	console.log("newOrEditPurchaseRecord");
	gDB.transaction(function(tx){
		tx.executeSql(newRecordSql('t_purchase_record', record),[new Date(), new Date()],(tx,okr)=>{
			console.log('newOrEditPurchaseRecord rowsAffected', okr.rowsAffected);
			resolve('ok');
		}, (tx,errr)=>{
			resolve(errr.message);
		});		
	});
});

window.delPurchaseRecord = (id) => new Promise(resolve => {
	console.log("delPurchaseRecord");
	gDB.transaction(function(tx){
		tx.executeSql('UPDATE t_purchase_record SET deleted=1 WHERE id=?',[id],(tx,okr)=>{
			console.log('delPurchaseRecord rowsAffected', okr.rowsAffected);
			resolve('ok');
		}, (tx,errr)=>{
			resolve(errr.message);
		});		
	});
});

function newRecordSql(table, record){
	let ssql_pre = 'INSERT INTO '+table+' (';
	let ssql_mid = ')VALUES (';
	let ssql_suf = ')';
	for (p in record ){
		if (p == 'edit_status' || p == 'Id'){
			continue;
		}
	
		ssql_pre += p.toLowerCase() + ',';
		let v = record[p];
		if (typeof v == 'string'){
			v = '"'+v+'"';
		}
		ssql_mid += v +','
	}
	ssql_pre += 'created,updated';
	ssql_mid += '?,?';
	
	console.log(ssql_pre, ssql_mid,ssql_suf);
	return ssql_pre + ssql_mid + ssql_suf;
}

window.searchCostRecords = () => { 
	return {
	total: 2,
	rows: [{}]
	}
};

var db_sql_func = function(){ /*
		CREATE TABLE IF NOT EXISTS `t_user` (
			`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			`name` varchar(255) NOT NULL DEFAULT '' ,
			`password` varchar(255) NOT NULL DEFAULT '' ,
			`deleted` integer NOT NULL DEFAULT 0 ,
			`created` datetime NOT NULL,
			`updated` datetime NOT NULL
		);
		
		CREATE TABLE IF NOT EXISTS `t_purchase_record` (
		    `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		    `date` varchar(255) NOT NULL DEFAULT '' ,
		    `provider` varchar(255) NOT NULL DEFAULT '' ,
		    `logistics` varchar(255) NOT NULL DEFAULT '' ,
		    `specifications` varchar(255) NOT NULL DEFAULT '' ,
		    `number` integer NOT NULL DEFAULT 0 ,
		    `weight` real NOT NULL DEFAULT 0 ,
		    `price` real NOT NULL DEFAULT 0 ,
		    `total` real NOT NULL DEFAULT 0 ,
		    `deleted` integer NOT NULL DEFAULT 0 ,
		    `created` datetime NOT NULL,
		    `updated` datetime NOT NULL
		);
				
		CREATE TABLE IF NOT EXISTS `t_sale_record` (
		    `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		    `date` varchar(255) NOT NULL DEFAULT '' ,
		    `customer` varchar(255) NOT NULL DEFAULT '' ,
		    `logistics` varchar(255) NOT NULL DEFAULT '' ,
		    `specifications` varchar(255) NOT NULL DEFAULT '' ,
		    `number` integer NOT NULL DEFAULT 0 ,
		    `weight` real NOT NULL DEFAULT 0 ,
		    `price` real NOT NULL DEFAULT 0 ,
		    `total` real NOT NULL DEFAULT 0 ,
		    `deleted` integer NOT NULL DEFAULT 0 ,
		    `created` datetime NOT NULL,
		    `updated` datetime NOT NULL
		);
				
		CREATE TABLE IF NOT EXISTS `t_cost_record` (
		    `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		    `date` varchar(255) NOT NULL DEFAULT '' ,
		    `tea_cost` real NOT NULL DEFAULT 0 ,
		    `labor_cost` real NOT NULL DEFAULT 0 ,
		    `freight` real NOT NULL DEFAULT 0 ,
		    `postage` real NOT NULL DEFAULT 0 ,
		    `carton_cost` real NOT NULL DEFAULT 0 ,
		    `consumables` real NOT NULL DEFAULT 0 ,
		    `packing_charges` real NOT NULL DEFAULT 0 ,
		    `total` real NOT NULL DEFAULT 0 ,
		    `deleted` integer NOT NULL DEFAULT 0 ,
		    `created` datetime NOT NULL,
		    `updated` datetime NOT NULL
		);
				
		CREATE TABLE IF NOT EXISTS `t_stock_record` (
		    `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		    `date` varchar(255) NOT NULL DEFAULT '' ,
		    `specifications` varchar(255) NOT NULL DEFAULT '' ,
		    `operation` varchar(255) NOT NULL DEFAULT '' ,
		    `weight` real NOT NULL DEFAULT 0 ,
		    `money` real NOT NULL DEFAULT 0 ,
		    `surplus_stock` real NOT NULL DEFAULT 0 ,
		    `deleted` integer NOT NULL DEFAULT 0 ,
		    `created` datetime NOT NULL,
		    `updated` datetime NOT NULL
		);
				
		CREATE TABLE IF NOT EXISTS `t_finance_statics` (
		    `id` integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		    `date` varchar(255) NOT NULL DEFAULT '' ,
		    `purchase` real NOT NULL DEFAULT 0 ,
		    `sale` real NOT NULL DEFAULT 0 ,
		    `cost` real NOT NULL DEFAULT 0 ,
		    `last_balance` real NOT NULL DEFAULT 0 ,
		    `total` real NOT NULL DEFAULT 0 ,
		    `purchased_stock` real NOT NULL DEFAULT 0 ,
		    `saled_stock` real NOT NULL DEFAULT 0 ,
		    `last_stock` real NOT NULL DEFAULT 0 ,
		    `total_stock` real NOT NULL DEFAULT 0 ,
		    `deleted` integer NOT NULL DEFAULT 0 ,
		    `created` datetime NOT NULL,
		    `updated` datetime NOT NULL
		);
*/};
/*
*/

function openAbout(){
	
}

var createdb_sqls= multiline(db_sql_func).split(';');
// ios 支持的大小只有不到50MB
console.log('before openDatabase');
var gDB = window.openDatabase('data_db', '1.0', 'Test DB', 20 * 1024 * 1024, function(){console.log('opening openDatabase')});
var initdb = function() {
	console.log('initDB begin');
	console.log('safari');
	gDB.transaction(function (tx) {
		// 创建表
	   $.each(createdb_sqls, (i, createdb_sql) => {
		   createdb_sql = createdb_sql.trim();
		   if (createdb_sql == ''){
			   return;
		   }
		   tx.executeSql(createdb_sql,[],
			function(tx,results){
			   // console.log(results);
			},function(tx,error){
			   alert('创建表失败:' + error.message);
		   });
		});
		
		// 检测是否存在用户
		tx.executeSql('SELECT * FROM t_user', [],
			function(tx,results){
			   if (results.rows.length == 0) {
				   tx.executeSql('INSERT INTO t_user (name, password, created, updated) VALUES ("yunxing", "750588",?,?)',[new Date(),new Date()],
					function(tx,results){
					   console.log(results);
					  },function(tx,error){
					   alert('插入表失败:' + error.message);
					  });
			   }
					},function(tx,error){
			   alert('查询表失败:' + error.message);
		   });
		
	});
	
	// gDB.transaction(function (tx) {
	// 	tx.executeSql('SELECT * FROM t_user', [], function (tx, results) {
	// 		var len = results.rows.length, i;
	// 		msg = "<p>查询记录条数: " + len + "</p>";
	// 		document.querySelector('#status').innerHTML +=  msg;
	// 		// console.log(len);
	// 		for (i = 0; i < len; i++){
	// 			msg = "<p><b>" + results.rows.item(i).name + "</b></p>";
	// 			document.querySelector('#status').innerHTML +=  msg;
	// 		}
	// 	}, function(tx, error){  
 //          alert( +  error.message); 
 //      });
	// });
};

window.verifyPassword = (name, password) => {
	console.log('verifyPassword start');
	return new Promise((resolve, reject) => {
		console.log('Promise start');
		gDB.transaction(function (tx) {
			tx.executeSql('SELECT * FROM t_user where name=? and password=?', [name,password],
			function(tx,results){
				console.log('verifyPassword succ');
				if (results.rows.length == 1){
					// console.log('login success');
					resolve( 'ok');
				}else{
					resolve('用户名或密码错误')
				}
			},function(tx,error){
			console.log('verifyPassword fail');
			   resolve( error.message);
			});
		});
	});
};


}else{
	var initdb = () => {};
	// window.verifyPassword = () => {return 'failed';};
}