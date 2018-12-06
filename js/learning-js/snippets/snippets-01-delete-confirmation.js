//---------------------------------------------------------
// Delete Confirmation (confirmAfterClick)

<a id="delete_link" href="/delete.php">Delete Script</a>

$('#delete_link').click(function(){
    // 只有系统消息框能起到作用, 但不美观
    return confirm("Are you sure you want to delete?");
    
    // 其它javascript库提供的弹出窗都是异步方式，不起作用
    //bootbox.confirm("Are you sure?", function(result) {
    //    Toast.show("Confirm result: "+result);
    //});
    
    //layer.confirm('Are you sure you want to delete?');

    //return false; // 只有加上这一句，javascript弹出窗才有机会显示
});

// 为了使用javascript弹出窗，只能采用这种变通的方法

<a onclick="mainSmartAdminDelete(this); return false;" 
   data-link="/country/delete/4" 
   href="javascript:void(0);">

function mainSmartAdminDelete(obj) {
   $.SmartMessageBox({
        title   : "Delete",
        content : "Are you sure?",
        buttons : '[No][Yes]'
    }, function(ButtonPressed) {
        if (ButtonPressed === "Yes") {
            redirect($(obj).attr("data-link"));
        }
        if (ButtonPressed === "No") {
            // do nothing
        }
    });
    return false; // 总是返回 false，这样弹出窗才能显示
}

function redirect(link) {
    window.location = link;
}
