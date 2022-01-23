$(document).ready(function()
{
    var s = ',leave,makeup,overtime,lieu,trip,egress,allAbsent,halfAbsent,';
    if(status == 'normal' || (reason && s.indexOf(',' + status + ',') == -1 ))
    {
        $('.editMode').hide();
        $('.viewMode').show();
    }
    else
    {
        $('.editMode').show();
        $('.viewMode').hide();
    }

    $('.edit').click(function()
    {
        $('.editMode').show();
        $('.viewMode').hide();
    })
})
