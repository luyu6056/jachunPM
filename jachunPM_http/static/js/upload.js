var add = 1024*1024*10//10mb传一次
function fileUpload(C,index,uploadSize,allSize){
	var file=C.find('input[type="file"]:enabled')[0].files[0];
	if(!file){
		C.submit();
		return;
	}
	var max=file.size;
	var begin=add*index;
	var end=begin+add;
	if(begin<max){
		if(end>max) end=max;
		for(var i=0;i<20;i++){
			$('.form-ajax').removeClass('percent_'+(i*5));
		}
		$('.form-ajax').addClass('percent_'+Math.floor(uploadSize/allSize*20)*5);
		$.ajax({
		    url: '/file/ajaxUploadTmp?name='+encodeURIComponent(file.name)+'&index='+index+'&blockSize='+add,
		    type: 'POST',
		    cache: false,
		    data: file.slice(begin,end),
		    processData: false,
		    contentType: false,
		    async:true,
		}).done(function(res) {
			if(res){
				C.enableForm();
				C.toggleClass("loading",false);
				$.zui.showMessager("上传失败:"+res)
			}else{
				uploadSize+=end-begin;
				fileUpload(C,index+1,uploadSize,allSize);
			}
			
		}).fail(function(res) {
			C.enableForm();
			C.toggleClass("loading",false);
			$.zui.showMessager("上传失败")
		});
	}else{
		C.find('input[type="file"]:enabled').eq(0).remove()
		C.append('<input type="hidden" name="uploadFileTmpName" value="'+file.name+'" />');
		fileUpload(C,0,uploadSize,allSize)
	}
}
							