function fileUpload(C,index,uploadSize,allSize,name){
	var add = 1024*1024*10//10mb传一次
	var file=C.find('input[type="file"]:enabled')[0].files[0];
	if(!file){
		C.submit();
		return;
	}
	if(index==0){
		name=file.name+"_。。_"+file.name+"_"+new Date().getTime();
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
		    url: '/file/ajaxUploadTmp?name='+encodeURIComponent(name)+'&index='+index+'&blockSize='+add,
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
				fileUpload(C,index+1,uploadSize,allSize,name);
			}
			
		}).fail(function(res) {
			C.enableForm();
			C.toggleClass("loading",false);
			$.zui.showMessager("上传失败")
		});
	}else{
		console.log($('input[type="file"]:enabled').eq(0));
		$('input[type="file"]:enabled').eq(0).parent().find('.file-input-normal .file-input-rename').attr("name",name)
		$('input[type="file"]:enabled').eq(0).parent().find('.file-input-normal .file-input-delete').attr("name",name)
		C.find('input[type="file"]:enabled').eq(0).remove()
		C.append('<input type="hidden" name="uploadFileTmpName" value="'+name+'" />');
		fileUpload(C,0,uploadSize,allSize,"")
	}
}