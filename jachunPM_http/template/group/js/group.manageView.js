function toggleProduct()
{ 
    $('#productBox').toggle($('#product').prop("checked"));
}

function toggleProject()
{
    $('#projectBox').toggle($('#project').prop("checked"));
}

$('input:checkbox[name^="allchecker"]').change(function()
{
    setTimeout(function(){toggleProduct(),toggleProject()}, 50);
});
$('#product').change(function(){toggleProduct();})
$('#project').change(function(){toggleProject();})

$(function()
{
    toggleProduct();
    toggleProject();
    $('.group-item :checkbox[name^="Acl"]').change(function()
    {
        var allChecked = true;
        $('.group-item :checkbox[name^="Acl"]').each(function()
        {
            if(!$(this).prop('checked')) allChecked = false;
        })
        $('input:checkbox[name^="allchecker"]').prop('checked', allChecked);
    })
})
