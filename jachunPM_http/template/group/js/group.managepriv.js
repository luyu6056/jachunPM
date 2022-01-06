function showPriv(value)
{
  location.href = createLink('group', 'managePriv', "type=byGroup&param="+ groupID + "&menu=&version=" + value);
}

/**
 * Control the actions select control for a module.
 * 
 * @param   string $module 
 * @access  public
 * @return  void
 */
function setModuleActions(module)
{
    $('#actionBox select').addClass('hidden');          // Hide all select first.
    $('#actionBox select').val('');                     // Unselect all select.
    $('.' + module + 'Actions').removeClass('hidden');  // Show the action control for current module.
}

function setNoChecked()
{
    var noCheckValue = '';
    $(':checkbox').each(function(){
        if(!$(this).prop('checked') && $(this).next('span').attr('id') != undefined) noCheckValue = noCheckValue + ',' + $(this).next('span').attr('id');
    })
    $('#noChecked').val(noCheckValue);
}

$(function()
{
    $('#privList > tbody > tr > th input[type=checkbox]').change(function()
    {
        var id      = $(this).attr('id');
        var checked = $(this).prop('checked');

        if(id == 'allChecker')
        {
            $('input[type=checkbox]').prop('checked', checked);
        }
        else
        {
            $(this).parents('tr').find('input[type=checkbox]').prop('checked', checked);
        }
    });
})
