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