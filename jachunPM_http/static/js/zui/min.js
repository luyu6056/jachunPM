/*!
 * ZUI: Zentao template - v1.8.1 - 2018-11-20
 * http://zui.sexy
 * GitHub: https://github.com/easysoft/zui.git
 * Copyright (c) 2018 cnezsoft.com; Licensed MIT
 */
!
function(t, e, i) {
	"use strict";
	if ("undefined" == typeof t) throw new Error("ZUI requires jQuery");
	t.zui || (t.zui = function(e) {
		t.isPlainObject(e) && t.extend(t.zui, e)
	});
	var n = {
		all: -1,
		left: 0,
		middle: 1,
		right: 2
	},
		o = 0;
	t.zui({
		uuid: function(t) {
			var e = 1e3 * (new Date).getTime() + o++ % 1e3;
			return t ? e : e.toString(36)
		},
		callEvent: function(e, n, o) {
			if (t.isFunction(e)) {
				o !== i && (e = t.proxy(e, o));
				var s = e(n);
				return n && (n.result = s), !(s !== i && !s)
			}
			return 1
		},
		clientLang: function() {
			var i, n = e.config;
			if ("undefined" != typeof n && n.clientLang && (i = n.clientLang), !i) {
				var o = t("html").attr("lang");
				i = o ? o : navigator.userLanguage || navigator.userLanguage || "zh_cn"
			}
			return i.replace("-", "_").toLowerCase()
		},
		strCode: function(t) {
			var e = 0;
			if (t && t.length) for (var i = 0; i < t.length; ++i) e += i * t.charCodeAt(i);
			return e
		},
		getMouseButtonCode: function(t) {
			return "number" != typeof t && (t = n[t]), t !== i && null !== t || (t = -1), t
		}
	}), t.fn.callEvent = function(e, n, o) {
		var s = t(this),
			a = e.indexOf(".zui."),
			r = a < 0 ? e : e.substring(0, a),
			l = t.Event(r, n);
		if (o === i && a > 0 && (o = s.data(e.substring(a + 1))), o && o.options) {
			var h = o.options[r];
			t.isFunction(h) && (l.result = t.zui.callEvent(h, l, o))
		}
		return s.trigger(l), l
	}, t.fn.callComEvent = function(e, n, o) {
		o === i || t.isArray(o) || (o = [o]);
		var s = this,
			a = s.triggerHandler(n, o),
			r = e.options[n];
		return r && (a = r.apply(e, o)), a
	}
}(jQuery, window, void 0), function(t) {
	"use strict";
	t.fn.fixOlPd = function(e) {
		return e = e || 10, this.each(function() {
			var i = t(this);
			i.css("paddingLeft", Math.ceil(Math.log10(i.children().length)) * e + 10)
		})
	}, t(function() {
		t(".ol-pd-fix,.article ol").fixOlPd()
	})
}(jQuery), +
function(t) {
	"use strict";
	var e = '[data-dismiss="alert"]',
		i = "zui.alert",
		n = function(i) {
			t(i).on("click", e, this.close)
		};
	n.prototype.close = function(e) {
		function n() {
			a.trigger("closed." + i).remove()
		}
		var o = t(this),
			s = o.attr("data-target");
		s || (s = o.attr("href"), s = s && s.replace(/.*(?=#[^\s]*$)/, ""));
		var a = t(s);
		e && e.preventDefault(), a.length || (a = o.hasClass("alert") ? o : o.parent()), a.trigger(e = t.Event("close." + i)), e.isDefaultPrevented() || (a.removeClass("in"), t.support.transition && a.hasClass("fade") ? a.one(t.support.transition.end, n).emulateTransitionEnd(150) : n())
	};
	var o = t.fn.alert;
	t.fn.alert = function(e) {
		return this.each(function() {
			var o = t(this),
				s = o.data(i);
			s || o.data(i, s = new n(this)), "string" == typeof e && s[e].call(o)
		})
	}, t.fn.alert.Constructor = n, t.fn.alert.noConflict = function() {
		return t.fn.alert = o, this
	}, t(document).on("click." + i + ".data-api", e, n.prototype.close)
}(window.jQuery), function(t, e) {
	"use strict";
	var i = "zui.pager",
		n = {
			page: 1,
			recTotal: 0,
			recPerPage: 10
		},
		o = {
			zh_cn: {
				pageOfText: "第 {0} 页",
				prev: "上一页",
				next: "下一页",
				first: "第一页",
				last: "最后一页",
				"goto": "跳转",
				pageOf: "第 <strong>{page}</strong> 页",
				totalPage: "共 <strong>{totalPage}</strong> 页",
				totalCount: "共 <strong>{recTotal}</strong> 项",
				pageSize: "每页 <strong>{recPerPage}</strong> 项",
				itemsRange: "第 <strong>{start}</strong> ~ <strong>{end}</strong> 项",
				pageOfTotal: "第 <strong>{page}</strong>/<strong>{totalPage}</strong> 页"
			},
			zh_tw: {
				pageOfText: "第 {0} 頁",
				prev: "上一頁",
				next: "下一頁",
				first: "第一頁",
				last: "最後一頁",
				"goto": "跳轉",
				pageOf: "第 <strong>{page}</strong> 頁",
				totalPage: "共 <strong>{totalPage}</strong> 頁",
				totalCount: "共 <strong>{recTotal}</strong> 項",
				pageSize: "每頁 <strong>{recPerPage}</strong> 項",
				itemsRange: "第 <strong>{start}</strong> ~ <strong>{end}</strong> 項",
				pageOfTotal: "第 <strong>{page}</strong>/<strong>{totalPage}</strong> 頁"
			},
			en: {
				pageOfText: "Page {0}",
				prev: "Prev",
				next: "Next",
				first: "First",
				last: "Last",
				"goto": "Goto",
				pageOf: "Page <strong>{page}</strong>",
				totalPage: "<strong>{totalPage}</strong> pages",
				totalCount: "<strong>{recTotal}</strong> in total",
				pageSize: "<strong>{recPerPage}</strong> per page",
				itemsRange: "From <strong>{start}</strong> to <strong>{end}</strong>",
				pageOfTotal: "Page <strong>{page}</strong> of <strong>{totalPage}</strong>"
			}
		},
		s = function(e, n) {
			var a = this;
			a.name = i, a.$ = t(e), n = a.options = t.extend({}, s.DEFAULTS, this.$.data(), n);
			var r = n.lang || t.zui.clientLang();
			a.lang = t.isPlainObject(r) ? t.extend(!0, {}, o[r.lang || t.zui.clientLang()], r) : o[r], a.state = {}, a.set(n.page, n.recTotal, n.recPerPage, !0), a.$.on("click", ".pager-goto-btn", function() {
				var e = t(this).closest(".pager-goto"),
					i = parseInt(e.find(".pager-goto-input").val());
				NaN !== i && a.set(i)
			}).on("click", ".pager-item", function() {
				var e = t(this).data("page");
				"number" == typeof e && e > 0 && a.set(e)
			}).on("click", ".pager-size-menu [data-size]", function() {
				var e = t(this).data("size");
				"number" == typeof e && e > 0 && a.set(-1, -1, e)
			})
		};
	s.prototype.set = function(e, i, o, s) {
		var a = this;
		"object" == typeof e && null !== e && (o = e.recPerPage, i = e.recTotal, e = e.page);
		var r = a.state;
		r || (r = t.extend({}, n));
		var l = t.extend({}, r);
		return "number" == typeof o && o > 0 && (r.recPerPage = o), "number" == typeof i && i >= 0 && (r.recTotal = i), "number" == typeof e && e >= 0 && (r.page = e), r.totalPage = r.recTotal && r.recPerPage ? Math.ceil(r.recTotal / r.recPerPage) : 1, r.page = Math.max(0, Math.min(r.page, r.totalPage)), r.pageRecCount = r.recTotal, r.page && r.recTotal && (r.page < r.totalPage ? r.pageRecCount = r.recPerPage : r.page > 1 && (r.pageRecCount = r.recTotal - r.recPerPage * (r.page - 1))), r.skip = r.page > 1 ? (r.page - 1) * r.recPerPage : 0, r.start = r.skip + 1, r.end = r.skip + r.pageRecCount, r.prev = r.page > 1 ? r.page - 1 : 0, r.next = r.page < r.totalPage ? r.page + 1 : 0, a.state = r, s || l.page === r.page && l.recTotal === r.recTotal && l.recPerPage === r.recPerPage || a.$.callComEvent(a, "onPageChange", [r, l]), a.render()
	}, s.prototype.createLinkItem = function(i, n, o) {
		var s = this;
		n === e && (n = i);
		var a = t('<a title="' + s.lang.pageOfText.format(i) + '" class="pager-item" data-page="' + i + '"/>').attr("href", i ? s.createLink(i, s.state) : "###").html(n);
		return o || (a = t("<li />").append(a).toggleClass("active", i === s.state.page).toggleClass("disabled", !i || i === s.state.page)), a
	}, s.prototype.createNavItems = function(t) {
		var i = this,
			n = i.$,
			o = i.state,
			s = o.totalPage,
			a = o.page,
			r = function(t, o) {
				if (t === !1) return void n.append(i.createLinkItem(0, o || i.options.navEllipsisItem));
				o === e && (o = t);
				for (var s = t; s <= o; ++s) n.append(i.createLinkItem(s))
			};
		t === e && (t = i.options.maxNavCount || 10), r(1), s > 1 && (s <= t ? r(2, s) : a < t - 2 ? (r(2, t - 2), r(!1), r(s)) : a > s - t + 2 ? (r(!1), r(s - t + 2, s)) : (r(!1), r(a - Math.ceil((t - 4) / 2), a + Math.floor((t - 4) / 2)), r(!1), r(s)))
	}, s.prototype.createGoto = function() {
		var e = this,
			i = this.state,
			n = t('<div class="input-group pager-goto" style="width: ' + (35 + 9 * (i.page + "").length + 25 + 12 * e.lang["goto"].length) + 'px"><input value="' + i.page + '" type="number" min="1" max="' + i.totalPage + '" placeholder="' + i.page + '" class="form-control pager-goto-input"><span class="input-group-btn"><button class="btn pager-goto-btn" type="button">' + e.lang["goto"] + "</button></span></div>");
		return n
	}, s.prototype.createSizeMenu = function() {
		var e = this,
			i = this.state,
			n = t('<ul class="dropdown-menu"></ul>'),
			o = e.options.pageSizeOptions;
		"string" == typeof o && (o = o.split(","));
		for (var s = 0; s < o.length; ++s) {
			var a = o[s];
			"string" == typeof a && (a = parseInt(a));
			var r = t('<li><a href="###" data-size="' + a + '">' + a + "</a></li>").toggleClass("active", a === i.recPerPage);
			n.append(r)
		}
		return t('<div class="btn-group pager-size-menu"><button type="button" class="btn dropdown-toggle" data-toggle="dropdown">' + e.lang.pageSize.format(i) + ' <span class="caret"></span></button></div>').addClass(e.options.menuDirection).append(n)
	}, s.prototype.createElement = function(e, i, n) {
		var o = this,
			s = t.proxy(o.createLinkItem, o),
			a = o.lang;
		switch (e) {
		case "prev":
			return s(n.prev, a.prev);
		case "prev_icon":
			return s(n.prev, '<i class="icon ' + o.options.prevIcon + '"></i>');
		case "next":
			return s(n.next, a.next);
		case "next_icon":
			return s(n.next, '<i class="icon ' + o.options.nextIcon + '"></i>');
		case "first":
			return s(1, a.first);
		case "first_icon":
			return s(1, '<i class="icon ' + o.options.firstIcon + '"></i>');
		case "last":
			return s(n.totalPage, a.last);
		case "last_icon":
			return s(n.totalPage, '<i class="icon ' + o.options.lastIcon + '"></i>');
		case "space":
		case "|":
			return t('<li class="space" />');
		case "nav":
		case "pages":
			return void o.createNavItems();
		case "total_text":
			return t(('<div class="pager-label">' + a.totalCount + "</div>").format(n));
		case "page_text":
			return t(('<div class="pager-label">' + a.pageOf + "</div>").format(n));
		case "total_page_text":
			return t(('<div class="pager-label">' + a.totalPage + "</div>").format(n));
		case "page_of_total_text":
			return t(('<div class="pager-label">' + a.pageOfTotal + "</div>").format(n));
		case "page_size_text":
			return t(('<div class="pager-label">' + a.pageSize + "</div>").format(n));
		case "items_range_text":
			return t(('<div class="pager-label">' + a.itemsRange + "</div>").format(n));
		case "goto":
			return o.createGoto();
		case "size_menu":
			return o.createSizeMenu();
		default:
			return t("<li/>").html(e.format(n))
		}
	}, s.prototype.createLink = function(i, n) {
		i === e && (i = this.state.page), n === e && (n = this.state);
		var o = this.options.linkCreator;
		return "string" == typeof o ? o.format(t.extend({}, n, {
			page: i
		})) : t.isFunction(o) ? o(i, n) : "#page=" + i
	}, s.prototype.render = function(e) {
		var i = this,
			n = i.state,
			o = i.options.elementCreator || i.createElement,
			s = t.isPlainObject(o);
		e = e || i.elements || i.options.elements, "string" == typeof e && (e = e.split(",")), i.elements = e, i.$.empty();
		for (var a = 0; a < e.length; ++a) {
			var r = t.trim(e[a]),
				l = s ? o[r] || o : o,
				h = l.call(i, r, i.$, n);
			h === !1 && (h = i.createElement(r, i.$, n)), h instanceof t && ("LI" !== h[0].tagName && (h = t("<li/>").append(h)), i.$.append(h))
		}
		var c = null;
		return i.$.children("li").each(function() {
			var e = t(this),
				i = !! e.children(".pager-item").length;
			c ? c.toggleClass("pager-item-right", !i) : i && e.addClass("pager-item-left"), c = i ? e : null
		}), c && c.addClass("pager-item-right"), i.$.callComEvent(i, "onRender", [n]), i
	}, s.DEFAULTS = t.extend({
		elements: ["first_icon", "prev_icon", "pages", "next_icon", "last_icon", "page_of_total_text", "items_range_text", "total_text"],
		prevIcon: "icon-double-angle-left",
		nextIcon: "icon-double-angle-right",
		firstIcon: "icon-step-backward",
		lastIcon: "icon-step-forward",
		navEllipsisItem: '<i class="icon icon-ellipsis-h"></i>',
		maxNavCount: 10,
		menuDirection: "dropdown",
		pageSizeOptions: [10, 20, 30, 50, 100]
	}, n), t.fn.pager = function(e) {
		return this.each(function() {
			var n = t(this),
				o = n.data(i),
				a = "object" == typeof e && e;
			o || n.data(i, o = new s(this, a)), "string" == typeof e && o[e]()
		})
	}, s.NAME = i, t.fn.pager.Constructor = s, t(function() {
		t('[data-ride="pager"]').pager()
	})
}(jQuery, void 0), +
function(t) {
	"use strict";
	var e = "zui.tab",
		i = function(e) {
			this.element = t(e)
		};
	i.prototype.show = function() {
		var i = this.element,
			n = i.closest("ul:not(.dropdown-menu)"),
			o = i.attr("data-target") || i.attr("data-tab");
		if (o || (o = i.attr("href"), o = o && o.replace(/.*(?=#[^\s]*$)/, "")), !i.parent("li").hasClass("active")) {
			var s = n.find(".active:last a")[0],
				a = t.Event("show." + e, {
					relatedTarget: s
				});
			if (i.trigger(a), !a.isDefaultPrevented()) {
				var r = t(o);
				this.activate(i.parent("li"), n), this.activate(r, r.parent(), function() {
					i.trigger({
						type: "shown." + e,
						relatedTarget: s
					})
				})
			}
		}
	}, i.prototype.activate = function(e, i, n) {
		function o() {
			s.removeClass("active").find("> .dropdown-menu > .active").removeClass("active"), e.addClass("active"), a ? (e[0].offsetWidth, e.addClass("in")) : e.removeClass("fade"), e.parent(".dropdown-menu") && e.closest("li.dropdown").addClass("active"), n && n()
		}
		var s = i.find("> .active"),
			a = n && t.support.transition && s.hasClass("fade");
		a ? s.one(t.support.transition.end, o).emulateTransitionEnd(150) : o(), s.removeClass("in")
	};
	var n = t.fn.tab;
	t.fn.tab = function(n) {
		return this.each(function() {
			var o = t(this),
				s = o.data(e);
			s || o.data(e, s = new i(this)), "string" == typeof n && s[n]()
		})
	}, t.fn.tab.Constructor = i, t.fn.tab.noConflict = function() {
		return t.fn.tab = n, this
	}, t(document).on("click.zui.tab.data-api", '[data-toggle="tab"], [data-tab]', function(e) {
		e.preventDefault(), t(this).tab("show")
	})
}(window.jQuery), +
function(t) {
	"use strict";

	function e() {
		var t = document.createElement("bootstrap"),
			e = {
				WebkitTransition: "webkitTransitionEnd",
				MozTransition: "transitionend",
				OTransition: "oTransitionEnd otransitionend",
				transition: "transitionend"
			};
		for (var i in e) if (void 0 !== t.style[i]) return {
			end: e[i]
		};
		return !1
	}
	t.fn.emulateTransitionEnd = function(e) {
		var i = !1,
			n = this;
		t(this).one("bsTransitionEnd", function() {
			i = !0
		});
		var o = function() {
				i || t(n).trigger(t.support.transition.end)
			};
		return setTimeout(o, e), this
	}, t(function() {
		t.support.transition = e(), t.support.transition && (t.event.special.bsTransitionEnd = {
			bindType: t.support.transition.end,
			delegateType: t.support.transition.end,
			handle: function(e) {
				if (t(e.target).is(this)) return e.handleObj.handler.apply(this, arguments)
			}
		})
	})
}(jQuery), +
function(t) {
	"use strict";
	var e = "zui.collapse",
		i = function(e, n) {
			this.$element = t(e), this.options = t.extend({}, i.DEFAULTS, n), this.transitioning = null, this.options.parent && (this.$parent = t(this.options.parent)), this.options.toggle && this.toggle()
		};
	i.DEFAULTS = {
		toggle: !0
	}, i.prototype.dimension = function() {
		var t = this.$element.hasClass("width");
		return t ? "width" : "height"
	}, i.prototype.show = function() {
		if (!this.transitioning && !this.$element.hasClass("in")) {
			var i = t.Event("show." + e);
			if (this.$element.trigger(i), !i.isDefaultPrevented()) {
				var n = this.$parent && this.$parent.find(".in");
				if (n && n.length) {
					var o = n.data(e);
					if (o && o.transitioning) return;
					n.collapse("hide"), o || n.data(e, null)
				}
				var s = this.dimension();
				this.$element.removeClass("collapse").addClass("collapsing")[s](0), this.transitioning = 1;
				var a = function() {
						this.$element.removeClass("collapsing").addClass("in")[s]("auto"), this.transitioning = 0, this.$element.trigger("shown." + e)
					};
				if (!t.support.transition) return a.call(this);
				var r = t.camelCase(["scroll", s].join("-"));
				this.$element.one(t.support.transition.end, t.proxy(a, this)).emulateTransitionEnd(350)[s](this.$element[0][r])
			}
		}
	}, i.prototype.hide = function() {
		if (!this.transitioning && this.$element.hasClass("in")) {
			var i = t.Event("hide." + e);
			if (this.$element.trigger(i), !i.isDefaultPrevented()) {
				var n = this.dimension();
				this.$element[n](this.$element[n]())[0].offsetHeight, this.$element.addClass("collapsing").removeClass("collapse").removeClass("in"), this.transitioning = 1;
				var o = function() {
						this.transitioning = 0, this.$element.trigger("hidden." + e).removeClass("collapsing").addClass("collapse")
					};
				return t.support.transition ? void this.$element[n](0).one(t.support.transition.end, t.proxy(o, this)).emulateTransitionEnd(350) : o.call(this)
			}
		}
	}, i.prototype.toggle = function() {
		this[this.$element.hasClass("in") ? "hide" : "show"]()
	};
	var n = t.fn.collapse;
	t.fn.collapse = function(n) {
		return this.each(function() {
			var o = t(this),
				s = o.data(e),
				a = t.extend({}, i.DEFAULTS, o.data(), "object" == typeof n && n);
			s || o.data(e, s = new i(this, a)), "string" == typeof n && s[n]()
		})
	}, t.fn.collapse.Constructor = i, t.fn.collapse.noConflict = function() {
		return t.fn.collapse = n, this
	}, t(document).on("click." + e + ".data-api", "[data-toggle=collapse]", function(i) {
		var n, o = t(this),
			s = o.attr("data-target") || i.preventDefault() || (n = o.attr("href")) && n.replace(/.*(?=#[^\s]+$)/, ""),
			a = t(s),
			r = a.data(e),
			l = r ? "toggle" : o.data(),
			h = o.attr("data-parent"),
			c = h && t(h);
		r && r.transitioning || (c && c.find('[data-toggle=collapse][data-parent="' + h + '"]').not(o).addClass("collapsed"), o[a.hasClass("in") ? "addClass" : "removeClass"]("collapsed")), a.collapse(l)
	})
}(window.jQuery), function(t, e) {
	"use strict";
	var i = 1200,
		n = 992,
		o = 768,
		s = e(t),
		a = function() {
			var t = s.width();
			e("html").toggleClass("screen-desktop", t >= n && t < i).toggleClass("screen-desktop-wide", t >= i).toggleClass("screen-tablet", t >= o && t < n).toggleClass("screen-phone", t < o).toggleClass("device-mobile", t < n).toggleClass("device-desktop", t >= n)
		},
		r = "",
		l = navigator.userAgent;
	l.match(/(iPad|iPhone|iPod)/i) ? r += " os-ios" : l.match(/android/i) ? r += " os-android" : l.match(/Win/i) ? r += " os-windows" : l.match(/Mac/i) ? r += " os-mac" : l.match(/Linux/i) ? r += " os-linux" : l.match(/X11/i) && (r += " os-unix"), "ontouchstart" in document.documentElement && (r += " is-touchable"), e("html").addClass(r), s.resize(a), a()
}(window, jQuery), function(t) {
	"use strict";
	var e = {
		zh_cn: '您的浏览器版本过低，无法体验所有功能，建议升级或者更换浏览器。 <a href="http://browsehappy.com/" target="_blank" class="alert-link">了解更多...</a>',
		zh_tw: '您的瀏覽器版本過低，無法體驗所有功能，建議升級或者更换瀏覽器。<a href="http://browsehappy.com/" target="_blank" class="alert-link">了解更多...</a>',
		en: 'Your browser is too old, it has been unable to experience the colorful internet. We strongly recommend that you upgrade a better one. <a href="http://browsehappy.com/" target="_blank" class="alert-link">Learn more...</a>'
	},
		i = function() {
			var t = this.isIE() || this.isIE10() || !1;
			if (t) for (var e = 10; e > 5; e--) if (this.isIE(e)) {
				t = e;
				break
			}
			this.ie = t, this.cssHelper()
		};
	i.prototype.cssHelper = function() {
		var e = this.ie,
			i = t("html");
		i.toggleClass("ie", e).removeClass("ie-6 ie-7 ie-8 ie-9 ie-10"), e && i.addClass("ie-" + e).toggleClass("gt-ie-7 gte-ie-8 support-ie", e >= 8).toggleClass("lte-ie-7 lt-ie-8 outdated-ie", e < 8).toggleClass("gt-ie-8 gte-ie-9", e >= 9).toggleClass("lte-ie-8 lt-ie-9", e < 9).toggleClass("gt-ie-9 gte-ie-10", e >= 10).toggleClass("lte-ie-9 lt-ie-10", e < 10)
	}, i.prototype.tip = function(i) {
		var n = t("#browseHappyTip");
		n.length || (n = t('<div id="browseHappyTip" class="alert alert-dismissable alert-danger-inverse alert-block" style="position: relative; z-index: 99999"><button type="button" class="close" data-dismiss="alert" aria-hidden="true">×</button><div class="container"><div class="content text-center"></div></div></div>'), n.prependTo("body")), n.find(".content").html(i || this.browseHappyTip || e[t.zui.clientLang() || "zh_cn"])
	}, i.prototype.isIE = function(t) {
		if (10 === t) return this.isIE10();
		var e = document.createElement("b");
		return e.innerHTML = "<!--[if IE " + (t || "") + "]><i></i><![endif]-->", 1 === e.getElementsByTagName("i").length
	}, i.prototype.isIE10 = function() {
		return !1
	}, t.zui({
		browser: new i
	}), t(function() {
		t("body").hasClass("disabled-browser-tip") || t.zui.browser.ie && t.zui.browser.ie < 8 && t.zui.browser.tip()
	})
}(jQuery), function() {
	"use strict";
	Date.ONEDAY_TICKS = 864e5, Date.prototype.format || (Date.prototype.format = function(t) {
		var e = {
			"M+": this.getMonth() + 1,
			"d+": this.getDate(),
			"h+": this.getHours(),
			"m+": this.getMinutes(),
			"s+": this.getSeconds(),
			"q+": Math.floor((this.getMonth() + 3) / 3),
			"S+": this.getMilliseconds()
		};
		/(y+)/i.test(t) && (t = t.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length)));
		for (var i in e) new RegExp("(" + i + ")").test(t) && (t = t.replace(RegExp.$1, 1 == RegExp.$1.length ? e[i] : ("00" + e[i]).substr(("" + e[i]).length)));
		return t
	}), Date.prototype.addMilliseconds || (Date.prototype.addMilliseconds = function(t) {
		return this.setTime(this.getTime() + t), this
	}), Date.prototype.addDays || (Date.prototype.addDays = function(t) {
		return this.addMilliseconds(t * Date.ONEDAY_TICKS), this
	}), Date.prototype.clone || (Date.prototype.clone = function() {
		var t = new Date;
		return t.setTime(this.getTime()), t
	}), Date.isLeapYear || (Date.isLeapYear = function(t) {
		return t % 4 === 0 && t % 100 !== 0 || t % 400 === 0
	}), Date.getDaysInMonth || (Date.getDaysInMonth = function(t, e) {
		return [31, Date.isLeapYear(t) ? 29 : 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31][e]
	}), Date.prototype.isLeapYear || (Date.prototype.isLeapYear = function() {
		return Date.isLeapYear(this.getFullYear())
	}), Date.prototype.clearTime || (Date.prototype.clearTime = function() {
		return this.setHours(0), this.setMinutes(0), this.setSeconds(0), this.setMilliseconds(0), this
	}), Date.prototype.getDaysInMonth || (Date.prototype.getDaysInMonth = function() {
		return Date.getDaysInMonth(this.getFullYear(), this.getMonth())
	}), Date.prototype.addMonths || (Date.prototype.addMonths = function(t) {
		var e = this.getDate();
		return this.setDate(1), this.setMonth(this.getMonth() + t), this.setDate(Math.min(e, this.getDaysInMonth())), this
	}), Date.prototype.getLastWeekday || (Date.prototype.getLastWeekday = function(t) {
		t = t || 1;
		for (var e = this.clone(); e.getDay() != t;) e.addDays(-1);
		return e.clearTime(), e
	}), Date.prototype.isSameDay || (Date.prototype.isSameDay = function(t) {
		return t.toDateString() === this.toDateString()
	}), Date.prototype.isSameWeek || (Date.prototype.isSameWeek = function(t) {
		var e = this.getLastWeekday(),
			i = e.clone().addDays(7);
		return t >= e && t < i
	}), Date.prototype.isSameYear || (Date.prototype.isSameYear = function(t) {
		return this.getFullYear() === t.getFullYear()
	}), Date.create || (Date.create = function(t) {
		return t instanceof Date || ("number" == typeof t && t < 1e10 && (t *= 1e3), t = new Date(t)), t
	}), Date.timestamp || (Date.timestamp = function(t) {
		return "number" == typeof t ? t < 1e10 && (t *= 1e3) : t = Date.create(t).getTime(), t
	})
}(), function() {
	"use strict";
	String.prototype.format || (String.prototype.format = function(t) {
		var e = this;
		if (arguments.length > 0) {
			var i;
			if (arguments.length <= 2 && "object" == typeof t) for (var n in t) void 0 !== t[n] && (i = new RegExp("(" + (arguments[1] ? arguments[1].replace("0", n) : "{" + n + "}") + ")", "g"), e = e.replace(i, t[n]));
			else for (var o = 0; o < arguments.length; o++) void 0 !== arguments[o] && (i = new RegExp("({[" + o + "]})", "g"), e = e.replace(i, arguments[o]))
		}
		return e
	}), String.prototype.isNum || (String.prototype.isNum = function(t) {
		if (null !== t) {
			var e, i;
			return i = /\d*/i, e = t.match(i), e == t
		}
		return !1
	}), String.prototype.endsWith || (String.prototype.endsWith = function(t, e) {
		var i = this.toString();
		(void 0 === e || e > i.length) && (e = i.length), e -= t.length;
		var n = i.indexOf(t, e);
		return n !== -1 && n === e
	}), String.prototype.startsWith || (String.prototype.startsWith = function(t, e) {
		return e = e || 0, this.lastIndexOf(t, e) === e
	}), String.prototype.includes || (String.prototype.includes = function() {
		return String.prototype.indexOf.apply(this, arguments) !== -1
	})
}(),
/*!
 * jQuery resize event - v1.1
 * http://benalman.com/projects/jquery-resize-plugin/
 * Copyright (c) 2010 "Cowboy" Ben Alman
 * MIT & GPL http://benalman.com/about/license/
 */

function(t, e, i) {
	"$:nomunge";

	function n() {
		o = e[r](function() {
			s.each(function() {
				var e = t(this),
					i = e.width(),
					n = e.height(),
					o = t.data(this, h);
				i === o.w && n === o.h || e.trigger(l, [o.w = i, o.h = n])
			}), n()
		}, a[c])
	}
	var o, s = t([]),
		a = t.resize = t.extend(t.resize, {}),
		r = "setTimeout",
		l = "resize",
		h = l + "-special-event",
		c = "delay",
		d = "throttleWindow";
	a[c] = 250, a[d] = !0, t.event.special[l] = {
		setup: function() {
			if (!a[d] && this[r]) return !1;
			var e = t(this);
			s = s.add(e), t.data(this, h, {
				w: e.width(),
				h: e.height()
			}), 1 === s.length && n()
		},
		teardown: function() {
			if (!a[d] && this[r]) return !1;
			var e = t(this);
			s = s.not(e), e.removeData(h), s.length || clearTimeout(o)
		},
		add: function(e) {
			function n(e, n, s) {
				var a = t(this),
					r = t.data(this, h) || {};
				r.w = n !== i ? n : a.width(), r.h = s !== i ? s : a.height(), o.apply(this, arguments)
			}
			if (!a[d] && this[r]) return !1;
			var o;
			return t.isFunction(e) ? (o = e, n) : (o = e.handler, void(e.handler = n))
		}
	}
}(jQuery, this),
/*!
 * jQuery Cookie Plugin v1.4.1
 * https://github.com/carhartl/jquery-cookie
 * Copyright 2006, 2014 Klaus Hartl
 * Released under the MIT license
 */

function(t) {
	"function" == typeof define && define.amd ? define(["jquery"], t) : t("object" == typeof exports ? require("jquery") : jQuery)
}(function(t) {
	function e(t) {
		return r.raw ? t : encodeURIComponent(t)
	}
	function i(t) {
		return r.raw ? t : decodeURIComponent(t)
	}
	function n(t) {
		return e(r.json ? JSON.stringify(t) : String(t))
	}
	function o(t) {
		0 === t.indexOf('"') && (t = t.slice(1, -1).replace(/\\"/g, '"').replace(/\\\\/g, "\\"));
		try {
			return t = decodeURIComponent(t.replace(a, " ")), r.json ? JSON.parse(t) : t
		} catch (e) {}
	}
	function s(e, i) {
		var n = r.raw ? e : o(e);
		return t.isFunction(i) ? i(n) : n
	}
	var a = /\+/g,
		r = t.cookie = function(o, a, l) {
			if (void 0 !== a && !t.isFunction(a)) {
				if (l = t.extend({}, r.defaults, l), "number" == typeof l.expires) {
					var h = l.expires,
						c = l.expires = new Date;
					c.setTime(+c + 864e5 * h)
				}
				return document.cookie = [e(o), "=", n(a), l.expires ? "; expires=" + l.expires.toUTCString() : "", l.path ? "; path=" + l.path : "", l.domain ? "; domain=" + l.domain : "", l.secure ? "; secure" : ""].join("")
			}
			for (var d = o ? void 0 : {}, u = document.cookie ? document.cookie.split("; ") : [], p = 0, f = u.length; p < f; p++) {
				var g = u[p].split("="),
					m = i(g.shift()),
					v = g.join("=");
				if (o && o === m) {
					d = s(v, a);
					break
				}
				o || void 0 === (v = s(v)) || (d[m] = v)
			}
			return d
		};
	r.defaults = {}, t.removeCookie = function(e, i) {
		return void 0 !== t.cookie(e) && (t.cookie(e, "", t.extend({}, i, {
			expires: -1
		})), !t.cookie(e))
	}
}), function(t, e) {
	"use strict";
	var i, n, o = "localStorage",
		s = "page_" + t.location.pathname + t.location.search,
		a = function() {
			this.slience = !0;
			try {
				o in t && t[o] && t[o].setItem && (this.enable = !0, i = t[o])
			} catch (a) {}
			this.enable || (n = {}, i = {
				getLength: function() {
					var t = 0;
					return e.each(n, function() {
						t++
					}), t
				},
				key: function(t) {
					var i, o = 0;
					return e.each(n, function(e) {
						return o === t ? (i = e, !1) : void o++
					}), i
				},
				removeItem: function(t) {
					delete n[t]
				},
				getItem: function(t) {
					return n[t]
				},
				setItem: function(t, e) {
					n[t] = e
				},
				clear: function() {
					n = {}
				}
			}), this.storage = i, this.page = this.get(s, {})
		};
	a.prototype.pageSave = function() {
		if (e.isEmptyObject(this.page)) this.remove(s);
		else {
			var t, i = [];
			for (t in this.page) {
				var n = this.page[t];
				null === n && i.push(t)
			}
			for (t = i.length - 1; t >= 0; t--) delete this.page[i[t]];
			this.set(s, this.page)
		}
	}, a.prototype.pageRemove = function(t) {
		"undefined" != typeof this.page[t] && (this.page[t] = null, this.pageSave())
	}, a.prototype.pageClear = function() {
		this.page = {}, this.pageSave()
	}, a.prototype.pageGet = function(t, e) {
		var i = this.page[t];
		return void 0 === e || null !== i && void 0 !== i ? i : e
	}, a.prototype.pageSet = function(t, i) {
		e.isPlainObject(t) ? e.extend(!0, this.page, t) : this.page[this.serialize(t)] = i, this.pageSave()
	}, a.prototype.check = function() {
		if (!this.enable && !this.slience) throw new Error("Browser not support localStorage or enable status been set true.");
		return this.enable
	}, a.prototype.length = function() {
		return this.check() ? i.getLength ? i.getLength() : i.length : 0
	}, a.prototype.removeItem = function(t) {
		return i.removeItem(t), this
	}, a.prototype.remove = function(t) {
		return this.removeItem(t)
	}, a.prototype.getItem = function(t) {
		return i.getItem(t)
	}, a.prototype.get = function(t, e) {
		var i = this.deserialize(this.getItem(t));
		return "undefined" != typeof i && null !== i || "undefined" == typeof e ? i : e
	}, a.prototype.key = function(t) {
		return i.key(t)
	}, a.prototype.setItem = function(t, e) {
		return i.setItem(t, e), this
	}, a.prototype.set = function(t, e) {
		return void 0 === e ? this.remove(t) : (this.setItem(t, this.serialize(e)), this)
	}, a.prototype.clear = function() {
		return i.clear(), this
	}, a.prototype.forEach = function(t) {
		for (var e = this.length(), n = e - 1; n >= 0; n--) {
			var o = i.key(n);
			t(o, this.get(o))
		}
		return this
	}, a.prototype.getAll = function() {
		var t = {};
		return this.forEach(function(e, i) {
			t[e] = i
		}), t
	}, a.prototype.serialize = function(t) {
		return "string" == typeof t ? t : JSON.stringify(t)
	}, a.prototype.deserialize = function(t) {
		if ("string" == typeof t) try {
			return JSON.parse(t)
		} catch (e) {
			return t || void 0
		}
	}, e.zui({
		store: new a
	})
}(window, jQuery), function(t) {
	"use strict";
	var e = "zui.searchBox",
		i = function(e, n) {
			var o = this;
			o.name = name, o.$ = t(e), o.options = n = t.extend({}, i.DEFAULTS, o.$.data(), n);
			var s = o.$.is(n.inputSelector) ? o.$ : o.$.find(n.inputSelector);
			if (s.length) {
				var a = function() {
						o.changeTimer && (clearTimeout(o.changeTimer), o.changeTimer = null)
					},
					r = function() {
						a();
						var t = o.getSearch();
						if (t !== o.lastValue) {
							var e = "" === t;
							s.toggleClass("empty", e), o.$.callComEvent(o, "onSearchChange", [t, e]), o.lastValue = t
						}
					};
				o.$input = s = s.first(), s.on(n.listenEvent, function(t) {
					o.changeTimer = setTimeout(function() {
						r()
					}, n.changeDelay)
				}).on("focus", function(t) {
					s.addClass("focus"), o.$.callComEvent(o, "onFocus", [t])
				}).on("blur", function(t) {
					s.removeClass("focus"), o.$.callComEvent(o, "onBlur", [t])
				}).on("keydown", function(t) {
					var e = 0,
						i = t.which;
					27 === i && n.escToClear ? (this.setSearch("", !0), r(), e = 1) : 13 === i && n.onPressEnter && (r(), o.$.callComEvent(o, "onPressEnter", [t]));
					var s = o.$.callComEvent(o, "onKeyDown", [t]);
					s === !1 && (e = 1), e && t.preventDefault()
				}), o.$.on("click", ".search-clear-btn", function(t) {
					o.setSearch("", !0), r(), o.focus(), t.preventDefault()
				}), r()
			} else console.error("ZUI: search box init error, cannot find search box input element.")
		};
	i.DEFAULTS = {
		inputSelector: 'input[type="search"],input[type="text"]',
		listenEvent: "change input paste",
		changeDelay: 500
	}, i.prototype.getSearch = function() {
		return this.$input && t.trim(this.$input.val())
	}, i.prototype.setSearch = function(t, e) {
		var i = this.$input;
		i && (i.val(t), e || i.trigger("change"))
	}, i.prototype.focus = function() {
		this.$input && this.$input.focus()
	}, t.fn.searchBox = function(n) {
		return this.each(function() {
			var o = t(this),
				s = o.data(e),
				a = "object" == typeof n && n;
			s || o.data(e, s = new i(this, a)), "string" == typeof n && s[n]()
		})
	}, i.NAME = e, t.fn.searchBox.Constructor = i
}(jQuery), function(t, e) {
	"use strict";
	var i = "zui.draggable",
		n = {
			container: "body",
			move: !0
		},
		o = 0,
		s = function(e, i) {
			var s = this;
			s.$ = t(e), s.id = o++, s.options = t.extend({}, n, s.$.data(), i), s.init()
		};
	s.DEFAULTS = n, s.NAME = i, s.prototype.init = function() {
		var n, o, s, a, r, l = this,
			h = l.$,
			c = "before",
			d = "drag",
			u = "finish",
			p = "." + i + "." + l.id,
			f = "mousedown" + p,
			g = "mouseup" + p,
			m = "mousemove" + p,
			v = l.options,
			y = v.selector,
			b = v.handle,
			w = h,
			x = t.isFunction(v.move),
			C = function(t) {
				var e = t.pageX,
					i = t.pageY;
				r = !0;
				var o = {
					left: e - s.x,
					top: i - s.y
				};
				w.removeClass("drag-ready").addClass("dragging"), v.move && (x ? v.move(o, w) : w.css(o)), v[d] && v[d]({
					event: t,
					element: w,
					startOffset: s,
					pos: o,
					offset: {
						x: e - n.x,
						y: i - n.y
					},
					smallOffset: {
						x: e - a.x,
						y: i - a.y
					}
				}), a.x = e, a.y = i, v.stopPropagation && t.stopPropagation()
			},
			_ = function(i) {
				if (t(e).off(p), !r) return void w.removeClass("drag-ready");
				var o = {
					left: i.pageX - s.x,
					top: i.pageY - s.y
				};
				w.removeClass("drag-ready dragging"), v.move && (x ? v.move(o, w) : w.css(o)), v[u] && v[u]({
					event: i,
					element: w,
					startOffset: s,
					pos: o,
					offset: {
						x: i.pageX - n.x,
						y: i.pageY - n.y
					},
					smallOffset: {
						x: i.pageX - a.x,
						y: i.pageY - a.y
					}
				}), i.preventDefault(), v.stopPropagation && i.stopPropagation()
			},
			k = function(i) {
				var l = t.zui.getMouseButtonCode(v.mouseButton);
				if (!(l > -1 && i.button !== l)) {
					var h = t(this);
					if (y && (w = b ? h.closest(y) : h), v[c]) {
						var d = v[c]({
							event: i,
							element: w
						});
						if (d === !1) return
					}
					var u = t(v.container),
						p = w.offset();
					o = u.offset(), n = {
						x: i.pageX,
						y: i.pageY
					}, s = {
						x: i.pageX - p.left + o.left,
						y: i.pageY - p.top + o.top
					}, a = t.extend({}, n), r = !1, w.addClass("drag-ready"), i.preventDefault(), v.stopPropagation && i.stopPropagation(), t(e).on(m, C).on(g, _)
				}
			};
		b ? h.on(f, b, k) : y ? h.on(f, y, k) : h.on(f, k)
	}, s.prototype.destroy = function() {
		var n = "." + i + "." + this.id;
		this.$.off(n), t(e).off(n), this.$.data(i, null)
	}, t.fn.draggable = function(e) {
		return this.each(function() {
			var n = t(this),
				o = n.data(i),
				a = "object" == typeof e && e;
			o || n.data(i, o = new s(this, a)), "string" == typeof e && o[e]()
		})
	}, t.fn.draggable.Constructor = s
}(jQuery, document), function(t, e, i) {
	"use strict";
	var n = "zui.droppable",
		o = {
			target: ".droppable-target",
			deviation: 5,
			sensorOffsetX: 0,
			sensorOffsetY: 0,
			dropToClass: "drop-to"
		},
		s = 0,
		a = function(e, i) {
			var n = this;
			n.id = s++, n.$ = t(e), n.options = t.extend({}, o, n.$.data(), i), n.init()
		};
	a.DEFAULTS = o, a.NAME = n, a.prototype.trigger = function(e, i) {
		return t.zui.callEvent(this.options[e], i, this)
	}, a.prototype.init = function() {
		var o, s, a, r, l, h, c, d, u, p, f, g, m, v = this,
			y = v.$,
			b = v.options,
			w = b.deviation,
			x = "." + n + "." + v.id,
			C = "mousedown" + x,
			_ = "mouseup" + x,
			k = "mousemove" + x,
			T = b.selector,
			S = b.handle,
			D = b.flex,
			M = b.container,
			P = b.canMoveHere,
			F = b.dropToClass,
			L = y,
			z = !1,
			I = M ? t(b.container).first() : T ? y : t("body"),
			$ = function(e) {
				if (z && (f = {
					left: e.pageX,
					top: e.pageY
				}, !(i.abs(f.left - d.left) < w && i.abs(f.top - d.top) < w))) {
					if (null === a) {
						var n = I.css("position");
						"absolute" != n && "relative" != n && "fixed" != n && (h = n, I.css("position", "relative")), a = L.clone().removeClass("drag-from").addClass("drag-shadow").css({
							position: "absolute",
							width: L.outerWidth(),
							transition: "none"
						}).appendTo(I), L.addClass("dragging"), v.trigger("start", {
							event: e,
							element: L,
							targets: o
						})
					}
					var c = {
						left: f.left - p.left,
						top: f.top - p.top
					},
						m = {
							left: c.left - u.left,
							top: c.top - u.top
						};
					a.css(m), t.extend(g, f);
					var y = !1;
					r = !1, D || o.removeClass(F);
					var x = null;
					if (o.each(function() {
						var e = t(this),
							i = e.offset(),
							n = e.outerWidth(),
							o = e.outerHeight(),
							s = i.left + b.sensorOffsetX,
							a = i.top + b.sensorOffsetY;
						if (f.left > s && f.top > a && f.left < s + n && f.top < a + o && (x && x.removeClass(F), x = e, !b.nested)) return !1
					}), x) {
						r = !0;
						var C = x.data("id");
						L.data("id") != C && (l = !1), (null === s || s.data("id") !== C && !l) && (y = !0), s = x, D && o.removeClass(F), s.addClass(F)
					}
					D ? null !== s && s.length && (r = !0) : (L.toggleClass("drop-in", r), a.toggleClass("drop-in", r)), P && P(L, s) === !1 || v.trigger("drag", {
						event: e,
						isIn: r,
						target: s,
						element: L,
						isNew: y,
						selfTarget: l,
						clickOffset: p,
						offset: c,
						position: {
							left: c.left - u.left,
							top: c.top - u.top
						},
						mouseOffset: f
					}), e.preventDefault()
				}
			},
			E = function(i) {
				if (t(e).off(x), clearTimeout(m), z) {
					if (z = !1, h && I.css("position", h), null === a) return L.removeClass("drag-from"), void v.trigger("always", {
						event: i,
						cancel: !0
					});
					r || (s = null);
					var n = !0;
					f = i ? {
						left: i.pageX,
						top: i.pageY
					} : g;
					var c = {
						left: f.left - p.left,
						top: f.top - p.top
					},
						d = {
							left: f.left - g.left,
							top: f.top - g.top
						};
					g.left = f.left, g.top = f.top;
					var y = {
						event: i,
						isIn: r,
						target: s,
						element: L,
						isNew: !l && null !== s,
						selfTarget: l,
						offset: c,
						mouseOffset: f,
						position: {
							left: c.left - u.left,
							top: c.top - u.top
						},
						lastMouseOffset: g,
						moveOffset: d
					};
					n = v.trigger("beforeDrop", y), n && r && v.trigger("drop", y), o.removeClass(F), L.removeClass("dragging").removeClass("drag-from"), a.remove(), a = null, v.trigger("finish", y), v.trigger("always", y), i && i.preventDefault()
				}
			},
			A = function(i) {
				var n = t.zui.getMouseButtonCode(b.mouseButton);
				if (!(n > -1 && i.button !== n)) {
					var f = t(this);
					T && (L = S ? f.closest(T) : f), L.hasClass("drag-shadow") || b.before && b.before({
						event: i,
						element: L
					}) === !1 || (z = !0, o = t.isFunction(b.target) ? b.target(L, y) : I.find(b.target), s = null, a = null, r = !1, l = !0, h = null, c = L.offset(), u = I.offset(), u.top = u.top - I.scrollTop(), u.left = u.left - I.scrollLeft(), d = {
						left: i.pageX,
						top: i.pageY
					}, g = t.extend({}, d), p = {
						left: d.left - c.left,
						top: d.top - c.top
					}, L.addClass("drag-from"), t(e).on(k, $).on(_, E), m = setTimeout(function() {
						t(e).on(C, E)
					}, 10), i.preventDefault(), b.stopPropagation && i.stopPropagation())
				}
			};
		S ? y.on(C, S, A) : T ? y.on(C, T, A) : y.on(C, A)
	}, a.prototype.destroy = function() {
		var i = "." + n + "." + this.id;
		this.$.off(i), t(e).off(i), this.$.data(n, null)
	}, a.prototype.reset = function() {
		this.destroy(), this.init()
	}, t.fn.droppable = function(e) {
		return this.each(function() {
			var i = t(this),
				o = i.data(n),
				s = "object" == typeof e && e;
			o || i.data(n, o = new a(this, s)), "string" == typeof e && o[e]()
		})
	}, t.fn.droppable.Constructor = a
}(jQuery, document, Math), +
function(t, e) {
	"use strict";

	function i(e, i, s) {
		return this.each(function() {
			var a = t(this),
				r = a.data(n),
				l = t.extend({}, o.DEFAULTS, a.data(), "object" == typeof e && e);
			r || a.data(n, r = new o(this, l)), "string" == typeof e ? r[e](i, s) : l.show && r.show(i, s)
		})
	}
	var n = "zui.modal",
		o = function(i, o) {
			var s = this;
			s.options = o, s.$body = t(document.body), s.$element = t(i), s.$backdrop = s.isShown = null, s.scrollbarWidth = 0, o.moveable === e && (s.options.moveable = s.$element.hasClass("modal-moveable")), o.remote && s.$element.find(".modal-content").load(o.remote, function() {
				s.$element.trigger("loaded." + n)
			})
		};
	o.VERSION = "3.2.0", o.TRANSITION_DURATION = 300, o.BACKDROP_TRANSITION_DURATION = 150, o.DEFAULTS = {
		backdrop: !0,
		keyboard: !0,
		show: !0,
		position: "fit"
	};
	var s = function(e, i) {
			var n = t(window);
			i.left = Math.max(0, Math.min(i.left, n.width() - e.outerWidth())), i.top = Math.max(0, Math.min(i.top, n.height() - e.outerHeight())), e.css(i)
		};
	o.prototype.toggle = function(t, e) {
		return this.isShown ? this.hide() : this.show(t, e)
	}, o.prototype.ajustPosition = function(i) {
		var o = this,
			a = o.options;
		if (i === e && (i = a.position), i !== e && null !== i) {
			t.isFunction(i) && (i = i(o));
			var r = o.$element.find(".modal-dialog"),
				l = t(window).height(),
				h = {
					maxHeight: "initial",
					overflow: "visible"
				},
				c = r.find(".modal-body").css(h);
			if (a.scrollInside) {
				var d = a.headerHeight;
				"number" != typeof d ? d = r.find(".modal-header").height() : t.isFunction(d) && (d = d($header)), h.maxHeight = l - d, c.outerHeight() > h.maxHeight && (h.overflow = "auto")
			}
			c.css(h);
			var u = Math.max(0, (l - r.outerHeight()) / 2);
			if ("fit" === i ? i = {
				top: u > 50 ? Math.floor(2 * u / 3) : u
			} : "center" === i ? i = {
				top: u
			} : t.isPlainObject(i) || (i = {
				top: i
			}), r.hasClass("modal-moveable")) {
				var p = null,
					f = a.rememberPos;
				f && (f === !0 ? p = o.$element.data("modal-pos") : t.zui.store && (p = t.zui.store.pageGet(n + ".rememberPos." + f))), i = t.extend(i, {
					left: Math.max(0, (t(window).width() - r.outerWidth()) / 2)
				}, p), "inside" === a.moveable ? s(r, i) : r.css(i)
			} else r.css(i)
		}
	}, o.prototype.setMoveale = function() {
		t.fn.draggable || console.error("Moveable modal requires draggable.js.");
		var e = this,
			i = e.options,
			o = e.$element.find(".modal-dialog").removeClass("modal-dragged");
		o.toggleClass("modal-moveable", !! i.moveable), e.$element.data("modal-moveable-setup") || o.draggable({
			container: e.$element,
			handle: ".modal-header",
			before: function() {
				var t = o.css("margin-top");
				t && "0px" !== t && o.css("top", t).css("margin-top", "").addClass("modal-dragged")
			},
			finish: function(o) {
				var s = i.rememberPos;
				s && (e.$element.data("modal-pos", o.pos), t.zui.store && s !== !0 && t.zui.store.pageSet(n + ".rememberPos." + s, o.pos))
			},
			move: "inside" !== i.moveable ||
			function(t) {
				s(o, t)
			}
		})
	}, o.prototype.show = function(e, i) {
		var s = this,
			a = t.Event("show." + n, {
				relatedTarget: e
			});
		s.$element.trigger(a), s.$element.toggleClass("modal-scroll-inside", !! s.options.scrollInside), s.isShown || a.isDefaultPrevented() || (s.isShown = !0, s.options.moveable && s.setMoveale(), s.checkScrollbar(), s.options.backdrop !== !1 && (s.$body.addClass("modal-open"), s.setScrollbar()), s.escape(), s.$element.on("click.dismiss." + n, '[data-dismiss="modal"]', function(t) {
			s.hide(), t.stopPropagation()
		}), s.backdrop(function() {
			var a = t.support.transition && s.$element.hasClass("fade");
			s.$element.parent().length || s.$element.appendTo(s.$body), s.$element.show().scrollTop(0), a && s.$element[0].offsetWidth, s.$element.addClass("in").attr("aria-hidden", !1), s.ajustPosition(i), s.enforceFocus();
			var r = t.Event("shown." + n, {
				relatedTarget: e
			});
			a ? s.$element.find(".modal-dialog").one("bsTransitionEnd", function() {
				s.$element.trigger("focus").trigger(r)
			}).emulateTransitionEnd(o.TRANSITION_DURATION) : s.$element.trigger("focus").trigger(r)
		}))
	}, o.prototype.hide = function(e) {
		e && e.preventDefault();
		var i = this;
		e = t.Event("hide." + n), i.$element.trigger(e), i.isShown && !e.isDefaultPrevented() && (i.isShown = !1, i.options.backdrop !== !1 && (i.$body.removeClass("modal-open"), i.resetScrollbar()), i.escape(), t(document).off("focusin." + n), i.$element.removeClass("in").attr("aria-hidden", !0).off("click.dismiss." + n), t.support.transition && i.$element.hasClass("fade") ? i.$element.one("bsTransitionEnd", t.proxy(i.hideModal, i)).emulateTransitionEnd(o.TRANSITION_DURATION) : i.hideModal())
	}, o.prototype.enforceFocus = function() {
		t(document).off("focusin." + n).on("focusin." + n, t.proxy(function(t) {
			this.$element[0] === t.target || this.$element.has(t.target).length || this.$element.trigger("focus")
		}, this))
	}, o.prototype.escape = function() {
		this.isShown && this.options.keyboard ? t(document).on("keydown.dismiss." + n, t.proxy(function(i) {
			if (27 == i.which) {
				var o = t.Event("escaping." + n),
					s = this.$element.triggerHandler(o, "esc");
				if (s != e && !s) return;
				this.hide()
			}
		}, this)) : this.isShown || t(document).off("keydown.dismiss." + n)
	}, o.prototype.hideModal = function() {
		var t = this;
		this.$element.hide(), this.backdrop(function() {
			t.$element.trigger("hidden." + n)
		})
	}, o.prototype.removeBackdrop = function() {
		this.$backdrop && this.$backdrop.remove(), this.$backdrop = null
	}, o.prototype.backdrop = function(e) {
		var i = this,
			s = this.$element.hasClass("fade") ? "fade" : "";
		if (this.isShown && this.options.backdrop) {
			var a = t.support.transition && s;
			if (this.$backdrop = t('<div class="modal-backdrop ' + s + '" />').appendTo(this.$body), this.$element.on("mousedown.dismiss." + n, t.proxy(function(t) {
				t.target === t.currentTarget && ("static" == this.options.backdrop ? this.$element[0].focus.call(this.$element[0]) : this.hide.call(this))
			}, this)), a && this.$backdrop[0].offsetWidth, this.$backdrop.addClass("in"), !e) return;
			a ? this.$backdrop.one("bsTransitionEnd", e).emulateTransitionEnd(o.BACKDROP_TRANSITION_DURATION) : e()
		} else if (!this.isShown && this.$backdrop) {
			this.$backdrop.removeClass("in");
			var r = function() {
					i.removeBackdrop(), e && e()
				};
			t.support.transition && this.$element.hasClass("fade") ? this.$backdrop.one("bsTransitionEnd", r).emulateTransitionEnd(o.BACKDROP_TRANSITION_DURATION) : r()
		} else e && e()
	}, o.prototype.checkScrollbar = function() {
		document.body.clientWidth >= window.innerWidth || (this.scrollbarWidth = this.scrollbarWidth || this.measureScrollbar())
	}, o.prototype.setScrollbar = function() {
		var t = parseInt(this.$body.css("padding-right") || 0, 10);
		this.scrollbarWidth && this.$body.css("padding-right", t + this.scrollbarWidth)
	}, o.prototype.resetScrollbar = function() {
		this.$body.css("padding-right", "")
	}, o.prototype.measureScrollbar = function() {
		var t = document.createElement("div");
		t.className = "modal-scrollbar-measure", this.$body.append(t);
		var e = t.offsetWidth - t.clientWidth;
		return this.$body[0].removeChild(t), e
	};
	var a = t.fn.modal;
	t.fn.modal = i, t.fn.modal.Constructor = o, t.fn.modal.noConflict = function() {
		return t.fn.modal = a, this
	}, t(document).on("click." + n + ".data-api", '[data-toggle="modal"]', function(e) {
		var o = t(this),
			s = o.attr("href"),
			a = null;
		try {
			a = t(o.attr("data-target") || s && s.replace(/.*(?=#[^\s]+$)/, ""))
		} catch (r) {
			return
		}
		if (a.length) {
			var l = a.data(n) ? "toggle" : t.extend({
				remote: !/#/.test(s) && s
			}, a.data(), o.data());
			o.is("a") && e.preventDefault(), a.one("show." + n, function(t) {
				t.isDefaultPrevented() || a.one("hidden." + n, function() {
					o.is(":visible") && o.trigger("focus")
				})
			}), i.call(a, l, this, o.data("position"))
		}
	})
}(jQuery, void 0), function(t, e, i) {
	"use strict";
	if (!t.fn.modal) throw new Error("Modal trigger requires modal.js");
	var n = "zui.modaltrigger",
		o = "ajax",
		s = ".zui.modal",
		a = "string",
		r = function(e, i) {
			e = t.extend({}, r.DEFAULTS, t.ModalTriggerDefaults, i ? i.data() : null, e), this.isShown, this.$trigger = i, this.options = e, this.id = t.zui.uuid()
		};
	r.DEFAULTS = {
		type: "custom",
		height: "auto",
		name: "triggerModal",
		fade: !0,
		position: "fit",
		showHeader: !0,
		delay: 0,
		backdrop: !0,
		keyboard: !0,
		waittime: 0,
		loadingIcon: "icon-spinner-indicator",
		scrollInside: !1
	}, r.prototype.init = function(i) {
		var r = this;
		if (i.url && (!i.type || i.type != o && "iframe" != i.type) && (i.type = o), i.remote) i.type = o, typeof i.remote === a && (i.url = i.remote);
		else if (i.iframe) i.type = "iframe", typeof i.iframe === a && (i.url = i.iframe);
		else if (i.custom && (i.type = "custom", typeof i.custom === a)) {
			var l;
			try {
				l = t(i.custom)
			} catch (h) {}
			l && l.length ? i.custom = l : t.isFunction(e[i.custom]) && (i.custom = e[i.custom])
		}
		var c = t("#" + i.name);
		c.length && (r.isShown || c.off(s), c.remove()), c = t('<div id="' + i.name + '" class="modal modal-trigger ' + (i.className || "") + '">' + ("string" == typeof i.loadingIcon && 0 === i.loadingIcon.indexOf("icon-") ? '<div class="icon icon-spin loader ' + i.loadingIcon + '"></div>' : i.loadingIcon) + '<div class="modal-dialog"><div class="modal-content"><div class="modal-header"><button class="close" data-dismiss="modal">×</button><h4 class="modal-title"><i class="modal-icon"></i> <span class="modal-title-name"></span></h4></div><div class="modal-body"></div></div></div></div>').appendTo("body").data(n, r);
		var d = function(e, n) {
				var o = i[e];
				t.isFunction(o) && c.on(n + s, o)
			};
		d("onShow", "show"), d("shown", "shown"), d("onHide", "hide"), d("hidden", "hidden"), d("loaded", "loaded"), c.on("shown" + s, function() {
			r.isShown = !0
		}).on("hidden" + s, function() {
			r.isShown = !1
		}), this.$modal = c, this.$dialog = c.find(".modal-dialog"), i.mergeOptions && (this.options = i)
	}, r.prototype.show = function(i) {
		var s = t.extend({}, this.options, {
			url: this.$trigger ? this.$trigger.attr("href") || this.$trigger.attr("data-url") || this.$trigger.data("url") : this.options.url
		}, i);
		this.init(s);
		var r = this,
			l = this.$modal,
			h = this.$dialog,
			c = s.custom,
			d = h.find(".modal-body").css("padding", ""),
			u = h.find(".modal-header"),
			p = h.find(".modal-content");
		l.toggleClass("fade", s.fade).addClass(s.className).toggleClass("modal-loading", !this.isShown).toggleClass("modal-scroll-inside", !! s.scrollInside), h.toggleClass("modal-md", "md" === s.size).toggleClass("modal-sm", "sm" === s.size).toggleClass("modal-lg", "lg" === s.size).toggleClass("modal-fullscreen", "fullscreen" === s.size), u.toggle(s.showHeader), u.find(".modal-icon").attr("class", "modal-icon icon-" + s.icon), u.find(".modal-title-name").text(s.title || ""), s.size && "fullscreen" === s.size && (s.width = "", s.height = "");
		var f = function() {
				clearTimeout(this.resizeTask), this.resizeTask = setTimeout(function() {
					r.ajustPosition(s.position)
				}, 100)
			},
			g = function(t, e) {
				return "undefined" == typeof t && (t = s.delay), setTimeout(function() {
					h = l.find(".modal-dialog"), s.width && "auto" != s.width && h.css("width", s.width), s.height && "auto" != s.height && (h.css("height", s.height), "iframe" === s.type && d.css("height", h.height() - u.outerHeight())), r.ajustPosition(s.position), l.removeClass("modal-loading"), "iframe" != s.type && h.off("resize." + n).on("resize." + n, f), e && e()
				}, t)
			};
		if ("custom" === s.type && c) if (t.isFunction(c)) {
			var m = c({
				modal: l,
				options: s,
				modalTrigger: r,
				ready: g
			});
			typeof m === a && (d.html(m), g())
		} else c instanceof t ? (d.html(t("<div>").append(c.clone()).html()), g()) : (d.html(c), g());
		else if (s.url) {
			var v = function() {
					var t = l.callComEvent(r, "broken");
					t && (d.html(t), g())
				};
			if (l.attr("ref", s.url), "iframe" === s.type) {
				l.addClass("modal-iframe"), this.firstLoad = !0;
				var y = "iframe-" + s.name;
				u.detach(), d.detach(), p.empty().append(u).append(d), d.css("padding", 0).html('<iframe id="' + y + '" name="' + y + '" src="' + s.url + '" frameborder="no"  allowfullscreen="true" mozallowfullscreen="true" webkitallowfullscreen="true"  allowtransparency="true" scrolling="auto" style="width: 100%; height: 100%; left: 0px;"></iframe>'), s.waittime > 0 && (r.waitTimeout = g(s.waittime, v));
				var b = document.getElementById(y);
				b.onload = b.onreadystatechange = function() {
					var i = !! s.scrollInside;
					if (r.firstLoad && l.addClass("modal-loading"), !this.readyState || "complete" == this.readyState) {
						r.firstLoad = !1, s.waittime > 0 && clearTimeout(r.waitTimeout);
						try {
							l.attr("ref", b.contentWindow.location.href);
							var o = e.frames[y].$;
							if (o && "auto" === s.height && "fullscreen" != s.size) {
								var a = o("body").addClass("body-modal").toggleClass("body-modal-scroll-inside", i);
								s.iframeBodyClass && a.addClass(s.iframeBodyClass);
								var h = [],
									c = function(n) {
										l.removeClass("fade");
										var o = a.outerHeight();
										if (n === !0 && s.onlyIncreaseHeight && (o = Math.max(o, d.data("minModalHeight") || 0), d.data("minModalHeight", o)), i) {
											var r = s.headerHeight;
											"number" != typeof r ? r = u.height() : t.isFunction(r) && (r = r(u));
											var c = t(e).height();
											o = Math.min(o, c - r)
										}
										for (h.length > 1 && o === h[0] && (o = Math.max(o, h[1])), h.push(o); h.length > 2;) h.shift();
										d.css("height", o), s.fade && l.addClass("fade"), g()
									};
								l.callComEvent(r, "loaded", {
									modalType: "iframe",
									jQuery: o
								}), setTimeout(c, 100), a.off("resize." + n).on("resize." + n, c), i && t(e).off("resize." + n).on("resize." + n, c)
							} else g()
						} catch (p) {
							g()
						}
					}
				}
			} else t.ajax(t.extend({
				url: s.url,
				success: function(i) {
					try {
						var n = t(i);
						n.filter(".modal-dialog").length ? h.replaceWith(n) : n.filter(".modal-content").length ? h.find(".modal-content").replaceWith(n) : d.wrapInner(n)
					} catch (s) {
						e.console && e.console.warn && console.warn("ZUI: Cannot recogernize remote content.", {
							error: s,
							data: i
						}), l.html(i)
					}
					l.callComEvent(r, "loaded", {
						modalType: o
					}), g()
				},
				error: v
			}, s.ajaxOptions))
		}
		l.modal({
			show: "show",
			backdrop: s.backdrop,
			moveable: s.moveable,
			rememberPos: s.rememberPos,
			keyboard: s.keyboard,
			scrollInside: s.scrollInside
		})
	}, r.prototype.close = function(i, n) {
		var o = this;
		(i || n) && o.$modal.on("hidden" + s, function() {
			t.isFunction(i) && i(), typeof n === a && n.length && !o.$modal.data("cancel-reload") && ("this" === n ? e.location.reload() : e.location = n)
		}), o.$modal.modal("hide")
	}, r.prototype.toggle = function(t) {
		this.isShown ? this.close() : this.show(t)
	}, r.prototype.ajustPosition = function(e) {
		e = e === i ? this.options.position : e, t.isFunction(e) && (e = e(this)), this.$modal.modal("ajustPosition", e)
	}, t.zui({
		ModalTrigger: r,
		modalTrigger: new r
	}), t.fn.modalTrigger = function(e, i) {
		return t(this).each(function() {
			var o = t(this),
				s = o.data(n),
				l = t.extend({
					title: o.attr("title") || o.text(),
					url: o.attr("href"),
					type: o.hasClass("iframe") ? "iframe" : ""
				}, o.data(), t.isPlainObject(e) && e);
			s || o.data(n, s = new r(l, o)), typeof e == a ? s[e](i) : l.show && s.show(i), o.on((l.trigger || "click") + ".toggle." + n, function(e) {
				l = t.extend(l, {
					url: o.attr("href") || o.attr("data-url") || o.data("url") || l.url
				}), s.toggle(l), o.is("a") && e.preventDefault()
			})
		})
	};
	var l = t.fn.modal;
	t.fn.modal = function(e, i) {
		return t(this).each(function() {
			var n = t(this);
			n.hasClass("modal") ? l.call(n, e, i) : n.modalTrigger(e, i)
		})
	}, t.fn.modal.bs = l;
	var h = function(e) {
			return e = t(e ? e : ".modal.modal-trigger"), e && e instanceof t ? e : null
		},
		c = function(i, o, s) {
			var a = i;
			if (t.isFunction(i)) {
				var r = s;
				s = o, o = i, i = r
			}
			i = h(i), i && i.length ? i.each(function() {
				t(this).data(n).close(o, s)
			}) : t("body").hasClass("modal-open") || t(".modal.in").length || t("body").hasClass("body-modal") && e.parent.$.zui.closeModal(a, o, s)
		},
		d = function(t, e) {
			e = h(e), e && e.length && e.modal("ajustPosition", t)
		};
	t.zui({
		closeModal: c,
		ajustModalPosition: d
	}), t(document).on("click." + n + ".data-api", '[data-toggle="modal"]', function(e) {
		var i = t(this),
			o = i.attr("href"),
			s = null;
		try {
			s = t(i.attr("data-target") || o && o.replace(/.*(?=#[^\s]+$)/, ""))
		} catch (a) {}
		s && s.length || (i.data(n) ? i.trigger(".toggle." + n) : i.modalTrigger({
			show: !0
		})), i.is("a") && e.preventDefault()
	}).on("click." + n + ".data-api", '[data-dismiss="modal"]', function() {
		t.zui.closeModal()
	})
}(window.jQuery, window, void 0), +
function(t) {
	"use strict";
	var e = function(t, e) {
			this.type = null, this.options = null, this.enabled = null, this.timeout = null, this.hoverState = null, this.$element = null, this.init("tooltip", t, e)
		};
	e.DEFAULTS = {
		animation: !0,
		placement: "top",
		selector: !1,
		template: '<div class="tooltip"><div class="tooltip-arrow"></div><div class="tooltip-inner"></div></div>',
		trigger: "hover focus",
		title: "",
		delay: 0,
		html: !1,
		container: !1
	}, e.prototype.init = function(e, i, n) {
		this.enabled = !0, this.type = e, this.$element = t(i), this.options = this.getOptions(n);
		for (var o = this.options.trigger.split(" "), s = o.length; s--;) {
			var a = o[s];
			if ("click" == a) this.$element.on("click." + this.type, this.options.selector, t.proxy(this.toggle, this));
			else if ("manual" != a) {
				var r = "hover" == a ? "mouseenter" : "focus",
					l = "hover" == a ? "mouseleave" : "blur";
				this.$element.on(r + "." + this.type, this.options.selector, t.proxy(this.enter, this)), this.$element.on(l + "." + this.type, this.options.selector, t.proxy(this.leave, this))
			}
		}
		this.options.selector ? this._options = t.extend({}, this.options, {
			trigger: "manual",
			selector: ""
		}) : this.fixTitle()
	}, e.prototype.getDefaults = function() {
		return e.DEFAULTS
	}, e.prototype.getOptions = function(e) {
		return e = t.extend({}, this.getDefaults(), this.$element.data(), e), e.delay && "number" == typeof e.delay && (e.delay = {
			show: e.delay,
			hide: e.delay
		}), e
	}, e.prototype.getDelegateOptions = function() {
		var e = {},
			i = this.getDefaults();
		return this._options && t.each(this._options, function(t, n) {
			i[t] != n && (e[t] = n)
		}), e
	}, e.prototype.enter = function(e) {
		var i = e instanceof this.constructor ? e : t(e.currentTarget)[this.type](this.getDelegateOptions()).data("zui." + this.type);
		return clearTimeout(i.timeout), i.hoverState = "in", i.options.delay && i.options.delay.show ? void(i.timeout = setTimeout(function() {
			"in" == i.hoverState && i.show()
		}, i.options.delay.show)) : i.show()
	}, e.prototype.leave = function(e) {
		var i = e instanceof this.constructor ? e : t(e.currentTarget)[this.type](this.getDelegateOptions()).data("zui." + this.type);
		return clearTimeout(i.timeout), i.hoverState = "out", i.options.delay && i.options.delay.hide ? void(i.timeout = setTimeout(function() {
			"out" == i.hoverState && i.hide()
		}, i.options.delay.hide)) : i.hide()
	}, e.prototype.show = function(e) {
		var i = t.Event("show.zui." + this.type);
		if ((e || this.hasContent()) && this.enabled) {
			var n = this;
			if (n.$element.trigger(i), i.isDefaultPrevented()) return;
			var o = n.tip();
			n.setContent(e), n.options.animation && o.addClass("fade");
			var s = "function" == typeof n.options.placement ? n.options.placement.call(n, o[0], n.$element[0]) : n.options.placement,
				a = /\s?auto?\s?/i,
				r = a.test(s);
			r && (s = s.replace(a, "") || "top"), o.detach().css({
				top: 0,
				left: 0,
				display: "block"
			}).addClass(s), n.options.container ? o.appendTo(n.options.container) : o.insertAfter(n.$element);
			var l = n.getPosition(),
				h = o[0].offsetWidth,
				c = o[0].offsetHeight;
			if (r) {
				var d = n.$element.parent(),
					u = s,
					p = document.documentElement.scrollTop || document.body.scrollTop,
					f = "body" == n.options.container ? window.innerWidth : d.outerWidth(),
					g = "body" == n.options.container ? window.innerHeight : d.outerHeight(),
					m = "body" == n.options.container ? 0 : d.offset().left;
				s = "bottom" == s && l.top + l.height + c - p > g ? "top" : "top" == s && l.top - p - c < 0 ? "bottom" : "right" == s && l.right + h > f ? "left" : "left" == s && l.left - h < m ? "right" : s, o.removeClass(u).addClass(s)
			}
			var v = n.getCalculatedOffset(s, l, h, c);
			n.applyPlacement(v, s);
			var y = function() {
					var t = n.hoverState;
					n.$element.trigger("shown.zui." + n.type), n.hoverState = null, "out" == t && n.leave(n)
				};
			t.support.transition && n.$tip.hasClass("fade") ? o.one("bsTransitionEnd", y).emulateTransitionEnd(150) : y()
		}
	}, e.prototype.applyPlacement = function(t, e) {
		var i, n = this.tip(),
			o = n[0].offsetWidth,
			s = n[0].offsetHeight,
			a = parseInt(n.css("margin-top"), 10),
			r = parseInt(n.css("margin-left"), 10);
		isNaN(a) && (a = 0), isNaN(r) && (r = 0), t.top = t.top + a, t.left = t.left + r, n.offset(t).addClass("in");
		var l = n[0].offsetWidth,
			h = n[0].offsetHeight;
		if ("top" == e && h != s && (i = !0, t.top = t.top + s - h), /bottom|top/.test(e)) {
			var c = 0;
			t.left < 0 && (c = t.left * -2, t.left = 0, n.offset(t), l = n[0].offsetWidth, h = n[0].offsetHeight), this.replaceArrow(c - o + l, l, "left")
		} else this.replaceArrow(h - s, h, "top");
		i && n.offset(t)
	}, e.prototype.replaceArrow = function(t, e, i) {
		this.arrow().css(i, t ? 50 * (1 - t / e) + "%" : "")
	}, e.prototype.setContent = function(t) {
		var e = this.tip(),
			i = t || this.getTitle();
		this.options.tipId && e.attr("id", this.options.tipId), this.options.tipClass && e.addClass(this.options.tipClass), e.find(".tooltip-inner")[this.options.html ? "html" : "text"](i), e.removeClass("fade in top bottom left right")
	}, e.prototype.hide = function() {
		function e() {
			"in" != i.hoverState && n.detach()
		}
		var i = this,
			n = this.tip(),
			o = t.Event("hide.zui." + this.type);
		if (this.$element.trigger(o), !o.isDefaultPrevented()) return n.removeClass("in"), t.support.transition && this.$tip.hasClass("fade") ? n.one(t.support.transition.end, e).emulateTransitionEnd(150) : e(), this.$element.trigger("hidden.zui." + this.type), this
	}, e.prototype.fixTitle = function() {
		var t = this.$element;
		(t.attr("title") || "string" != typeof t.attr("data-original-title")) && t.attr("data-original-title", t.attr("title") || "").attr("title", "")
	}, e.prototype.hasContent = function() {
		return this.getTitle()
	}, e.prototype.getPosition = function() {
		var e = this.$element[0];
		return t.extend({}, "function" == typeof e.getBoundingClientRect ? e.getBoundingClientRect() : {
			width: e.offsetWidth,
			height: e.offsetHeight
		}, this.$element.offset())
	}, e.prototype.getCalculatedOffset = function(t, e, i, n) {
		return "bottom" == t ? {
			top: e.top + e.height,
			left: e.left + e.width / 2 - i / 2
		} : "top" == t ? {
			top: e.top - n,
			left: e.left + e.width / 2 - i / 2
		} : "left" == t ? {
			top: e.top + e.height / 2 - n / 2,
			left: e.left - i
		} : {
			top: e.top + e.height / 2 - n / 2,
			left: e.left + e.width
		}
	}, e.prototype.getTitle = function() {
		var t, e = this.$element,
			i = this.options;
		return t = e.attr("data-original-title") || ("function" == typeof i.title ? i.title.call(e[0]) : i.title)
	}, e.prototype.tip = function() {
		return this.$tip = this.$tip || t(this.options.template)
	}, e.prototype.arrow = function() {
		return this.$arrow = this.$arrow || this.tip().find(".tooltip-arrow")
	}, e.prototype.validate = function() {
		this.$element[0].parentNode || (this.hide(), this.$element = null, this.options = null)
	}, e.prototype.enable = function() {
		this.enabled = !0
	}, e.prototype.disable = function() {
		this.enabled = !1
	}, e.prototype.toggleEnabled = function() {
		this.enabled = !this.enabled
	}, e.prototype.toggle = function(e) {
		var i = e ? t(e.currentTarget)[this.type](this.getDelegateOptions()).data("zui." + this.type) : this;
		i.tip().hasClass("in") ? i.leave(i) : i.enter(i)
	}, e.prototype.destroy = function() {
		this.hide().$element.off("." + this.type).removeData("zui." + this.type)
	};
	var i = t.fn.tooltip;
	t.fn.tooltip = function(i, n) {
		return this.each(function() {
			var o = t(this),
				s = o.data("zui.tooltip"),
				a = "object" == typeof i && i;
			s || o.data("zui.tooltip", s = new e(this, a)), "string" == typeof i && s[i](n)
		})
	}, t.fn.tooltip.Constructor = e, t.fn.tooltip.noConflict = function() {
		return t.fn.tooltip = i, this
	}
}(window.jQuery), +
function(t) {
	"use strict";
	var e = function(t, e) {
			this.init("popover", t, e)
		};
	if (!t.fn.tooltip) throw new Error("Popover requires tooltip.js");
	e.DEFAULTS = t.extend({}, t.fn.tooltip.Constructor.DEFAULTS, {
		placement: "right",
		trigger: "click",
		content: "",
		template: '<div class="popover"><div class="arrow"></div><h3 class="popover-title"></h3><div class="popover-content"></div></div>'
	}), e.prototype = t.extend({}, t.fn.tooltip.Constructor.prototype), e.prototype.constructor = e, e.prototype.getDefaults = function() {
		return e.DEFAULTS
	}, e.prototype.setContent = function() {
		var t = this.tip(),
			e = this.getTarget();
		if (e) return e.find(".arrow").length < 1 && t.addClass("no-arrow"), void t.html(e.html());
		var i = this.getTitle(),
			n = this.getContent();
		t.find(".popover-title")[this.options.html ? "html" : "text"](i), t.find(".popover-content")[this.options.html ? "html" : "text"](n), t.removeClass("fade top bottom left right in"), this.options.tipId && t.attr("id", this.options.tipId), this.options.tipClass && t.addClass(this.options.tipClass), t.find(".popover-title").html() || t.find(".popover-title").hide()
	}, e.prototype.hasContent = function() {
		return this.getTarget() || this.getTitle() || this.getContent()
	}, e.prototype.getContent = function() {
		var t = this.$element,
			e = this.options;
		return t.attr("data-content") || ("function" == typeof e.content ? e.content.call(t[0]) : e.content)
	}, e.prototype.getTarget = function() {
		var e = this.$element,
			i = this.options,
			n = e.attr("data-target") || ("function" == typeof i.target ? i.target.call(e[0]) : i.target);
		return !!n && ("$next" == n ? e.next(".popover") : t(n))
	}, e.prototype.arrow = function() {
		return this.$arrow = this.$arrow || this.tip().find(".arrow")
	}, e.prototype.tip = function() {
		return this.$tip || (this.$tip = t(this.options.template)), this.$tip
	};
	var i = t.fn.popover;
	t.fn.popover = function(i) {
		return this.each(function() {
			var n = t(this),
				o = n.data("zui.popover"),
				s = "object" == typeof i && i;
			o || n.data("zui.popover", o = new e(this, s)), "string" == typeof i && o[i]()
		})
	}, t.fn.popover.Constructor = e, t.fn.popover.noConflict = function() {
		return t.fn.popover = i, this
	}
}(window.jQuery), +
function(t) {
	"use strict";

	function e() {
		t(o).remove(), t(s).each(function(e) {
			var o = i(t(this));
			o.hasClass("open") && (o.trigger(e = t.Event("hide." + n)), e.isDefaultPrevented() || o.removeClass("open").trigger("hidden." + n))
		})
	}
	function i(e) {
		var i = e.attr("data-target");
		i || (i = e.attr("href"), i = i && /#/.test(i) && i.replace(/.*(?=#[^\s]*$)/, ""));
		var n;
		try {
			n = i && t(i)
		} catch (o) {}
		return n && n.length ? n : e.parent()
	}
	var n = "zui.dropdown",
		o = ".dropdown-backdrop",
		s = "[data-toggle=dropdown]",
		a = function(e) {
			t(e).on("click." + n, this.toggle)
		};
	a.prototype.toggle = function(o) {
		var s = t(this);
		if (!s.is(".disabled, :disabled")) {
			var a = i(s),
				r = a.hasClass("open");
			if (e(), !r) {
				if ("ontouchstart" in document.documentElement && !a.closest(".navbar-nav").length && t('<div class="dropdown-backdrop"/>').insertAfter(t(this)).on("click", e), a.trigger(o = t.Event("show." + n)), o.isDefaultPrevented()) return;
				a.toggleClass("open").trigger("shown." + n), s.focus()
			}
			return !1
		}
	}, a.prototype.keydown = function(e) {
		if (/(38|40|27)/.test(e.keyCode)) {
			var n = t(this);
			if (e.preventDefault(), e.stopPropagation(), !n.is(".disabled, :disabled")) {
				var o = i(n),
					a = o.hasClass("open");
				if (!a || a && 27 == e.keyCode) return 27 == e.which && o.find(s).focus(), n.click();
				var r = t("[role=menu] li:not(.divider):visible a", o);
				if (r.length) {
					var l = r.index(r.filter(":focus"));
					38 == e.keyCode && l > 0 && l--, 40 == e.keyCode && l < r.length - 1 && l++, ~l || (l = 0), r.eq(l).focus()
				}
			}
		}
	};
	var r = t.fn.dropdown;
	t.fn.dropdown = function(e) {
		return this.each(function() {
			var i = t(this),
				n = i.data("dropdown");
			n || i.data("dropdown", n = new a(this)), "string" == typeof e && n[e].call(i)
		})
	}, t.fn.dropdown.Constructor = a, t.fn.dropdown.noConflict = function() {
		return t.fn.dropdown = r, this
	};
	var l = n + ".data-api";
	t(document).on("click." + l, e).on("click." + l, ".dropdown form", function(t) {
		t.stopPropagation()
	}).on("click." + l, s, a.prototype.toggle).on("keydown." + l, s + ", [role=menu]", a.prototype.keydown)
}(window.jQuery), function(t, e, i) {
	"use strict";
	var n = 0,
		o = '<div class="messager messager-{type} {placement}" style="display: none"><div class="messager-content"></div><div class="messager-actions"></div></div>',
		s = {
			type: "default",
			placement: "top",
			time: 4e3,
			parent: "body",
			icon: null,
			close: !0,
			fade: !0,
			scale: !0
		},
		a = {},
		r = function(e, r) {
			t.isPlainObject(e) && (r = e, e = r.message);
			var l = this;
			r = l.options = t.extend({}, s, r), l.id = r.id || n++;
			var h = a[l.id];
			h && h.destroy(), a[l.id] = l, l.message = (r.icon ? '<i class="icon-' + r.icon + ' icon"></i> ' : "") + e, l.$ = t(o.format(r)).toggleClass("fade", r.fade).toggleClass("scale", r.scale).attr("id", "messager-" + l.id), r.cssClass && l.$.addClass(r.cssClass);
			var c = !1,
				d = l.$.find(".messager-actions"),
				u = function(e) {
					var n = t('<button type="button" class="action action-' + e.name + '"/>');
					"close" === e.name && n.addClass("close"), e.html !== i && n.html(e.html), e.icon !== i && n.append('<i class="action-icon icon-' + e.icon + '"/>'), e.text !== i && n.append('<span class="action-text">' + e.text + "</span>"), e.tooltip !== i && n.attr("title", e.tooltip).tooltip(), n.data("action", e), d.append(n)
				};
			r.actions && t.each(r.actions, function(t, e) {
				e.name === i && (e.name = t), "close" == e.name && (c = !0), u(e)
			}), !c && r.close && u({
				name: "close",
				html: "&times;"
			}), l.$.on("click", ".action", function(e) {
				var i, n = t(this).data("action");
				r.onAction && (i = r.onAction.call(this, n.name, n, l), i === !1) || t.isFunction(n.action) && (i = n.action.call(this, l), i === !1) || (l.hide(), e.stopPropagation())
			}), l.$.on("click", function(t) {
				if (r.onAction) {
					var e = r.onAction.call(this, "content", null, l);
					e === !0 && l.hide()
				}
			});
			var p = l.$.find(".messager-content").html(l.message);
			r.contentClass && p.addClass(r.contentClass), l.$.data("zui.messager", l), r.show && l.message !== i && l.show()
		};
	r.prototype.update = function(e, i) {
		var n = this,
			o = n.options;
		n.$.removeClass("messager-" + o.type), i && (o = t.extend(o, i)), n.$.addClass("messager-" + o.type), e && (n.message = (o.icon ? '<i class="icon-' + o.icon + ' icon"></i> ' : "") + e, n.$.find(".messager-content").html(n.message))
	}, r.prototype.show = function(n, o) {
		var s = this,
			a = this.options;
		if (t.isFunction(n)) {
			var r = o;
			o = n, r !== i && (n = r)
		}
		if (s.isShow) return void s.hide(function() {
			s.show(n, o)
		});
		s.hiding && (clearTimeout(s.hiding), s.hiding = null), s.update(n);
		var l = a.placement,
			h = t(a.parent),
			c = h.children(".messagers-holder." + l);
		if (c.length || (c = t("<div/>").attr("class", "messagers-holder " + l).appendTo(h)), c.append(s.$), "center" === l) {
			var d = t(e).height() - c.height();
			c.css("top", Math.max(-d, d / 2))
		}
		return s.$.show().addClass("in"), a.time && (s.hiding = setTimeout(function() {
			s.hide()
		}, a.time)), s.isShow = !0, o && o(), s
	}, r.prototype.hide = function(t, e) {
		t === !0 && (e = !0, t = null);
		var i = this;
		if (i.$.hasClass("in")) {
			i.$.removeClass("in");
			var n = function() {
					var e = i.$.parent();
					i.$.detach(), e.children().length || e.remove(), t && t(!0)
				};
			e ? n() : setTimeout(n, 200)
		} else t && t(!1);
		i.isShow = !1
	}, r.prototype.destroy = function() {
		var t = this;
		t.hide(function() {
			t.$.remove(), t.$ = null
		}, !0), delete a[t.id]
	}, r.all = a, r.DEFAULTS = s;
	var l = function() {
			t(".messager").each(function() {
				var e = t(this).data("zui.messager");
				e && e.hide && e.hide(!0)
			})
		},
		h = function(e, n) {
			"string" == typeof n && (n = {
				type: n
			}), n = t.extend({}, n), n.id === i && l();
			var o = a[n.id] || new r(e, n);
			return o.show(), o
		},
		c = function(t) {
			return "string" == typeof t ? {
				placement: t
			} : t
		},
		d = {
			show: h,
			hide: l
		};
	t.each({
		primary: 0,
		success: "ok-sign",
		info: "info-sign",
		warning: "warning-sign",
		danger: "exclamation-sign",
		important: 0,
		special: 0
	}, function(e, i) {
		d[e] = function(n, o) {
			return h(n, t.extend({
				type: e,
				icon: i || null
			}, c(o)))
		}
	}), t.zui({
		Messager: r,
		showMessager: h,
		messager: d
	})
}(jQuery, window, void 0), function(t, e, i, n) {
	"use strict";

	function o(t) {
		if (t = t.toLowerCase(), t && c.test(t)) {
			var e;
			if (4 === t.length) {
				var i = "#";
				for (e = 1; e < 4; e += 1) i += t.slice(e, e + 1).concat(t.slice(e, e + 1));
				t = i
			}
			var n = [];
			for (e = 1; e < 7; e += 2) n.push(b("0x" + t.slice(e, e + 2)));
			return {
				r: n[0],
				g: n[1],
				b: n[2],
				a: 1
			}
		}
		throw new Error("Wrong hex string! (hex: " + t + ")")
	}
	function s(e) {
		return typeof e === f && ("transparent" === e.toLowerCase() || m[e.toLowerCase()] || c.test(t.trim(e.toLowerCase())))
	}
	function a(t) {
		function e(t) {
			return t = t < 0 ? t + 1 : t > 1 ? t - 1 : t, 6 * t < 1 ? r + (a - r) * t * 6 : 2 * t < 1 ? a : 3 * t < 2 ? r + (a - r) * (2 / 3 - t) * 6 : r
		}
		var i = t.h,
			n = t.s,
			o = t.l,
			s = t.a;
		i = h(i) % u / u, n = l(h(n)), o = l(h(o)), s = l(h(s));
		var a = o <= .5 ? o * (n + 1) : o + n - o * n,
			r = 2 * o - a,
			c = {
				r: e(i + 1 / 3) * d,
				g: e(i) * d,
				b: e(i - 1 / 3) * d,
				a: s
			};
		return c
	}
	function r(t, i, n) {
		return v(n) && (n = 0), v(i) && (i = d), e.min(e.max(t, n), i)
	}
	function l(t, e) {
		return r(t, e)
	}
	function h(t) {
		return "number" == typeof t ? t : parseFloat(t)
	}
	var c = /^#([0-9a-fA-f]{3}|[0-9a-fA-f]{6})$/,
		d = 255,
		u = 360,
		p = 100,
		f = "string",
		g = "object",
		m = {
			aliceblue: "#f0f8ff",
			antiquewhite: "#faebd7",
			aqua: "#00ffff",
			aquamarine: "#7fffd4",
			azure: "#f0ffff",
			beige: "#f5f5dc",
			bisque: "#ffe4c4",
			black: "#000000",
			blanchedalmond: "#ffebcd",
			blue: "#0000ff",
			blueviolet: "#8a2be2",
			brown: "#a52a2a",
			burlywood: "#deb887",
			cadetblue: "#5f9ea0",
			chartreuse: "#7fff00",
			chocolate: "#d2691e",
			coral: "#ff7f50",
			cornflowerblue: "#6495ed",
			cornsilk: "#fff8dc",
			crimson: "#dc143c",
			cyan: "#00ffff",
			darkblue: "#00008b",
			darkcyan: "#008b8b",
			darkgoldenrod: "#b8860b",
			darkgray: "#a9a9a9",
			darkgreen: "#006400",
			darkkhaki: "#bdb76b",
			darkmagenta: "#8b008b",
			darkolivegreen: "#556b2f",
			darkorange: "#ff8c00",
			darkorchid: "#9932cc",
			darkred: "#8b0000",
			darksalmon: "#e9967a",
			darkseagreen: "#8fbc8f",
			darkslateblue: "#483d8b",
			darkslategray: "#2f4f4f",
			darkturquoise: "#00ced1",
			darkviolet: "#9400d3",
			deeppink: "#ff1493",
			deepskyblue: "#00bfff",
			dimgray: "#696969",
			dodgerblue: "#1e90ff",
			firebrick: "#b22222",
			floralwhite: "#fffaf0",
			forestgreen: "#228b22",
			fuchsia: "#ff00ff",
			gainsboro: "#dcdcdc",
			ghostwhite: "#f8f8ff",
			gold: "#ffd700",
			goldenrod: "#daa520",
			gray: "#808080",
			green: "#008000",
			greenyellow: "#adff2f",
			honeydew: "#f0fff0",
			hotpink: "#ff69b4",
			indianred: "#cd5c5c",
			indigo: "#4b0082",
			ivory: "#fffff0",
			khaki: "#f0e68c",
			lavender: "#e6e6fa",
			lavenderblush: "#fff0f5",
			lawngreen: "#7cfc00",
			lemonchiffon: "#fffacd",
			lightblue: "#add8e6",
			lightcoral: "#f08080",
			lightcyan: "#e0ffff",
			lightgoldenrodyellow: "#fafad2",
			lightgray: "#d3d3d3",
			lightgreen: "#90ee90",
			lightpink: "#ffb6c1",
			lightsalmon: "#ffa07a",
			lightseagreen: "#20b2aa",
			lightskyblue: "#87cefa",
			lightslategray: "#778899",
			lightsteelblue: "#b0c4de",
			lightyellow: "#ffffe0",
			lime: "#00ff00",
			limegreen: "#32cd32",
			linen: "#faf0e6",
			magenta: "#ff00ff",
			maroon: "#800000",
			mediumaquamarine: "#66cdaa",
			mediumblue: "#0000cd",
			mediumorchid: "#ba55d3",
			mediumpurple: "#9370db",
			mediumseagreen: "#3cb371",
			mediumslateblue: "#7b68ee",
			mediumspringgreen: "#00fa9a",
			mediumturquoise: "#48d1cc",
			mediumvioletred: "#c71585",
			midnightblue: "#191970",
			mintcream: "#f5fffa",
			mistyrose: "#ffe4e1",
			moccasin: "#ffe4b5",
			navajowhite: "#ffdead",
			navy: "#000080",
			oldlace: "#fdf5e6",
			olive: "#808000",
			olivedrab: "#6b8e23",
			orange: "#ffa500",
			orangered: "#ff4500",
			orchid: "#da70d6",
			palegoldenrod: "#eee8aa",
			palegreen: "#98fb98",
			paleturquoise: "#afeeee",
			palevioletred: "#db7093",
			papayawhip: "#ffefd5",
			peachpuff: "#ffdab9",
			peru: "#cd853f",
			pink: "#ffc0cb",
			plum: "#dda0dd",
			powderblue: "#b0e0e6",
			purple: "#800080",
			red: "#ff0000",
			rosybrown: "#bc8f8f",
			royalblue: "#4169e1",
			saddlebrown: "#8b4513",
			salmon: "#fa8072",
			sandybrown: "#f4a460",
			seagreen: "#2e8b57",
			seashell: "#fff5ee",
			sienna: "#a0522d",
			silver: "#c0c0c0",
			skyblue: "#87ceeb",
			slateblue: "#6a5acd",
			slategray: "#708090",
			snow: "#fffafa",
			springgreen: "#00ff7f",
			steelblue: "#4682b4",
			tan: "#d2b48c",
			teal: "#008080",
			thistle: "#d8bfd8",
			tomato: "#ff6347",
			turquoise: "#40e0d0",
			violet: "#ee82ee",
			wheat: "#f5deb3",
			white: "#ffffff",
			whitesmoke: "#f5f5f5",
			yellow: "#ffff00",
			yellowgreen: "#9acd32"
		},
		v = function(t) {
			return t === n
		},
		y = function(t) {
			return !v(t)
		},
		b = function(t) {
			return parseInt(t)
		},
		w = function(t) {
			return b(l(h(t), d))
		},
		x = function(t, e, i, n) {
			var s = this;
			if (s.r = s.g = s.b = 0, s.a = 1, y(n) && (s.a = l(h(n), 1)), y(t) && y(e) && y(i)) s.r = w(t), s.g = w(e), s.b = w(i);
			else if (y(t)) {
				var r = typeof t;
				if (r == f) if (t = t.toLowerCase(), "transparent" === t) s.a = 0;
				else if (m[t]) this.rgb(o(m[t]));
				else if (0 === t.indexOf("rgb")) {
					var c = t.substring(t.indexOf("(") + 1, t.lastIndexOf(")")).split(",", 4);
					s.rgb({
						r: c[0],
						g: c[1],
						b: c[2],
						a: c[3]
					})
				} else s.rgb(o(t));
				else if ("number" == r && v(e)) s.r = s.g = s.b = w(t);
				else if (r == g && y(t.r)) s.r = w(t.r), y(t.g) && (s.g = w(t.g)), y(t.b) && (s.b = w(t.b)), y(t.a) && (s.a = l(h(t.a), 1));
				else if (r == g && y(t.h)) {
					var d = {
						h: l(h(t.h), u),
						s: 1,
						l: 1,
						a: 1
					};
					y(t.s) && (d.s = l(h(t.s), 1)), y(t.l) && (d.l = l(h(t.l), 1)), y(t.a) && (d.a = l(h(t.a), 1)), s.rgb(a(d))
				}
			}
		};
	x.prototype.rgb = function(t) {
		var e = this;
		if (y(t)) {
			if (typeof t == g) y(t.r) && (e.r = w(t.r)), y(t.g) && (e.g = w(t.g)), y(t.b) && (e.b = w(t.b)), y(t.a) && (e.a = l(h(t.a), 1));
			else {
				var i = b(h(t));
				e.r = i, e.g = i, e.b = i
			}
			return e
		}
		return {
			r: e.r,
			g: e.g,
			b: e.b,
			a: e.a
		}
	}, x.prototype.hue = function(t) {
		var e = this,
			i = e.toHsl();
		return v(t) ? i.h : (i.h = l(h(t), u), e.rgb(a(i)), e)
	}, x.prototype.darken = function(t) {
		var e = this,
			i = e.toHsl();
		return i.l -= t / p, i.l = l(i.l, 1), e.rgb(a(i)), e
	}, x.prototype.clone = function() {
		var t = this;
		return new x(t.r, t.g, t.b, t.a)
	}, x.prototype.lighten = function(t) {
		return this.darken(-t)
	}, x.prototype.fade = function(t) {
		return this.a = l(t / p, 1), this
	}, x.prototype.spin = function(t) {
		var e = this.toHsl(),
			i = (e.h + t) % u;
		return e.h = i < 0 ? u + i : i, this.rgb(a(e))
	}, x.prototype.toHsl = function() {
		var t, i, n = this,
			o = n.r / d,
			s = n.g / d,
			a = n.b / d,
			r = n.a,
			l = e.max(o, s, a),
			h = e.min(o, s, a),
			c = (l + h) / 2,
			p = l - h;
		if (l === h) t = i = 0;
		else {
			switch (i = c > .5 ? p / (2 - l - h) : p / (l + h), l) {
			case o:
				t = (s - a) / p + (s < a ? 6 : 0);
				break;
			case s:
				t = (a - o) / p + 2;
				break;
			case a:
				t = (o - s) / p + 4
			}
			t /= 6
		}
		return {
			h: t * u,
			s: i,
			l: c,
			a: r
		}
	}, x.prototype.luma = function() {
		var t = this.r / d,
			i = this.g / d,
			n = this.b / d;
		return t = t <= .03928 ? t / 12.92 : e.pow((t + .055) / 1.055, 2.4), i = i <= .03928 ? i / 12.92 : e.pow((i + .055) / 1.055, 2.4), n = n <= .03928 ? n / 12.92 : e.pow((n + .055) / 1.055, 2.4), .2126 * t + .7152 * i + .0722 * n
	}, x.prototype.saturate = function(t) {
		var e = this.toHsl();
		return e.s += t / p, e.s = l(e.s), this.rgb(a(e))
	}, x.prototype.desaturate = function(t) {
		return this.saturate(-t)
	}, x.prototype.contrast = function(t, e, i) {
		if (e = v(e) ? new x(d, d, d, 1) : new x(e), t = v(t) ? new x(0, 0, 0, 1) : new x(t), t.luma() > e.luma()) {
			var n = e;
			e = t, t = n
		}
		return this.a < .5 ? t : (i = v(i) ? .43 : h(i), this.luma() < i ? e : t)
	}, x.prototype.hexStr = function() {
		var t = this.r.toString(16),
			e = this.g.toString(16),
			i = this.b.toString(16);
		return 1 == t.length && (t = "0" + t), 1 == e.length && (e = "0" + e), 1 == i.length && (i = "0" + i), "#" + t + e + i
	}, x.prototype.toCssStr = function() {
		var t = this;
		return t.a > 0 ? t.a < 1 ? "rgba(" + t.r + "," + t.g + "," + t.b + "," + t.a + ")" : t.hexStr() : "transparent"
	}, x.isColor = s, x.names = m, x.get = function(t) {
		return new x(t)
	}, t.zui({
		Color: x
	})
}(jQuery, Math, window, void 0),
/*!
 * Chart.js 1.0.2
 * Copyright 2015 Nick Downie
 * Released under the MIT license
 * http://chartjs.org/
 */

function(t) {
	"use strict";
	var e = t && t.zui ? t.zui : this,
		i = (e.Chart, function(t) {
			this.canvas = t.canvas, this.ctx = t;
			var e = function(t, e) {
					return t["offset" + e] ? t["offset" + e] : document.defaultView.getComputedStyle(t).getPropertyValue(e)
				},
				i = this.width = e(t.canvas, "Width"),
				o = this.height = e(t.canvas, "Height");
			t.canvas.width = i, t.canvas.height = o;
			var i = this.width = t.canvas.width,
				o = this.height = t.canvas.height;
			return this.aspectRatio = this.width / this.height, n.retinaScale(this), this
		});
	i.defaults = {
		global: {
			animation: !0,
			animationSteps: 60,
			animationEasing: "easeOutQuart",
			showScale: !0,
			scaleOverride: !1,
			scaleSteps: null,
			scaleStepWidth: null,
			scaleStartValue: null,
			scaleLineColor: "rgba(0,0,0,.1)",
			scaleLineWidth: 1,
			scaleShowLabels: !0,
			scaleLabel: "<%=value%>",
			scaleIntegersOnly: !0,
			scaleBeginAtZero: !1,
			scaleFontFamily: "'Helvetica Neue', 'Helvetica', 'Arial', sans-serif",
			scaleFontSize: 12,
			scaleFontStyle: "normal",
			scaleFontColor: "#666",
			responsive: !1,
			maintainAspectRatio: !0,
			showTooltips: !0,
			customTooltips: !1,
			tooltipEvents: ["mousemove", "touchstart", "touchmove", "mouseout"],
			tooltipFillColor: "rgba(0,0,0,0.8)",
			tooltipFontFamily: "'Helvetica Neue', 'Helvetica', 'Arial', sans-serif",
			tooltipFontSize: 14,
			tooltipFontStyle: "normal",
			tooltipFontColor: "#fff",
			tooltipTitleFontFamily: "'Helvetica Neue', 'Helvetica', 'Arial', sans-serif",
			tooltipTitleFontSize: 14,
			tooltipTitleFontStyle: "bold",
			tooltipTitleFontColor: "#fff",
			tooltipYPadding: 6,
			tooltipXPadding: 6,
			tooltipCaretSize: 8,
			tooltipCornerRadius: 6,
			tooltipXOffset: 10,
			tooltipTemplate: "<%if (label){%><%=label%>: <%}%><%= value %>",
			multiTooltipTemplate: "<%if (datasetLabel){%><%=datasetLabel%>: <%}%><%= value %>",
			multiTooltipTitleTemplate: "<%= label %>",
			multiTooltipKeyBackground: "#fff",
			onAnimationProgress: function() {},
			onAnimationComplete: function() {}
		}
	}, i.types = {};
	var n = i.helpers = {},
		o = n.each = function(t, e, i) {
			var n = Array.prototype.slice.call(arguments, 3);
			if (t) if (t.length === +t.length) {
				var o;
				for (o = 0; o < t.length; o++) e.apply(i, [t[o], o].concat(n))
			} else for (var s in t) e.apply(i, [t[s], s].concat(n))
		},
		s = n.clone = function(t) {
			var e = {};
			return o(t, function(i, n) {
				t.hasOwnProperty(n) && (e[n] = i)
			}), e
		},
		a = n.extend = function(t) {
			return o(Array.prototype.slice.call(arguments, 1), function(e) {
				o(e, function(i, n) {
					e.hasOwnProperty(n) && (t[n] = i)
				})
			}), t
		},
		r = n.merge = function(t, e) {
			var i = Array.prototype.slice.call(arguments, 0);
			return i.unshift({}), a.apply(null, i)
		},
		l = n.indexOf = function(t, e) {
			if (Array.prototype.indexOf) return t.indexOf(e);
			for (var i = 0; i < t.length; i++) if (t[i] === e) return i;
			return -1
		},
		h = (n.where = function(t, e) {
			var i = [];
			return n.each(t, function(t) {
				e(t) && i.push(t)
			}), i
		}, n.findNextWhere = function(t, e, i) {
			i || (i = -1);
			for (var n = i + 1; n < t.length; n++) {
				var o = t[n];
				if (e(o)) return o
			}
		}, n.findPreviousWhere = function(t, e, i) {
			i || (i = t.length);
			for (var n = i - 1; n >= 0; n--) {
				var o = t[n];
				if (e(o)) return o
			}
		}, n.inherits = function(t) {
			var e = this,
				i = t && t.hasOwnProperty("constructor") ? t.constructor : function() {
					return e.apply(this, arguments)
				},
				n = function() {
					this.constructor = i
				};
			return n.prototype = e.prototype, i.prototype = new n, i.extend = h, t && a(i.prototype, t), i.__super__ = e.prototype, i
		}),
		c = n.noop = function() {},
		d = n.uid = function() {
			var t = 0;
			return function() {
				return "chart-" + t++
			}
		}(),
		u = n.warn = function(t) {
			window.console && "function" == typeof window.console.warn && console.warn(t)
		},
		p = n.amd = "function" == typeof define && define.amd,
		f = n.isNumber = function(t) {
			return !isNaN(parseFloat(t)) && isFinite(t)
		},
		g = n.max = function(t) {
			return Math.max.apply(Math, t)
		},
		m = n.min = function(t) {
			return Math.min.apply(Math, t)
		},
		v = (n.cap = function(t, e, i) {
			if (f(e)) {
				if (t > e) return e
			} else if (f(i) && t < i) return i;
			return t
		}, n.getDecimalPlaces = function(t) {
			return t % 1 !== 0 && f(t) ? t.toString().split(".")[1].length : 0
		}),
		y = n.radians = function(t) {
			return t * (Math.PI / 180)
		},
		b = (n.getAngleFromPoint = function(t, e) {
			var i = e.x - t.x,
				n = e.y - t.y,
				o = Math.sqrt(i * i + n * n),
				s = 2 * Math.PI + Math.atan2(n, i);
			return i < 0 && n < 0 && (s += 2 * Math.PI), {
				angle: s,
				distance: o
			}
		}, n.aliasPixel = function(t) {
			return t % 2 === 0 ? 0 : .5
		}),
		w = (n.splineCurve = function(t, e, i, n) {
			var o = Math.sqrt(Math.pow(e.x - t.x, 2) + Math.pow(e.y - t.y, 2)),
				s = Math.sqrt(Math.pow(i.x - e.x, 2) + Math.pow(i.y - e.y, 2)),
				a = n * o / (o + s),
				r = n * s / (o + s);
			return {
				inner: {
					x: e.x - a * (i.x - t.x),
					y: e.y - a * (i.y - t.y)
				},
				outer: {
					x: e.x + r * (i.x - t.x),
					y: e.y + r * (i.y - t.y)
				}
			}
		}, n.calculateOrderOfMagnitude = function(t) {
			return Math.floor(Math.log(t) / Math.LN10)
		}),
		x = (n.calculateScaleRange = function(t, e, i, n, o) {
			var s = 2,
				a = Math.floor(e / (1.5 * i)),
				r = s >= a,
				l = g(t),
				h = m(t);
			l === h && (l += .5, h >= .5 && !n ? h -= .5 : l += .5);
			for (var c = Math.abs(l - h), d = w(c), u = Math.ceil(l / (1 * Math.pow(10, d))) * Math.pow(10, d), p = n ? 0 : Math.floor(h / (1 * Math.pow(10, d))) * Math.pow(10, d), f = u - p, v = Math.pow(10, d), y = Math.round(f / v);
			(y > a || 2 * y < a) && !r;) if (y > a) v *= 2, y = Math.round(f / v), y % 1 !== 0 && (r = !0);
			else if (o && d >= 0) {
				if (v / 2 % 1 !== 0) break;
				v /= 2, y = Math.round(f / v)
			} else v /= 2, y = Math.round(f / v);
			return r && (y = s, v = f / y), {
				steps: y,
				stepValue: v,
				min: p,
				max: p + y * v
			}
		}, n.template = function(t, e) {
			function i(t, e) {
				var i = /\W/.test(t) ? new Function("obj", "var p=[],print=function(){p.push.apply(p,arguments);};with(obj){p.push('" + t.replace(/[\r\t\n]/g, " ").split("<%").join("\t").replace(/((^|%>)[^\t]*)'/g, "$1\r").replace(/\t=(.*?)%>/g, "',$1,'").split("\t").join("');").split("%>").join("p.push('").split("\r").join("\\'") + "');}return p.join('');") : n[t] = n[t];
				return e ? i(e) : i
			}
			if (t instanceof Function) return t(e);
			var n = {};
			return i(t, e)
		}),
		C = (n.generateLabels = function(t, e, i, n) {
			var s = new Array(e);
			return labelTemplateString && o(s, function(e, o) {
				s[o] = x(t, {
					value: i + n * (o + 1)
				})
			}), s
		}, n.easingEffects = {
			linear: function(t) {
				return t
			},
			easeInQuad: function(t) {
				return t * t
			},
			easeOutQuad: function(t) {
				return -1 * t * (t - 2)
			},
			easeInOutQuad: function(t) {
				return (t /= .5) < 1 ? .5 * t * t : -.5 * (--t * (t - 2) - 1)
			},
			easeInCubic: function(t) {
				return t * t * t
			},
			easeOutCubic: function(t) {
				return 1 * ((t = t / 1 - 1) * t * t + 1)
			},
			easeInOutCubic: function(t) {
				return (t /= .5) < 1 ? .5 * t * t * t : .5 * ((t -= 2) * t * t + 2)
			},
			easeInQuart: function(t) {
				return t * t * t * t
			},
			easeOutQuart: function(t) {
				return -1 * ((t = t / 1 - 1) * t * t * t - 1)
			},
			easeInOutQuart: function(t) {
				return (t /= .5) < 1 ? .5 * t * t * t * t : -.5 * ((t -= 2) * t * t * t - 2)
			},
			easeInQuint: function(t) {
				return 1 * (t /= 1) * t * t * t * t
			},
			easeOutQuint: function(t) {
				return 1 * ((t = t / 1 - 1) * t * t * t * t + 1)
			},
			easeInOutQuint: function(t) {
				return (t /= .5) < 1 ? .5 * t * t * t * t * t : .5 * ((t -= 2) * t * t * t * t + 2)
			},
			easeInSine: function(t) {
				return -1 * Math.cos(t / 1 * (Math.PI / 2)) + 1
			},
			easeOutSine: function(t) {
				return 1 * Math.sin(t / 1 * (Math.PI / 2))
			},
			easeInOutSine: function(t) {
				return -.5 * (Math.cos(Math.PI * t / 1) - 1)
			},
			easeInExpo: function(t) {
				return 0 === t ? 1 : 1 * Math.pow(2, 10 * (t / 1 - 1))
			},
			easeOutExpo: function(t) {
				return 1 === t ? 1 : 1 * (-Math.pow(2, -10 * t / 1) + 1)
			},
			easeInOutExpo: function(t) {
				return 0 === t ? 0 : 1 === t ? 1 : (t /= .5) < 1 ? .5 * Math.pow(2, 10 * (t - 1)) : .5 * (-Math.pow(2, -10 * --t) + 2)
			},
			easeInCirc: function(t) {
				return t >= 1 ? t : -1 * (Math.sqrt(1 - (t /= 1) * t) - 1)
			},
			easeOutCirc: function(t) {
				return 1 * Math.sqrt(1 - (t = t / 1 - 1) * t)
			},
			easeInOutCirc: function(t) {
				return (t /= .5) < 1 ? -.5 * (Math.sqrt(1 - t * t) - 1) : .5 * (Math.sqrt(1 - (t -= 2) * t) + 1)
			},
			easeInElastic: function(t) {
				var e = 1.70158,
					i = 0,
					n = 1;
				return 0 === t ? 0 : 1 == (t /= 1) ? 1 : (i || (i = .3), n < Math.abs(1) ? (n = 1, e = i / 4) : e = i / (2 * Math.PI) * Math.asin(1 / n), -(n * Math.pow(2, 10 * (t -= 1)) * Math.sin((1 * t - e) * (2 * Math.PI) / i)))
			},
			easeOutElastic: function(t) {
				var e = 1.70158,
					i = 0,
					n = 1;
				return 0 === t ? 0 : 1 == (t /= 1) ? 1 : (i || (i = .3), n < Math.abs(1) ? (n = 1, e = i / 4) : e = i / (2 * Math.PI) * Math.asin(1 / n), n * Math.pow(2, -10 * t) * Math.sin((1 * t - e) * (2 * Math.PI) / i) + 1)
			},
			easeInOutElastic: function(t) {
				var e = 1.70158,
					i = 0,
					n = 1;
				return 0 === t ? 0 : 2 == (t /= .5) ? 1 : (i || (i = 1 * (.3 * 1.5)), n < Math.abs(1) ? (n = 1, e = i / 4) : e = i / (2 * Math.PI) * Math.asin(1 / n), t < 1 ? -.5 * (n * Math.pow(2, 10 * (t -= 1)) * Math.sin((1 * t - e) * (2 * Math.PI) / i)) : n * Math.pow(2, -10 * (t -= 1)) * Math.sin((1 * t - e) * (2 * Math.PI) / i) * .5 + 1)
			},
			easeInBack: function(t) {
				var e = 1.70158;
				return 1 * (t /= 1) * t * ((e + 1) * t - e)
			},
			easeOutBack: function(t) {
				var e = 1.70158;
				return 1 * ((t = t / 1 - 1) * t * ((e + 1) * t + e) + 1)
			},
			easeInOutBack: function(t) {
				var e = 1.70158;
				return (t /= .5) < 1 ? .5 * (t * t * (((e *= 1.525) + 1) * t - e)) : .5 * ((t -= 2) * t * (((e *= 1.525) + 1) * t + e) + 2)
			},
			easeInBounce: function(t) {
				return 1 - C.easeOutBounce(1 - t)
			},
			easeOutBounce: function(t) {
				return (t /= 1) < 1 / 2.75 ? 1 * (7.5625 * t * t) : t < 2 / 2.75 ? 1 * (7.5625 * (t -= 1.5 / 2.75) * t + .75) : t < 2.5 / 2.75 ? 1 * (7.5625 * (t -= 2.25 / 2.75) * t + .9375) : 1 * (7.5625 * (t -= 2.625 / 2.75) * t + .984375)
			},
			easeInOutBounce: function(t) {
				return t < .5 ? .5 * C.easeInBounce(2 * t) : .5 * C.easeOutBounce(2 * t - 1) + .5
			}
		}),
		_ = n.requestAnimFrame = function() {
			return window.requestAnimationFrame || window.webkitRequestAnimationFrame || window.mozRequestAnimationFrame || window.oRequestAnimationFrame || window.msRequestAnimationFrame ||
			function(t) {
				return window.setTimeout(t, 1e3 / 60)
			}
		}(),
		k = n.cancelAnimFrame = function() {
			return window.cancelAnimationFrame || window.webkitCancelAnimationFrame || window.mozCancelAnimationFrame || window.oCancelAnimationFrame || window.msCancelAnimationFrame ||
			function(t) {
				return window.clearTimeout(t, 1e3 / 60)
			}
		}(),
		T = (n.animationLoop = function(t, e, i, n, o, s) {
			var a = 0,
				r = C[i] || C.linear,
				l = function() {
					a++;
					var i = a / e,
						h = r(i);
					t.call(s, h, i, a), n.call(s, h, i), a < e ? s.animationFrame = _(l) : o.apply(s)
				};
			_(l)
		}, n.getRelativePosition = function(t) {
			var e, i, n = t.originalEvent || t,
				o = t.currentTarget || t.srcElement,
				s = o.getBoundingClientRect();
			return n.touches ? (e = n.touches[0].clientX - s.left, i = n.touches[0].clientY - s.top) : (e = n.clientX - s.left, i = n.clientY - s.top), {
				x: e,
				y: i
			}
		}, n.addEvent = function(t, e, i) {
			t.addEventListener ? t.addEventListener(e, i) : t.attachEvent ? t.attachEvent("on" + e, i) : t["on" + e] = i
		}),
		S = n.removeEvent = function(t, e, i) {
			t.removeEventListener ? t.removeEventListener(e, i, !1) : t.detachEvent ? t.detachEvent("on" + e, i) : t["on" + e] = c
		},
		D = (n.bindEvents = function(t, e, i) {
			t.events || (t.events = {}), o(e, function(e) {
				t.events[e] = function() {
					i.apply(t, arguments)
				}, T(t.chart.canvas, e, t.events[e])
			})
		}, n.unbindEvents = function(t, e) {
			o(e, function(e, i) {
				S(t.chart.canvas, i, e)
			})
		}),
		M = n.getMaximumWidth = function(t) {
			var e = t.parentNode;
			return e.clientWidth
		},
		P = n.getMaximumHeight = function(t) {
			var e = t.parentNode;
			return e.clientHeight
		},
		F = (n.getMaximumSize = n.getMaximumWidth, n.retinaScale = function(t) {
			var e = t.ctx,
				i = t.canvas.width,
				n = t.canvas.height;
			window.devicePixelRatio && (e.canvas.style.width = i + "px", e.canvas.style.height = n + "px", e.canvas.height = n * window.devicePixelRatio, e.canvas.width = i * window.devicePixelRatio, e.scale(window.devicePixelRatio, window.devicePixelRatio))
		}),
		L = n.clear = function(t) {
			t.ctx.clearRect(0, 0, t.width, t.height)
		},
		z = n.fontString = function(t, e, i) {
			return e + " " + t + "px " + i
		},
		I = n.longestText = function(t, e, i) {
			t.font = e;
			var n = 0;
			return o(i, function(e) {
				var i = t.measureText(e).width;
				n = i > n ? i : n
			}), n
		},
		$ = n.drawRoundedRectangle = function(t, e, i, n, o, s) {
			t.beginPath(), t.moveTo(e + s, i), t.lineTo(e + n - s, i), t.quadraticCurveTo(e + n, i, e + n, i + s), t.lineTo(e + n, i + o - s), t.quadraticCurveTo(e + n, i + o, e + n - s, i + o), t.lineTo(e + s, i + o), t.quadraticCurveTo(e, i + o, e, i + o - s), t.lineTo(e, i + s), t.quadraticCurveTo(e, i, e + s, i), t.closePath()
		};
	i.instances = {}, i.Type = function(t, e, n) {
		this.options = e, this.chart = n, this.id = d(), i.instances[this.id] = this, e.responsive && this.resize(), this.initialize.call(this, t)
	}, a(i.Type.prototype, {
		initialize: function() {
			return this
		},
		clear: function() {
			return L(this.chart), this
		},
		stop: function() {
			return k(this.animationFrame), this
		},
		resize: function(t) {
			this.stop();
			var e = this.chart.canvas,
				i = M(this.chart.canvas),
				n = this.options.maintainAspectRatio ? i / this.chart.aspectRatio : P(this.chart.canvas);
			return e.width = this.chart.width = i, e.height = this.chart.height = n, F(this.chart), "function" == typeof t && t.apply(this, Array.prototype.slice.call(arguments, 1)), this
		},
		reflow: c,
		render: function(t) {
			return t && this.reflow(), this.options.animation && !t ? n.animationLoop(this.draw, this.options.animationSteps, this.options.animationEasing, this.options.onAnimationProgress, this.options.onAnimationComplete, this) : (this.draw(), this.options.onAnimationComplete.call(this)), this
		},
		generateLegend: function() {
			return x(this.options.legendTemplate, this)
		},
		destroy: function() {
			this.clear(), D(this, this.events);
			var t = this.chart.canvas;
			t.width = this.chart.width, t.height = this.chart.height, t.style.removeProperty ? (t.style.removeProperty("width"), t.style.removeProperty("height")) : (t.style.removeAttribute("width"), t.style.removeAttribute("height")), delete i.instances[this.id]
		},
		showTooltip: function(t, e) {
			"undefined" == typeof this.activeElements && (this.activeElements = []);
			var s = function(t) {
					var e = !1;
					return t.length !== this.activeElements.length ? e = !0 : (o(t, function(t, i) {
						t !== this.activeElements[i] && (e = !0)
					}, this), e)
				}.call(this, t);
			if (s || e) {
				if (this.activeElements = t, this.draw(), this.options.customTooltips && this.options.customTooltips(!1), t.length > 0) if (this.datasets && this.datasets.length > 1) {
					for (var a, r, h = this.datasets.length - 1; h >= 0 && (a = this.datasets[h].points || this.datasets[h].bars || this.datasets[h].segments, r = l(a, t[0]), r === -1); h--);
					var c = [],
						d = [],
						u = function(t) {
							var e, i, o, s, a, l = [],
								h = [],
								u = [];
							return n.each(this.datasets, function(t) {
								t.showTooltips !== !1 && (e = t.points || t.bars || t.segments, e[r] && e[r].hasValue() && l.push(e[r]))
							}), n.each(l, function(t) {
								h.push(t.x), u.push(t.y), c.push(n.template(this.options.multiTooltipTemplate, t)), d.push({
									fill: t._saved.fillColor || t.fillColor,
									stroke: t._saved.strokeColor || t.strokeColor
								})
							}, this), a = m(u), o = g(u), s = m(h), i = g(h), {
								x: s > this.chart.width / 2 ? s : i,
								y: (a + o) / 2
							}
						}.call(this, r);
					new i.MultiTooltip({
						x: u.x,
						y: u.y,
						xPadding: this.options.tooltipXPadding,
						yPadding: this.options.tooltipYPadding,
						xOffset: this.options.tooltipXOffset,
						fillColor: this.options.tooltipFillColor,
						textColor: this.options.tooltipFontColor,
						fontFamily: this.options.tooltipFontFamily,
						fontStyle: this.options.tooltipFontStyle,
						fontSize: this.options.tooltipFontSize,
						titleTextColor: this.options.tooltipTitleFontColor,
						titleFontFamily: this.options.tooltipTitleFontFamily,
						titleFontStyle: this.options.tooltipTitleFontStyle,
						titleFontSize: this.options.tooltipTitleFontSize,
						cornerRadius: this.options.tooltipCornerRadius,
						labels: c,
						legendColors: d,
						legendColorBackground: this.options.multiTooltipKeyBackground,
						title: x(this.options.multiTooltipTitleTemplate, t[0]),
						chart: this.chart,
						ctx: this.chart.ctx,
						custom: this.options.customTooltips
					}).draw()
				} else o(t, function(t) {
					var e = t.tooltipPosition();
					new i.Tooltip({
						x: Math.round(e.x),
						y: Math.round(e.y),
						xPadding: this.options.tooltipXPadding,
						yPadding: this.options.tooltipYPadding,
						fillColor: this.options.tooltipFillColor,
						textColor: this.options.tooltipFontColor,
						fontFamily: this.options.tooltipFontFamily,
						fontStyle: this.options.tooltipFontStyle,
						fontSize: this.options.tooltipFontSize,
						caretHeight: this.options.tooltipCaretSize,
						cornerRadius: this.options.tooltipCornerRadius,
						text: x(this.options.tooltipTemplate, t),
						chart: this.chart,
						custom: this.options.customTooltips
					}).draw()
				}, this);
				return this
			}
		},
		toBase64Image: function() {
			return this.chart.canvas.toDataURL.apply(this.chart.canvas, arguments)
		}
	}), i.Type.extend = function(t) {
		var e = this,
			n = function() {
				return e.apply(this, arguments)
			};
		if (n.prototype = s(e.prototype), a(n.prototype, t), n.extend = i.Type.extend, t.name || e.prototype.name) {
			var o = t.name || e.prototype.name,
				l = i.defaults[e.prototype.name] ? s(i.defaults[e.prototype.name]) : {};
			i.defaults[o] = a(l, t.defaults), i.types[o] = n, i.prototype[o] = function(t, e) {
				var s = r(i.defaults.global, i.defaults[o], e || {});
				return new n(t, s, this)
			}
		} else u("Name not provided for this chart, so it hasn't been registered");
		return e
	}, i.Element = function(t) {
		a(this, t), this.initialize.apply(this, arguments), this.save()
	}, a(i.Element.prototype, {
		initialize: function() {},
		restore: function(t) {
			return t ? o(t, function(t) {
				this[t] = this._saved[t]
			}, this) : a(this, this._saved), this
		},
		save: function() {
			return this._saved = s(this), delete this._saved._saved, this
		},
		update: function(t) {
			return o(t, function(t, e) {
				this._saved[e] = this[e], this[e] = t
			}, this), this
		},
		transition: function(t, e) {
			return o(t, function(t, i) {
				this[i] = (t - this._saved[i]) * e + this._saved[i]
			}, this), this
		},
		tooltipPosition: function() {
			return {
				x: this.x,
				y: this.y
			}
		},
		hasValue: function() {
			return f(this.value)
		}
	}), i.Element.extend = h, i.Point = i.Element.extend({
		display: !0,
		inRange: function(t, e) {
			var i = this.hitDetectionRadius + this.radius;
			return Math.pow(t - this.x, 2) + Math.pow(e - this.y, 2) < Math.pow(i, 2)
		},
		draw: function() {
			if (this.display) {
				var t = this.ctx;
				t.beginPath(), t.arc(this.x, this.y, this.radius, 0, 2 * Math.PI), t.closePath(), t.strokeStyle = this.strokeColor, t.lineWidth = this.strokeWidth, t.fillStyle = this.fillColor, t.fill(), t.stroke()
			}
		}
	}), i.Arc = i.Element.extend({
		inRange: function(t, e) {
			var i = n.getAngleFromPoint(this, {
				x: t,
				y: e
			}),
				o = i.angle >= this.startAngle && i.angle <= this.endAngle,
				s = i.distance >= this.innerRadius && i.distance <= this.outerRadius;
			return o && s
		},
		tooltipPosition: function() {
			var t = this.startAngle + (this.endAngle - this.startAngle) / 2,
				e = (this.outerRadius - this.innerRadius) / 2 + this.innerRadius;
			return {
				x: this.x + Math.cos(t) * e,
				y: this.y + Math.sin(t) * e
			}
		},
		draw: function(t) {
			var e = this.ctx;
			if (e.beginPath(), e.arc(this.x, this.y, this.outerRadius, this.startAngle, this.endAngle), e.arc(this.x, this.y, this.innerRadius, this.endAngle, this.startAngle, !0), e.closePath(), e.strokeStyle = this.strokeColor, e.lineWidth = this.strokeWidth, e.fillStyle = this.fillColor, e.fill(), e.lineJoin = "bevel", this.showStroke && e.stroke(), this.circleBeginEnd) {
				var i = (this.outerRadius + this.innerRadius) / 2,
					n = (this.outerRadius - this.innerRadius) / 2;
				e.beginPath(), e.arc(this.x + Math.cos(this.startAngle) * i, this.y + Math.sin(this.startAngle) * i, n, 0, 2 * Math.PI), e.closePath(), e.fill(), e.beginPath(), e.arc(this.x + Math.cos(this.endAngle) * i, this.y + Math.sin(this.endAngle) * i, n, 0, 2 * Math.PI), e.closePath(), e.fill()
			}
		}
	}), i.Rectangle = i.Element.extend({
		draw: function() {
			var t = this.ctx,
				e = this.width / 2,
				i = this.x - e,
				n = this.x + e,
				o = this.base - (this.base - this.y),
				s = this.strokeWidth / 2;
			this.showStroke && (i += s, n -= s, o += s), t.beginPath(), t.fillStyle = this.fillColor, t.strokeStyle = this.strokeColor, t.lineWidth = this.strokeWidth, t.moveTo(i, this.base), t.lineTo(i, o), t.lineTo(n, o), t.lineTo(n, this.base), t.fill(), this.showStroke && t.stroke()
		},
		height: function() {
			return this.base - this.y
		},
		inRange: function(t, e) {
			return t >= this.x - this.width / 2 && t <= this.x + this.width / 2 && e >= this.y && e <= this.base
		}
	}), i.Tooltip = i.Element.extend({
		draw: function() {
			var t = this.chart.ctx;
			t.font = z(this.fontSize, this.fontStyle, this.fontFamily), this.xAlign = "center", this.yAlign = "above";
			var e = this.caretPadding = 2,
				i = t.measureText(this.text).width + 2 * this.xPadding,
				n = this.fontSize + 2 * this.yPadding,
				o = n + this.caretHeight + e;
			this.x + i / 2 > this.chart.width ? this.xAlign = "left" : this.x - i / 2 < 0 && (this.xAlign = "right"), this.y - o < 0 && (this.yAlign = "below");
			var s = this.x - i / 2,
				a = this.y - o;
			if (t.fillStyle = this.fillColor, this.custom) this.custom(this);
			else {
				switch (this.yAlign) {
				case "above":
					t.beginPath(), t.moveTo(this.x, this.y - e), t.lineTo(this.x + this.caretHeight, this.y - (e + this.caretHeight)), t.lineTo(this.x - this.caretHeight, this.y - (e + this.caretHeight)), t.closePath(), t.fill();
					break;
				case "below":
					a = this.y + e + this.caretHeight, t.beginPath(), t.moveTo(this.x, this.y + e), t.lineTo(this.x + this.caretHeight, this.y + e + this.caretHeight), t.lineTo(this.x - this.caretHeight, this.y + e + this.caretHeight), t.closePath(), t.fill()
				}
				switch (this.xAlign) {
				case "left":
					s = this.x - i + (this.cornerRadius + this.caretHeight);
					break;
				case "right":
					s = this.x - (this.cornerRadius + this.caretHeight)
				}
				$(t, s, a, i, n, this.cornerRadius), t.fill(), t.fillStyle = this.textColor, t.textAlign = "center", t.textBaseline = "middle", t.fillText(this.text, s + i / 2, a + n / 2)
			}
		}
	}), i.MultiTooltip = i.Element.extend({
		initialize: function() {
			this.font = z(this.fontSize, this.fontStyle, this.fontFamily), this.titleFont = z(this.titleFontSize, this.titleFontStyle, this.titleFontFamily), this.height = this.labels.length * this.fontSize + (this.labels.length - 1) * (this.fontSize / 2) + 2 * this.yPadding + 1.5 * this.titleFontSize, this.ctx.font = this.titleFont;
			var t = this.ctx.measureText(this.title).width,
				e = I(this.ctx, this.font, this.labels) + this.fontSize + 3,
				i = g([e, t]);
			this.width = i + 2 * this.xPadding;
			var n = this.height / 2;
			this.y - n < 0 ? this.y = n : this.y + n > this.chart.height && (this.y = this.chart.height - n), this.x > this.chart.width / 2 ? this.x -= this.xOffset + this.width : this.x += this.xOffset
		},
		getLineHeight: function(t) {
			var e = this.y - this.height / 2 + this.yPadding,
				i = t - 1;
			return 0 === t ? e + this.titleFontSize / 2 : e + (1.5 * this.fontSize * i + this.fontSize / 2) + 1.5 * this.titleFontSize
		},
		draw: function() {
			if (this.custom) this.custom(this);
			else {
				$(this.ctx, this.x, this.y - this.height / 2, this.width, this.height, this.cornerRadius);
				var t = this.ctx;
				t.fillStyle = this.fillColor, t.fill(), t.closePath(), t.textAlign = "left", t.textBaseline = "middle", t.fillStyle = this.titleTextColor, t.font = this.titleFont, t.fillText(this.title, this.x + this.xPadding, this.getLineHeight(0)), t.font = this.font, n.each(this.labels, function(e, i) {
					t.fillStyle = this.textColor, t.fillText(e, this.x + this.xPadding + this.fontSize + 3, this.getLineHeight(i + 1)), t.fillStyle = this.legendColorBackground, t.fillRect(this.x + this.xPadding, this.getLineHeight(i + 1) - this.fontSize / 2, this.fontSize, this.fontSize), t.fillStyle = this.legendColors[i].fill, t.fillRect(this.x + this.xPadding, this.getLineHeight(i + 1) - this.fontSize / 2, this.fontSize, this.fontSize)
				}, this)
			}
		}
	}), i.Scale = i.Element.extend({
		initialize: function() {
			this.fit()
		},
		buildYLabels: function() {
			this.yLabels = [];
			for (var t = v(this.stepValue), e = 0; e <= this.steps; e++) this.yLabels.push(x(this.templateString, {
				value: (this.min + e * this.stepValue).toFixed(t)
			}));
			this.yLabelWidth = this.display && this.showLabels ? I(this.ctx, this.font, this.yLabels) : 0
		},
		addXLabel: function(t) {
			this.xLabels.push(t), this.valuesCount++, this.fit()
		},
		removeXLabel: function() {
			this.xLabels.shift(), this.valuesCount--, this.fit()
		},
		fit: function() {
			this.startPoint = this.display ? this.fontSize : 0, this.endPoint = this.display ? this.height - 1.5 * this.fontSize - 5 : this.height, this.startPoint += this.padding, this.endPoint -= this.padding;
			var t, e = this.endPoint - this.startPoint;
			for (this.calculateYRange(e), this.buildYLabels(), this.calculateXLabelRotation(); e > this.endPoint - this.startPoint;) e = this.endPoint - this.startPoint, t = this.yLabelWidth, this.calculateYRange(e), this.buildYLabels(), t < this.yLabelWidth && this.calculateXLabelRotation()
		},
		calculateXLabelRotation: function() {
			this.ctx.font = this.font;
			var t, e, i = this.ctx.measureText(this.xLabels[0]).width,
				n = this.ctx.measureText(this.xLabels[this.xLabels.length - 1]).width;
			if (this.xScalePaddingRight = n / 2 + 3, this.xScalePaddingLeft = i / 2 > this.yLabelWidth + 10 ? i / 2 : this.yLabelWidth + 10, this.xLabelRotation = 0, this.display) {
				var o, s = I(this.ctx, this.font, this.xLabels);
				this.xLabelWidth = s;
				for (var a = Math.floor(this.calculateX(1) - this.calculateX(0)) - 6; this.xLabelWidth > a && 0 === this.xLabelRotation || this.xLabelWidth > a && this.xLabelRotation <= 90 && this.xLabelRotation > 0;) o = Math.cos(y(this.xLabelRotation)), t = o * i, e = o * n, t + this.fontSize / 2 > this.yLabelWidth + 8 && (this.xScalePaddingLeft = t + this.fontSize / 2), this.xScalePaddingRight = this.fontSize / 2, this.xLabelRotation++, this.xLabelWidth = o * s;
				this.xLabelRotation > 0 && (this.endPoint -= Math.sin(y(this.xLabelRotation)) * s + 3)
			} else this.xLabelWidth = 0, this.xScalePaddingRight = this.padding, this.xScalePaddingLeft = this.padding
		},
		calculateYRange: c,
		drawingArea: function() {
			return this.startPoint - this.endPoint
		},
		calculateY: function(t) {
			var e = this.drawingArea() / (this.min - this.max);
			return this.endPoint - e * (t - this.min)
		},
		calculateX: function(t) {
			var e = (this.xLabelRotation > 0, this.width - (this.xScalePaddingLeft + this.xScalePaddingRight)),
				i = e / Math.max(this.valuesCount - (this.offsetGridLines ? 0 : 1), 1),
				n = i * t + this.xScalePaddingLeft;
			return this.offsetGridLines && (n += i / 2), Math.round(n)
		},
		update: function(t) {
			n.extend(this, t), this.fit()
		},
		draw: function() {
			var t = this.ctx,
				e = (this.endPoint - this.startPoint) / this.steps,
				i = Math.round(this.xScalePaddingLeft);
			if (this.display) {
				t.fillStyle = this.textColor, t.font = this.font;
				var s = this.showBeyondLine ? 5 : 0;
				o(this.yLabels, function(o, a) {
					var r = this.endPoint - e * a,
						l = Math.round(r),
						h = this.showHorizontalLines;
					t.textAlign = "right", t.textBaseline = "middle", this.showLabels && t.fillText(o, i - 10, r), 0 !== a || h || (h = !0), h && t.beginPath(), a > 0 ? (t.lineWidth = this.gridLineWidth, t.strokeStyle = this.gridLineColor) : (t.lineWidth = this.lineWidth, t.strokeStyle = this.lineColor), l += n.aliasPixel(t.lineWidth), h && (t.moveTo(i, l), t.lineTo(this.width, l), t.stroke(), t.closePath()), t.lineWidth = this.lineWidth, t.strokeStyle = this.lineColor, t.beginPath(), t.moveTo(i - s, l), t.lineTo(i, l), t.stroke(), t.closePath()
				}, this), o(this.xLabels, function(e, i) {
					var n = this.calculateX(i) + b(this.lineWidth),
						o = this.calculateX(i - (this.offsetGridLines ? .5 : 0)) + b(this.lineWidth),
						a = this.xLabelRotation > 0,
						r = this.showVerticalLines;
					0 !== i || r || (r = !0), r && t.beginPath(), i > 0 ? (t.lineWidth = this.gridLineWidth, t.strokeStyle = this.gridLineColor) : (t.lineWidth = this.lineWidth, t.strokeStyle = this.lineColor), r && (t.moveTo(o, this.endPoint), t.lineTo(o, this.startPoint - 3), t.stroke(), t.closePath()), t.lineWidth = this.lineWidth, t.strokeStyle = this.lineColor, t.beginPath(), t.moveTo(o, this.endPoint), t.lineTo(o, this.endPoint + s), t.stroke(), t.closePath(), t.save(), t.translate(n, a ? this.endPoint + 12 : this.endPoint + 8), t.rotate(y(this.xLabelRotation) * -1), t.font = this.font, t.textAlign = a ? "right" : "center", t.textBaseline = a ? "middle" : "top", t.fillText(e, 0, 0), t.restore()
				}, this)
			}
		}
	}), i.RadialScale = i.Element.extend({
		initialize: function() {
			this.size = m([this.height, this.width]), this.drawingArea = this.display ? this.size / 2 - (this.fontSize / 2 + this.backdropPaddingY) : this.size / 2
		},
		calculateCenterOffset: function(t) {
			var e = this.drawingArea / (this.max - this.min);
			return (t - this.min) * e
		},
		update: function() {
			this.lineArc ? this.drawingArea = this.display ? this.size / 2 - (this.fontSize / 2 + this.backdropPaddingY) : this.size / 2 : this.setScaleSize(), this.buildYLabels()
		},
		buildYLabels: function() {
			this.yLabels = [];
			for (var t = v(this.stepValue), e = 0; e <= this.steps; e++) this.yLabels.push(x(this.templateString, {
				value: (this.min + e * this.stepValue).toFixed(t)
			}))
		},
		getCircumference: function() {
			return 2 * Math.PI / this.valuesCount
		},
		setScaleSize: function() {
			var t, e, i, n, o, s, a, r, l, h, c, d, u = m([this.height / 2 - this.pointLabelFontSize - 5, this.width / 2]),
				p = this.width,
				g = 0;
			for (this.ctx.font = z(this.pointLabelFontSize, this.pointLabelFontStyle, this.pointLabelFontFamily), e = 0; e < this.valuesCount; e++) t = this.getPointPosition(e, u), i = this.ctx.measureText(x(this.templateString, {
				value: this.labels[e]
			})).width + 5, 0 === e || e === this.valuesCount / 2 ? (n = i / 2, t.x + n > p && (p = t.x + n, o = e), t.x - n < g && (g = t.x - n, a = e)) : e < this.valuesCount / 2 ? t.x + i > p && (p = t.x + i, o = e) : e > this.valuesCount / 2 && t.x - i < g && (g = t.x - i, a = e);
			l = g, h = Math.ceil(p - this.width), s = this.getIndexAngle(o), r = this.getIndexAngle(a), c = h / Math.sin(s + Math.PI / 2), d = l / Math.sin(r + Math.PI / 2), c = f(c) ? c : 0, d = f(d) ? d : 0, this.drawingArea = u - (d + c) / 2, this.setCenterPoint(d, c)
		},
		setCenterPoint: function(t, e) {
			var i = this.width - e - this.drawingArea,
				n = t + this.drawingArea;
			this.xCenter = (n + i) / 2, this.yCenter = this.height / 2
		},
		getIndexAngle: function(t) {
			var e = 2 * Math.PI / this.valuesCount;
			return t * e - Math.PI / 2
		},
		getPointPosition: function(t, e) {
			var i = this.getIndexAngle(t);
			return {
				x: Math.cos(i) * e + this.xCenter,
				y: Math.sin(i) * e + this.yCenter
			}
		},
		draw: function() {
			if (this.display) {
				var t = this.ctx;
				if (o(this.yLabels, function(e, i) {
					if (i > 0) {
						var n, o = i * (this.drawingArea / this.steps),
							s = this.yCenter - o;
						if (this.lineWidth > 0) if (t.strokeStyle = this.lineColor, t.lineWidth = this.lineWidth, this.lineArc) t.beginPath(), t.arc(this.xCenter, this.yCenter, o, 0, 2 * Math.PI), t.closePath(), t.stroke();
						else {
							t.beginPath();
							for (var a = 0; a < this.valuesCount; a++) n = this.getPointPosition(a, this.calculateCenterOffset(this.min + i * this.stepValue)), 0 === a ? t.moveTo(n.x, n.y) : t.lineTo(n.x, n.y);
							t.closePath(), t.stroke()
						}
						if (this.showLabels) {
							if (t.font = z(this.fontSize, this.fontStyle, this.fontFamily), this.showLabelBackdrop) {
								var r = t.measureText(e).width;
								t.fillStyle = this.backdropColor, t.fillRect(this.xCenter - r / 2 - this.backdropPaddingX, s - this.fontSize / 2 - this.backdropPaddingY, r + 2 * this.backdropPaddingX, this.fontSize + 2 * this.backdropPaddingY)
							}
							t.textAlign = "center", t.textBaseline = "middle", t.fillStyle = this.fontColor, t.fillText(e, this.xCenter, s)
						}
					}
				}, this), !this.lineArc) {
					t.lineWidth = this.angleLineWidth, t.strokeStyle = this.angleLineColor;
					for (var e = this.valuesCount - 1; e >= 0; e--) {
						if (this.angleLineWidth > 0) {
							var i = this.getPointPosition(e, this.calculateCenterOffset(this.max));
							t.beginPath(), t.moveTo(this.xCenter, this.yCenter), t.lineTo(i.x, i.y), t.stroke(), t.closePath()
						}
						var n = this.getPointPosition(e, this.calculateCenterOffset(this.max) + 5);
						t.font = z(this.pointLabelFontSize, this.pointLabelFontStyle, this.pointLabelFontFamily), t.fillStyle = this.pointLabelFontColor;
						var s = this.labels.length,
							a = this.labels.length / 2,
							r = a / 2,
							l = e < r || e > s - r,
							h = e === r || e === s - r;
						0 === e ? t.textAlign = "center" : e === a ? t.textAlign = "center" : e < a ? t.textAlign = "left" : t.textAlign = "right", h ? t.textBaseline = "middle" : l ? t.textBaseline = "bottom" : t.textBaseline = "top", t.fillText(this.labels[e], n.x, n.y)
					}
				}
			}
		}
	}), n.addEvent(window, "resize", function() {
		var t;
		return function() {
			clearTimeout(t), t = setTimeout(function() {
				o(i.instances, function(t) {
					t.options.responsive && t.resize(t.render, !0)
				})
			}, 50)
		}
	}()), p ? define(function() {
		return i
	}) : "object" == typeof module && module.exports && (module.exports = i), e.Chart = i, t.fn.chart = function() {
		var t = [];
		return this.each(function() {
			t.push(new i(this.getContext("2d")))
		}), 1 === t.length ? t[0] : t
	}
}.call(this, jQuery), function(t) {
	"use strict";
	var e = t && t.zui ? t.zui : this,
		i = e.Chart,
		n = i.helpers,
		o = {
			scaleShowGridLines: !0,
			scaleGridLineColor: "rgba(0,0,0,.05)",
			scaleGridLineWidth: 1,
			scaleShowHorizontalLines: !0,
			scaleShowBeyondLine: !0,
			scaleShowVerticalLines: !0,
			bezierCurve: !0,
			bezierCurveTension: .4,
			pointDot: !0,
			pointDotRadius: 4,
			pointDotStrokeWidth: 1,
			pointHitDetectionRadius: 20,
			datasetStroke: !0,
			datasetStrokeWidth: 2,
			datasetFill: !0,
			legendTemplate: '<ul class="<%=name.toLowerCase()%>-legend"><% for (var i=0; i<datasets.length; i++){%><li><span style="background-color:<%=datasets[i].strokeColor%>"></span><%if(datasets[i].label){%><%=datasets[i].label%><%}%></li><%}%></ul>'
		};
	i.Type.extend({
		name: "Line",
		defaults: o,
		initialize: function(e) {
			this.PointClass = i.Point.extend({
				strokeWidth: this.options.pointDotStrokeWidth,
				radius: this.options.pointDotRadius,
				display: this.options.pointDot,
				hitDetectionRadius: this.options.pointHitDetectionRadius,
				ctx: this.chart.ctx,
				inRange: function(t) {
					return Math.pow(t - this.x, 2) < Math.pow(this.radius + this.hitDetectionRadius, 2)
				}
			}), this.datasets = [], this.options.showTooltips && n.bindEvents(this, this.options.tooltipEvents, function(t) {
				var e = "mouseout" !== t.type ? this.getPointsAtEvent(t) : [];
				this.eachPoints(function(t) {
					t.restore(["fillColor", "strokeColor"])
				}), n.each(e, function(t) {
					t.fillColor = t.highlightFill, t.strokeColor = t.highlightStroke
				}), this.showTooltip(e)
			}), n.each(e.datasets, function(i) {
				if (t.zui && t.zui.Color && t.zui.Color.get) {
					var o = t.zui.Color.get(i.color),
						s = o.toCssStr();
					i.fillColor || (i.fillColor = o.clone().fade(20).toCssStr()), i.strokeColor || (i.strokeColor = s), i.pointColor || (i.pointColor = s), i.pointStrokeColor || (i.pointStrokeColor = "#fff"), i.pointHighlightFill || (i.pointHighlightFill = "#fff"), i.pointHighlightStroke || (i.pointHighlightStroke = s)
				}
				var a = {
					label: i.label || null,
					fillColor: i.fillColor,
					strokeColor: i.strokeColor,
					pointColor: i.pointColor,
					pointStrokeColor: i.pointStrokeColor,
					showTooltips: i.showTooltips !== !1,
					points: []
				};
				this.datasets.push(a), n.each(i.data, function(t, n) {
					a.points.push(new this.PointClass({
						value: t,
						label: e.labels[n],
						datasetLabel: i.label,
						strokeColor: i.pointStrokeColor,
						fillColor: i.pointColor,
						highlightFill: i.pointHighlightFill || i.pointColor,
						highlightStroke: i.pointHighlightStroke || i.pointStrokeColor
					}))
				}, this), this.buildScale(e.labels), this.eachPoints(function(t, e) {
					n.extend(t, {
						x: this.scale.calculateX(e),
						y: this.scale.endPoint
					}), t.save()
				}, this)
			}, this), this.render()
		},
		update: function() {
			this.scale.update(), n.each(this.activeElements, function(t) {
				t.restore(["fillColor", "strokeColor"])
			}), this.eachPoints(function(t) {
				t.save()
			}), this.render()
		},
		eachPoints: function(t) {
			n.each(this.datasets, function(e) {
				n.each(e.points, t, this)
			}, this)
		},
		getPointsAtEvent: function(t) {
			var e = [],
				i = n.getRelativePosition(t);
			return n.each(this.datasets, function(t) {
				n.each(t.points, function(t) {
					t.inRange(i.x, i.y) && e.push(t)
				})
			}, this), e
		},
		buildScale: function(t) {
			var e = this,
				o = function() {
					var t = [];
					return e.eachPoints(function(e) {
						t.push(e.value)
					}), t
				},
				s = {
					templateString: this.options.scaleLabel,
					height: this.chart.height,
					width: this.chart.width,
					ctx: this.chart.ctx,
					textColor: this.options.scaleFontColor,
					fontSize: this.options.scaleFontSize,
					fontStyle: this.options.scaleFontStyle,
					fontFamily: this.options.scaleFontFamily,
					valuesCount: t.length,
					beginAtZero: this.options.scaleBeginAtZero,
					integersOnly: this.options.scaleIntegersOnly,
					calculateYRange: function(t) {
						var e = n.calculateScaleRange(o(), t, this.fontSize, this.beginAtZero, this.integersOnly);
						n.extend(this, e)
					},
					xLabels: t,
					font: n.fontString(this.options.scaleFontSize, this.options.scaleFontStyle, this.options.scaleFontFamily),
					lineWidth: this.options.scaleLineWidth,
					lineColor: this.options.scaleLineColor,
					showHorizontalLines: this.options.scaleShowHorizontalLines,
					showVerticalLines: this.options.scaleShowVerticalLines,
					showBeyondLine: this.options.scaleShowBeyondLine,
					gridLineWidth: this.options.scaleShowGridLines ? this.options.scaleGridLineWidth : 0,
					gridLineColor: this.options.scaleShowGridLines ? this.options.scaleGridLineColor : "rgba(0,0,0,0)",
					padding: this.options.showScale ? 0 : this.options.pointDotRadius + this.options.pointDotStrokeWidth,
					showLabels: this.options.scaleShowLabels,
					display: this.options.showScale
				};
			this.options.scaleOverride && n.extend(s, {
				calculateYRange: n.noop,
				steps: this.options.scaleSteps,
				stepValue: this.options.scaleStepWidth,
				min: this.options.scaleStartValue,
				max: this.options.scaleStartValue + this.options.scaleSteps * this.options.scaleStepWidth
			}), this.scale = new i.Scale(s)
		},
		addData: function(t, e) {
			n.each(t, function(t, i) {
				this.datasets[i].points.push(new this.PointClass({
					value: t,
					label: e,
					datasetLabel: this.datasets[i].label,
					x: this.scale.calculateX(this.scale.valuesCount + 1),
					y: this.scale.endPoint,
					strokeColor: this.datasets[i].pointStrokeColor,
					fillColor: this.datasets[i].pointColor
				}))
			}, this), this.scale.addXLabel(e), this.update()
		},
		removeData: function() {
			this.scale.removeXLabel(), n.each(this.datasets, function(t) {
				t.points.shift()
			}, this), this.update()
		},
		reflow: function() {
			var t = n.extend({
				height: this.chart.height,
				width: this.chart.width
			});
			this.scale.update(t)
		},
		draw: function(t) {
			var e = t || 1;
			this.clear();
			var i = this.chart.ctx,
				o = function(t) {
					return null !== t.value
				},
				s = function(t, e, i) {
					return n.findNextWhere(e, o, i) || t
				},
				a = function(t, e, i) {
					return n.findPreviousWhere(e, o, i) || t
				};
			this.scale.draw(e), n.each(this.datasets, function(t) {
				var r = n.where(t.points, o);
				n.each(t.points, function(t, i) {
					t.hasValue() && t.transition({
						y: this.scale.calculateY(t.value),
						x: this.scale.calculateX(i)
					}, e)
				}, this), this.options.bezierCurve && n.each(r, function(t, e) {
					var i = e > 0 && e < r.length - 1 ? this.options.bezierCurveTension : 0;
					t.controlPoints = n.splineCurve(a(t, r, e), t, s(t, r, e), i), t.controlPoints.outer.y > this.scale.endPoint ? t.controlPoints.outer.y = this.scale.endPoint : t.controlPoints.outer.y < this.scale.startPoint && (t.controlPoints.outer.y = this.scale.startPoint), t.controlPoints.inner.y > this.scale.endPoint ? t.controlPoints.inner.y = this.scale.endPoint : t.controlPoints.inner.y < this.scale.startPoint && (t.controlPoints.inner.y = this.scale.startPoint)
				}, this), i.lineWidth = this.options.datasetStrokeWidth, i.strokeStyle = t.strokeColor, i.beginPath(), n.each(r, function(t, e) {
					if (0 === e) i.moveTo(t.x, t.y);
					else if (this.options.bezierCurve) {
						var n = a(t, r, e);
						i.bezierCurveTo(n.controlPoints.outer.x, n.controlPoints.outer.y, t.controlPoints.inner.x, t.controlPoints.inner.y, t.x, t.y)
					} else i.lineTo(t.x, t.y)
				}, this), i.stroke(), this.options.datasetFill && r.length > 0 && (i.lineTo(r[r.length - 1].x, this.scale.endPoint), i.lineTo(r[0].x, this.scale.endPoint), i.fillStyle = t.fillColor, i.closePath(), i.fill()), n.each(r, function(t) {
					t.draw()
				})
			}, this)
		}
	}), t.fn.lineChart = function(e, n) {
		var o = [];
		return this.each(function() {
			var s = t(this);
			o.push(new i(this.getContext("2d")).Line(e, t.extend(s.data(), n)))
		}), 1 === o.length ? o[0] : o
	}
}.call(this, jQuery), function(t) {
	"use strict";
	var e = t && t.zui ? t.zui : this,
		i = e.Chart,
		n = i.helpers,
		o = {
			segmentShowStroke: !0,
			segmentStrokeColor: "#fff",
			segmentStrokeWidth: 1,
			percentageInnerCutout: 50,
			scaleShowLabels: !1,
			scaleLabel: "<%=value%>",
			scaleLabelPlacement: "auto",
			animationSteps: 60,
			animationEasing: "easeOutBounce",
			animateRotate: !0,
			animateScale: !1,
			legendTemplate: '<ul class="<%=name.toLowerCase()%>-legend"><% for (var i=0; i<segments.length; i++){%><li><span style="background-color:<%=segments[i].fillColor%>"></span><%if(segments[i].label){%><%=segments[i].label%><%}%></li><%}%></ul>'
		};
	i.Type.extend({
		name: "Doughnut",
		defaults: o,
		initialize: function(t) {
			this.segments = [], this.outerRadius = (n.min([this.chart.width, this.chart.height]) - this.options.segmentStrokeWidth / 2) / 2, this.SegmentArc = i.Arc.extend({
				ctx: this.chart.ctx,
				x: this.chart.width / 2,
				y: this.chart.height / 2
			}), this.options.showTooltips && n.bindEvents(this, this.options.tooltipEvents, function(t) {
				var e = "mouseout" !== t.type ? this.getSegmentsAtEvent(t) : [];
				n.each(this.segments, function(t) {
					t.restore(["fillColor"])
				}), n.each(e, function(t) {
					t.fillColor = t.highlightColor
				}), this.showTooltip(e)
			}), this.calculateTotal(t), n.each(t, function(t, e) {
				this.addData(t, e, !0)
			}, this), this.render()
		},
		getSegmentsAtEvent: function(t) {
			var e = [],
				i = n.getRelativePosition(t);
			return n.each(this.segments, function(t) {
				t.inRange(i.x, i.y) && e.push(t)
			}, this), e
		},
		addData: function(e, i, n) {
			if (t.zui && t.zui.Color && t.zui.Color.get) {
				var o = new t.zui.Color.get(e.color);
				e.color = o.toCssStr(), e.highlight || (e.highlight = o.lighten(5).toCssStr())
			}
			var s = i || this.segments.length;
			this.segments.splice(s, 0, new this.SegmentArc({
				id: "undefined" == typeof e.id ? s : e.id,
				value: e.value,
				outerRadius: this.options.animateScale ? 0 : this.outerRadius,
				innerRadius: this.options.animateScale ? 0 : this.outerRadius / 100 * this.options.percentageInnerCutout,
				fillColor: e.color,
				highlightColor: e.highlight || e.color,
				showStroke: this.options.segmentShowStroke,
				strokeWidth: this.options.segmentStrokeWidth,
				strokeColor: this.options.segmentStrokeColor,
				startAngle: 1.5 * Math.PI,
				circumference: this.options.animateRotate ? 0 : this.calculateCircumference(e.value),
				showLabel: e.showLabel !== !1,
				circleBeginEnd: e.circleBeginEnd,
				label: e.label
			})), n || (this.reflow(), this.update())
		},
		calculateCircumference: function(t) {
			return 2 * Math.PI * (Math.abs(t) / this.total)
		},
		calculateTotal: function(t) {
			this.total = 0, n.each(t, function(t) {
				this.total += Math.abs(t.value)
			}, this)
		},
		update: function() {
			this.calculateTotal(this.segments), n.each(this.activeElements, function(t) {
				t.restore(["fillColor"])
			}), n.each(this.segments, function(t) {
				t.save()
			}), this.render()
		},
		removeData: function(t) {
			var e = n.isNumber(t) ? t : this.segments.length - 1;
			this.segments.splice(e, 1), this.reflow(), this.update()
		},
		reflow: function() {
			n.extend(this.SegmentArc.prototype, {
				x: this.chart.width / 2,
				y: this.chart.height / 2
			}), this.outerRadius = (n.min([this.chart.width, this.chart.height]) - this.options.segmentStrokeWidth / 2) / 2, n.each(this.segments, function(t) {
				t.update({
					outerRadius: this.outerRadius,
					innerRadius: this.outerRadius / 100 * this.options.percentageInnerCutout
				})
			}, this)
		},
		drawLabel: function(e, i, o) {
			var s = this.options,
				a = (e.endAngle + e.startAngle) / 2,
				r = s.scaleLabelPlacement;
			"inside" !== r && "outside" !== r && this.chart.width - this.chart.height > 50 && e.circumference < Math.PI / 18 && (r = "outside");
			var l = Math.cos(a) * e.outerRadius,
				h = Math.sin(a) * e.outerRadius,
				c = n.template(s.scaleLabel, {
					value: "undefined" == typeof i ? e.value : Math.round(i * e.value),
					label: e.label
				}),
				d = this.chart.ctx;
			d.font = n.fontString(s.scaleFontSize, s.scaleFontStyle, s.scaleFontFamily), d.textBaseline = "middle", d.textAlign = "center";
			var u = (d.measureText(c).width, this.chart.width / 2),
				p = this.chart.height / 2;
			if ("outside" === r) {
				var f = l >= 0,
					g = l + u,
					m = h + p;
				d.textAlign = f ? "left" : "right", d.measureText(c).width, l = f ? Math.max(u + e.outerRadius + 10, l + 30 + u) : Math.min(u - e.outerRadius - 10, l - 30 + u);
				var v = s.scaleFontSize * (s.scaleLineHeight || 1),
					y = Math.round((.8 * h + p) / v) + 1,
					b = (Math.floor(this.chart.width / v) + 1, f ? 1 : -1);
				if (o[y * b] && (y > 1 ? y-- : y++), o[y * b]) return;
				h = (y - 1) * v + s.scaleFontSize / 2, o[y * b] = !0, d.beginPath(), d.moveTo(g, m), d.lineTo(l, h), l = f ? l + 5 : l - 5, d.lineTo(l, h), d.strokeStyle = t.zui && t.zui.Color ? new t.zui.Color(e.fillColor).fade(40).toCssStr() : e.fillColor, d.strokeWidth = s.scaleLineWidth, d.stroke(), d.fillStyle = e.fillColor
			} else l = .7 * l + u, h = .7 * h + p, d.fillStyle = t.zui && t.zui.Color ? new t.zui.Color(e.fillColor).contrast().toCssStr() : "#fff";
			d.fillText(c, l, h)
		},
		draw: function(t) {
			var e = t ? t : 1;
			this.clear();
			var i;
			if (n.each(this.segments, function(t, i) {
				t.transition({
					circumference: this.calculateCircumference(t.value),
					outerRadius: this.outerRadius,
					innerRadius: this.outerRadius / 100 * this.options.percentageInnerCutout
				}, e), t.endAngle = t.startAngle + t.circumference, this.options.reverseDrawOrder || t.draw(), 0 === i && (t.startAngle = 1.5 * Math.PI), i < this.segments.length - 1 && (this.segments[i + 1].startAngle = t.endAngle)
			}, this), this.options.reverseDrawOrder && n.each(this.segments.slice().reverse(), function(t, e) {
				t.draw()
			}, this), this.options.scaleShowLabels) {
				var o = this.segments.slice().sort(function(t, e) {
					return e.value - t.value
				}),
					i = {};
				n.each(o, function(e, n) {
					e.showLabel && this.drawLabel(e, t, i)
				}, this)
			}
		}
	}), i.types.Doughnut.extend({
		name: "Pie",
		defaults: n.merge(o, {
			percentageInnerCutout: 0
		})
	}), t.fn.pieChart = function(e, n) {
		var o = [];
		return this.each(function() {
			var s = t(this);
			o.push(new i(this.getContext("2d")).Pie(e, t.extend(s.data(), n)))
		}), 1 === o.length ? o[0] : o
	}, t.fn.doughnutChart = function(e, n) {
		var o = [];
		return this.each(function() {
			var s = t(this);
			o.push(new i(this.getContext("2d")).Doughnut(e, t.extend(s.data(), n)))
		}), 1 === o.length ? o[0] : o
	}
}.call(this, jQuery), function(t) {
	"use strict";
	var e = t && t.zui ? t.zui : this,
		i = e.Chart,
		n = i.helpers,
		o = {
			scaleBeginAtZero: !0,
			scaleShowGridLines: !0,
			scaleGridLineColor: "rgba(0,0,0,.05)",
			scaleGridLineWidth: 1,
			scaleShowHorizontalLines: !0,
			scaleShowVerticalLines: !0,
			scaleShowBeyondLine: !0,
			barShowStroke: !0,
			barStrokeWidth: 1,
			scaleValuePlacement: "auto",
			barValueSpacing: 5,
			barDatasetSpacing: 1,
			legendTemplate: '<ul class="<%=name.toLowerCase()%>-legend"><% for (var i=0; i<datasets.length; i++){%><li><span style="background-color:<%=datasets[i].fillColor%>"></span><%if(datasets[i].label){%><%=datasets[i].label%><%}%></li><%}%></ul>'
		};
	i.Type.extend({
		name: "Bar",
		defaults: o,
		initialize: function(e) {
			var o = this.options;
			this.ScaleClass = i.Scale.extend({
				offsetGridLines: !0,
				calculateBarX: function(t, e, i) {
					var n = this.calculateBaseWidth(),
						s = this.calculateX(i) - n / 2,
						a = this.calculateBarWidth(t);
					return s + a * e + e * o.barDatasetSpacing + a / 2
				},
				calculateBaseWidth: function() {
					return this.calculateX(1) - this.calculateX(0) - 2 * o.barValueSpacing
				},
				calculateBarWidth: function(t) {
					var e = this.calculateBaseWidth() - (t - 1) * o.barDatasetSpacing;
					return e / t
				}
			}), this.datasets = [], this.options.showTooltips && n.bindEvents(this, this.options.tooltipEvents, function(t) {
				var e = "mouseout" !== t.type ? this.getBarsAtEvent(t) : [];
				this.eachBars(function(t) {
					t.restore(["fillColor", "strokeColor"])
				}), n.each(e, function(t) {
					t.fillColor = t.highlightFill, t.strokeColor = t.highlightStroke
				}), this.showTooltip(e)
			}), this.BarClass = i.Rectangle.extend({
				strokeWidth: this.options.barStrokeWidth,
				showStroke: this.options.barShowStroke,
				ctx: this.chart.ctx
			}), n.each(e.datasets, function(i, o) {
				if (t.zui && t.zui.Color && t.zui.Color.get) {
					var s = t.zui.Color.get(i.color),
						a = s.toCssStr();
					i.fillColor || (i.fillColor = s.clone().fade(50).toCssStr()), i.strokeColor || (i.strokeColor = a)
				}
				var r = {
					label: i.label || null,
					fillColor: i.fillColor,
					strokeColor: i.strokeColor,
					bars: []
				};
				this.datasets.push(r), n.each(i.data, function(t, n) {
					r.bars.push(new this.BarClass({
						value: t,
						label: e.labels[n],
						datasetLabel: i.label,
						strokeColor: i.strokeColor,
						fillColor: i.fillColor,
						highlightFill: i.highlightFill || i.fillColor,
						highlightStroke: i.highlightStroke || i.strokeColor
					}))
				}, this)
			}, this), this.buildScale(e.labels), this.BarClass.prototype.base = this.scale.endPoint, this.eachBars(function(t, e, i) {
				n.extend(t, {
					width: this.scale.calculateBarWidth(this.datasets.length),
					x: this.scale.calculateBarX(this.datasets.length, i, e),
					y: this.scale.endPoint
				}), t.save()
			}, this), this.render()
		},
		update: function() {
			this.scale.update(), n.each(this.activeElements, function(t) {
				t.restore(["fillColor", "strokeColor"])
			}), this.eachBars(function(t) {
				t.save()
			}), this.render()
		},
		eachBars: function(t) {
			n.each(this.datasets, function(e, i) {
				n.each(e.bars, t, this, i)
			}, this)
		},
		getBarsAtEvent: function(t) {
			for (var e, i = [], o = n.getRelativePosition(t), s = function(t) {
					i.push(t.bars[e])
				}, a = 0; a < this.datasets.length; a++) for (e = 0; e < this.datasets[a].bars.length; e++) if (this.datasets[a].bars[e].inRange(o.x, o.y)) return n.each(this.datasets, s), i;
			return i
		},
		buildScale: function(t) {
			var e = this,
				i = function() {
					var t = [];
					return e.eachBars(function(e) {
						t.push(e.value)
					}), t
				},
				o = {
					templateString: this.options.scaleLabel,
					height: this.chart.height,
					width: this.chart.width,
					ctx: this.chart.ctx,
					textColor: this.options.scaleFontColor,
					fontSize: this.options.scaleFontSize,
					fontStyle: this.options.scaleFontStyle,
					fontFamily: this.options.scaleFontFamily,
					valuesCount: t.length,
					beginAtZero: this.options.scaleBeginAtZero,
					integersOnly: this.options.scaleIntegersOnly,
					calculateYRange: function(t) {
						var e = n.calculateScaleRange(i(), t, this.fontSize, this.beginAtZero, this.integersOnly);
						n.extend(this, e)
					},
					xLabels: t,
					font: n.fontString(this.options.scaleFontSize, this.options.scaleFontStyle, this.options.scaleFontFamily),
					lineWidth: this.options.scaleLineWidth,
					lineColor: this.options.scaleLineColor,
					showHorizontalLines: this.options.scaleShowHorizontalLines,
					showVerticalLines: this.options.scaleShowVerticalLines,
					showBeyondLine: this.options.scaleShowBeyondLine,
					gridLineWidth: this.options.scaleShowGridLines ? this.options.scaleGridLineWidth : 0,
					gridLineColor: this.options.scaleShowGridLines ? this.options.scaleGridLineColor : "rgba(0,0,0,0)",
					padding: this.options.showScale ? 0 : this.options.barShowStroke ? this.options.barStrokeWidth : 0,
					showLabels: this.options.scaleShowLabels,
					display: this.options.showScale
				};
			this.options.scaleOverride && n.extend(o, {
				calculateYRange: n.noop,
				steps: this.options.scaleSteps,
				stepValue: this.options.scaleStepWidth,
				min: this.options.scaleStartValue,
				max: this.options.scaleStartValue + this.options.scaleSteps * this.options.scaleStepWidth
			}), this.scale = new this.ScaleClass(o)
		},
		addData: function(t, e) {
			n.each(t, function(t, i) {
				this.datasets[i].bars.push(new this.BarClass({
					value: t,
					label: e,
					x: this.scale.calculateBarX(this.datasets.length, i, this.scale.valuesCount + 1),
					y: this.scale.endPoint,
					width: this.scale.calculateBarWidth(this.datasets.length),
					base: this.scale.endPoint,
					strokeColor: this.datasets[i].strokeColor,
					fillColor: this.datasets[i].fillColor
				}))
			}, this), this.scale.addXLabel(e), this.update()
		},
		removeData: function() {
			this.scale.removeXLabel(), n.each(this.datasets, function(t) {
				t.bars.shift()
			}, this), this.update()
		},
		reflow: function() {
			n.extend(this.BarClass.prototype, {
				y: this.scale.endPoint,
				base: this.scale.endPoint
			});
			var t = n.extend({
				height: this.chart.height,
				width: this.chart.width
			});
			this.scale.update(t)
		},
		drawLabel: function(t, e) {
			var i = this.options;
			e = e || i.scaleValuePlacement, e = e ? e.toLowerCase() : "auto", "auto" === e && (e = t.y < 15 ? "insdie" : "outside");
			var o = "insdie" === e ? t.y + 10 : t.y - 10,
				s = this.chart.ctx;
			s.font = n.fontString(i.scaleFontSize, i.scaleFontStyle, i.scaleFontFamily), s.textBaseline = "middle", s.textAlign = "center", s.fillStyle = i.scaleFontColor, s.fillText(t.value, t.x, o)
		},
		draw: function(t) {
			var e = t || 1;
			this.clear();
			this.chart.ctx;
			this.scale.draw(e);
			var i = this.options.scaleShowLabels && this.options.scaleValuePlacement;
			n.each(this.datasets, function(t, o) {
				n.each(t.bars, function(t, n) {
					t.hasValue() && (t.base = this.scale.endPoint, t.transition({
						x: this.scale.calculateBarX(this.datasets.length, o, n),
						y: this.scale.calculateY(t.value),
						width: this.scale.calculateBarWidth(this.datasets.length)
					}, e).draw()), i && this.drawLabel(t)
				}, this)
			}, this)
		}
	}), t.fn.barChart = function(e, n) {
		var o = [];
		return this.each(function() {
			var s = t(this);
			o.push(new i(this.getContext("2d")).Bar(e, t.extend(s.data(), n)))
		}), 1 === o.length ? o[0] : o
	}
}.call(this, jQuery),
/*!
 * Datetimepicker for Bootstrap
 * Copyright 2012 Stefan Petre
 * Licensed under the Apache License v2.0
 */
!
function(t) {
	function e() {
		return new Date(Date.UTC.apply(Date, arguments))
	}
	var i = function(e, i) {
			var s = this;
			this.element = t(e), this.language = (i.language || this.element.data("date-language") || (t.zui && t.zui.clientLang ? t.zui.clientLang().replace("_", "-") : "zh-cn")).toLowerCase(), this.language = this.language in n ? this.language : "en", this.isRTL = n[this.language].rtl || !1, this.formatType = i.formatType || this.element.data("format-type") || "standard", this.format = o.parseFormat(i.format || this.element.data("date-format") || n[this.language].format || o.getDefaultFormat(this.formatType, "input"), this.formatType), this.isInline = !1, this.isVisible = !1, this.isInput = this.element.is("input"), this.component = !! this.element.is(".date") && this.element.find(".input-group-addon .icon-th, .input-group-addon .icon-time, .input-group-addon .icon-calendar").parent(), this.componentReset = !! this.element.is(".date") && this.element.find(".input-group-addon .icon-remove").parent(), this.hasInput = this.component && this.element.find("input").length, this.component && 0 === this.component.length && (this.component = !1), this.linkField = i.linkField || this.element.data("link-field") || !1, this.linkFormat = o.parseFormat(i.linkFormat || this.element.data("link-format") || o.getDefaultFormat(this.formatType, "link"), this.formatType), this.minuteStep = i.minuteStep || this.element.data("minute-step") || 5, this.pickerPosition = i.pickerPosition || this.element.data("picker-position") || "bottom-right", this.showMeridian = i.showMeridian || this.element.data("show-meridian") || !1, this.initialDate = i.initialDate || new Date, this.pickerClass = i.eleClass, this.pickerId = i.eleId, this._attachEvents(), this.formatViewType = "datetime", "formatViewType" in i ? this.formatViewType = i.formatViewType : "formatViewType" in this.element.data() && (this.formatViewType = this.element.data("formatViewType")), this.minView = 0, "minView" in i ? this.minView = i.minView : "minView" in this.element.data() && (this.minView = this.element.data("min-view")), this.minView = o.convertViewMode(this.minView), this.maxView = o.modes.length - 1, "maxView" in i ? this.maxView = i.maxView : "maxView" in this.element.data() && (this.maxView = this.element.data("max-view")), this.maxView = o.convertViewMode(this.maxView), this.wheelViewModeNavigation = !1, "wheelViewModeNavigation" in i ? this.wheelViewModeNavigation = i.wheelViewModeNavigation : "wheelViewModeNavigation" in this.element.data() && (this.wheelViewModeNavigation = this.element.data("view-mode-wheel-navigation")), this.wheelViewModeNavigationInverseDirection = !1, "wheelViewModeNavigationInverseDirection" in i ? this.wheelViewModeNavigationInverseDirection = i.wheelViewModeNavigationInverseDirection : "wheelViewModeNavigationInverseDirection" in this.element.data() && (this.wheelViewModeNavigationInverseDirection = this.element.data("view-mode-wheel-navigation-inverse-dir")), this.wheelViewModeNavigationDelay = 100, "wheelViewModeNavigationDelay" in i ? this.wheelViewModeNavigationDelay = i.wheelViewModeNavigationDelay : "wheelViewModeNavigationDelay" in this.element.data() && (this.wheelViewModeNavigationDelay = this.element.data("view-mode-wheel-navigation-delay")), this.startViewMode = 2, "startView" in i ? this.startViewMode = i.startView : "startView" in this.element.data() && (this.startViewMode = this.element.data("start-view")), this.startViewMode = o.convertViewMode(this.startViewMode), this.viewMode = this.startViewMode, this.viewSelect = this.minView, "viewSelect" in i ? this.viewSelect = i.viewSelect : "viewSelect" in this.element.data() && (this.viewSelect = this.element.data("view-select")), this.viewSelect = o.convertViewMode(this.viewSelect), this.forceParse = !0, "forceParse" in i ? this.forceParse = i.forceParse : "dateForceParse" in this.element.data() && (this.forceParse = this.element.data("date-force-parse")), this.picker = t(o.template).appendTo(this.isInline ? this.element : "body").on({
				click: t.proxy(this.click, this),
				mousedown: t.proxy(this.mousedown, this)
			}), this.wheelViewModeNavigation && (t.fn.mousewheel ? this.picker.on({
				mousewheel: t.proxy(this.mousewheel, this)
			}) : console.log("Mouse Wheel event is not supported. Please include the jQuery Mouse Wheel plugin before enabling this option")), this.isInline ? this.picker.addClass("datetimepicker-inline") : this.picker.addClass("datetimepicker-dropdown-" + this.pickerPosition + " dropdown-menu"), this.isRTL && (this.picker.addClass("datetimepicker-rtl"), this.picker.find(".prev span, .next span").toggleClass("icon-arrow-left icon-arrow-right")), t(document).on("mousedown", function(e) {
				0 === t(e.target).closest(".datetimepicker").length && s.hide()
			}), this.autoclose = !1, "autoclose" in i ? this.autoclose = i.autoclose : "dateAutoclose" in this.element.data() && (this.autoclose = this.element.data("date-autoclose")), this.keyboardNavigation = !0, "keyboardNavigation" in i ? this.keyboardNavigation = i.keyboardNavigation : "dateKeyboardNavigation" in this.element.data() && (this.keyboardNavigation = this.element.data("date-keyboard-navigation")), this.todayBtn = i.todayBtn || this.element.data("date-today-btn") || !1, this.todayHighlight = i.todayHighlight || this.element.data("date-today-highlight") || !1, this.weekStart = (i.weekStart || this.element.data("date-weekstart") || n[this.language].weekStart || 0) % 7, this.weekEnd = (this.weekStart + 6) % 7, this.startDate = -(1 / 0), this.endDate = 1 / 0, this.daysOfWeekDisabled = [], this.setStartDate(i.startDate || this.element.data("date-startdate")), this.setEndDate(i.endDate || this.element.data("date-enddate")), this.setDaysOfWeekDisabled(i.daysOfWeekDisabled || this.element.data("date-days-of-week-disabled")), this.fillDow(), this.fillMonths(), this.update(), this.showMode(), this.isInline && this.show()
		};
	i.prototype = {
		constructor: i,
		_events: [],
		_attachEvents: function() {
			this._detachEvents(), this.isInput ? this._events = [
				[this.element,
				{
					focus: t.proxy(this.show, this),
					keyup: t.proxy(this.update, this),
					keydown: t.proxy(this.keydown, this)
				}]
			] : this.component && this.hasInput ? (this._events = [
				[this.element.find("input"),
				{
					focus: t.proxy(this.show, this),
					keyup: t.proxy(this.update, this),
					keydown: t.proxy(this.keydown, this)
				}],
				[this.component,
				{
					click: t.proxy(this.show, this)
				}]
			], this.componentReset && this._events.push([this.componentReset,
			{
				click: t.proxy(this.reset, this)
			}])) : this.element.is("div") ? this.isInline = !0 : this._events = [
				[this.element,
				{
					click: t.proxy(this.show, this)
				}]
			];
			for (var e, i, n = 0; n < this._events.length; n++) e = this._events[n][0], i = this._events[n][1], e.on(i)
		},
		_detachEvents: function() {
			for (var t, e, i = 0; i < this._events.length; i++) t = this._events[i][0], e = this._events[i][1], t.off(e);
			this._events = []
		},
		show: function(e) {
			this.picker.show(), this.height = this.component ? this.component.outerHeight() : this.element.outerHeight(), this.forceParse && this.update(), this.place(), t(window).on("resize", t.proxy(this.place, this)), e && (e.stopPropagation(), e.preventDefault()), this.isVisible = !0, this.element.trigger({
				type: "show",
				date: this.date
			})
		},
		hide: function(e) {
			this.isVisible && (this.isInline || (this.picker.hide(), t(window).off("resize", this.place), this.viewMode = this.startViewMode, this.showMode(), this.isInput || t(document).off("mousedown", this.hide), this.forceParse && (this.isInput && this.element.val() || this.hasInput && this.element.find("input").val()) && this.setValue(), this.isVisible = !1, this.element.trigger({
				type: "hide",
				date: this.date
			})))
		},
		remove: function() {
			this._detachEvents(), this.picker.remove(), delete this.picker, delete this.element.data().datetimepicker
		},
		getDate: function() {
			var t = this.getUTCDate();
			return new Date(t.getTime() + 6e4 * t.getTimezoneOffset())
		},
		getUTCDate: function() {
			return this.date
		},
		setDate: function(t) {
			this.setUTCDate(new Date(t.getTime() - 6e4 * t.getTimezoneOffset()))
		},
		setUTCDate: function(t) {
			t >= this.startDate && t <= this.endDate ? (this.date = t, this.setValue(), this.viewDate = this.date, this.fill()) : this.element.trigger({
				type: "outOfRange",
				date: t,
				startDate: this.startDate,
				endDate: this.endDate
			})
		},
		setFormat: function(t) {
			this.format = o.parseFormat(t, this.formatType);
			var e;
			this.isInput ? e = this.element : this.component && (e = this.element.find("input")), e && e.val() && this.setValue()
		},
		setValue: function() {
			var e = this.getFormattedDate();
			this.isInput ? this.element.val(e) : (this.component && this.element.find("input").val(e), this.element.data("date", e)), this.linkField && t("#" + this.linkField).val(this.getFormattedDate(this.linkFormat))
		},
		getFormattedDate: function(t) {
			return void 0 == t && (t = this.format), o.formatDate(this.date, t, this.language, this.formatType)
		},
		setStartDate: function(t) {
			this.startDate = t || -(1 / 0), this.startDate !== -(1 / 0) && (this.startDate = o.parseDate(this.startDate, this.format, this.language, this.formatType)), this.update(), this.updateNavArrows()
		},
		setEndDate: function(t) {
			this.endDate = t || 1 / 0, this.endDate !== 1 / 0 && (this.endDate = o.parseDate(this.endDate, this.format, this.language, this.formatType)), this.update(), this.updateNavArrows()
		},
		setDaysOfWeekDisabled: function(e) {
			this.daysOfWeekDisabled = e || [], t.isArray(this.daysOfWeekDisabled) || (this.daysOfWeekDisabled = this.daysOfWeekDisabled.split(/,\s*/)), this.daysOfWeekDisabled = t.map(this.daysOfWeekDisabled, function(t) {
				return parseInt(t, 10)
			}), this.update(), this.updateNavArrows()
		},
		place: function() {
			if (!this.isInline) {
				var e = 0;
				t("div").each(function() {
					var i = parseInt(t(this).css("zIndex"), 10);
					i > e && (e = i)
				});
				var i, n, o, s = e + 10;
				this.component ? (i = this.component.offset(), o = i.left, "bottom-left" !== this.pickerPosition && "top-left" !== this.pickerPosition && "auto-left" !== this.pickerPosition || (o += this.component.outerWidth() - this.picker.outerWidth())) : (i = this.element.offset(), o = i.left);
				var a = 0 === this.pickerPosition.indexOf("auto-"),
					r = a ? (i.top + this.picker.outerHeight() > t(window).height() + t(window).scrollTop() ? "top" : "bottom") + (0 === this.pickerPosition.lastIndexOf("-left") ? "-left" : "-right") : this.pickerPosition;
				n = "top-left" === r || "top-right" === r ? i.top - this.picker.outerHeight() : i.top + this.height, this.picker.css({
					top: n,
					left: o,
					zIndex: s
				}).attr("class", "datetimepicker dropdown-menu datetimepicker-dropdown-" + r), this.pickerClass && this.picker.addClass(this.pickerClass), this.pickerId && this.picker.attr("id", this.pickerId)
			}
		},
		update: function() {
			var t, e = !1;
			arguments && arguments.length && ("string" == typeof arguments[0] || arguments[0] instanceof Date) ? (t = arguments[0], e = !0) : (t = this.element.data("date") || (this.isInput ? this.element.val() : this.element.find("input").val()) || this.initialDate, ("string" == typeof t || t instanceof String) && (t = t.replace(/^\s+|\s+$/g, ""))), t || (t = new Date, e = !1), this.date = o.parseDate(t, this.format, this.language, this.formatType), e && this.setValue(), this.date < this.startDate ? this.viewDate = new Date(this.startDate) : this.date > this.endDate ? this.viewDate = new Date(this.endDate) : this.viewDate = new Date(this.date), this.fill()
		},
		fillDow: function() {
			for (var t = this.weekStart, e = "<tr>"; t < this.weekStart + 7;) e += '<th class="dow">' + n[this.language].daysMin[t++ % 7] + "</th>";
			e += "</tr>", this.picker.find(".datetimepicker-days thead").append(e)
		},
		fillMonths: function() {
			for (var t = "", e = 0; e < 12;) t += '<span class="month">' + n[this.language].monthsShort[e++] + "</span>";
			this.picker.find(".datetimepicker-months td").html(t)
		},
		fill: function() {
			if (null != this.date && null != this.viewDate) {
				var i = new Date(this.viewDate),
					s = i.getUTCFullYear(),
					a = i.getUTCMonth(),
					r = i.getUTCDate(),
					l = i.getUTCHours(),
					h = i.getUTCMinutes(),
					c = this.startDate !== -(1 / 0) ? this.startDate.getUTCFullYear() : -(1 / 0),
					d = this.startDate !== -(1 / 0) ? this.startDate.getUTCMonth() : -(1 / 0),
					u = this.endDate !== 1 / 0 ? this.endDate.getUTCFullYear() : 1 / 0,
					p = this.endDate !== 1 / 0 ? this.endDate.getUTCMonth() : 1 / 0,
					f = new e(this.date.getUTCFullYear(), this.date.getUTCMonth(), this.date.getUTCDate()).valueOf(),
					g = new Date;
				if (this.picker.find(".datetimepicker-days thead th:eq(1)").text(n[this.language].months[a] + " " + s), "time" == this.formatViewType) {
					var m = l % 12 ? l % 12 : 12,
						v = (m < 10 ? "0" : "") + m,
						y = (h < 10 ? "0" : "") + h,
						b = n[this.language].meridiem[l < 12 ? 0 : 1];
					this.picker.find(".datetimepicker-hours thead th:eq(1)").text(v + ":" + y + " " + b.toUpperCase()), this.picker.find(".datetimepicker-minutes thead th:eq(1)").text(v + ":" + y + " " + b.toUpperCase())
				} else this.picker.find(".datetimepicker-hours thead th:eq(1)").text(r + " " + n[this.language].months[a] + " " + s), this.picker.find(".datetimepicker-minutes thead th:eq(1)").text(r + " " + n[this.language].months[a] + " " + s);
				this.picker.find("tfoot th.today").text(n[this.language].today).toggle(this.todayBtn !== !1), this.updateNavArrows(), this.fillMonths();
				var w = e(s, a - 1, 28, 0, 0, 0, 0),
					x = o.getDaysInMonth(w.getUTCFullYear(), w.getUTCMonth());
				w.setUTCDate(x), w.setUTCDate(x - (w.getUTCDay() - this.weekStart + 7) % 7);
				var C = new Date(w);
				C.setUTCDate(C.getUTCDate() + 42), C = C.valueOf();
				for (var _, k = []; w.valueOf() < C;) w.getUTCDay() == this.weekStart && k.push("<tr>"), _ = "", w.getUTCFullYear() < s || w.getUTCFullYear() == s && w.getUTCMonth() < a ? _ += " old" : (w.getUTCFullYear() > s || w.getUTCFullYear() == s && w.getUTCMonth() > a) && (_ += " new"), this.todayHighlight && w.getUTCFullYear() == g.getFullYear() && w.getUTCMonth() == g.getMonth() && w.getUTCDate() == g.getDate() && (_ += " today"), w.valueOf() == f && (_ += " active"), (w.valueOf() + 864e5 <= this.startDate || w.valueOf() > this.endDate || t.inArray(w.getUTCDay(), this.daysOfWeekDisabled) !== -1) && (_ += " disabled"), k.push('<td class="day' + _ + '">' + w.getUTCDate() + "</td>"), w.getUTCDay() == this.weekEnd && k.push("</tr>"), w.setUTCDate(w.getUTCDate() + 1);
				this.picker.find(".datetimepicker-days tbody").empty().append(k.join("")), k = [];
				for (var T = "", S = "", D = "", M = 0; M < 24; M++) {
					var P = e(s, a, r, M);
					_ = "", P.valueOf() + 36e5 <= this.startDate || P.valueOf() > this.endDate ? _ += " disabled" : l == M && (_ += " active"), this.showMeridian && 2 == n[this.language].meridiem.length ? (S = M < 12 ? n[this.language].meridiem[0] : n[this.language].meridiem[1], S != D && ("" != D && k.push("</fieldset>"), k.push('<fieldset class="hour"><legend>' + S.toUpperCase() + "</legend>")), D = S, T = M % 12 ? M % 12 : 12, k.push('<span class="hour' + _ + " hour_" + (M < 12 ? "am" : "pm") + '">' + T + "</span>"), 23 == M && k.push("</fieldset>")) : (T = M + ":00", k.push('<span class="hour' + _ + '">' + T + "</span>"))
				}
				this.picker.find(".datetimepicker-hours td").html(k.join("")), k = [], T = "", S = "", D = "";
				for (var M = 0; M < 60; M += this.minuteStep) {
					var P = e(s, a, r, l, M, 0);
					_ = "", P.valueOf() < this.startDate || P.valueOf() > this.endDate ? _ += " disabled" : Math.floor(h / this.minuteStep) == Math.floor(M / this.minuteStep) && (_ += " active"), this.showMeridian && 2 == n[this.language].meridiem.length ? (S = l < 12 ? n[this.language].meridiem[0] : n[this.language].meridiem[1], S != D && ("" != D && k.push("</fieldset>"), k.push('<fieldset class="minute"><legend>' + S.toUpperCase() + "</legend>")), D = S, T = l % 12 ? l % 12 : 12, k.push('<span class="minute' + _ + '">' + T + ":" + (M < 10 ? "0" + M : M) + "</span>"), 59 == M && k.push("</fieldset>")) : (T = M + ":00", k.push('<span class="minute' + _ + '">' + l + ":" + (M < 10 ? "0" + M : M) + "</span>"))
				}
				this.picker.find(".datetimepicker-minutes td").html(k.join(""));
				var F = this.date.getUTCFullYear(),
					L = this.picker.find(".datetimepicker-months").find("th:eq(1)").text(s).end().find("span").removeClass("active");
				F == s && L.eq(this.date.getUTCMonth()).addClass("active"), (s < c || s > u) && L.addClass("disabled"), s == c && L.slice(0, d).addClass("disabled"), s == u && L.slice(p + 1).addClass("disabled"), k = "", s = 10 * parseInt(s / 10, 10);
				var z = this.picker.find(".datetimepicker-years").find("th:eq(1)").text(s + "-" + (s + 9)).end().find("td");
				s -= 1;
				for (var M = -1; M < 11; M++) k += '<span class="year' + (M == -1 || 10 == M ? " old" : "") + (F == s ? " active" : "") + (s < c || s > u ? " disabled" : "") + '">' + s + "</span>", s += 1;
				z.html(k), this.place()
			}
		},
		updateNavArrows: function() {
			var t = new Date(this.viewDate),
				e = t.getUTCFullYear(),
				i = t.getUTCMonth(),
				n = t.getUTCDate(),
				o = t.getUTCHours();
			switch (this.viewMode) {
			case 0:
				this.startDate !== -(1 / 0) && e <= this.startDate.getUTCFullYear() && i <= this.startDate.getUTCMonth() && n <= this.startDate.getUTCDate() && o <= this.startDate.getUTCHours() ? this.picker.find(".prev").css({
					visibility: "hidden"
				}) : this.picker.find(".prev").css({
					visibility: "visible"
				}), this.endDate !== 1 / 0 && e >= this.endDate.getUTCFullYear() && i >= this.endDate.getUTCMonth() && n >= this.endDate.getUTCDate() && o >= this.endDate.getUTCHours() ? this.picker.find(".next").css({
					visibility: "hidden"
				}) : this.picker.find(".next").css({
					visibility: "visible"
				});
				break;
			case 1:
				this.startDate !== -(1 / 0) && e <= this.startDate.getUTCFullYear() && i <= this.startDate.getUTCMonth() && n <= this.startDate.getUTCDate() ? this.picker.find(".prev").css({
					visibility: "hidden"
				}) : this.picker.find(".prev").css({
					visibility: "visible"
				}), this.endDate !== 1 / 0 && e >= this.endDate.getUTCFullYear() && i >= this.endDate.getUTCMonth() && n >= this.endDate.getUTCDate() ? this.picker.find(".next").css({
					visibility: "hidden"
				}) : this.picker.find(".next").css({
					visibility: "visible"
				});
				break;
			case 2:
				this.startDate !== -(1 / 0) && e <= this.startDate.getUTCFullYear() && i <= this.startDate.getUTCMonth() ? this.picker.find(".prev").css({
					visibility: "hidden"
				}) : this.picker.find(".prev").css({
					visibility: "visible"
				}), this.endDate !== 1 / 0 && e >= this.endDate.getUTCFullYear() && i >= this.endDate.getUTCMonth() ? this.picker.find(".next").css({
					visibility: "hidden"
				}) : this.picker.find(".next").css({
					visibility: "visible"
				});
				break;
			case 3:
			case 4:
				this.startDate !== -(1 / 0) && e <= this.startDate.getUTCFullYear() ? this.picker.find(".prev").css({
					visibility: "hidden"
				}) : this.picker.find(".prev").css({
					visibility: "visible"
				}), this.endDate !== 1 / 0 && e >= this.endDate.getUTCFullYear() ? this.picker.find(".next").css({
					visibility: "hidden"
				}) : this.picker.find(".next").css({
					visibility: "visible"
				})
			}
		},
		mousewheel: function(e) {
			if (e.preventDefault(), e.stopPropagation(), !this.wheelPause) {
				this.wheelPause = !0;
				var i = e.originalEvent,
					n = i.wheelDelta,
					o = n > 0 ? 1 : 0 === n ? 0 : -1;
				this.wheelViewModeNavigationInverseDirection && (o = -o), this.showMode(o), setTimeout(t.proxy(function() {
					this.wheelPause = !1
				}, this), this.wheelViewModeNavigationDelay)
			}
		},
		click: function(i) {
			i.stopPropagation(), i.preventDefault();
			var n = t(i.target).closest("span, td, th, legend");
			if (1 == n.length) {
				if (n.is(".disabled")) return void this.element.trigger({
					type: "outOfRange",
					date: this.viewDate,
					startDate: this.startDate,
					endDate: this.endDate
				});
				switch (n[0].nodeName.toLowerCase()) {
				case "th":
					switch (n[0].className) {
					case "switch":
						this.showMode(1);
						break;
					case "prev":
					case "next":
						var s = o.modes[this.viewMode].navStep * ("prev" == n[0].className ? -1 : 1);
						switch (this.viewMode) {
						case 0:
							this.viewDate = this.moveHour(this.viewDate, s);
							break;
						case 1:
							this.viewDate = this.moveDate(this.viewDate, s);
							break;
						case 2:
							this.viewDate = this.moveMonth(this.viewDate, s);
							break;
						case 3:
						case 4:
							this.viewDate = this.moveYear(this.viewDate, s)
						}
						this.fill();
						break;
					case "today":
						var a = new Date;
						a = e(a.getFullYear(), a.getMonth(), a.getDate(), a.getHours(), a.getMinutes(), a.getSeconds(), 0), a < this.startDate ? a = this.startDate : a > this.endDate && (a = this.endDate), this.viewMode = this.startViewMode, this.showMode(0), this._setDate(a), this.fill(), this.autoclose && this.hide()
					}
					break;
				case "span":
					if (!n.is(".disabled")) {
						var r = this.viewDate.getUTCFullYear(),
							l = this.viewDate.getUTCMonth(),
							h = this.viewDate.getUTCDate(),
							c = this.viewDate.getUTCHours(),
							d = this.viewDate.getUTCMinutes(),
							u = this.viewDate.getUTCSeconds();
						if (n.is(".month") ? (this.viewDate.setUTCDate(1), l = n.parent().find("span").index(n), h = this.viewDate.getUTCDate(), this.viewDate.setUTCMonth(l), this.element.trigger({
							type: "changeMonth",
							date: this.viewDate
						}), this.viewSelect >= 3 && this._setDate(e(r, l, h, c, d, u, 0))) : n.is(".year") ? (this.viewDate.setUTCDate(1), r = parseInt(n.text(), 10) || 0, this.viewDate.setUTCFullYear(r), this.element.trigger({
							type: "changeYear",
							date: this.viewDate
						}), this.viewSelect >= 4 && this._setDate(e(r, l, h, c, d, u, 0))) : n.is(".hour") ? (c = parseInt(n.text(), 10) || 0, (n.hasClass("hour_am") || n.hasClass("hour_pm")) && (12 == c && n.hasClass("hour_am") ? c = 0 : 12 != c && n.hasClass("hour_pm") && (c += 12)), this.viewDate.setUTCHours(c), this.element.trigger({
							type: "changeHour",
							date: this.viewDate
						}), this.viewSelect >= 1 && this._setDate(e(r, l, h, c, d, u, 0))) : n.is(".minute") && (d = parseInt(n.text().substr(n.text().indexOf(":") + 1), 10) || 0, this.viewDate.setUTCMinutes(d), this.element.trigger({
							type: "changeMinute",
							date: this.viewDate
						}), this.viewSelect >= 0 && this._setDate(e(r, l, h, c, d, u, 0))), 0 != this.viewMode) {
							var p = this.viewMode;
							this.showMode(-1), this.fill(), p == this.viewMode && this.autoclose && this.hide()
						} else this.fill(), this.autoclose && this.hide()
					}
					break;
				case "td":
					if (n.is(".day") && !n.is(".disabled")) {
						var h = parseInt(n.text(), 10) || 1,
							r = this.viewDate.getUTCFullYear(),
							l = this.viewDate.getUTCMonth(),
							c = this.viewDate.getUTCHours(),
							d = this.viewDate.getUTCMinutes(),
							u = this.viewDate.getUTCSeconds();
						n.is(".old") ? 0 === l ? (l = 11, r -= 1) : l -= 1 : n.is(".new") && (11 == l ? (l = 0, r += 1) : l += 1), this.viewDate.setUTCFullYear(r), this.viewDate.setUTCMonth(l, h), this.element.trigger({
							type: "changeDay",
							date: this.viewDate
						}), this.viewSelect >= 2 && this._setDate(e(r, l, h, c, d, u, 0))
					}
					var p = this.viewMode;
					this.showMode(-1), this.fill(), p == this.viewMode && this.autoclose && this.hide()
				}
			}
		},
		_setDate: function(t, e) {
			e && "date" != e || (this.date = t), e && "view" != e || (this.viewDate = t), this.fill(), this.setValue();
			var i;
			this.isInput ? i = this.element : this.component && (i = this.element.find("input")), i && (i.change(), this.autoclose && (!e || "date" == e)), this.element.trigger({
				type: "changeDate",
				date: this.date
			}), null === t && (this.date = this.viewDate)
		},
		moveMinute: function(t, e) {
			if (!e) return t;
			var i = new Date(t.valueOf());
			return i.setUTCMinutes(i.getUTCMinutes() + e * this.minuteStep), i
		},
		moveHour: function(t, e) {
			if (!e) return t;
			var i = new Date(t.valueOf());
			return i.setUTCHours(i.getUTCHours() + e), i
		},
		moveDate: function(t, e) {
			if (!e) return t;
			var i = new Date(t.valueOf());
			return i.setUTCDate(i.getUTCDate() + e), i
		},
		moveMonth: function(t, e) {
			if (!e) return t;
			var i, n, o = new Date(t.valueOf()),
				s = o.getUTCDate(),
				a = o.getUTCMonth(),
				r = Math.abs(e);
			if (e = e > 0 ? 1 : -1, 1 == r) n = e == -1 ?
			function() {
				return o.getUTCMonth() == a
			} : function() {
				return o.getUTCMonth() != i
			}, i = a + e, o.setUTCMonth(i), (i < 0 || i > 11) && (i = (i + 12) % 12);
			else {
				for (var l = 0; l < r; l++) o = this.moveMonth(o, e);
				i = o.getUTCMonth(), o.setUTCDate(s), n = function() {
					return i != o.getUTCMonth()
				}
			}
			for (; n();) o.setUTCDate(--s), o.setUTCMonth(i);
			return o
		},
		moveYear: function(t, e) {
			return this.moveMonth(t, 12 * e)
		},
		dateWithinRange: function(t) {
			return t >= this.startDate && t <= this.endDate
		},
		keydown: function(t) {
			if (this.picker.is(":not(:visible)")) return void(27 == t.keyCode && this.show());
			var e, i, n, o = !1;
			switch (t.keyCode) {
			case 27:
				this.hide(), t.preventDefault();
				break;
			case 37:
			case 39:
				if (!this.keyboardNavigation) break;
				e = 37 == t.keyCode ? -1 : 1, viewMode = this.viewMode, t.ctrlKey ? viewMode += 2 : t.shiftKey && (viewMode += 1), 4 == viewMode ? (i = this.moveYear(this.date, e), n = this.moveYear(this.viewDate, e)) : 3 == viewMode ? (i = this.moveMonth(this.date, e), n = this.moveMonth(this.viewDate, e)) : 2 == viewMode ? (i = this.moveDate(this.date, e), n = this.moveDate(this.viewDate, e)) : 1 == viewMode ? (i = this.moveHour(this.date, e), n = this.moveHour(this.viewDate, e)) : 0 == viewMode && (i = this.moveMinute(this.date, e), n = this.moveMinute(this.viewDate, e)), this.dateWithinRange(i) && (this.date = i, this.viewDate = n, this.setValue(), this.update(), t.preventDefault(), o = !0);
				break;
			case 38:
			case 40:
				if (!this.keyboardNavigation) break;
				e = 38 == t.keyCode ? -1 : 1, viewMode = this.viewMode, t.ctrlKey ? viewMode += 2 : t.shiftKey && (viewMode += 1), 4 == viewMode ? (i = this.moveYear(this.date, e), n = this.moveYear(this.viewDate, e)) : 3 == viewMode ? (i = this.moveMonth(this.date, e), n = this.moveMonth(this.viewDate, e)) : 2 == viewMode ? (i = this.moveDate(this.date, 7 * e), n = this.moveDate(this.viewDate, 7 * e)) : 1 == viewMode ? this.showMeridian ? (i = this.moveHour(this.date, 6 * e), n = this.moveHour(this.viewDate, 6 * e)) : (i = this.moveHour(this.date, 4 * e), n = this.moveHour(this.viewDate, 4 * e)) : 0 == viewMode && (i = this.moveMinute(this.date, 4 * e), n = this.moveMinute(this.viewDate, 4 * e)), this.dateWithinRange(i) && (this.date = i, this.viewDate = n, this.setValue(), this.update(), t.preventDefault(), o = !0);
				break;
			case 13:
				if (0 != this.viewMode) {
					var s = this.viewMode;
					this.showMode(-1), this.fill(), s == this.viewMode && this.autoclose && this.hide()
				} else this.fill(), this.autoclose && this.hide();
				t.preventDefault();
				break;
			case 9:
				this.hide()
			}
			if (o) {
				var a;
				this.isInput ? a = this.element : this.component && (a = this.element.find("input")), a && a.change(), this.element.trigger({
					type: "changeDate",
					date: this.date
				})
			}
		},
		showMode: function(t) {
			if (t) {
				var e = Math.max(0, Math.min(o.modes.length - 1, this.viewMode + t));
				e >= this.minView && e <= this.maxView && (this.element.trigger({
					type: "changeMode",
					date: this.viewDate,
					oldViewMode: this.viewMode,
					newViewMode: e
				}), this.viewMode = e)
			}
			this.picker.find(">div").hide().filter(".datetimepicker-" + o.modes[this.viewMode].clsName).css("display", "block"), this.updateNavArrows()
		},
		reset: function(t) {
			this._setDate(null, "date")
		}
	}, t.fn.datetimepicker = function(e) {
		var n = Array.apply(null, arguments);
		return n.shift(), this.each(function() {
			var o = t(this),
				s = o.data("datetimepicker"),
				a = "object" == typeof e && e;
			s || o.data("datetimepicker", s = new i(this, t.extend({}, t.fn.datetimepicker.defaults, o.data(), a))), "string" == typeof e && "function" == typeof s[e] && s[e].apply(s, n)
		})
	}, t.fn.datetimepicker.defaults = {
		pickerPosition: "auto-right"
	}, t.fn.datetimepicker.Constructor = i;
	var n = t.fn.datetimepicker.dates = {
		en: {
			days: ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"],
			daysShort: ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"],
			daysMin: ["Su", "Mo", "Tu", "We", "Th", "Fr", "Sa", "Su"],
			months: ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"],
			monthsShort: ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"],
			meridiem: ["am", "pm"],
			suffix: ["st", "nd", "rd", "th"],
			today: "Today"
		}
	};
	n["zh-cn"] = {
		days: ["星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"],
		daysShort: ["周日", "周一", "周二", "周三", "周四", "周五", "周六", "周日"],
		daysMin: ["日", "一", "二", "三", "四", "五", "六", "日"],
		months: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
		monthsShort: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
		today: "今日",
		suffix: [],
		meridiem: []
	}, n["zh-tw"] = {
		days: ["星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"],
		daysShort: ["周日", "周一", "周二", "周三", "周四", "周五", "周六", "周日"],
		daysMin: ["日", "一", "二", "三", "四", "五", "六", "日"],
		months: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
		monthsShort: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
		today: "今天",
		suffix: [],
		meridiem: ["上午", "下午"]
	};
	var o = {
		modes: [{
			clsName: "minutes",
			navFnc: "Hours",
			navStep: 1
		}, {
			clsName: "hours",
			navFnc: "Date",
			navStep: 1
		}, {
			clsName: "days",
			navFnc: "Month",
			navStep: 1
		}, {
			clsName: "months",
			navFnc: "FullYear",
			navStep: 1
		}, {
			clsName: "years",
			navFnc: "FullYear",
			navStep: 10
		}],
		isLeapYear: function(t) {
			return t % 4 === 0 && t % 100 !== 0 || t % 400 === 0
		},
		getDaysInMonth: function(t, e) {
			return [31, o.isLeapYear(t) ? 29 : 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31][e]
		},
		getDefaultFormat: function(t, e) {
			if ("standard" == t) return "input" == e ? "yyyy-mm-dd hh:ii" : "yyyy-mm-dd hh:ii:ss";
			if ("php" == t) return "input" == e ? "Y-m-d H:i" : "Y-m-d H:i:s";
			throw new Error("Invalid format type.")
		},
		validParts: function(t) {
			if ("standard" == t) return /hh?|HH?|p|P|ii?|ss?|dd?|DD?|mm?|MM?|yy(?:yy)?/g;
			if ("php" == t) return /[dDjlNwzFmMnStyYaABgGhHis]/g;
			throw new Error("Invalid format type.")
		},
		nonpunctuation: /[^ -\/:-@\[-`{-~\t\n\rTZ]+/g,
		parseFormat: function(t, e) {
			var i = t.replace(this.validParts(e), "\0").split("\0"),
				n = t.match(this.validParts(e));
			if (!i || !i.length || !n || 0 == n.length) throw new Error("Invalid date format.");
			return {
				separators: i,
				parts: n
			}
		},
		parseDate: function(o, s, a, r) {
			if (o instanceof Date) {
				var l = new Date(o.valueOf() - 6e4 * o.getTimezoneOffset());
				return l.setMilliseconds(0), l
			}
			if (/^\d{4}\-\d{1,2}\-\d{1,2}$/.test(o) && (s = this.parseFormat("yyyy-mm-dd", r)), /^\d{4}\-\d{1,2}\-\d{1,2}[T ]\d{1,2}\:\d{1,2}$/.test(o) && (s = this.parseFormat("yyyy-mm-dd hh:ii", r)), /^\d{4}\-\d{1,2}\-\d{1,2}[T ]\d{1,2}\:\d{1,2}\:\d{1,2}[Z]{0,1}$/.test(o) && (s = this.parseFormat("yyyy-mm-dd hh:ii:ss", r)), /^[-+]\d+[dmwy]([\s,]+[-+]\d+[dmwy])*$/.test(o)) {
				var h, c, d = /([-+]\d+)([dmwy])/,
					u = o.match(/([-+]\d+)([dmwy])/g);
				o = new Date;
				for (var p = 0; p < u.length; p++) switch (h = d.exec(u[p]), c = parseInt(h[1]), h[2]) {
				case "d":
					o.setUTCDate(o.getUTCDate() + c);
					break;
				case "m":
					o = i.prototype.moveMonth.call(i.prototype, o, c);
					break;
				case "w":
					o.setUTCDate(o.getUTCDate() + 7 * c);
					break;
				case "y":
					o = i.prototype.moveYear.call(i.prototype, o, c)
				}
				return e(o.getUTCFullYear(), o.getUTCMonth(), o.getUTCDate(), o.getUTCHours(), o.getUTCMinutes(), o.getUTCSeconds(), 0)
			}
			var f, g, h, u = o && o.match(this.nonpunctuation) || [],
				o = new Date(0, 0, 0, 0, 0, 0, 0),
				m = {},
				v = ["hh", "h", "ii", "i", "ss", "s", "yyyy", "yy", "M", "MM", "m", "mm", "D", "DD", "d", "dd", "H", "HH", "p", "P"],
				y = {
					hh: function(t, e) {
						return t.setUTCHours(e)
					},
					h: function(t, e) {
						return t.setUTCHours(e)
					},
					HH: function(t, e) {
						return t.setUTCHours(12 == e ? 0 : e)
					},
					H: function(t, e) {
						return t.setUTCHours(12 == e ? 0 : e)
					},
					ii: function(t, e) {
						return t.setUTCMinutes(e)
					},
					i: function(t, e) {
						return t.setUTCMinutes(e)
					},
					ss: function(t, e) {
						return t.setUTCSeconds(e)
					},
					s: function(t, e) {
						return t.setUTCSeconds(e)
					},
					yyyy: function(t, e) {
						return t.setUTCFullYear(e)
					},
					yy: function(t, e) {
						return t.setUTCFullYear(2e3 + e)
					},
					m: function(t, e) {
						for (e -= 1; e < 0;) e += 12;
						for (e %= 12, t.setUTCMonth(e); t.getUTCMonth() != e;) t.setUTCDate(t.getUTCDate() - 1);
						return t
					},
					d: function(t, e) {
						return t.setUTCDate(e)
					},
					p: function(t, e) {
						return t.setUTCHours(1 == e ? t.getUTCHours() + 12 : t.getUTCHours())
					}
				};
			if (y.M = y.MM = y.mm = y.m, y.dd = y.d, y.P = y.p, o = e(o.getFullYear(), o.getMonth(), o.getDate(), o.getHours(), o.getMinutes(), o.getSeconds()), u.length == s.parts.length) {
				for (var p = 0, b = s.parts.length; p < b; p++) {
					if (f = parseInt(u[p], 10), h = s.parts[p], isNaN(f)) switch (h) {
					case "MM":
						g = t(n[a].months).filter(function() {
							var t = this.slice(0, u[p].length),
								e = u[p].slice(0, t.length);
							return t == e
						}), f = t.inArray(g[0], n[a].months) + 1;
						break;
					case "M":
						g = t(n[a].monthsShort).filter(function() {
							var t = this.slice(0, u[p].length),
								e = u[p].slice(0, t.length);
							return t == e
						}), f = t.inArray(g[0], n[a].monthsShort) + 1;
						break;
					case "p":
					case "P":
						f = t.inArray(u[p].toLowerCase(), n[a].meridiem)
					}
					m[h] = f
				}
				for (var w, p = 0; p < v.length; p++) w = v[p], w in m && !isNaN(m[w]) && y[w](o, m[w])
			}
			return o
		},
		formatDate: function(e, i, s, a) {
			if (null == e) return "";
			var r;
			if ("standard" == a) r = {
				yy: e.getUTCFullYear().toString().substring(2),
				yyyy: e.getUTCFullYear(),
				m: e.getUTCMonth() + 1,
				M: n[s].monthsShort[e.getUTCMonth()],
				MM: n[s].months[e.getUTCMonth()],
				d: e.getUTCDate(),
				D: n[s].daysShort[e.getUTCDay()],
				DD: n[s].days[e.getUTCDay()],
				p: 2 == n[s].meridiem.length ? n[s].meridiem[e.getUTCHours() < 12 ? 0 : 1] : "",
				h: e.getUTCHours(),
				i: e.getUTCMinutes(),
				s: e.getUTCSeconds()
			}, 2 == n[s].meridiem.length ? r.H = r.h % 12 == 0 ? 12 : r.h % 12 : r.H = r.h, r.HH = (r.H < 10 ? "0" : "") + r.H, r.P = r.p.toUpperCase(), r.hh = (r.h < 10 ? "0" : "") + r.h, r.ii = (r.i < 10 ? "0" : "") + r.i, r.ss = (r.s < 10 ? "0" : "") + r.s, r.dd = (r.d < 10 ? "0" : "") + r.d, r.mm = (r.m < 10 ? "0" : "") + r.m;
			else {
				if ("php" != a) throw new Error("Invalid format type.");
				r = {
					y: e.getUTCFullYear().toString().substring(2),
					Y: e.getUTCFullYear(),
					F: n[s].months[e.getUTCMonth()],
					M: n[s].monthsShort[e.getUTCMonth()],
					n: e.getUTCMonth() + 1,
					t: o.getDaysInMonth(e.getUTCFullYear(), e.getUTCMonth()),
					j: e.getUTCDate(),
					l: n[s].days[e.getUTCDay()],
					D: n[s].daysShort[e.getUTCDay()],
					w: e.getUTCDay(),
					N: 0 == e.getUTCDay() ? 7 : e.getUTCDay(),
					S: e.getUTCDate() % 10 <= n[s].suffix.length ? n[s].suffix[e.getUTCDate() % 10 - 1] : "",
					a: 2 == n[s].meridiem.length ? n[s].meridiem[e.getUTCHours() < 12 ? 0 : 1] : "",
					g: e.getUTCHours() % 12 == 0 ? 12 : e.getUTCHours() % 12,
					G: e.getUTCHours(),
					i: e.getUTCMinutes(),
					s: e.getUTCSeconds()
				}, r.m = (r.n < 10 ? "0" : "") + r.n, r.d = (r.j < 10 ? "0" : "") + r.j, r.A = r.a.toString().toUpperCase(), r.h = (r.g < 10 ? "0" : "") + r.g, r.H = (r.G < 10 ? "0" : "") + r.G, r.i = (r.i < 10 ? "0" : "") + r.i, r.s = (r.s < 10 ? "0" : "") + r.s
			}
			for (var e = [], l = t.extend([], i.separators), h = 0, c = i.parts.length; h < c; h++) l.length && e.push(l.shift()), e.push(r[i.parts[h]]);
			return l.length && e.push(l.shift()), e.join("")
		},
		convertViewMode: function(t) {
			switch (t) {
			case 4:
			case "decade":
				t = 4;
				break;
			case 3:
			case "year":
				t = 3;
				break;
			case 2:
			case "month":
				t = 2;
				break;
			case 1:
			case "day":
				t = 1;
				break;
			case 0:
			case "hour":
				t = 0
			}
			return t
		},
		headTemplate: '<thead><tr><th class="prev"><i class="icon-arrow-left"/></th><th colspan="5" class="switch"></th><th class="next"><i class="icon-arrow-right"/></th></tr></thead>',
		contTemplate: '<tbody><tr><td colspan="7"></td></tr></tbody>',
		footTemplate: '<tfoot><tr><th colspan="7" class="today"></th></tr></tfoot>'
	};
	o.template = '<div class="datetimepicker"><div class="datetimepicker-minutes"><table class=" table-condensed">' + o.headTemplate + o.contTemplate + o.footTemplate + '</table></div><div class="datetimepicker-hours"><table class=" table-condensed">' + o.headTemplate + o.contTemplate + o.footTemplate + '</table></div><div class="datetimepicker-days"><table class=" table-condensed">' + o.headTemplate + "<tbody></tbody>" + o.footTemplate + '</table></div><div class="datetimepicker-months"><table class="table-condensed">' + o.headTemplate + o.contTemplate + o.footTemplate + '</table></div><div class="datetimepicker-years"><table class="table-condensed">' + o.headTemplate + o.contTemplate + o.footTemplate + "</table></div></div>", t.fn.datetimepicker.DPGlobal = o, t.fn.datetimepicker.noConflict = function() {
		return t.fn.datetimepicker = old, this
	}, t(document).on("focus.datetimepicker.data-api click.datetimepicker.data-api", '[data-provide="datetimepicker"]', function(e) {
		var i = t(this);
		i.data("datetimepicker") || (e.preventDefault(), i.datetimepicker("show"))
	}), t(function() {
		t('[data-provide="datetimepicker-inline"]').datetimepicker()
	})
}(window.jQuery), /*! bootbox.js v4.4.0 http://bootboxjs.com/license.txt */

function(t, e) {
	"use strict";
	"function" == typeof define && define.amd ? define(["jquery"], e) : "object" == typeof exports ? module.exports = e(require("jquery")) : t.bootbox = e(t.jQuery)
}(this, function t(e, i) {
	"use strict";

	function n(t) {
		var e = m[f.locale];
		return e ? e[t] : m.en[t]
	}
	function o(t, i, n) {
		t.stopPropagation(), t.preventDefault();
		var o = e.isFunction(n) && n.call(i, t) === !1;
		o || i.modal("hide")
	}
	function s(t) {
		var e, i = 0;
		for (e in t) i++;
		return i
	}
	function a(t, i) {
		var n = 0;
		e.each(t, function(t, e) {
			i(t, e, n++)
		})
	}
	function r(t) {
		var i, n;
		if ("object" != typeof t) throw new Error("Please supply an object of options");
		if (!t.message) throw new Error("Please specify a message");
		return t = e.extend({}, f, t), t.buttons || (t.buttons = {}), i = t.buttons, n = s(i), a(i, function(t, o, s) {
			if (e.isFunction(o) && (o = i[t] = {
				callback: o
			}), "object" !== e.type(o)) throw new Error("button with key " + t + " must be an object");
			o.label || (o.label = t), o.className || (2 === n && ("ok" === t || "confirm" === t) || 1 === n ? o.className = "btn-primary" : o.className = "btn-default")
		}), t
	}
	function l(t, e) {
		var i = t.length,
			n = {};
		if (i < 1 || i > 2) throw new Error("Invalid argument length");
		return 2 === i || "string" == typeof t[0] ? (n[e[0]] = t[0], n[e[1]] = t[1]) : n = t[0], n
	}
	function h(t, i, n) {
		return e.extend(!0, {}, t, l(i, n))
	}
	function c(t, e, i, n) {
		var o = {
			className: "bootbox-" + t,
			buttons: d.apply(null, e)
		};
		return u(h(o, n, i), e)
	}
	function d() {
		for (var t = {}, e = 0, i = arguments.length; e < i; e++) {
			var o = arguments[e],
				s = o.toLowerCase(),
				a = o.toUpperCase();
			t[s] = {
				label: n(a)
			}
		}
		return t
	}
	function u(t, e) {
		var n = {};
		return a(e, function(t, e) {
			n[e] = !0
		}), a(t.buttons, function(t) {
			if (n[t] === i) throw new Error("button key " + t + " is not allowed (options are " + e.join("\n") + ")")
		}), t
	}
	var p = {
		dialog: "<div class='bootbox modal' tabindex='-1' role='dialog'><div class='modal-dialog'><div class='modal-content'><div class='modal-body'><div class='bootbox-body'></div></div></div></div></div>",
		header: "<div class='modal-header'><h4 class='modal-title'></h4></div>",
		footer: "<div class='modal-footer'></div>",
		closeButton: "<button type='button' class='bootbox-close-button close' data-dismiss='modal' aria-hidden='true'>&times;</button>",
		form: "<form class='bootbox-form'></form>",
		inputs: {
			text: "<input class='bootbox-input bootbox-input-text form-control' autocomplete=off type=text />",
			textarea: "<textarea class='bootbox-input bootbox-input-textarea form-control'></textarea>",
			email: "<input class='bootbox-input bootbox-input-email form-control' autocomplete='off' type='email' />",
			select: "<select class='bootbox-input bootbox-input-select form-control'></select>",
			checkbox: "<div class='checkbox'><label><input class='bootbox-input bootbox-input-checkbox' type='checkbox' /></label></div>",
			date: "<input class='bootbox-input bootbox-input-date form-control' autocomplete=off type='date' />",
			time: "<input class='bootbox-input bootbox-input-time form-control' autocomplete=off type='time' />",
			number: "<input class='bootbox-input bootbox-input-number form-control' autocomplete=off type='number' />",
			password: "<input class='bootbox-input bootbox-input-password form-control' autocomplete='off' type='password' />"
		}
	},
		f = {
			locale: e.zui && e.zui.clientLang ? e.zui.clientLang() : "zh_cn",
			backdrop: "static",
			animate: !0,
			className: null,
			closeButton: !0,
			show: !0,
			container: "body"
		},
		g = {};
	g.alert = function() {
		var t;
		if (t = c("alert", ["ok"], ["message", "callback"], arguments), t.callback && !e.isFunction(t.callback)) throw new Error("alert requires callback property to be a function when provided");
		return t.buttons.ok.callback = t.onEscape = function() {
			return !e.isFunction(t.callback) || t.callback.call(this)
		}, g.dialog(t)
	}, g.confirm = function() {
		var t;
		if (t = c("confirm", ["confirm", "cancel"], ["message", "callback"], arguments), t.buttons.cancel.callback = t.onEscape = function() {
			return t.callback.call(this, !1)
		}, t.buttons.confirm.callback = function() {
			return t.callback.call(this, !0)
		}, !e.isFunction(t.callback)) throw new Error("confirm requires a callback");
		return g.dialog(t)
	}, g.prompt = function() {
		var t, n, o, s, r, l, c;
		if (s = e(p.form), n = {
			className: "bootbox-prompt",
			buttons: d("cancel", "confirm"),
			value: "",
			inputType: "text"
		}, t = u(h(n, arguments, ["title", "callback"]), ["confirm", "cancel"]), l = t.show === i || t.show, t.message = s, t.buttons.cancel.callback = t.onEscape = function() {
			return t.callback.call(this, null)
		}, t.buttons.confirm.callback = function() {
			var i;
			switch (t.inputType) {
			case "text":
			case "textarea":
			case "email":
			case "select":
			case "date":
			case "time":
			case "number":
			case "password":
				i = r.val();
				break;
			case "checkbox":
				var n = r.find("input:checked");
				i = [], a(n, function(t, n) {
					i.push(e(n).val())
				})
			}
			return t.callback.call(this, i)
		}, t.show = !1, !t.title) throw new Error("prompt requires a title");
		if (!e.isFunction(t.callback)) throw new Error("prompt requires a callback");
		if (!p.inputs[t.inputType]) throw new Error("invalid prompt type");
		switch (r = e(p.inputs[t.inputType]), t.inputType) {
		case "text":
		case "textarea":
		case "email":
		case "date":
		case "time":
		case "number":
		case "password":
			r.val(t.value);
			break;
		case "select":
			var f = {};
			if (c = t.inputOptions || [], !e.isArray(c)) throw new Error("Please pass an array of input options");
			if (!c.length) throw new Error("prompt with select requires options");
			a(c, function(t, n) {
				var o = r;
				if (n.value === i || n.text === i) throw new Error("given options in wrong format");
				n.group && (f[n.group] || (f[n.group] = e("<optgroup/>").attr("label", n.group)), o = f[n.group]), o.append("<option value='" + n.value + "'>" + n.text + "</option>")
			}), a(f, function(t, e) {
				r.append(e)
			}), r.val(t.value);
			break;
		case "checkbox":
			var m = e.isArray(t.value) ? t.value : [t.value];
			if (c = t.inputOptions || [], !c.length) throw new Error("prompt with checkbox requires options");
			if (!c[0].value || !c[0].text) throw new Error("given options in wrong format");
			r = e("<div/>"), a(c, function(i, n) {
				var o = e(p.inputs[t.inputType]);
				o.find("input").attr("value", n.value), o.find("label").append(n.text), a(m, function(t, e) {
					e === n.value && o.find("input").prop("checked", !0)
				}), r.append(o)
			})
		}
		return t.placeholder && r.attr("placeholder", t.placeholder), t.pattern && r.attr("pattern", t.pattern), t.maxlength && r.attr("maxlength", t.maxlength), s.append(r), s.on("submit", function(t) {
			t.preventDefault(), t.stopPropagation(), o.find(".btn-primary").click()
		}), o = g.dialog(t), o.off("shown.zui.modal"), o.on("shown.zui.modal", function() {
			r.focus()
		}), l === !0 && o.modal("show"), o
	}, g.dialog = function(t) {
		t = r(t);
		var n = e(p.dialog),
			s = n.find(".modal-dialog"),
			l = n.find(".modal-body"),
			h = t.buttons,
			c = "",
			d = {
				onEscape: t.onEscape
			};
		if (e.fn.modal === i) throw new Error("$.fn.modal is not defined; please double check you have included the Bootstrap JavaScript library. See http://getbootstrap.com/javascript/ for more details.");
		if (a(h, function(t, e) {
			c += "<button data-bb-handler='" + t + "' type='button' class='btn " + e.className + "'>" + e.label + "</button>", d[t] = e.callback
		}), l.find(".bootbox-body").html(t.message), t.animate === !0 && n.addClass("fade"), t.className && n.addClass(t.className), "large" === t.size ? s.addClass("modal-lg") : "small" === t.size && s.addClass("modal-sm"), t.title && l.before(p.header), t.closeButton) {
			var u = e(p.closeButton);
			t.title ? n.find(".modal-header").prepend(u) : u.css("margin-top", "-10px").prependTo(l)
		}
		return t.title && n.find(".modal-title").html(t.title), c.length && (l.after(p.footer), n.find(".modal-footer").html(c)), n.on("hidden.zui.modal", function(t) {
			t.target === this && n.remove()
		}), n.on("shown.zui.modal", function() {
			n.find(".btn-primary:first").focus()
		}), "static" !== t.backdrop && n.on("click.dismiss.zui.modal", function(t) {
			n.children(".modal-backdrop").length && (t.currentTarget = n.children(".modal-backdrop").get(0)), t.target === t.currentTarget && n.trigger("escape.close.bb")
		}), n.on("escape.close.bb", function(t) {
			d.onEscape && o(t, n, d.onEscape)
		}), n.on("click", ".modal-footer button", function(t) {
			var i = e(this).data("bb-handler");
			o(t, n, d[i])
		}), n.on("click", ".bootbox-close-button", function(t) {
			o(t, n, d.onEscape)
		}), n.on("keyup", function(t) {
			27 === t.which && n.trigger("escape.close.bb")
		}), e(t.container).append(n), n.modal({
			backdrop: !! t.backdrop && "static",
			keyboard: !1,
			show: !1
		}), t.show && n.modal("show"), n
	}, g.setDefaults = function() {
		var t = {};
		2 === arguments.length ? t[arguments[0]] = arguments[1] : t = arguments[0], e.extend(f, t)
	}, g.hideAll = function() {
		return e(".bootbox").modal("hide"), g
	};
	var m = {
		en: {
			OK: "OK",
			CANCEL: "Cancel",
			CONFIRM: "OK"
		},
		zh_cn: {
			OK: "确认",
			CANCEL: "取消",
			CONFIRM: "确认"
		},
		zh_tw: {
			OK: "確認",
			CANCEL: "取消",
			CONFIRM: "確認"
		}
	};
	return g.addLocale = function(t, i) {
		return e.each(["OK", "CANCEL", "CONFIRM"], function(t, e) {
			if (!i[e]) throw new Error("Please supply a translation for '" + e + "'")
		}), m[t] = {
			OK: i.OK,
			CANCEL: i.CANCEL,
			CONFIRM: i.CONFIRM
		}, g
	}, g.removeLocale = function(t) {
		return delete m[t], g
	}, g.setLocale = function(t) {
		return g.setDefaults("locale", t)
	}, g.init = function(i) {
		return t(i || e)
	}, g
}),
/*!
Chosen, a Select Box Enhancer for jQuery and Prototype
by Patrick Filler for Harvest, http://getharvest.com

Version 1.1.0
Full source at https://github.com/harvesthq/chosen
Copyright (c) 2011 Harvest http://getharvest.com

MIT License, https://github.com/harvesthq/chosen/blob/master/LICENSE.md
*/
function() {
	var t, e, i, n, o, s = {}.hasOwnProperty,
		a = function(t, e) {
			function i() {
				this.constructor = t
			}
			for (var n in e) s.call(e, n) && (t[n] = e[n]);
			return i.prototype = e.prototype, t.prototype = new i, t.__super__ = e.prototype, t
		},
		r = {
			zh_cn: {
				no_results_text: "没有找到"
			},
			zh_tw: {
				no_results_text: "沒有找到"
			},
			en: {
				no_results_text: "No results match"
			}
		};
	n = function() {
		function e() {
			this.options_index = 0, this.parsed = []
		}
		return e.prototype.add_node = function(t) {
			return "OPTGROUP" === t.nodeName.toUpperCase() ? this.add_group(t) : this.add_option(t)
		}, e.prototype.add_group = function(e) {
			var i, n, o, s, a, r;
			for (i = this.parsed.length, this.parsed.push({
				array_index: i,
				group: !0,
				label: this.escapeExpression(e.label),
				children: 0,
				disabled: e.disabled,
				title: e.title,
				search_keys: t.trim(e.getAttribute("data-keys") || "").replace(/,/g, " ")
			}), a = e.childNodes, r = [], o = 0, s = a.length; o < s; o++) n = a[o], r.push(this.add_option(n, i, e.disabled));
			return r
		}, e.prototype.add_option = function(e, i, n) {
			if ("OPTION" === e.nodeName.toUpperCase()) return "" !== e.text ? (null != i && (this.parsed[i].children += 1), this.parsed.push({
				array_index: this.parsed.length,
				options_index: this.options_index,
				value: e.value,
				text: e.text,
				title: e.title,
				html: e.innerHTML,
				selected: e.selected,
				disabled: n === !0 ? n : e.disabled,
				group_array_index: i,
				classes: e.className,
				style: e.style.cssText,
				data: e.getAttribute("data-data"),
				search_keys: (t.trim(e.getAttribute("data-keys") || "") + e.value).replace(/,/, " ")
			})) : this.parsed.push({
				array_index: this.parsed.length,
				options_index: this.options_index,
				empty: !0
			}), this.options_index += 1
		}, e.prototype.escapeExpression = function(t) {
			var e, i;
			return null == t || t === !1 ? "" : /[\&\<\>\"\'\`]/.test(t) ? (e = {
				"<": "&lt;",
				">": "&gt;",
				'"': "&quot;",
				"'": "&#x27;",
				"`": "&#x60;"
			}, i = /&(?!\w+;)|[\<\>\"\'\`]/g, t.replace(i, function(t) {
				return e[t] || "&amp;"
			})) : t
		}, e
	}(), n.select_to_array = function(t) {
		var e, i, o, s, a;
		for (i = new n, a = t.childNodes, o = 0, s = a.length; o < s; o++) e = a[o], i.add_node(e);
		return i.parsed
	}, e = function() {
		function e(i, n) {
			this.form_field = i, this.options = null != n ? n : {}, e.browser_is_supported() && (this.lang = r[this.options.lang || (t.zui.clientLang ? t.zui.clientLang() : "zh_cn")], this.is_multiple = this.form_field.multiple, this.set_default_text(), this.set_default_values(), this.setup(), this.set_up_html(), this.register_observers())
		}
		return e.prototype.set_default_values = function() {
			var t = this;
			return this.click_test_action = function(e) {
				return t.test_active_click(e)
			}, this.activate_action = function(e) {
				return t.activate_field(e)
			}, this.active_field = !1, this.mouse_on_container = !1, this.results_showing = !1, this.result_highlighted = null, this.allow_single_deselect = null != this.options.allow_single_deselect && null != this.form_field.options[0] && "" === this.form_field.options[0].text && this.options.allow_single_deselect, this.disable_search_threshold = this.options.disable_search_threshold || 0, this.disable_search = this.options.disable_search || !1, this.enable_split_word_search = null == this.options.enable_split_word_search || this.options.enable_split_word_search, this.group_search = null == this.options.group_search || this.options.group_search, this.search_contains = this.options.search_contains || !1, this.single_backstroke_delete = null == this.options.single_backstroke_delete || this.options.single_backstroke_delete, this.max_selected_options = this.options.max_selected_options || 1 / 0, this.drop_direction = this.options.drop_direction || "auto", this.middle_highlight = this.options.middle_highlight, this.compact_search = this.options.compact_search || !0, this.inherit_select_classes = this.options.inherit_select_classes || !1, this.display_selected_options = null == this.options.display_selected_options || this.options.display_selected_options, this.display_disabled_options = null == this.options.display_disabled_options || this.options.display_disabled_options
		}, e.prototype.set_default_text = function() {
			return this.form_field.getAttribute("data-placeholder") ? this.default_text = this.form_field.getAttribute("data-placeholder") : this.is_multiple ? this.default_text = this.options.placeholder_text_multiple || this.options.placeholder_text || e.default_multiple_text : this.default_text = this.options.placeholder_text_single || this.options.placeholder_text || e.default_single_text, this.results_none_found = this.form_field.getAttribute("data-no_results_text") || this.options.no_results_text || this.lang.no_results_text || e.default_no_result_text
		}, e.prototype.mouse_enter = function() {
			return this.mouse_on_container = !0
		}, e.prototype.mouse_leave = function() {
			return this.mouse_on_container = !1
		}, e.prototype.input_focus = function(t) {
			var e = this;
			if (this.is_multiple) {
				if (!this.active_field) return setTimeout(function() {
					return e.container_mousedown()
				}, 50)
			} else if (!this.active_field) return this.activate_field()
		}, e.prototype.input_blur = function(t) {
			var e = this;
			if (!this.mouse_on_container) return this.active_field = !1, setTimeout(function() {
				return e.blur_test()
			}, 100)
		}, e.prototype.results_option_build = function(t) {
			var e, i, n, o, s;
			for (e = "", s = this.results_data, n = 0, o = s.length; n < o; n++) i = s[n], e += i.group ? this.result_add_group(i) : this.result_add_option(i), (null != t ? t.first : void 0) && (i.selected && this.is_multiple ? this.choice_build(i) : i.selected && !this.is_multiple && this.single_set_selected_text(i.text));
			return e
		}, e.prototype.result_add_option = function(t) {
			var e, i;
			return t.search_match && this.include_option_in_results(t) ? (e = [], t.disabled || t.selected && this.is_multiple || e.push("active-result"), !t.disabled || t.selected && this.is_multiple || e.push("disabled-result"), t.selected && e.push("result-selected"), null != t.group_array_index && e.push("group-option"), "" !== t.classes && e.push(t.classes), i = document.createElement("li"), i.className = e.join(" "), i.style.cssText = t.style, i.title = t.title, i.setAttribute("data-option-array-index", t.array_index), i.setAttribute("data-data", t.data), i.innerHTML = t.search_text, this.outerHTML(i)) : ""
		}, e.prototype.result_add_group = function(t) {
			var e;
			return (t.search_match || t.group_match) && t.active_options > 0 ? (e = document.createElement("li"), e.className = "group-result", e.title = t.title, e.innerHTML = t.search_text, this.outerHTML(e)) : ""
		}, e.prototype.results_update_field = function() {
			if (this.set_default_text(), this.is_multiple || this.results_reset_cleanup(), this.result_clear_highlight(), this.results_build(), this.results_showing) return this.winnow_results()
		}, e.prototype.reset_single_select_options = function() {
			var t, e, i, n, o;
			for (n = this.results_data, o = [], e = 0, i = n.length; e < i; e++) t = n[e], t.selected ? o.push(t.selected = !1) : o.push(void 0);
			return o
		}, e.prototype.results_toggle = function() {
			return this.results_showing ? this.results_hide() : this.results_show()
		}, e.prototype.results_search = function(t) {
			return this.results_showing ? this.winnow_results(1) : this.results_show()
		}, e.prototype.winnow_results = function(t) {
			var e, i, n, o, s, a, r, l, h, c, d, u, p;
			for (this.no_results_clear(), s = 0, r = this.get_search_text(), e = r.replace(/[-[\]{}()*+?.,\\^$|#\s]/g, "\\$&"), o = this.search_contains ? "" : "^", n = new RegExp(o + e, "i"), c = new RegExp(e, "i"), p = this.results_data, d = 0, u = p.length; d < u; d++) i = p[d], i.search_match = !1, a = null, this.include_option_in_results(i) && (i.group && (i.group_match = !1, i.active_options = 0), null != i.group_array_index && this.results_data[i.group_array_index] && (a = this.results_data[i.group_array_index], 0 === a.active_options && a.search_match && (s += 1), a.active_options += 1), i.group && !this.group_search || (i.search_text = i.group ? i.label : i.html, i.search_keys_match = this.search_string_match(i.search_keys, n), i.search_text_match = this.search_string_match(i.search_text, n), i.search_match = i.search_text_match || i.search_keys_match, i.search_match && !i.group && (s += 1), i.search_match ? (i.search_text_match && i.search_text.length ? (l = i.search_text.search(c), h = i.search_text.substr(0, l + r.length) + "</em>" + i.search_text.substr(l + r.length), i.search_text = h.substr(0, l) + "<em>" + h.substr(l)) : i.search_keys_match && i.search_keys.length && (l = i.search_keys.search(c), h = i.search_keys.substr(0, l + r.length) + "</em>" + i.search_keys.substr(l + r.length), i.search_text += '&nbsp; <small style="opacity: 0.7">' + h.substr(0, l) + "<em>" + h.substr(l) + "</small>"), null != a && (a.group_match = !0)) : null != i.group_array_index && this.results_data[i.group_array_index].search_match && (i.search_match = !0)));
			return this.result_clear_highlight(), s < 1 && r.length ? (this.update_results_content(""), this.no_results(r)) : (this.update_results_content(this.results_option_build()), this.winnow_results_set_highlight(t))
		}, e.prototype.search_string_match = function(t, e) {
			var i, n, o, s;
			if (e.test(t)) return !0;
			if (this.enable_split_word_search && (t.indexOf(" ") >= 0 || 0 === t.indexOf("[")) && (n = t.replace(/\[|\]/g, "").split(" "), n.length)) for (o = 0, s = n.length; o < s; o++) if (i = n[o], e.test(i)) return !0
		}, e.prototype.choices_count = function() {
			var t, e, i, n;
			if (null != this.selected_option_count) return this.selected_option_count;
			for (this.selected_option_count = 0, n = this.form_field.options, e = 0, i = n.length; e < i; e++) t = n[e], t.selected && "" != t.value && (this.selected_option_count += 1);
			return this.selected_option_count
		}, e.prototype.choices_click = function(t) {
			if (t.preventDefault(), !this.results_showing && !this.is_disabled) return this.results_show()
		}, e.prototype.keyup_checker = function(t) {
			var e, i;
			switch (e = null != (i = t.which) ? i : t.keyCode, this.search_field_scale(), e) {
			case 8:
				if (this.is_multiple && this.backstroke_length < 1 && this.choices_count() > 0) return this.keydown_backstroke();
				if (!this.pending_backstroke) return this.result_clear_highlight(), this.results_search();
				break;
			case 13:
				if (t.preventDefault(), this.results_showing) return this.result_select(t);
				break;
			case 27:
				return this.results_showing && this.results_hide(), !0;
			case 9:
			case 38:
			case 40:
			case 16:
			case 91:
			case 17:
				break;
			default:
				return this.results_search()
			}
		}, e.prototype.clipboard_event_checker = function(t) {
			var e = this;
			return setTimeout(function() {
				return e.results_search()
			}, 50)
		}, e.prototype.container_width = function() {
			return null != this.options.width ? this.options.width : this.form_field && this.form_field.classList && this.form_field.classList.contains("form-control") ? "100%" : "" + this.form_field.offsetWidth + "px"
		}, e.prototype.include_option_in_results = function(t) {
			return !(this.is_multiple && !this.display_selected_options && t.selected) && (!(!this.display_disabled_options && t.disabled) && !t.empty)
		}, e.prototype.search_results_touchstart = function(t) {
			return this.touch_started = !0, this.search_results_mouseover(t)
		}, e.prototype.search_results_touchmove = function(t) {
			return this.touch_started = !1, this.search_results_mouseout(t)
		}, e.prototype.search_results_touchend = function(t) {
			if (this.touch_started) return this.search_results_mouseup(t)
		}, e.prototype.outerHTML = function(t) {
			var e;
			return t.outerHTML ? t.outerHTML : (e = document.createElement("div"), e.appendChild(t), e.innerHTML)
		}, e.browser_is_supported = function() {
			return "Microsoft Internet Explorer" === window.navigator.appName ? document.documentMode >= 8 : !/iP(od|hone)/i.test(window.navigator.userAgent) && (!/Android/i.test(window.navigator.userAgent) || !/Mobile/i.test(window.navigator.userAgent))
		}, e.default_multiple_text = "", e.default_single_text = "", e.default_no_result_text = "No results match", e
	}(), t = jQuery, t.fn.extend({
		chosen: function(n) {
			return e.browser_is_supported() ? this.each(function(e) {
				var o, s;
				o = t(this), s = o.data("chosen"), "destroy" === n && s ? s.destroy() : s || o.data("chosen", new i(this, n))
			}) : this
		}
	}), i = function(e) {
		function i() {
			return o = i.__super__.constructor.apply(this, arguments)
		}
		return a(i, e), i.prototype.setup = function() {
			return this.form_field_jq = t(this.form_field), this.current_selectedIndex = this.form_field.selectedIndex, this.is_rtl = this.form_field_jq.hasClass("chosen-rtl")
		}, i.prototype.set_up_html = function() {
			var e, i;
			e = ["chosen-container"], e.push("chosen-container-" + (this.is_multiple ? "multi" : "single")), this.inherit_select_classes && this.form_field.className && e.push(this.form_field.className), this.is_rtl && e.push("chosen-rtl");
			var n = this.form_field.getAttribute("data-css-class");
			return n && e.push(n), i = {
				"class": e.join(" "),
				style: "width: " + this.container_width() + ";",
				title: this.form_field.title
			}, this.form_field.id.length && (i.id = this.form_field.id.replace(/[^\w]/g, "_") + "_chosen"), this.container = t("<div />", i), this.is_multiple ? this.container.html('<ul class="chosen-choices"><li class="search-field"><input type="text" value="' + this.default_text + '" class="default" autocomplete="off" style="width:25px;" /></li></ul><div class="chosen-drop"><ul class="chosen-results"></ul></div>') : (this.container.html('<a class="chosen-single chosen-default" tabindex="-1"><span>' + this.default_text + '</span><div><b></b></div><div class="chosen-search"><input type="text" autocomplete="off" /></div></a><div class="chosen-drop"><ul class="chosen-results"></ul></div>'), this.compact_search && this.container.find(".chosen-search").appendTo(this.container.find(".chosen-single"))), this.form_field_jq.hide().after(this.container), this.dropdown = this.container.find("div.chosen-drop").first(), this.search_field = this.container.find("input").first(), this.search_results = this.container.find("ul.chosen-results").first(), this.search_field_scale(), this.search_no_results = this.container.find("li.no-results").first(), this.is_multiple ? (this.search_choices = this.container.find("ul.chosen-choices").first(), this.search_container = this.container.find("li.search-field").first()) : (this.search_container = this.container.find("div.chosen-search").first(), this.selected_item = this.container.find(".chosen-single").first()), this.options.drop_width && this.dropdown.css("width", this.options.drop_width).addClass("chosen-drop-size-limited"), this.results_build(), this.set_tab_index(), this.set_label_behavior(), this.form_field_jq.trigger("chosen:ready", {
				chosen: this
			})
		}, i.prototype.register_observers = function() {
			var t = this;
			return this.container.bind("mousedown.chosen", function(e) {
				t.container_mousedown(e)
			}), this.container.bind("mouseup.chosen", function(e) {
				t.container_mouseup(e)
			}), this.container.bind("mouseenter.chosen", function(e) {
				t.mouse_enter(e)
			}), this.container.bind("mouseleave.chosen", function(e) {
				t.mouse_leave(e)
			}), this.search_results.bind("mouseup.chosen", function(e) {
				t.search_results_mouseup(e)
			}), this.search_results.bind("mouseover.chosen", function(e) {
				t.search_results_mouseover(e)
			}), this.search_results.bind("mouseout.chosen", function(e) {
				t.search_results_mouseout(e)
			}), this.search_results.bind("mousewheel.chosen DOMMouseScroll.chosen", function(e) {
				t.search_results_mousewheel(e)
			}), this.search_results.bind("touchstart.chosen", function(e) {
				t.search_results_touchstart(e)
			}), this.search_results.bind("touchmove.chosen", function(e) {
				t.search_results_touchmove(e)
			}), this.search_results.bind("touchend.chosen", function(e) {
				t.search_results_touchend(e)
			}), this.form_field_jq.bind("chosen:updated.chosen", function(e) {
				t.results_update_field(e)
			}), this.form_field_jq.bind("chosen:activate.chosen", function(e) {
				t.activate_field(e)
			}), this.form_field_jq.bind("chosen:open.chosen", function(e) {
				t.container_mousedown(e)
			}), this.form_field_jq.bind("chosen:close.chosen", function(e) {
				t.input_blur(e)
			}), this.search_field.bind("blur.chosen", function(e) {
				t.input_blur(e)
			}), this.search_field.bind("keyup.chosen", function(e) {
				t.keyup_checker(e)
			}), this.search_field.bind("keydown.chosen", function(e) {
				t.keydown_checker(e)
			}), this.search_field.bind("focus.chosen", function(e) {
				t.input_focus(e)
			}), this.search_field.bind("cut.chosen", function(e) {
				t.clipboard_event_checker(e)
			}), this.search_field.bind("paste.chosen", function(e) {
				t.clipboard_event_checker(e)
			}), this.is_multiple ? this.search_choices.bind("click.chosen", function(e) {
				t.choices_click(e)
			}) : this.container.bind("click.chosen", function(t) {
				t.preventDefault()
			})
		}, i.prototype.destroy = function() {
			return t(this.container[0].ownerDocument).unbind("click.chosen", this.click_test_action), this.search_field[0].tabIndex && (this.form_field_jq[0].tabIndex = this.search_field[0].tabIndex), this.container.remove(), this.form_field_jq.removeData("chosen"), this.form_field_jq.show()
		}, i.prototype.search_field_disabled = function() {
			return this.is_disabled = this.form_field_jq[0].disabled, this.is_disabled ? (this.container.addClass("chosen-disabled"), this.search_field[0].disabled = !0, this.is_multiple || this.selected_item.unbind("focus.chosen", this.activate_action), this.close_field()) : (this.container.removeClass("chosen-disabled"), this.search_field[0].disabled = !1, this.is_multiple ? void 0 : this.selected_item.bind("focus.chosen", this.activate_action))
		}, i.prototype.container_mousedown = function(e) {
			if (!this.is_disabled && (e && "mousedown" === e.type && !this.results_showing && e.preventDefault(), null == e || !t(e.target).hasClass("search-choice-close"))) return this.active_field ? this.is_multiple || !e || t(e.target)[0] !== this.selected_item[0] && !t(e.target).parents("a.chosen-single").length || (e.preventDefault(), this.results_toggle()) : (this.is_multiple && this.search_field.val(""), t(this.container[0].ownerDocument).bind("click.chosen", this.click_test_action), this.results_show()), this.activate_field()
		}, i.prototype.container_mouseup = function(t) {
			if ("ABBR" === t.target.nodeName && !this.is_disabled) return this.results_reset(t)
		}, i.prototype.search_results_mousewheel = function(t) {
			var e;
			if (t.originalEvent && (e = -t.originalEvent.wheelDelta || t.originalEvent.detail), null != e) return t.preventDefault(), "DOMMouseScroll" === t.type && (e = 40 * e), this.search_results.scrollTop(e + this.search_results.scrollTop())
		}, i.prototype.blur_test = function(t) {
			if (!this.active_field && this.container.hasClass("chosen-container-active")) return this.close_field()
		}, i.prototype.close_field = function() {
			return t(this.container[0].ownerDocument).unbind("click.chosen", this.click_test_action), this.active_field = !1, this.results_hide(), this.container.removeClass("chosen-container-active"), this.clear_backstroke(), this.show_search_field_default(), this.search_field_scale()
		}, i.prototype.activate_field = function() {
			return this.container.addClass("chosen-container-active"), this.active_field = !0, this.search_field.val(this.search_field.val()), this.search_field.focus()
		}, i.prototype.test_active_click = function(e) {
			var i;
			return i = t(e.target).closest(".chosen-container"), i.length && this.container[0] === i[0] ? this.active_field = !0 : this.close_field()
		}, i.prototype.results_build = function() {
			return this.parsing = !0, this.selected_option_count = null, this.results_data = n.select_to_array(this.form_field), this.is_multiple ? this.search_choices.find("li.search-choice").remove() : this.is_multiple || (this.single_set_selected_text(), this.disable_search || this.form_field.options.length <= this.disable_search_threshold ? (this.search_field[0].readOnly = !0, this.container.addClass("chosen-container-single-nosearch"), this.container.removeClass("chosen-with-search")) : (this.search_field[0].readOnly = !1, this.container.removeClass("chosen-container-single-nosearch"), this.container.addClass("chosen-with-search"))), this.update_results_content(this.results_option_build({
				first: !0
			})), this.search_field_disabled(), this.show_search_field_default(), this.search_field_scale(), this.parsing = !1
		}, i.prototype.result_do_highlight = function(t, e) {
			var i, n, o, s, a, r, l = -1;
			t.length && (this.result_clear_highlight(), this.result_highlight = t, this.result_highlight.addClass("highlighted"), o = parseInt(this.search_results.css("maxHeight"), 10), r = this.result_highlight.outerHeight(), a = this.search_results.scrollTop(), s = o + a, n = this.result_highlight.position().top + this.search_results.scrollTop(), i = n + r, this.middle_highlight && (e || "always" === this.middle_highlight || i >= s || n < a) ? l = Math.min(n - r, Math.max(0, n - (o - r) / 2)) : i >= s ? l = i - o > 0 ? i - o : 0 : n < a && (l = n), l > -1 && this.search_results.scrollTop(l))
		}, i.prototype.result_clear_highlight = function() {
			return this.result_highlight && this.result_highlight.removeClass("highlighted"), this.result_highlight = null
		}, i.prototype.results_show = function() {
			if (this.is_multiple && this.max_selected_options <= this.choices_count()) return this.form_field_jq.trigger("chosen:maxselected", {
				chosen: this
			}), !1;
			this.results_showing = !0, this.search_field.focus(), this.search_field.val(this.search_field.val()), this.winnow_results(1);
			var e = this.drop_direction;
			if (t.isFunction(e) && (e = e.call(this)), "auto" === e) if (this.drop_directionFixed) e = this.drop_directionFixed;
			else {
				var i = this.container.find(".chosen-drop"),
					n = this.container.offset();
				n.top + i.outerHeight() + 30 > t(window).height() + t(window).scrollTop() && (e = "up"), this.drop_directionFixed = e
			}
			return this.container.toggleClass("chosen-up", "up" === e).addClass("chosen-with-drop"), this.form_field_jq.trigger("chosen:showing_dropdown", {
				chosen: this
			})
		}, i.prototype.update_results_content = function(t) {
			return this.search_results.html(t)
		}, i.prototype.results_hide = function() {
			return this.results_showing && (this.result_clear_highlight(), this.container.removeClass("chosen-with-drop"), this.form_field_jq.trigger("chosen:hiding_dropdown", {
				chosen: this
			}), this.drop_directionFixed = 0), this.results_showing = !1
		}, i.prototype.set_tab_index = function(t) {
			var e;
			if (this.form_field.tabIndex) return e = this.form_field.tabIndex, this.form_field.tabIndex = -1, this.search_field[0].tabIndex = e
		}, i.prototype.set_label_behavior = function() {
			var e = this;
			if (this.form_field_label = this.form_field_jq.parents("label"), !this.form_field_label.length && this.form_field.id.length && (this.form_field_label = t("label[for='" + this.form_field.id + "']")), this.form_field_label.length > 0) return this.form_field_label.bind("click.chosen", function(t) {
				return e.is_multiple ? e.container_mousedown(t) : e.activate_field()
			})
		}, i.prototype.show_search_field_default = function() {
			return this.is_multiple && this.choices_count() < 1 && !this.active_field ? (this.search_field.val(this.default_text), this.search_field.addClass("default")) : (this.search_field.val(""), this.search_field.removeClass("default"))
		}, i.prototype.search_results_mouseup = function(e) {
			var i;
			if (i = t(e.target).hasClass("active-result") ? t(e.target) : t(e.target).parents(".active-result").first(), i.length) return this.result_highlight = i, this.result_select(e), this.search_field.focus()
		}, i.prototype.search_results_mouseover = function(e) {
			var i;
			if (i = t(e.target).hasClass("active-result") ? t(e.target) : t(e.target).parents(".active-result").first()) return this.result_do_highlight(i)
		}, i.prototype.search_results_mouseout = function(e) {
			if (t(e.target).hasClass("active-result")) return this.result_clear_highlight()
		}, i.prototype.choice_build = function(e) {
			var i, n, o = this;
			return i = t("<li />", {
				"class": "search-choice"
			}).html("<span title='" + e.html + "'>" + e.html + "</span>"), e.disabled ? i.addClass("search-choice-disabled") : (n = t("<a />", {
				"class": "search-choice-close",
				"data-option-array-index": e.array_index
			}), n.bind("click.chosen", function(t) {
				return o.choice_destroy_link_click(t)
			}), i.append(n)), this.search_container.before(i)
		}, i.prototype.choice_destroy_link_click = function(e) {
			if (e.preventDefault(), e.stopPropagation(), !this.is_disabled) return this.choice_destroy(t(e.target))
		}, i.prototype.choice_destroy = function(t) {
			if (this.result_deselect(t[0].getAttribute("data-option-array-index"))) return this.show_search_field_default(), this.is_multiple && this.choices_count() > 0 && this.search_field.val().length < 1 && this.results_hide(), t.parents("li").first().remove(), this.search_field_scale()
		}, i.prototype.results_reset = function() {
			if (this.reset_single_select_options(), this.form_field.options[0].selected = !0, this.single_set_selected_text(), this.show_search_field_default(), this.results_reset_cleanup(), this.form_field_jq.trigger("change"), this.active_field) return this.results_hide()
		}, i.prototype.results_reset_cleanup = function() {
			return this.current_selectedIndex = this.form_field.selectedIndex, this.selected_item.find("abbr").remove()
		}, i.prototype.result_select = function(t) {
			var e, i;
			if (this.result_highlight) return e = this.result_highlight, this.result_clear_highlight(), this.is_multiple && this.max_selected_options <= this.choices_count() ? (this.form_field_jq.trigger("chosen:maxselected", {
				chosen: this
			}), !1) : (this.is_multiple ? e.removeClass("active-result") : this.reset_single_select_options(), i = this.results_data[e[0].getAttribute("data-option-array-index")], i.selected = !0, this.form_field.options[i.options_index].selected = !0, this.selected_option_count = null, this.is_multiple ? this.choice_build(i) : this.single_set_selected_text(i.text), (t.metaKey || t.ctrlKey) && this.is_multiple || this.results_hide(), this.search_field.val(""), (this.is_multiple || this.form_field.selectedIndex !== this.current_selectedIndex) && this.form_field_jq.trigger("change", {
				selected: this.form_field.options[i.options_index].value
			}), this.current_selectedIndex = this.form_field.selectedIndex, this.search_field_scale())
		}, i.prototype.single_set_selected_text = function(t) {
			return null == t && (t = this.default_text), t === this.default_text ? this.selected_item.addClass("chosen-default") : (this.single_deselect_control_build(), this.selected_item.removeClass("chosen-default")), this.compact_search && this.search_field.attr("placeholder", t), this.selected_item.find("span").attr("title", t).text(t)
		}, i.prototype.result_deselect = function(t) {
			var e;
			return e = this.results_data[t], !this.form_field.options[e.options_index].disabled && (e.selected = !1, this.form_field.options[e.options_index].selected = !1, this.selected_option_count = null, this.result_clear_highlight(), this.results_showing && this.winnow_results(), this.form_field_jq.trigger("change", {
				deselected: this.form_field.options[e.options_index].value
			}), this.search_field_scale(), !0)
		}, i.prototype.single_deselect_control_build = function() {
			if (this.allow_single_deselect) return this.selected_item.find("abbr").length || this.selected_item.find("span").first().after('<abbr class="search-choice-close"></abbr>'), this.selected_item.addClass("chosen-single-with-deselect")
		}, i.prototype.get_search_text = function() {
			return this.search_field.val() === this.default_text ? "" : t("<div/>").text(t.trim(this.search_field.val())).html()
		}, i.prototype.winnow_results_set_highlight = function(t) {
			var e, i;
			if (i = this.is_multiple ? [] : this.search_results.find(".result-selected.active-result"), e = i.length ? i.first() : this.search_results.find(".active-result").first(), null != e) return this.result_do_highlight(e, t)
		}, i.prototype.no_results = function(e) {
			var i;
			return i = t('<li class="no-results">' + this.results_none_found + ' "<span></span>"</li>'), i.find("span").first().html(e), this.search_results.append(i), this.form_field_jq.trigger("chosen:no_results", {
				chosen: this
			})
		}, i.prototype.no_results_clear = function() {
			return this.search_results.find(".no-results").remove()
		}, i.prototype.keydown_arrow = function() {
			var t;
			return this.results_showing && this.result_highlight ? (t = this.result_highlight.nextAll("li.active-result").first()) ? this.result_do_highlight(t) : void 0 : this.results_show()
		}, i.prototype.keyup_arrow = function() {
			var t;
			return this.results_showing || this.is_multiple ? this.result_highlight ? (t = this.result_highlight.prevAll("li.active-result"), t.length ? this.result_do_highlight(t.first()) : (this.choices_count() > 0 && this.results_hide(), this.result_clear_highlight())) : void 0 : this.results_show()
		}, i.prototype.keydown_backstroke = function() {
			var t;
			return this.pending_backstroke ? (this.choice_destroy(this.pending_backstroke.find("a").first()), this.clear_backstroke()) : (t = this.search_container.siblings("li.search-choice").last(), t.length && !t.hasClass("search-choice-disabled") ? (this.pending_backstroke = t, this.single_backstroke_delete ? this.keydown_backstroke() : this.pending_backstroke.addClass("search-choice-focus")) : void 0)
		}, i.prototype.clear_backstroke = function() {
			return this.pending_backstroke && this.pending_backstroke.removeClass("search-choice-focus"), this.pending_backstroke = null
		}, i.prototype.keydown_checker = function(t) {
			var e, i;
			switch (e = null != (i = t.which) ? i : t.keyCode, this.search_field_scale(), 8 !== e && this.pending_backstroke && this.clear_backstroke(), e) {
			case 8:
				this.backstroke_length = this.search_field.val().length;
				break;
			case 9:
				this.results_showing && !this.is_multiple && this.result_select(t), this.mouse_on_container = !1;
				break;
			case 13:
				t.preventDefault();
				break;
			case 38:
				t.preventDefault(), this.keyup_arrow();
				break;
			case 40:
				t.preventDefault(), this.keydown_arrow()
			}
		}, i.prototype.search_field_scale = function() {
			var e, i, n, o, s, a, r, l, h;
			if (this.is_multiple) {
				for (n = 0, r = 0, s = "position:absolute; left: -1000px; top: -1000px; display:none;", a = ["font-size", "font-style", "font-weight", "font-family", "line-height", "text-transform", "letter-spacing"], l = 0, h = a.length; l < h; l++) o = a[l], s += o + ":" + this.search_field.css(o) + ";";
				return e = t("<div />", {
					style: s
				}), e.text(this.search_field.val()), t("body").append(e), r = e.width() + 25, e.remove(), i = this.container.outerWidth(), r > i - 10 && (r = i - 10), this.search_field.css({
					width: r + "px"
				})
			}
		}, i
	}(e)
}.call(this), function(t) {
	"use strict";
	var e = "zui.selectable",
		i = function(i, n) {
			this.name = e, this.$ = t(i), this.id = t.zui.uuid(), this.selectOrder = 1, this.selections = {}, this.getOptions(n), this._init()
		},
		n = function(t, e, i) {
			return t >= i.left && t <= i.left + i.width && e >= i.top && e <= i.top + i.height
		},
		o = function(t, e) {
			var i = Math.max(t.left, e.left),
				o = Math.max(t.top, e.top),
				s = Math.min(t.left + t.width, e.left + e.width),
				a = Math.min(t.top + t.height, e.top + e.height);
			return n(i, o, t) && n(s, a, t) && n(i, o, e) && n(s, a, e)
		};
	i.DEFAULTS = {
		selector: "li,tr,div",
		trigger: "",
		selectClass: "active",
		rangeStyle: {
			border: "1px solid " + (t.zui.colorset ? t.zui.colorset.primary : "#3280fc"),
			backgroundColor: t.zui.colorset ? new t.zui.Color(t.zui.colorset.primary).fade(20).toCssStr() : "rgba(50, 128, 252, 0.2)"
		},
		clickBehavior: "toggle",
		ignoreVal: 3,
		listenClick: !0
	}, i.prototype.getOptions = function(e) {
		this.options = t.extend({}, i.DEFAULTS, this.$.data(), e)
	}, i.prototype.select = function(t) {
		this.toggle(t, !0)
	}, i.prototype.unselect = function(t) {
		this.toggle(t, !1)
	}, i.prototype.toggle = function(e, i, n) {
		var o, s, a = this.options.selector,
			r = this;
		if (void 0 === e) return void this.$.find(a).each(function() {
			r.toggle(this, i)
		});
		if ("object" == typeof e ? (o = t(e).closest(a), s = o.data("id")) : (s = e, o = r.$.find('.slectable-item[data-id="' + s + '"]')), o && o.length) {
			if (s || (s = t.zui.uuid(), o.attr("data-id", s)), void 0 !== i && null !== i || (i = !r.selections[s]), !! i != !! r.selections[s]) {
				var l;
				t.isFunction(n) && (l = n(i)), l !== !0 && (r.selections[s] = !! i && r.selectOrder++, r.callEvent(i ? "select" : "unselect", {
					id: s,
					selections: r.selections,
					target: o,
					selected: r.getSelectedArray()
				}, r))
			}
			r.options.selectClass && o.toggleClass(r.options.selectClass, i)
		}
	}, i.prototype.getSelectedArray = function() {
		var e = [];
		return t.each(this.selections, function(t, i) {
			i && e.push(t)
		}), e
	}, i.prototype.syncSelectionsFromClass = function() {
		var e = this;
		e.$children = e.$.find(e.options.selector);
		e.selections = {}, e.$children.each(function() {
			var i = t(this);
			e.selections[i.data("id")] = i.hasClass(e.options.selectClass)
		})
	}, i.prototype._init = function() {
		var e, i, n, s, a, r, l, h = this.options,
			c = this,
			d = h.ignoreVal,
			u = !0,
			p = "." + this.name + "." + this.id,
			f = t.isFunction(h.checkFunc) ? h.checkFunc : null,
			g = t.isFunction(h.rangeFunc) ? h.rangeFunc : null,
			m = !1,
			v = null,
			y = "mousedown" + p,
			b = function() {
				s && c.$children.each(function() {
					var e = t(this),
						i = e.offset();
					i.width = e.outerWidth(), i.height = e.outerHeight();
					var n = g ? g.call(this, s, i) : o(s, i);
					if (f) {
						var a = f.call(c, {
							intersect: n,
							target: e,
							range: s,
							targetRange: i
						});
						a === !0 ? c.select(e) : a === !1 && c.unselect(e)
					} else n ? c.select(e) : c.multiKey || c.unselect(e)
				})
			},
			w = function(o) {
				m && (a = o.pageX, r = o.pageY, s = {
					width: Math.abs(a - e),
					height: Math.abs(r - i),
					left: a > e ? e : a,
					top: r > i ? i : r
				}, u && s.width < d && s.height < d || (n || (n = t('.selectable-range[data-id="' + c.id + '"]'), n.length || (n = t('<div class="selectable-range" data-id="' + c.id + '"></div>').css(t.extend({
					zIndex: 1060,
					position: "absolute",
					top: e,
					left: i,
					pointerEvents: "none"
				}, c.options.rangeStyle)).appendTo(t("body")))), n.css(s), clearTimeout(l), l = setTimeout(b, 10), u = !1))
			},
			x = function(e) {
				t(document).off(p), clearTimeout(v), m && (m = !1, n && n.remove(), u || s && (clearTimeout(l), b(), s = null), c.callEvent("finish", {
					selections: c.selections,
					selected: c.getSelectedArray()
				}), e.preventDefault())
			},
			C = function(o) {
				if (m) return x(o);
				var s = t.zui.getMouseButtonCode(h.mouseButton);
				if (!(s > -1 && o.button !== s || c.altKey || 3 === o.which || c.callEvent("start", o) === !1)) {
					var a = c.$children = c.$.find(h.selector);
					a.addClass("slectable-item");
					var r = c.multiKey ? "multi" : h.clickBehavior;
					if ("single" === r && c.unselect(), h.listenClick && ("multi" === r ? c.toggle(o.target) : "single" === r ? c.select(o.target) : "toggle" === r && c.toggle(o.target, null, function(t) {
						c.unselect()
					})), c.callEvent("startDrag", o) === !1) return void c.callEvent("finish", {
						selections: c.selections,
						selected: c.getSelectedArray()
					});
					e = o.pageX, i = o.pageY, n = null, u = !0, m = !0, t(document).on("mousemove" + p, w).on("mouseup" + p, x), v = setTimeout(function() {
						t(document).on(y, x)
					}, 10), o.preventDefault()
				}
			},
			_ = h.container && "default" !== h.container ? t(h.container) : this.$;
		h.trigger ? _.on(y, h.trigger, C) : _.on(y, C), t(document).on("keydown", function(t) {
			var e = t.keyCode;
			17 === e || 91 == e ? c.multiKey = e : 18 === e && (c.altKey = !0);
		}).on("keyup", function(t) {
			c.multiKey = !1, c.altKey = !1
		})
	}, i.prototype.callEvent = function(e, i) {
		var n = t.Event(e + "." + this.name);
		this.$.trigger(n, i);
		var o = n.result,
			s = this.options[e];
		return t.isFunction(s) && (o = s.apply(this, t.isArray(i) ? i : [i])), o
	}, t.fn.selectable = function(n) {
		return this.each(function() {
			var o = t(this),
				s = o.data(e),
				a = "object" == typeof n && n;
			s || o.data(e, s = new i(this, a)), "string" == typeof n && s[n]()
		})
	}, t.fn.selectable.Constructor = i, t(function() {
		t('[data-ride="selectable"]').selectable()
	})
}(jQuery), +
function(t, e, i) {
	"use strict";
	if (!t.fn.droppable) return void console.error("Sortable requires droppable.js");
	var n = "zui.sortable",
		o = {
			selector: "li,div",
			dragCssClass: "invisible",
			sortingClass: "sortable-sorting"
		},
		s = "order",
		a = function(e, i) {
			var n = this;
			n.$ = t(e), n.options = t.extend({}, o, n.$.data(), i), n.init()
		};
	a.DEFAULTS = o, a.NAME = n, a.prototype.init = function() {
		var e, i = this,
			n = i.$,
			o = i.options,
			a = o.selector,
			r = o.containerSelector,
			l = o.sortingClass,
			h = o.dragCssClass,
			c = o.targetSelector,
			d = o.reverse,
			u = function(e) {
				e = e || i.getItems(1);
				var n = e.length;
				n && e.each(function(e) {
					var i = d ? n - e : e;
					t(this).attr("data-" + s, i).data(s, i)
				})
			};
		u(), n.droppable({
			handle: o.trigger,
			target: c ? c : r ? a + "," + r : a,
			selector: a,
			container: n,
			always: o.always,
			flex: !0,
			lazy: o.lazy,
			canMoveHere: o.canMoveHere,
			dropToClass: o.dropToClass,
			before: o.before,
			nested: !! r,
			mouseButton: o.mouseButton,
			stopPropagation: o.stopPropagation,
			start: function(t) {
				h && t.element.addClass(h), e = !1, i.trigger("start", t)
			},
			drag: function(t) {
				if (n.addClass(l), t.isIn) {
					var o = t.element,
						h = t.target,
						c = r && h.is(r);
					if (c) {
						if (!h.children(a).filter(".dragging").length) {
							h.append(o);
							var p = i.getItems(1);
							u(p), i.trigger(s, {
								list: p,
								element: o
							})
						}
						return
					}
					var f = o.data(s),
						g = h.data(s);
					if (f === g) return u(p);
					f > g ? h[d ? "after" : "before"](o) : h[d ? "before" : "after"](o), e = !0;
					var p = i.getItems(1);
					u(p), i.trigger(s, {
						list: p,
						element: o
					})
				}
			},
			finish: function(t) {
				h && t.element && t.element.removeClass(h), n.removeClass(l), i.trigger("finish", {
					list: i.getItems(),
					element: t.element,
					changed: e
				})
			}
		})
	}, a.prototype.destroy = function() {
		this.$.droppable("destroy"), this.$.data(n, null)
	}, a.prototype.reset = function() {
		this.destroy(), this.init()
	}, a.prototype.getItems = function(e) {
		var i = this.$.find(this.options.selector).not(".drag-shadow");
		return e ? i : i.map(function() {
			var e = t(this);
			return {
				item: e,
				order: e.data("order")
			}
		})
	}, a.prototype.trigger = function(e, i) {
		return t.zui.callEvent(this.options[e], i, this)
	}, t.fn.sortable = function(e) {
		return this.each(function() {
			var i = t(this),
				o = i.data(n),
				s = "object" == typeof e && e;
			o ? "object" == typeof e && o.reset() : i.data(n, o = new a(this, s)), "string" == typeof e && o[e]()
		})
	}, t.fn.sortable.Constructor = a
}(jQuery, window, document),
/*!
 * jQuery Form Plugin
 * version: 4.2.2
 * Requires jQuery v1.7.2 or later
 * Project repository: https://github.com/jquery-form/form

 * Copyright 2017 Kevin Morris
 * Copyright 2006 M. Alsup

 * Dual licensed under the LGPL-2.1+ or MIT licenses
 * https://github.com/jquery-form/form#license

 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 */
function(t) {
	"function" == typeof define && define.amd ? define(["jquery"], t) : "object" == typeof module && module.exports ? module.exports = function(e, i) {
		return "undefined" == typeof i && (i = "undefined" != typeof window ? require("jquery") : require("jquery")(e)), t(i), i
	} : t(jQuery)
}(function(t) {
	"use strict";

	function e(e) {
		var i = e.data;
		e.isDefaultPrevented() || (e.preventDefault(), t(e.target).closest("form").ajaxSubmit(i))
	}
	function i(e) {
		var i = e.target,
			n = t(i);
		if (!n.is("[type=submit],[type=image]")) {
			var o = n.closest("[type=submit]");
			if (0 === o.length) return;
			i = o[0]
		}
		var s = i.form;
		if (s.clk = i, "image" === i.type) if ("undefined" != typeof e.offsetX) s.clk_x = e.offsetX, s.clk_y = e.offsetY;
		else if ("function" == typeof t.fn.offset) {
			var a = n.offset();
			s.clk_x = e.pageX - a.left, s.clk_y = e.pageY - a.top
		} else s.clk_x = e.pageX - i.offsetLeft, s.clk_y = e.pageY - i.offsetTop;
		setTimeout(function() {
			s.clk = s.clk_x = s.clk_y = null
		}, 100)
	}
	function n() {
		if (t.fn.ajaxSubmit.debug) {
			var e = "[jquery.form] " + Array.prototype.join.call(arguments, "");
			window.console && window.console.log ? window.console.log(e) : window.opera && window.opera.postError && window.opera.postError(e)
		}
	}
	var o = /\r?\n/g,
		s = {};
	s.fileapi = void 0 !== t('<input type="file">').get(0).files, s.formdata = "undefined" != typeof window.FormData;
	var a = !! t.fn.prop;
	t.fn.attr2 = function() {
		if (!a) return this.attr.apply(this, arguments);
		var t = this.prop.apply(this, arguments);
		return t && t.jquery || "string" == typeof t ? t : this.attr.apply(this, arguments)
	}, t.fn.ajaxSubmit = function(e, i, o, r) {
		function l(i) {
			var n, o, s = t.param(i, e.traditional).split("&"),
				a = s.length,
				r = [];
			for (n = 0; n < a; n++) s[n] = s[n].replace(/\+/g, " "), o = s[n].split("="), r.push([decodeURIComponent(o[0]), decodeURIComponent(o[1])]);
			return r
		}
		function h(i) {
			for (var n = new FormData, o = 0; o < i.length; o++) n.append(i[o].name, i[o].value);
			if (e.extraData) {
				var s = l(e.extraData);
				for (o = 0; o < s.length; o++) s[o] && n.append(s[o][0], s[o][1])
			}
			e.data = null;
			var a = t.extend(!0, {}, t.ajaxSettings, e, {
				contentType: !1,
				processData: !1,
				cache: !1,
				type: d || "POST"
			});
			e.uploadProgress && (a.xhr = function() {
				var i = t.ajaxSettings.xhr();
				return i.upload && i.upload.addEventListener("progress", function(t) {
					var i = 0,
						n = t.loaded || t.position,
						o = t.total;
					t.lengthComputable && (i = Math.ceil(n / o * 100)), e.uploadProgress(t, n, o, i)
				}, !1), i
			}), a.data = null;
			var r = a.beforeSend;
			return a.beforeSend = function(t, i) {
				e.formData ? i.data = e.formData : i.data = n, r && r.call(this, t, i)
			}, t.ajax(a)
		}
		function c(i) {
			function o(t) {
				var e = null;
				try {
					t.contentWindow && (e = t.contentWindow.document)
				} catch (i) {
					n("cannot get iframe.contentWindow document: " + i)
				}
				if (e) return e;
				try {
					e = t.contentDocument ? t.contentDocument : t.document
				} catch (i) {
					n("cannot get iframe.contentDocument: " + i), e = t.document
				}
				return e
			}
			function s() {
				function e() {
					try {
						var t = o(m).readyState;
						n("state = " + t), t && "uninitialized" === t.toLowerCase() && setTimeout(e, 50)
					} catch (i) {
						n("Server abort: ", i, " (", i.name, ")"), r(M), C && clearTimeout(C), C = void 0
					}
				}
				var i = f.attr2("target"),
					s = f.attr2("action"),
					a = "multipart/form-data",
					l = f.attr("enctype") || f.attr("encoding") || a;
				_.setAttribute("target", p), d && !/post/i.test(d) || _.setAttribute("method", "POST"), s !== c.url && _.setAttribute("action", c.url), c.skipEncodingOverride || d && !/post/i.test(d) || f.attr({
					encoding: "multipart/form-data",
					enctype: "multipart/form-data"
				}), c.timeout && (C = setTimeout(function() {
					x = !0, r(D)
				}, c.timeout));
				var h = [];
				try {
					if (c.extraData) for (var u in c.extraData) c.extraData.hasOwnProperty(u) && (t.isPlainObject(c.extraData[u]) && c.extraData[u].hasOwnProperty("name") && c.extraData[u].hasOwnProperty("value") ? h.push(t('<input type="hidden" name="' + c.extraData[u].name + '">', T).val(c.extraData[u].value).appendTo(_)[0]) : h.push(t('<input type="hidden" name="' + u + '">', T).val(c.extraData[u]).appendTo(_)[0]));
					c.iframeTarget || g.appendTo(S), m.attachEvent ? m.attachEvent("onload", r) : m.addEventListener("load", r, !1), setTimeout(e, 15);
					try {
						_.submit()
					} catch (v) {
						var y = document.createElement("form").submit;
						y.apply(_)
					}
				} finally {
					_.setAttribute("action", s), _.setAttribute("enctype", l), i ? _.setAttribute("target", i) : f.removeAttr("target"), t.each(h, function() {
						this.remove()
					})
				}
			}
			function r(e) {
				if (!v.aborted && !I) {
					if (z = o(m), z || (n("cannot access response document"), e = M), e === D && v) return v.abort("timeout"), void k.reject(v, "timeout");
					if (e === M && v) return v.abort("server abort"), void k.reject(v, "error", "server abort");
					if (z && z.location.href !== c.iframeSrc || x) {
						m.detachEvent ? m.detachEvent("onload", r) : m.removeEventListener("load", r, !1);
						var i, s = "success";
						try {
							if (x) throw "timeout";
							var a = "xml" === c.dataType || z.XMLDocument || t.isXMLDoc(z);
							if (n("isXml=" + a), !a && window.opera && (null === z.body || !z.body.innerHTML) && --$) return n("requeing onLoad callback, DOM not available"), void setTimeout(r, 250);
							var l = z.body ? z.body : z.documentElement;
							v.responseText = l ? l.innerHTML : null, v.responseXML = z.XMLDocument ? z.XMLDocument : z, a && (c.dataType = "xml"), v.getResponseHeader = function(t) {
								var e = {
									"content-type": c.dataType
								};
								return e[t.toLowerCase()]
							}, l && (v.status = Number(l.getAttribute("status")) || v.status, v.statusText = l.getAttribute("statusText") || v.statusText);
							var h = (c.dataType || "").toLowerCase(),
								d = /(json|script|text)/.test(h);
							if (d || c.textarea) {
								var p = z.getElementsByTagName("textarea")[0];
								if (p) v.responseText = p.value, v.status = Number(p.getAttribute("status")) || v.status, v.statusText = p.getAttribute("statusText") || v.statusText;
								else if (d) {
									var f = z.getElementsByTagName("pre")[0],
										y = z.getElementsByTagName("body")[0];
									f ? v.responseText = f.textContent ? f.textContent : f.innerText : y && (v.responseText = y.textContent ? y.textContent : y.innerText)
								}
							} else "xml" === h && !v.responseXML && v.responseText && (v.responseXML = E(v.responseText));
							try {
								L = j(v, h, c)
							} catch (b) {
								s = "parsererror", v.error = i = b || s
							}
						} catch (b) {
							n("error caught: ", b), s = "error", v.error = i = b || s
						}
						v.aborted && (n("upload aborted"), s = null), v.status && (s = v.status >= 200 && v.status < 300 || 304 === v.status ? "success" : "error"), "success" === s ? (c.success && c.success.call(c.context, L, "success", v), k.resolve(v.responseText, "success", v), u && t.event.trigger("ajaxSuccess", [v, c])) : s && ("undefined" == typeof i && (i = v.statusText), c.error && c.error.call(c.context, v, s, i), k.reject(v, "error", i), u && t.event.trigger("ajaxError", [v, c, i])), u && t.event.trigger("ajaxComplete", [v, c]), u && !--t.active && t.event.trigger("ajaxStop"), c.complete && c.complete.call(c.context, v, s), I = !0, c.timeout && clearTimeout(C), setTimeout(function() {
							c.iframeTarget ? g.attr("src", c.iframeSrc) : g.remove(), v.responseXML = null
						}, 100)
					}
				}
			}
			var l, h, c, u, p, g, m, v, b, w, x, C, _ = f[0],
				k = t.Deferred();
			if (k.abort = function(t) {
				v.abort(t)
			}, i) for (h = 0; h < y.length; h++) l = t(y[h]), a ? l.prop("disabled", !1) : l.removeAttr("disabled");
			c = t.extend(!0, {}, t.ajaxSettings, e), c.context = c.context || c, p = "jqFormIO" + (new Date).getTime();
			var T = _.ownerDocument,
				S = f.closest("body");
			if (c.iframeTarget ? (g = t(c.iframeTarget, T), w = g.attr2("name"), w ? p = w : g.attr2("name", p)) : (g = t('<iframe name="' + p + '" src="' + c.iframeSrc + '" />', T), g.css({
				position: "absolute",
				top: "-1000px",
				left: "-1000px"
			})), m = g[0], v = {
				aborted: 0,
				responseText: null,
				responseXML: null,
				status: 0,
				statusText: "n/a",
				getAllResponseHeaders: function() {},
				getResponseHeader: function() {},
				setRequestHeader: function() {},
				abort: function(e) {
					var i = "timeout" === e ? "timeout" : "aborted";
					n("aborting upload... " + i), this.aborted = 1;
					try {
						m.contentWindow.document.execCommand && m.contentWindow.document.execCommand("Stop")
					} catch (o) {}
					g.attr("src", c.iframeSrc), v.error = i, c.error && c.error.call(c.context, v, i, e), u && t.event.trigger("ajaxError", [v, c, i]), c.complete && c.complete.call(c.context, v, i)
				}
			}, u = c.global, u && 0 === t.active++ && t.event.trigger("ajaxStart"), u && t.event.trigger("ajaxSend", [v, c]), c.beforeSend && c.beforeSend.call(c.context, v, c) === !1) return c.global && t.active--, k.reject(), k;
			if (v.aborted) return k.reject(), k;
			b = _.clk, b && (w = b.name, w && !b.disabled && (c.extraData = c.extraData || {}, c.extraData[w] = b.value, "image" === b.type && (c.extraData[w + ".x"] = _.clk_x, c.extraData[w + ".y"] = _.clk_y)));
			var D = 1,
				M = 2,
				P = t("meta[name=csrf-token]").attr("content"),
				F = t("meta[name=csrf-param]").attr("content");
			F && P && (c.extraData = c.extraData || {}, c.extraData[F] = P), c.forceSync ? s() : setTimeout(s, 10);
			var L, z, I, $ = 50,
				E = t.parseXML ||
			function(t, e) {
				return window.ActiveXObject ? (e = new ActiveXObject("Microsoft.XMLDOM"), e.async = "false", e.loadXML(t)) : e = (new DOMParser).parseFromString(t, "text/xml"), e && e.documentElement && "parsererror" !== e.documentElement.nodeName ? e : null
			}, A = t.parseJSON ||
			function(t) {
				return window.eval("(" + t + ")")
			}, j = function(e, i, n) {
				var o = e.getResponseHeader("content-type") || "",
					s = ("xml" === i || !i) && o.indexOf("xml") >= 0,
					a = s ? e.responseXML : e.responseText;
				return s && "parsererror" === a.documentElement.nodeName && t.error && t.error("parsererror"), n && n.dataFilter && (a = n.dataFilter(a, i)), "string" == typeof a && (("json" === i || !i) && o.indexOf("json") >= 0 ? a = A(a) : ("script" === i || !i) && o.indexOf("javascript") >= 0 && t.globalEval(a)), a
			};
			return k
		}
		if (!this.length) return n("ajaxSubmit: skipping submit process - no element selected"), this;
		var d, u, p, f = this;
		"function" == typeof e ? e = {
			success: e
		} : "string" == typeof e || e === !1 && arguments.length > 0 ? (e = {
			url: e,
			data: i,
			dataType: o
		}, "function" == typeof r && (e.success = r)) : "undefined" == typeof e && (e = {}), d = e.method || e.type || this.attr2("method"), u = e.url || this.attr2("action"), p = "string" == typeof u ? t.trim(u) : "", p = p || window.location.href || "", p && (p = (p.match(/^([^#]+)/) || [])[1]), e = t.extend(!0, {
			url: p,
			success: t.ajaxSettings.success,
			type: d || t.ajaxSettings.type,
			iframeSrc: /^https/i.test(window.location.href || "") ? "javascript:false" : "about:blank"
		}, e);
		var g = {};
		if (this.trigger("form-pre-serialize", [this, e, g]), g.veto) return n("ajaxSubmit: submit vetoed via form-pre-serialize trigger"), this;
		if (e.beforeSerialize && e.beforeSerialize(this, e) === !1) return n("ajaxSubmit: submit aborted via beforeSerialize callback"), this;
		var m = e.traditional;
		"undefined" == typeof m && (m = t.ajaxSettings.traditional);
		var v, y = [],
			b = this.formToArray(e.semantic, y, e.filtering);
		if (e.data) {
			var w = t.isFunction(e.data) ? e.data(b) : e.data;
			e.extraData = w, v = t.param(w, m)
		}
		if (e.beforeSubmit && e.beforeSubmit(b, this, e) === !1) return n("ajaxSubmit: submit aborted via beforeSubmit callback"), this;
		if (this.trigger("form-submit-validate", [b, this, e, g]), g.veto) return n("ajaxSubmit: submit vetoed via form-submit-validate trigger"), this;
		var x = t.param(b, m);
		v && (x = x ? x + "&" + v : v), "GET" === e.type.toUpperCase() ? (e.url += (e.url.indexOf("?") >= 0 ? "&" : "?") + x, e.data = null) : e.data = x;
		var C = [];
		if (e.resetForm && C.push(function() {
			f.resetForm()
		}), e.clearForm && C.push(function() {
			f.clearForm(e.includeHidden)
		}), !e.dataType && e.target) {
			var _ = e.success ||
			function() {};
			C.push(function(i, n, o) {
				var s = arguments,
					a = e.replaceTarget ? "replaceWith" : "html";
				t(e.target)[a](i).each(function() {
					_.apply(this, s)
				})
			})
		} else e.success && (t.isArray(e.success) ? t.merge(C, e.success) : C.push(e.success));
		if (e.success = function(t, i, n) {
			for (var o = e.context || this, s = 0, a = C.length; s < a; s++) C[s].apply(o, [t, i, n || f, f])
		}, e.error) {
			var k = e.error;
			e.error = function(t, i, n) {
				var o = e.context || this;
				k.apply(o, [t, i, n, f])
			}
		}
		if (e.complete) {
			var T = e.complete;
			e.complete = function(t, i) {
				var n = e.context || this;
				T.apply(n, [t, i, f])
			}
		}
		var S = t("input[type=file]:enabled", this).filter(function() {
			return "" !== t(this).val()
		}),
			D = S.length > 0,
			M = "multipart/form-data",
			P = f.attr("enctype") === M || f.attr("encoding") === M,
			F = s.fileapi && s.formdata;
		n("fileAPI :" + F);
		var L, z = (D || P) && !F;
		e.iframe !== !1 && (e.iframe || z) ? e.closeKeepAlive ? t.get(e.closeKeepAlive, function() {
			L = c(b)
		}) : L = c(b) : L = (D || P) && F ? h(b) : t.ajax(e), f.removeData("jqxhr").data("jqxhr", L);
		for (var I = 0; I < y.length; I++) y[I] = null;
		return this.trigger("form-submit-notify", [this, e]), this
	}, t.fn.ajaxForm = function(o, s, a, r) {
		if (("string" == typeof o || o === !1 && arguments.length > 0) && (o = {
			url: o,
			data: s,
			dataType: a
		}, "function" == typeof r && (o.success = r)), o = o || {}, o.delegation = o.delegation && t.isFunction(t.fn.on), !o.delegation && 0 === this.length) {
			var l = {
				s: this.selector,
				c: this.context
			};
			return !t.isReady && l.s ? (n("DOM not ready, queuing ajaxForm"), t(function() {
				t(l.s, l.c).ajaxForm(o)
			}), this) : (n("terminating; zero elements found by selector" + (t.isReady ? "" : " (DOM not ready)")), this)
		}
		return o.delegation ? (t(document).off("submit.form-plugin", this.selector, e).off("click.form-plugin", this.selector, i).on("submit.form-plugin", this.selector, o, e).on("click.form-plugin", this.selector, o, i), this) : this.ajaxFormUnbind().on("submit.form-plugin", o, e).on("click.form-plugin", o, i)
	}, t.fn.ajaxFormUnbind = function() {
		return this.off("submit.form-plugin click.form-plugin")
	}, t.fn.formToArray = function(e, i, n) {
		var o = [];
		if (0 === this.length) return o;
		var a, r = this[0],
			l = this.attr("id"),
			h = e || "undefined" == typeof r.elements ? r.getElementsByTagName("*") : r.elements;
		if (h && (h = t.makeArray(h)), l && (e || /(Edge|Trident)\//.test(navigator.userAgent)) && (a = t(':input[form="' + l + '"]').get(), a.length && (h = (h || []).concat(a))), !h || !h.length) return o;
		t.isFunction(n) && (h = t.map(h, n));
		var c, d, u, p, f, g, m;
		for (c = 0, g = h.length; c < g; c++) if (f = h[c], u = f.name, u && !f.disabled) if (e && r.clk && "image" === f.type) r.clk === f && (o.push({
			name: u,
			value: t(f).val(),
			type: f.type
		}), o.push({
			name: u + ".x",
			value: r.clk_x
		}, {
			name: u + ".y",
			value: r.clk_y
		}));
		else if (p = t.fieldValue(f, !0), p && p.constructor === Array) for (i && i.push(f), d = 0, m = p.length; d < m; d++) o.push({
			name: u,
			value: p[d]
		});
		else if (s.fileapi && "file" === f.type) {
			i && i.push(f);
			var v = f.files;
			if (v.length) for (d = 0; d < v.length; d++) o.push({
				name: u,
				value: v[d],
				type: f.type
			});
			else o.push({
				name: u,
				value: "",
				type: f.type
			})
		} else null !== p && "undefined" != typeof p && (i && i.push(f), o.push({
			name: u,
			value: p,
			type: f.type,
			required: f.required
		}));
		if (!e && r.clk) {
			var y = t(r.clk),
				b = y[0];
			u = b.name, u && !b.disabled && "image" === b.type && (o.push({
				name: u,
				value: y.val()
			}), o.push({
				name: u + ".x",
				value: r.clk_x
			}, {
				name: u + ".y",
				value: r.clk_y
			}))
		}
		return o
	}, t.fn.formSerialize = function(e) {
		return t.param(this.formToArray(e))
	}, t.fn.fieldSerialize = function(e) {
		var i = [];
		return this.each(function() {
			var n = this.name;
			if (n) {
				var o = t.fieldValue(this, e);
				if (o && o.constructor === Array) for (var s = 0, a = o.length; s < a; s++) i.push({
					name: n,
					value: o[s]
				});
				else null !== o && "undefined" != typeof o && i.push({
					name: this.name,
					value: o
				})
			}
		}), t.param(i)
	}, t.fn.fieldValue = function(e) {
		for (var i = [], n = 0, o = this.length; n < o; n++) {
			var s = this[n],
				a = t.fieldValue(s, e);
			null === a || "undefined" == typeof a || a.constructor === Array && !a.length || (a.constructor === Array ? t.merge(i, a) : i.push(a))
		}
		return i
	}, t.fieldValue = function(e, i) {
		var n = e.name,
			s = e.type,
			a = e.tagName.toLowerCase();
		if ("undefined" == typeof i && (i = !0), i && (!n || e.disabled || "reset" === s || "button" === s || ("checkbox" === s || "radio" === s) && !e.checked || ("submit" === s || "image" === s) && e.form && e.form.clk !== e || "select" === a && e.selectedIndex === -1)) return null;
		if ("select" === a) {
			var r = e.selectedIndex;
			if (r < 0) return null;
			for (var l = [], h = e.options, c = "select-one" === s, d = c ? r + 1 : h.length, u = c ? r : 0; u < d; u++) {
				var p = h[u];
				if (p.selected && !p.disabled) {
					var f = p.value;
					if (f || (f = p.attributes && p.attributes.value && !p.attributes.value.specified ? p.text : p.value), c) return f;
					l.push(f)
				}
			}
			return l
		}
		return t(e).val().replace(o, "\r\n")
	}, t.fn.clearForm = function(e) {
		return this.each(function() {
			t("input,select,textarea", this).clearFields(e)
		})
	}, t.fn.clearFields = t.fn.clearInputs = function(e) {
		var i = /^(?:color|date|datetime|email|month|number|password|range|search|tel|text|time|url|week)$/i;
		return this.each(function() {
			var n = this.type,
				o = this.tagName.toLowerCase();
			i.test(n) || "textarea" === o ? this.value = "" : "checkbox" === n || "radio" === n ? this.checked = !1 : "select" === o ? this.selectedIndex = -1 : "file" === n ? /MSIE/.test(navigator.userAgent) ? t(this).replaceWith(t(this).clone(!0)) : t(this).val("") : e && (e === !0 && /hidden/.test(n) || "string" == typeof e && t(this).is(e)) && (this.value = "")
		})
	}, t.fn.resetForm = function() {
		return this.each(function() {
			var e = t(this),
				i = this.tagName.toLowerCase();
			switch (i) {
			case "input":
				this.checked = this.defaultChecked;
			case "textarea":
				return this.value = this.defaultValue, !0;
			case "option":
			case "optgroup":
				var n = e.parents("select");
				return n.length && n[0].multiple ? "option" === i ? this.selected = this.defaultSelected : e.find("option").resetForm() : n.resetForm(), !0;
			case "select":
				return e.find("option").each(function(t) {
					if (this.selected = this.defaultSelected, this.defaultSelected && !e[0].multiple) return e[0].selectedIndex = t, !1
				}), !0;
			case "label":
				var o = t(e.attr("for")),
					s = e.find("input,select,textarea");
				return o[0] && s.unshift(o[0]), s.resetForm(), !0;
			case "form":
				return ("function" == typeof this.reset || "object" == typeof this.reset && !this.reset.nodeType) && this.reset(), !0;
			default:
				return e.find("form,input,label,select,textarea").resetForm(), !0
			}
		})
	}, t.fn.enable = function(t) {
		return "undefined" == typeof t && (t = !0), this.each(function() {
			this.disabled = !t
		})
	}, t.fn.selected = function(e) {
		return "undefined" == typeof e && (e = !0), this.each(function() {
			var i = this.type;
			if ("checkbox" === i || "radio" === i) this.checked = e;
			else if ("option" === this.tagName.toLowerCase()) {
				var n = t(this).parent("select");
				e && n[0] && "select-one" === n[0].type && n.find("option").selected(!1), this.selected = e
			}
		})
	}, t.fn.ajaxSubmit.debug = !1
}),
/*!
 * jQuery Hotkeys Plugin
 * Copyright 2010, John Resig
 * Dual licensed under the MIT or GPL Version 2 licenses.
 *
 * Based upon the plugin by Tzury Bar Yochay:
 * http://github.com/tzuryby/hotkeys
 *
 * Original idea by:
 * Binny V A, http://www.openjs.com/scripts/events/keyboard_shortcuts/
 */

function(t) {
	function e(e) {
		if ("string" == typeof e.data) {
			var i = e.handler,
				n = e.data.toLowerCase().split(" ");
			e.handler = function(e) {
				if (this === e.target || !/textarea|select/i.test(e.target.nodeName) && "text" !== e.target.type) {
					var o = "keypress" !== e.type && t.hotkeys.specialKeys[e.which],
						s = String.fromCharCode(e.which).toLowerCase(),
						a = "",
						r = {};
					e.altKey && "alt" !== o && (a += "alt+"), e.ctrlKey && "ctrl" !== o && (a += "ctrl+"), e.metaKey && !e.ctrlKey && "meta" !== o && (a += "meta+"), e.shiftKey && "shift" !== o && (a += "shift+"), o ? r[a + o] = !0 : (r[a + s] = !0, r[a + t.hotkeys.shiftNums[s]] = !0, "shift+" === a && (r[t.hotkeys.shiftNums[s]] = !0));
					for (var l = 0, h = n.length; l < h; l++) if (r[n[l]]) return i.apply(this, arguments)
				}
			}
		}
	}
	t.hotkeys = {
		version: "0.8",
		specialKeys: {
			8: "backspace",
			9: "tab",
			13: "return",
			16: "shift",
			17: "ctrl",
			18: "alt",
			19: "pause",
			20: "capslock",
			27: "esc",
			32: "space",
			33: "pageup",
			34: "pagedown",
			35: "end",
			36: "home",
			37: "left",
			38: "up",
			39: "right",
			40: "down",
			45: "insert",
			46: "del",
			96: "0",
			97: "1",
			98: "2",
			99: "3",
			100: "4",
			101: "5",
			102: "6",
			103: "7",
			104: "8",
			105: "9",
			106: "*",
			107: "+",
			109: "-",
			110: ".",
			111: "/",
			112: "f1",
			113: "f2",
			114: "f3",
			115: "f4",
			116: "f5",
			117: "f6",
			118: "f7",
			119: "f8",
			120: "f9",
			121: "f10",
			122: "f11",
			123: "f12",
			144: "numlock",
			145: "scroll",
			191: "/",
			224: "meta"
		},
		shiftNums: {
			"`": "~",
			1: "!",
			2: "@",
			3: "#",
			4: "$",
			5: "%",
			6: "^",
			7: "&",
			8: "*",
			9: "(",
			0: ")",
			"-": "_",
			"=": "+",
			";": ": ",
			"'": '"',
			",": "<",
			".": ">",
			"/": "?",
			"\\": "|"
		}
	}, t.each(["keydown", "keyup", "keypress"], function() {
		t.event.special[this] = {
			add: e
		}
	})
}(jQuery), function(t) {
	"use strict";

	function e(e, i) {
		if (e === !1) return e;
		if (!e) return i;
		e === !0 ? e = {
			add: !0,
			"delete": !0,
			edit: !0,
			sort: !0
		} : "string" == typeof e && (e = e.split(","));
		var n;
		return t.isArray(e) && (n = {}, t.each(e, function(e, i) {
			t.isPlainObject(i) ? n[i.action] = i : n[i] = !0
		}), e = n), t.isPlainObject(e) && (n = {}, t.each(e, function(e, i) {
			i ? n[e] = t.extend({
				type: e
			}, a[e], t.isPlainObject(i) ? i : null) : n[e] = !1
		}), e = n), i ? t.extend(!0, {}, i, e) : e
	}
	function i(e, i, n) {
		return i = i || e.type, t(n || e.template).addClass("tree-action").attr(t.extend({
			"data-type": i,
			title: e.title || ""
		}, e.attr)).data("action", e)
	}
	var n = "zui.tree",
		o = 0,
		s = function(e, i) {
			this.name = n, this.$ = t(e), this.getOptions(i), this._init()
		},
		a = {
			sort: {
				template: '<a class="sort-handler" href="javascript:;"><i class="icon icon-move"></i></a>'
			},
			add: {
				template: '<a href="javascript:;"><i class="icon icon-plus"></i></a>'
			},
			edit: {
				template: '<a href="javascript:;"><i class="icon icon-pencil"></i></a>'
			},
			"delete": {
				template: '<a href="javascript:;"><i class="icon icon-trash"></i></a>'
			}
		};
	s.DEFAULTS = {
		animate: null,
		initialState: "normal",
		toggleTemplate: '<i class="list-toggle icon"></i>'
	}, s.prototype.add = function(e, i, n, o, s) {
		var a, r = t(e),
			l = this.options;
		if (r.is("li") ? (a = r.children("ul"), a.length || (a = t("<ul/>"), r.append(a), this._initList(a, r))) : a = r, a) {
			var h = this;
			t.isArray(i) || (i = [i]), t.each(i, function(e, i) {
				var n = t("<li/>").data(i).appendTo(a);
				void 0 !== i.id && n.attr("data-id", i.id);
				var o = l.itemWrapper ? t(l.itemWrapper === !0 ? '<div class="tree-item-wrapper"/>' : l.itemWrapper).appendTo(n) : n;
				if (i.html) o.html(i.html);
				else if (t.isFunction(h.options.itemCreator)) {
					var s = h.options.itemCreator(n, i);
					s !== !0 && s !== !1 && o.html(s)
				} else i.url ? o.append(t("<a/>", {
					href: i.url
				}).text(i.title || i.name)) : o.append(t("<span/>").text(i.title || i.name));
				h._initItem(n, i.idx || e, a, i), i.children && i.children.length && h.add(n, i.children)
			}), this._initList(a), n && !a.hasClass("tree") && h.expand(a.parent("li"), o, s)
		}
	}, s.prototype.reload = function(e) {
		var i = this;
		e && (i.$.empty(), i.add(i.$, e)), i.isPreserve && i.store.time && i.$.find("li:not(.tree-action-item)").each(function() {
			var e = t(this);
			i[i.store[e.data("id")] ? "expand" : "collapse"](e, !0, !0)
		})
	}, s.prototype._initList = function(n, o, s, a) {
		var r = this;
		n.hasClass("tree") ? (s = 0, o = null) : (o = (o || n.closest("li")).addClass("has-list"), o.find(".list-toggle").length || o.prepend(this.options.toggleTemplate), s = s || o.data("idx")), n.removeClass("has-active-item");
		var l = n.attr("data-idx", s || 0).children("li:not(.tree-action-item)").each(function(e) {
			r._initItem(t(this), e + 1, n)
		});
		1 !== l.length || l.find("ul").length || l.addClass("tree-single-item"), a = a || (o ? o.data() : null);
		var h = e(a ? a.actions : null, this.actions);
		if (h) {
			if (h.add && h.add.templateInList !== !1) {
				var c = n.children("li.tree-action-item");
				c.length ? c.detach().appendTo(n) : t('<li class="tree-action-item"/>').append(i(h.add, "add", h.add.templateInList)).appendTo(n)
			}
			h.sort && n.sortable(t.extend({
				dragCssClass: "tree-drag-holder",
				trigger: ".sort-handler",
				selector: "li:not(.tree-action-item)",
				finish: function(t) {
					r.callEvent("action", {
						action: h.sort,
						$list: n,
						target: t.target,
						item: a
					})
				}
			}, h.sort.options, t.isPlainObject(this.options.sortable) ? this.options.sortable : null))
		}
		o && (o.hasClass("open") || a && a.open) && o.addClass("open in")
	}, s.prototype._initItem = function(n, o, s, a) {
		if (void 0 === o) {
			var r = n.prev("li");
			o = r.length ? r.data("idx") + 1 : 1
		}
		if (s = s || n.closest("ul"), n.attr("data-idx", o).removeClass("tree-single-item"), !n.data("id")) {
			var l = o;
			s.hasClass("tree") || (l = s.parent("li").data("id") + "-" + l), n.attr("data-id", l)
		}
		n.hasClass("active") && s.parent("li").addClass("has-active-item"), a = a || n.data();
		var h = e(a.actions, this.actions);
		if (h) {
			var c = n.find(".tree-actions");
			c.length || (c = t('<div class="tree-actions"/>').appendTo(this.options.itemWrapper ? n.find(".tree-item-wrapper") : n), t.each(h, function(t, e) {
				e && c.append(i(e, t))
			}))
		}
		var d = n.children("ul");
		d.length && this._initList(d, n, o, a)
	}, s.prototype._init = function() {
		var i = this.options,
			s = this;
		this.actions = e(i.actions), this.$.addClass("tree"), i.animate && this.$.addClass("tree-animate"), this._initList(this.$);
		var a = i.initialState,
			r = t.zui && t.zui.store && t.zui.store.enable;
		r && (this.selector = n + "::" + (i.name || "") + "#" + (this.$.attr("id") || o++), this.store = t.zui.store[i.name ? "get" : "pageGet"](this.selector, {})), "preserve" === a && (r ? this.isPreserve = !0 : this.options.initialState = a = "normal"), this.reload(i.data), r && (this.isPreserve = !0), "expand" === a ? this.expand() : "collapse" === a && this.collapse(), this.$.on("click", '.list-toggle,a[href="#"],.tree-toggle', function(e) {
			var i = t(this),
				n = i.parent("li");
			s.callEvent("hit", {
				target: n,
				item: n.data()
			}), s.toggle(n), i.is("a") && e.preventDefault()
		}).on("click", ".tree-action", function() {
			var e = t(this),
				i = e.data();
			if (i.action && (i = i.action), "sort" !== i.type) {
				var n = e.closest("li:not(.tree-action-item)");
				s.callEvent("action", {
					action: i,
					target: this,
					$item: n,
					item: n.data()
				})
			}
		})
	}, s.prototype.preserve = function(e, i, n) {
		if (this.isPreserve) if (e) i = i || e.data("id"), n = void 0 === n && e.hasClass("open"), n ? this.store[i] = n : delete this.store[i], this.store.time = (new Date).getTime(), t.zui.store[this.options.name ? "set" : "pageSet"](this.selector, this.store);
		else {
			var o = this;
			this.store = {}, this.$.find("li").each(function() {
				o.preserve(t(this))
			})
		}
	}, s.prototype.expand = function(t, e, i) {
		t ? (t.addClass("open"), !e && this.options.animate ? setTimeout(function() {
			t.addClass("in")
		}, 10) : t.addClass("in")) : t = this.$.find("li.has-list").addClass("open in"), i || this.preserve(t), this.callEvent("expand", t, this)
	}, s.prototype.show = function(e, i, n) {
		var o = this;
		e.each(function() {
			var e = t(this);
			if (o.expand(e, i, n), e) for (var s = e.parent("ul"); s && s.length && !s.hasClass("tree");) {
				var a = s.parent("li");
				a.length ? (o.expand(a, i, n), s = a.parent("ul")) : s = !1
			}
		})
	}, s.prototype.collapse = function(t, e, i) {
		t ? !e && this.options.animate ? (t.removeClass("in"), setTimeout(function() {
			t.removeClass("open")
		}, 300)) : t.removeClass("open in") : t = this.$.find("li.has-list").removeClass("open in"), i || this.preserve(t), this.callEvent("collapse", t, this)
	}, s.prototype.toggle = function(t) {
		var e = t && t.hasClass("open") || t === !1 || void 0 === t && this.$.find("li.has-list.open").length;
		this[e ? "collapse" : "expand"](t)
	}, s.prototype.getOptions = function(e) {
		this.options = t.extend({}, s.DEFAULTS, this.$.data(), e), null === this.options.animate && this.$.hasClass("tree-animate") && (this.options.animate = !0)
	}, s.prototype.toData = function(e, i) {
		t.isFunction(e) && (i = e, e = null), e = e || this.$;
		var n = this;
		return e.children("li:not(.tree-action-item)").map(function() {
			var e = t(this),
				o = e.data();
			delete o["zui.droppable"];
			var s = e.children("ul");
			return s.length && (o.children = n.toData(s)), t.isFunction(i) ? i(o, e) : o
		}).get()
	}, s.prototype.callEvent = function(e, i) {
		var n;
		return t.isFunction(this.options[e]) && (n = this.options[e](i, this)), this.$.trigger(t.Event(e + "." + this.name, i)), n
	}, t.fn.tree = function(e, i) {
		return this.each(function() {
			var o = t(this),
				a = o.data(n),
				r = "object" == typeof e && e;
			a || o.data(n, a = new s(this, r)), "string" == typeof e && a[e](i)
		})
	}, t.fn.tree.Constructor = s, t(function() {
		t('[data-ride="tree"]').tree()
	})
}(jQuery), function(t) {
	"use strict";
	var e = "zui.colorPicker",
		i = '<div class="colorpicker"><button type="button" class="btn dropdown-toggle" data-toggle="dropdown"><span class="cp-title"></span><i class="ic"></i></button><ul class="dropdown-menu clearfix"></ul></div>',
		n = {
			zh_cn: {
				errorTip: "不是有效的颜色值"
			},
			zh_tw: {
				errorTip: "不是有效的顏色值"
			},
			en: {
				errorTip: "Not a valid color value"
			}
		},
		o = function(i, n) {
			this.name = e, this.$ = t(i), this.getOptions(n), this.init()
		};
	o.DEFAULTS = {
		colors: ["#00BCD4", "#388E3C", "#3280fc", "#3F51B5", "#9C27B0", "#795548", "#F57C00", "#F44336", "#E91E63"],
		pullMenuRight: !0,
		wrapper: "btn-wrapper",
		tileSize: 30,
		lineCount: 5,
		optional: !0,
		tooltip: "top",
		icon: "caret-down",
		updateBtn: "auto"
	}, o.prototype.init = function() {
		var e = this,
			n = e.options,
			o = e.$,
			s = o.parent(),
			a = !1;
		s.hasClass("colorpicker") ? e.$picker = s : (e.$picker = t(n.template || i), a = !0), e.$picker.addClass(n.wrapper).find(".cp-title").toggle(void 0 !== n.title).text(n.title), e.$menu = e.$picker.find(".dropdown-menu").toggleClass("pull-right", n.pullMenuRight), e.$btn = e.$picker.find(".btn.dropdown-toggle"), e.$btn.find(".ic").addClass("icon-" + n.icon), n.btnTip && e.$picker.attr("data-toggle", "tooltip").tooltip({
			title: n.btnTip,
			placement: n.tooltip,
			container: "body"
		}), o.attr("data-provide", null), a && o.after(e.$picker), e.colors = {}, t.each(n.colors, function(i, n) {
			if (t.zui.Color.isColor(n)) {
				var o = new t.zui.Color(n);
				e.colors[o.toCssStr()] = o
			}
		}), e.updateColors(), e.$picker.on("click", ".cp-tile", function() {
			e.setValue(t(this).data("color"))
		});
		var r = function() {
				var i = o.val(),
					s = t.zui.Color.isColor(i);
				o.parent().toggleClass("has-error", !(s || n.optional && "" === i)), s ? e.setValue(i, !0) : n.optional && "" === i ? o.tooltip("hide") : o.is(":focus") || o.tooltip("show", n.errorTip)
			};
		o.is("input:not([type=hidden])") ? (n.tooltip && o.attr("data-toggle", "tooltip").tooltip({
			trigger: "manual",
			placement: n.tooltip,
			tipClass: "tooltip-danger",
			container: "body"
		}), o.on("keyup paste input change", r)) : o.appendTo(e.$picker), r()
	}, o.prototype.addColor = function(e) {
		e instanceof t.zui.Color || (e = new t.zui.Color(e));
		var i = e.toCssStr(),
			n = this.options;
		this.colors[i] || (this.colors[i] = e);
		var o = t('<a href="###" class="cp-tile"></a>', {
			titile: e
		}).data("color", e).css({
			color: e.contrast().toCssStr(),
			background: i,
			"border-color": e.luma() > .43 ? "#ccc" : "transparent"
		}).attr("data-color", i);
		this.$menu.append(t("<li/>").css({
			width: n.tileSize,
			height: n.tileSize
		}).append(o)), n.optional && this.$menu.find(".cp-tile.empty").parent().detach().appendTo(this.$menu)
	}, o.prototype.updateColors = function(e) {
		var i = (this.$picker, this.$menu.children("li:not(.heading)").remove()),
			n = this.options,
			e = e || this.colors,
			o = this,
			s = 0;
		if (t.each(e, function(t, e) {
			o.addColor(e), s++
		}), n.optional) {
			var a = t('<li><a class="cp-tile empty" href="###"></a></li>').css({
				width: n.tileSize,
				height: n.tileSize
			});
			this.$menu.append(a), s++
		}
		i.css("width", Math.min(s, n.lineCount) * n.tileSize + 6)
	}, o.prototype.setValue = function(e, i) {
		var n = this,
			o = n.options,
			s = n.$btn;
		n.$menu.find(".cp-tile.active").removeClass("active");
		var a = "",
			r = o.updateBtn;
		if ("auto" === r) {
			var l = s.find(".color-bar");
			r = !l.length ||
			function(t) {
				l.css("background", t || "")
			}
		}
		if (e) {
			var h = new t.zui.Color(e);
			a = h.toCssStr().toLowerCase(), r && (t.isFunction(r) ? r(a, s, n) : s.css({
				background: a,
				color: h.contrast().toCssStr(),
				borderColor: h.luma() > .43 ? "#ccc" : a
			})), n.colors[a] || n.addColor(h), i || n.$.val().toLowerCase() === a || n.$.val(a).trigger("change"), n.$menu.find('.cp-tile[data-color="' + a + '"]').addClass("active"), n.$.tooltip("hide"), n.$.trigger("colorchange", h)
		} else r && (t.isFunction(r) ? r(null, s, n) : s.attr("style", null)), i || "" === n.$.val() || n.$.val(a).trigger("change"), o.optional && n.$.tooltip("hide"), n.$menu.find(".cp-tile.empty").addClass("active"), n.$.trigger("colorchange", null);
		o.updateBorder && t(o.updateBorder).css("border-color", a), o.updateBackground && t(o.updateBackground).css("background-color", a), o.updateColor && t(o.updateColor).css("color", a), o.updateText && t(o.updateText).text(a)
	}, o.prototype.getOptions = function(e) {
		var i = t.extend({}, o.DEFAULTS, this.$.data(), e);
		"string" == typeof i.colors && (i.colors = i.colors.split(","));
		var s = (i.lang || t.zui.clientLang()).toLowerCase();
		i.errorTip || (i.errorTip = n[s].errorTip), t.fn.tooltip || (i.btnTip = !1), this.options = i
	}, t.fn.colorPicker = function(i) {
		return this.each(function() {
			var n = t(this),
				s = n.data(e),
				a = "object" == typeof i && i;
			s || n.data(e, s = new o(this, a)), "string" == typeof i && s[i]()
		})
	}, t.fn.colorPicker.Constructor = o, t(function() {
		t('[data-provide="colorpicker"]').colorPicker()
	})
}(jQuery), function(t, e) {
	function i(t) {
		return t === e && (t = n += 1), o[t % o.length]
	}
	var n = 0,
		o = ["#00a9fc", "#ff5d5d", "#fdc137", "#00da88", "#7ec5ff", "#8666b8", "#bd7b46", "#ff9100", "#ff3d00", "#f57f17", "#00e5ff", "#00b0ff", "#2979ff", "#3d5afe", "#651fff", "#d500f9", "#f50057", "#ff1744"];
	jQuery.fn.tableChart = function() {
		t(this).each(function() {
			var e = t(this),
				n = e.data(),
				o = n.chart || "pie",
				s = t(n.target);
			if (s.length) {
				var a = null;
				if ("pie" === o) {
					n = t.extend({
						scaleShowLabels: !0,
						scaleLabel: "<%=label%>: <%=value%>"
					}, n);
					var r = [],
						l = e.find("tbody > tr").each(function(e) {
							var n = t(this),
								o = i();
							n.attr("data-id", e).find(".chart-color-dot").css("background", o), r.push({
								label: n.find(".chart-label").text(),
								value: parseInt(n.find(".chart-value").text()),
								color: o,
								id: e
							})
						});
					r.length > 1 ? n.scaleLabelPlacement = "outside" : 1 === r.length && (n.scaleLabelPlacement = "inside", r.push({
						label: "",
						value: r[0].value / 2e3,
						color: "#fff",
						showLabel: !1
					})), a = s.pieChart(r, n), s.on("mousemove", function(t) {
						var e = a.getSegmentsAtEvent(t);
						l.removeClass("active"), e.length && l.filter('[data-id="' + e[0].id + '"]').addClass("active")
					})
				} else if ("bar" === o) {
					var h = i(),
						c = [],
						d = {
							label: e.find("thead .chart-label").text(),
							color: h,
							data: []
						},
						l = e.find("tbody > tr").each(function(e) {
							var i = t(this);
							c.push(i.find(".chart-label").text()), d.data.push(parseInt(i.find(".chart-value").text())), i.find(".chart-color-dot").css("background", h)
						}),
						r = {
							labels: c,
							datasets: [d]
						};
					c.length && (n.barValueSpacing = 5), a = s.barChart(r, n)
				} else if ("line" === o) {
					var h = i(),
						c = [],
						d = {
							label: e.find("thead .chart-label").text(),
							color: h,
							data: []
						},
						l = e.find("tbody > tr").each(function(e) {
							var i = t(this);
							c.push(i.find(".chart-label").text()), d.data.push(parseInt(i.find(".chart-value").text())), i.find(".chart-color-dot").css("background", h)
						}),
						r = {
							labels: c,
							datasets: [d]
						};
					c.length && (n.barValueSpacing = 5), a = s.lineChart(r, n)
				}
				null !== a && e.data("zui.chart", a)
			}
		})
	}, t(".table-chart").tableChart();
	var s = function(i, n) {
			var o = t(i);
			if (!o.data("initProgressPie")) {
				o.data("initProgressPie", 1);
				var s = o.is("canvas") ? o : o.find("canvas"),
					a = t.extend({
						value: 0,
						color: t.getThemeColor("primary") || "#006af1",
						backColor: t.getThemeColor("pale") || "#E9F2FB",
						doughnut: !0,
						doughnutSize: 85,
						width: 20,
						height: 20,
						showTip: !1,
						name: "",
						tipTemplate: "<%=value%>%",
						animation: "auto",
						realValue: parseFloat(o.find(".progress-value").text())
					}, n, o.data()),
					r = s.length;
				r || (s = t("<canvas>").appendTo(o)), s.attr("width") !== e ? a.width = s.attr("width") : s.attr("width", a.width), s.attr("height") !== e ? a.height = s.attr("height") : s.attr("height", a.height), r || 8 != t.zui.browser.ie || G_vmlCanvasManager.initElement(s[0]), "auto" === a.animation && (a.animation = a.width > 30), a.value = Math.max(0, Math.min(100, a.value)), o.addClass("progress-pie-" + a.width);
				var l = [{
					value: a.value,
					label: a.name,
					color: a.color,
					circleBeginEnd: !0
				}, {
					value: 100 - a.value,
					label: "",
					color: a.backColor
				}];
				s[a.doughnut ? "doughnutChart" : "pieChart"](l, t.extend({
					segmentShowStroke: !1,
					animation: a.animation,
					showTooltips: a.showTip,
					tooltipTemplate: a.tipTemplate,
					percentageInnerCutout: a.doughnutSize,
					reverseDrawOrder: !0,
					animationEasing: "easeInOutQuart",
					onAnimationProgress: a.realValue ?
					function(t) {
						o.find(".progress-value").text(Math.floor(a.realValue * t))
					} : e,
					onAnimationComplete: a.realValue ?
					function(t) {
						o.find(".progress-value").text(a.realValue)
					} : e
				}, a.chartOptions))
			}
		};
	jQuery.fn.progressPie = function(e) {
		t(this).each(function() {
			var i = t(this);
			if (!i.closest(".hidden").length) {
				var n = i.closest(".tab-pane");
				n.length && !n.hasClass("active") ? t('[data-toggle="tab"][data-target="#' + n.attr("id") + '"]').one("shown.zui.tab", function() {
					s(i, e)
				}) : s(this, e)
			}
		})
	}, t(function() {
		t(".table-chart").tableChart();
		var e = t(".progress-pie");
		e.length > 100 ? setTimeout(function() {
			e.progressPie()
		}, 1e3) : e.progressPie()
	})
}(jQuery, void 0), function(t) {
	jQuery.fn.sparkline = function(e) {
		t(this).each(function() {
			var i = t(this),
				n = t.extend({
					values: i.attr("values"),
					width: i.width() - 4,
					height: i.height() - 4
				}, i.data(), e),
				o = n.height,
				s = [],
				a = n.width,
				r = n.values.split(","),
				l = 0;
			for (var h in r) {
				var c = parseFloat(r[h]);
				NaN != c && (s.push(c), l = Math.max(c, l))
			}
			var d = (Math.min(l, 30), Math.min(a, Math.max(10, s.length * a / 30))),
				u = i.children("canvas");
			u.length || (i.append('<canvas class="projectline-canvas"></canvas>'), u = i.children("canvas")), u.attr("width", d).attr("height", o);
			var p = {
				labels: s,
				datasets: [{
					fillColor: t.getThemeColor("pale") || "rgba(0,0,255,0.05)",
					strokeColor: t.getThemeColor("primary") || "#0054EC",
					pointColor: t.getThemeColor("secondary") || "rgba(255,136,0,1)",
					pointStrokeColor: "#fff",
					data: s
				}]
			},
				f = {
					animation: !0,
					scaleOverride: !0,
					scaleStepWidth: Math.ceil(l / 10),
					scaleSteps: 10,
					scaleStartValue: 0,
					showScale: !1,
					showTooltips: !1,
					pointDot: !1,
					scaleShowGridLines: !1,
					datasetStrokeWidth: 1
				},
				g = t(u).lineChart(p, f);
			i.data("sparklineChart", g)
		})
	}, t(function() {
		t(".sparkline").sparkline()
	})
}(jQuery), function(t) {
	t(function() {
		t.fn.fixedDate = function() {
			return t(this).each(function() {
				var e = t(this).attr("autocomplete", "off");
				"0000-00-00" == e.val() && e.focus(function() {
					"0000-00-00" == e.val() && e.val("").datetimepicker("update")
				}).blur(function() {
					"" == e.val() && e.val("0000-00-00")
				})
			})
		};
		var e = {
			language: t("html").attr("lang"),
			weekStart: 1,
			todayBtn: 1,
			autoclose: 1,
			todayHighlight: 1,
			startView: 2,
			forceParse: 0,
			showMeridian: 1,
			format: "yyyy-mm-dd hh:ii",
			startDate: "1970-1-1"
		},
			i = t.extend({}, e, {
				minView: 2,
				format: "yyyy-mm-dd"
			}),
			n = t.extend({}, e, {
				startView: 1,
				minView: 0,
				maxView: 1,
				format: "hh:ii"
			});
		t(".datepicker-wrapper").click(function() {
			t(this).find(".form-date, .form-datetime, .form-time").datetimepicker("show").focus()
		}), window.datepickerOptions = e, t.fn.datepicker = function(e) {
			return this.datetimepicker(t.extend({}, i, e))
		}, t.fn.timepicker = function(e) {
			return this.datetimepicker(t.extend({}, n, e))
		}, t.fn.datepickerAll = function() {
			return this.find(".form-datetime").fixedDate().datetimepicker(e), this.find(".form-date").fixedDate().datepicker(), this.find(".form-time").fixedDate().timepicker(), this
		}, t("body").datepickerAll()
	})
}(jQuery), function(t) {
	var e = function(e, i) {
			i = t.extend({
				idStart: 0,
				idEnd: 9,
				chosen: !0,
				datetimepicker: !0,
				colorPicker: !0,
				hotkeys: !0
			}, i, e.data());
			var n = e.find(".template");
			!n.length && i.template && (n = t(i.template));
			var o = 0,
				s = 0,
				a = function(t) {
					t.is("select.chosen") ? t.next(".chosen-container").find("input").focus() : t.focus()
				},
				r = function(t) {
					var i = e.find("[data-ctrl-index]:focus,.chosen-container-active").first();
					if (i.length) {
						if (i.is(".chosen-container-active")) {
							if (i.hasClass("chosen-with-drop") && ("down" === t || "up" === t)) return;
							i = i.prev("select.chosen")
						}
						var n = i.data("ctrlIndex"),
							r = i.closest("tr").data("row");
						"down" === t ? r < s - 1 ? r += 1 : r = 0 : "up" === t ? r > 0 ? r -= 1 : r = s - 1 : "left" === t ? n > 0 ? n -= 1 : n = o - 1 : "right" === t && (n < o - 1 ? n += 1 : n = 0), a(e.find('tr[data-row="' + r + '"]').find('[data-ctrl-index="' + n + '"]'))
					}
				},
				l = {
					options: i,
					focusNext: r,
					focusControl: a
				},
				h = e.find("tbody,.batch-rows"),
				c = function(e) {
					t.fn.chosen && i.chosen && e.find(".chosen").chosen(t.isPlainObject(i.chosen) ? i.chosen : null), t.fn.datetimepicker && i.datetimepicker && e.datepickerAll(t.isPlainObject(i.datetimepicker) ? i.datetimepicker : null), t.fn.colorPicker && i.colorPicker && e.find("input.colorpicker").colorPicker(t.isPlainObject(i.colorPicker) ? i.colorPicker : null);
					var n = 0;
					e.find('input[type!="hidden"],textarea,select').each(function() {
						var e = t(this);
						e.parent().hasClass("chosen-search") || e.attr("data-ctrl-index", n++)
					}), o = Math.max(o, n)
				};
			if (n.length) {
				var d = n.remove().html(),
					u = function(e, o) {
						var a = d;
						"number" != typeof e && (e = s), s = Math.max(e + 1, s), a = a.replace(/\$idPlus/g, e + 1).replace(/\$id/g, e);
						var r = t("<" + n[0].tagName.toLowerCase() + " />").html(a);
						return r.attr("data-row", e).addClass(n.attr("class")).removeClass("template"), i.rowCreator && i.rowCreator(r, e, i), o ? o.after(r) : h.append(r), c(r), r
					};
				t.extend(l, {
					createRow: u,
					template: d
				});
				for (var p = i.idStart; p <= i.idEnd; ++p) u(p)
			} else c(e);
			e.on("click", ".btn-copy", function() {
				var e = t(this),
					i = t(e.data("copyFrom")).val(),
					n = t(e.data("copyTo")).val(i).addClass("highlight");
				setTimeout(function() {
					n.removeClass("highlight")
				}, 2e3)
			}), i.hotkeys && t(document).on("keydown", function(t) {
				var e = {
					"Ctrl+#37": "left",
					"Ctrl+#39": "right",
					"#38": "up",
					"#40": "down",
					"Ctrl+#38": "up",
					"Ctrl+#40": "down"
				},
					i = [];
				t.ctrlKey && i.push("Ctrl"), i.push("#" + t.keyCode);
				var n = e[i.join("+")];
				n && (r(n), t.ctrlKey && (t.stopPropagation(), t.preventDefault()))
			}), e.data("zui.batchActionForm", l)
		};
	t.fn.batchActionForm = function(i) {
		return this.each(function() {
			e(t(this), i)
		})
	}
}(jQuery), function(t, e) {
	"use strict";
	var i = "zui.table",
		n = {
			zh_cn: {
				selectedItems: "已选择 <strong>{0}</strong> 项",
				attrTotal: "{0}总计 <strong>{1}</strong>"
			},
			zh_tw: {
				selectedItems: "已选择 <strong>{0}</strong> 项",
				attrTotal: "{0}总计 <strong>{1}</strong>"
			},
			en: {
				selectedItems: "Seleted <strong>{0}</strong> items",
				attrTotal: "{0} total <strong>{1}</strong>"
			}
		},
		o = /^((?!chrome|android).)*safari/i.test(navigator.userAgent),
		s = function(e, o) {
			var a = this;
			a.name = i;
			var r = a.$ = t(e);
			o = a.options = t.extend({}, s.DEFAULTS, this.$.data(), o);
			var l = o.lang || "zh_cn";
			a.lang = t.isPlainObject(l) ? t.extend(!0, {}, n[l.lang || t.zui.clientLang()], l) : n[l], r.attr("id") || (r.attr("id", "table-" + t.zui.uuid()), o.hot && console.warn("ZUI: table hot replace id not defined, the element id attribute should be set.")), r.attr("data-ride") || r.attr("data-ride", "table"), a.getTable().find("thead>tr>th").each(function() {
				var e = t(this);
				if (!e.attr("title")) {
					var i = t.trim(e.find("a").text() || e.text() || "");
					i.length && e.attr("title", i)
				}
			}), o.checkable && (r.on("click", ".check-all", function() {
				a.checkAll(!t(this).hasClass("checked"))
			}).on("click", "tbody>tr", function(e) {
				t(e.target).closest('.btn,a,.not-check,.form-control,input[type="text"],.chosen-container').length || a.checkRow(t(this))
			}).on("click", 'tbody input[type="checkbox"],tbody label[for]', function(e) {
				e.stopPropagation();
				var i = t(this);
				i.is("label") && (i = i.closest(".checkbox-primary").find('input[type="checkbox"]')), a.checkRow(i.closest("tr"), i.is(":checked"))
			}), o.selectable && r.selectable(t.extend({}, {
				selector: a.isDataTable ? ".fixed-left tbody>tr" : "tbody>tr",
				selectClass: "",
				trigger: ".c-id",
				clickBehavior: "multi",
				listenClick: !1,
				select: function(e) {
					a.checkRow(e.target, !0), t.cookie("ajax_dragSelected") || (t.cookie("ajax_dragSelected", "on", {
						expires: config.cookieLife,
						path: config.webRoot
					}), t.ajaxSendScore("dragSelected"))
				},
				unselect: function(t) {
					a.checkRow(t.target, !1)
				},
				rangeStyle: {
					border: "1px solid #006af1",
					backgroundColor: "rgba(50,128,252,0.2)",
					borderRadius: "2px"
				}
			}, t.isPlainObject(o.selectable) ? o.selectable : null)));
			var h = a.$form = r.is("form") ? r : r.find("form");
			h.length && (o.ajaxForm ? h.ajaxForm(t.isPlainObject(o.ajaxForm) ? o.ajaxForm : null) : h.on("click", "[data-form-action]", function() {
				h.attr("action", t(this).data("formAction")).submit()
			})), (o.fixFooter || o.fixHeader) && (a.pageFooterHeight = t("#footer").outerHeight(), a.updateFixUI(), t(window).on("scroll resize", function() {
				a.updateFixUI()
			}).on("sidebar.toggle", function() {
				setTimeout(function() {
					a.updateFixUI()
				}, 200)
			})), o.group && (r.on("click", ".group-toggle", function() {
				a.toggleRowGroup(t(this).closest("tr").data("id"))
			}), t(document).on("click", ".group-collapse-all", function() {
				a.toggleGroups(!1)
			}).on("click", ".group-expand-all", function() {
				a.toggleGroups(!0)
			})), a.defaultStatistic = r.find(".table-statistic").html(), a.updateStatistic(), a.initModals(), a.checkItems = {}, a.updateCheckUI()
		};
	s.prototype.reload = function(e) {
		var i = this,
			n = i.options,
			o = n.replaceId;
		if (!o) return e && e();
		"self" === o && (o = i.$.attr("id"));
		var s = t("<div></div>");
		i.$.addClass("load-indicator loading"), s.load(window.location.href + " #" + o, function() {
			i.$.empty().html(s.children().html()).removeClass("load-indicator loading"), i.$.trigger("beforeTableReload"), i.updateStatistic(), i.initModals(), i.$.datepickerAll();
			var o = i.$.find("tbody>tr"),
				a = !1;
			t.each(i.checkItems, function(t, e) {
				e && (i.checkRow(o.filter('[data-id="' + t + '"]'), !0, !0), a = !0)
			}), a && i.updateCheckUI(), i.$.trigger("tableReload");
			var r = t("#mainMenu>.btn-toolbar>.btn-active-text>.label");
			r.length && r.text(i.getTable().find("tbody:first>tr:not(.table-children)").length), i.$.find('[data-ride="pager"]').pager(), e && e(), n.afterReload && n.afterReload()
		})
	}, s.prototype.initModals = function() {
		var e = this,
			i = e.options,
			n = e.$.find(i.iframeModalTrigger);
		if (n.length) {
			var o = {
				type: "iframe",
				onHide: i.replaceId ?
				function() {
					var n = t.cookie("selfClose");
					(1 == n || i.hot) && (t("#triggerModal").data("cancel-reload", 1), e.reload(function() {
						t.cookie("selfClose", 0)
					}))
				} : null
			};
			n.modalTrigger(o)
		}
	}, s.prototype.getTable = function() {
		var t = this.$;
		if (this.isDataTable) return t.find("div.datatable");
		var e = t.is("table") ? t : t.find("table:not(.fixed-header-copy)").first();
		return e.is(".datatable") && (this.isDataTable = !0, e = t.find("div.datatable")), e
	}, s.prototype.toggleGroups = function(e) {
		var i = this,
			n = {};
		i.$.find("tbody>tr").each(function() {
			var o = t(this).closest("tr").data("id");
			n[o] || i.toggleRowGroup(o, e)
		})
	}, s.prototype.toggleRowGroup = function(i, n) {
		var o = this.$.find('tbody>tr[data-id="' + i + '"]'),
			s = o.filter(".group-summary"),
			a = n === e ? !s.hasClass("hidden") : !! n;
		o.not(".group-summary").toggleClass("hidden", !a), s.toggleClass("hidden", a), t("body").toggleClass("table-group-collapsed", !this.$.find("tbody>tr.group-summary.hidden").length)
	}, s.prototype.updateStatistic = function() {
		var i = this,
			n = i.$.find(".table-statistic");
		if (n.length) {
			if (i.defaultStatistic === e && (i.defaultStatistic = n.html()), i.options.statisticCreator) return void n.html(i.options.statisticCreator(i) || i.defaultStatistic);
			var o = i.statisticCols;
			if (!o && o !== !1) {
				o = {};
				var s = !1;
				i.getTable().find("thead th").each(function(e) {
					var i = t(this),
						n = i.data("statistic");
					n && (s = !0, o[e] = {
						format: n,
						name: i.text()
					})
				}), i.statisticCols = !! s && o
			}
			var a = 0;
			o && t.each(o, function(t) {
				o[t].total = 0, o[t].checkedTotal = 0
			}), i.$.find(i.isDataTable ? ".fixed-left tbody>tr" : "tbody>tr").each(function() {
				var e = t(this),
					i = e.hasClass("checked"),
					n = e.children("td");
				i && a++, o && t.each(o, function(t) {
					var e = parseFloat(n.eq(t).text());
					isNaN(e) && (e = 0), o[t].total += e, i && (o[t].checkedTotal += e)
				})
			});
			var r = [];
			if (a) r.push(i.lang.selectedItems.format(a));
			else if (i.defaultStatistic) return void n.html(i.defaultStatistic);
			o && t.each(o, function(t) {
				var e = o[t],
					n = e[a ? "checkedTotal" : "total"];
				e.format && (n = e.format.format(n)), r.push(i.lang.attrTotal.format(e.name, n))
			}), n.html(r.join(", "))
		}
	}, s.prototype.updateFixUI = function(e) {
		var i = this,
			n = (new Date).getTime();
		if (!e && (i.lastUpdateCall && clearTimeout(i.lastUpdateCall), !i.lastUpdateTime || n - i.lastUpdateTime < 100)) return void(i.lastUpdateCall = setTimeout(function() {
			i.updateFixUI(!0)
		}, 30));
		if (i.lastUpdateTime = n, i.lastUpdateCall && (clearTimeout(i.lastUpdateCall), i.lastUpdateCall = null), o) {
			var s = i.getTable();
			if (s.parent().is(".table-responsive")) {
				var a = s.find("thead"),
					r = 0;
				a.find("th").each(function() {
					r += t(this).outerWidth()
				}), s.css("min-width", r)
			}
		}
		i.options.fixHeader && !i.isDataTable && i.fixHeader(), i.options.fixFooter && i.fixFooter()
	}, s.prototype.fixHeader = function() {
		var e = this,
			i = e.getTable(),
			n = i.find("thead"),
			o = n[0].getBoundingClientRect(),
			s = e.options.fixFooter,
			a = t.isFunction(s) ? s(o, n) : o.top < ("number" == typeof s ? s : -5),
			r = e.$.find(".fix-table-copy-wrapper"),
			l = i.parent(),
			h = l.is(".table-responsive");
		if (a) {
			if (r.length || (r = t('<div class="fix-table-copy-wrapper" style="overflow: hidden; position:fixed; z-index: 3; top: 0;"></div>').append(t('<table class="fixed-header-copy"></table>').addClass(i.attr("class")).append(n.clone())).insertAfter(i)), h) {
				var c = l[0].getBoundingClientRect();
				r.css({
					left: c.left,
					width: l.width()
				}), r.find(".fixed-header-copy").css({
					left: o.left - c.left,
					position: "relative",
					minWidth: i.width()
				}), l.data("fixHeaderScroll") || l.data("fixHeaderScroll", 1).on("scroll", function() {
					e.fixHeader()
				})
			} else r.css({
				left: o.left,
				width: o.width
			});
			var d = r.find("th");
			n.find("th").each(function(e) {
				d.eq(e).css("width", t(this).outerWidth())
			})
		} else r.remove()
	}, s.prototype.fixFooter = function() {
		var e, i = this,
			n = i.getTable(),
			o = i.$.find(".table-footer");
		if (i.isDataTable) e = n[0].getBoundingClientRect();
		else {
			var s = n.find("tbody");
			if (!s.length) return;
			e = s[0].getBoundingClientRect(), e = s[0].getBoundingClientRect()
		}
		var a = i.options.fixFooter;
		o.toggleClass("fixed-footer", !! r);
		var r = t.isFunction(a) ? a(e, o) : e.bottom > window.innerHeight - 50 - ("number" == typeof a ? a : i.pageFooterHeight || 5);
		o.toggleClass("fixed-footer", !! r), n.toggleClass("with-footer-fixed", !! r), n.trigger("fixFooter", r);
		var l = t("body"),
			h = l.hasClass("body-modal");
		if (r) {
			var c = n.parent(),
				d = c.is(".table-responsive");
			o.css({
				bottom: i.pageFooterHeight || 0,
				left: d ? c[0].getBoundingClientRect().left : e.left,
				width: d ? c.width() : e.width
			}), h && l.css("padding-bottom", 40)
		} else o.css({
			width: "",
			left: 0,
			bottom: 0
		}), h && l.css("padding-bottom", 0)
	}, s.prototype.checkAll = function(e) {
		var i = this,
			n = i.$.find(i.isDataTable ? ".fixed-left tbody>tr" : "tbody>tr");
		n.each(function() {
			i.checkRow(t(this), e, !0)
		}), i.updateCheckUI()
	}, s.prototype.checkRow = function(t, i, n) {
		var o = this;
		o.isDataTable && !t.is(".datatable-row-left") && (t = o.getTable().find('.datatable-row-left[data-index="' + t.data("index") + '"]'));
		var s = t.find('input[type="checkbox"]');
		s.length && (i === e && (i = !s.is(":checked")), o.isDataTable ? o.getTable().find('.datatable-row[data-index="' + t.data("index") + '"]').toggleClass("checked", i) : t.toggleClass("checked", i), this.checkItems[t.data("id")] = i, s.prop("checked", i).trigger("change"), n || o.updateCheckUI())
	}, s.prototype.updateCheckUI = function() {
		var e = this,
			i = e.getTable(),
			n = i.find(e.isDataTable ? ".fixed-left tbody>tr" : "tbody>tr").not(".group-summary"),
			o = !1,
			s = null,
			a = 0,
			r = !1,
			l = n.length;
		n.each(function(n) {
			var h = t(this),
				c = h.find('input[type="checkbox"]');
			r = c.is(":checked");
			var d = e.isDataTable ? i.find('.datatable-row[data-index="' + h.data("index") + '"]') : h;
			d.toggleClass("checked", r), d.toggleClass("row-check-begin", r && !o), s && s.toggleClass("row-check-end", !r && o), r && (a += 1), s = d, o = r, l === n + 1 && d.toggleClass("row-check-end", r)
		}), e.$.toggleClass("has-row-checked", a > 0).find(".check-all").toggleClass("checked", !(!l || a !== l)), e.updateStatistic(), e.options.onCheckChange && e.options.onCheckChange()
	}, s.DEFAULTS = {
		checkable: !0,
		ajaxForm: !1,
		selectable: !0,
		fixHeader: !0,
		fixFooter: !0,
		iframeWidth: 900,
		replaceId: "self",
		hot: !1,
		iframeModalTrigger: ".iframe"
	}, t.fn.table = function(e) {
		return this.each(function() {
			var n = t(this),
				o = n.data(i),
				a = "object" == typeof e && e;
			o || n.data(i, o = new s(this, a)), "string" == typeof e && o[e]()
		})
	}, s.NAME = i, t.fn.table.Constructor = s, t(function() {
		t('[data-ride="table"]').table()
	})
}(jQuery, void 0), function(t, e, i) {
	t.fn._ajaxForm = t.fn.ajaxForm;
	var n = {
		timeout: e.config ? e.config.timeout : 0,
		dataType: "json",
		method: "post"
	},
		o = "";
	t.fn.enableForm = function(e, n, o) {
		return e === i && (e = !0), this.each(function() {
			var i = t(this);
			n || i.find('[type="submit"]').attr("disabled", e ? null : "disabled"), !o && i.hasClass("load-indicator") && i.toggleClass("loading", !e), i.toggleClass("form-disabled", !e)
		})
	}, t.enableForm = function(e, i, n) {
		e === !1 ? t("form").enableForm(e, i, n) : t("form.form-disabled").enableForm(!0, i, n)
	};
	var s = function(e, i, n) {
			"string" == typeof i && (n = i, i = null), n = n || "show", t.zui.messager ? t.zui.messager[n](e, i) : alert(e)
		};
	t.ajaxForm = function(a, r) {
		var l = t(a);
		if (l.length > 1) return l.each(function() {
			t.ajaxForm(this, r)
		});
		t.isFunction(r) && (r = {
			complete: r
		}), r = t.extend({}, n, l.data(), r);
		var h = r.beforeSubmit,
			c = r.error,
			d = r.success,
			u = r.finish;
		delete r.finish, delete r.success, delete r.onError, delete r.beforeSubmit, r = t.extend({
			beforeSubmit: function(n, s, a) {
				if (l.enableForm(!1), (h && h(n, s, a)) !== !1) {
					var r = {},
						c = s.find('[type="file"]');
					r.fileapi = c.length && c[0].files !== i, r.formdata = e.FormData !== i;
					var d = r.fileapi && s.find('input[type="file"]:enabled').filter(function() {
						return "" !== t(this).val()
					}),
						u = d.length,
						p = "multipart/form-data",
						f = s.attr("enctype") == p || s.attr("encoding") == p,
						g = r.fileapi && r.formdata,
						m = (u || f) && !g;
					m && ("" == o && (o = a.url), a.url != o && (a.url = o), a.url = a.url.indexOf("&") >= 0 ? a.url + "&HTTP_X_REQUESTED_WITH=XMLHttpRequest" : a.url + "?HTTP_X_REQUESTED_WITH=XMLHttpRequest");
				}
			},
			success: function(i, n, o) {
				if ((d && d(i, n, o, l)) !== !1) {
					try {
						"string" == typeof i && (i = JSON.parse(i))
					} catch (a) {}
					if (null === i || "object" != typeof i) return i ? alert(i) : s("No response.", "danger");
					var h = r.responser ? t(r.responser) : l.find(".form-responser");
					h.length || (h = t("#responser"));
					var c = i.message,
						p = function() {
							var n = i.callback;
							if (n) {
								var o = n.indexOf("("),
									s = (o > 0 ? n.substr(0, o) : n).split("."),
									a = e,
									r = s[0];
								s.length > 1 && (r = s[1], "top" === s[0] ? a = e.top : "parent" === s[0] && (a = e.parent));
								var h = a[r];
								if (t.isFunction(h)) {
									var c = [];
									return o > 0 && ")" == n[n.length - 1] && (c = t.parseJSON("[" + n.substring(o + 1, n.length - 1) + "]")), c.push(i), h.apply(l, c)
								}
							}
						};
					if ("success" === i.result) {
						if (l.enableForm(!0, 1), c) {
							var f = l.find('[type="submit"]'),
								g = !1;
							f.length && (f.popover({
								container: "body",
								trigger: "manual",
								content: c,
								tipClass: "popover-in-modal popover-success popover-form-result",
								placement: i.placement || r.popoverPlacement || "right"
							}).popover("show"), setTimeout(function() {
								f.popover("destroy")
							}, r.popoverTime || 2e3), g = !0), h.length && (h.html('<span class="small text-success">' + c + "</span>").show().delay(3e3).fadeOut(100), g = !0), g || s(c, "success")
						}
						if (u) return u(i, !0, l);
						if ((r.closeModal || i.closeModal) && setTimeout(t.zui.closeModal, r.closeModalTime || 2e3), p() === !1) return;
						var m = r.locate || i.locate;
						if (m) if ("loadInModal" == m) {
							var v = t(".modal");
							setTimeout(function() {
								v.load(v.attr("ref"), function() {
									t(this).find(".modal-dialog").css("width", t(this).data("width")), t.zui.ajustModalPosition()
								})
							}, 1e3)
						} else {
							var y = "reload" == m ? e.location.href : m;
							setTimeout(function() {
								e.location.href = y
							}, 1200)
						}
						var b = r.ajaxReload || i.ajaxReload;
						if (b) {
							var w = t(b);
							w.length && w.load(e.location.href + " " + b, function() {
								w.find('[data-toggle="modal"]').modalTrigger()
							})
						}
					} else {
						if (l.enableForm(), "string" == typeof c) h.length ? h.html('<span class="text-small text-red">' + c + "</span>").show().delay(3e3).fadeOut(100) : s(c, "danger");
						else if ("object" == typeof c) {
							var x = !1,
								C = [];
							t.each(c, function(e, i) {
								var n = t.isArray(i) ? i.join(";") : i,
									o = t("#" + e);
								if (!o.length) return void C.push(n);
								var s = e + "Label",
									a = t("#" + s);
								if (!a.length) {
									var r = o.closest(".input-group").length,
										l = o.closest("td").length;
									a = t('<div id="' + s + '" class="text-danger help-text" />').appendTo(l ? o.closest("td") : r ? o.closest(".input-group").parent() : o.parent())
								}
								a.empty().append(n), o.addClass("has-error");
								var h = function() {
										var e = t("#" + s);
										if (e.length) return e.remove(), o.removeClass("has-error"), !0
									};
								o.on("change input mousedown", h);
								var c = t("#" + e + "_chosen");
								c.length && c.find(".chosen-single,.chosen-choices").addClass("has-error").on("mousedown", function() {
									h() === !0 && t(this).removeClass("has-error")
								}), x || (o.focus(), x = !0)
							}), C.length && s(C.join(";"), "danger")
						}
						if (u) return u(i, !1, l);
						if (p() === !1) return
					}
				}
			},
			error: function(t, i, n) {
				if ((c && c(t, i, n, l)) !== !1) {
					l.enableForm();
					var o = "timeout" == i || "error" == i ? e.lang ? e.lang.timeout : i : t.responseText + i + n;
					s(o, "danger")
				}
			}
		}, r), l._ajaxForm(r).data("zui.ajaxform", !0), l.on("click", "[data-form-action]", function() {
			l.attr("action", t(this).data("formAction")).submit()
		})
	}, t.fn.ajaxForm = function(e) {
		return this.each(function() {
			t.ajaxForm(this, e)
		})
	}, t.fn.setInputRequired = function() {
		return this.each(function() {
			var e = t(this),
				i = e.parent();
			i.is(".input-control,td") ? i.addClass("required") : e.is(".chosen") ? e.attr("required", null).next(".chosen-container").addClass("required") : i.addClass("required"), e.attr("required", null);
			var n = i.closest(".input-group");
			n.length && 1 === n.find(".required,input[required],select[required]").length && n.addClass("required")
		})
	}, t(function() {
		t('.form-ajax,form[data-type="ajax"]').ajaxForm(), setTimeout(function() {
			var i = e.config.requiredFields,
				n = t("form");
			i && (i = i.split(",")), i && i.length && t.each(i, function(t, e) {
				n.find("#" + e).attr("required", "required")
			}), n.find("input[required],select[required],textarea[required]").setInputRequired()
		}, 400), t("#hiddenwin"), t('form[target="hiddenwin"]').on("submit", function() {
			var e = t(this);
			e.data("zui.ajaxform") || e.enableForm(!1).data("disabledTime", (new Date).getTime())
		}).on("click", function() {
			var e = t(this),
				i = e.data("disabledTime");
			i && (new Date).getTime() - i > 1e4 && e.enableForm(!0).data("disabledTime", null)
		})
	})
}(jQuery, window, void 0), function(t) {
	"use strict";
	var e = "zui.searchList",
		i = function(t, e) {
			if (t && t.length) for (var i = 0; i < t.length; ++i) if (e.indexOf(t[i]) < 0) return !1;
			return !0
		},
		n = function(i, o) {
			var s = this;
			s.name = e;
			var a = s.$ = t(i);
			o = s.options = t.extend({}, n.DEFAULTS, this.$.data(), o);
			var r = a.find(o.searchBox);
			r.length && (r.searchBox({
				onSearchChange: function(t) {
					s.search(t)
				},
				onKeyDown: function(t) {
					var e = t.which;
					if (13 === e) {
						var i = s.getActiveItem();
						i.length && (o.onSelectItem ? o.onSelectItem(i) : window.location.href = i.attr("href")), t.preventDefault()
					} else if (38 === e) {
						var i = s.getActiveItem();
						i.removeClass("active");
						for (var n = i.prev(); n.length && !n.is(".search-list-item:not(.hidden)");) n = n.prev();
						n.length || (n = s.getItems().not(".hidden").last()), s.scrollTo(n.addClass("active")), t.preventDefault()
					} else if (40 === e) {
						var i = s.getActiveItem();
						i.removeClass("active");
						for (var a = i.next(); a.length && !a.is(".search-list-item:not(.hidden)");) a = a.next();
						a.length || (a = s.getItems().not(".hidden").first()), s.scrollTo(a.addClass("active")), t.preventDefault()
					}
				}
			}), s.searchBox = r.data("zui.searchBox"), s.search(s.searchBox.getSearch()));
			var l = s.$menu = a.closest(".dropdown-menu");
			if (l.length) {
				s.isDropdown = !0, a.on("click", function(e) {
					t(e.target).closest(o.selector).length || e.stopPropagation()
				});
				var h = l.parent();
				h.on(h.hasClass("dropdown-hover") ? "mouseenter" : "shown.zui.dropdown", function() {
					s.tryLoadRemote(function() {
						setTimeout(function() {
							s.searchBox && s.searchBox.focus()
						}, 50)
					})
				})
			}
			a.on("mouseenter", o.selector, function() {
				a.find(s.options.selector).not(".hidden").removeClass("active"), t(this).addClass("active")
			})
		};
	n.prototype.tryLoadRemote = function(t) {
		var e = this,
			i = e.options;
		i.url || i.ajax ? e.isLoaded ? t() : e.loadRemote(t) : t()
	}, n.prototype.loadRemote = function(e) {
		var i = this,
			n = i.options;
		i.$menu.addClass("load-indicator loading").find(".list-group").remove(), i.isLoaded = !1, t.ajax(t.extend({
			url: n.url,
			type: "GET",
			dataType: "html",
			success: function(n, o, s) {
				var a = t(n);
				a.hasClass("list-group") || (a = t('<div class="list-group"></div>').append(a)), i.$menu.append(a), i.$menu.removeClass("loading"), i.isLoaded = !0, e && e(!0)
			},
			error: function() {
				i.$menu.removeClass("loading").append('<div class="list-group"><div class="text-error has-padding">' + (n.errorText || window.lang && window.lang.timeout) + "</div></div>"), e && e(!1)
			}
		}, n.ajax))
	}, n.prototype.scrollTo = function(t) {
		t.length && t[0].scrollIntoView({
			behavior: "smooth"
		})
	}, n.prototype.getItems = function() {
		return this.$.find(this.options.selector).addClass("search-list-item")
	}, n.prototype.getActiveItem = function() {
		return this.getItems().filter(".active:first")
	}, n.prototype.search = function(e) {
		var n = this,
			o = void 0 === e || null === e || "" === e,
			s = n.getItems().removeClass("active");
		if (o) s.removeClass("hidden");
		else {
			var a = t.trim(e).split(" ");
			s.each(function() {
				var e = t(this),
					n = e.text() + " " + (e.data("key") || e.data("filter"));
				e.toggleClass("hidden", !i(a, n))
			})
		}
		n.scrollTo(s.not(".hidden").first().addClass("active"))
	}, n.DEFAULTS = {
		selector: ".list-group a",
		searchBox: ".search-box",
		onSelectItem: null
	}, t.fn.searchList = function(i) {
		return this.each(function() {
			var o = t(this),
				s = o.data(e),
				a = "object" == typeof i && i;
			s || o.data(e, s = new n(this, a)), "string" == typeof i && s[i]()
		})
	}, n.NAME = e, t.fn.searchList.Constructor = n, t(function() {
		t('[data-ride="searchList"]').searchList()
	})
}(jQuery), function(t) {
	"use strict";
	var e = "zui.labelSelector",
		i = function(n, o) {
			var s = this;
			s.name = e, s.$ = t(n), o = s.options = t.extend({}, i.DEFAULTS, this.$.data(), o), s.$.hide(), s.update()
		};
	i.prototype.select = function(t) {
		t += "", this.$wrapper.find(".label.active").removeClass("active"), this.$wrapper.find('.label[data-value="' + t + '"]').addClass("active"), this.$.val(t).trigger("change")
	}, i.prototype.update = function() {
		var e = this,
			i = e.options,
			n = e.$wrapper;
		if (!n) {
			if (i.wrapper) n = t(i.wrapper);
			else {
				var o = e.$.next();
				n = o.hasClass(".label-selector") ? o : t('<div class="label-selector"></div>')
			}
			n.parent().length || e.$.after(n), e.$wrapper = n, n.on("click", ".label", function(i) {
				var n = e.$.val(),
					o = t(this).data("value");
				e.hasEmptyValue !== !1 && o == n && (o = e.hasEmptyValue), e.select(o), i.preventDefault()
			})
		}
		n.empty();
		var s = e.$.val();
		e.hasEmptyValue = !1, e.$.children("option").each(function() {
			var e = t(this),
				o = {
					label: e.text(),
					value: e.val()
				},
				a = "" === o.value || "0" === o.value,
				r = t(i.labelTemplate || '<span class="label"></span>');
			i.labelClass && !a && r.addClass(i.labelClass), i.labelCreator ? r = i.labelCreator(r) : (r.data("option", o).attr("data-value", o.value), a && !o.label ? r.addClass("empty").append('<i class="icon icon-close"></i>') : r.text(o.label).toggleClass("active", s === o.value)), n.append(r)
		})
	}, i.DEFAULTS = {}, t.fn.labelSelector = function(n) {
		return this.each(function() {
			var o = t(this),
				s = o.data(e),
				a = "object" == typeof n && n;
			s || o.data(e, s = new i(this, a)), "string" == typeof n && s[n]()
		})
	}, i.NAME = e, t.fn.labelSelector.Constructor = i, t(function() {
		t('[data-provide="labelSelector"]').labelSelector()
	})
}(jQuery), function(t) {
	"use strict";
	var e = "zui.fileInput",
		i = t.BYTE_UNITS = {
			B: 1,
			KB: 1024,
			MB: 1048576,
			GB: 1073741824,
			TB: 1099511627776
		},
		n = t.formatBytes = function(t, e, n) {
			return void 0 === e && (e = 2), n || (n = t < i.KB ? "B" : t < i.MB ? "KB" : t < i.GB ? "MB" : t < i.TB ? "GB" : "TB"), (t / i[n]).toFixed(e) + n
		},
		o = function(t) {
			if ("string" == typeof t) {
				t = t.toUpperCase();
				var e = t.replace(/\d+/, "");
				t = parseFloat(t.replace(e, "")), t *= i[e] || i[e + "B"], t = Math.floor(t)
			}
			return t
		},
		s = function(i, a) {
			var r = this;
			r.name = e;
			var l = r.$ = t(i);
			a = r.options = t.extend({}, s.DEFAULTS, this.$.data(), a), a.fileMaxSize && "string" == typeof a.fileMaxSize && (a.fileMaxSize = o(a.fileMaxSize));
			var h = r.$input = l.find('input[type="file"]');
			l.on("click", ".file-input-btn", function() {
				h.trigger("click")
			}).on("click", ".file-input-rename", function() {
				r.oldName = l.addClass("edit").find(".file-editbox").focus().val()
			}).on("click", ".file-input-delete", function() {
				h.val(""), r.update(), a.onDelete && a.onDelete(r)
			}).on("click", ".file-name-cancel", function() {
				l.removeClass("edit").find(".file-editbox").focus().val(r.oldName)
			}).on("click", ".file-name-confirm", function() {
				var e = l.find(".file-editbox"),
					i = t.trim(e.val());
				i.length ? l.removeClass("edit").find(".file-title").text(i) : e.focus()
			}).on("change input paste", ".file-editbox", function() {
				var e = t(this);
				e.attr("size", Math.max(5, e.val().length))
			}), h.on("change", function() {
				var t = r.getFile();
				t && a.fileMaxSize && t.size > a.fileMaxSize && (h.val(""), (window.bootbox || window).alert(a.fileSizeError.format(n(a.fileMaxSize)))), r.update()
			}), r.update()
		};
	s.prototype.getFile = function() {
		var t = this.$input.prop("files");
		return t && t[0]
	}, s.prototype.update = function(t) {
		var e = this,
			i = e.$,
			o = e.getFile(),
			s = !o;
		i.toggleClass("normal", !s).toggleClass("empty", s), o ? (e.oldName = o.name, i.find(".file-title").text(o.name).attr("title", o.name), i.find(".file-size").text(n(o.size)), i.find(".file-editbox").val(o.name).attr("size", o.name.length), e.options.onSelect && e.options.onSelect(o, e)) : i.find(".file-editbox").val("")
	}, s.DEFAULTS = {
		fileMaxSize: 0,
		fileSizeError: "无法上传大于 {0} 的文件。"
	}, t.fn.fileInput = function(i) {
		return this.each(function() {
			var n = t(this),
				o = n.data(e),
				a = "object" == typeof i && i;
			o || n.data(e, o = new s(this, a)), "string" == typeof i && o[i]()
		})
	}, s.NAME = e, t.fn.fileInput.Constructor = s, t(function() {
		t('[data-provide="fileInput"]').fileInput()
	});
	var a = "zui.fileInputList",
		r = function(e, i) {
			var n = this;
			n.name = a;
			var o = n.$ = t(e);
			i = n.options = t.extend({}, r.DEFAULTS, this.$.data(), i), n.$template = o.find(".file-input").detach(), n.add()
		};
	r.prototype.add = function() {
		var t = this,
			e = t.options,
			i = t.$template.clone();
		"before" === e.appendWay ? t.$.prepend(i) : t.$.append(i), i.fileInput({
			fileMaxSize: e.eachFileMaxSize,
			fileSizeError: e.fileSizeError,
			onDelete: function(e) {
				e.$.remove(), t.options.onDelete && t.options.onDelete(e, t)
			},
			onSelect: function(e, i) {
				t.add(), t.options.onSelect && t.options.onSelect(e, i, t)
			}
		})
	}, r.DEFAULTS = {
		fileMaxSize: 0,
		eachFileMaxSize: 0,
		appendWay: "after",
		fileSizeError: "无法上传大于 {0} 的文件。"
	}, t.fn.fileInputList = function(e) {
		return this.each(function() {
			var i = t(this),
				n = i.data(a),
				o = "object" == typeof e && e;
			n || i.data(a, n = new r(this, o)), "string" == typeof e && n[e]()
		})
	}, r.NAME = a, t.fn.fileInputList.Constructor = r, t(function() {
		t('[data-provide="fileInputList"]').fileInputList()
	})
}(jQuery), function(t) {
	window.config || (window.config = {}), t.createLink = window.createLink = function(t, e, n, o, s) {
		if (o || (o = config.defaultView), s || (s = !1), n) for (n = n.split("&"), i = 0; i < n.length; i++) {
			var a = n[i].split("=");
			n[i] = [a.shift(), a.join("=")]
		}
		var r;
		if ("GET" != config.requestType) {
			if ("PATH_INFO" == config.requestType && (r = config.webRoot + t + config.requestFix + e), "PATH_INFO2" == config.requestType && (r = config.webRoot + "index.php/" + t + config.requestFix + e), n) for (i = 0; i < n.length; i++) r += config.requestFix + n[i][1];
			r += "." + o
		} else if (r = config.router + "?" + config.moduleVar + "=" + t + "&" + config.methodVar + "=" + e + "&" + config.viewVar + "=" + o, n) for (i = 0; i < n.length; i++) r += "&" + n[i][0] + "=" + n[i][1];
		if (void 0 !== config.onlybody && "yes" == config.onlybody || s) {
			var l = "GET" != config.requestType ? "?onlybody=yes" : "&onlybody=yes";
			r += l
		}
		return r
	}, t(function() {
		var e = t("#main,#mainContent,#mainRow,.auto-fade-in");
		e.length && e.hasClass("fade") && setTimeout(function() {
			e.addClass("in")
		}, e.data("fadeTime") || 200)
	}), t.ajaxSendScore = function(e) {
		t.get(t.createLink("score", "ajax", "method=" + e))
	};
	var e = function(t) {
			var e = 0;
			if (t) {
				var i = t.split(":");
				e += 60 * parseInt(i[0]), e += parseInt(i[1])
			}
			return e
		},
		n = function(t) {
			t %= 1440;
			var e = Math.floor(t / 60),
				i = t % 60;
			return e < 10 && (e = "0" + e), i < 10 && (i = "0" + i), e + ":" + i
		},
		o = function(t) {
			if ("string" == typeof t && (t = e(t)), "number" == typeof t) if (t < 1e5) {
				var i = new Date;
				i.setHours(Math.floor(t / 60) % 24), i.setMinutes(t % 60), t = i
			} else t = new Date(t);
			return t
		},
		s = function(t, i) {
			for (var s = i ? o(i) : new Date, a = s.getHours(), r = 10 * Math.floor(s.getMinutes() / 10) + 10, l = 0; l < 24; ++l) {
				var h = (l + a) % 24;
				if (!(h < 5)) for (var c = 0; c < 6; ++c) {
					var d = n(60 * h + 10 * c + r);
					t.append('<option value="' + d + '">' + d + "</option>")
				}
			}
			t.val() || (time = e(s.format("hh:mm")), time = time - time % 10 + 10, t.val(n(time)))
		};
	t.fn.timeSpanControl = function(i) {
		return this.each(function() {
			var a = t(this),
				r = t.extend({}, i, a.data()),
				l = a.find('[name="begin"],.control-time-begin'),
				h = a.find('[name="end"],.control-time-end'),
				c = function() {
					var t = l.val();
					if (a.find(".hide-empty-begin").toggleClass("hide", !t), t) {
						var i = n(e(t) + 30);
						h.find('option[value="' + i + '"]').length && h.val(i), r.onChange && r.onChange(h, i)
					}
				};
			if (a.data("timeSpanControlInit")) {
				if (r.begin) {
					var d = o(r.begin).format("hh:mm");
					l.find('option[value="' + d + '"]').length && l.val(d), r.onChange && r.onChange(l, d)
				}
				if (r.end) {
					var u = o(r.end).format("hh:mm");
					h.find('option[value="' + u + '"]').length && h.val(u), r.onChange && r.onChange(h, u)
				}
			} else l.on("change", c), s(l, r.begin), s(h, r.end), a.data("timeSpanControlInit", !0);
			r.end || c()
		})
	}, t.timeSpanControl = {
		convertTimeToNum: e,
		convertNumToTime: n,
		initTimeSelect: s,
		createTime: o
	};
	var a = t.setSearchType = function(e, i) {
			var n = t("#searchType");
			e || (e = n.val()), e = e || "bug", n.val(e);
			var o = t("#searchTypeMenu");
			o.find("li.selected").removeClass("selected");
			var s = o.find('a[data-value="' + e + '"]'),
				a = s.text();
			s.parent().addClass("selected"), t("#searchTypeName").text(a), i || t("#searchInput").focus()
		};
	t.gotoObject = function(e, i) {
		e || (e = t("#searchType").val()), i || (i = t("#searchInput").val()), i && e && (window.location.href = t.createLink(e, "testsuite" === e ? "library" : "view", "id=" + i))
	}, t(function() {
		a(null, !0), t(document).on("keydown", function(e) {
			e.ctrlKey && 71 === e.keyCode && (t("#searchInput").val("").focus(), e.stopPropagation(), e.preventDefault())
		})
	}), t.removeAnchor = window.removeAnchor = function(t) {
		var e = t.lastIndexOf("#");
		return e > -1 ? t.substr(0, e) : t
	}, t.refreshPage = function() {
		location.href = removeAnchor(location.href)
	}, t.selectLang = window.selectLang = function(e) {
		t.cookie("lang", e, {
			expires: config.cookieLife,
			path: config.webRoot
		}), t.ajaxSendScore("selectLang"), t.refreshPage()
	}, t.selectTheme = window.selectTheme = function(e) {
		t.cookie("theme", e, {
			expires: config.cookieLife,
			path: config.webRoot
		}), t.ajaxSendScore("selectTheme"), t.refreshPage()
	}, t.chosenDefaultOptions = {
		disable_search_threshold: 1,
		compact_search: !0,
		allow_single_deselect: !0,
		placeholder_text_single: " ",
		placeholder_text_multiple: " ",
		search_contains: !0,
		drop_direction: function() {
			var e = t(this.container).closest(".table-responsive:not(.scroll-none)");
			if (e.length) {
				if (this.drop_directionFixed) return this.drop_directionFixed;
				var i = "down",
					n = this.container.find(".chosen-drop"),
					o = this.container.position(),
					s = n.outerHeight();
				return o.top >= s && o.top + s < e.outerHeight() && (i = "up"), this.drop_directionFixed = i, i
			}
			return "auto"
		}
	}, t.chosenSimpleOptions = t.extend({}, t.chosenDefaultOptions, {
		disable_search_threshold: 6
	}), t.fn._chosen = t.fn.chosen, t.fn.chosen = function(e) {
		return "string" == typeof e ? this._chosen(e) : this.each(function() {
			var i = t(this).addClass("chosen-controled");
			return i._chosen(t.extend({}, i.hasClass("chosen-simple") ? t.chosenSimpleOptions : t.chosenDefaultOptions, i.data(), e))
		})
	}, t(function() {
		t(".chosen,.chosen-simple").each(function() {
			var e = t(this);
			e.closest(".template").length || e.chosen()
		})
	}), t.extend(t.fn.pager.Constructor.DEFAULTS, {
		maxNavCount: 8,
		prevIcon: "icon-angle-left",
		nextIcon: "icon-angle-right",
		firstIcon: "icon-first-page",
		lastIcon: "icon-last-page",
		navEllipsisItem: "…",
		menuDirection: "dropup",
		pageSizeOptions: [5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 100, 200, 500, 1e3, 2e3],
		elements: ["total_text", "size_menu", "first_icon", "prev_icon", '<div class="pager-label"><strong>{page}</strong>/<strong>{totalPage}</strong></div>', "next_icon", "last_icon"],
		onPageChange: function(e, i) {
			e.recPerPage !== i.recPerPage && t.cookie(this.options.pageCookie, e.recPerPage, {
				expires: config.cookieLife,
				path: config.webRoot
			}), e.recPerPage !== i.recPerPage && (window.location.href = this.createLink())
		}
	}), t.zui.Messager.DEFAULTS.cssClass = "messagger-zt", t.fn.reverseOrder = function() {
		return this.each(function() {
			var e = t(this);
			e.prependTo(e.parent())
		})
	};
	var r = function(e, i) {
			var n = t(e);
			i = t.extend({}, n.data(), i);
			var o = n.find(".histories-list"),
				s = !0,
				a = !1;
			n.on("click", ".btn-reverse", function() {
				o.children("li").reverseOrder(), s = !s, t(this).find(".icon").toggleClass("icon-arrow-up", s).toggleClass("icon-arrow-down", !s)
			}).on("click", ".btn-expand-all", function() {
				var e = t(this).find(".icon");
				a = !a, e.toggleClass("icon-plus", !a).toggleClass("icon-minus", a), o.children("li").toggleClass("show-changes", a)
			}).on("click", ".btn-expand", function() {
				t(this).closest("li").toggleClass("show-changes")
			}).on("click", ".btn-strip", function() {
				var e = t(this),
					n = e.find(".icon"),
					o = n.hasClass("icon-code");
				n.toggleClass("icon-code", !o).toggleClass("icon-text", o), e.attr("title", o ? i.original : i.textdiff), e.closest("li").toggleClass("show-original", o)
			}), o.find(".btn-strip").attr("title", i.original);
			var r = n.find(".modal-comment").modal({
				show: !1
			}).on("shown.zui.modal", function() {
				var t = r.find("#comment");
				t.length && (t.focus(), window.editor && window.editor.comment && window.editor.comment.focus())
			}).on("show.zui.modal", function() {
				var e = r.find("#comment");
				e.length && !e.data("keditor") && t.fn.kindeditor && e.kindeditor()
			});
			n.on("click", ".btn-comment", function(t) {
				r.modal("toggle"), t.preventDefault()
			}).on("click", ".btn-edit-comment,.btn-hide-form", function() {
				t(this).closest("li").toggleClass("show-form")
			});
			var l = n.find(".comment-edit-form");
			l.ajaxForm({
				success: function(t, e, i, n) {
					setTimeout(function() {
						l.closest("li").removeClass("show-form")
					}, 2e3)
				}
			})
		};
	t.fn.histories = function(t) {
		return this.each(function() {
			r(this, t)
		})
	}, t(function() {
		t(".histories").histories()
	});
	var l = 0,
		h = 0;
	t.toggleSidebar = function(e) {
		var i = t("#sidebar");
		if (i.length) {
			var n = t("main");
			if (void 0 === e) e = n.hasClass("hide-sidebar");
			else if (e && !n.hasClass("hide-sidebar")) return;
			n.toggleClass("hide-sidebar", !e), clearTimeout(l), t.zui.store.set(h, e);
			var o = i.children(".cell"),
				s = {
					overflow: "visible",
					maxHeight: "initial"
				};
			e ? (i.addClass("showing"), l = setTimeout(function() {
				i.removeClass("showing"), i.trigger("sidebar.toggle", e)
			}, 210)) : (i.trigger("sidebar.toggle", e), t(window).width() < 1900 && (s = {
				overflow: "hidden",
				maxHeight: t(window).height() - 45
			})), o.css(s)
		}
	};
	var c = t.initSidebar = function() {
			var e = t("#sidebar");
			if (e.length) {
				if (e.data("init")) return !0;
				h = "sidebar:" + (e.data("id") || config.currentModule + "/" + config.currentMethod);
				var i = t("main");
				i.on("click", ".sidebar-toggle", function() {
					t.toggleSidebar(i.hasClass("hide-sidebar"))
				});
				var n = t.zui.store.get(h, e.data("hide") !== !1);
				n === !1 && e.addClass("no-animate"), t.toggleSidebar(n), n === !1 && setTimeout(function() {
					e.removeClass("no-animate")
				}, 500);
				var o = function() {
						var i = e.find(".sidebar-toggle");
						if (i.length) {
							var n = i[0].getBoundingClientRect(),
								o = t(window).height(),
								s = Math.max(0, Math.floor(Math.min(o - 40, n.top + n.height) - Math.max(n.top, 0)) / 2) + (n.top < 0 ? 0 - n.top : 0);
							i.find(".icon").css("top", s)
						}
					};
				return o(), e.on("sidebar.toggle", o), t(window).on("resize", o).on("scroll", o), e.data("init", 1), !0
			}
		};
	c() || t(c), t.toggleQueryBox = function(e, i) {
		var n = t(i || "#queryBox");
		n.length && (void 0 === e && (e = !n.hasClass("show")), n.toggleClass("show", !! e), n.data("init") || (n.addClass("load-indicator loading").data("init", 1), t.get(t.createLink("search", "buildForm"), function(t) {
			n.html(t).removeClass("loading")
		})), t(".querybox-toggle").toggleClass("querybox-opened", e))
	}, t(function() {
		var e = t("#queryBox");
		e.length && (t(document).on("click", ".querybox-toggle", function() {
			t.toggleQueryBox()
		}), e.hasClass("show") && t.toggleQueryBox(!0))
	}), t.extend(t.fn.colorPicker.Constructor.DEFAULTS, {
		colors: ["#3DA7F5", "#75C941", "#2DBDB2", "#797EC9", "#FFAF38", "#FF4E3E"]
	}), window.setCheckedCookie = function() {
		var e = [],
			i = t('#mainContent .main-table tbody>tr input[type="checkbox"]:checked');
		i.each(function() {
			var i = parseInt(t(this).val(), 10);
			NaN !== i && e.push(i)
		}), t.cookie("checkedItem", e.join(","), {
			expires: config.cookieLife,
			path: config.webRoot
		})
	}, t.extend(t.fn.modal.bs.Constructor.DEFAULTS, {
		scrollInside: !0,
		backdrop: "static",
		headerHeight: 100
	}), t.extend(t.zui.ModalTrigger.DEFAULTS, {
		scrollInside: !0,
		backdrop: "static",
		headerHeight: 40
	}), t.fn.initIframeModal = function() {
		return this.each(function() {
			var e = t(this);
			if (!e.parents('[data-ride="table"],.skip-iframe-modal').length) {
				var i = {
					type: "iframe"
				};
				e.hasClass("export") && t.extend(i, {
					width: 800,
					shown: setCheckedCookie
				}, e.data()), e.modalTrigger(i)
			}
		})
	}, t(function() {
		t("a.iframe,.export").initIframeModal()
	});
	var d = function() {
			var e, i, n = t(this),
				o = t.extend({
					limitSize: 40,
					suffix: "…"
				}, n.data()),
				s = n.text();
			if (s.length > o.limitSize) {
				e = s, i = s.substr(0, o.limitSize) + o.suffix, n.text(i).addClass("limit-text-on");
				var a = o.toggleBtn ? t(o.toggleBtn) : n.next(".text-limit-toggle");
				a.text(a.data("textExpand")), a.on("click", function() {
					var t = n.toggleClass("limit-text-on").hasClass("limit-text-on");
					n.text(t ? i : e), a.text(a.data(t ? "textExpand" : "textCollapse"))
				})
			} else(o.toggleBtn ? t(o.toggleBtn) : n.next(".text-limit-toggle")).hide()
		};
	t.fn.textLimit = function() {
		return this.each(d)
	}, t(function() {
		t(".text-limit").textLimit()
	}), t.fixedTableHead = window.fixedTableHead = function(e, i) {
		var n = t(e);
		if (n.is("table") || (n = n.find("table")), n.length) {
			var o = t(i || window),
				s = null,
				a = function() {
					var e = n.children("thead"),
						i = e[0].getBoundingClientRect(),
						o = n.next(".fixed-head-table");
					if (i.top < 0) {
						var a = e.width();
						if (o.length) {
							if (s !== a) {
								s = a;
								var r = o.find("th");
								e.find("th").each(function(e) {
									r.eq(e).width(t(this).width())
								})
							}
						} else {
							var o = t("<table class='table fixed-head-table' style='position:fixed; top: 0;'></table>").addClass(n.attr("class")),
								l = e.clone(),
								r = l.find("th");
							e.find("th").each(function(e) {
								r.eq(e).width(t(this).width())
							}), o.append(l).insertAfter(n)
						}
						o.css({
							left: i.left,
							width: i.width
						}).show()
					} else o.hide()
				};
			o.on("scroll", a).on("resize", a), a()
		}
	}, t(document).on("click", "tr[data-url]", function() {
		var e = t(this),
			i = e.data("href") || e.data("url");
		i && (window.location.href = i)
	}), "yes" === config.onlybody && self === parent && (window.location.href = window.location.href.replace("?onlybody=yes", "").replace("&onlybody=yes", "")), t(function() {
		t("body").addClass("m-{currentModule}-{currentMethod}".format(config))
	});
	var u, p, f, g, m, v = function() {
			u || (u = t("#subNavbar"), p = t("#pageNav"), f = t("#pageActions"), g = u.children(".nav"), m = g.outerWidth());
			var e = u.outerWidth(),
				i = p.outerWidth() || 0,
				n = f.outerWidth() || 0;
			if (i = i ? i + 15 : 0, n = n ? n + 15 : 0, !i && !n) return void g.css({
				maxWidth: null,
				left: null,
				position: "static"
			});
			var o = Math.max(300, e - i - n),
				s = Math.min(o, m),
				a = (e - s) / 2,
				r = i && a < i ? i : n && a < n ? e - s - n : 0;
			g.css({
				maxWidth: o,
				left: r ? r - a : 0,
				position: "relative"
			})
		},
		y = function() {
			t.cookie("windowWidth", window.innerWidth), t.cookie("windowHeight", window.innerHeight), v()
		};
	t(y), t(window).on("resize", y);
	var b = function() {
			var e = t("#back").attr("href");
			e && (window.location.href = e)
		},
		w = function() {
			t.cookie("ajax_lastNext") || (t.cookie("ajax_lastNext", "on", {
				expires: config.cookieLife,
				path: config.webRoot
			}), t.ajaxSendScore("lastNext"))
		},
		x = function() {
			var e = t("#prevPage").attr("href");
			e && (window.location.href = e), w()
		},
		C = function() {
			var e = t("#nextPage").attr("href");
			e && (window.location.href = e), w()
		};
	t(document).on("keydown", function(t) {
		t.altKey && 38 === t.keyCode ? b() : 37 === t.keyCode ? x() : 39 === t.keyCode && C()
	}), t.fn.tree.Constructor.DEFAULTS.initialState = "preserve", t.closeModal = function(e, i, n) {
		t.zui.closeModal(n, e, i)
	}, t.getThemeColor = function(e) {
		if (!t.themeColor) {
			var i = t("#mainHeader");
			i.length && (t.themeColor = {
				primary: i.css("border-top-color"),
				pale: i.css("border-bottom-color"),
				secondary: i.css("background-color")
			})
		}
		return e ? t.themeColor && t.themeColor[e] : t.themeColor
	};
	var _ = function(e) {
			var i, n, o = t(e),
				s = o.children(".input-group-addon,.form-control:not(.chosen-controled),.chosen-container,.btn,.input-control,.input-group-btn,.datepicker-wrapper");
			s.each(function(e) {
				var o = t(this),
					a = o.is(".input-group-addon") ? "addon" : o.is(".chosen-container") ? "chosen" : o.is(".btn") ? "btn" : o.is(".input-control,.datepicker-wrapper") ? "insideInput" : o.is(".input-group-btn") ? "insideBtn" : "input",
					r = !i,
					l = e === s.length - 1,
					h = {};
				h.borderTopLeftRadius = 0, h.borderBottomLeftRadius = 0, h.borderTopRightRadius = 0, h.borderBottomRightRadius = 0, r && ("addon" === a && (h.borderLeftWidth = 1), h.borderTopLeftRadius = 2, h.borderBottomLeftRadius = 2), l && ("addon" === a && (h.borderRightWidth = 1), h.borderTopRightRadius = 2, h.borderBottomRightRadius = 2), n && ("chosen" !== n && "input" !== n && "btn" !== n && "insideInput" !== n && "insideBtn" !== n || "chosen" !== a && "input" !== n && "btn" !== a && "insideInput" !== a && "insideBtn" !== a || (h.borderLeftColor = "transparent")), ("insideBtn" === a ? o.find(".btn") : "insideInput" === a ? o.find(".form-control") : "chosen" === a ? o.find(".chosen-single,.chosen-choices") : o).css(h), i = o, n = a
			})
		};
	t.fn.fixInputGroup = function() {
		return this.each(function() {
			_(this)
		})
	};
	var k = function() {
			var e = t(".main-actions>.btn-toolbar");
			if (e.length) {
				var i, n, o = !1,
					s = null,
					a = e.children(),
					r = a.length;
				for (a.each(function(e) {
					i = t(this), n = i.is(".divider"), n && !s && i.hide(), o || n || (o = !0), s = n ? null : i, !n || e !== r - 1 && 0 !== e || i.hide()
				}); i.length && i.is(".divider");) i = i.hide().prev();
				o || e.hide()
			}
		};
	t(function() {
		t(".input-group,.btn-group").fixInputGroup(), k()
	}), window.holders && t.each(window.holders, function(e) {
		var i = t("#" + e);
		i.length && i.is("input") && i.attr("placeholder", window.holders[e])
	});
	var T = function() {
			var e, i = "en" == config.clientLang ? "http://www.zentao.pm/book/zentaomanual/8.html?fullScreen=zentao" : "http://www.zentao.net/book/zentaopmshelp.html?fullScreen=zentao",
				n = t("#navbar > .nav").first(),
				o = 1e4,
				s = function() {
					clearTimeout(e), t("#helpContent").removeClass("show-error")
				},
				a = t.openHelp = function() {
					s(), n.children("li.active:not(#helpMenuItem)").removeClass("active").addClass("close-help-tab"), t("#helpMenuItem").addClass("active");
					var a = t("#helpContent");
					if (a.length) {
						if (t("body").hasClass("show-help-tab")) return void t("#helpIframe").get(0).contentWindow.location.replace(i)
					} else {
						a = t('<div id="helpContent"><div class="load-error text-center"><h4 class="text-danger">' + lang.timeout + '</h4><p><a href="###" class="open-help-tab"><i class="icon icon-arrow-right"></i> ' + i + '</a></p></div><iframe id="helpIframe" name="helpIframe" src="' + i + '" frameborder="no" allowtransparency="true" scrolling="auto" hidefocus="" style="width: 100%; height: 100%; left: 0px;"></iframe></div>'), t("#header").after(a);
						var r = t("#helpIframe").get(0);
						e = setTimeout(function() {
							t("#helpContent").addClass("show-error")
						}, o), r.onload = r.onreadystatechange = function() {
							this.readyState && "complete" != this.readyState || s()
						}
					}
					t("body").addClass("show-help-tab")
				},
				r = t.closeHelp = function() {
					t("body").removeClass("show-help-tab"), t("#helpMenuItem").removeClass("active"), n.find("li.close-help-tab").removeClass("close-help-tab").addClass("active").find("a").focus()
				};
			t(document).on("click", ".open-help-tab", function(e) {
				var i = t("#helpMenuItem");
				i.length || (i = t('<li id="helpMenuItem"><a href="javascript:;" class="open-help-tab">' + t(this).text() + ' <i class="icon icon-close close-help-tab icon-sm"></i></a></li>'), n.append('<li class="divider"></li>').append(i)), a(), e.preventDefault()
			}).on("click", ".close-help-tab", function(t) {
				r(), t.stopPropagation(), t.preventDefault()
			})
		};
	t(T), t(function() {
		var e = t(".table-responsive"),
			i = function() {
				e.each(function() {
					this.scrollHeight - 3 <= this.clientHeight && this.scrollWidth - 3 <= this.clientWidth ? t(this).addClass("scroll-none").css("overflow", "visible") : t(this).removeClass("scroll-none").css("overflow", "auto")
				})
			};
		e.length && (i(), t(window).on("resize", i))
	});
	var S = function() {
			var e = this.value ? this.scrollHeight + 2 + "px" : "32px";
			this.style.height = "auto", this.style.height = e, t(this).closest("tr").find("textarea").each(function() {
				this.style.height = e
			})
		};
	t.autoResizeTextarea = function(e) {
		t(e).each(S)
	}, t(function() {
		t("textarea.autosize").each(S), t(document).on("input keyup paste change", "textarea.autosize", S)
	}), t(function() {
		var e = t("#dropMenu");
		e.length && e.on("click", ".toggle-right-col", function(t) {
			e.toggleClass("show-right-col"), t.stopPropagation(), t.preventDefault()
		})
	});
	var D = "undefined" != typeof InstallTrigger;
	t.zui.browser.firefox = D, t("html").toggleClass("is-firefox", D).toggleClass("not-firefox", !D), t(function() {
		var e = t("#mainContent>.main-col"),
			i = e.children(".main-actions"),
			n = i.prev();
		if (i.length && n.length) {
			t('<div class="main-actions-holder"></div>').css("height", i.outerHeight()).insertAfter(i);
			var o = function() {
					var e = n[0].getBoundingClientRect(),
						o = e.top + e.height + 120 > t(window).height();
					t("body").toggleClass("main-actions-fixed", o), o && i.width(n.width())
				};
			t.resetToolbarPosition = o, o(), t(window).on("resize scroll", o)
		}
	}), t(document).on("show.zui.modal", function() {
		t("body.body-modal").length && window.parent && window.parent !== window && window.parent.$("body").addClass("hide-modal-close")
	}).on("hidden.zui.modal", function() {
		t("body.body-modal").length && window.parent && window.parent !== window && window.parent.$("body").removeClass("hide-modal-close")
	})
}(jQuery);