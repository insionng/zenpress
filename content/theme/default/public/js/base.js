function initFastSection() {
	function e() {
		if ("desktop" == deviceType) {
			var e = window.innerHeight - 210,
			t = $(".J_fastSectionList .wrap").height(); ! $(".J_fastSectionList .wrap").data("origin") && $(".J_fastSectionList .wrap").data("origin", t),
			$(".J_fastSectionList .wrap").data("origin") >= e && $(".J_fastSectionList .wrap").height(e),
			$(".J_fastSection").trigger("sticky_kit:detach"),
			$(".J_fastSection").stick_in_parent({
				parent: ".main-section"
			}),
			$(window).trigger("scroll")
		} else $(".J_fastSectionList .wrap").height("auto"),
		$(".J_fastSection").trigger("sticky_kit:detach");
		$(".J_fastSectionList .wrap").perfectScrollbar("update")
	}
	function t() {
		$(".J_fastSectionList .wrap").bind("scroll",
		function() {
			var e = ($(this).scrollTop(), $(this).find(".panel:visible .load-more")),
			t = e.offset().top - $(this).offset().top - $(this).height();
			if (0 > t) {
				var n = e,
				i = n.attr("href");
				if (n.attr("href", "javascript:void(0)"), n.hasClass("no-data")) return;
				if (n.hasClass("loading")) return;
				n.addClass("loading"),
				$.get(i,
				function(e) {
					setTimeout(function() {
						n.parent();
						$(e).insertAfter(n),
						n.remove(),
						$(".J_fastSectionList .wrap").perfectScrollbar("update")
					},
					0)
				},
				"html")
			}
		}).delegate(".panel:visible .load-more", "click",
		function(e) {
			if (e.preventDefault(), "desktop" != deviceType) {
				var t = $(this),
				n = t.attr("href");
				t.attr("href", "javascript:void(0)"),
				t.hasClass("no-data") || t.hasClass("loading") || (t.addClass("loading"), $.get(n,
				function(e) {
					setTimeout(function() {
						t.parent();
						$(e).insertAfter(t),
						t.remove(),
						$(".J_fastSectionList .wrap").perfectScrollbar("update")
					},
					0)
				},
				"html"))
			}
		})
	}
	$(".J_fastSectionList .wrap").perfectScrollbar({
		wheelPropagation: !0
	}),
	e(),
	$(window).resize(e),
	$(".J_fastSectionList .tabs a").each(function(t) {
		$(this).click(function(n) {
			n.preventDefault(),
			$(this).addClass("active").siblings().removeClass("active"),
			$(".J_fastSectionList .panel").eq(t).show().siblings().hide(),
			$(".J_fastSectionList .wrap").perfectScrollbar("update"),
			$(".J_fastSectionList .wrap").data("origin", "").css("height", "auto"),
			e(),
			$(window).trigger("resize")
		})
	});
	var n = function() {
		$(this).hasClass("active") ? $(this).removeClass("active") : $(this).addClass("active").siblings().removeClass("active"),
		$(".J_fastSectionList .wrap").perfectScrollbar("update")
	};
	$(".J_fastSectionList").delegate("section", "click", n),
	t(),
	$(".J_fastSectionList .wrap .product").magnificPopup({
		delegate: "img",
		type: "image"
	}),
	$(".J_fastSectionList .wrap .product").delegate("img", "click",
	function(e) {
		e.stopPropagation(),
		$.magnificPopup.open({
			items: {
				src: $(this).attr("src")
			},
			type: "image"
		},
		0)
	})
}
function initMobileNav(e) {
	/*$(".J_mobileNav a").each(function(t) {
		$(this).click(function(n) {
			n.preventDefault(),
			$(this).addClass("active").siblings().removeClass("active");
			var i = $(".J_fastSection").add(e);
			switch (i.addClass("mobile-hide"), t) {
			case 0:
				$(e).removeClass("mobile-hide");
				break;
			case 1:
				$(".J_fastSection").removeClass("mobile-hide"),
				$(".J_fastSectionList .tabs a").eq(0).trigger("click");
				break;
			case 2:
				$(".J_fastSection").removeClass("mobile-hide"),
				$(".J_fastSectionList .tabs a").eq(1).trigger("click")
			}
		})
	})*/
}
function initLazyLoad() {
	function e() {
		var e = $(window).scrollTop() + $(window).height();
		$("[data-lazyload]").each(function() {
			var t = $(this).offset().top,
			n = $(this);
			if ($(this).data("lazyload") && (n.addClass("before-fade-in"), e > t && $(this).data("lazyload"))) {
				var i = $(this).data("lazyload"),
				r = new Image;
				console.log(1 == n.data("fitMobile"), "desktop" != deviceType),
				1 == n.data("fitMobile") && "desktop" != deviceType && (i = i.replace(/\!.+$|$/, "")),
				n.data("lazyload", null).removeAttr("data-lazyload"),
				console.log(i),
				r.onload = function() {
					console.log(n[0].tagName),
					"img" == n[0].tagName.toLowerCase() ? n[0].src = i: n.css({
						"background-image": "url(" + i + ")"
					}),
					n.addClass("after-fade-in")
				},
				r.src = i
			}
		})
	}
	$(document).ready(function() {
		e(),
		$(window).scroll(function() {
			e()
		})
	})
} !
function(e, t) {
	"object" == typeof module && "object" == typeof module.exports ? module.exports = e.document ? t(e, !0) : function(e) {
		if (!e.document) throw new Error("jQuery requires a window with a document");
		return t(e)
	}: t(e)
} ("undefined" != typeof window ? window: this,
function(e, t) {
	function n(e) {
		var t = e.length,
		n = Z.type(e);
		return "function" === n || Z.isWindow(e) ? !1 : 1 === e.nodeType && t ? !0 : "array" === n || 0 === t || "number" == typeof t && t > 0 && t - 1 in e
	}
	function i(e, t, n) {
		if (Z.isFunction(t)) return Z.grep(e,
		function(e, i) {
			return !! t.call(e, i, e) !== n
		});
		if (t.nodeType) return Z.grep(e,
		function(e) {
			return e === t !== n
		});
		if ("string" == typeof t) {
			if (st.test(t)) return Z.filter(t, e, n);
			t = Z.filter(t, e)
		}
		return Z.grep(e,
		function(e) {
			return z.call(t, e) >= 0 !== n
		})
	}
	function r(e, t) {
		for (; (e = e[t]) && 1 !== e.nodeType;);
		return e
	}
	function o(e) {
		var t = ht[e] = {};
		return Z.each(e.match(pt) || [],
		function(e, n) {
			t[n] = !0
		}),
		t
	}
	function a() {
		V.removeEventListener("DOMContentLoaded", a, !1),
		e.removeEventListener("load", a, !1),
		Z.ready()
	}
	function s() {
		Object.defineProperty(this.cache = {},
		0, {
			get: function() {
				return {}
			}
		}),
		this.expando = Z.expando + s.uid++
	}
	function l(e, t, n) {
		var i;
		if (void 0 === n && 1 === e.nodeType) if (i = "data-" + t.replace(wt, "-$1").toLowerCase(), n = e.getAttribute(i), "string" == typeof n) {
			try {
				n = "true" === n ? !0 : "false" === n ? !1 : "null" === n ? null: +n + "" === n ? +n: bt.test(n) ? Z.parseJSON(n) : n
			} catch(r) {}
			yt.set(e, t, n)
		} else n = void 0;
		return n
	}
	function c() {
		return ! 0
	}
	function u() {
		return ! 1
	}
	function d() {
		try {
			return V.activeElement
		} catch(e) {}
	}
	function f(e, t) {
		return Z.nodeName(e, "table") && Z.nodeName(11 !== t.nodeType ? t: t.firstChild, "tr") ? e.getElementsByTagName("tbody")[0] || e.appendChild(e.ownerDocument.createElement("tbody")) : e
	}
	function p(e) {
		return e.type = (null !== e.getAttribute("type")) + "/" + e.type,
		e
	}
	function h(e) {
		var t = Mt.exec(e.type);
		return t ? e.type = t[1] : e.removeAttribute("type"),
		e
	}
	function g(e, t) {
		for (var n = 0,
		i = e.length; i > n; n++) vt.set(e[n], "globalEval", !t || vt.get(t[n], "globalEval"))
	}
	function m(e, t) {
		var n, i, r, o, a, s, l, c;
		if (1 === t.nodeType) {
			if (vt.hasData(e) && (o = vt.access(e), a = vt.set(t, o), c = o.events)) {
				delete a.handle,
				a.events = {};
				for (r in c) for (n = 0, i = c[r].length; i > n; n++) Z.event.add(t, r, c[r][n])
			}
			yt.hasData(e) && (s = yt.access(e), l = Z.extend({},
			s), yt.set(t, l))
		}
	}
	function v(e, t) {
		var n = e.getElementsByTagName ? e.getElementsByTagName(t || "*") : e.querySelectorAll ? e.querySelectorAll(t || "*") : [];
		return void 0 === t || t && Z.nodeName(e, t) ? Z.merge([e], n) : n
	}
	function y(e, t) {
		var n = t.nodeName.toLowerCase();
		"input" === n && Tt.test(e.type) ? t.checked = e.checked: ("input" === n || "textarea" === n) && (t.defaultValue = e.defaultValue)
	}
	function b(t, n) {
		var i, r = Z(n.createElement(t)).appendTo(n.body),
		o = e.getDefaultComputedStyle && (i = e.getDefaultComputedStyle(r[0])) ? i.display: Z.css(r[0], "display");
		return r.detach(),
		o
	}
	function w(e) {
		var t = V,
		n = Rt[e];

		return n || (n = b(e, t), "none" !== n && n || (Ot = (Ot || Z("<iframe frameborder='0' width='0' height='0'/>")).appendTo(t.documentElement), t = Ot[0].contentDocument, t.write(), t.close(), n = b(e, t), Ot.detach()), Rt[e] = n),
		n
	}
	function x(e, t, n) {
		var i, r, o, a, s = e.style;
		return n = n || Ft(e),
		n && (a = n.getPropertyValue(t) || n[t]),
		n && ("" !== a || Z.contains(e.ownerDocument, e) || (a = Z.style(e, t)), Xt.test(a) && qt.test(t) && (i = s.width, r = s.minWidth, o = s.maxWidth, s.minWidth = s.maxWidth = s.width = a, a = n.width, s.width = i, s.minWidth = r, s.maxWidth = o)),
		void 0 !== a ? a + "": a
	}
	function k(e, t) {
		return {
			get: function() {
				return e() ? void delete this.get: (this.get = t).apply(this, arguments)
			}
		}
	}
	function C(e, t) {
		if (t in e) return t;
		for (var n = t[0].toUpperCase() + t.slice(1), i = t, r = Kt.length; r--;) if (t = Kt[r] + n, t in e) return t;
		return i
	}
	function T(e, t, n) {
		var i = Yt.exec(t);
		return i ? Math.max(0, i[1] - (n || 0)) + (i[2] || "px") : t
	}
	function S(e, t, n, i, r) {
		for (var o = n === (i ? "border": "content") ? 4 : "width" === t ? 1 : 0, a = 0; 4 > o; o += 2)"margin" === n && (a += Z.css(e, n + kt[o], !0, r)),
		i ? ("content" === n && (a -= Z.css(e, "padding" + kt[o], !0, r)), "margin" !== n && (a -= Z.css(e, "border" + kt[o] + "Width", !0, r))) : (a += Z.css(e, "padding" + kt[o], !0, r), "padding" !== n && (a += Z.css(e, "border" + kt[o] + "Width", !0, r)));
		return a
	}
	function _(e, t, n) {
		var i = !0,
		r = "width" === t ? e.offsetWidth: e.offsetHeight,
		o = Ft(e),
		a = "border-box" === Z.css(e, "boxSizing", !1, o);
		if (0 >= r || null == r) {
			if (r = x(e, t, o), (0 > r || null == r) && (r = e.style[t]), Xt.test(r)) return r;
			i = a && (Q.boxSizingReliable() || r === e.style[t]),
			r = parseFloat(r) || 0
		}
		return r + S(e, t, n || (a ? "border": "content"), i, o) + "px"
	}
	function $(e, t) {
		for (var n, i, r, o = [], a = 0, s = e.length; s > a; a++) i = e[a],
		i.style && (o[a] = vt.get(i, "olddisplay"), n = i.style.display, t ? (o[a] || "none" !== n || (i.style.display = ""), "" === i.style.display && Ct(i) && (o[a] = vt.access(i, "olddisplay", w(i.nodeName)))) : (r = Ct(i), "none" === n && r || vt.set(i, "olddisplay", r ? n: Z.css(i, "display"))));
		for (a = 0; s > a; a++) i = e[a],
		i.style && (t && "none" !== i.style.display && "" !== i.style.display || (i.style.display = t ? o[a] || "": "none"));
		return e
	}
	function L(e, t, n, i, r) {
		return new L.prototype.init(e, t, n, i, r)
	}
	function E() {
		return setTimeout(function() {
			Qt = void 0
		}),
		Qt = Z.now()
	}
	function I(e, t) {
		var n, i = 0,
		r = {
			height: e
		};
		for (t = t ? 1 : 0; 4 > i; i += 2 - t) n = kt[i],
		r["margin" + n] = r["padding" + n] = e;
		return t && (r.opacity = r.width = e),
		r
	}
	function N(e, t, n) {
		for (var i, r = (nn[t] || []).concat(nn["*"]), o = 0, a = r.length; a > o; o++) if (i = r[o].call(n, t, e)) return i
	}
	function A(e, t, n) {
		var i, r, o, a, s, l, c, u, d = this,
		f = {},
		p = e.style,
		h = e.nodeType && Ct(e),
		g = vt.get(e, "fxshow");
		n.queue || (s = Z._queueHooks(e, "fx"), null == s.unqueued && (s.unqueued = 0, l = s.empty.fire, s.empty.fire = function() {
			s.unqueued || l()
		}), s.unqueued++, d.always(function() {
			d.always(function() {
				s.unqueued--,
				Z.queue(e, "fx").length || s.empty.fire()
			})
		})),
		1 === e.nodeType && ("height" in t || "width" in t) && (n.overflow = [p.overflow, p.overflowX, p.overflowY], c = Z.css(e, "display"), u = "none" === c ? vt.get(e, "olddisplay") || w(e.nodeName) : c, "inline" === u && "none" === Z.css(e, "float") && (p.display = "inline-block")),
		n.overflow && (p.overflow = "hidden", d.always(function() {
			p.overflow = n.overflow[0],
			p.overflowX = n.overflow[1],
			p.overflowY = n.overflow[2]
		}));
		for (i in t) if (r = t[i], Gt.exec(r)) {
			if (delete t[i], o = o || "toggle" === r, r === (h ? "hide": "show")) {
				if ("show" !== r || !g || void 0 === g[i]) continue;
				h = !0
			}
			f[i] = g && g[i] || Z.style(e, i)
		} else c = void 0;
		if (Z.isEmptyObject(f))"inline" === ("none" === c ? w(e.nodeName) : c) && (p.display = c);
		else {
			g ? "hidden" in g && (h = g.hidden) : g = vt.access(e, "fxshow", {}),
			o && (g.hidden = !h),
			h ? Z(e).show() : d.done(function() {
				Z(e).hide()
			}),
			d.done(function() {
				var t;
				vt.remove(e, "fxshow");
				for (t in f) Z.style(e, t, f[t])
			});
			for (i in f) a = N(h ? g[i] : 0, i, d),
			i in g || (g[i] = a.start, h && (a.end = a.start, a.start = "width" === i || "height" === i ? 1 : 0))
		}
	}
	function D(e, t) {
		var n, i, r, o, a;
		for (n in e) if (i = Z.camelCase(n), r = t[i], o = e[n], Z.isArray(o) && (r = o[1], o = e[n] = o[0]), n !== i && (e[i] = o, delete e[n]), a = Z.cssHooks[i], a && "expand" in a) {
			o = a.expand(o),
			delete e[i];
			for (n in o) n in e || (e[n] = o[n], t[n] = r)
		} else t[i] = r
	}
	function j(e, t, n) {
		var i, r, o = 0,
		a = tn.length,
		s = Z.Deferred().always(function() {
			delete l.elem
		}),
		l = function() {
			if (r) return ! 1;
			for (var t = Qt || E(), n = Math.max(0, c.startTime + c.duration - t), i = n / c.duration || 0, o = 1 - i, a = 0, l = c.tweens.length; l > a; a++) c.tweens[a].run(o);
			return s.notifyWith(e, [c, o, n]),
			1 > o && l ? n: (s.resolveWith(e, [c]), !1)
		},
		c = s.promise({
			elem: e,
			props: Z.extend({},
			t),
			opts: Z.extend(!0, {
				specialEasing: {}
			},
			n),
			originalProperties: t,
			originalOptions: n,
			startTime: Qt || E(),
			duration: n.duration,
			tweens: [],
			createTween: function(t, n) {
				var i = Z.Tween(e, c.opts, t, n, c.opts.specialEasing[t] || c.opts.easing);
				return c.tweens.push(i),
				i
			},
			stop: function(t) {
				var n = 0,
				i = t ? c.tweens.length: 0;
				if (r) return this;
				for (r = !0; i > n; n++) c.tweens[n].run(1);
				return t ? s.resolveWith(e, [c, t]) : s.rejectWith(e, [c, t]),
				this
			}
		}),
		u = c.props;
		for (D(u, c.opts.specialEasing); a > o; o++) if (i = tn[o].call(c, e, u, c.opts)) return i;
		return Z.map(u, N, c),
		Z.isFunction(c.opts.start) && c.opts.start.call(e, c),
		Z.fx.timer(Z.extend(l, {
			elem: e,
			anim: c,
			queue: c.opts.queue
		})),
		c.progress(c.opts.progress).done(c.opts.done, c.opts.complete).fail(c.opts.fail).always(c.opts.always)
	}
	function H(e) {
		return function(t, n) {
			"string" != typeof t && (n = t, t = "*");
			var i, r = 0,
			o = t.toLowerCase().match(pt) || [];
			if (Z.isFunction(n)) for (; i = o[r++];)"+" === i[0] ? (i = i.slice(1) || "*", (e[i] = e[i] || []).unshift(n)) : (e[i] = e[i] || []).push(n)
		}
	}
	function M(e, t, n, i) {
		function r(s) {
			var l;
			return o[s] = !0,
			Z.each(e[s] || [],
			function(e, s) {
				var c = s(t, n, i);
				return "string" != typeof c || a || o[c] ? a ? !(l = c) : void 0 : (t.dataTypes.unshift(c), r(c), !1)
			}),
			l
		}
		var o = {},
		a = e === wn;
		return r(t.dataTypes[0]) || !o["*"] && r("*")
	}
	function P(e, t) {
		var n, i, r = Z.ajaxSettings.flatOptions || {};
		for (n in t) void 0 !== t[n] && ((r[n] ? e: i || (i = {}))[n] = t[n]);
		return i && Z.extend(!0, e, i),
		e
	}
	function W(e, t, n) {
		for (var i, r, o, a, s = e.contents,
		l = e.dataTypes;
		"*" === l[0];) l.shift(),
		void 0 === i && (i = e.mimeType || t.getResponseHeader("Content-Type"));
		if (i) for (r in s) if (s[r] && s[r].test(i)) {
			l.unshift(r);
			break
		}
		if (l[0] in n) o = l[0];
		else {
			for (r in n) {
				if (!l[0] || e.converters[r + " " + l[0]]) {
					o = r;
					break
				}
				a || (a = r)
			}
			o = o || a
		}
		return o ? (o !== l[0] && l.unshift(o), n[o]) : void 0
	}
	function O(e, t, n, i) {
		var r, o, a, s, l, c = {},
		u = e.dataTypes.slice();
		if (u[1]) for (a in e.converters) c[a.toLowerCase()] = e.converters[a];
		for (o = u.shift(); o;) if (e.responseFields[o] && (n[e.responseFields[o]] = t), !l && i && e.dataFilter && (t = e.dataFilter(t, e.dataType)), l = o, o = u.shift()) if ("*" === o) o = l;
		else if ("*" !== l && l !== o) {
			if (a = c[l + " " + o] || c["* " + o], !a) for (r in c) if (s = r.split(" "), s[1] === o && (a = c[l + " " + s[0]] || c["* " + s[0]])) {
				a === !0 ? a = c[r] : c[r] !== !0 && (o = s[0], u.unshift(s[1]));
				break
			}
			if (a !== !0) if (a && e["throws"]) t = a(t);
			else try {
				t = a(t)
			} catch(d) {
				return {
					state: "parsererror",
					error: a ? d: "No conversion from " + l + " to " + o
				}
			}
		}
		return {
			state: "success",
			data: t
		}
	}
	function R(e, t, n, i) {
		var r;
		if (Z.isArray(t)) Z.each(t,
		function(t, r) {
			n || Sn.test(e) ? i(e, r) : R(e + "[" + ("object" == typeof r ? t: "") + "]", r, n, i)
		});
		else if (n || "object" !== Z.type(t)) i(e, t);
		else for (r in t) R(e + "[" + r + "]", t[r], n, i)
	}
	function q(e) {
		return Z.isWindow(e) ? e: 9 === e.nodeType && e.defaultView
	}
	var X = [],
	F = X.slice,
	B = X.concat,
	Y = X.push,
	z = X.indexOf,
	J = {},
	U = J.toString,
	K = J.hasOwnProperty,
	Q = {},
	V = e.document,
	G = "2.1.3",
	Z = function(e, t) {
		return new Z.fn.init(e, t)
	},
	et = /^[\s\uFEFF\xA0]+|[\s\uFEFF\xA0]+$/g,
	tt = /^-ms-/,
	nt = /-([\da-z])/gi,
	it = function(e, t) {
		return t.toUpperCase()
	};
	Z.fn = Z.prototype = {
		jquery: G,
		constructor: Z,
		selector: "",
		length: 0,
		toArray: function() {
			return F.call(this)
		},
		get: function(e) {
			return null != e ? 0 > e ? this[e + this.length] : this[e] : F.call(this)
		},
		pushStack: function(e) {
			var t = Z.merge(this.constructor(), e);
			return t.prevObject = this,
			t.context = this.context,
			t
		},
		each: function(e, t) {
			return Z.each(this, e, t)
		},
		map: function(e) {
			return this.pushStack(Z.map(this,
			function(t, n) {
				return e.call(t, n, t)
			}))
		},
		slice: function() {
			return this.pushStack(F.apply(this, arguments))
		},
		first: function() {
			return this.eq(0)
		},
		last: function() {
			return this.eq( - 1)
		},
		eq: function(e) {
			var t = this.length,
			n = +e + (0 > e ? t: 0);
			return this.pushStack(n >= 0 && t > n ? [this[n]] : [])
		},
		end: function() {
			return this.prevObject || this.constructor(null)
		},
		push: Y,
		sort: X.sort,
		splice: X.splice
	},
	Z.extend = Z.fn.extend = function() {
		var e, t, n, i, r, o, a = arguments[0] || {},
		s = 1,
		l = arguments.length,
		c = !1;
		for ("boolean" == typeof a && (c = a, a = arguments[s] || {},
		s++), "object" == typeof a || Z.isFunction(a) || (a = {}), s === l && (a = this, s--); l > s; s++) if (null != (e = arguments[s])) for (t in e) n = a[t],
		i = e[t],
		a !== i && (c && i && (Z.isPlainObject(i) || (r = Z.isArray(i))) ? (r ? (r = !1, o = n && Z.isArray(n) ? n: []) : o = n && Z.isPlainObject(n) ? n: {},
		a[t] = Z.extend(c, o, i)) : void 0 !== i && (a[t] = i));
		return a
	},
	Z.extend({
		expando: "jQuery" + (G + Math.random()).replace(/\D/g, ""),
		isReady: !0,
		error: function(e) {
			throw new Error(e)
		},
		noop: function() {},
		isFunction: function(e) {
			return "function" === Z.type(e)
		},
		isArray: Array.isArray,
		isWindow: function(e) {
			return null != e && e === e.window
		},
		isNumeric: function(e) {
			return ! Z.isArray(e) && e - parseFloat(e) + 1 >= 0
		},
		isPlainObject: function(e) {
			return "object" !== Z.type(e) || e.nodeType || Z.isWindow(e) ? !1 : e.constructor && !K.call(e.constructor.prototype, "isPrototypeOf") ? !1 : !0
		},
		isEmptyObject: function(e) {
			var t;
			for (t in e) return ! 1;
			return ! 0
		},
		type: function(e) {
			return null == e ? e + "": "object" == typeof e || "function" == typeof e ? J[U.call(e)] || "object": typeof e
		},
		globalEval: function(e) {
			var t, n = eval;
			e = Z.trim(e),
			e && (1 === e.indexOf("use strict") ? (t = V.createElement("script"), t.text = e, V.head.appendChild(t).parentNode.removeChild(t)) : n(e))
		},
		camelCase: function(e) {
			return e.replace(tt, "ms-").replace(nt, it)
		},
		nodeName: function(e, t) {
			return e.nodeName && e.nodeName.toLowerCase() === t.toLowerCase()
		},
		each: function(e, t, i) {
			var r, o = 0,
			a = e.length,
			s = n(e);
			if (i) {
				if (s) for (; a > o && (r = t.apply(e[o], i), r !== !1); o++);
				else for (o in e) if (r = t.apply(e[o], i), r === !1) break
			} else if (s) for (; a > o && (r = t.call(e[o], o, e[o]), r !== !1); o++);
			else for (o in e) if (r = t.call(e[o], o, e[o]), r === !1) break;
			return e
		},
		trim: function(e) {
			return null == e ? "": (e + "").replace(et, "")
		},
		makeArray: function(e, t) {
			var i = t || [];
			return null != e && (n(Object(e)) ? Z.merge(i, "string" == typeof e ? [e] : e) : Y.call(i, e)),
			i
		},
		inArray: function(e, t, n) {
			return null == t ? -1 : z.call(t, e, n)
		},
		merge: function(e, t) {
			for (var n = +t.length,
			i = 0,
			r = e.length; n > i; i++) e[r++] = t[i];
			return e.length = r,
			e
		},
		grep: function(e, t, n) {
			for (var i, r = [], o = 0, a = e.length, s = !n; a > o; o++) i = !t(e[o], o),
			i !== s && r.push(e[o]);
			return r
		},
		map: function(e, t, i) {
			var r, o = 0,
			a = e.length,
			s = n(e),
			l = [];
			if (s) for (; a > o; o++) r = t(e[o], o, i),
			null != r && l.push(r);
			else for (o in e) r = t(e[o], o, i),
			null != r && l.push(r);
			return B.apply([], l)
		},
		guid: 1,
		proxy: function(e, t) {
			var n, i, r;
			return "string" == typeof t && (n = e[t], t = e, e = n),
			Z.isFunction(e) ? (i = F.call(arguments, 2), r = function() {
				return e.apply(t || this, i.concat(F.call(arguments)))
			},
			r.guid = e.guid = e.guid || Z.guid++, r) : void 0
		},
		now: Date.now,
		support: Q
	}),
	Z.each("Boolean Number String Function Array Date RegExp Object Error".split(" "),
	function(e, t) {
		J["[object " + t + "]"] = t.toLowerCase()
	});
	var rt = function(e) {
		function t(e, t, n, i) {
			var r, o, a, s, l, c, d, p, h, g;
			if ((t ? t.ownerDocument || t: R) !== A && N(t), t = t || A, n = n || [], s = t.nodeType, "string" != typeof e || !e || 1 !== s && 9 !== s && 11 !== s) return n;
			if (!i && j) {
				if (11 !== s && (r = yt.exec(e))) if (a = r[1]) {
					if (9 === s) {
						if (o = t.getElementById(a), !o || !o.parentNode) return n;
						if (o.id === a) return n.push(o),
						n
					} else if (t.ownerDocument && (o = t.ownerDocument.getElementById(a)) && W(t, o) && o.id === a) return n.push(o),
					n
				} else {
					if (r[2]) return G.apply(n, t.getElementsByTagName(e)),
					n;
					if ((a = r[3]) && x.getElementsByClassName) return G.apply(n, t.getElementsByClassName(a)),
					n
				}
				if (x.qsa && (!H || !H.test(e))) {
					if (p = d = O, h = t, g = 1 !== s && e, 1 === s && "object" !== t.nodeName.toLowerCase()) {
						for (c = S(e), (d = t.getAttribute("id")) ? p = d.replace(wt, "\\$&") : t.setAttribute("id", p), p = "[id='" + p + "'] ", l = c.length; l--;) c[l] = p + f(c[l]);
						h = bt.test(e) && u(t.parentNode) || t,
						g = c.join(",")
					}
					if (g) try {
						return G.apply(n, h.querySelectorAll(g)),
						n
					} catch(m) {} finally {
						d || t.removeAttribute("id")
					}
				}
			}
			return $(e.replace(lt, "$1"), t, n, i)
		}
		function n() {
			function e(n, i) {
				return t.push(n + " ") > k.cacheLength && delete e[t.shift()],
				e[n + " "] = i
			}
			var t = [];
			return e
		}
		function i(e) {
			return e[O] = !0,
			e
		}
		function r(e) {
			var t = A.createElement("div");
			try {
				return !! e(t)
			} catch(n) {
				return ! 1
			} finally {
				t.parentNode && t.parentNode.removeChild(t),
				t = null
			}
		}
		function o(e, t) {
			for (var n = e.split("|"), i = e.length; i--;) k.attrHandle[n[i]] = t
		}
		function a(e, t) {
			var n = t && e,
			i = n && 1 === e.nodeType && 1 === t.nodeType && (~t.sourceIndex || J) - (~e.sourceIndex || J);
			if (i) return i;
			if (n) for (; n = n.nextSibling;) if (n === t) return - 1;
			return e ? 1 : -1
		}
		function s(e) {
			return function(t) {
				var n = t.nodeName.toLowerCase();
				return "input" === n && t.type === e
			}
		}
		function l(e) {
			return function(t) {
				var n = t.nodeName.toLowerCase();
				return ("input" === n || "button" === n) && t.type === e
			}
		}
		function c(e) {
			return i(function(t) {
				return t = +t,
				i(function(n, i) {
					for (var r, o = e([], n.length, t), a = o.length; a--;) n[r = o[a]] && (n[r] = !(i[r] = n[r]))
				})
			})
		}
		function u(e) {
			return e && "undefined" != typeof e.getElementsByTagName && e
		}
		function d() {}
		function f(e) {
			for (var t = 0,
			n = e.length,
			i = ""; n > t; t++) i += e[t].value;
			return i
		}
		function p(e, t, n) {
			var i = t.dir,
			r = n && "parentNode" === i,
			o = X++;
			return t.first ?
			function(t, n, o) {
				for (; t = t[i];) if (1 === t.nodeType || r) return e(t, n, o)
			}: function(t, n, a) {
				var s, l, c = [q, o];
				if (a) {
					for (; t = t[i];) if ((1 === t.nodeType || r) && e(t, n, a)) return ! 0
				} else for (; t = t[i];) if (1 === t.nodeType || r) {
					if (l = t[O] || (t[O] = {}), (s = l[i]) && s[0] === q && s[1] === o) return c[2] = s[2];
					if (l[i] = c, c[2] = e(t, n, a)) return ! 0
				}
			}
		}
		function h(e) {
			return e.length > 1 ?
			function(t, n, i) {
				for (var r = e.length; r--;) if (!e[r](t, n, i)) return ! 1;
				return ! 0
			}: e[0]
		}
		function g(e, n, i) {
			for (var r = 0,
			o = n.length; o > r; r++) t(e, n[r], i);
			return i
		}
		function m(e, t, n, i, r) {
			for (var o, a = [], s = 0, l = e.length, c = null != t; l > s; s++)(o = e[s]) && (!n || n(o, i, r)) && (a.push(o), c && t.push(s));
			return a
		}
		function v(e, t, n, r, o, a) {
			return r && !r[O] && (r = v(r)),
			o && !o[O] && (o = v(o, a)),
			i(function(i, a, s, l) {
				var c, u, d, f = [],
				p = [],
				h = a.length,
				v = i || g(t || "*", s.nodeType ? [s] : s, []),
				y = !e || !i && t ? v: m(v, f, e, s, l),
				b = n ? o || (i ? e: h || r) ? [] : a: y;
				if (n && n(y, b, s, l), r) for (c = m(b, p), r(c, [], s, l), u = c.length; u--;)(d = c[u]) && (b[p[u]] = !(y[p[u]] = d));
				if (i) {
					if (o || e) {
						if (o) {
							for (c = [], u = b.length; u--;)(d = b[u]) && c.push(y[u] = d);
							o(null, b = [], c, l)
						}
						for (u = b.length; u--;)(d = b[u]) && (c = o ? et(i, d) : f[u]) > -1 && (i[c] = !(a[c] = d))
					}
				} else b = m(b === a ? b.splice(h, b.length) : b),
				o ? o(null, a, b, l) : G.apply(a, b)
			})
		}
		function y(e) {
			for (var t, n, i, r = e.length,
			o = k.relative[e[0].type], a = o || k.relative[" "], s = o ? 1 : 0, l = p(function(e) {
				return e === t
			},
			a, !0), c = p(function(e) {
				return et(t, e) > -1
			},
			a, !0), u = [function(e, n, i) {
				var r = !o && (i || n !== L) || ((t = n).nodeType ? l(e, n, i) : c(e, n, i));
				return t = null,
				r
			}]; r > s; s++) if (n = k.relative[e[s].type]) u = [p(h(u), n)];
			else {
				if (n = k.filter[e[s].type].apply(null, e[s].matches), n[O]) {
					for (i = ++s; r > i && !k.relative[e[i].type]; i++);
					return v(s > 1 && h(u), s > 1 && f(e.slice(0, s - 1).concat({
						value: " " === e[s - 2].type ? "*": ""
					})).replace(lt, "$1"), n, i > s && y(e.slice(s, i)), r > i && y(e = e.slice(i)), r > i && f(e))
				}
				u.push(n)
			}
			return h(u)
		}
		function b(e, n) {
			var r = n.length > 0,
			o = e.length > 0,
			a = function(i, a, s, l, c) {
				var u, d, f, p = 0,
				h = "0",
				g = i && [],
				v = [],
				y = L,
				b = i || o && k.find.TAG("*", c),
				w = q += null == y ? 1 : Math.random() || .1,
				x = b.length;
				for (c && (L = a !== A && a); h !== x && null != (u = b[h]); h++) {
					if (o && u) {
						for (d = 0; f = e[d++];) if (f(u, a, s)) {
							l.push(u);
							break
						}
						c && (q = w)
					}
					r && ((u = !f && u) && p--, i && g.push(u))
				}
				if (p += h, r && h !== p) {
					for (d = 0; f = n[d++];) f(g, v, a, s);
					if (i) {
						if (p > 0) for (; h--;) g[h] || v[h] || (v[h] = Q.call(l));
						v = m(v)
					}
					G.apply(l, v),
					c && !i && v.length > 0 && p + n.length > 1 && t.uniqueSort(l)
				}
				return c && (q = w, L = y),
				g
			};
			return r ? i(a) : a
		}
		var w, x, k, C, T, S, _, $, L, E, I, N, A, D, j, H, M, P, W, O = "sizzle" + 1 * new Date,
		R = e.document,
		q = 0,
		X = 0,
		F = n(),
		B = n(),
		Y = n(),
		z = function(e, t) {
			return e === t && (I = !0),
			0
		},
		J = 1 << 31,
		U = {}.hasOwnProperty,
		K = [],
		Q = K.pop,
		V = K.push,
		G = K.push,
		Z = K.slice,
		et = function(e, t) {
			for (var n = 0,
			i = e.length; i > n; n++) if (e[n] === t) return n;
			return - 1
		},
		tt = "checked|selected|async|autofocus|autoplay|controls|defer|disabled|hidden|ismap|loop|multiple|open|readonly|required|scoped",
		nt = "[\\x20\\t\\r\\n\\f]",
		it = "(?:\\\\.|[\\w-]|[^\\x00-\\xa0])+",
		rt = it.replace("w", "w#"),
		ot = "\\[" + nt + "*(" + it + ")(?:" + nt + "*([*^$|!~]?=)" + nt + "*(?:'((?:\\\\.|[^\\\\'])*)'|\"((?:\\\\.|[^\\\\\"])*)\"|(" + rt + "))|)" + nt + "*\\]",
		at = ":(" + it + ")(?:\\((('((?:\\\\.|[^\\\\'])*)'|\"((?:\\\\.|[^\\\\\"])*)\")|((?:\\\\.|[^\\\\()[\\]]|" + ot + ")*)|.*)\\)|)",
		st = new RegExp(nt + "+", "g"),
		lt = new RegExp("^" + nt + "+|((?:^|[^\\\\])(?:\\\\.)*)" + nt + "+$", "g"),
		ct = new RegExp("^" + nt + "*," + nt + "*"),
		ut = new RegExp("^" + nt + "*([>+~]|" + nt + ")" + nt + "*"),
		dt = new RegExp("=" + nt + "*([^\\]'\"]*?)" + nt + "*\\]", "g"),
		ft = new RegExp(at),
		pt = new RegExp("^" + rt + "$"),
		ht = {
			ID: new RegExp("^#(" + it + ")"),
			CLASS: new RegExp("^\\.(" + it + ")"),
			TAG: new RegExp("^(" + it.replace("w", "w*") + ")"),
			ATTR: new RegExp("^" + ot),
			PSEUDO: new RegExp("^" + at),
			CHILD: new RegExp("^:(only|first|last|nth|nth-last)-(child|of-type)(?:\\(" + nt + "*(even|odd|(([+-]|)(\\d*)n|)" + nt + "*(?:([+-]|)" + nt + "*(\\d+)|))" + nt + "*\\)|)", "i"),
			bool: new RegExp("^(?:" + tt + ")$", "i"),
			needsContext: new RegExp("^" + nt + "*[>+~]|:(even|odd|eq|gt|lt|nth|first|last)(?:\\(" + nt + "*((?:-\\d)?\\d*)" + nt + "*\\)|)(?=[^-]|$)", "i")
		},
		gt = /^(?:input|select|textarea|button)$/i,
		mt = /^h\d$/i,
		vt = /^[^{]+\{\s*\[native \w/,
		yt = /^(?:#([\w-]+)|(\w+)|\.([\w-]+))$/,
		bt = /[+~]/,
		wt = /'|\\/g,
		xt = new RegExp("\\\\([\\da-f]{1,6}" + nt + "?|(" + nt + ")|.)", "ig"),
		kt = function(e, t, n) {
			var i = "0x" + t - 65536;
			return i !== i || n ? t: 0 > i ? String.fromCharCode(i + 65536) : String.fromCharCode(i >> 10 | 55296, 1023 & i | 56320)
		},
		Ct = function() {
			N()
		};
		try {
			G.apply(K = Z.call(R.childNodes), R.childNodes),
			K[R.childNodes.length].nodeType
		} catch(Tt) {
			G = {
				apply: K.length ?
				function(e, t) {
					V.apply(e, Z.call(t))
				}: function(e, t) {
					for (var n = e.length,
					i = 0; e[n++] = t[i++];);
					e.length = n - 1
				}
			}
		}
		x = t.support = {},
		T = t.isXML = function(e) {
			var t = e && (e.ownerDocument || e).documentElement;
			return t ? "HTML" !== t.nodeName: !1
		},
		N = t.setDocument = function(e) {
			var t, n, i = e ? e.ownerDocument || e: R;
			return i !== A && 9 === i.nodeType && i.documentElement ? (A = i, D = i.documentElement, n = i.defaultView, n && n !== n.top && (n.addEventListener ? n.addEventListener("unload", Ct, !1) : n.attachEvent && n.attachEvent("onunload", Ct)), j = !T(i), x.attributes = r(function(e) {
				return e.className = "i",
				!e.getAttribute("className")
			}), x.getElementsByTagName = r(function(e) {
				return e.appendChild(i.createComment("")),
				!e.getElementsByTagName("*").length
			}), x.getElementsByClassName = vt.test(i.getElementsByClassName), x.getById = r(function(e) {
				return D.appendChild(e).id = O,
				!i.getElementsByName || !i.getElementsByName(O).length
			}), x.getById ? (k.find.ID = function(e, t) {
				if ("undefined" != typeof t.getElementById && j) {
					var n = t.getElementById(e);
					return n && n.parentNode ? [n] : []
				}
			},
			k.filter.ID = function(e) {
				var t = e.replace(xt, kt);
				return function(e) {
					return e.getAttribute("id") === t
				}
			}) : (delete k.find.ID, k.filter.ID = function(e) {
				var t = e.replace(xt, kt);
				return function(e) {
					var n = "undefined" != typeof e.getAttributeNode && e.getAttributeNode("id");
					return n && n.value === t
				}
			}), k.find.TAG = x.getElementsByTagName ?
			function(e, t) {
				return "undefined" != typeof t.getElementsByTagName ? t.getElementsByTagName(e) : x.qsa ? t.querySelectorAll(e) : void 0
			}: function(e, t) {
				var n, i = [],
				r = 0,
				o = t.getElementsByTagName(e);
				if ("*" === e) {
					for (; n = o[r++];) 1 === n.nodeType && i.push(n);
					return i
				}
				return o
			},
			k.find.CLASS = x.getElementsByClassName &&
			function(e, t) {
				return j ? t.getElementsByClassName(e) : void 0
			},
			M = [], H = [], (x.qsa = vt.test(i.querySelectorAll)) && (r(function(e) {
				D.appendChild(e).innerHTML = "<a id='" + O + "'></a><select id='" + O + "-\f]' msallowcapture=''><option selected=''></option></select>",
				e.querySelectorAll("[msallowcapture^='']").length && H.push("[*^$]=" + nt + "*(?:''|\"\")"),
				e.querySelectorAll("[selected]").length || H.push("\\[" + nt + "*(?:value|" + tt + ")"),
				e.querySelectorAll("[id~=" + O + "-]").length || H.push("~="),
				e.querySelectorAll(":checked").length || H.push(":checked"),
				e.querySelectorAll("a#" + O + "+*").length || H.push(".#.+[+~]")
			}), r(function(e) {
				var t = i.createElement("input");
				t.setAttribute("type", "hidden"),
				e.appendChild(t).setAttribute("name", "D"),
				e.querySelectorAll("[name=d]").length && H.push("name" + nt + "*[*^$|!~]?="),
				e.querySelectorAll(":enabled").length || H.push(":enabled", ":disabled"),
				e.querySelectorAll("*,:x"),
				H.push(",.*:")
			})), (x.matchesSelector = vt.test(P = D.matches || D.webkitMatchesSelector || D.mozMatchesSelector || D.oMatchesSelector || D.msMatchesSelector)) && r(function(e) {
				x.disconnectedMatch = P.call(e, "div"),
				P.call(e, "[s!='']:x"),
				M.push("!=", at)
			}), H = H.length && new RegExp(H.join("|")), M = M.length && new RegExp(M.join("|")), t = vt.test(D.compareDocumentPosition), W = t || vt.test(D.contains) ?
			function(e, t) {
				var n = 9 === e.nodeType ? e.documentElement: e,
				i = t && t.parentNode;
				return e === i || !(!i || 1 !== i.nodeType || !(n.contains ? n.contains(i) : e.compareDocumentPosition && 16 & e.compareDocumentPosition(i)))
			}: function(e, t) {
				if (t) for (; t = t.parentNode;) if (t === e) return ! 0;
				return ! 1
			},
			z = t ?
			function(e, t) {
				if (e === t) return I = !0,
				0;
				var n = !e.compareDocumentPosition - !t.compareDocumentPosition;
				return n ? n: (n = (e.ownerDocument || e) === (t.ownerDocument || t) ? e.compareDocumentPosition(t) : 1, 1 & n || !x.sortDetached && t.compareDocumentPosition(e) === n ? e === i || e.ownerDocument === R && W(R, e) ? -1 : t === i || t.ownerDocument === R && W(R, t) ? 1 : E ? et(E, e) - et(E, t) : 0 : 4 & n ? -1 : 1)
			}: function(e, t) {
				if (e === t) return I = !0,
				0;
				var n, r = 0,
				o = e.parentNode,
				s = t.parentNode,
				l = [e],
				c = [t];
				if (!o || !s) return e === i ? -1 : t === i ? 1 : o ? -1 : s ? 1 : E ? et(E, e) - et(E, t) : 0;
				if (o === s) return a(e, t);
				for (n = e; n = n.parentNode;) l.unshift(n);
				for (n = t; n = n.parentNode;) c.unshift(n);
				for (; l[r] === c[r];) r++;
				return r ? a(l[r], c[r]) : l[r] === R ? -1 : c[r] === R ? 1 : 0
			},
			i) : A
		},
		t.matches = function(e, n) {
			return t(e, null, null, n)
		},
		t.matchesSelector = function(e, n) {
			if ((e.ownerDocument || e) !== A && N(e), n = n.replace(dt, "='$1']"), !(!x.matchesSelector || !j || M && M.test(n) || H && H.test(n))) try {
				var i = P.call(e, n);
				if (i || x.disconnectedMatch || e.document && 11 !== e.document.nodeType) return i
			} catch(r) {}
			return t(n, A, null, [e]).length > 0
		},
		t.contains = function(e, t) {
			return (e.ownerDocument || e) !== A && N(e),
			W(e, t)
		},
		t.attr = function(e, t) { (e.ownerDocument || e) !== A && N(e);
			var n = k.attrHandle[t.toLowerCase()],
			i = n && U.call(k.attrHandle, t.toLowerCase()) ? n(e, t, !j) : void 0;
			return void 0 !== i ? i: x.attributes || !j ? e.getAttribute(t) : (i = e.getAttributeNode(t)) && i.specified ? i.value: null
		},
		t.error = function(e) {
			throw new Error("Syntax error, unrecognized expression: " + e)
		},
		t.uniqueSort = function(e) {
			var t, n = [],
			i = 0,
			r = 0;
			if (I = !x.detectDuplicates, E = !x.sortStable && e.slice(0), e.sort(z), I) {
				for (; t = e[r++];) t === e[r] && (i = n.push(r));
				for (; i--;) e.splice(n[i], 1)
			}
			return E = null,
			e
		},
		C = t.getText = function(e) {
			var t, n = "",
			i = 0,
			r = e.nodeType;
			if (r) {
				if (1 === r || 9 === r || 11 === r) {
					if ("string" == typeof e.textContent) return e.textContent;
					for (e = e.firstChild; e; e = e.nextSibling) n += C(e)
				} else if (3 === r || 4 === r) return e.nodeValue
			} else for (; t = e[i++];) n += C(t);
			return n
		},
		k = t.selectors = {
			cacheLength: 50,
			createPseudo: i,
			match: ht,
			attrHandle: {},
			find: {},
			relative: {
				">": {
					dir: "parentNode",
					first: !0
				},
				" ": {
					dir: "parentNode"
				},
				"+": {
					dir: "previousSibling",
					first: !0
				},
				"~": {
					dir: "previousSibling"
				}
			},
			preFilter: {
				ATTR: function(e) {
					return e[1] = e[1].replace(xt, kt),
					e[3] = (e[3] || e[4] || e[5] || "").replace(xt, kt),
					"~=" === e[2] && (e[3] = " " + e[3] + " "),
					e.slice(0, 4)
				},
				CHILD: function(e) {
					return e[1] = e[1].toLowerCase(),
					"nth" === e[1].slice(0, 3) ? (e[3] || t.error(e[0]), e[4] = +(e[4] ? e[5] + (e[6] || 1) : 2 * ("even" === e[3] || "odd" === e[3])), e[5] = +(e[7] + e[8] || "odd" === e[3])) : e[3] && t.error(e[0]),
					e
				},
				PSEUDO: function(e) {
					var t, n = !e[6] && e[2];
					return ht.CHILD.test(e[0]) ? null: (e[3] ? e[2] = e[4] || e[5] || "": n && ft.test(n) && (t = S(n, !0)) && (t = n.indexOf(")", n.length - t) - n.length) && (e[0] = e[0].slice(0, t), e[2] = n.slice(0, t)), e.slice(0, 3))
				}
			},
			filter: {
				TAG: function(e) {
					var t = e.replace(xt, kt).toLowerCase();
					return "*" === e ?
					function() {
						return ! 0
					}: function(e) {
						return e.nodeName && e.nodeName.toLowerCase() === t
					}
				},
				CLASS: function(e) {
					var t = F[e + " "];
					return t || (t = new RegExp("(^|" + nt + ")" + e + "(" + nt + "|$)")) && F(e,
					function(e) {
						return t.test("string" == typeof e.className && e.className || "undefined" != typeof e.getAttribute && e.getAttribute("class") || "")
					})
				},
				ATTR: function(e, n, i) {
					return function(r) {
						var o = t.attr(r, e);
						return null == o ? "!=" === n: n ? (o += "", "=" === n ? o === i: "!=" === n ? o !== i: "^=" === n ? i && 0 === o.indexOf(i) : "*=" === n ? i && o.indexOf(i) > -1 : "$=" === n ? i && o.slice( - i.length) === i: "~=" === n ? (" " + o.replace(st, " ") + " ").indexOf(i) > -1 : "|=" === n ? o === i || o.slice(0, i.length + 1) === i + "-": !1) : !0
					}
				},
				CHILD: function(e, t, n, i, r) {
					var o = "nth" !== e.slice(0, 3),
					a = "last" !== e.slice( - 4),
					s = "of-type" === t;
					return 1 === i && 0 === r ?
					function(e) {
						return !! e.parentNode
					}: function(t, n, l) {
						var c, u, d, f, p, h, g = o !== a ? "nextSibling": "previousSibling",
						m = t.parentNode,
						v = s && t.nodeName.toLowerCase(),
						y = !l && !s;
						if (m) {
							if (o) {
								for (; g;) {
									for (d = t; d = d[g];) if (s ? d.nodeName.toLowerCase() === v: 1 === d.nodeType) return ! 1;
									h = g = "only" === e && !h && "nextSibling"
								}
								return ! 0
							}
							if (h = [a ? m.firstChild: m.lastChild], a && y) {
								for (u = m[O] || (m[O] = {}), c = u[e] || [], p = c[0] === q && c[1], f = c[0] === q && c[2], d = p && m.childNodes[p]; d = ++p && d && d[g] || (f = p = 0) || h.pop();) if (1 === d.nodeType && ++f && d === t) {
									u[e] = [q, p, f];
									break
								}
							} else if (y && (c = (t[O] || (t[O] = {}))[e]) && c[0] === q) f = c[1];
							else for (; (d = ++p && d && d[g] || (f = p = 0) || h.pop()) && ((s ? d.nodeName.toLowerCase() !== v: 1 !== d.nodeType) || !++f || (y && ((d[O] || (d[O] = {}))[e] = [q, f]), d !== t)););
							return f -= r,
							f === i || f % i === 0 && f / i >= 0
						}
					}
				},
				PSEUDO: function(e, n) {
					var r, o = k.pseudos[e] || k.setFilters[e.toLowerCase()] || t.error("unsupported pseudo: " + e);
					return o[O] ? o(n) : o.length > 1 ? (r = [e, e, "", n], k.setFilters.hasOwnProperty(e.toLowerCase()) ? i(function(e, t) {
						for (var i, r = o(e, n), a = r.length; a--;) i = et(e, r[a]),
						e[i] = !(t[i] = r[a])
					}) : function(e) {
						return o(e, 0, r)
					}) : o
				}
			},
			pseudos: {
				not: i(function(e) {
					var t = [],
					n = [],
					r = _(e.replace(lt, "$1"));
					return r[O] ? i(function(e, t, n, i) {
						for (var o, a = r(e, null, i, []), s = e.length; s--;)(o = a[s]) && (e[s] = !(t[s] = o))
					}) : function(e, i, o) {
						return t[0] = e,
						r(t, null, o, n),
						t[0] = null,
						!n.pop()
					}
				}),
				has: i(function(e) {
					return function(n) {
						return t(e, n).length > 0
					}
				}),
				contains: i(function(e) {
					return e = e.replace(xt, kt),
					function(t) {
						return (t.textContent || t.innerText || C(t)).indexOf(e) > -1
					}
				}),
				lang: i(function(e) {
					return pt.test(e || "") || t.error("unsupported lang: " + e),
					e = e.replace(xt, kt).toLowerCase(),
					function(t) {
						var n;
						do
						if (n = j ? t.lang: t.getAttribute("xml:lang") || t.getAttribute("lang")) return n = n.toLowerCase(),
						n === e || 0 === n.indexOf(e + "-");
						while ((t = t.parentNode) && 1 === t.nodeType);
						return ! 1
					}
				}),
				target: function(t) {
					var n = e.location && e.location.hash;
					return n && n.slice(1) === t.id
				},
				root: function(e) {
					return e === D
				},
				focus: function(e) {
					return e === A.activeElement && (!A.hasFocus || A.hasFocus()) && !!(e.type || e.href || ~e.tabIndex)
				},
				enabled: function(e) {
					return e.disabled === !1
				},
				disabled: function(e) {
					return e.disabled === !0
				},
				checked: function(e) {
					var t = e.nodeName.toLowerCase();
					return "input" === t && !!e.checked || "option" === t && !!e.selected
				},
				selected: function(e) {
					return e.parentNode && e.parentNode.selectedIndex,
					e.selected === !0
				},
				empty: function(e) {
					for (e = e.firstChild; e; e = e.nextSibling) if (e.nodeType < 6) return ! 1;
					return ! 0
				},
				parent: function(e) {
					return ! k.pseudos.empty(e)
				},
				header: function(e) {
					return mt.test(e.nodeName)
				},
				input: function(e) {
					return gt.test(e.nodeName)
				},
				button: function(e) {
					var t = e.nodeName.toLowerCase();
					return "input" === t && "button" === e.type || "button" === t
				},
				text: function(e) {
					var t;
					return "input" === e.nodeName.toLowerCase() && "text" === e.type && (null == (t = e.getAttribute("type")) || "text" === t.toLowerCase())
				},
				first: c(function() {
					return [0]
				}),
				last: c(function(e, t) {
					return [t - 1]
				}),
				eq: c(function(e, t, n) {
					return [0 > n ? n + t: n]
				}),
				even: c(function(e, t) {
					for (var n = 0; t > n; n += 2) e.push(n);
					return e
				}),
				odd: c(function(e, t) {
					for (var n = 1; t > n; n += 2) e.push(n);
					return e
				}),
				lt: c(function(e, t, n) {
					for (var i = 0 > n ? n + t: n; --i >= 0;) e.push(i);
					return e
				}),
				gt: c(function(e, t, n) {
					for (var i = 0 > n ? n + t: n; ++i < t;) e.push(i);
					return e
				})
			}
		},
		k.pseudos.nth = k.pseudos.eq;
		for (w in {
			radio: !0,
			checkbox: !0,
			file: !0,
			password: !0,
			image: !0
		}) k.pseudos[w] = s(w);
		for (w in {
			submit: !0,
			reset: !0
		}) k.pseudos[w] = l(w);
		return d.prototype = k.filters = k.pseudos,
		k.setFilters = new d,
		S = t.tokenize = function(e, n) {
			var i, r, o, a, s, l, c, u = B[e + " "];
			if (u) return n ? 0 : u.slice(0);
			for (s = e, l = [], c = k.preFilter; s;) { (!i || (r = ct.exec(s))) && (r && (s = s.slice(r[0].length) || s), l.push(o = [])),
				i = !1,
				(r = ut.exec(s)) && (i = r.shift(), o.push({
					value: i,
					type: r[0].replace(lt, " ")
				}), s = s.slice(i.length));
				for (a in k.filter) ! (r = ht[a].exec(s)) || c[a] && !(r = c[a](r)) || (i = r.shift(), o.push({
					value: i,
					type: a,
					matches: r
				}), s = s.slice(i.length));
				if (!i) break
			}
			return n ? s.length: s ? t.error(e) : B(e, l).slice(0)
		},
		_ = t.compile = function(e, t) {
			var n, i = [],
			r = [],
			o = Y[e + " "];
			if (!o) {
				for (t || (t = S(e)), n = t.length; n--;) o = y(t[n]),
				o[O] ? i.push(o) : r.push(o);
				o = Y(e, b(r, i)),
				o.selector = e
			}
			return o
		},
		$ = t.select = function(e, t, n, i) {
			var r, o, a, s, l, c = "function" == typeof e && e,
			d = !i && S(e = c.selector || e);
			if (n = n || [], 1 === d.length) {
				if (o = d[0] = d[0].slice(0), o.length > 2 && "ID" === (a = o[0]).type && x.getById && 9 === t.nodeType && j && k.relative[o[1].type]) {
					if (t = (k.find.ID(a.matches[0].replace(xt, kt), t) || [])[0], !t) return n;
					c && (t = t.parentNode),
					e = e.slice(o.shift().value.length)
				}
				for (r = ht.needsContext.test(e) ? 0 : o.length; r--&&(a = o[r], !k.relative[s = a.type]);) if ((l = k.find[s]) && (i = l(a.matches[0].replace(xt, kt), bt.test(o[0].type) && u(t.parentNode) || t))) {
					if (o.splice(r, 1), e = i.length && f(o), !e) return G.apply(n, i),
					n;
					break
				}
			}
			return (c || _(e, d))(i, t, !j, n, bt.test(e) && u(t.parentNode) || t),
			n
		},
		x.sortStable = O.split("").sort(z).join("") === O,
		x.detectDuplicates = !!I,
		N(),
		x.sortDetached = r(function(e) {
			return 1 & e.compareDocumentPosition(A.createElement("div"))
		}),
		r(function(e) {
			return e.innerHTML = "<a href='#'></a>",
			"#" === e.firstChild.getAttribute("href")
		}) || o("type|href|height|width",
		function(e, t, n) {
			return n ? void 0 : e.getAttribute(t, "type" === t.toLowerCase() ? 1 : 2)
		}),
		x.attributes && r(function(e) {
			return e.innerHTML = "<input/>",
			e.firstChild.setAttribute("value", ""),
			"" === e.firstChild.getAttribute("value")
		}) || o("value",
		function(e, t, n) {
			return n || "input" !== e.nodeName.toLowerCase() ? void 0 : e.defaultValue
		}),
		r(function(e) {
			return null == e.getAttribute("disabled")
		}) || o(tt,
		function(e, t, n) {
			var i;
			return n ? void 0 : e[t] === !0 ? t.toLowerCase() : (i = e.getAttributeNode(t)) && i.specified ? i.value: null
		}),
		t
	} (e);
	Z.find = rt,
	Z.expr = rt.selectors,
	Z.expr[":"] = Z.expr.pseudos,
	Z.unique = rt.uniqueSort,
	Z.text = rt.getText,
	Z.isXMLDoc = rt.isXML,
	Z.contains = rt.contains;
	var ot = Z.expr.match.needsContext,
	at = /^<(\w+)\s*\/?>(?:<\/\1>|)$/,
	st = /^.[^:#\[\.,]*$/;
	Z.filter = function(e, t, n) {
		var i = t[0];
		return n && (e = ":not(" + e + ")"),
		1 === t.length && 1 === i.nodeType ? Z.find.matchesSelector(i, e) ? [i] : [] : Z.find.matches(e, Z.grep(t,
		function(e) {
			return 1 === e.nodeType
		}))
	},
	Z.fn.extend({
		find: function(e) {
			var t, n = this.length,
			i = [],
			r = this;
			if ("string" != typeof e) return this.pushStack(Z(e).filter(function() {
				for (t = 0; n > t; t++) if (Z.contains(r[t], this)) return ! 0
			}));
			for (t = 0; n > t; t++) Z.find(e, r[t], i);
			return i = this.pushStack(n > 1 ? Z.unique(i) : i),
			i.selector = this.selector ? this.selector + " " + e: e,
			i
		},
		filter: function(e) {
			return this.pushStack(i(this, e || [], !1))
		},
		not: function(e) {
			return this.pushStack(i(this, e || [], !0))
		},
		is: function(e) {
			return !! i(this, "string" == typeof e && ot.test(e) ? Z(e) : e || [], !1).length
		}
	});
	var lt, ct = /^(?:\s*(<[\w\W]+>)[^>]*|#([\w-]*))$/,
	ut = Z.fn.init = function(e, t) {
		var n, i;
		if (!e) return this;
		if ("string" == typeof e) {
			if (n = "<" === e[0] && ">" === e[e.length - 1] && e.length >= 3 ? [null, e, null] : ct.exec(e), !n || !n[1] && t) return ! t || t.jquery ? (t || lt).find(e) : this.constructor(t).find(e);
			if (n[1]) {
				if (t = t instanceof Z ? t[0] : t, Z.merge(this, Z.parseHTML(n[1], t && t.nodeType ? t.ownerDocument || t: V, !0)), at.test(n[1]) && Z.isPlainObject(t)) for (n in t) Z.isFunction(this[n]) ? this[n](t[n]) : this.attr(n, t[n]);
				return this
			}
			return i = V.getElementById(n[2]),
			i && i.parentNode && (this.length = 1, this[0] = i),
			this.context = V,
			this.selector = e,
			this
		}
		return e.nodeType ? (this.context = this[0] = e, this.length = 1, this) : Z.isFunction(e) ? "undefined" != typeof lt.ready ? lt.ready(e) : e(Z) : (void 0 !== e.selector && (this.selector = e.selector, this.context = e.context), Z.makeArray(e, this))
	};
	ut.prototype = Z.fn,
	lt = Z(V);
	var dt = /^(?:parents|prev(?:Until|All))/,
	ft = {
		children: !0,
		contents: !0,
		next: !0,
		prev: !0
	};
	Z.extend({
		dir: function(e, t, n) {
			for (var i = [], r = void 0 !== n; (e = e[t]) && 9 !== e.nodeType;) if (1 === e.nodeType) {
				if (r && Z(e).is(n)) break;
				i.push(e)
			}
			return i
		},
		sibling: function(e, t) {
			for (var n = []; e; e = e.nextSibling) 1 === e.nodeType && e !== t && n.push(e);
			return n
		}
	}),
	Z.fn.extend({
		has: function(e) {
			var t = Z(e, this),
			n = t.length;
			return this.filter(function() {
				for (var e = 0; n > e; e++) if (Z.contains(this, t[e])) return ! 0
			})
		},
		closest: function(e, t) {
			for (var n, i = 0,
			r = this.length,
			o = [], a = ot.test(e) || "string" != typeof e ? Z(e, t || this.context) : 0; r > i; i++) for (n = this[i]; n && n !== t; n = n.parentNode) if (n.nodeType < 11 && (a ? a.index(n) > -1 : 1 === n.nodeType && Z.find.matchesSelector(n, e))) {
				o.push(n);
				break
			}
			return this.pushStack(o.length > 1 ? Z.unique(o) : o)
		},
		index: function(e) {
			return e ? "string" == typeof e ? z.call(Z(e), this[0]) : z.call(this, e.jquery ? e[0] : e) : this[0] && this[0].parentNode ? this.first().prevAll().length: -1
		},
		add: function(e, t) {
			return this.pushStack(Z.unique(Z.merge(this.get(), Z(e, t))))
		},
		addBack: function(e) {
			return this.add(null == e ? this.prevObject: this.prevObject.filter(e))
		}
	}),
	Z.each({
		parent: function(e) {
			var t = e.parentNode;
			return t && 11 !== t.nodeType ? t: null
		},
		parents: function(e) {
			return Z.dir(e, "parentNode")
		},
		parentsUntil: function(e, t, n) {
			return Z.dir(e, "parentNode", n)
		},
		next: function(e) {
			return r(e, "nextSibling")
		},
		prev: function(e) {
			return r(e, "previousSibling")
		},
		nextAll: function(e) {
			return Z.dir(e, "nextSibling")
		},
		prevAll: function(e) {
			return Z.dir(e, "previousSibling")
		},
		nextUntil: function(e, t, n) {
			return Z.dir(e, "nextSibling", n)
		},
		prevUntil: function(e, t, n) {
			return Z.dir(e, "previousSibling", n)
		},
		siblings: function(e) {
			return Z.sibling((e.parentNode || {}).firstChild, e)
		},
		children: function(e) {
			return Z.sibling(e.firstChild)
		},
		contents: function(e) {
			return e.contentDocument || Z.merge([], e.childNodes)
		}
	},
	function(e, t) {
		Z.fn[e] = function(n, i) {
			var r = Z.map(this, t, n);
			return "Until" !== e.slice( - 5) && (i = n),
			i && "string" == typeof i && (r = Z.filter(i, r)),
			this.length > 1 && (ft[e] || Z.unique(r), dt.test(e) && r.reverse()),
			this.pushStack(r)
		}
	});
	var pt = /\S+/g,
	ht = {};
	Z.Callbacks = function(e) {
		e = "string" == typeof e ? ht[e] || o(e) : Z.extend({},
		e);
		var t, n, i, r, a, s, l = [],
		c = !e.once && [],
		u = function(o) {
			for (t = e.memory && o, n = !0, s = r || 0, r = 0, a = l.length, i = !0; l && a > s; s++) if (l[s].apply(o[0], o[1]) === !1 && e.stopOnFalse) {
				t = !1;
				break
			}
			i = !1,
			l && (c ? c.length && u(c.shift()) : t ? l = [] : d.disable())
		},
		d = {
			add: function() {
				if (l) {
					var n = l.length; !
					function o(t) {
						Z.each(t,
						function(t, n) {
							var i = Z.type(n);
							"function" === i ? e.unique && d.has(n) || l.push(n) : n && n.length && "string" !== i && o(n)
						})
					} (arguments),
					i ? a = l.length: t && (r = n, u(t))
				}
				return this
			},
			remove: function() {
				return l && Z.each(arguments,
				function(e, t) {
					for (var n; (n = Z.inArray(t, l, n)) > -1;) l.splice(n, 1),
					i && (a >= n && a--, s >= n && s--)
				}),
				this
			},
			has: function(e) {
				return e ? Z.inArray(e, l) > -1 : !(!l || !l.length)
			},
			empty: function() {
				return l = [],
				a = 0,
				this
			},
			disable: function() {
				return l = c = t = void 0,
				this
			},
			disabled: function() {
				return ! l
			},
			lock: function() {
				return c = void 0,
				t || d.disable(),
				this
			},
			locked: function() {
				return ! c
			},
			fireWith: function(e, t) {
				return ! l || n && !c || (t = t || [], t = [e, t.slice ? t.slice() : t], i ? c.push(t) : u(t)),
				this
			},
			fire: function() {
				return d.fireWith(this, arguments),
				this
			},
			fired: function() {
				return !! n
			}
		};
		return d
	},
	Z.extend({
		Deferred: function(e) {
			var t = [["resolve", "done", Z.Callbacks("once memory"), "resolved"], ["reject", "fail", Z.Callbacks("once memory"), "rejected"], ["notify", "progress", Z.Callbacks("memory")]],
			n = "pending",
			i = {
				state: function() {
					return n
				},
				always: function() {
					return r.done(arguments).fail(arguments),
					this
				},
				then: function() {
					var e = arguments;
					return Z.Deferred(function(n) {
						Z.each(t,
						function(t, o) {
							var a = Z.isFunction(e[t]) && e[t];
							r[o[1]](function() {
								var e = a && a.apply(this, arguments);
								e && Z.isFunction(e.promise) ? e.promise().done(n.resolve).fail(n.reject).progress(n.notify) : n[o[0] + "With"](this === i ? n.promise() : this, a ? [e] : arguments)
							})
						}),
						e = null
					}).promise()
				},
				promise: function(e) {
					return null != e ? Z.extend(e, i) : i
				}
			},
			r = {};
			return i.pipe = i.then,
			Z.each(t,
			function(e, o) {
				var a = o[2],
				s = o[3];
				i[o[1]] = a.add,
				s && a.add(function() {
					n = s
				},
				t[1 ^ e][2].disable, t[2][2].lock),
				r[o[0]] = function() {
					return r[o[0] + "With"](this === r ? i: this, arguments),
					this
				},
				r[o[0] + "With"] = a.fireWith
			}),
			i.promise(r),
			e && e.call(r, r),
			r
		},
		when: function(e) {
			var t, n, i, r = 0,
			o = F.call(arguments),
			a = o.length,
			s = 1 !== a || e && Z.isFunction(e.promise) ? a: 0,
			l = 1 === s ? e: Z.Deferred(),
			c = function(e, n, i) {
				return function(r) {
					n[e] = this,
					i[e] = arguments.length > 1 ? F.call(arguments) : r,
					i === t ? l.notifyWith(n, i) : --s || l.resolveWith(n, i)
				}
			};
			if (a > 1) for (t = new Array(a), n = new Array(a), i = new Array(a); a > r; r++) o[r] && Z.isFunction(o[r].promise) ? o[r].promise().done(c(r, i, o)).fail(l.reject).progress(c(r, n, t)) : --s;
			return s || l.resolveWith(i, o),
			l.promise()
		}
	});
	var gt;
	Z.fn.ready = function(e) {
		return Z.ready.promise().done(e),
		this
	},
	Z.extend({
		isReady: !1,
		readyWait: 1,
		holdReady: function(e) {
			e ? Z.readyWait++:Z.ready(!0)
		},
		ready: function(e) { (e === !0 ? --Z.readyWait: Z.isReady) || (Z.isReady = !0, e !== !0 && --Z.readyWait > 0 || (gt.resolveWith(V, [Z]), Z.fn.triggerHandler && (Z(V).triggerHandler("ready"), Z(V).off("ready"))))
		}
	}),
	Z.ready.promise = function(t) {
		return gt || (gt = Z.Deferred(), "complete" === V.readyState ? setTimeout(Z.ready) : (V.addEventListener("DOMContentLoaded", a, !1), e.addEventListener("load", a, !1))),
		gt.promise(t)
	},
	Z.ready.promise();
	var mt = Z.access = function(e, t, n, i, r, o, a) {
		var s = 0,
		l = e.length,
		c = null == n;
		if ("object" === Z.type(n)) {
			r = !0;
			for (s in n) Z.access(e, t, s, n[s], !0, o, a)
		} else if (void 0 !== i && (r = !0, Z.isFunction(i) || (a = !0), c && (a ? (t.call(e, i), t = null) : (c = t, t = function(e, t, n) {
			return c.call(Z(e), n)
		})), t)) for (; l > s; s++) t(e[s], n, a ? i: i.call(e[s], s, t(e[s], n)));
		return r ? e: c ? t.call(e) : l ? t(e[0], n) : o
	};
	Z.acceptData = function(e) {
		return 1 === e.nodeType || 9 === e.nodeType || !+e.nodeType
	},
	s.uid = 1,
	s.accepts = Z.acceptData,
	s.prototype = {
		key: function(e) {
			if (!s.accepts(e)) return 0;
			var t = {},
			n = e[this.expando];
			if (!n) {
				n = s.uid++;
				try {
					t[this.expando] = {
						value: n
					},
					Object.defineProperties(e, t)
				} catch(i) {
					t[this.expando] = n,
					Z.extend(e, t)
				}
			}
			return this.cache[n] || (this.cache[n] = {}),
			n
		},
		set: function(e, t, n) {
			var i, r = this.key(e),
			o = this.cache[r];
			if ("string" == typeof t) o[t] = n;
			else if (Z.isEmptyObject(o)) Z.extend(this.cache[r], t);
			else for (i in t) o[i] = t[i];
			return o
		},
		get: function(e, t) {
			var n = this.cache[this.key(e)];
			return void 0 === t ? n: n[t]
		},
		access: function(e, t, n) {
			var i;
			return void 0 === t || t && "string" == typeof t && void 0 === n ? (i = this.get(e, t), void 0 !== i ? i: this.get(e, Z.camelCase(t))) : (this.set(e, t, n), void 0 !== n ? n: t)
		},
		remove: function(e, t) {
			var n, i, r, o = this.key(e),
			a = this.cache[o];
			if (void 0 === t) this.cache[o] = {};
			else {
				Z.isArray(t) ? i = t.concat(t.map(Z.camelCase)) : (r = Z.camelCase(t), t in a ? i = [t, r] : (i = r, i = i in a ? [i] : i.match(pt) || [])),
				n = i.length;
				for (; n--;) delete a[i[n]]
			}
		},
		hasData: function(e) {
			return ! Z.isEmptyObject(this.cache[e[this.expando]] || {})
		},
		discard: function(e) {
			e[this.expando] && delete this.cache[e[this.expando]]
		}
	};
	var vt = new s,
	yt = new s,
	bt = /^(?:\{[\w\W]*\}|\[[\w\W]*\])$/,
	wt = /([A-Z])/g;
	Z.extend({
		hasData: function(e) {
			return yt.hasData(e) || vt.hasData(e)
		},
		data: function(e, t, n) {
			return yt.access(e, t, n)
		},
		removeData: function(e, t) {
			yt.remove(e, t)
		},
		_data: function(e, t, n) {
			return vt.access(e, t, n)
		},
		_removeData: function(e, t) {
			vt.remove(e, t)
		}
	}),
	Z.fn.extend({
		data: function(e, t) {
			var n, i, r, o = this[0],
			a = o && o.attributes;
			if (void 0 === e) {
				if (this.length && (r = yt.get(o), 1 === o.nodeType && !vt.get(o, "hasDataAttrs"))) {
					for (n = a.length; n--;) a[n] && (i = a[n].name, 0 === i.indexOf("data-") && (i = Z.camelCase(i.slice(5)), l(o, i, r[i])));
					vt.set(o, "hasDataAttrs", !0)
				}
				return r
			}
			return "object" == typeof e ? this.each(function() {
				yt.set(this, e)
			}) : mt(this,
			function(t) {
				var n, i = Z.camelCase(e);
				if (o && void 0 === t) {
					if (n = yt.get(o, e), void 0 !== n) return n;
					if (n = yt.get(o, i), void 0 !== n) return n;
					if (n = l(o, i, void 0), void 0 !== n) return n
				} else this.each(function() {
					var n = yt.get(this, i);
					yt.set(this, i, t),
					-1 !== e.indexOf("-") && void 0 !== n && yt.set(this, e, t)
				})
			},
			null, t, arguments.length > 1, null, !0)
		},
		removeData: function(e) {
			return this.each(function() {
				yt.remove(this, e)
			})
		}
	}),
	Z.extend({
		queue: function(e, t, n) {
			var i;
			return e ? (t = (t || "fx") + "queue", i = vt.get(e, t), n && (!i || Z.isArray(n) ? i = vt.access(e, t, Z.makeArray(n)) : i.push(n)), i || []) : void 0
		},
		dequeue: function(e, t) {
			t = t || "fx";
			var n = Z.queue(e, t),
			i = n.length,
			r = n.shift(),
			o = Z._queueHooks(e, t),
			a = function() {
				Z.dequeue(e, t)
			};
			"inprogress" === r && (r = n.shift(), i--),
			r && ("fx" === t && n.unshift("inprogress"), delete o.stop, r.call(e, a, o)),
			!i && o && o.empty.fire()
		},
		_queueHooks: function(e, t) {
			var n = t + "queueHooks";
			return vt.get(e, n) || vt.access(e, n, {
				empty: Z.Callbacks("once memory").add(function() {
					vt.remove(e, [t + "queue", n])
				})
			})
		}
	}),
	Z.fn.extend({
		queue: function(e, t) {
			var n = 2;
			return "string" != typeof e && (t = e, e = "fx", n--),
			arguments.length < n ? Z.queue(this[0], e) : void 0 === t ? this: this.each(function() {
				var n = Z.queue(this, e, t);
				Z._queueHooks(this, e),
				"fx" === e && "inprogress" !== n[0] && Z.dequeue(this, e)
			})
		},
		dequeue: function(e) {
			return this.each(function() {
				Z.dequeue(this, e)
			})
		},
		clearQueue: function(e) {
			return this.queue(e || "fx", [])
		},
		promise: function(e, t) {
			var n, i = 1,
			r = Z.Deferred(),
			o = this,
			a = this.length,
			s = function() {--i || r.resolveWith(o, [o])
			};
			for ("string" != typeof e && (t = e, e = void 0), e = e || "fx"; a--;) n = vt.get(o[a], e + "queueHooks"),
			n && n.empty && (i++, n.empty.add(s));
			return s(),
			r.promise(t)
		}
	});
	var xt = /[+-]?(?:\d*\.|)\d+(?:[eE][+-]?\d+|)/.source,
	kt = ["Top", "Right", "Bottom", "Left"],
	Ct = function(e, t) {
		return e = t || e,
		"none" === Z.css(e, "display") || !Z.contains(e.ownerDocument, e)
	},
	Tt = /^(?:checkbox|radio)$/i; !
	function() {
		var e = V.createDocumentFragment(),
		t = e.appendChild(V.createElement("div")),
		n = V.createElement("input");
		n.setAttribute("type", "radio"),
		n.setAttribute("checked", "checked"),
		n.setAttribute("name", "t"),
		t.appendChild(n),
		Q.checkClone = t.cloneNode(!0).cloneNode(!0).lastChild.checked,
		t.innerHTML = "<textarea>x</textarea>",
		Q.noCloneChecked = !!t.cloneNode(!0).lastChild.defaultValue
	} ();
	var St = "undefined";
	Q.focusinBubbles = "onfocusin" in e;
	var _t = /^key/,
	$t = /^(?:mouse|pointer|contextmenu)|click/,
	Lt = /^(?:focusinfocus|focusoutblur)$/,
	Et = /^([^.]*)(?:\.(.+)|)$/;
	Z.event = {
		global: {},
		add: function(e, t, n, i, r) {
			var o, a, s, l, c, u, d, f, p, h, g, m = vt.get(e);
			if (m) for (n.handler && (o = n, n = o.handler, r = o.selector), n.guid || (n.guid = Z.guid++), (l = m.events) || (l = m.events = {}), (a = m.handle) || (a = m.handle = function(t) {
				return typeof Z !== St && Z.event.triggered !== t.type ? Z.event.dispatch.apply(e, arguments) : void 0
			}), t = (t || "").match(pt) || [""], c = t.length; c--;) s = Et.exec(t[c]) || [],
			p = g = s[1],
			h = (s[2] || "").split(".").sort(),
			p && (d = Z.event.special[p] || {},
			p = (r ? d.delegateType: d.bindType) || p, d = Z.event.special[p] || {},
			u = Z.extend({
				type: p,
				origType: g,
				data: i,
				handler: n,
				guid: n.guid,
				selector: r,
				needsContext: r && Z.expr.match.needsContext.test(r),
				namespace: h.join(".")
			},
			o), (f = l[p]) || (f = l[p] = [], f.delegateCount = 0, d.setup && d.setup.call(e, i, h, a) !== !1 || e.addEventListener && e.addEventListener(p, a, !1)), d.add && (d.add.call(e, u), u.handler.guid || (u.handler.guid = n.guid)), r ? f.splice(f.delegateCount++, 0, u) : f.push(u), Z.event.global[p] = !0)
		},
		remove: function(e, t, n, i, r) {
			var o, a, s, l, c, u, d, f, p, h, g, m = vt.hasData(e) && vt.get(e);
			if (m && (l = m.events)) {
				for (t = (t || "").match(pt) || [""], c = t.length; c--;) if (s = Et.exec(t[c]) || [], p = g = s[1], h = (s[2] || "").split(".").sort(), p) {
					for (d = Z.event.special[p] || {},
					p = (i ? d.delegateType: d.bindType) || p, f = l[p] || [], s = s[2] && new RegExp("(^|\\.)" + h.join("\\.(?:.*\\.|)") + "(\\.|$)"), a = o = f.length; o--;) u = f[o],
					!r && g !== u.origType || n && n.guid !== u.guid || s && !s.test(u.namespace) || i && i !== u.selector && ("**" !== i || !u.selector) || (f.splice(o, 1), u.selector && f.delegateCount--, d.remove && d.remove.call(e, u));
					a && !f.length && (d.teardown && d.teardown.call(e, h, m.handle) !== !1 || Z.removeEvent(e, p, m.handle), delete l[p])
				} else for (p in l) Z.event.remove(e, p + t[c], n, i, !0);
				Z.isEmptyObject(l) && (delete m.handle, vt.remove(e, "events"))
			}
		},
		trigger: function(t, n, i, r) {
			var o, a, s, l, c, u, d, f = [i || V],
			p = K.call(t, "type") ? t.type: t,
			h = K.call(t, "namespace") ? t.namespace.split(".") : [];
			if (a = s = i = i || V, 3 !== i.nodeType && 8 !== i.nodeType && !Lt.test(p + Z.event.triggered) && (p.indexOf(".") >= 0 && (h = p.split("."), p = h.shift(), h.sort()), c = p.indexOf(":") < 0 && "on" + p, t = t[Z.expando] ? t: new Z.Event(p, "object" == typeof t && t), t.isTrigger = r ? 2 : 3, t.namespace = h.join("."), t.namespace_re = t.namespace ? new RegExp("(^|\\.)" + h.join("\\.(?:.*\\.|)") + "(\\.|$)") : null, t.result = void 0, t.target || (t.target = i), n = null == n ? [t] : Z.makeArray(n, [t]), d = Z.event.special[p] || {},
			r || !d.trigger || d.trigger.apply(i, n) !== !1)) {
				if (!r && !d.noBubble && !Z.isWindow(i)) {
					for (l = d.delegateType || p, Lt.test(l + p) || (a = a.parentNode); a; a = a.parentNode) f.push(a),
					s = a;
					s === (i.ownerDocument || V) && f.push(s.defaultView || s.parentWindow || e)
				}
				for (o = 0; (a = f[o++]) && !t.isPropagationStopped();) t.type = o > 1 ? l: d.bindType || p,
				u = (vt.get(a, "events") || {})[t.type] && vt.get(a, "handle"),
				u && u.apply(a, n),
				u = c && a[c],
				u && u.apply && Z.acceptData(a) && (t.result = u.apply(a, n), t.result === !1 && t.preventDefault());
				return t.type = p,
				r || t.isDefaultPrevented() || d._default && d._default.apply(f.pop(), n) !== !1 || !Z.acceptData(i) || c && Z.isFunction(i[p]) && !Z.isWindow(i) && (s = i[c], s && (i[c] = null), Z.event.triggered = p, i[p](), Z.event.triggered = void 0, s && (i[c] = s)),
				t.result
			}
		},
		dispatch: function(e) {
			e = Z.event.fix(e);
			var t, n, i, r, o, a = [],
			s = F.call(arguments),
			l = (vt.get(this, "events") || {})[e.type] || [],
			c = Z.event.special[e.type] || {};
			if (s[0] = e, e.delegateTarget = this, !c.preDispatch || c.preDispatch.call(this, e) !== !1) {
				for (a = Z.event.handlers.call(this, e, l), t = 0; (r = a[t++]) && !e.isPropagationStopped();) for (e.currentTarget = r.elem, n = 0; (o = r.handlers[n++]) && !e.isImmediatePropagationStopped();)(!e.namespace_re || e.namespace_re.test(o.namespace)) && (e.handleObj = o, e.data = o.data, i = ((Z.event.special[o.origType] || {}).handle || o.handler).apply(r.elem, s), void 0 !== i && (e.result = i) === !1 && (e.preventDefault(), e.stopPropagation()));
				return c.postDispatch && c.postDispatch.call(this, e),
				e.result
			}
		},
		handlers: function(e, t) {
			var n, i, r, o, a = [],
			s = t.delegateCount,
			l = e.target;
			if (s && l.nodeType && (!e.button || "click" !== e.type)) for (; l !== this; l = l.parentNode || this) if (l.disabled !== !0 || "click" !== e.type) {
				for (i = [], n = 0; s > n; n++) o = t[n],
				r = o.selector + " ",
				void 0 === i[r] && (i[r] = o.needsContext ? Z(r, this).index(l) >= 0 : Z.find(r, this, null, [l]).length),
				i[r] && i.push(o);
				i.length && a.push({
					elem: l,
					handlers: i
				})
			}
			return s < t.length && a.push({
				elem: this,
				handlers: t.slice(s)
			}),
			a
		},
		props: "altKey bubbles cancelable ctrlKey currentTarget eventPhase metaKey relatedTarget shiftKey target timeStamp view which".split(" "),
		fixHooks: {},
		keyHooks: {
			props: "char charCode key keyCode".split(" "),
			filter: function(e, t) {
				return null == e.which && (e.which = null != t.charCode ? t.charCode: t.keyCode),
				e
			}
		},
		mouseHooks: {
			props: "button buttons clientX clientY offsetX offsetY pageX pageY screenX screenY toElement".split(" "),
			filter: function(e, t) {
				var n, i, r, o = t.button;
				return null == e.pageX && null != t.clientX && (n = e.target.ownerDocument || V, i = n.documentElement, r = n.body, e.pageX = t.clientX + (i && i.scrollLeft || r && r.scrollLeft || 0) - (i && i.clientLeft || r && r.clientLeft || 0), e.pageY = t.clientY + (i && i.scrollTop || r && r.scrollTop || 0) - (i && i.clientTop || r && r.clientTop || 0)),
				e.which || void 0 === o || (e.which = 1 & o ? 1 : 2 & o ? 3 : 4 & o ? 2 : 0),
				e
			}
		},
		fix: function(e) {
			if (e[Z.expando]) return e;
			var t, n, i, r = e.type,
			o = e,
			a = this.fixHooks[r];
			for (a || (this.fixHooks[r] = a = $t.test(r) ? this.mouseHooks: _t.test(r) ? this.keyHooks: {}), i = a.props ? this.props.concat(a.props) : this.props, e = new Z.Event(o), t = i.length; t--;) n = i[t],
			e[n] = o[n];
			return e.target || (e.target = V),
			3 === e.target.nodeType && (e.target = e.target.parentNode),
			a.filter ? a.filter(e, o) : e
		},
		special: {
			load: {
				noBubble: !0
			},
			focus: {
				trigger: function() {
					return this !== d() && this.focus ? (this.focus(), !1) : void 0
				},
				delegateType: "focusin"
			},
			blur: {
				trigger: function() {
					return this === d() && this.blur ? (this.blur(), !1) : void 0
				},
				delegateType: "focusout"
			},
			click: {
				trigger: function() {
					return "checkbox" === this.type && this.click && Z.nodeName(this, "input") ? (this.click(), !1) : void 0
				},
				_default: function(e) {
					return Z.nodeName(e.target, "a")
				}
			},
			beforeunload: {
				postDispatch: function(e) {
					void 0 !== e.result && e.originalEvent && (e.originalEvent.returnValue = e.result)
				}
			}
		},
		simulate: function(e, t, n, i) {
			var r = Z.extend(new Z.Event, n, {
				type: e,
				isSimulated: !0,
				originalEvent: {}
			});
			i ? Z.event.trigger(r, null, t) : Z.event.dispatch.call(t, r),
			r.isDefaultPrevented() && n.preventDefault()
		}
	},
	Z.removeEvent = function(e, t, n) {
		e.removeEventListener && e.removeEventListener(t, n, !1)
	},
	Z.Event = function(e, t) {
		return this instanceof Z.Event ? (e && e.type ? (this.originalEvent = e, this.type = e.type, this.isDefaultPrevented = e.defaultPrevented || void 0 === e.defaultPrevented && e.returnValue === !1 ? c: u) : this.type = e, t && Z.extend(this, t), this.timeStamp = e && e.timeStamp || Z.now(), void(this[Z.expando] = !0)) : new Z.Event(e, t)
	},
	Z.Event.prototype = {
		isDefaultPrevented: u,
		isPropagationStopped: u,
		isImmediatePropagationStopped: u,
		preventDefault: function() {
			var e = this.originalEvent;
			this.isDefaultPrevented = c,
			e && e.preventDefault && e.preventDefault()
		},
		stopPropagation: function() {
			var e = this.originalEvent;
			this.isPropagationStopped = c,
			e && e.stopPropagation && e.stopPropagation()
		},
		stopImmediatePropagation: function() {
			var e = this.originalEvent;
			this.isImmediatePropagationStopped = c,
			e && e.stopImmediatePropagation && e.stopImmediatePropagation(),
			this.stopPropagation()
		}
	},
	Z.each({
		mouseenter: "mouseover",
		mouseleave: "mouseout",
		pointerenter: "pointerover",
		pointerleave: "pointerout"
	},
	function(e, t) {
		Z.event.special[e] = {
			delegateType: t,
			bindType: t,
			handle: function(e) {
				var n, i = this,
				r = e.relatedTarget,
				o = e.handleObj;
				return (!r || r !== i && !Z.contains(i, r)) && (e.type = o.origType, n = o.handler.apply(this, arguments), e.type = t),
				n
			}
		}
	}),
	Q.focusinBubbles || Z.each({
		focus: "focusin",
		blur: "focusout"
	},
	function(e, t) {
		var n = function(e) {
			Z.event.simulate(t, e.target, Z.event.fix(e), !0)
		};
		Z.event.special[t] = {
			setup: function() {
				var i = this.ownerDocument || this,
				r = vt.access(i, t);
				r || i.addEventListener(e, n, !0),
				vt.access(i, t, (r || 0) + 1)
			},
			teardown: function() {
				var i = this.ownerDocument || this,
				r = vt.access(i, t) - 1;
				r ? vt.access(i, t, r) : (i.removeEventListener(e, n, !0), vt.remove(i, t))
			}
		}
	}),
	Z.fn.extend({
		on: function(e, t, n, i, r) {
			var o, a;
			if ("object" == typeof e) {
				"string" != typeof t && (n = n || t, t = void 0);
				for (a in e) this.on(a, t, n, e[a], r);
				return this
			}
			if (null == n && null == i ? (i = t, n = t = void 0) : null == i && ("string" == typeof t ? (i = n, n = void 0) : (i = n, n = t, t = void 0)), i === !1) i = u;
			else if (!i) return this;
			return 1 === r && (o = i, i = function(e) {
				return Z().off(e),
				o.apply(this, arguments)
			},
			i.guid = o.guid || (o.guid = Z.guid++)),
			this.each(function() {
				Z.event.add(this, e, i, n, t)
			})
		},
		one: function(e, t, n, i) {
			return this.on(e, t, n, i, 1)
		},
		off: function(e, t, n) {
			var i, r;
			if (e && e.preventDefault && e.handleObj) return i = e.handleObj,
			Z(e.delegateTarget).off(i.namespace ? i.origType + "." + i.namespace: i.origType, i.selector, i.handler),
			this;
			if ("object" == typeof e) {
				for (r in e) this.off(r, t, e[r]);
				return this
			}
			return (t === !1 || "function" == typeof t) && (n = t, t = void 0),
			n === !1 && (n = u),
			this.each(function() {
				Z.event.remove(this, e, n, t)
			})
		},
		trigger: function(e, t) {
			return this.each(function() {
				Z.event.trigger(e, t, this)
			})
		},
		triggerHandler: function(e, t) {
			var n = this[0];
			return n ? Z.event.trigger(e, t, n, !0) : void 0
		}
	});
	var It = /<(?!area|br|col|embed|hr|img|input|link|meta|param)(([\w:]+)[^>]*)\/>/gi,
	Nt = /<([\w:]+)/,
	At = /<|&#?\w+;/,
	Dt = /<(?:script|style|link)/i,
	jt = /checked\s*(?:[^=]|=\s*.checked.)/i,
	Ht = /^$|\/(?:java|ecma)script/i,
	Mt = /^true\/(.*)/,
	Pt = /^\s*<!(?:\[CDATA\[|--)|(?:\]\]|--)>\s*$/g,
	Wt = {
		option: [1, "<select multiple='multiple'>", "</select>"],
		thead: [1, "<table>", "</table>"],
		col: [2, "<table><colgroup>", "</colgroup></table>"],
		tr: [2, "<table><tbody>", "</tbody></table>"],
		td: [3, "<table><tbody><tr>", "</tr></tbody></table>"],
		_default: [0, "", ""]
	};
	Wt.optgroup = Wt.option,
	Wt.tbody = Wt.tfoot = Wt.colgroup = Wt.caption = Wt.thead,
	Wt.th = Wt.td,
	Z.extend({
		clone: function(e, t, n) {
			var i, r, o, a, s = e.cloneNode(!0),
			l = Z.contains(e.ownerDocument, e);
			if (! (Q.noCloneChecked || 1 !== e.nodeType && 11 !== e.nodeType || Z.isXMLDoc(e))) for (a = v(s), o = v(e), i = 0, r = o.length; r > i; i++) y(o[i], a[i]);
			if (t) if (n) for (o = o || v(e), a = a || v(s), i = 0, r = o.length; r > i; i++) m(o[i], a[i]);
			else m(e, s);
			return a = v(s, "script"),
			a.length > 0 && g(a, !l && v(e, "script")),
			s
		},
		buildFragment: function(e, t, n, i) {
			for (var r, o, a, s, l, c, u = t.createDocumentFragment(), d = [], f = 0, p = e.length; p > f; f++) if (r = e[f], r || 0 === r) if ("object" === Z.type(r)) Z.merge(d, r.nodeType ? [r] : r);
			else if (At.test(r)) {
				for (o = o || u.appendChild(t.createElement("div")), a = (Nt.exec(r) || ["", ""])[1].toLowerCase(), s = Wt[a] || Wt._default, o.innerHTML = s[1] + r.replace(It, "<$1></$2>") + s[2], c = s[0]; c--;) o = o.lastChild;
				Z.merge(d, o.childNodes),
				o = u.firstChild,
				o.textContent = ""
			} else d.push(t.createTextNode(r));
			for (u.textContent = "", f = 0; r = d[f++];) if ((!i || -1 === Z.inArray(r, i)) && (l = Z.contains(r.ownerDocument, r), o = v(u.appendChild(r), "script"), l && g(o), n)) for (c = 0; r = o[c++];) Ht.test(r.type || "") && n.push(r);
			return u
		},
		cleanData: function(e) {
			for (var t, n, i, r, o = Z.event.special,
			a = 0; void 0 !== (n = e[a]); a++) {
				if (Z.acceptData(n) && (r = n[vt.expando], r && (t = vt.cache[r]))) {
					if (t.events) for (i in t.events) o[i] ? Z.event.remove(n, i) : Z.removeEvent(n, i, t.handle);
					vt.cache[r] && delete vt.cache[r]
				}
				delete yt.cache[n[yt.expando]]
			}
		}
	}),
	Z.fn.extend({
		text: function(e) {
			return mt(this,
			function(e) {
				return void 0 === e ? Z.text(this) : this.empty().each(function() { (1 === this.nodeType || 11 === this.nodeType || 9 === this.nodeType) && (this.textContent = e)
				})
			},
			null, e, arguments.length)
		},
		append: function() {
			return this.domManip(arguments,
			function(e) {
				if (1 === this.nodeType || 11 === this.nodeType || 9 === this.nodeType) {
					var t = f(this, e);
					t.appendChild(e)
				}
			})
		},
		prepend: function() {
			return this.domManip(arguments,
			function(e) {
				if (1 === this.nodeType || 11 === this.nodeType || 9 === this.nodeType) {
					var t = f(this, e);
					t.insertBefore(e, t.firstChild)
				}
			})
		},
		before: function() {
			return this.domManip(arguments,
			function(e) {
				this.parentNode && this.parentNode.insertBefore(e, this)
			})
		},
		after: function() {
			return this.domManip(arguments,
			function(e) {
				this.parentNode && this.parentNode.insertBefore(e, this.nextSibling)
			})
		},
		remove: function(e, t) {
			for (var n, i = e ? Z.filter(e, this) : this, r = 0; null != (n = i[r]); r++) t || 1 !== n.nodeType || Z.cleanData(v(n)),
			n.parentNode && (t && Z.contains(n.ownerDocument, n) && g(v(n, "script")), n.parentNode.removeChild(n));
			return this
		},
		empty: function() {
			for (var e, t = 0; null != (e = this[t]); t++) 1 === e.nodeType && (Z.cleanData(v(e, !1)), e.textContent = "");
			return this
		},
		clone: function(e, t) {
			return e = null == e ? !1 : e,
			t = null == t ? e: t,
			this.map(function() {
				return Z.clone(this, e, t)
			})
		},
		html: function(e) {
			return mt(this,
			function(e) {
				var t = this[0] || {},
				n = 0,
				i = this.length;
				if (void 0 === e && 1 === t.nodeType) return t.innerHTML;
				if ("string" == typeof e && !Dt.test(e) && !Wt[(Nt.exec(e) || ["", ""])[1].toLowerCase()]) {
					e = e.replace(It, "<$1></$2>");
					try {
						for (; i > n; n++) t = this[n] || {},
						1 === t.nodeType && (Z.cleanData(v(t, !1)), t.innerHTML = e);
						t = 0
					} catch(r) {}
				}
				t && this.empty().append(e)
			},
			null, e, arguments.length)
		},
		replaceWith: function() {
			var e = arguments[0];
			return this.domManip(arguments,
			function(t) {
				e = this.parentNode,
				Z.cleanData(v(this)),
				e && e.replaceChild(t, this)
			}),
			e && (e.length || e.nodeType) ? this: this.remove()
		},
		detach: function(e) {
			return this.remove(e, !0)
		},
		domManip: function(e, t) {
			e = B.apply([], e);
			var n, i, r, o, a, s, l = 0,
			c = this.length,
			u = this,
			d = c - 1,
			f = e[0],
			g = Z.isFunction(f);
			if (g || c > 1 && "string" == typeof f && !Q.checkClone && jt.test(f)) return this.each(function(n) {
				var i = u.eq(n);
				g && (e[0] = f.call(this, n, i.html())),
				i.domManip(e, t)
			});
			if (c && (n = Z.buildFragment(e, this[0].ownerDocument, !1, this), i = n.firstChild, 1 === n.childNodes.length && (n = i), i)) {
				for (r = Z.map(v(n, "script"), p), o = r.length; c > l; l++) a = n,
				l !== d && (a = Z.clone(a, !0, !0), o && Z.merge(r, v(a, "script"))),
				t.call(this[l], a, l);
				if (o) for (s = r[r.length - 1].ownerDocument, Z.map(r, h), l = 0; o > l; l++) a = r[l],
				Ht.test(a.type || "") && !vt.access(a, "globalEval") && Z.contains(s, a) && (a.src ? Z._evalUrl && Z._evalUrl(a.src) : Z.globalEval(a.textContent.replace(Pt, "")))
			}
			return this
		}
	}),
	Z.each({
		appendTo: "append",
		prependTo: "prepend",
		insertBefore: "before",
		insertAfter: "after",
		replaceAll: "replaceWith"
	},
	function(e, t) {
		Z.fn[e] = function(e) {
			for (var n, i = [], r = Z(e), o = r.length - 1, a = 0; o >= a; a++) n = a === o ? this: this.clone(!0),
			Z(r[a])[t](n),
			Y.apply(i, n.get());
			return this.pushStack(i)
		}
	});
	var Ot, Rt = {},
	qt = /^margin/,
	Xt = new RegExp("^(" + xt + ")(?!px)[a-z%]+$", "i"),
	Ft = function(t) {
		return t.ownerDocument.defaultView.opener ? t.ownerDocument.defaultView.getComputedStyle(t, null) : e.getComputedStyle(t, null)
	}; !
	function() {
		function t() {
			a.style.cssText = "-webkit-box-sizing:border-box;-moz-box-sizing:border-box;box-sizing:border-box;display:block;margin-top:1%;top:1%;border:1px;padding:1px;width:4px;position:absolute",
			a.innerHTML = "",
			r.appendChild(o);
			var t = e.getComputedStyle(a, null);
			n = "1%" !== t.top,
			i = "4px" === t.width,
			r.removeChild(o)
		}
		var n, i, r = V.documentElement,
		o = V.createElement("div"),
		a = V.createElement("div");
		a.style && (a.style.backgroundClip = "content-box", a.cloneNode(!0).style.backgroundClip = "", Q.clearCloneStyle = "content-box" === a.style.backgroundClip, o.style.cssText = "border:0;width:0;height:0;top:0;left:-9999px;margin-top:1px;position:absolute", o.appendChild(a), e.getComputedStyle && Z.extend(Q, {
			pixelPosition: function() {
				return t(),
				n
			},
			boxSizingReliable: function() {
				return null == i && t(),
				i
			},
			reliableMarginRight: function() {
				var t, n = a.appendChild(V.createElement("div"));
				return n.style.cssText = a.style.cssText = "-webkit-box-sizing:content-box;-moz-box-sizing:content-box;box-sizing:content-box;display:block;margin:0;border:0;padding:0",
				n.style.marginRight = n.style.width = "0",
				a.style.width = "1px",
				r.appendChild(o),
				t = !parseFloat(e.getComputedStyle(n, null).marginRight),
				r.removeChild(o),
				a.removeChild(n),
				t
			}
		}))
	} (),
	Z.swap = function(e, t, n, i) {
		var r, o, a = {};
		for (o in t) a[o] = e.style[o],
		e.style[o] = t[o];
		r = n.apply(e, i || []);
		for (o in t) e.style[o] = a[o];
		return r
	};
	var Bt = /^(none|table(?!-c[ea]).+)/,
	Yt = new RegExp("^(" + xt + ")(.*)$", "i"),
	zt = new RegExp("^([+-])=(" + xt + ")", "i"),
	Jt = {
		position: "absolute",
		visibility: "hidden",
		display: "block"
	},
	Ut = {
		letterSpacing: "0",
		fontWeight: "400"
	},
	Kt = ["Webkit", "O", "Moz", "ms"];
	Z.extend({
		cssHooks: {
			opacity: {
				get: function(e, t) {
					if (t) {
						var n = x(e, "opacity");
						return "" === n ? "1": n
					}
				}
			}
		},
		cssNumber: {
			columnCount: !0,
			fillOpacity: !0,
			flexGrow: !0,
			flexShrink: !0,
			fontWeight: !0,
			lineHeight: !0,
			opacity: !0,
			order: !0,
			orphans: !0,
			widows: !0,
			zIndex: !0,
			zoom: !0
		},
		cssProps: {
			"float": "cssFloat"
		},
		style: function(e, t, n, i) {
			if (e && 3 !== e.nodeType && 8 !== e.nodeType && e.style) {
				var r, o, a, s = Z.camelCase(t),
				l = e.style;
				return t = Z.cssProps[s] || (Z.cssProps[s] = C(l, s)),
				a = Z.cssHooks[t] || Z.cssHooks[s],
				void 0 === n ? a && "get" in a && void 0 !== (r = a.get(e, !1, i)) ? r: l[t] : (o = typeof n, "string" === o && (r = zt.exec(n)) && (n = (r[1] + 1) * r[2] + parseFloat(Z.css(e, t)), o = "number"), void(null != n && n === n && ("number" !== o || Z.cssNumber[s] || (n += "px"), Q.clearCloneStyle || "" !== n || 0 !== t.indexOf("background") || (l[t] = "inherit"), a && "set" in a && void 0 === (n = a.set(e, n, i)) || (l[t] = n))))
			}
		},
		css: function(e, t, n, i) {
			var r, o, a, s = Z.camelCase(t);
			return t = Z.cssProps[s] || (Z.cssProps[s] = C(e.style, s)),
			a = Z.cssHooks[t] || Z.cssHooks[s],
			a && "get" in a && (r = a.get(e, !0, n)),
			void 0 === r && (r = x(e, t, i)),
			"normal" === r && t in Ut && (r = Ut[t]),
			"" === n || n ? (o = parseFloat(r), n === !0 || Z.isNumeric(o) ? o || 0 : r) : r
		}
	}),
	Z.each(["height", "width"],
	function(e, t) {
		Z.cssHooks[t] = {
			get: function(e, n, i) {
				return n ? Bt.test(Z.css(e, "display")) && 0 === e.offsetWidth ? Z.swap(e, Jt,
				function() {
					return _(e, t, i)
				}) : _(e, t, i) : void 0
			},
			set: function(e, n, i) {
				var r = i && Ft(e);
				return T(e, n, i ? S(e, t, i, "border-box" === Z.css(e, "boxSizing", !1, r), r) : 0)
			}
		}
	}),
	Z.cssHooks.marginRight = k(Q.reliableMarginRight,
	function(e, t) {
		return t ? Z.swap(e, {
			display: "inline-block"
		},
		x, [e, "marginRight"]) : void 0
	}),
	Z.each({
		margin: "",
		padding: "",
		border: "Width"
	},
	function(e, t) {
		Z.cssHooks[e + t] = {
			expand: function(n) {
				for (var i = 0,
				r = {},
				o = "string" == typeof n ? n.split(" ") : [n]; 4 > i; i++) r[e + kt[i] + t] = o[i] || o[i - 2] || o[0];
				return r
			}
		},
		qt.test(e) || (Z.cssHooks[e + t].set = T)
	}),
	Z.fn.extend({
		css: function(e, t) {
			return mt(this,
			function(e, t, n) {
				var i, r, o = {},
				a = 0;
				if (Z.isArray(t)) {
					for (i = Ft(e), r = t.length; r > a; a++) o[t[a]] = Z.css(e, t[a], !1, i);
					return o
				}
				return void 0 !== n ? Z.style(e, t, n) : Z.css(e, t)
			},
			e, t, arguments.length > 1)
		},
		show: function() {
			return $(this, !0)
		},
		hide: function() {
			return $(this)
		},
		toggle: function(e) {
			return "boolean" == typeof e ? e ? this.show() : this.hide() : this.each(function() {
				Ct(this) ? Z(this).show() : Z(this).hide()
			})
		}
	}),
	Z.Tween = L,
	L.prototype = {
		constructor: L,
		init: function(e, t, n, i, r, o) {
			this.elem = e,
			this.prop = n,
			this.easing = r || "swing",
			this.options = t,
			this.start = this.now = this.cur(),
			this.end = i,
			this.unit = o || (Z.cssNumber[n] ? "": "px")
		},
		cur: function() {
			var e = L.propHooks[this.prop];
			return e && e.get ? e.get(this) : L.propHooks._default.get(this)
		},
		run: function(e) {
			var t, n = L.propHooks[this.prop];
			return this.pos = t = this.options.duration ? Z.easing[this.easing](e, this.options.duration * e, 0, 1, this.options.duration) : e,
			this.now = (this.end - this.start) * t + this.start,
			this.options.step && this.options.step.call(this.elem, this.now, this),
			n && n.set ? n.set(this) : L.propHooks._default.set(this),
			this
		}
	},
	L.prototype.init.prototype = L.prototype,
	L.propHooks = {
		_default: {
			get: function(e) {
				var t;
				return null == e.elem[e.prop] || e.elem.style && null != e.elem.style[e.prop] ? (t = Z.css(e.elem, e.prop, ""), t && "auto" !== t ? t: 0) : e.elem[e.prop]
			},
			set: function(e) {
				Z.fx.step[e.prop] ? Z.fx.step[e.prop](e) : e.elem.style && (null != e.elem.style[Z.cssProps[e.prop]] || Z.cssHooks[e.prop]) ? Z.style(e.elem, e.prop, e.now + e.unit) : e.elem[e.prop] = e.now
			}
		}
	},
	L.propHooks.scrollTop = L.propHooks.scrollLeft = {
		set: function(e) {
			e.elem.nodeType && e.elem.parentNode && (e.elem[e.prop] = e.now)
		}
	},
	Z.easing = {
		linear: function(e) {
			return e
		},
		swing: function(e) {
			return.5 - Math.cos(e * Math.PI) / 2
		}
	},
	Z.fx = L.prototype.init,
	Z.fx.step = {};
	var Qt, Vt, Gt = /^(?:toggle|show|hide)$/,
	Zt = new RegExp("^(?:([+-])=|)(" + xt + ")([a-z%]*)$", "i"),
	en = /queueHooks$/,
	tn = [A],
	nn = {
		"*": [function(e, t) {
			var n = this.createTween(e, t),
			i = n.cur(),
			r = Zt.exec(t),
			o = r && r[3] || (Z.cssNumber[e] ? "": "px"),
			a = (Z.cssNumber[e] || "px" !== o && +i) && Zt.exec(Z.css(n.elem, e)),
			s = 1,
			l = 20;
			if (a && a[3] !== o) {
				o = o || a[3],
				r = r || [],
				a = +i || 1;
				do s = s || ".5",
				a /= s,
				Z.style(n.elem, e, a + o);
				while (s !== (s = n.cur() / i) && 1 !== s && --l)
			}
			return r && (a = n.start = +a || +i || 0, n.unit = o, n.end = r[1] ? a + (r[1] + 1) * r[2] : +r[2]),
			n
		}]
	};
	Z.Animation = Z.extend(j, {
		tweener: function(e, t) {
			Z.isFunction(e) ? (t = e, e = ["*"]) : e = e.split(" ");
			for (var n, i = 0,
			r = e.length; r > i; i++) n = e[i],
			nn[n] = nn[n] || [],
			nn[n].unshift(t)
		},
		prefilter: function(e, t) {
			t ? tn.unshift(e) : tn.push(e)
		}
	}),
	Z.speed = function(e, t, n) {
		var i = e && "object" == typeof e ? Z.extend({},
		e) : {
			complete: n || !n && t || Z.isFunction(e) && e,
			duration: e,
			easing: n && t || t && !Z.isFunction(t) && t
		};
		return i.duration = Z.fx.off ? 0 : "number" == typeof i.duration ? i.duration: i.duration in Z.fx.speeds ? Z.fx.speeds[i.duration] : Z.fx.speeds._default,
		(null == i.queue || i.queue === !0) && (i.queue = "fx"),
		i.old = i.complete,
		i.complete = function() {
			Z.isFunction(i.old) && i.old.call(this),
			i.queue && Z.dequeue(this, i.queue)
		},
		i
	},
	Z.fn.extend({
		fadeTo: function(e, t, n, i) {
			return this.filter(Ct).css("opacity", 0).show().end().animate({
				opacity: t
			},
			e, n, i)
		},
		animate: function(e, t, n, i) {
			var r = Z.isEmptyObject(e),
			o = Z.speed(t, n, i),
			a = function() {
				var t = j(this, Z.extend({},
				e), o); (r || vt.get(this, "finish")) && t.stop(!0)
			};
			return a.finish = a,
			r || o.queue === !1 ? this.each(a) : this.queue(o.queue, a)
		},
		stop: function(e, t, n) {
			var i = function(e) {
				var t = e.stop;
				delete e.stop,
				t(n)
			};
			return "string" != typeof e && (n = t, t = e, e = void 0),
			t && e !== !1 && this.queue(e || "fx", []),
			this.each(function() {
				var t = !0,
				r = null != e && e + "queueHooks",
				o = Z.timers,
				a = vt.get(this);
				if (r) a[r] && a[r].stop && i(a[r]);
				else for (r in a) a[r] && a[r].stop && en.test(r) && i(a[r]);
				for (r = o.length; r--;) o[r].elem !== this || null != e && o[r].queue !== e || (o[r].anim.stop(n), t = !1, o.splice(r, 1)); (t || !n) && Z.dequeue(this, e)
			})
		},
		finish: function(e) {
			return e !== !1 && (e = e || "fx"),
			this.each(function() {
				var t, n = vt.get(this),
				i = n[e + "queue"],
				r = n[e + "queueHooks"],
				o = Z.timers,
				a = i ? i.length: 0;
				for (n.finish = !0, Z.queue(this, e, []), r && r.stop && r.stop.call(this, !0), t = o.length; t--;) o[t].elem === this && o[t].queue === e && (o[t].anim.stop(!0), o.splice(t, 1));
				for (t = 0; a > t; t++) i[t] && i[t].finish && i[t].finish.call(this);
				delete n.finish
			})
		}
	}),
	Z.each(["toggle", "show", "hide"],
	function(e, t) {
		var n = Z.fn[t];
		Z.fn[t] = function(e, i, r) {
			return null == e || "boolean" == typeof e ? n.apply(this, arguments) : this.animate(I(t, !0), e, i, r)
		}
	}),
	Z.each({
		slideDown: I("show"),
		slideUp: I("hide"),
		slideToggle: I("toggle"),
		fadeIn: {
			opacity: "show"
		},
		fadeOut: {
			opacity: "hide"
		},
		fadeToggle: {
			opacity: "toggle"
		}
	},
	function(e, t) {
		Z.fn[e] = function(e, n, i) {
			return this.animate(t, e, n, i)
		}
	}),
	Z.timers = [],
	Z.fx.tick = function() {
		var e, t = 0,
		n = Z.timers;
		for (Qt = Z.now(); t < n.length; t++) e = n[t],
		e() || n[t] !== e || n.splice(t--, 1);
		n.length || Z.fx.stop(),
		Qt = void 0
	},
	Z.fx.timer = function(e) {
		Z.timers.push(e),
		e() ? Z.fx.start() : Z.timers.pop()
	},
	Z.fx.interval = 13,
	Z.fx.start = function() {
		Vt || (Vt = setInterval(Z.fx.tick, Z.fx.interval))
	},
	Z.fx.stop = function() {
		clearInterval(Vt),
		Vt = null
	},
	Z.fx.speeds = {
		slow: 600,
		fast: 200,
		_default: 400
	},
	Z.fn.delay = function(e, t) {
		return e = Z.fx ? Z.fx.speeds[e] || e: e,
		t = t || "fx",
		this.queue(t,
		function(t, n) {
			var i = setTimeout(t, e);
			n.stop = function() {
				clearTimeout(i)
			}
		})
	},
	function() {
		var e = V.createElement("input"),
		t = V.createElement("select"),
		n = t.appendChild(V.createElement("option"));
		e.type = "checkbox",
		Q.checkOn = "" !== e.value,
		Q.optSelected = n.selected,
		t.disabled = !0,
		Q.optDisabled = !n.disabled,
		e = V.createElement("input"),
		e.value = "t",
		e.type = "radio",
		Q.radioValue = "t" === e.value
	} ();
	var rn, on, an = Z.expr.attrHandle;
	Z.fn.extend({
		attr: function(e, t) {
			return mt(this, Z.attr, e, t, arguments.length > 1)
		},
		removeAttr: function(e) {
			return this.each(function() {
				Z.removeAttr(this, e)
			})
		}
	}),
	Z.extend({
		attr: function(e, t, n) {
			var i, r, o = e.nodeType;
			return e && 3 !== o && 8 !== o && 2 !== o ? typeof e.getAttribute === St ? Z.prop(e, t, n) : (1 === o && Z.isXMLDoc(e) || (t = t.toLowerCase(), i = Z.attrHooks[t] || (Z.expr.match.bool.test(t) ? on: rn)), void 0 === n ? i && "get" in i && null !== (r = i.get(e, t)) ? r: (r = Z.find.attr(e, t), null == r ? void 0 : r) : null !== n ? i && "set" in i && void 0 !== (r = i.set(e, n, t)) ? r: (e.setAttribute(t, n + ""), n) : void Z.removeAttr(e, t)) : void 0
		},
		removeAttr: function(e, t) {
			var n, i, r = 0,
			o = t && t.match(pt);
			if (o && 1 === e.nodeType) for (; n = o[r++];) i = Z.propFix[n] || n,
			Z.expr.match.bool.test(n) && (e[i] = !1),
			e.removeAttribute(n)
		},
		attrHooks: {
			type: {
				set: function(e, t) {
					if (!Q.radioValue && "radio" === t && Z.nodeName(e, "input")) {
						var n = e.value;
						return e.setAttribute("type", t),
						n && (e.value = n),
						t
					}
				}
			}
		}
	}),
	on = {
		set: function(e, t, n) {
			return t === !1 ? Z.removeAttr(e, n) : e.setAttribute(n, n),
			n
		}
	},
	Z.each(Z.expr.match.bool.source.match(/\w+/g),
	function(e, t) {
		var n = an[t] || Z.find.attr;
		an[t] = function(e, t, i) {
			var r, o;
			return i || (o = an[t], an[t] = r, r = null != n(e, t, i) ? t.toLowerCase() : null, an[t] = o),
			r
		}
	});
	var sn = /^(?:input|select|textarea|button)$/i;
	Z.fn.extend({
		prop: function(e, t) {
			return mt(this, Z.prop, e, t, arguments.length > 1)
		},
		removeProp: function(e) {
			return this.each(function() {
				delete this[Z.propFix[e] || e]
			})
		}
	}),
	Z.extend({
		propFix: {
			"for": "htmlFor",
			"class": "className"
		},
		prop: function(e, t, n) {
			var i, r, o, a = e.nodeType;
			return e && 3 !== a && 8 !== a && 2 !== a ? (o = 1 !== a || !Z.isXMLDoc(e), o && (t = Z.propFix[t] || t, r = Z.propHooks[t]), void 0 !== n ? r && "set" in r && void 0 !== (i = r.set(e, n, t)) ? i: e[t] = n: r && "get" in r && null !== (i = r.get(e, t)) ? i: e[t]) : void 0
		},
		propHooks: {
			tabIndex: {
				get: function(e) {
					return e.hasAttribute("tabindex") || sn.test(e.nodeName) || e.href ? e.tabIndex: -1
				}
			}
		}
	}),
	Q.optSelected || (Z.propHooks.selected = {
		get: function(e) {
			var t = e.parentNode;
			return t && t.parentNode && t.parentNode.selectedIndex,
			null
		}
	}),
	Z.each(["tabIndex", "readOnly", "maxLength", "cellSpacing", "cellPadding", "rowSpan", "colSpan", "useMap", "frameBorder", "contentEditable"],
	function() {
		Z.propFix[this.toLowerCase()] = this
	});
	var ln = /[\t\r\n\f]/g;
	Z.fn.extend({
		addClass: function(e) {
			var t, n, i, r, o, a, s = "string" == typeof e && e,
			l = 0,
			c = this.length;
			if (Z.isFunction(e)) return this.each(function(t) {
				Z(this).addClass(e.call(this, t, this.className))
			});
			if (s) for (t = (e || "").match(pt) || []; c > l; l++) if (n = this[l], i = 1 === n.nodeType && (n.className ? (" " + n.className + " ").replace(ln, " ") : " ")) {
				for (o = 0; r = t[o++];) i.indexOf(" " + r + " ") < 0 && (i += r + " ");
				a = Z.trim(i),
				n.className !== a && (n.className = a)
			}
			return this
		},
		removeClass: function(e) {
			var t, n, i, r, o, a, s = 0 === arguments.length || "string" == typeof e && e,
			l = 0,
			c = this.length;
			if (Z.isFunction(e)) return this.each(function(t) {
				Z(this).removeClass(e.call(this, t, this.className))
			});
			if (s) for (t = (e || "").match(pt) || []; c > l; l++) if (n = this[l], i = 1 === n.nodeType && (n.className ? (" " + n.className + " ").replace(ln, " ") : "")) {
				for (o = 0; r = t[o++];) for (; i.indexOf(" " + r + " ") >= 0;) i = i.replace(" " + r + " ", " ");
				a = e ? Z.trim(i) : "",
				n.className !== a && (n.className = a)
			}
			return this
		},
		toggleClass: function(e, t) {
			var n = typeof e;
			return "boolean" == typeof t && "string" === n ? t ? this.addClass(e) : this.removeClass(e) : this.each(Z.isFunction(e) ?
			function(n) {
				Z(this).toggleClass(e.call(this, n, this.className, t), t)
			}: function() {
				if ("string" === n) for (var t, i = 0,
				r = Z(this), o = e.match(pt) || []; t = o[i++];) r.hasClass(t) ? r.removeClass(t) : r.addClass(t);
				else(n === St || "boolean" === n) && (this.className && vt.set(this, "__className__", this.className), this.className = this.className || e === !1 ? "": vt.get(this, "__className__") || "")
			})
		},
		hasClass: function(e) {
			for (var t = " " + e + " ",
			n = 0,
			i = this.length; i > n; n++) if (1 === this[n].nodeType && (" " + this[n].className + " ").replace(ln, " ").indexOf(t) >= 0) return ! 0;
			return ! 1
		}
	});
	var cn = /\r/g;
	Z.fn.extend({
		val: function(e) {
			var t, n, i, r = this[0];
			return arguments.length ? (i = Z.isFunction(e), this.each(function(n) {
				var r;
				1 === this.nodeType && (r = i ? e.call(this, n, Z(this).val()) : e, null == r ? r = "": "number" == typeof r ? r += "": Z.isArray(r) && (r = Z.map(r,
				function(e) {
					return null == e ? "": e + ""
				})), t = Z.valHooks[this.type] || Z.valHooks[this.nodeName.toLowerCase()], t && "set" in t && void 0 !== t.set(this, r, "value") || (this.value = r))
			})) : r ? (t = Z.valHooks[r.type] || Z.valHooks[r.nodeName.toLowerCase()], t && "get" in t && void 0 !== (n = t.get(r, "value")) ? n: (n = r.value, "string" == typeof n ? n.replace(cn, "") : null == n ? "": n)) : void 0
		}
	}),
	Z.extend({
		valHooks: {
			option: {
				get: function(e) {
					var t = Z.find.attr(e, "value");
					return null != t ? t: Z.trim(Z.text(e))
				}
			},
			select: {
				get: function(e) {
					for (var t, n, i = e.options,
					r = e.selectedIndex,
					o = "select-one" === e.type || 0 > r,
					a = o ? null: [], s = o ? r + 1 : i.length, l = 0 > r ? s: o ? r: 0; s > l; l++) if (n = i[l], !(!n.selected && l !== r || (Q.optDisabled ? n.disabled: null !== n.getAttribute("disabled")) || n.parentNode.disabled && Z.nodeName(n.parentNode, "optgroup"))) {
						if (t = Z(n).val(), o) return t;
						a.push(t)
					}
					return a
				},
				set: function(e, t) {
					for (var n, i, r = e.options,
					o = Z.makeArray(t), a = r.length; a--;) i = r[a],
					(i.selected = Z.inArray(i.value, o) >= 0) && (n = !0);
					return n || (e.selectedIndex = -1),
					o
				}
			}
		}
	}),
	Z.each(["radio", "checkbox"],
	function() {
		Z.valHooks[this] = {
			set: function(e, t) {
				return Z.isArray(t) ? e.checked = Z.inArray(Z(e).val(), t) >= 0 : void 0
			}
		},
		Q.checkOn || (Z.valHooks[this].get = function(e) {
			return null === e.getAttribute("value") ? "on": e.value
		})
	}),
	Z.each("blur focus focusin focusout load resize scroll unload click dblclick mousedown mouseup mousemove mouseover mouseout mouseenter mouseleave change select submit keydown keypress keyup error contextmenu".split(" "),
	function(e, t) {
		Z.fn[t] = function(e, n) {
			return arguments.length > 0 ? this.on(t, null, e, n) : this.trigger(t)
		}
	}),
	Z.fn.extend({
		hover: function(e, t) {
			return this.mouseenter(e).mouseleave(t || e)
		},
		bind: function(e, t, n) {
			return this.on(e, null, t, n)
		},
		unbind: function(e, t) {
			return this.off(e, null, t)
		},
		delegate: function(e, t, n, i) {
			return this.on(t, e, n, i)
		},
		undelegate: function(e, t, n) {
			return 1 === arguments.length ? this.off(e, "**") : this.off(t, e || "**", n)
		}
	});
	var un = Z.now(),
	dn = /\?/;
	Z.parseJSON = function(e) {
		return JSON.parse(e + "")
	},
	Z.parseXML = function(e) {
		var t, n;
		if (!e || "string" != typeof e) return null;
		try {
			n = new DOMParser,
			t = n.parseFromString(e, "text/xml")
		} catch(i) {
			t = void 0
		}
		return (!t || t.getElementsByTagName("parsererror").length) && Z.error("Invalid XML: " + e),
		t
	};
	var fn = /#.*$/,
	pn = /([?&])_=[^&]*/,
	hn = /^(.*?):[ \t]*([^\r\n]*)$/gm,
	gn = /^(?:about|app|app-storage|.+-extension|file|res|widget):$/,
	mn = /^(?:GET|HEAD)$/,
	vn = /^\/\//,
	yn = /^([\w.+-]+:)(?:\/\/(?:[^\/?#]*@|)([^\/?#:]*)(?::(\d+)|)|)/,
	bn = {},
	wn = {},
	xn = "*/".concat("*"),
	kn = e.location.href,
	Cn = yn.exec(kn.toLowerCase()) || [];
	Z.extend({
		active: 0,
		lastModified: {},
		etag: {},
		ajaxSettings: {
			url: kn,
			type: "GET",
			isLocal: gn.test(Cn[1]),
			global: !0,
			processData: !0,
			async: !0,
			contentType: "application/x-www-form-urlencoded; charset=UTF-8",
			accepts: {
				"*": xn,
				text: "text/plain",
				html: "text/html",
				xml: "application/xml, text/xml",
				json: "application/json, text/javascript"
			},
			contents: {
				xml: /xml/,
				html: /html/,
				json: /json/
			},
			responseFields: {
				xml: "responseXML",
				text: "responseText",
				json: "responseJSON"
			},
			converters: {
				"* text": String,
				"text html": !0,
				"text json": Z.parseJSON,
				"text xml": Z.parseXML
			},
			flatOptions: {
				url: !0,
				context: !0
			}
		},
		ajaxSetup: function(e, t) {
			return t ? P(P(e, Z.ajaxSettings), t) : P(Z.ajaxSettings, e)
		},
		ajaxPrefilter: H(bn),
		ajaxTransport: H(wn),
		ajax: function(e, t) {
			function n(e, t, n, a) {
				var l, u, v, y, w, k = t;
				2 !== b && (b = 2, s && clearTimeout(s), i = void 0, o = a || "", x.readyState = e > 0 ? 4 : 0, l = e >= 200 && 300 > e || 304 === e, n && (y = W(d, x, n)), y = O(d, y, x, l), l ? (d.ifModified && (w = x.getResponseHeader("Last-Modified"), w && (Z.lastModified[r] = w), w = x.getResponseHeader("etag"), w && (Z.etag[r] = w)), 204 === e || "HEAD" === d.type ? k = "nocontent": 304 === e ? k = "notmodified": (k = y.state, u = y.data, v = y.error, l = !v)) : (v = k, (e || !k) && (k = "error", 0 > e && (e = 0))), x.status = e, x.statusText = (t || k) + "", l ? h.resolveWith(f, [u, k, x]) : h.rejectWith(f, [x, k, v]), x.statusCode(m), m = void 0, c && p.trigger(l ? "ajaxSuccess": "ajaxError", [x, d, l ? u: v]), g.fireWith(f, [x, k]), c && (p.trigger("ajaxComplete", [x, d]), --Z.active || Z.event.trigger("ajaxStop")))
			}
			"object" == typeof e && (t = e, e = void 0),
			t = t || {};
			var i, r, o, a, s, l, c, u, d = Z.ajaxSetup({},
			t),
			f = d.context || d,
			p = d.context && (f.nodeType || f.jquery) ? Z(f) : Z.event,
			h = Z.Deferred(),
			g = Z.Callbacks("once memory"),
			m = d.statusCode || {},
			v = {},
			y = {},
			b = 0,
			w = "canceled",
			x = {
				readyState: 0,
				getResponseHeader: function(e) {
					var t;
					if (2 === b) {
						if (!a) for (a = {}; t = hn.exec(o);) a[t[1].toLowerCase()] = t[2];
						t = a[e.toLowerCase()]
					}
					return null == t ? null: t
				},
				getAllResponseHeaders: function() {
					return 2 === b ? o: null
				},
				setRequestHeader: function(e, t) {
					var n = e.toLowerCase();
					return b || (e = y[n] = y[n] || e, v[e] = t),
					this
				},
				overrideMimeType: function(e) {
					return b || (d.mimeType = e),
					this
				},
				statusCode: function(e) {
					var t;
					if (e) if (2 > b) for (t in e) m[t] = [m[t], e[t]];
					else x.always(e[x.status]);
					return this
				},
				abort: function(e) {
					var t = e || w;
					return i && i.abort(t),
					n(0, t),
					this
				}
			};
			if (h.promise(x).complete = g.add, x.success = x.done, x.error = x.fail, d.url = ((e || d.url || kn) + "").replace(fn, "").replace(vn, Cn[1] + "//"), d.type = t.method || t.type || d.method || d.type, d.dataTypes = Z.trim(d.dataType || "*").toLowerCase().match(pt) || [""], null == d.crossDomain && (l = yn.exec(d.url.toLowerCase()), d.crossDomain = !(!l || l[1] === Cn[1] && l[2] === Cn[2] && (l[3] || ("http:" === l[1] ? "80": "443")) === (Cn[3] || ("http:" === Cn[1] ? "80": "443")))), d.data && d.processData && "string" != typeof d.data && (d.data = Z.param(d.data, d.traditional)), M(bn, d, t, x), 2 === b) return x;
			c = Z.event && d.global,
			c && 0 === Z.active++&&Z.event.trigger("ajaxStart"),
			d.type = d.type.toUpperCase(),
			d.hasContent = !mn.test(d.type),
			r = d.url,
			d.hasContent || (d.data && (r = d.url += (dn.test(r) ? "&": "?") + d.data, delete d.data), d.cache === !1 && (d.url = pn.test(r) ? r.replace(pn, "$1_=" + un++) : r + (dn.test(r) ? "&": "?") + "_=" + un++)),
			d.ifModified && (Z.lastModified[r] && x.setRequestHeader("If-Modified-Since", Z.lastModified[r]), Z.etag[r] && x.setRequestHeader("If-None-Match", Z.etag[r])),
			(d.data && d.hasContent && d.contentType !== !1 || t.contentType) && x.setRequestHeader("Content-Type", d.contentType),
			x.setRequestHeader("Accept", d.dataTypes[0] && d.accepts[d.dataTypes[0]] ? d.accepts[d.dataTypes[0]] + ("*" !== d.dataTypes[0] ? ", " + xn + "; q=0.01": "") : d.accepts["*"]);
			for (u in d.headers) x.setRequestHeader(u, d.headers[u]);
			if (d.beforeSend && (d.beforeSend.call(f, x, d) === !1 || 2 === b)) return x.abort();
			w = "abort";
			for (u in {
				success: 1,
				error: 1,
				complete: 1
			}) x[u](d[u]);
			if (i = M(wn, d, t, x)) {
				x.readyState = 1,
				c && p.trigger("ajaxSend", [x, d]),
				d.async && d.timeout > 0 && (s = setTimeout(function() {
					x.abort("timeout")
				},
				d.timeout));
				try {
					b = 1,
					i.send(v, n)
				} catch(k) {
					if (! (2 > b)) throw k;
					n( - 1, k)
				}
			} else n( - 1, "No Transport");
			return x
		},
		getJSON: function(e, t, n) {
			return Z.get(e, t, n, "json")
		},
		getScript: function(e, t) {
			return Z.get(e, void 0, t, "script")
		}
	}),
	Z.each(["get", "post"],
	function(e, t) {
		Z[t] = function(e, n, i, r) {
			return Z.isFunction(n) && (r = r || i, i = n, n = void 0),
			Z.ajax({
				url: e,
				type: t,
				dataType: r,
				data: n,
				success: i
			})
		}
	}),
	Z._evalUrl = function(e) {
		return Z.ajax({
			url: e,
			type: "GET",
			dataType: "script",
			async: !1,
			global: !1,
			"throws": !0
		})
	},
	Z.fn.extend({
		wrapAll: function(e) {
			var t;
			return Z.isFunction(e) ? this.each(function(t) {
				Z(this).wrapAll(e.call(this, t))
			}) : (this[0] && (t = Z(e, this[0].ownerDocument).eq(0).clone(!0), this[0].parentNode && t.insertBefore(this[0]), t.map(function() {
				for (var e = this; e.firstElementChild;) e = e.firstElementChild;
				return e
			}).append(this)), this)
		},
		wrapInner: function(e) {
			return this.each(Z.isFunction(e) ?
			function(t) {
				Z(this).wrapInner(e.call(this, t))
			}: function() {
				var t = Z(this),
				n = t.contents();
				n.length ? n.wrapAll(e) : t.append(e)
			})
		},
		wrap: function(e) {
			var t = Z.isFunction(e);
			return this.each(function(n) {
				Z(this).wrapAll(t ? e.call(this, n) : e)
			})
		},
		unwrap: function() {
			return this.parent().each(function() {
				Z.nodeName(this, "body") || Z(this).replaceWith(this.childNodes)
			}).end()
		}
	}),
	Z.expr.filters.hidden = function(e) {
		return e.offsetWidth <= 0 && e.offsetHeight <= 0
	},
	Z.expr.filters.visible = function(e) {
		return ! Z.expr.filters.hidden(e)
	};
	var Tn = /%20/g,
	Sn = /\[\]$/,
	_n = /\r?\n/g,
	$n = /^(?:submit|button|image|reset|file)$/i,
	Ln = /^(?:input|select|textarea|keygen)/i;
	Z.param = function(e, t) {
		var n, i = [],
		r = function(e, t) {
			t = Z.isFunction(t) ? t() : null == t ? "": t,
			i[i.length] = encodeURIComponent(e) + "=" + encodeURIComponent(t)
		};
		if (void 0 === t && (t = Z.ajaxSettings && Z.ajaxSettings.traditional), Z.isArray(e) || e.jquery && !Z.isPlainObject(e)) Z.each(e,
		function() {
			r(this.name, this.value)
		});
		else for (n in e) R(n, e[n], t, r);
		return i.join("&").replace(Tn, "+")
	},
	Z.fn.extend({
		serialize: function() {
			return Z.param(this.serializeArray())
		},
		serializeArray: function() {
			return this.map(function() {
				var e = Z.prop(this, "elements");
				return e ? Z.makeArray(e) : this
			}).filter(function() {
				var e = this.type;
				return this.name && !Z(this).is(":disabled") && Ln.test(this.nodeName) && !$n.test(e) && (this.checked || !Tt.test(e))
			}).map(function(e, t) {
				var n = Z(this).val();
				return null == n ? null: Z.isArray(n) ? Z.map(n,
				function(e) {
					return {
						name: t.name,
						value: e.replace(_n, "\r\n")
					}
				}) : {
					name: t.name,
					value: n.replace(_n, "\r\n")
				}
			}).get()
		}
	}),
	Z.ajaxSettings.xhr = function() {
		try {
			return new XMLHttpRequest
		} catch(e) {}
	};
	var En = 0,
	In = {},
	Nn = {
		0 : 200,
		1223 : 204
	},
	An = Z.ajaxSettings.xhr();
	e.attachEvent && e.attachEvent("onunload",
	function() {
		for (var e in In) In[e]()
	}),
	Q.cors = !!An && "withCredentials" in An,
	Q.ajax = An = !!An,
	Z.ajaxTransport(function(e) {
		var t;
		return Q.cors || An && !e.crossDomain ? {
			send: function(n, i) {
				var r, o = e.xhr(),
				a = ++En;
				if (o.open(e.type, e.url, e.async, e.username, e.password), e.xhrFields) for (r in e.xhrFields) o[r] = e.xhrFields[r];
				e.mimeType && o.overrideMimeType && o.overrideMimeType(e.mimeType),
				e.crossDomain || n["X-Requested-With"] || (n["X-Requested-With"] = "XMLHttpRequest");
				for (r in n) o.setRequestHeader(r, n[r]);
				t = function(e) {
					return function() {
						t && (delete In[a], t = o.onload = o.onerror = null, "abort" === e ? o.abort() : "error" === e ? i(o.status, o.statusText) : i(Nn[o.status] || o.status, o.statusText, "string" == typeof o.responseText ? {
							text: o.responseText
						}: void 0, o.getAllResponseHeaders()))
					}
				},
				o.onload = t(),
				o.onerror = t("error"),
				t = In[a] = t("abort");
				try {
					o.send(e.hasContent && e.data || null)
				} catch(s) {
					if (t) throw s
				}
			},
			abort: function() {
				t && t()
			}
		}: void 0
	}),
	Z.ajaxSetup({
		accepts: {
			script: "text/javascript, application/javascript, application/ecmascript, application/x-ecmascript"
		},
		contents: {
			script: /(?:java|ecma)script/
		},
		converters: {
			"text script": function(e) {
				return Z.globalEval(e),
				e
			}
		}
	}),
	Z.ajaxPrefilter("script",
	function(e) {
		void 0 === e.cache && (e.cache = !1),
		e.crossDomain && (e.type = "GET")
	}),
	Z.ajaxTransport("script",
	function(e) {
		if (e.crossDomain) {
			var t, n;
			return {
				send: function(i, r) {
					t = Z("<script>").prop({
						async: !0,
						charset: e.scriptCharset,
						src: e.url
					}).on("load error", n = function(e) {
						t.remove(),
						n = null,
						e && r("error" === e.type ? 404 : 200, e.type)
					}),
					V.head.appendChild(t[0])
				},
				abort: function() {
					n && n()
				}
			}
		}
	});
	var Dn = [],
	jn = /(=)\?(?=&|$)|\?\?/;
	Z.ajaxSetup({
		jsonp: "callback",
		jsonpCallback: function() {
			var e = Dn.pop() || Z.expando + "_" + un++;
			return this[e] = !0,
			e
		}
	}),
	Z.ajaxPrefilter("json jsonp",
	function(t, n, i) {
		var r, o, a, s = t.jsonp !== !1 && (jn.test(t.url) ? "url": "string" == typeof t.data && !(t.contentType || "").indexOf("application/x-www-form-urlencoded") && jn.test(t.data) && "data");
		return s || "jsonp" === t.dataTypes[0] ? (r = t.jsonpCallback = Z.isFunction(t.jsonpCallback) ? t.jsonpCallback() : t.jsonpCallback, s ? t[s] = t[s].replace(jn, "$1" + r) : t.jsonp !== !1 && (t.url += (dn.test(t.url) ? "&": "?") + t.jsonp + "=" + r), t.converters["script json"] = function() {
			return a || Z.error(r + " was not called"),
			a[0]
		},
		t.dataTypes[0] = "json", o = e[r], e[r] = function() {
			a = arguments
		},
		i.always(function() {
			e[r] = o,
			t[r] && (t.jsonpCallback = n.jsonpCallback, Dn.push(r)),
			a && Z.isFunction(o) && o(a[0]),
			a = o = void 0
		}), "script") : void 0
	}),
	Z.parseHTML = function(e, t, n) {
		if (!e || "string" != typeof e) return null;
		"boolean" == typeof t && (n = t, t = !1),
		t = t || V;
		var i = at.exec(e),
		r = !n && [];
		return i ? [t.createElement(i[1])] : (i = Z.buildFragment([e], t, r), r && r.length && Z(r).remove(), Z.merge([], i.childNodes))
	};
	var Hn = Z.fn.load;
	Z.fn.load = function(e, t, n) {
		if ("string" != typeof e && Hn) return Hn.apply(this, arguments);
		var i, r, o, a = this,
		s = e.indexOf(" ");
		return s >= 0 && (i = Z.trim(e.slice(s)), e = e.slice(0, s)),
		Z.isFunction(t) ? (n = t, t = void 0) : t && "object" == typeof t && (r = "POST"),
		a.length > 0 && Z.ajax({
			url: e,
			type: r,
			dataType: "html",
			data: t
		}).done(function(e) {
			o = arguments,
			a.html(i ? Z("<div>").append(Z.parseHTML(e)).find(i) : e)
		}).complete(n &&
		function(e, t) {
			a.each(n, o || [e.responseText, t, e])
		}),
		this
	},
	Z.each(["ajaxStart", "ajaxStop", "ajaxComplete", "ajaxError", "ajaxSuccess", "ajaxSend"],
	function(e, t) {
		Z.fn[t] = function(e) {
			return this.on(t, e)
		}
	}),
	Z.expr.filters.animated = function(e) {
		return Z.grep(Z.timers,
		function(t) {
			return e === t.elem
		}).length
	};
	var Mn = e.document.documentElement;
	Z.offset = {
		setOffset: function(e, t, n) {
			var i, r, o, a, s, l, c, u = Z.css(e, "position"),
			d = Z(e),
			f = {};
			"static" === u && (e.style.position = "relative"),
			s = d.offset(),
			o = Z.css(e, "top"),
			l = Z.css(e, "left"),
			c = ("absolute" === u || "fixed" === u) && (o + l).indexOf("auto") > -1,
			c ? (i = d.position(), a = i.top, r = i.left) : (a = parseFloat(o) || 0, r = parseFloat(l) || 0),
			Z.isFunction(t) && (t = t.call(e, n, s)),
			null != t.top && (f.top = t.top - s.top + a),
			null != t.left && (f.left = t.left - s.left + r),
			"using" in t ? t.using.call(e, f) : d.css(f)
		}
	},
	Z.fn.extend({
		offset: function(e) {
			if (arguments.length) return void 0 === e ? this: this.each(function(t) {
				Z.offset.setOffset(this, e, t)
			});
			var t, n, i = this[0],
			r = {
				top: 0,
				left: 0
			},
			o = i && i.ownerDocument;
			return o ? (t = o.documentElement, Z.contains(t, i) ? (typeof i.getBoundingClientRect !== St && (r = i.getBoundingClientRect()), n = q(o), {
				top: r.top + n.pageYOffset - t.clientTop,
				left: r.left + n.pageXOffset - t.clientLeft
			}) : r) : void 0
		},
		position: function() {
			if (this[0]) {
				var e, t, n = this[0],
				i = {
					top: 0,
					left: 0
				};
				return "fixed" === Z.css(n, "position") ? t = n.getBoundingClientRect() : (e = this.offsetParent(), t = this.offset(), Z.nodeName(e[0], "html") || (i = e.offset()), i.top += Z.css(e[0], "borderTopWidth", !0), i.left += Z.css(e[0], "borderLeftWidth", !0)),
				{
					top: t.top - i.top - Z.css(n, "marginTop", !0),
					left: t.left - i.left - Z.css(n, "marginLeft", !0)
				}
			}
		},
		offsetParent: function() {
			return this.map(function() {
				for (var e = this.offsetParent || Mn; e && !Z.nodeName(e, "html") && "static" === Z.css(e, "position");) e = e.offsetParent;
				return e || Mn
			})
		}
	}),
	Z.each({
		scrollLeft: "pageXOffset",
		scrollTop: "pageYOffset"
	},
	function(t, n) {
		var i = "pageYOffset" === n;
		Z.fn[t] = function(r) {
			return mt(this,
			function(t, r, o) {
				var a = q(t);
				return void 0 === o ? a ? a[n] : t[r] : void(a ? a.scrollTo(i ? e.pageXOffset: o, i ? o: e.pageYOffset) : t[r] = o)
			},
			t, r, arguments.length, null)
		}
	}),
	Z.each(["top", "left"],
	function(e, t) {
		Z.cssHooks[t] = k(Q.pixelPosition,
		function(e, n) {
			return n ? (n = x(e, t), Xt.test(n) ? Z(e).position()[t] + "px": n) : void 0
		})
	}),
	Z.each({
		Height: "height",
		Width: "width"
	},
	function(e, t) {
		Z.each({
			padding: "inner" + e,
			content: t,
			"": "outer" + e
		},
		function(n, i) {
			Z.fn[i] = function(i, r) {
				var o = arguments.length && (n || "boolean" != typeof i),
				a = n || (i === !0 || r === !0 ? "margin": "border");
				return mt(this,
				function(t, n, i) {
					var r;
					return Z.isWindow(t) ? t.document.documentElement["client" + e] : 9 === t.nodeType ? (r = t.documentElement, Math.max(t.body["scroll" + e], r["scroll" + e], t.body["offset" + e], r["offset" + e], r["client" + e])) : void 0 === i ? Z.css(t, n, a) : Z.style(t, n, i, a)
				},
				t, o ? i: void 0, o, null)
			}
		})
	}),
	Z.fn.size = function() {
		return this.length
	},
	Z.fn.andSelf = Z.fn.addBack,
	"function" == typeof define && define.amd && define("jquery", [],
	function() {
		return Z
	});
	var Pn = e.jQuery,
	Wn = e.$;
	return Z.noConflict = function(t) {
		return e.$ === Z && (e.$ = Wn),
		t && e.jQuery === Z && (e.jQuery = Pn),
		Z
	},
	typeof t === St && (e.jQuery = e.$ = Z),
	Z
}),
function() {
	var e, t;
	e = this.jQuery || window.jQuery,
	t = e(window),
	e.fn.stick_in_parent = function(n) {
		var i, r, o, a, s, l, c, u, d, f, p;
		for (null == n && (n = {}), p = n.sticky_class, s = n.inner_scrolling, f = n.recalc_every, d = n.parent, u = n.offset_top, c = n.spacer, r = n.bottoming, null == u && (u = 0), null == d && (d = void 0), null == s && (s = !0), null == p && (p = "is_stuck"), i = e(document), null == r && (r = !0), o = function(n, o, a, l, h, g, m, v) {
			var y, b, w, x, k, C, T, S, _, $, L, E;
			if (!n.data("sticky_kit")) {
				if (n.data("sticky_kit", !0), k = i.height(), T = n.parent(), null != d && (T = T.closest(d)), !T.length) throw "failed to find stick parent";
				if (y = w = !1, (L = null != c ? c && n.closest(c) : e("<div />")) && L.css("position", n.css("position")), S = function() {
					var e, t, r;
					return ! v && (k = i.height(), e = parseInt(T.css("border-top-width"), 10), t = parseInt(T.css("padding-top"), 10), o = parseInt(T.css("padding-bottom"), 10), a = T.offset().top + e + t, l = T.height(), w && (y = w = !1, null == c && (n.insertAfter(L), L.detach()), n.css({
						position: "",
						top: "",
						width: "",
						bottom: ""
					}).removeClass(p), r = !0), h = n.offset().top - (parseInt(n.css("margin-top"), 10) || 0) - u, g = n.outerHeight(!0), m = n.css("float"), L && L.css({
						width: n.outerWidth(!0),
						height: g,
						display: n.css("display"),
						"vertical-align": n.css("vertical-align"),
						"float": m
					}), r) ? E() : void 0
				},
				S(), g !== l) return x = void 0,
				C = u,
				$ = f,
				E = function() {
					var e, d, b, _;
					return ! v && (b = !1, null != $ && (--$, 0 >= $ && ($ = f, S(), b = !0)), b || i.height() === k || S(), b = t.scrollTop(), null != x && (d = b - x), x = b, w ? (r && (_ = b + g + C > l + a, y && !_ && (y = !1, n.css({
						position: "fixed",
						bottom: "",
						top: C
					}).trigger("sticky_kit:unbottom"))), h > b && (w = !1, C = u, null == c && ("left" !== m && "right" !== m || n.insertAfter(L), L.detach()), e = {
						position: "",
						width: "",
						top: ""
					},
					n.css(e).removeClass(p).trigger("sticky_kit:unstick")), s && (e = t.height(), g + u > e && !y && (C -= d, C = Math.max(e - g, C), C = Math.min(u, C), w && n.css({
						top: C + "px"
					})))) : b > h && (w = !0, e = {
						position: "fixed",
						top: C
					},
					e.width = "border-box" === n.css("box-sizing") ? n.outerWidth() + "px": n.width() + "px", n.css(e).addClass(p), null == c && (n.after(L), "left" !== m && "right" !== m || L.append(n)), n.trigger("sticky_kit:stick")), w && r && (null == _ && (_ = b + g + C > l + a), !y && _)) ? (y = !0, "static" === T.css("position") && T.css({
						position: "relative"
					}), n.css({
						position: "absolute",
						bottom: o,
						top: "auto"
					}).trigger("sticky_kit:bottom")) : void 0
				},
				_ = function() {
					return S(),
					E()
				},
				b = function() {
					return v = !0,
					t.off("touchmove", E),
					t.off("scroll", E),
					t.off("resize", _),
					e(document.body).off("sticky_kit:recalc", _),
					n.off("sticky_kit:detach", b),
					n.removeData("sticky_kit"),
					n.css({
						position: "",
						bottom: "",
						top: "",
						width: ""
					}),
					T.position("position", ""),
					w ? (null == c && ("left" !== m && "right" !== m || n.insertAfter(L), L.remove()), n.removeClass(p)) : void 0
				},
				t.on("touchmove", E),
				t.on("scroll", E),
				t.on("resize", _),
				e(document.body).on("sticky_kit:recalc", _),
				n.on("sticky_kit:detach", b),
				setTimeout(E, 0)
			}
		},
		a = 0, l = this.length; l > a; a++) n = this[a],
		o(e(n));
		return this
	}
}.call(this),
!
function e(t, n, i) {
	function r(a, s) {
		if (!n[a]) {
			if (!t[a]) {
				var l = "function" == typeof require && require;
				if (!s && l) return l(a, !0);
				if (o) return o(a, !0);
				var c = new Error("Cannot find module '" + a + "'");
				throw c.code = "MODULE_NOT_FOUND",
				c
			}
			var u = n[a] = {
				exports: {}
			};
			t[a][0].call(u.exports,
			function(e) {
				var n = t[a][1][e];
				return r(n ? n: e)
			},
			u, u.exports, e, t, n, i)
		}
		return n[a].exports
	}
	for (var o = "function" == typeof require && require,
	a = 0; a < i.length; a++) r(i[a]);
	return r
} ({
	1 : [function(e, t) {
		"use strict";
		function n(e) {
			e.fn.perfectScrollbar = function(t) {
				return this.each(function() {
					if ("object" == typeof t || "undefined" == typeof t) {
						var n = t;
						r.get(this) || i.initialize(this, n)
					} else {
						var o = t;
						"update" === o ? i.update(this) : "destroy" === o && i.destroy(this)
					}
					return e(this)
				})
			}
		}
		var i = e("../main"),
		r = e("../plugin/instances");
		if ("function" == typeof define && define.amd) define(["jquery"], n);
		else {
			var o = window.jQuery ? window.jQuery: window.$;
			"undefined" != typeof o && n(o)
		}
		t.exports = n
	},
	{
		"../main": 7,
		"../plugin/instances": 18
	}],
	2 : [function(e, t, n) {
		"use strict";
		function i(e, t) {
			var n = e.className.split(" ");
			n.indexOf(t) < 0 && n.push(t),
			e.className = n.join(" ")
		}
		function r(e, t) {
			var n = e.className.split(" "),
			i = n.indexOf(t);
			i >= 0 && n.splice(i, 1),
			e.className = n.join(" ")
		}
		n.add = function(e, t) {
			e.classList ? e.classList.add(t) : i(e, t)
		},
		n.remove = function(e, t) {
			e.classList ? e.classList.remove(t) : r(e, t)
		},
		n.list = function(e) {
			return e.classList ? e.classList: e.className.split(" ")
		}
	},
	{}],
	3 : [function(e, t, n) {
		"use strict";
		function i(e, t) {
			return window.getComputedStyle(e)[t]
		}
		function r(e, t, n) {
			return "number" == typeof n && (n = n.toString() + "px"),
			e.style[t] = n,
			e
		}
		function o(e, t) {
			for (var n in t) {
				var i = t[n];
				"number" == typeof i && (i = i.toString() + "px"),
				e.style[n] = i
			}
			return e
		}
		n.e = function(e, t) {
			var n = document.createElement(e);
			return n.className = t,
			n
		},
		n.appendTo = function(e, t) {
			return t.appendChild(e),
			e
		},
		n.css = function(e, t, n) {
			return "object" == typeof t ? o(e, t) : "undefined" == typeof n ? i(e, t) : r(e, t, n)
		},
		n.matches = function(e, t) {
			return "undefined" != typeof e.matches ? e.matches(t) : "undefined" != typeof e.matchesSelector ? e.matchesSelector(t) : "undefined" != typeof e.webkitMatchesSelector ? e.webkitMatchesSelector(t) : "undefined" != typeof e.mozMatchesSelector ? e.mozMatchesSelector(t) : "undefined" != typeof e.msMatchesSelector ? e.msMatchesSelector(t) : void 0
		},
		n.remove = function(e) {
			"undefined" != typeof e.remove ? e.remove() : e.parentNode && e.parentNode.removeChild(e)
		}
	},
	{}],
	4 : [function(e, t) {
		"use strict";
		var n = function(e) {
			this.element = e,
			this.events = {}
		};
		n.prototype.bind = function(e, t) {
			"undefined" == typeof this.events[e] && (this.events[e] = []),
			this.events[e].push(t),
			this.element.addEventListener(e, t, !1)
		},
		n.prototype.unbind = function(e, t) {
			var n = "undefined" != typeof t;
			this.events[e] = this.events[e].filter(function(i) {
				return n && i !== t ? !0 : (this.element.removeEventListener(e, i, !1), !1)
			},
			this)
		},
		n.prototype.unbindAll = function() {
			for (var e in this.events) this.unbind(e)
		};
		var i = function() {
			this.eventElements = []
		};
		i.prototype.eventElement = function(e) {
			var t = this.eventElements.filter(function(t) {
				return t.element === e
			})[0];
			return "undefined" == typeof t && (t = new n(e), this.eventElements.push(t)),
			t
		},
		i.prototype.bind = function(e, t, n) {
			this.eventElement(e).bind(t, n)
		},
		i.prototype.unbind = function(e, t, n) {
			this.eventElement(e).unbind(t, n)
		},
		i.prototype.unbindAll = function() {
			for (var e = 0; e < this.eventElements.length; e++) this.eventElements[e].unbindAll()
		},
		i.prototype.once = function(e, t, n) {
			var i = this.eventElement(e),
			r = function(e) {
				i.unbind(t, r),
				n(e)
			};
			i.bind(t, r)
		},
		t.exports = i
	},
	{}],
	5 : [function(e, t) {
		"use strict";
		t.exports = function() {
			function e() {
				return Math.floor(65536 * (1 + Math.random())).toString(16).substring(1)
			}
			return function() {
				return e() + e() + "-" + e() + "-" + e() + "-" + e() + "-" + e() + e() + e()
			}
		} ()
	},
	{}],
	6 : [function(e, t, n) {
		"use strict";
		var i = e("./class"),
		r = e("./dom");
		n.toInt = function(e) {
			return "string" == typeof e ? parseInt(e, 10) : ~~e
		},
		n.clone = function(e) {
			if (null === e) return null;
			if ("object" == typeof e) {
				var t = {};
				for (var n in e) t[n] = this.clone(e[n]);
				return t
			}
			return e
		},
		n.extend = function(e, t) {
			var n = this.clone(e);
			for (var i in t) n[i] = this.clone(t[i]);
			return n
		},
		n.isEditable = function(e) {
			return r.matches(e, "input,[contenteditable]") || r.matches(e, "select,[contenteditable]") || r.matches(e, "textarea,[contenteditable]") || r.matches(e, "button,[contenteditable]")
		},
		n.removePsClasses = function(e) {
			for (var t = i.list(e), n = 0; n < t.length; n++) {
				var r = t[n];
				0 === r.indexOf("ps-") && i.remove(e, r)
			}
		},
		n.outerWidth = function(e) {
			return this.toInt(r.css(e, "width")) + this.toInt(r.css(e, "paddingLeft")) + this.toInt(r.css(e, "paddingRight")) + this.toInt(r.css(e, "borderLeftWidth")) + this.toInt(r.css(e, "borderRightWidth"))
		},
		n.startScrolling = function(e, t) {
			i.add(e, "ps-in-scrolling"),
			"undefined" != typeof t ? i.add(e, "ps-" + t) : (i.add(e, "ps-x"), i.add(e, "ps-y"))
		},
		n.stopScrolling = function(e, t) {
			i.remove(e, "ps-in-scrolling"),
			"undefined" != typeof t ? i.remove(e, "ps-" + t) : (i.remove(e, "ps-x"), i.remove(e, "ps-y"))
		},
		n.env = {
			isWebKit: "WebkitAppearance" in document.documentElement.style,
			supportsTouch: "ontouchstart" in window || window.DocumentTouch && document instanceof window.DocumentTouch,
			supportsIePointer: null !== window.navigator.msMaxTouchPoints
		}
	},
	{
		"./class": 2,
		"./dom": 3
	}],
	7 : [function(e, t) {
		"use strict";
		var n = e("./plugin/destroy"),
		i = e("./plugin/initialize"),
		r = e("./plugin/update");
		t.exports = {
			initialize: i,
			update: r,
			destroy: n
		}
	},
	{
		"./plugin/destroy": 9,
		"./plugin/initialize": 17,
		"./plugin/update": 20
	}],
	8 : [function(e, t) {
		"use strict";
		t.exports = {
			wheelSpeed: 1,
			wheelPropagation: !1,
			swipePropagation: !0,
			minScrollbarLength: null,
			maxScrollbarLength: null,
			useBothWheelAxes: !1,
			useKeyboard: !0,
			suppressScrollX: !1,
			suppressScrollY: !1,
			scrollXMarginOffset: 0,
			scrollYMarginOffset: 0
		}
	},
	{}],
	9 : [function(e, t) {
		"use strict";
		var n = e("../lib/dom"),
		i = e("../lib/helper"),
		r = e("./instances");
		t.exports = function(e) {
			var t = r.get(e);
			t.event.unbindAll(),
			n.remove(t.scrollbarX),
			n.remove(t.scrollbarY),
			n.remove(t.scrollbarXRail),
			n.remove(t.scrollbarYRail),
			i.removePsClasses(e),
			r.remove(e)
		}
	},
	{
		"../lib/dom": 3,
		"../lib/helper": 6,
		"./instances": 18
	}],
	10 : [function(e, t) {
		"use strict";
		function n(e, t) {
			function n(e) {
				return e.getBoundingClientRect()
			}
			var r = window.Event.prototype.stopPropagation.bind;
			t.event.bind(t.scrollbarY, "click", r),
			t.event.bind(t.scrollbarYRail, "click",
			function(r) {
				var a = i.toInt(t.scrollbarYHeight / 2),
				s = r.pageY - n(t.scrollbarYRail).top - a,
				l = t.containerHeight - t.scrollbarYHeight,
				c = s / l;
				0 > c ? c = 0 : c > 1 && (c = 1),
				e.scrollTop = (t.contentHeight - t.containerHeight) * c,
				o(e)
			}),
			t.event.bind(t.scrollbarX, "click", r),
			t.event.bind(t.scrollbarXRail, "click",
			function(r) {
				var a = i.toInt(t.scrollbarXWidth / 2),
				s = r.pageX - n(t.scrollbarXRail).left - a;
				console.log(r.pageX, t.scrollbarXRail.offsetLeft);
				var l = t.containerWidth - t.scrollbarXWidth,
				c = s / l;
				0 > c ? c = 0 : c > 1 && (c = 1),
				e.scrollLeft = (t.contentWidth - t.containerWidth) * c,
				o(e)
			})
		}
		var i = e("../../lib/helper"),
		r = e("../instances"),
		o = e("../update-geometry");
		t.exports = function(e) {
			var t = r.get(e);
			n(e, t)
		}
	},
	{
		"../../lib/helper": 6,
		"../instances": 18,
		"../update-geometry": 19
	}],
	11 : [function(e, t) {
		"use strict";
		function n(e, t) {
			function n(n) {
				var r = i + n,
				a = t.containerWidth - t.scrollbarXWidth;
				t.scrollbarXLeft = 0 > r ? 0 : r > a ? a: r;
				var s = o.toInt(t.scrollbarXLeft * (t.contentWidth - t.containerWidth) / (t.containerWidth - t.scrollbarXWidth));
				e.scrollLeft = s
			}
			var i = null,
			a = null,
			l = function(t) {
				n(t.pageX - a),
				s(e),
				t.stopPropagation(),
				t.preventDefault()
			},
			c = function() {
				o.stopScrolling(e, "x"),
				t.event.unbind(t.ownerDocument, "mousemove", l)
			};
			t.event.bind(t.scrollbarX, "mousedown",
			function(n) {
				a = n.pageX,
				i = o.toInt(r.css(t.scrollbarX, "left")),
				o.startScrolling(e, "x"),
				t.event.bind(t.ownerDocument, "mousemove", l),
				t.event.once(t.ownerDocument, "mouseup", c),
				n.stopPropagation(),
				n.preventDefault()
			})
		}
		function i(e, t) {
			function n(n) {
				var r = i + n,
				a = t.containerHeight - t.scrollbarYHeight;
				t.scrollbarYTop = 0 > r ? 0 : r > a ? a: r;
				var s = o.toInt(t.scrollbarYTop * (t.contentHeight - t.containerHeight) / (t.containerHeight - t.scrollbarYHeight));
				e.scrollTop = s
			}
			var i = null,
			a = null,
			l = function(t) {
				n(t.pageY - a),
				s(e),
				t.stopPropagation(),
				t.preventDefault()
			},
			c = function() {
				o.stopScrolling(e, "y"),
				t.event.unbind(t.ownerDocument, "mousemove", l)
			};
			t.event.bind(t.scrollbarY, "mousedown",
			function(n) {
				a = n.pageY,
				i = o.toInt(r.css(t.scrollbarY, "top")),
				o.startScrolling(e, "y"),
				t.event.bind(t.ownerDocument, "mousemove", l),
				t.event.once(t.ownerDocument, "mouseup", c),
				n.stopPropagation(),
				n.preventDefault()
			})
		}
		var r = e("../../lib/dom"),
		o = e("../../lib/helper"),
		a = e("../instances"),
		s = e("../update-geometry");
		t.exports = function(e) {
			var t = a.get(e);
			n(e, t),
			i(e, t)
		}
	},
	{
		"../../lib/dom": 3,
		"../../lib/helper": 6,
		"../instances": 18,
		"../update-geometry": 19
	}],
	12 : [function(e, t) {
		"use strict";
		function n(e, t) {
			function n(n, i) {
				var r = e.scrollTop;
				if (0 === n) {
					if (!t.scrollbarYActive) return ! 1;
					if (0 === r && i > 0 || r >= t.contentHeight - t.containerHeight && 0 > i) return ! t.settings.wheelPropagation
				}
				var o = e.scrollLeft;
				if (0 === i) {
					if (!t.scrollbarXActive) return ! 1;
					if (0 === o && 0 > n || o >= t.contentWidth - t.containerWidth && n > 0) return ! t.settings.wheelPropagation
				}
				return ! 0
			}
			var r = !1;
			t.event.bind(e, "mouseenter",
			function() {
				r = !0
			}),
			t.event.bind(e, "mouseleave",
			function() {
				r = !1
			});
			var a = !1;
			t.event.bind(t.ownerDocument, "keydown",
			function(s) {
				if ((!s.isDefaultPrevented || !s.isDefaultPrevented()) && r) {
					var l = document.activeElement ? document.activeElement: t.ownerDocument.activeElement;
					if (l) {
						for (; l.shadowRoot;) l = l.shadowRoot.activeElement;
						if (i.isEditable(l)) return
					}
					var c = 0,
					u = 0;
					switch (s.which) {
					case 37:
						c = -30;
						break;
					case 38:
						u = 30;
						break;
					case 39:
						c = 30;
						break;
					case 40:
						u = -30;
						break;
					case 33:
						u = 90;
						break;
					case 32:
					case 34:
						u = -90;
						break;
					case 35:
						u = s.ctrlKey ? -t.contentHeight: -t.containerHeight;
						break;
					case 36:
						u = s.ctrlKey ? e.scrollTop: t.containerHeight;
						break;
					default:
						return
					}
					e.scrollTop = e.scrollTop - u,
					e.scrollLeft = e.scrollLeft + c,
					o(e),
					a = n(c, u),
					a && s.preventDefault()
				}
			})
		}
		var i = e("../../lib/helper"),
		r = e("../instances"),
		o = e("../update-geometry");
		t.exports = function(e) {
			var t = r.get(e);
			n(e, t)
		}
	},
	{
		"../../lib/helper": 6,
		"../instances": 18,
		"../update-geometry": 19
	}],
	13 : [function(e, t) {
		"use strict";
		function n(e, t) {
			function n(n, i) {
				var r = e.scrollTop;
				if (0 === n) {
					if (!t.scrollbarYActive) return ! 1;
					if (0 === r && i > 0 || r >= t.contentHeight - t.containerHeight && 0 > i) return ! t.settings.wheelPropagation
				}
				var o = e.scrollLeft;
				if (0 === i) {
					if (!t.scrollbarXActive) return ! 1;
					if (0 === o && 0 > n || o >= t.contentWidth - t.containerWidth && n > 0) return ! t.settings.wheelPropagation
				}
				return ! 0
			}
			function r(e) {
				var t = e.deltaX,
				n = -1 * e.deltaY;
				return ("undefined" == typeof t || "undefined" == typeof n) && (t = -1 * e.wheelDeltaX / 6, n = e.wheelDeltaY / 6),
				e.deltaMode && 1 === e.deltaMode && (t *= 10, n *= 10),
				t !== t && n !== n && (t = 0, n = e.wheelDelta),
				[t, n]
			}
			function a(t, n) {
				var i = e.querySelector("textarea:hover");
				if (i) {
					var r = i.scrollHeight - i.clientHeight;
					if (r > 0 && !(0 === i.scrollTop && n > 0 || i.scrollTop === r && 0 > n)) return ! 0;
					var o = i.scrollLeft - i.clientWidth;
					if (o > 0 && !(0 === i.scrollLeft && 0 > t || i.scrollLeft === o && t > 0)) return ! 0
				}
				return ! 1
			}
			function s(s) {
				if (i.env.isWebKit || !e.querySelector("select:focus")) {
					var c = r(s),
					u = c[0],
					d = c[1];
					a(u, d) || (l = !1, t.settings.useBothWheelAxes ? t.scrollbarYActive && !t.scrollbarXActive ? (e.scrollTop = d ? e.scrollTop - d * t.settings.wheelSpeed: e.scrollTop + u * t.settings.wheelSpeed, l = !0) : t.scrollbarXActive && !t.scrollbarYActive && (e.scrollLeft = u ? e.scrollLeft + u * t.settings.wheelSpeed: e.scrollLeft - d * t.settings.wheelSpeed, l = !0) : (e.scrollTop = e.scrollTop - d * t.settings.wheelSpeed, e.scrollLeft = e.scrollLeft + u * t.settings.wheelSpeed), o(e), l = l || n(u, d), l && (s.stopPropagation(), s.preventDefault()))
				}
			}
			var l = !1;
			"undefined" != typeof window.onwheel ? t.event.bind(e, "wheel", s) : "undefined" != typeof window.onmousewheel && t.event.bind(e, "mousewheel", s)
		}
		var i = e("../../lib/helper"),
		r = e("../instances"),
		o = e("../update-geometry");
		t.exports = function(e) {
			var t = r.get(e);
			n(e, t)
		}
	},
	{
		"../../lib/helper": 6,
		"../instances": 18,
		"../update-geometry": 19
	}],
	14 : [function(e, t) {
		"use strict";
		function n(e, t) {
			t.event.bind(e, "scroll",
			function() {
				r(e)
			})
		}
		var i = e("../instances"),
		r = e("../update-geometry");
		t.exports = function(e) {
			var t = i.get(e);
			n(e, t)
		}
	},
	{
		"../instances": 18,
		"../update-geometry": 19
	}],
	15 : [function(e, t) {
		"use strict";
		function n(e, t) {
			function n() {
				var e = window.getSelection ? window.getSelection() : document.getSelection ? document.getSelection() : "";
				return 0 === e.toString().length ? null: e.getRangeAt(0).commonAncestorContainer
			}
			function a() {
				l || (l = setInterval(function() {
					return r.get(e) ? (e.scrollTop = e.scrollTop + c.top, e.scrollLeft = e.scrollLeft + c.left, void o(e)) : void clearInterval(l)
				},
				50))
			}
			function s() {
				l && (clearInterval(l), l = null),
				i.stopScrolling(e)
			}
			var l = null,
			c = {
				top: 0,
				left: 0
			},
			u = !1;
			t.event.bind(t.ownerDocument, "selectionchange",
			function() {
				e.contains(n()) ? u = !0 : (u = !1, s())
			}),
			t.event.bind(window, "mouseup",
			function() {
				u && (u = !1, s())
			}),
			t.event.bind(window, "mousemove",
			function(t) {
				if (u) {
					var n = {
						x: t.pageX,
						y: t.pageY
					},
					r = {
						left: e.offsetLeft,
						right: e.offsetLeft + e.offsetWidth,
						top: e.offsetTop,
						bottom: e.offsetTop + e.offsetHeight
					};
					n.x < r.left + 3 ? (c.left = -5, i.startScrolling(e, "x")) : n.x > r.right - 3 ? (c.left = 5, i.startScrolling(e, "x")) : c.left = 0,
					n.y < r.top + 3 ? (c.top = r.top + 3 - n.y < 5 ? -5 : -20, i.startScrolling(e, "y")) : n.y > r.bottom - 3 ? (c.top = n.y - r.bottom + 3 < 5 ? 5 : 20, i.startScrolling(e, "y")) : c.top = 0,
					0 === c.top && 0 === c.left ? s() : a()
				}
			})
		}
		var i = e("../../lib/helper"),
		r = e("../instances"),
		o = e("../update-geometry");
		t.exports = function(e) {
			var t = r.get(e);
			n(e, t)
		}
	},
	{
		"../../lib/helper": 6,
		"../instances": 18,
		"../update-geometry": 19
	}],
	16 : [function(e, t) {
		"use strict";
		function n(e, t, n, o) {
			function a(n, i) {
				var r = e.scrollTop,
				o = e.scrollLeft,
				a = Math.abs(n),
				s = Math.abs(i);
				if (s > a) {
					if (0 > i && r === t.contentHeight - t.containerHeight || i > 0 && 0 === r) return ! t.settings.swipePropagation
				} else if (a > s && (0 > n && o === t.contentWidth - t.containerWidth || n > 0 && 0 === o)) return ! t.settings.swipePropagation;
				return ! 0
			}
			function s(t, n) {
				e.scrollTop = e.scrollTop - n,
				e.scrollLeft = e.scrollLeft - t,
				r(e)
			}
			function l() {
				b = !0
			}
			function c() {
				b = !1
			}
			function u(e) {
				return e.targetTouches ? e.targetTouches[0] : e
			}
			function d(e) {
				return e.targetTouches && 1 === e.targetTouches.length ? !0 : e.pointerType && "mouse" !== e.pointerType && e.pointerType !== e.MSPOINTER_TYPE_MOUSE ? !0 : !1
			}
			function f(e) {
				if (d(e)) {
					w = !0;
					var t = u(e);
					g.pageX = t.pageX,
					g.pageY = t.pageY,
					m = (new Date).getTime(),
					null !== y && clearInterval(y),
					e.stopPropagation()
				}
			}
			function p(e) {
				if (!b && w && d(e)) {
					var t = u(e),
					n = {
						pageX: t.pageX,
						pageY: t.pageY
					},
					i = n.pageX - g.pageX,
					r = n.pageY - g.pageY;
					s(i, r),
					g = n;
					var o = (new Date).getTime(),
					l = o - m;
					l > 0 && (v.x = i / l, v.y = r / l, m = o),
					a(i, r) && (e.stopPropagation(), e.preventDefault())
				}
			}
			function h() { ! b && w && (w = !1, clearInterval(y), y = setInterval(function() {
					return i.get(e) ? Math.abs(v.x) < .01 && Math.abs(v.y) < .01 ? void clearInterval(y) : (s(30 * v.x, 30 * v.y), v.x *= .8, void(v.y *= .8)) : void clearInterval(y)
				},
				10))
			}
			var g = {},
			m = 0,
			v = {},
			y = null,
			b = !1,
			w = !1;
			n && (t.event.bind(window, "touchstart", l), t.event.bind(window, "touchend", c), t.event.bind(e, "touchstart", f), t.event.bind(e, "touchmove", p), t.event.bind(e, "touchend", h)),
			o && (window.PointerEvent ? (t.event.bind(window, "pointerdown", l), t.event.bind(window, "pointerup", c), t.event.bind(e, "pointerdown", f), t.event.bind(e, "pointermove", p), t.event.bind(e, "pointerup", h)) : window.MSPointerEvent && (t.event.bind(window, "MSPointerDown", l), t.event.bind(window, "MSPointerUp", c), t.event.bind(e, "MSPointerDown", f), t.event.bind(e, "MSPointerMove", p), t.event.bind(e, "MSPointerUp", h)))
		}
		var i = e("../instances"),
		r = e("../update-geometry");
		t.exports = function(e, t, r) {
			var o = i.get(e);
			n(e, o, t, r)
		}
	},
	{
		"../instances": 18,
		"../update-geometry": 19
	}],
	17 : [function(e, t) {
		"use strict";
		var n = e("../lib/class"),
		i = e("../lib/helper"),
		r = e("./instances"),
		o = e("./update-geometry"),
		a = e("./handler/click-rail"),
		s = e("./handler/drag-scrollbar"),
		l = e("./handler/keyboard"),
		c = e("./handler/mouse-wheel"),
		u = e("./handler/native-scroll"),
		d = e("./handler/selection"),
		f = e("./handler/touch");
		t.exports = function(e, t) {
			t = "object" == typeof t ? t: {},
			n.add(e, "ps-container");
			var p = r.add(e);
			p.settings = i.extend(p.settings, t),
			a(e),
			s(e),
			c(e),
			u(e),
			d(e),
			(i.env.supportsTouch || i.env.supportsIePointer) && f(e, i.env.supportsTouch, i.env.supportsIePointer),
			p.settings.useKeyboard && l(e),
			o(e)
		}
	},
	{
		"../lib/class": 2,
		"../lib/helper": 6,
		"./handler/click-rail": 10,
		"./handler/drag-scrollbar": 11,
		"./handler/keyboard": 12,
		"./handler/mouse-wheel": 13,
		"./handler/native-scroll": 14,
		"./handler/selection": 15,
		"./handler/touch": 16,
		"./instances": 18,
		"./update-geometry": 19
	}],
	18 : [function(e, t, n) {
		"use strict";
		function i(e) {
			var t = this;
			t.settings = d.clone(l),
			t.containerWidth = null,
			t.containerHeight = null,
			t.contentWidth = null,
			t.contentHeight = null,
			t.isRtl = "rtl" === s.css(e, "direction"),
			t.event = new c,
			t.ownerDocument = e.ownerDocument || document,
			t.scrollbarXRail = s.appendTo(s.e("div", "ps-scrollbar-x-rail"), e),
			t.scrollbarX = s.appendTo(s.e("div", "ps-scrollbar-x"), t.scrollbarXRail),
			t.scrollbarXActive = null,
			t.scrollbarXWidth = null,
			t.scrollbarXLeft = null,
			t.scrollbarXBottom = d.toInt(s.css(t.scrollbarXRail, "bottom")),
			t.isScrollbarXUsingBottom = t.scrollbarXBottom === t.scrollbarXBottom,
			t.scrollbarXTop = t.isScrollbarXUsingBottom ? null: d.toInt(s.css(t.scrollbarXRail, "top")),
			t.railBorderXWidth = d.toInt(s.css(t.scrollbarXRail, "borderLeftWidth")) + d.toInt(s.css(t.scrollbarXRail, "borderRightWidth")),
			t.railXMarginWidth = d.toInt(s.css(t.scrollbarXRail, "marginLeft")) + d.toInt(s.css(t.scrollbarXRail, "marginRight")),
			t.railXWidth = null,
			t.scrollbarYRail = s.appendTo(s.e("div", "ps-scrollbar-y-rail"), e),
			t.scrollbarY = s.appendTo(s.e("div", "ps-scrollbar-y"), t.scrollbarYRail),
			t.scrollbarYActive = null,
			t.scrollbarYHeight = null,
			t.scrollbarYTop = null,
			t.scrollbarYRight = d.toInt(s.css(t.scrollbarYRail, "right")),
			t.isScrollbarYUsingRight = t.scrollbarYRight === t.scrollbarYRight,
			t.scrollbarYLeft = t.isScrollbarYUsingRight ? null: d.toInt(s.css(t.scrollbarYRail, "left")),
			t.scrollbarYOuterWidth = t.isRtl ? d.outerWidth(t.scrollbarY) : null,
			t.railBorderYWidth = d.toInt(s.css(t.scrollbarYRail, "borderTopWidth")) + d.toInt(s.css(t.scrollbarYRail, "borderBottomWidth")),
			t.railYMarginHeight = d.toInt(s.css(t.scrollbarYRail, "marginTop")) + d.toInt(s.css(t.scrollbarYRail, "marginBottom")),
			t.railYHeight = null
		}
		function r(e) {
			return "undefined" == typeof e.dataset ? e.getAttribute("data-ps-id") : e.dataset.psId
		}
		function o(e, t) {
			"undefined" == typeof e.dataset ? e.setAttribute("data-ps-id", t) : e.dataset.psId = t
		}
		function a(e) {
			"undefined" == typeof e.dataset ? e.removeAttribute("data-ps-id") : delete e.dataset.psId
		}
		var s = e("../lib/dom"),
		l = e("./default-setting"),
		c = e("../lib/event-manager"),
		u = e("../lib/guid"),
		d = e("../lib/helper"),
		f = {};
		n.add = function(e) {
			var t = u();
			return o(e, t),
			f[t] = new i(e),
			f[t]
		},
		n.remove = function(e) {
			delete f[r(e)],
			a(e)
		},
		n.get = function(e) {
			return f[r(e)]
		}
	},
	{
		"../lib/dom": 3,
		"../lib/event-manager": 4,
		"../lib/guid": 5,
		"../lib/helper": 6,
		"./default-setting": 8
	}],
	19 : [function(e, t) {
		"use strict";
		function n(e, t) {
			return e.settings.minScrollbarLength && (t = Math.max(t, e.settings.minScrollbarLength)),
			e.settings.maxScrollbarLength && (t = Math.min(t, e.settings.maxScrollbarLength)),
			t
		}
		function i(e, t) {
			var n = {
				width: t.railXWidth
			};
			n.left = t.isRtl ? e.scrollLeft + t.containerWidth - t.contentWidth: e.scrollLeft,
			t.isScrollbarXUsingBottom ? n.bottom = t.scrollbarXBottom - e.scrollTop: n.top = t.scrollbarXTop + e.scrollTop,
			o.css(t.scrollbarXRail, n);
			var i = {
				top: e.scrollTop,
				height: t.railYHeight
			};
			t.isScrollbarYUsingRight ? i.right = t.isRtl ? t.contentWidth - e.scrollLeft - t.scrollbarYRight - t.scrollbarYOuterWidth: t.scrollbarYRight - e.scrollLeft: i.left = t.isRtl ? e.scrollLeft + 2 * t.containerWidth - t.contentWidth - t.scrollbarYLeft - t.scrollbarYOuterWidth: t.scrollbarYLeft + e.scrollLeft,
			o.css(t.scrollbarYRail, i),
			o.css(t.scrollbarX, {
				left: t.scrollbarXLeft,
				width: t.scrollbarXWidth - t.railBorderXWidth
			}),
			o.css(t.scrollbarY, {
				top: t.scrollbarYTop,
				height: t.scrollbarYHeight - t.railBorderYWidth
			})
		}
		var r = e("../lib/class"),
		o = e("../lib/dom"),
		a = e("../lib/helper"),
		s = e("./instances");
		t.exports = function(e) {
			var t = s.get(e);
			t.containerWidth = e.clientWidth,
			t.containerHeight = e.clientHeight,
			t.contentWidth = e.scrollWidth,
			t.contentHeight = e.scrollHeight,
			e.contains(t.scrollbarXRail) || o.appendTo(t.scrollbarXRail, e),
			e.contains(t.scrollbarYRail) || o.appendTo(t.scrollbarYRail, e),
			!t.settings.suppressScrollX && t.containerWidth + t.settings.scrollXMarginOffset < t.contentWidth ? (t.scrollbarXActive = !0, t.railXWidth = t.containerWidth - t.railXMarginWidth, t.scrollbarXWidth = n(t, a.toInt(t.railXWidth * t.containerWidth / t.contentWidth)), t.scrollbarXLeft = a.toInt(e.scrollLeft * (t.railXWidth - t.scrollbarXWidth) / (t.contentWidth - t.containerWidth))) : (t.scrollbarXActive = !1, t.scrollbarXWidth = 0, t.scrollbarXLeft = 0, e.scrollLeft = 0),
			!t.settings.suppressScrollY && t.containerHeight + t.settings.scrollYMarginOffset < t.contentHeight ? (t.scrollbarYActive = !0, t.railYHeight = t.containerHeight - t.railYMarginHeight, t.scrollbarYHeight = n(t, a.toInt(t.railYHeight * t.containerHeight / t.contentHeight)), t.scrollbarYTop = a.toInt(e.scrollTop * (t.railYHeight - t.scrollbarYHeight) / (t.contentHeight - t.containerHeight))) : (t.scrollbarYActive = !1, t.scrollbarYHeight = 0, t.scrollbarYTop = 0, e.scrollTop = 0),
			t.scrollbarXLeft >= t.railXWidth - t.scrollbarXWidth && (t.scrollbarXLeft = t.railXWidth - t.scrollbarXWidth),
			t.scrollbarYTop >= t.railYHeight - t.scrollbarYHeight && (t.scrollbarYTop = t.railYHeight - t.scrollbarYHeight),
			i(e, t),
			r[t.scrollbarXActive ? "add": "remove"](e, "ps-active-x"),
			r[t.scrollbarYActive ? "add": "remove"](e, "ps-active-y")
		}
	},
	{
		"../lib/class": 2,
		"../lib/dom": 3,
		"../lib/helper": 6,
		"./instances": 18
	}],
	20 : [function(e, t) {
		"use strict";
		var n = e("../lib/dom"),
		i = e("./instances"),
		r = e("./update-geometry");
		t.exports = function(e) {
			var t = i.get(e);
			n.css(t.scrollbarXRail, "display", "none"),
			n.css(t.scrollbarYRail, "display", "none"),
			r(e),
			n.css(t.scrollbarXRail, "display", "block"),
			n.css(t.scrollbarYRail, "display", "block")
		}
	},
	{
		"../lib/dom": 3,
		"./instances": 18,
		"./update-geometry": 19
	}]
},
{},
[1]),
function(e) {
	"function" == typeof define && define.amd ? define(["jquery"], e) : e("object" == typeof exports ? require("jquery") : window.jQuery || window.Zepto)
} (function(e) {
	var t, n, i, r, o, a, s, l = "Close",
	c = "BeforeClose",
	u = "AfterClose",
	d = "BeforeAppend",
	f = "MarkupParse",
	p = "Open",
	h = "Change",
	g = "mfp",
	m = "." + g,
	v = "mfp-ready",
	y = "mfp-removing",
	b = "mfp-prevent-close",
	w = function() {},
	x = !!window.jQuery,
	k = e(window),
	C = function(e, n) {
		t.ev.on(g + e + m, n)
	},
	T = function(t, n, i, r) {
		var o = document.createElement("div");
		return o.className = "mfp-" + t,
		i && (o.innerHTML = i),
		r ? n && n.appendChild(o) : (o = e(o), n && o.appendTo(n)),
		o
	},
	S = function(n, i) {
		t.ev.triggerHandler(g + n, i),
		t.st.callbacks && (n = n.charAt(0).toLowerCase() + n.slice(1), t.st.callbacks[n] && t.st.callbacks[n].apply(t, e.isArray(i) ? i: [i]))
	},
	_ = function(n) {
		return n === s && t.currTemplate.closeBtn || (t.currTemplate.closeBtn = e(t.st.closeMarkup.replace("%title%", t.st.tClose)), s = n),
		t.currTemplate.closeBtn
	},
	$ = function() {
		e.magnificPopup.instance || (t = new w, t.init(), e.magnificPopup.instance = t)
	},
	L = function() {
		var e = document.createElement("p").style,
		t = ["ms", "O", "Moz", "Webkit"];
		if (void 0 !== e.transition) return ! 0;
		for (; t.length;) if (t.pop() + "Transition" in e) return ! 0;
		return ! 1
	};
	w.prototype = {
		constructor: w,
		init: function() {
			var n = navigator.appVersion;
			t.isIE7 = -1 !== n.indexOf("MSIE 7."),
			t.isIE8 = -1 !== n.indexOf("MSIE 8."),
			t.isLowIE = t.isIE7 || t.isIE8,
			t.isAndroid = /android/gi.test(n),
			t.isIOS = /iphone|ipad|ipod/gi.test(n),
			t.supportsTransition = L(),
			t.probablyMobile = t.isAndroid || t.isIOS || /(Opera Mini)|Kindle|webOS|BlackBerry|(Opera Mobi)|(Windows Phone)|IEMobile/i.test(navigator.userAgent),
			r = e(document),
			t.popupsCache = {}
		},
		open: function(n) {
			i || (i = e(document.body));
			var o;
			if (n.isObj === !1) {
				t.items = n.items.toArray(),
				t.index = 0;
				var s, l = n.items;
				for (o = 0; l.length > o; o++) if (s = l[o], s.parsed && (s = s.el[0]), s === n.el[0]) {
					t.index = o;
					break
				}
			} else t.items = e.isArray(n.items) ? n.items: [n.items],
			t.index = n.index || 0;
			if (t.isOpen) return void t.updateItemHTML();
			t.types = [],
			a = "",
			t.ev = n.mainEl && n.mainEl.length ? n.mainEl.eq(0) : r,
			n.key ? (t.popupsCache[n.key] || (t.popupsCache[n.key] = {}), t.currTemplate = t.popupsCache[n.key]) : t.currTemplate = {},
			t.st = e.extend(!0, {},
			e.magnificPopup.defaults, n),
			t.fixedContentPos = "auto" === t.st.fixedContentPos ? !t.probablyMobile: t.st.fixedContentPos,
			t.st.modal && (t.st.closeOnContentClick = !1, t.st.closeOnBgClick = !1, t.st.showCloseBtn = !1, t.st.enableEscapeKey = !1),
			t.bgOverlay || (t.bgOverlay = T("bg").on("click" + m,
			function() {
				t.close()
			}), t.wrap = T("wrap").attr("tabindex", -1).on("click" + m,
			function(e) {
				t._checkIfClose(e.target) && t.close()
			}), t.container = T("container", t.wrap)),
			t.contentContainer = T("content"),
			t.st.preloader && (t.preloader = T("preloader", t.container, t.st.tLoading));
			var c = e.magnificPopup.modules;
			for (o = 0; c.length > o; o++) {
				var u = c[o];
				u = u.charAt(0).toUpperCase() + u.slice(1),
				t["init" + u].call(t)
			}
			S("BeforeOpen"),
			t.st.showCloseBtn && (t.st.closeBtnInside ? (C(f,
			function(e, t, n, i) {
				n.close_replaceWith = _(i.type)
			}), a += " mfp-close-btn-in") : t.wrap.append(_())),
			t.st.alignTop && (a += " mfp-align-top"),
			t.wrap.css(t.fixedContentPos ? {
				overflow: t.st.overflowY,
				overflowX: "hidden",
				overflowY: t.st.overflowY
			}: {
				top: k.scrollTop(),
				position: "absolute"
			}),
			(t.st.fixedBgPos === !1 || "auto" === t.st.fixedBgPos && !t.fixedContentPos) && t.bgOverlay.css({
				height: r.height(),
				position: "absolute"
			}),
			t.st.enableEscapeKey && r.on("keyup" + m,
			function(e) {
				27 === e.keyCode && t.close()
			}),
			k.on("resize" + m,
			function() {
				t.updateSize()
			}),
			t.st.closeOnContentClick || (a += " mfp-auto-cursor"),
			a && t.wrap.addClass(a);
			var d = t.wH = k.height(),
			h = {};
			if (t.fixedContentPos && t._hasScrollBar(d)) {
				var g = t._getScrollbarSize();
				g && (h.marginRight = g)
			}
			t.fixedContentPos && (t.isIE7 ? e("body, html").css("overflow", "hidden") : h.overflow = "hidden");
			var y = t.st.mainClass;
			return t.isIE7 && (y += " mfp-ie7"),
			y && t._addClassToMFP(y),
			t.updateItemHTML(),
			S("BuildControls"),
			e("html").css(h),
			t.bgOverlay.add(t.wrap).prependTo(t.st.prependTo || i),
			t._lastFocusedEl = document.activeElement,
			setTimeout(function() {
				t.content ? (t._addClassToMFP(v), t._setFocus()) : t.bgOverlay.addClass(v),
				r.on("focusin" + m, t._onFocusIn)
			},
			16),
			t.isOpen = !0,
			t.updateSize(d),
			S(p),
			n
		},
		close: function() {
			t.isOpen && (S(c), t.isOpen = !1, t.st.removalDelay && !t.isLowIE && t.supportsTransition ? (t._addClassToMFP(y), setTimeout(function() {
				t._close()
			},
			t.st.removalDelay)) : t._close())
		},
		_close: function() {
			S(l);
			var n = y + " " + v + " ";
			if (t.bgOverlay.detach(), t.wrap.detach(), t.container.empty(), t.st.mainClass && (n += t.st.mainClass + " "), t._removeClassFromMFP(n), t.fixedContentPos) {
				var i = {
					marginRight: ""
				};
				t.isIE7 ? e("body, html").css("overflow", "") : i.overflow = "",
				e("html").css(i)
			}
			r.off("keyup" + m + " focusin" + m),
			t.ev.off(m),
			t.wrap.attr("class", "mfp-wrap").removeAttr("style"),
			t.bgOverlay.attr("class", "mfp-bg"),
			t.container.attr("class", "mfp-container"),
			!t.st.showCloseBtn || t.st.closeBtnInside && t.currTemplate[t.currItem.type] !== !0 || t.currTemplate.closeBtn && t.currTemplate.closeBtn.detach(),
			t._lastFocusedEl && e(t._lastFocusedEl).focus(),
			t.currItem = null,
			t.content = null,
			t.currTemplate = null,
			t.prevHeight = 0,
			S(u)
		},
		updateSize: function(e) {
			if (t.isIOS) {
				var n = document.documentElement.clientWidth / window.innerWidth,
				i = window.innerHeight * n;
				t.wrap.css("height", i),
				t.wH = i
			} else t.wH = e || k.height();
			t.fixedContentPos || t.wrap.css("height", t.wH),
			S("Resize")
		},
		updateItemHTML: function() {
			var n = t.items[t.index];
			t.contentContainer.detach(),
			t.content && t.content.detach(),
			n.parsed || (n = t.parseEl(t.index));
			var i = n.type;
			if (S("BeforeChange", [t.currItem ? t.currItem.type: "", i]), t.currItem = n, !t.currTemplate[i]) {
				var r = t.st[i] ? t.st[i].markup: !1;
				S("FirstMarkupParse", r),
				t.currTemplate[i] = r ? e(r) : !0
			}
			o && o !== n.type && t.container.removeClass("mfp-" + o + "-holder");
			var a = t["get" + i.charAt(0).toUpperCase() + i.slice(1)](n, t.currTemplate[i]);
			t.appendContent(a, i),
			n.preloaded = !0,
			S(h, n),
			o = n.type,
			t.container.prepend(t.contentContainer),
			S("AfterChange")
		},
		appendContent: function(e, n) {
			t.content = e,
			e ? t.st.showCloseBtn && t.st.closeBtnInside && t.currTemplate[n] === !0 ? t.content.find(".mfp-close").length || t.content.append(_()) : t.content = e: t.content = "",
			S(d),
			t.container.addClass("mfp-" + n + "-holder"),
			t.contentContainer.append(t.content)
		},
		parseEl: function(n) {
			var i, r = t.items[n];
			if (r.tagName ? r = {
				el: e(r)
			}: (i = r.type, r = {
				data: r,
				src: r.src
			}), r.el) {
				for (var o = t.types,
				a = 0; o.length > a; a++) if (r.el.hasClass("mfp-" + o[a])) {
					i = o[a];
					break
				}
				r.src = r.el.attr("data-mfp-src"),
				r.src || (r.src = r.el.attr("href"))
			}
			return r.type = i || t.st.type || "inline",
			r.index = n,
			r.parsed = !0,
			t.items[n] = r,
			S("ElementParse", r),
			t.items[n]
		},
		addGroup: function(e, n) {
			var i = function(i) {
				i.mfpEl = this,
				t._openClick(i, e, n)
			};
			n || (n = {});
			var r = "click.magnificPopup";
			n.mainEl = e,
			n.items ? (n.isObj = !0, e.off(r).on(r, i)) : (n.isObj = !1, n.delegate ? e.off(r).on(r, n.delegate, i) : (n.items = e, e.off(r).on(r, i)))
		},
		_openClick: function(n, i, r) {
			var o = void 0 !== r.midClick ? r.midClick: e.magnificPopup.defaults.midClick;
			if (o || 2 !== n.which && !n.ctrlKey && !n.metaKey) {
				var a = void 0 !== r.disableOn ? r.disableOn: e.magnificPopup.defaults.disableOn;
				if (a) if (e.isFunction(a)) {
					if (!a.call(t)) return ! 0
				} else if (a > k.width()) return ! 0;
				n.type && (n.preventDefault(), t.isOpen && n.stopPropagation()),
				r.el = e(n.mfpEl),
				r.delegate && (r.items = i.find(r.delegate)),
				t.open(r)
			}
		},
		updateStatus: function(e, i) {
			if (t.preloader) {
				n !== e && t.container.removeClass("mfp-s-" + n),
				i || "loading" !== e || (i = t.st.tLoading);
				var r = {
					status: e,
					text: i
				};
				S("UpdateStatus", r),
				e = r.status,
				i = r.text,
				t.preloader.html(i),
				t.preloader.find("a").on("click",
				function(e) {
					e.stopImmediatePropagation()
				}),
				t.container.addClass("mfp-s-" + e),
				n = e
			}
		},
		_checkIfClose: function(n) {
			if (!e(n).hasClass(b)) {
				var i = t.st.closeOnContentClick,
				r = t.st.closeOnBgClick;
				if (i && r) return ! 0;
				if (!t.content || e(n).hasClass("mfp-close") || t.preloader && n === t.preloader[0]) return ! 0;
				if (n === t.content[0] || e.contains(t.content[0], n)) {
					if (i) return ! 0
				} else if (r && e.contains(document, n)) return ! 0;
				return ! 1
			}
		},
		_addClassToMFP: function(e) {
			t.bgOverlay.addClass(e),
			t.wrap.addClass(e)
		},
		_removeClassFromMFP: function(e) {
			this.bgOverlay.removeClass(e),
			t.wrap.removeClass(e)
		},
		_hasScrollBar: function(e) {
			return (t.isIE7 ? r.height() : document.body.scrollHeight) > (e || k.height())
		},
		_setFocus: function() { (t.st.focus ? t.content.find(t.st.focus).eq(0) : t.wrap).focus()
		},
		_onFocusIn: function(n) {
			return n.target === t.wrap[0] || e.contains(t.wrap[0], n.target) ? void 0 : (t._setFocus(), !1)
		},
		_parseMarkup: function(t, n, i) {
			var r;
			i.data && (n = e.extend(i.data, n)),
			S(f, [t, n, i]),
			e.each(n,
			function(e, n) {
				if (void 0 === n || n === !1) return ! 0;
				if (r = e.split("_"), r.length > 1) {
					var i = t.find(m + "-" + r[0]);
					if (i.length > 0) {
						var o = r[1];
						"replaceWith" === o ? i[0] !== n[0] && i.replaceWith(n) : "img" === o ? i.is("img") ? i.attr("src", n) : i.replaceWith('<img src="' + n + '" class="' + i.attr("class") + '" />') : i.attr(r[1], n)
					}
				} else t.find(m + "-" + e).html(n)
			})
		},
		_getScrollbarSize: function() {
			if (void 0 === t.scrollbarSize) {
				var e = document.createElement("div");
				e.style.cssText = "width: 99px; height: 99px; overflow: scroll; position: absolute; top: -9999px;",
				document.body.appendChild(e),
				t.scrollbarSize = e.offsetWidth - e.clientWidth,
				document.body.removeChild(e)
			}
			return t.scrollbarSize
		}
	},
	e.magnificPopup = {
		instance: null,
		proto: w.prototype,
		modules: [],
		open: function(t, n) {
			return $(),
			t = t ? e.extend(!0, {},
			t) : {},
			t.isObj = !0,
			t.index = n || 0,
			this.instance.open(t)
		},
		close: function() {
			return e.magnificPopup.instance && e.magnificPopup.instance.close()
		},
		registerModule: function(t, n) {
			n.options && (e.magnificPopup.defaults[t] = n.options),
			e.extend(this.proto, n.proto),
			this.modules.push(t)
		},
		defaults: {
			disableOn: 0,
			key: null,
			midClick: !1,
			mainClass: "",
			preloader: !0,
			focus: "",
			closeOnContentClick: !1,
			closeOnBgClick: !0,
			closeBtnInside: !0,
			showCloseBtn: !0,
			enableEscapeKey: !0,
			modal: !1,
			alignTop: !1,
			removalDelay: 0,
			prependTo: null,
			fixedContentPos: "auto",
			fixedBgPos: "auto",
			overflowY: "auto",
			closeMarkup: '<button title="%title%" type="button" class="mfp-close">&times;</button>',
			tClose: "Close (Esc)",
			tLoading: "Loading..."
		}
	},
	e.fn.magnificPopup = function(n) {
		$();
		var i = e(this);
		if ("string" == typeof n) if ("open" === n) {
			var r, o = x ? i.data("magnificPopup") : i[0].magnificPopup,
			a = parseInt(arguments[1], 10) || 0;
			o.items ? r = o.items[a] : (r = i, o.delegate && (r = r.find(o.delegate)), r = r.eq(a)),
			t._openClick({
				mfpEl: r
			},
			i, o)
		} else t.isOpen && t[n].apply(t, Array.prototype.slice.call(arguments, 1));
		else n = e.extend(!0, {},
		n),
		x ? i.data("magnificPopup", n) : i[0].magnificPopup = n,
		t.addGroup(i, n);
		return i
	};
	var E, I, N, A = "inline",
	D = function() {
		N && (I.after(N.addClass(E)).detach(), N = null)
	};
	e.magnificPopup.registerModule(A, {
		options: {
			hiddenClass: "hide",
			markup: "",
			tNotFound: "Content not found"
		},
		proto: {
			initInline: function() {
				t.types.push(A),
				C(l + "." + A,
				function() {
					D()
				})
			},
			getInline: function(n, i) {
				if (D(), n.src) {
					var r = t.st.inline,
					o = e(n.src);
					if (o.length) {
						var a = o[0].parentNode;
						a && a.tagName && (I || (E = r.hiddenClass, I = T(E), E = "mfp-" + E), N = o.after(I).detach().removeClass(E)),
						t.updateStatus("ready")
					} else t.updateStatus("error", r.tNotFound),
					o = e("<div>");
					return n.inlineElement = o,
					o
				}
				return t.updateStatus("ready"),
				t._parseMarkup(i, {},
				n),
				i
			}
		}
	});
	var j, H = "ajax",
	M = function() {
		j && i.removeClass(j)
	},
	P = function() {
		M(),
		t.req && t.req.abort()
	};
	e.magnificPopup.registerModule(H, {
		options: {
			settings: null,
			cursor: "mfp-ajax-cur",
			tError: '<a href="%url%">The content</a> could not be loaded.'
		},
		proto: {
			initAjax: function() {
				t.types.push(H),
				j = t.st.ajax.cursor,
				C(l + "." + H, P),
				C("BeforeChange." + H, P)
			},
			getAjax: function(n) {
				j && i.addClass(j),
				t.updateStatus("loading");
				var r = e.extend({
					url: n.src,
					success: function(i, r, o) {
						var a = {
							data: i,
							xhr: o
						};
						S("ParseAjax", a),
						t.appendContent(e(a.data), H),
						n.finished = !0,
						M(),
						t._setFocus(),
						setTimeout(function() {
							t.wrap.addClass(v)
						},
						16),
						t.updateStatus("ready"),
						S("AjaxContentAdded")
					},
					error: function() {
						M(),
						n.finished = n.loadError = !0,
						t.updateStatus("error", t.st.ajax.tError.replace("%url%", n.src))
					}
				},
				t.st.ajax.settings);
				return t.req = e.ajax(r),
				""
			}
		}
	});
	var W, O = function(n) {
		if (n.data && void 0 !== n.data.title) return n.data.title;
		var i = t.st.image.titleSrc;
		if (i) {
			if (e.isFunction(i)) return i.call(t, n);
			if (n.el) return n.el.attr(i) || ""
		}
		return ""
	};
	e.magnificPopup.registerModule("image", {
		options: {
			markup: '<div class="mfp-figure"><div class="mfp-close"></div><figure><div class="mfp-img"></div><figcaption><div class="mfp-bottom-bar"><div class="mfp-title"></div><div class="mfp-counter"></div></div></figcaption></figure></div>',
			cursor: "mfp-zoom-out-cur",
			titleSrc: "title",
			verticalFit: !0,
			tError: '<a href="%url%">The image</a> could not be loaded.'
		},
		proto: {
			initImage: function() {
				var e = t.st.image,
				n = ".image";
				t.types.push("image"),
				C(p + n,
				function() {
					"image" === t.currItem.type && e.cursor && i.addClass(e.cursor)
				}),
				C(l + n,
				function() {
					e.cursor && i.removeClass(e.cursor),
					k.off("resize" + m)
				}),
				C("Resize" + n, t.resizeImage),
				t.isLowIE && C("AfterChange", t.resizeImage)
			},
			resizeImage: function() {
				var e = t.currItem;
				if (e && e.img && t.st.image.verticalFit) {
					var n = 0;
					t.isLowIE && (n = parseInt(e.img.css("padding-top"), 10) + parseInt(e.img.css("padding-bottom"), 10)),
					e.img.css("max-height", t.wH - n)
				}
			},
			_onImageHasSize: function(e) {
				e.img && (e.hasSize = !0, W && clearInterval(W), e.isCheckingImgSize = !1, S("ImageHasSize", e), e.imgHidden && (t.content && t.content.removeClass("mfp-loading"), e.imgHidden = !1))
			},
			findImageSize: function(e) {
				var n = 0,
				i = e.img[0],
				r = function(o) {
					W && clearInterval(W),
					W = setInterval(function() {
						return i.naturalWidth > 0 ? void t._onImageHasSize(e) : (n > 200 && clearInterval(W), n++, void(3 === n ? r(10) : 40 === n ? r(50) : 100 === n && r(500)))
					},
					o)
				};
				r(1)
			},
			getImage: function(n, i) {
				var r = 0,
				o = function() {
					n && (n.img[0].complete ? (n.img.off(".mfploader"), n === t.currItem && (t._onImageHasSize(n), t.updateStatus("ready")), n.hasSize = !0, n.loaded = !0, S("ImageLoadComplete")) : (r++, 200 > r ? setTimeout(o, 100) : a()))
				},
				a = function() {
					n && (n.img.off(".mfploader"), n === t.currItem && (t._onImageHasSize(n), t.updateStatus("error", s.tError.replace("%url%", n.src))), n.hasSize = !0, n.loaded = !0, n.loadError = !0)
				},
				s = t.st.image,
				l = i.find(".mfp-img");
				if (l.length) {
					var c = document.createElement("img");
					c.className = "mfp-img",
					n.el && n.el.find("img").length && (c.alt = n.el.find("img").attr("alt")),
					n.img = e(c).on("load.mfploader", o).on("error.mfploader", a),
					c.src = n.src,
					l.is("img") && (n.img = n.img.clone()),
					c = n.img[0],
					c.naturalWidth > 0 ? n.hasSize = !0 : c.width || (n.hasSize = !1)
				}
				return t._parseMarkup(i, {
					title: O(n),
					img_replaceWith: n.img
				},
				n),
				t.resizeImage(),
				n.hasSize ? (W && clearInterval(W), n.loadError ? (i.addClass("mfp-loading"), t.updateStatus("error", s.tError.replace("%url%", n.src))) : (i.removeClass("mfp-loading"), t.updateStatus("ready")), i) : (t.updateStatus("loading"), n.loading = !0, n.hasSize || (n.imgHidden = !0, i.addClass("mfp-loading"), t.findImageSize(n)), i)
			}
		}
	});
	var R, q = function() {
		return void 0 === R && (R = void 0 !== document.createElement("p").style.MozTransform),
		R
	};
	e.magnificPopup.registerModule("zoom", {
		options: {
			enabled: !1,
			easing: "ease-in-out",
			duration: 300,
			opener: function(e) {
				return e.is("img") ? e: e.find("img")
			}
		},
		proto: {
			initZoom: function() {
				var e, n = t.st.zoom,
				i = ".zoom";
				if (n.enabled && t.supportsTransition) {
					var r, o, a = n.duration,
					s = function(e) {
						var t = e.clone().removeAttr("style").removeAttr("class").addClass("mfp-animated-image"),
						i = "all " + n.duration / 1e3 + "s " + n.easing,
						r = {
							position: "fixed",
							zIndex: 9999,
							left: 0,
							top: 0,
							"-webkit-backface-visibility": "hidden"
						},
						o = "transition";
						return r["-webkit-" + o] = r["-moz-" + o] = r["-o-" + o] = r[o] = i,
						t.css(r),
						t
					},
					u = function() {
						t.content.css("visibility", "visible")
					};
					C("BuildControls" + i,
					function() {
						if (t._allowZoom()) {
							if (clearTimeout(r), t.content.css("visibility", "hidden"), e = t._getItemToZoom(), !e) return void u();
							o = s(e),
							o.css(t._getOffset()),
							t.wrap.append(o),
							r = setTimeout(function() {
								o.css(t._getOffset(!0)),
								r = setTimeout(function() {
									u(),
									setTimeout(function() {
										o.remove(),
										e = o = null,
										S("ZoomAnimationEnded")
									},
									16)
								},
								a)
							},
							16)
						}
					}),
					C(c + i,
					function() {
						if (t._allowZoom()) {
							if (clearTimeout(r), t.st.removalDelay = a, !e) {
								if (e = t._getItemToZoom(), !e) return;
								o = s(e)
							}
							o.css(t._getOffset(!0)),
							t.wrap.append(o),
							t.content.css("visibility", "hidden"),
							setTimeout(function() {
								o.css(t._getOffset())
							},
							16)
						}
					}),
					C(l + i,
					function() {
						t._allowZoom() && (u(), o && o.remove(), e = null)
					})
				}
			},
			_allowZoom: function() {
				return "image" === t.currItem.type
			},
			_getItemToZoom: function() {
				return t.currItem.hasSize ? t.currItem.img: !1
			},
			_getOffset: function(n) {
				var i;
				i = n ? t.currItem.img: t.st.zoom.opener(t.currItem.el || t.currItem);
				var r = i.offset(),
				o = parseInt(i.css("padding-top"), 10),
				a = parseInt(i.css("padding-bottom"), 10);
				r.top -= e(window).scrollTop() - o;
				var s = {
					width: i.width(),
					height: (x ? i.innerHeight() : i[0].offsetHeight) - a - o
				};
				return q() ? s["-moz-transform"] = s.transform = "translate(" + r.left + "px," + r.top + "px)": (s.left = r.left, s.top = r.top),
				s
			}
		}
	});
	var X = "iframe",
	F = "//about:blank",
	B = function(e) {
		if (t.currTemplate[X]) {
			var n = t.currTemplate[X].find("iframe");
			n.length && (e || (n[0].src = F), t.isIE8 && n.css("display", e ? "block": "none"))
		}
	};
	e.magnificPopup.registerModule(X, {
		options: {
			markup: '<div class="mfp-iframe-scaler"><div class="mfp-close"></div><iframe class="mfp-iframe" src="//about:blank" frameborder="0" allowfullscreen></iframe></div>',
			srcAction: "iframe_src",
			patterns: {
				youtube: {
					index: "youtube.com",
					id: "v=",
					src: "//www.youtube.com/embed/%id%?autoplay=1"
				},
				vimeo: {
					index: "vimeo.com/",
					id: "/",
					src: "//player.vimeo.com/video/%id%?autoplay=1"
				},
				gmaps: {
					index: "//maps.google.",
					src: "%id%&output=embed"
				}
			}
		},
		proto: {
			initIframe: function() {
				t.types.push(X),
				C("BeforeChange",
				function(e, t, n) {
					t !== n && (t === X ? B() : n === X && B(!0))
				}),
				C(l + "." + X,
				function() {
					B()
				})
			},
			getIframe: function(n, i) {
				var r = n.src,
				o = t.st.iframe;
				e.each(o.patterns,
				function() {
					return r.indexOf(this.index) > -1 ? (this.id && (r = "string" == typeof this.id ? r.substr(r.lastIndexOf(this.id) + this.id.length, r.length) : this.id.call(this, r)), r = this.src.replace("%id%", r), !1) : void 0
				});
				var a = {};
				return o.srcAction && (a[o.srcAction] = r),
				t._parseMarkup(i, a, n),
				t.updateStatus("ready"),
				i
			}
		}
	});
	var Y = function(e) {
		var n = t.items.length;
		return e > n - 1 ? e - n: 0 > e ? n + e: e
	},
	z = function(e, t, n) {
		return e.replace(/%curr%/gi, t + 1).replace(/%total%/gi, n)
	};
	e.magnificPopup.registerModule("gallery", {
		options: {
			enabled: !1,
			arrowMarkup: '<button title="%title%" type="button" class="mfp-arrow mfp-arrow-%dir%"></button>',
			preload: [0, 2],
			navigateByImgClick: !0,
			arrows: !0,
			tPrev: "Previous (Left arrow key)",
			tNext: "Next (Right arrow key)",
			tCounter: "%curr% of %total%"
		},
		proto: {
			initGallery: function() {
				var n = t.st.gallery,
				i = ".mfp-gallery",
				o = Boolean(e.fn.mfpFastClick);
				return t.direction = !0,
				n && n.enabled ? (a += " mfp-gallery", C(p + i,
				function() {
					n.navigateByImgClick && t.wrap.on("click" + i, ".mfp-img",
					function() {
						return t.items.length > 1 ? (t.next(), !1) : void 0
					}),
					r.on("keydown" + i,
					function(e) {
						37 === e.keyCode ? t.prev() : 39 === e.keyCode && t.next()
					})
				}), C("UpdateStatus" + i,
				function(e, n) {
					n.text && (n.text = z(n.text, t.currItem.index, t.items.length))
				}), C(f + i,
				function(e, i, r, o) {
					var a = t.items.length;
					r.counter = a > 1 ? z(n.tCounter, o.index, a) : ""
				}), C("BuildControls" + i,
				function() {
					if (t.items.length > 1 && n.arrows && !t.arrowLeft) {
						var i = n.arrowMarkup,
						r = t.arrowLeft = e(i.replace(/%title%/gi, n.tPrev).replace(/%dir%/gi, "left")).addClass(b),
						a = t.arrowRight = e(i.replace(/%title%/gi, n.tNext).replace(/%dir%/gi, "right")).addClass(b),
						s = o ? "mfpFastClick": "click";
						r[s](function() {
							t.prev()
						}),
						a[s](function() {
							t.next()
						}),
						t.isIE7 && (T("b", r[0], !1, !0), T("a", r[0], !1, !0), T("b", a[0], !1, !0), T("a", a[0], !1, !0)),
						t.container.append(r.add(a))
					}
				}), C(h + i,
				function() {
					t._preloadTimeout && clearTimeout(t._preloadTimeout),
					t._preloadTimeout = setTimeout(function() {
						t.preloadNearbyImages(),
						t._preloadTimeout = null
					},
					16)
				}), void C(l + i,
				function() {
					r.off(i),
					t.wrap.off("click" + i),
					t.arrowLeft && o && t.arrowLeft.add(t.arrowRight).destroyMfpFastClick(),
					t.arrowRight = t.arrowLeft = null
				})) : !1
			},
			next: function() {
				t.direction = !0,
				t.index = Y(t.index + 1),
				t.updateItemHTML()
			},
			prev: function() {
				t.direction = !1,
				t.index = Y(t.index - 1),
				t.updateItemHTML()
			},
			goTo: function(e) {
				t.direction = e >= t.index,
				t.index = e,
				t.updateItemHTML()
			},
			preloadNearbyImages: function() {
				var e, n = t.st.gallery.preload,
				i = Math.min(n[0], t.items.length),
				r = Math.min(n[1], t.items.length);
				for (e = 1; (t.direction ? r: i) >= e; e++) t._preloadItem(t.index + e);
				for (e = 1; (t.direction ? i: r) >= e; e++) t._preloadItem(t.index - e)
			},
			_preloadItem: function(n) {
				if (n = Y(n), !t.items[n].preloaded) {
					var i = t.items[n];
					i.parsed || (i = t.parseEl(n)),
					S("LazyLoad", i),
					"image" === i.type && (i.img = e('<img class="mfp-img" />').on("load.mfploader",
					function() {
						i.hasSize = !0
					}).on("error.mfploader",
					function() {
						i.hasSize = !0,
						i.loadError = !0,
						S("LazyLoadError", i)
					}).attr("src", i.src)),
					i.preloaded = !0
				}
			}
		}
	});
	var J = "retina";
	e.magnificPopup.registerModule(J, {
		options: {
			replaceSrc: function(e) {
				return e.src.replace(/\.\w+$/,
				function(e) {
					return "@2x" + e
				})
			},
			ratio: 1
		},
		proto: {
			initRetina: function() {
				if (window.devicePixelRatio > 1) {
					var e = t.st.retina,
					n = e.ratio;
					n = isNaN(n) ? n() : n,
					n > 1 && (C("ImageHasSize." + J,
					function(e, t) {
						t.img.css({
							"max-width": t.img[0].naturalWidth / n,
							width: "100%"
						})
					}), C("ElementParse." + J,
					function(t, i) {
						i.src = e.replaceSrc(i, n)
					}))
				}
			}
		}
	}),
	function() {
		var t = 1e3,
		n = "ontouchstart" in window,
		i = function() {
			k.off("touchmove" + o + " touchend" + o)
		},
		r = "mfpFastClick",
		o = "." + r;
		e.fn.mfpFastClick = function(r) {
			return e(this).each(function() {
				var a, s = e(this);
				if (n) {
					var l, c, u, d, f, p;
					s.on("touchstart" + o,
					function(e) {
						d = !1,
						p = 1,
						f = e.originalEvent ? e.originalEvent.touches[0] : e.touches[0],
						c = f.clientX,
						u = f.clientY,
						k.on("touchmove" + o,
						function(e) {
							f = e.originalEvent ? e.originalEvent.touches: e.touches,
							p = f.length,
							f = f[0],
							(Math.abs(f.clientX - c) > 10 || Math.abs(f.clientY - u) > 10) && (d = !0, i())
						}).on("touchend" + o,
						function(e) {
							i(),
							d || p > 1 || (a = !0, e.preventDefault(), clearTimeout(l), l = setTimeout(function() {
								a = !1
							},
							t), r())
						})
					})
				}
				s.on("click" + o,
				function() {
					a || r()
				})
			})
		},
		e.fn.destroyMfpFastClick = function() {
			e(this).off("touchstart" + o + " click" + o),
			n && k.off("touchmove" + o + " touchend" + o)
		}
	} (),
	$()
}),
window.deviceType = $(window).width() >= 800 ? "desktop": "mobile",
$(window).resize(function() {
	window.deviceType = $(window).width() >= 800 ? "desktop": "mobile"
}),
$(document).ready(function() {
	var e = $(".J_fixedTools");
	$(window).scrollTop() > 200 && !e.hasClass("show") ? e.addClass("show") : $(window).scrollTop() <= 200 && e.hasClass("show") && e.removeClass("show"),
	$(window).scroll(function() {
		$(window).scrollTop() > 200 && !e.hasClass("show") ? e.addClass("show") : $(window).scrollTop() <= 200 && e.hasClass("show") && e.removeClass("show")
	}),
	e.find(".J_up").click(function(e) {
		e.preventDefault(),
		$("html, body").animate({
			scrollTop: 0
		},
		300)
	})
}),
initLazyLoad(),
$(document).ready(function() {
	function e(t) {
		t = t || $(".J_articleList").eq(0),
		$(".J_listLoadMore", t).unbind().click(function(t) {
			t.preventDefault();
			var n = $(this);
			n.hasClass("no-data") || n.hasClass("loading") || (n.addClass("loading"), $.get(n.attr("href"),
			function(t) {
				var i = n.parent();
				$(t).insertAfter(n),
				n.remove(),
				$(window).trigger("scroll"),
				e(i)
			},
			"html"))
		})
	}
	initFastSection(),
	initMobileNav(".J_headImages, .J_articleListWrap"),
	$(".J_newsListNavBar").stick_in_parent(),
	e();
	var t = $(".J_newsListNavBar").length && $(".J_newsListNavBar").offset().top;
	$(window).resize(function() {
		t = $(".J_newsListNavBar").length && $(".J_newsListNavBar").offset().top
	}),
	$(".J_newsListNavBar a").eq(0).data("J_articleListWrap", $(".J_articleList").eq(0)),
	$(".J_newsListNavBar a").click(function(n) {
		n.preventDefault();
		var listid = "articles-list";
		var url = $(this).attr("href")+" #"+listid;
		$("#"+listid).load(url);
		$("#"+listid).addClass("articles");
		$("#"+listid).addClass("J_articleList");  
		$("#"+listid).addClass("ias_container");
	})
}),
$(document).ready(function() {
	var _loginTipstimer
	function logtips(str){
		if( !str ) return false
		_loginTipstimer && clearTimeout(_loginTipstimer)
		$('.sign-tips').html(str).animate({
			height: 29
		}, 220)
		_loginTipstimer = setTimeout(function(){
			$('.sign-tips').animate({
				height: 0
			}, 220)
		}, 5000)
	}
	
	function is_check_name(str) {    
		return /^[\w]{3,16}$/.test(str) 
	}
	function is_check_mail(str) {
		return /^[_a-z0-9-]+(\.[_a-z0-9-]+)*@[a-z0-9-]+(\.[a-z0-9-]+)*(\.[a-z]{2,4})$/.test(str)
	}
		
	var b = $("#sign");
	$("#register-active").on("click",
	function() {
		b.removeClass("sign").addClass("register")
	}),
	$("#login-active").on("click",
	function() {
		b.removeClass("register").addClass("sign")
	}),
	$(".user-login,.user-reg").on("click",
	function(c) {
		$('html').removeClass('holding');
		$('body').removeClass('holding');
		$('body').removeClass('slide-left');
		$('body').removeClass('holding-right');
		$('header').removeClass('show-menu');
		$("div.overlay").length <= 0 ? $(".header-wrap").before('<div class="overlay"></div>') : $("div.overlay").show(),
		$("body").addClass("fadeIn"),
		1 == $(this).attr("data-sign") ? b.removeClass("sign").addClass("register") : b.removeClass("register").addClass("sign"),
		$("div.overlay, form a.close").on("click",
		function() {
			return $("body").removeClass("fadeIn"),
			$.removeAttr("class"),
			$("div.overlay").remove(),
			!1
		}),
		c.preventDefault()
	})
	$('.captcha-clk').bind('click',function(){
		if( !is_check_mail($("#user_email").val()) ){
			logtips('邮箱格式错误')
			return
		}
		
		var captcha = $(this);
		if(captcha.hasClass("disabled")){
			logtips('您操作太快了，等等吧')
		}else{
			captcha.addClass("disabled");
			captcha.html("发送中...");
			$.post(
				_MBT.uri+'/includes/action/captcha.php?'+Math.random(),
				{
					action: "mobantu_captcha",
					email:$("#user_email").val()
				},
				function (data) {
					if($.trim(data) == "1"){
						logtips('已发送验证码至邮箱，可能会出现在垃圾箱里哦~')
						var countdown=60; 
						settime()
						function settime() { 
							if (countdown == 0) { 
								captcha.removeClass("disabled");   
								captcha.html("发送验证码");
								countdown = 60; 
								return;
							} else { 
								captcha.addClass("disabled");
								captcha.html("重新发送(" + countdown + ")"); 
								countdown--; 
							} 
							setTimeout(function() { settime() },1000) 
						}

					}else if($.trim(data) == "2"){
						logtips('邮箱已存在')
						captcha.html("发送验证码");
						captcha.removeClass("disabled"); 
					}else{
						logtips('验证码发送失败，请稍后重试')
						captcha.html("发送验证码");
						captcha.removeClass("disabled"); 
					}
				}
			);
		}
	});
	$('.login-loader').on('click', function(){
		logtips("登录中，请稍等...");									
		$.post(
			_MBT.uri+'signin',
			{
				usr: $("#username").val(),
				pwd: $("#password").val(),
				rememberme: $('#rememberme').val(),
				action: "mobantu_login"
				
			},
			function (data) {
				if (data != "1") {
					logtips("用户名或密码错误");
				}
				else {
					logtips("登录成功，跳转中...");
					location.reload();                     
				}
			}
		);
	})
	$('.register-loader').on('click', function(){
		if( !is_check_name($("#user_name").val()) ){
			logtips('用户名只能由字母数字或下划线组成的3-16位字符')
			return
		}
		
		if( !is_check_mail($("#user_email").val()) ){
			logtips('邮箱格式错误')
			return
		}

		if( $("#user_pass").val().length < 6 ){
			logtips('密码太短，至少6位')
			return
		}
		
		if( $("#user_pass").val() != $("#user_pass2").val()){
			logtips('两次输入密码不一致')
			return
		}
		
		logtips("注册中，请稍等...");
		$.post(
			_MBT.uri+'/includes/action/signup',
			{
				user_register: $("#user_name").val(),
				user_email: $("#user_email").val(),
				password: $("#user_pass").val(),
				captcha: $("#captcha").val(),
				action: "mobantu_register"
			},
			function (data) {
				if (data == "1") {
					logtips("注册成功，登录中...");
					location.reload(); 
				}
				else {
					logtips(data);
				}
			}
		);										   
	})
});
$(document).ready(function() {
	if ($("body").hasClass('single')) {
		$(".article iframe").each(function(){
			  $(this).width("100%");
			  $(this).height($(this).width()*9/16);
		});
		$(".article embed").each(function(){
			  $(this).width("100%");
			  $(this).height($(this).width()*9/16);
		});
		$(".article video").each(function(){
			  $(this).width("100%");
			  $(this).height($(this).width()*9/16);
		});
	}
    initScroll();
	$('.J_commonHeaderWrapper .J_menuTrigger').bind('mouseup touchend',function(e){
		e.preventDefault();
		e.stopPropagation();
		var wrapper = $('.J_commonHeaderWrapper');
		var trigger = $(this);

		function cancel(e){
			var cur = $(e.target);

			if(!cur.parents('.J_commonHeaderWrapper').length || cur.parents('.triggers').length){
				e.preventDefault && e.preventDefault();
				$('body').removeClass('slide-left');
				$('body').unbind('mouseup touchstart', cancel);
				$('.J_commonHeaderWrapper .J_navList a').unbind('mouseup touchend');
				setTimeout(function(){
					trigger.removeClass('active');
					releaseBody();
					wrapper.removeClass('show-menu');
				}, 400);
			}
		}

		if(trigger.hasClass('active')){
			//cancel(e);
			return;
		}
		trigger.addClass('active');
		holdBody(true);
		wrapper.addClass('show-menu');
		setTimeout(function(){
			$('body').bind('mouseup touchstart', cancel);
			/*$('.J_commonHeaderWrapper .J_navList a').bind('mouseup touchend', function(){
				$('.J_commonHeaderWrapper .J_navList a').unbind('mouseup touchend', arguments.callee);
				cancel({
					target:$()
				});
			});*/
		},100);
	}),
	$('.J_commonHeaderWrapper .J_searchTrigger').bind('mouseup touchend', function bindSearch(e){
		e.preventDefault();
		var wrapper = $('.J_commonHeaderWrapper');
		var trigger = $(this);
		holdBody();
		wrapper.addClass('show-search');
		wrapper.find('.searchbar input').trigger('focus');

	});

	$('.J_commonHeaderWrapper .searchbar .close-icon').bind('mouseup touchend', function(e){
		e.preventDefault();
		setTimeout(function(){
			var wrapper = $('.J_commonHeaderWrapper');
			wrapper.removeClass('show-search');
			releaseBody();
		}, 100);
	});
	
	function holdBody (slide){
		$('html').addClass('holding');
		$('body').addClass('holding');
		if(slide){
			$('body').addClass('slide-left');
			$('body').addClass('holding-right');
		}
		window.scrollTo(0,0);
	}
	function releaseBody (){
		$('html').removeClass('holding');
		$('body').removeClass('holding');
		$('body').removeClass('slide-left');
		$('body').removeClass('holding-right');

	}
	/**
     * 初始化滚动动作
     */
    function initScroll(){
        /*if(fixed=="false"){
            $('body').css({
                paddingTop:0
            });
            $('.J_commonHeaderWrapper').css({
                position:'relative'
            });
            return;
        }*/
        $(window).scroll(function(){
            var top = $(window).scrollTop();
            if($(window).width()<=820)return;
            var wrap = $('.J_commonHeaderWrapper');
            if(top>0 && !wrap.hasClass('scrolling')){
                wrap.addClass('scrolling');
            }else if(top==0){
                wrap.removeClass('scrolling');
            }
        });
    }
}),
$(document).ready(function() {
	initFastSection(),
	initMobileNav(".content-wrapper"),
	$(".J_addCommentBtn").click(function(e) {
		e.preventDefault(),
		$("body, html").animate({
			scrollTop: $(".single-post-comment").offset().top
		},
		400),
		$(".single-post-comment textarea").trigger("focus")
	}),
	$(".J_showAllShareBtn").click(function(e) {
		e.preventDefault(),
		$(this).siblings(".external").removeClass("external"),
		$(this).hide(),
		$(this).parents(".hide-external").removeClass("hide-external")
	})
}),
$(document).ready(function() {
	function e(t) {
		t = t || $(".J_articleList").eq(0),
		$(".J_listLoadMore", t).unbind().click(function(t) {
			t.preventDefault();
			var n = $(this);
			n.hasClass("no-data") || n.hasClass("loading") || (n.addClass("loading"), $.get(n.attr("href"),
			function(t) {
				var i = n.parent();
				$(t).insertAfter(n),
				n.remove(),
				$(window).trigger("scroll"),
				e(i)
			},
			"html"))
		})
	}
	initFastSection(),
	initMobileNav(".J_resultListWrap"),
	$(".J_newsListNavBar").stick_in_parent(),
	e()
}),
$(document).ready(function() {
	initFastSection(),
	initMobileNav(".J_fastNewsDetailWrap")
}),
function() {
	jQuery(function() {
		return window.ga ? (ga("send", "pageview"), $("#headline_top").click(function() {
			return ga("send", "event", "link", "headline#top", ga_user_id)
		}), $("#headline_one").click(function() {
			return ga("send", "event", "link", "headline#one", ga_user_id)
		}), $("#headline_two").click(function() {
			return ga("send", "event", "link", "headline#two", ga_user_id)
		}), $("#headline_three").click(function() {
			return ga("send", "event", "link", "headline#three", ga_user_id)
		}), $("#headline_four").click(function() {
			return ga("send", "event", "link", "headline#four", ga_user_id)
		}), $("#headline_five").click(function() {
			return ga("send", "event", "link", "headline#five", ga_user_id)
		}), $(".J_login").click(function() {
			return ga("send", "event", "link", "login#top_nav", ga_user_id)
		}), $(".login_before_comment").click(function() {
			return ga("send", "event", "link", "login#comment", ga_user_id)
		}), $(".J_searchForm input").focus(function() {
			return ga("send", "event", "link", "search#focus", ga_user_id)
		}), $("#info_flows_next_link").click(function() {
			return ga("send", "event", "link", "paginate#info_flows#next", ga_user_id)
		}), $("#columns_next_link").click(function() {
			return ga("send", "event", "link", "paginate#columns#next", ga_user_id)
		}), $("#search_next_link").click(function() {
			return ga("send", "event", "link", "paginate#search#next", ga_user_id)
		}), $("#user_domain_next_link").click(function() {
			return ga("send", "event", "link", "paginate#newsflashes#next", ga_user_id)
		}), $(".newsflashes_next_link").click(function() {
			return ga("send", "event", "link", "paginate#product_notes#next", ga_user_id)
		}), $(".product_notes_next_link").click(function() {
			return ga("send", "event", "link", "paginate#user_domain#next", ga_user_id)
		}), $("#tag_next_link").click(function() {
			return ga("send", "event", "link", "paginate#tag#next", ga_user_id)
		}), $(".info_flow_news_title").click(function() {
			return ga("send", "event", "link", "news#info_flows#title", ga_user_id)
		}), $(".info_flow_news_image").click(function() {
			return ga("send", "event", "link", "news#info_flows#image", ga_user_id)
		}), $(".J_logout").click(function() {
			return ga("send", "event", "link", "logout#top_nav", ga_user_id)
		}), $(".categories a").click(function() {
			return ga("send", "event", "link", "nav#categories", ga_user_id)
		}), $(".app-download").hover(function() {
			return ga("send", "event", "link", "app_download#top_nav", ga_user_id)
		}), $(".app_download_footer").click(function() {
			return ga("send", "event", "link", "app_download#footer", ga_user_id)
		}), $(".compiled.rong-company-link").click(function() {
			return ga("send", "event", "link", "rong-company-link#click", ga_user_id)
		}), $(".compiled.rong-company-link").hover(function() {
			return ga("send", "event", "hover", "rong-company-link#hover", ga_user_id)
		}), $(".newsflashes_nav").click(function() {
			return ga("send", "event", "link", "aside-nav#newsflashes", ga_user_id)
		}), $(".product_notes_nav").click(function() {
			return ga("send", "event", "link", "aside-nav#product_notes", ga_user_id)
		})) : void 0
	})
}.call(this),
$(document).ready(function() {
	function e(t) {
		t = t || $(".J_articleList").eq(0),
		$(".J_listLoadMore", t).unbind().click(function(t) {
			t.preventDefault();
			var n = $(this);
			$.get(n.attr("href"),
			function(t) {
				var i = n.parent();
				$(t).insertAfter(n),
				n.remove(),
				$(window).trigger("scroll"),
				e(i)
			},
			"html")
		})
	}
	initFastSection(),
	initMobileNav(".J_resultListWrap"),
	$(".J_newsListNavBar").stick_in_parent(),
	e()
}),
function() {
	var AsymcRender;
	AsymcRender = function() {
		function AsymcRender(e) {
			this.stack = e,
			this.stack = $("div[async]")
		}
		return AsymcRender.prototype.render = function() {
			return $.each(this.stack,
			function(key, obj) {
				return $.get($(obj).attr("async-url"), {},
				function(data) {
					return $(obj).hide().html(data).fadeIn(1500),
					$(obj).attr("async-callback") ? eval($(obj).attr("async-callback")) : void 0
				})
			})
		},
		AsymcRender
	} (),
	jQuery(function() {
		var e;
		return e = new AsymcRender,
		e.render()
	})
}.call(this),
function() {
	var e;
	e = function() {
		function e(e, t) {
			this.startInterval = 6e4,
			this.init(e, t)
		}
		return e.prototype.init = function(e, t) {
			return this.$element = $(e),
			this.options = $.extend({},
			$.fn.timeago.defaults, t),
			this.updateTime(),
			this.startTimer()
		},
		e.prototype.startTimer = function() {
			var e;
			return e = this,
			this.interval = setInterval(function() {
				return e.refresh()
			},
			this.startInterval)
		},
		e.prototype.stopTimer = function() {
			return clearInterval(this.interval)
		},
		e.prototype.restartTimer = function() {
			return this.stopTimer(),
			this.startTimer()
		},
		e.prototype.refresh = function() {
			return this.updateTime(),
			this.updateInterval()
		},
		e.prototype.updateTime = function() {
			var e;
			return e = this,
			this.$element.findAndSelf(this.options.selector).each(function() {
				var t;
				return t = e.timeAgoInWords($(this).attr(e.options.attr)),
				$(this).html(t)
			})
		},
		e.prototype.updateInterval = function() {
			var e, t, n, i;
			if (this.$element.findAndSelf(this.options.selector).length > 0) {
				if ("up" === this.options.dir ? e = ":first": "down" === this.options.dir && (e = ":last"), i = this.$element.findAndSelf(this.options.selector).filter(e).attr(this.options.attr), t = this.parse(i), n = this.getTimeDistanceInMinutes(t), n >= 0 && 44 >= n && 6e4 !== this.startInterval) return this.startInterval = 6e4,
				this.restartTimer();
				if (n >= 45 && 89 >= n && 132e4 !== this.startInterval) return this.startInterval = 132e4,
				this.restartTimer();
				if (n >= 90 && 2519 >= n && 18e5 !== this.startInterval) return this.startInterval = 18e5,
				this.restartTimer();
				if (n >= 2520 && 432e5 !== this.startInterval) return this.startInterval = 432e5,
				this.restartTimer()
			}
		},
		e.prototype.timeAgoInWords = function(e) {
			var t;
			return t = this.parse(e),
			"" + this.options.lang.prefixes.ago + this.distanceOfTimeInWords(t) + this.options.lang.suffix
		},
		e.prototype.parse = function(e) {
			var t;
			return t = $.trim(e),
			t = t.replace(/\.\d\d\d+/, ""),
			t = t.replace(/-/, "/").replace(/-/, "/"),
			t = t.replace(/T/, " ").replace(/Z/, " UTC"),
			t = t.replace(/([\+\-]\d\d)\:?(\d\d)/, " $1$2"),
			new Date(t)
		},
		e.prototype.getTimeDistanceInMinutes = function(e) {
			var t;
			return t = (new Date).getTime() - e.getTime(),
			Math.round(Math.abs(t) / 1e3 / 60)
		},
		e.prototype.distanceOfTimeInWords = function(e) {
			var t;
			return t = this.getTimeDistanceInMinutes(e),
			0 === t ? "" + this.options.lang.prefixes.lt + " " + this.options.lang.units.minute: 1 === t ? "1 " + this.options.lang.units.minute: t >= 2 && 44 >= t ? "" + t + " " + this.options.lang.units.minutes: t >= 45 && 89 >= t ? "" + this.options.lang.prefixes.about + " 1 " + this.options.lang.units.hour: t >= 90 && 1439 >= t ? "" + this.options.lang.prefixes.about + " " + Math.round(t / 60) + " " + this.options.lang.units.hours: t >= 1440 && 2519 >= t ? "1 " + this.options.lang.units.day: t >= 2520 && 43199 >= t ? "" + Math.round(t / 1440) + " " + this.options.lang.units.days: t >= 43200 && 86399 >= t ? "" + this.options.lang.prefixes.about + " 1 " + this.options.lang.units.month: t >= 86400 && 525599 >= t ? "" + Math.round(t / 43200) + " " + this.options.lang.units.months: t >= 525600 && 655199 >= t ? "" + this.options.lang.prefixes.about + " 1 " + this.options.lang.units.year: t >= 655200 && 914399 >= t ? "" + this.options.lang.prefixes.over + " 1 " + this.options.lang.units.year: t >= 914400 && 1051199 >= t ? "" + this.options.lang.prefixes.almost + " 2 " + this.options.lang.units.years: "" + this.options.lang.prefixes.about + " " + Math.round(t / 525600) + " " + this.options.lang.units.years
		},
		e
	} (),
	$.fn.timeago = function(t) {
		return null == t && (t = {}),
		this.each(function() {
			var n, i;
			return n = $(this),
			i = n.data("timeago"),
			i ? "string" == typeof t ? i[t]() : void 0 : n.data("timeago", new e(this, t))
		})
	},
	$.fn.findAndSelf = function(e) {
		return this.find(e).add(this.filter(e))
	},
	$.fn.timeago.Constructor = e,
	$.fn.timeago.defaults = {
		selector: "time.timeago",
		attr: "datetime",
		dir: "up",
		lang: {
			units: {
				second: "second",
				seconds: "seconds",
				minute: "minute",
				minutes: "minutes",
				hour: "hour",
				hours: "hours",
				day: "day",
				days: "days",
				month: "month",
				months: "months",
				year: "year",
				years: "years"
			},
			prefixes: {
				lt: "less than a",
				about: "about",
				over: "over",
				almost: "almost",
				ago: ""
			},
			suffix: " ago"
		}
	}
}.call(this),
function() {
	$.fn.timeago.defaults.lang = {
		units: {
			second: "\u79d2",
			seconds: "\u79d2",
			minute: "\u5206\u949f",
			minutes: "\u5206\u949f",
			hour: "\u5c0f\u65f6",
			hours: "\u5c0f\u65f6",
			day: "\u5929",
			days: "\u5929",
			month: "\u6708",
			months: "\u6708",
			year: "\u5e74",
			years: "\u5e74"
		},
		prefixes: {
			lt: "1",
			about: "",
			over: "",
			almost: "",
			ago: ""
		},
		suffix: "\u524d"
	}
}.call(this),
function() { (!window.BAIDU_CLB_fillSlot || $("html").is(".portable-device")) && (window.BAIDU_CLB_fillSlot = function() {}),
	function() {
		var e;
		return e = alert,
		window.alert = function(t) {
			return t.match("malicious javascript") ? void 0 : e(t)
		}
	} (),
	window.mobilecheck = function() {
		var e;
		return e = !1,
		function(t) { (/(android|bb\d+|meego).+mobile|avantgo|bada\/|blackberry|blazer|compal|elaine|fennec|hiptop|iemobile|ip(hone|od)|iris|kindle|lge |maemo|midp|mmp|mobile.+firefox|netfront|opera m(ob|in)i|palm( os)?|phone|p(ixi|re)\/|plucker|pocket|psp|series(4|6)0|symbian|treo|up\.(browser|link)|vodafone|wap|windows ce|xda|xiino/i.test(t) || /1207|6310|6590|3gso|4thp|50[1-6]i|770s|802s|a wa|abac|ac(er|oo|s\-)|ai(ko|rn)|al(av|ca|co)|amoi|an(ex|ny|yw)|aptu|ar(ch|go)|as(te|us)|attw|au(di|\-m|r |s )|avan|be(ck|ll|nq)|bi(lb|rd)|bl(ac|az)|br(e|v)w|bumb|bw\-(n|u)|c55\/|capi|ccwa|cdm\-|cell|chtm|cldc|cmd\-|co(mp|nd)|craw|da(it|ll|ng)|dbte|dc\-s|devi|dica|dmob|do(c|p)o|ds(12|\-d)|el(49|ai)|em(l2|ul)|er(ic|k0)|esl8|ez([4-7]0|os|wa|ze)|fetc|fly(\-|_)|g1 u|g560|gene|gf\-5|g\-mo|go(\.w|od)|gr(ad|un)|haie|hcit|hd\-(m|p|t)|hei\-|hi(pt|ta)|hp( i|ip)|hs\-c|ht(c(\-| |_|a|g|p|s|t)|tp)|hu(aw|tc)|i\-(20|go|ma)|i230|iac( |\-|\/)|ibro|idea|ig01|ikom|im1k|inno|ipaq|iris|ja(t|v)a|jbro|jemu|jigs|kddi|keji|kgt( |\/)|klon|kpt |kwc\-|kyo(c|k)|le(no|xi)|lg( g|\/(k|l|u)|50|54|\-[a-w])|libw|lynx|m1\-w|m3ga|m50\/|ma(te|ui|xo)|mc(01|21|ca)|m\-cr|me(rc|ri)|mi(o8|oa|ts)|mmef|mo(01|02|bi|de|do|t(\-| |o|v)|zz)|mt(50|p1|v )|mwbp|mywa|n10[0-2]|n20[2-3]|n30(0|2)|n50(0|2|5)|n7(0(0|1)|10)|ne((c|m)\-|on|tf|wf|wg|wt)|nok(6|i)|nzph|o2im|op(ti|wv)|oran|owg1|p800|pan(a|d|t)|pdxg|pg(13|\-([1-8]|c))|phil|pire|pl(ay|uc)|pn\-2|po(ck|rt|se)|prox|psio|pt\-g|qa\-a|qc(07|12|21|32|60|\-[2-7]|i\-)|qtek|r380|r600|raks|rim9|ro(ve|zo)|s55\/|sa(ge|ma|mm|ms|ny|va)|sc(01|h\-|oo|p\-)|sdk\/|se(c(\-|0|1)|47|mc|nd|ri)|sgh\-|shar|sie(\-|m)|sk\-0|sl(45|id)|sm(al|ar|b3|it|t5)|so(ft|ny)|sp(01|h\-|v\-|v )|sy(01|mb)|t2(18|50)|t6(00|10|18)|ta(gt|lk)|tcl\-|tdg\-|tel(i|m)|tim\-|t\-mo|to(pl|sh)|ts(70|m\-|m3|m5)|tx\-9|up(\.b|g1|si)|utst|v400|v750|veri|vi(rg|te)|vk(40|5[0-3]|\-v)|vm40|voda|vulc|vx(52|53|60|61|70|80|81|83|85|98)|w3c(\-| )|webc|whit|wi(g |nc|nw)|wmlb|wonu|x700|yas\-|your|zeto|zte\-/i.test(t.substr(0, 4))) && (e = !0)
		} (navigator.userAgent || navigator.vendor || window.opera),
		e
	},
	window.doFavorite = function(e) {
		$.post("/asynces/favorites", {
			url_code: e,
			authenticity_token: window._token
		},
		function(e) {
			var t;
			t = $(".J_addFavorite"),
			t.hasClass("is-favorite") ? (t.removeClass("is-favorite"), t.removeClass("icon-fly")) : (t.addClass("is-favorite"), setTimeout(function() {
				t.addClass("icon-fly")
			},
			0)),
			"add" === e.success ? (t.addClass("is-favorite"), setTimeout(function() {
				t.addClass("icon-fly"),
				$("#star-count").text(e.count)
			},
			0)) : "del" === e.success && (t.removeClass("is-favorite"), t.removeClass("icon-fly"), $("#star-count").text(e.count))
		})
	},
	jQuery(function() {
		$(".dropdown_login_out_link").on("click",
		function() {
			$(".real_login_out_link").trigger("click")
		}),
		$(".require-login").on("click",
		function() {
			void 0 !== $(".require-login").data().uid && "" !== $(".require-login").data().uid || !confirm("\u8bf7\u767b\u5f55\u540e\u7ee7\u7eed\u64cd\u4f5c\uff01 (\u25cf\u2014\u25cf)") || (window.location.href = "/users/auth/krypton")
		}),
		$(".single-post-header__headline img[src*=yestone]").after('<small><a href="http://yestone.com/?utm_source=36kr.com">\u56fe\u7247: Yestone.com \u7248\u6743\u56fe\u7247\u5e93</a></small>')
	})
}.call(this);