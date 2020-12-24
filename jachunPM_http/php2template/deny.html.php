<?php
/**
 * The html template file of deny method of user module of ZenTaoPMS.
 *
 * @copyright   Copyright 2009-2015 青岛易软天创网络科技有限公司(QingDao Nature Easy Soft Network Technology Co,LTD, www.cnezsoft.com)
 * @license     ZPL (http://zpl.pub/page/zplv12.html)
 * @author      Chunsheng Wang <chunsheng@cnezsoft.com>
 * @package     ZenTaoPMS
 * @version     $Id: deny.html.php 4129 2013-01-18 01:58:14Z wwccss $
 */
include '../../common/view/header.lite.html.php';
?>
<div class='container'>
  <div class='modal-dialog'>
    <div class='modal-header'><strong><?php echo $app->user->account, ' ', $lang->user->deny;?></strong></div>
    <div class='modal-body'>
      <div class='alert with-icon alert-pure'>
        <i class='icon-exclamation-sign'></i>
        <div class='content'>
        <?php
        if($denyType == 'nopriv')
        {
            $moduleName = isset($lang->$module->common)  ? $lang->$module->common  : $module;
            $methodName = isset($lang->$module->$method) ? $lang->$module->$method : $method;

            /* find method name if method is lowercase letter. */
            if(!isset($lang->$module->$method))
            {
                $tmpLang = array();
                foreach($lang->$module as $key => $value) $tmpLang[strtolower($key)] = $value;
                $methodName = isset($tmpLang[$method]) ? $tmpLang[$method] : $method;
            }

            printf($lang->user->errorDeny, $moduleName, $methodName);
        }

        if($denyType == 'noview')
        {
            $menuName = $menu;
            if(isset($lang->menu->$menu))list($menuName) = explode('|', $lang->menu->$menu);
            printf($lang->user->errorView, $menuName);
        }
        ?>
        </div>
      </div>
    </div>
    <div class='modal-footer'>
    <?php
    $isOnlybody = helper::inOnlyBodyMode();
    unset($_GET['onlybody']);
    echo html::a($this->createLink($config->default->module), $lang->my->common, ($isOnlybody ? '_parent' : ''), "class='btn'");
    if($refererBeforeDeny) echo html::a(helper::safe64Decode($refererBeforeDeny), $lang->user->goback, ($isOnlybody ? '_parent' : ''), "class='btn'");
    echo html::a($this->createLink('user', 'logout', "referer=" . helper::safe64Encode($denyPage)), $lang->user->relogin, ($isOnlybody ? '_parent' : ''), "class='btn btn-primary'");
    ?>
    </div>
  </div>
</div>
</body>
</html>
