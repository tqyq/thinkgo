function timeFmt(v) {
        	var date = new Date();
        	date.setTime(v * 1000);
        	return date.getFullYear() + '-' + (date.getMonth() + 1) + '-' + date.getDate() + ' ' + date.getHours() + ':' + date.getMinutes() + ':' + date.getSeconds();
        }

function dateFmt(v) {
	var date = new Date();
	date.setTime(v * 1000);
	return date.getFullYear() + '-' + (date.getMonth() + 1) + '-' + date.getDate();
}
