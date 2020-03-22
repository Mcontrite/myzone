// 表单快捷键提交 CTRL+ENTER   / form quick submit
$('form').keyup(function(e) {
	var jthis = $(this);
	if ((e.ctrlKey && (e.which == 13 || e.which == 10)) || (e.altKey && e.which == 83)) {
		jthis.trigger('submit');
		return false;
	}
});
// 点击响应整行：方便手机浏览  / check response line
$('.tap').on('click', function(e) {
	var href = $(this).attr('href') || $(this).data('href');
	if (e.target.nodeName == 'INPUT') return true;
	if ($(window).width() > 992) return;
	if (e.ctrlKey) {
		window.open(href);
		return false;
	} else {
		window.location = href;
	}
});
// 点击响应整行：导航栏下拉菜单   / check response line
$('ul.nav > li').on('click', function(e) {
	var jthis = $(this);
	var href = jthis.children('a').attr('href');
	if (e.ctrlKey) {
		window.open(href);
		return false;
	}
});
// 点击响应整行：，但是不响应 checkbox 的点击  / check response line, without checkbox
$('.article input[type="checkbox"]').parents('td').on('click', function(e) {
	e.stopPropagation();
})
// 确定框 / confirm / GET / POST
$('a.confirm').on('click', function() {
	var jthis = $(this);
	var text = jthis.data('confirm-text');
	$.confirm(text, function() {
		var method = xn.strtolower(jthis.data('method'));
		var href = jthis.data('href') || jthis.attr('href');
		if (method == 'post') {
			$.xpost(href, function(code, message) {
				if (code == 0) {
					window.location.reload();
				} else {
					alert(message);
				}
			});
		} else {
			//window.location = jthis.attr('href');
		}
	})
	return false;
});
// 选中所有 / check all
// <input class="checkall" data-target=".tid" />
$('input.checkall').on('click', function() {
	var jthis = $(this);
	var target = jthis.data('target');
	jtarget = $(target);
	jtarget.prop('checked', this.checked);
});
// 删除 / Delete post
$('body').on('click', '.post_delete', function() {
	var jthis = $(this);
	var href = jthis.data('href');
	var isfirst = jthis.attr('isfirst');
	if (window.confirm(lang.confirm_delete)) {
		$.xpost(href, function(code, message) {
			var isfirst = jthis.attr('isfirst');
			if (code == 0) {
				if (isfirst == '1') {
					$.location('<?php echo url("forum-$fid");?>');
				} else {
					// 删掉楼层
					jthis.parents('.post').remove();
					// 回复数 -1
					var jposts = $('.posts');
					jposts.html(xn.intval(jposts.html()) - 1);
				}
			} else {
				$.alert(message);
			}
		});
	}
	return false;
});
// 引用 / Quote
$('body').on('click', '.post_reply', function() {
	var jthis = $(this);
	var tid = jthis.data('tid');
	var pid = jthis.data('pid');
	var jmessage = $('#message');
	var jli = jthis.closest('.post');
	var jpostlist = jli.closest('.postlist');
	var jadvanced_reply = $('#advanced_reply');
	var jform = $('#quick_reply_form');
	if (jli.hasClass('quote')) {
		jli.removeClass('quote');
		jform.find('input[name="quotepid"]').val(0);
		jadvanced_reply.attr('href', xn.url('post-create-' + tid));
	} else {
		jpostlist.find('.post').removeClass('quote');
		jli.addClass('quote');
		var s = jmessage.val();
		jform.find('input[name="quotepid"]').val(pid);
		jadvanced_reply.attr('href', xn.url('post-create-' + tid + '-0-' + pid));
	}
	jmessage.focus();
	return false;
});
