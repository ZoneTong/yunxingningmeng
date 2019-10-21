;;
const INFO = 0;
const ALERT = 1;
function MessageBox(level, message) {
    if ($('body>#messageboxModal').length == 0){
        var op = document.createElement("div");
        op.innerHTML = `
		<div class="modal fade" id="messageboxModal" tabindex="-1" role="dialog" aria-labelledby="messageboxLabel">
			<div class="modal-dialog" role="document">
				<div class="modal-content">
					<div class="modal-header">
						<button type="button" class="close" data-dismiss="modal" aria-label="Close"><span
								aria-hidden="true" >&times;</span></button>
						<h4 class="modal-title" id="messageboxLabel">警告</h4>
					</div>
					<div class="modal-body">
								<label id="messageboxText" class="control-label">message1</label>
					</div>
					<div class="modal-footer">
						<button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
					</div>
				</div>
			</div>
		</div>`
        document.body.appendChild(op);
    }
    switch (level) {
        case ALERT:
            $('#messageboxLabel').html('警告').css('color', 'red');   
            break;
    
        default:
            $('#messageboxLabel').html('信息').css('color','blue');
            break;
    }
    $('#messageboxText').html(message)
    $('#messageboxModal').modal();
}

function parseDateString(s ){
	let year = parseInt(s.substr(0,4));
	let month = parseInt(s.substr(4,2)) - 1;
	let date = parseInt(s.substr(6,2));
	return new Date(year,month,date);
}

Date.prototype.toDateInputValue = (function () {
	let local = new Date(this);
	local.setMinutes(this.getMinutes() - this.getTimezoneOffset());
	return local.toJSON().slice(0, 10);
});

function toDBDate(d ){
	let year = d.getFullYear(); //获取完整的年份(4位,1970-????)
	let month = d.getMonth() + 1; //获取当前月份(0-11,0代表1月)
	let day = d.getDate(); //获取当前日(1-31)
	return ""+year+month+day;
}