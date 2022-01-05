/*!
 * jQuery JavaScript Library v3.6.0
 * https://jquery.com/
 *
 * Includes Sizzle.js
 * https://sizzlejs.com/
 *
 * Copyright OpenJS Foundation and other contributors
 * Released under the MIT license
 * https://jquery.org/license
 *
 * Date: 2021-03-02T17:08Z
 */
( function( global, factory ) {

	"use strict";

	if ( typeof module === "object" && typeof module.exports === "object" ) {

		// For CommonJS and CommonJS-like environments where a proper `window`
		// is present, execute the factory and get jQuery.
		// For environments that do not have a `window` with a `document`
		// (such as Node.js), expose a factory as module.exports.
		// This accentuates the need for the creation of a real `window`.
		// e.g. var jQuery = require("jquery")(window);
		// See ticket #14549 for more info.
		module.exports = global.document ?
			factory( global, true ) :
			function( w ) {
				if ( !w.document ) {
					throw new Error( "jQuery requires a window with a document" );
				}
				return factory( w );
			};
	} else {
		factory( global );
	}

// Pass this if window is not defined yet
} )( typeof window !== "undefined" ? window : this, function( window, noGlobal ) {

// Edge <= 12 - 13+, Firefox <=18 - 45+, IE 10 - 11, Safari 5.1 - 9+, iOS 6 - 9.1
// throw exceptions when non-strict code (e.g., ASP.NET 4.5) accesses strict mode
// arguments.callee.caller (trac-13335). But as of jQuery 3.0 (2016), strict mode should be common
// enough that all such attempts are guarded in a try block.
"use strict";

var arr = [];

var getProto = Object.getPrototypeOf;

var slice = arr.slice;

var flat = arr.flat ? function( array ) {
	return arr.flat.call( array );
} : function( array ) {
	return arr.concat.apply( [], array );
};


var push = arr.push;

var indexOf = arr.indexOf;

var class2type = {};

var toString = class2type.toString;

var hasOwn = class2type.hasOwnProperty;

var fnToString = hasOwn.toString;

var ObjectFunctionString = fnToString.call( Object );

var support = {};

var isFunction = function isFunction( obj ) {

		// Support: Chrome <=57, Firefox <=52
		// In some browsers, typeof returns "function" for HTML <object> elements
		// (i.e., `typeof document.createElement( "object" ) === "function"`).
		// We don't want to classify *any* DOM node as a function.
		// Support: QtWeb <=3.8.5, WebKit <=534.34, wkhtmltopdf tool <=0.12.5
		// Plus for old WebKit, typeof returns "function" for HTML collections
		// (e.g., `typeof document.getElementsByTagName("div") === "function"`). (gh-4756)
		return typeof obj === "function" && typeof obj.nodeType !== "number" &&
			typeof obj.item !== "function";
	};


var isWindow = function isWindow( obj ) {
		return obj != null && obj === obj.window;
	};


var document = window.document;



	var preservedScriptAttributes = {
		type: true,
		src: true,
		nonce: true,
		noModule: true
	};

	function DOMEval( code, node, doc ) {
		doc = doc || document;

		var i, val,
			script = doc.createElement( "script" );

		script.text = code;
		if ( node ) {
			for ( i in preservedScriptAttributes ) {

				// Support: Firefox 64+, Edge 18+
				// Some browsers don't support the "nonce" property on scripts.
				// On the other hand, just using `getAttribute` is not enough as
				// the `nonce` attribute is reset to an empty string whenever it
				// becomes browsing-context connected.
				// See https://github.com/whatwg/html/issues/2369
				// See https://html.spec.whatwg.org/#nonce-attributes
				// The `node.getAttribute` check was added for the sake of
				// `jQuery.globalEval` so that it can fake a nonce-containing node
				// via an object.
				val = node[ i ] || node.getAttribute && node.getAttribute( i );
				if ( val ) {
					script.setAttribute( i, val );
				}
			}
		}
		doc.head.appendChild( script ).parentNode.removeChild( script );
	}


function toType( obj ) {
	if ( obj == null ) {
		return obj + "";
	}

	// Support: Android <=2.3 only (functionish RegExp)
	return typeof obj === "object" || typeof obj === "function" ?
		class2type[ toString.call( obj ) ] || "object" :
		typeof obj;
}
/* global Symbol */
// Defining this global in .eslintrc.json would create a danger of using the global
// unguarded in another place, it seems safer to define global only for this module



var
	version = "3.6.0",

	// Define a local copy of jQuery
	jQuery = function( selector, context ) {

		// The jQuery object is actually just the init constructor 'enhanced'
		// Need init if jQuery is called (just allow error to be thrown if not included)
		return new jQuery.fn.init( selector, context );
	};

jQuery.fn = jQuery.prototype = {

	// The current version of jQuery being used
	jquery: version,

	constructor: jQuery,

	// The default length of a jQuery object is 0
	length: 0,

	toArray: function() {
		return slice.call( this );
	},

	// Get the Nth element in the matched element set OR
	// Get the whole matched element set as a clean array
	get: function( num ) {

		// Return all the elements in a clean array
		if ( num == null ) {
			return slice.call( this );
		}

		// Return just the one element from the set
		return num < 0 ? this[ num + this.length ] : this[ num ];
	},

	// Take an array of elements and push it onto the stack
	// (returning the new matched element set)
	pushStack: function( elems ) {

		// Build a new jQuery matched element set
		var ret = jQuery.merge( this.constructor(), elems );

		// Add the old object onto the stack (as a reference)
		ret.prevObject = this;

		// Return the newly-formed element set
		return ret;
	},

	// Execute a callback for every element in the matched set.
	each: function( callback ) {
		return jQuery.each( this, callback );
	},

	map: function( callback ) {
		return this.pushStack( jQuery.map( this, function( elem, i ) {
			return callback.call( elem, i, elem );
		} ) );
	},

	slice: function() {
		return this.pushStack( slice.apply( this, arguments ) );
	},

	first: function() {
		return this.eq( 0 );
	},

	last: function() {
		return this.eq( -1 );
	},

	even: function() {
		return this.pushStack( jQuery.grep( this, function( _elem, i ) {
			return ( i + 1 ) % 2;
		} ) );
	},

	odd: function() {
		return this.pushStack( jQuery.grep( this, function( _elem, i ) {
			return i % 2;
		} ) );
	},

	eq: function( i ) {
		var len = this.length,
			j = +i + ( i < 0 ? len : 0 );
		return this.pushStack( j >= 0 && j < len ? [ this[ j ] ] : [] );
	},

	end: function() {
		return this.prevObject || this.constructor();
	},

	// For internal use only.
	// Behaves like an Array's method, not like a jQuery method.
	push: push,
	sort: arr.sort,
	splice: arr.splice
};

jQuery.extend = jQuery.fn.extend = function() {
	var options, name, src, copy, copyIsArray, clone,
		target = arguments[ 0 ] || {},
		i = 1,
		length = arguments.length,
		deep = false;

	// Handle a deep copy situation
	if ( typeof target === "boolean" ) {
		deep = target;

		// Skip the boolean and the target
		target = arguments[ i ] || {};
		i++;
	}

	// Handle case when target is a string or something (possible in deep copy)
	if ( typeof target !== "object" && !isFunction( target ) ) {
		target = {};
	}

	// Extend jQuery itself if only one argument is passed
	if ( i === length ) {
		target = this;
		i--;
	}

	for ( ; i < length; i++ ) {

		// Only deal with non-null/undefined values
		if ( ( options = arguments[ i ] ) != null ) {

			// Extend the base object
			for ( name in options ) {
				copy = options[ name ];

				// Prevent Object.prototype pollution
				// Prevent never-ending loop
				if ( name === "__proto__" || target === copy ) {
					continue;
				}

				// Recurse if we're merging plain objects or arrays
				if ( deep && copy && ( jQuery.isPlainObject( copy ) ||
					( copyIsArray = Array.isArray( copy ) ) ) ) {
					src = target[ name ];

					// Ensure proper type for the source value
					if ( copyIsArray && !Array.isArray( src ) ) {
						clone = [];
					} else if ( !copyIsArray && !jQuery.isPlainObject( src ) ) {
						clone = {};
					} else {
						clone = src;
					}
					copyIsArray = false;

					// Never move original objects, clone them
					target[ name ] = jQuery.extend( deep, clone, copy );

				// Don't bring in undefined values
				} else if ( copy !== undefined ) {
					target[ name ] = copy;
				}
			}
		}
	}

	// Return the modified object
	return target;
};

jQuery.extend( {

	// Unique for each copy of jQuery on the page
	expando: "jQuery" + ( version + Math.random() ).replace( /\D/g, "" ),

	// Assume jQuery is ready without the ready module
	isReady: true,

	error: function( msg ) {
		throw new Error( msg );
	},

	noop: function() {},

	isPlainObject: function( obj ) {
		var proto, Ctor;

		// Detect obvious negatives
		// Use toString instead of jQuery.type to catch host objects
		if ( !obj || toString.call( obj ) !== "[object Object]" ) {
			return false;
		}

		proto = getProto( obj );

		// Objects with no prototype (e.g., `Object.create( null )`) are plain
		if ( !proto ) {
			return true;
		}

		// Objects with prototype are plain iff they were constructed by a global Object function
		Ctor = hasOwn.call( proto, "constructor" ) && proto.constructor;
		return typeof Ctor === "function" && fnToString.call( Ctor ) === ObjectFunctionString;
	},

	isEmptyObject: function( obj ) {
		var name;

		for ( name in obj ) {
			return false;
		}
		return true;
	},

	// Evaluates a script in a provided context; falls back to the global one
	// if not specified.
	globalEval: function( code, options, doc ) {
		DOMEval( code, { nonce: options && options.nonce }, doc );
	},

	each: function( obj, callback ) {
		var length, i = 0;

		if ( isArrayLike( obj ) ) {
			length = obj.length;
			for ( ; i < length; i++ ) {
				if ( callback.call( obj[ i ], i, obj[ i ] ) === false ) {
					break;
				}
			}
		} else {
			for ( i in obj ) {
				if ( callback.call( obj[ i ], i, obj[ i ] ) === false ) {
					break;
				}
			}
		}

		return obj;
	},

	// results is for internal usage only
	makeArray: function( arr, results ) {
		var ret = results || [];

		if ( arr != null ) {
			if ( isArrayLike( Object( arr ) ) ) {
				jQuery.merge( ret,
					typeof arr === "string" ?
						[ arr ] : arr
				);
			} else {
				push.call( ret, arr );
			}
		}

		return ret;
	},

	inArray: function( elem, arr, i ) {
		return arr == null ? -1 : indexOf.call( arr, elem, i );
	},

	// Support: Android <=4.0 only, PhantomJS 1 only
	// push.apply(_, arraylike) throws on ancient WebKit
	merge: function( first, second ) {
		var len = +second.length,
			j = 0,
			i = first.length;

		for ( ; j < len; j++ ) {
			first[ i++ ] = second[ j ];
		}

		first.length = i;

		return first;
	},

	grep: function( elems, callback, invert ) {
		var callbackInverse,
			matches = [],
			i = 0,
			length = elems.length,
			callbackExpect = !invert;

		// Go through the array, only saving the items
		// that pass the validator function
		for ( ; i < length; i++ ) {
			callbackInverse = !callback( elems[ i ], i );
			if ( callbackInverse !== callbackExpect ) {
				matches.push( elems[ i ] );
			}
		}

		return matches;
	},

	// arg is for internal usage only
	map: function( elems, callback, arg ) {
		var length, value,
			i = 0,
			ret = [];

		// Go through the array, translating each of the items to their new values
		if ( isArrayLike( elems ) ) {
			length = elems.length;
			for ( ; i < length; i++ ) {
				value = callback( elems[ i ], i, arg );

				if ( value != null ) {
					ret.push( value );
				}
			}

		// Go through every key on the object,
		} else {
			for ( i in elems ) {
				value = callback( elems[ i ], i, arg );

				if ( value != null ) {
					ret.push( value );
				}
			}
		}

		// Flatten any nested arrays
		return flat( ret );
	},

	// A global GUID counter for objects
	guid: 1,

	// jQuery.support is not used in Core but other projects attach their
	// properties to it so it needs to exist.
	support: support
} );

if ( typeof Symbol === "function" ) {
	jQuery.fn[ Symbol.iterator ] = arr[ Symbol.iterator ];
}

// Populate the class2type map
jQuery.each( "Boolean Number String Function Array Date RegExp Object Error Symbol".split( " " ),
	function( _i, name ) {
		class2type[ "[object " + name + "]" ] = name.toLowerCase();
	} );

function isArrayLike( obj ) {

	// Support: real iOS 8.2 only (not reproducible in simulator)
	// `in` check used to prevent JIT error (gh-2145)
	// hasOwn isn't used here due to false negatives
	// regarding Nodelist length in IE
	var length = !!obj && "length" in obj && obj.length,
		type = toType( obj );

	if ( isFunction( obj ) || isWindow( obj ) ) {
		return false;
	}

	return type === "array" || length === 0 ||
		typeof length === "number" && length > 0 && ( length - 1 ) in obj;
}
var Sizzle =
/*!
 * Sizzle CSS Selector Engine v2.3.6
 * https://sizzlejs.com/
 *
 * Copyright JS Foundation and other contributors
 * Released under the MIT license
 * https://js.foundation/
 *
 * Date: 2021-02-16
 */
( function( window ) {
var i,
	support,
	Expr,
	getText,
	isXML,
	tokenize,
	compile,
	select,
	outermostContext,
	sortInput,
	hasDuplicate,

	// Local document vars
	setDocument,
	document,
	docElem,
	documentIsHTML,
	rbuggyQSA,
	rbuggyMatches,
	matches,
	contains,

	// Instance-specific data
	expando = "sizzle" + 1 * new Date(),
	preferredDoc = window.document,
	dirruns = 0,
	done = 0,
	classCache = createCache(),
	tokenCache = createCache(),
	compilerCache = createCache(),
	nonnativeSelectorCache = createCache(),
	sortOrder = function( a, b ) {
		if ( a === b ) {
			hasDuplicate = true;
		}
		return 0;
	},

	// Instance methods
	hasOwn = ( {} ).hasOwnProperty,
	arr = [],
	pop = arr.pop,
	pushNative = arr.push,
	push = arr.push,
	slice = arr.slice,

	// Use a stripped-down indexOf as it's faster than native
	// https://jsperf.com/thor-indexof-vs-for/5
	indexOf = function( list, elem ) {
		var i = 0,
			len = list.length;
		for ( ; i < len; i++ ) {
			if ( list[ i ] === elem ) {
				return i;
			}
		}
		return -1;
	},

	booleans = "checked|selected|async|autofocus|autoplay|controls|defer|disabled|hidden|" +
		"ismap|loop|multiple|open|readonly|required|scoped",

	// Regular expressions

	// http://www.w3.org/TR/css3-selectors/#whitespace
	whitespace = "[\\x20\\t\\r\\n\\f]",

	// https://www.w3.org/TR/css-syntax-3/#ident-token-diagram
	identifier = "(?:\\\\[\\da-fA-F]{1,6}" + whitespace +
		"?|\\\\[^\\r\\n\\f]|[\\w-]|[^\0-\\x7f])+",

	// Attribute selectors: http://www.w3.org/TR/selectors/#attribute-selectors
	attributes = "\\[" + whitespace + "*(" + identifier + ")(?:" + whitespace +

		// Operator (capture 2)
		"*([*^$|!~]?=)" + whitespace +

		// "Attribute values must be CSS identifiers [capture 5]
		// or strings [capture 3 or capture 4]"
		"*(?:'((?:\\\\.|[^\\\\'])*)'|\"((?:\\\\.|[^\\\\\"])*)\"|(" + identifier + "))|)" +
		whitespace + "*\\]",

	pseudos = ":(" + identifier + ")(?:\\((" +

		// To reduce the number of selectors needing tokenize in the preFilter, prefer arguments:
		// 1. quoted (capture 3; capture 4 or capture 5)
		"('((?:\\\\.|[^\\\\'])*)'|\"((?:\\\\.|[^\\\\\"])*)\")|" +

		// 2. simple (capture 6)
		"((?:\\\\.|[^\\\\()[\\]]|" + attributes + ")*)|" +

		// 3. anything else (capture 2)
		".*" +
		")\\)|)",

	// Leading and non-escaped trailing whitespace, capturing some non-whitespace characters preceding the latter
	rwhitespace = new RegExp( whitespace + "+", "g" ),
	rtrim = new RegExp( "^" + whitespace + "+|((?:^|[^\\\\])(?:\\\\.)*)" +
		whitespace + "+$", "g" ),

	rcomma = new RegExp( "^" + whitespace + "*," + whitespace + "*" ),
	rcombinators = new RegExp( "^" + whitespace + "*([>+~]|" + whitespace + ")" + whitespace +
		"*" ),
	rdescend = new RegExp( whitespace + "|>" ),

	rpseudo = new RegExp( pseudos ),
	ridentifier = new RegExp( "^" + identifier + "$" ),

	matchExpr = {
		"ID": new RegExp( "^#(" + identifier + ")" ),
		"CLASS": new RegExp( "^\\.(" + identifier + ")" ),
		"TAG": new RegExp( "^(" + identifier + "|[*])" ),
		"ATTR": new RegExp( "^" + attributes ),
		"PSEUDO": new RegExp( "^" + pseudos ),
		"CHILD": new RegExp( "^:(only|first|last|nth|nth-last)-(child|of-type)(?:\\(" +
			whitespace + "*(even|odd|(([+-]|)(\\d*)n|)" + whitespace + "*(?:([+-]|)" +
			whitespace + "*(\\d+)|))" + whitespace + "*\\)|)", "i" ),
		"bool": new RegExp( "^(?:" + booleans + ")$", "i" ),

		// For use in libraries implementing .is()
		// We use this for POS matching in `select`
		"needsContext": new RegExp( "^" + whitespace +
			"*[>+~]|:(even|odd|eq|gt|lt|nth|first|last)(?:\\(" + whitespace +
			"*((?:-\\d)?\\d*)" + whitespace + "*\\)|)(?=[^-]|$)", "i" )
	},

	rhtml = /HTML$/i,
	rinputs = /^(?:input|select|textarea|button)$/i,
	rheader = /^h\d$/i,

	rnative = /^[^{]+\{\s*\[native \w/,

	// Easily-parseable/retrievable ID or TAG or CLASS selectors
	rquickExpr = /^(?:#([\w-]+)|(\w+)|\.([\w-]+))$/,

	rsibling = /[+~]/,

	// CSS escapes
	// http://www.w3.org/TR/CSS21/syndata.html#escaped-characters
	runescape = new RegExp( "\\\\[\\da-fA-F]{1,6}" + whitespace + "?|\\\\([^\\r\\n\\f])", "g" ),
	funescape = function( escape, nonHex ) {
		var high = "0x" + escape.slice( 1 ) - 0x10000;

		return nonHex ?

			// Strip the backslash prefix from a non-hex escape sequence
			nonHex :

			// Replace a hexadecimal escape sequence with the encoded Unicode code point
			// Support: IE <=11+
			// For values outside the Basic Multilingual Plane (BMP), manually construct a
			// surrogate pair
			high < 0 ?
				String.fromCharCode( high + 0x10000 ) :
				String.fromCharCode( high >> 10 | 0xD800, high & 0x3FF | 0xDC00 );
	},

	// CSS string/identifier serialization
	// https://drafts.csswg.org/cssom/#common-serializing-idioms
	rcssescape = /([\0-\x1f\x7f]|^-?\d)|^-$|[^\0-\x1f\x7f-\uFFFF\w-]/g,
	fcssescape = function( ch, asCodePoint ) {
		if ( asCodePoint ) {

			// U+0000 NULL becomes U+FFFD REPLACEMENT CHARACTER
			if ( ch === "\0" ) {
				return "\uFFFD";
			}

			// Control characters and (dependent upon position) numbers get escaped as code points
			return ch.slice( 0, -1 ) + "\\" +
				ch.charCodeAt( ch.length - 1 ).toString( 16 ) + " ";
		}

		// Other potentially-special ASCII characters get backslash-escaped
		return "\\" + ch;
	},

	// Used for iframes
	// See setDocument()
	// Removing the function wrapper causes a "Permission Denied"
	// error in IE
	unloadHandler = function() {
		setDocument();
	},

	inDisabledFieldset = addCombinator(
		function( elem ) {
			return elem.disabled === true && elem.nodeName.toLowerCase() === "fieldset";
		},
		{ dir: "parentNode", next: "legend" }
	);

// Optimize for push.apply( _, NodeList )
try {
	push.apply(
		( arr = slice.call( preferredDoc.childNodes ) ),
		preferredDoc.childNodes
	);

	// Support: Android<4.0
	// Detect silently failing push.apply
	// eslint-disable-next-line no-unused-expressions
	arr[ preferredDoc.childNodes.length ].nodeType;
} catch ( e ) {
	push = { apply: arr.length ?

		// Leverage slice if possible
		function( target, els ) {
			pushNative.apply( target, slice.call( els ) );
		} :

		// Support: IE<9
		// Otherwise append directly
		function( target, els ) {
			var j = target.length,
				i = 0;

			// Can't trust NodeList.length
			while ( ( target[ j++ ] = els[ i++ ] ) ) {}
			target.length = j - 1;
		}
	};
}

function Sizzle( selector, context, results, seed ) {
	var m, i, elem, nid, match, groups, newSelector,
		newContext = context && context.ownerDocument,

		// nodeType defaults to 9, since context defaults to document
		nodeType = context ? context.nodeType : 9;

	results = results || [];

	// Return early from calls with invalid selector or context
	if ( typeof selector !== "string" || !selector ||
		nodeType !== 1 && nodeType !== 9 && nodeType !== 11 ) {

		return results;
	}

	// Try to shortcut find operations (as opposed to filters) in HTML documents
	if ( !seed ) {
		setDocument( context );
		context = context || document;

		if ( documentIsHTML ) {

			// If the selector is sufficiently simple, try using a "get*By*" DOM method
			// (excepting DocumentFragment context, where the methods don't exist)
			if ( nodeType !== 11 && ( match = rquickExpr.exec( selector ) ) ) {

				// ID selector
				if ( ( m = match[ 1 ] ) ) {

					// Document context
					if ( nodeType === 9 ) {
						if ( ( elem = context.getElementById( m ) ) ) {

							// Support: IE, Opera, Webkit
							// TODO: identify versions
							// getElementById can match elements by name instead of ID
							if ( elem.id === m ) {
								results.push( elem );
								return results;
							}
						} else {
							return results;
						}

					// Element context
					} else {

						// Support: IE, Opera, Webkit
						// TODO: identify versions
						// getElementById can match elements by name instead of ID
						if ( newContext && ( elem = newContext.getElementById( m ) ) &&
							contains( context, elem ) &&
							elem.id === m ) {

							results.push( elem );
							return results;
						}
					}

				// Type selector
				} else if ( match[ 2 ] ) {
					push.apply( results, context.getElementsByTagName( selector ) );
					return results;

				// Class selector
				} else if ( ( m = match[ 3 ] ) && support.getElementsByClassName &&
					context.getElementsByClassName ) {

					push.apply( results, context.getElementsByClassName( m ) );
					return results;
				}
			}

			// Take advantage of querySelectorAll
			if ( support.qsa &&
				!nonnativeSelectorCache[ selector + " " ] &&
				( !rbuggyQSA || !rbuggyQSA.test( selector ) ) &&

				// Support: IE 8 only
				// Exclude object elements
				( nodeType !== 1 || context.nodeName.toLowerCase() !== "object" ) ) {

				newSelector = selector;
				newContext = context;

				// qSA considers elements outside a scoping root when evaluating child or
				// descendant combinators, which is not what we want.
				// In such cases, we work around the behavior by prefixing every selector in the
				// list with an ID selector referencing the scope context.
				// The technique has to be used as well when a leading combinator is used
				// as such selectors are not recognized by querySelectorAll.
				// Thanks to Andrew Dupont for this technique.
				if ( nodeType === 1 &&
					( rdescend.test( selector ) || rcombinators.test( selector ) ) ) {

					// Expand context for sibling selectors
					newContext = rsibling.test( selector ) && testContext( context.parentNode ) ||
						context;

					// We can use :scope instead of the ID hack if the browser
					// supports it & if we're not changing the context.
					if ( newContext !== context || !support.scope ) {

						// Capture the context ID, setting it first if necessary
						if ( ( nid = context.getAttribute( "id" ) ) ) {
							nid = nid.replace( rcssescape, fcssescape );
						} else {
							context.setAttribute( "id", ( nid = expando ) );
						}
					}

					// Prefix every selector in the list
					groups = tokenize( selector );
					i = groups.length;
					while ( i-- ) {
						groups[ i ] = ( nid ? "#" + nid : ":scope" ) + " " +
							toSelector( groups[ i ] );
					}
					newSelector = groups.join( "," );
				}

				try {
					push.apply( results,
						newContext.querySelectorAll( newSelector )
					);
					return results;
				} catch ( qsaError ) {
					nonnativeSelectorCache( selector, true );
				} finally {
					if ( nid === expando ) {
						context.removeAttribute( "id" );
					}
				}
			}
		}
	}

	// All others
	return select( selector.replace( rtrim, "$1" ), context, results, seed );
}

/**
 * Create key-value caches of limited size
 * @returns {function(string, object)} Returns the Object data after storing it on itself with
 *	property name the (space-suffixed) string and (if the cache is larger than Expr.cacheLength)
 *	deleting the oldest entry
 */
function createCache() {
	var keys = [];

	function cache( key, value ) {

		// Use (key + " ") to avoid collision with native prototype properties (see Issue #157)
		if ( keys.push( key + " " ) > Expr.cacheLength ) {

			// Only keep the most recent entries
			delete cache[ keys.shift() ];
		}
		return ( cache[ key + " " ] = value );
	}
	return cache;
}

/**
 * Mark a function for special use by Sizzle
 * @param {Function} fn The function to mark
 */
function markFunction( fn ) {
	fn[ expando ] = true;
	return fn;
}

/**
 * Support testing using an element
 * @param {Function} fn Passed the created element and returns a boolean result
 */
function assert( fn ) {
	var el = document.createElement( "fieldset" );

	try {
		return !!fn( el );
	} catch ( e ) {
		return false;
	} finally {

		// Remove from its parent by default
		if ( el.parentNode ) {
			el.parentNode.removeChild( el );
		}

		// release memory in IE
		el = null;
	}
}

/**
 * Adds the same handler for all of the specified attrs
 * @param {String} attrs Pipe-separated list of attributes
 * @param {Function} handler The method that will be applied
 */
function addHandle( attrs, handler ) {
	var arr = attrs.split( "|" ),
		i = arr.length;

	while ( i-- ) {
		Expr.attrHandle[ arr[ i ] ] = handler;
	}
}

/**
 * Checks document order of two siblings
 * @param {Element} a
 * @param {Element} b
 * @returns {Number} Returns less than 0 if a precedes b, greater than 0 if a follows b
 */
function siblingCheck( a, b ) {
	var cur = b && a,
		diff = cur && a.nodeType === 1 && b.nodeType === 1 &&
			a.sourceIndex - b.sourceIndex;

	// Use IE sourceIndex if available on both nodes
	if ( diff ) {
		return diff;
	}

	// Check if b follows a
	if ( cur ) {
		while ( ( cur = cur.nextSibling ) ) {
			if ( cur === b ) {
				return -1;
			}
		}
	}

	return a ? 1 : -1;
}

/**
 * Returns a function to use in pseudos for input types
 * @param {String} type
 */
function createInputPseudo( type ) {
	return function( elem ) {
		var name = elem.nodeName.toLowerCase();
		return name === "input" && elem.type === type;
	};
}

/**
 * Returns a function to use in pseudos for buttons
 * @param {String} type
 */
function createButtonPseudo( type ) {
	return function( elem ) {
		var name = elem.nodeName.toLowerCase();
		return ( name === "input" || name === "button" ) && elem.type === type;
	};
}

/**
 * Returns a function to use in pseudos for :enabled/:disabled
 * @param {Boolean} disabled true for :disabled; false for :enabled
 */
function createDisabledPseudo( disabled ) {

	// Known :disabled false positives: fieldset[disabled] > legend:nth-of-type(n+2) :can-disable
	return function( elem ) {

		// Only certain elements can match :enabled or :disabled
		// https://html.spec.whatwg.org/multipage/scripting.html#selector-enabled
		// https://html.spec.whatwg.org/multipage/scripting.html#selector-disabled
		if ( "form" in elem ) {

			// Check for inherited disabledness on relevant non-disabled elements:
			// * listed form-associated elements in a disabled fieldset
			//   https://html.spec.whatwg.org/multipage/forms.html#category-listed
			//   https://html.spec.whatwg.org/multipage/forms.html#concept-fe-disabled
			// * option elements in a disabled optgroup
			//   https://html.spec.whatwg.org/multipage/forms.html#concept-option-disabled
			// All such elements have a "form" property.
			if ( elem.parentNode && elem.disabled === false ) {

				// Option elements defer to a parent optgroup if present
				if ( "label" in elem ) {
					if ( "label" in elem.parentNode ) {
						return elem.parentNode.disabled === disabled;
					} else {
						return elem.disabled === disabled;
					}
				}

				// Support: IE 6 - 11
				// Use the isDisabled shortcut property to check for disabled fieldset ancestors
				return elem.isDisabled === disabled ||

					// Where there is no isDisabled, check manually
					/* jshint -W018 */
					elem.isDisabled !== !disabled &&
					inDisabledFieldset( elem ) === disabled;
			}

			return elem.disabled === disabled;

		// Try to winnow out elements that can't be disabled before trusting the disabled property.
		// Some victims get caught in our net (label, legend, menu, track), but it shouldn't
		// even exist on them, let alone have a boolean value.
		} else if ( "label" in elem ) {
			return elem.disabled === disabled;
		}

		// Remaining elements are neither :enabled nor :disabled
		return false;
	};
}

/**
 * Returns a function to use in pseudos for positionals
 * @param {Function} fn
 */
function createPositionalPseudo( fn ) {
	return markFunction( function( argument ) {
		argument = +argument;
		return markFunction( function( seed, matches ) {
			var j,
				matchIndexes = fn( [], seed.length, argument ),
				i = matchIndexes.length;

			// Match elements found at the specified indexes
			while ( i-- ) {
				if ( seed[ ( j = matchIndexes[ i ] ) ] ) {
					seed[ j ] = !( matches[ j ] = seed[ j ] );
				}
			}
		} );
	} );
}

/**
 * Checks a node for validity as a Sizzle context
 * @param {Element|Object=} context
 * @returns {Element|Object|Boolean} The input node if acceptable, otherwise a falsy value
 */
function testContext( context ) {
	return context && typeof context.getElementsByTagName !== "undefined" && context;
}

// Expose support vars for convenience
support = Sizzle.support = {};

/**
 * Detects XML nodes
 * @param {Element|Object} elem An element or a document
 * @returns {Boolean} True iff elem is a non-HTML XML node
 */
isXML = Sizzle.isXML = function( elem ) {
	var namespace = elem && elem.namespaceURI,
		docElem = elem && ( elem.ownerDocument || elem ).documentElement;

	// Support: IE <=8
	// Assume HTML when documentElement doesn't yet exist, such as inside loading iframes
	// https://bugs.jquery.com/ticket/4833
	return !rhtml.test( namespace || docElem && docElem.nodeName || "HTML" );
};

/**
 * Sets document-related variables once based on the current document
 * @param {Element|Object} [doc] An element or document object to use to set the document
 * @returns {Object} Returns the current document
 */
setDocument = Sizzle.setDocument = function( node ) {
	var hasCompare, subWindow,
		doc = node ? node.ownerDocument || node : preferredDoc;

	// Return early if doc is invalid or already selected
	// Support: IE 11+, Edge 17 - 18+
	// IE/Edge sometimes throw a "Permission denied" error when strict-comparing
	// two documents; shallow comparisons work.
	// eslint-disable-next-line eqeqeq
	if ( doc == document || doc.nodeType !== 9 || !doc.documentElement ) {
		return document;
	}

	// Update global variables
	document = doc;
	docElem = document.documentElement;
	documentIsHTML = !isXML( document );

	// Support: IE 9 - 11+, Edge 12 - 18+
	// Accessing iframe documents after unload throws "permission denied" errors (jQuery #13936)
	// Support: IE 11+, Edge 17 - 18+
	// IE/Edge sometimes throw a "Permission denied" error when strict-comparing
	// two documents; shallow comparisons work.
	// eslint-disable-next-line eqeqeq
	if ( preferredDoc != document &&
		( subWindow = document.defaultView ) && subWindow.top !== subWindow ) {

		// Support: IE 11, Edge
		if ( subWindow.addEventListener ) {
			subWindow.addEventListener( "unload", unloadHandler, false );

		// Support: IE 9 - 10 only
		} else if ( subWindow.attachEvent ) {
			subWindow.attachEvent( "onunload", unloadHandler );
		}
	}

	// Support: IE 8 - 11+, Edge 12 - 18+, Chrome <=16 - 25 only, Firefox <=3.6 - 31 only,
	// Safari 4 - 5 only, Opera <=11.6 - 12.x only
	// IE/Edge & older browsers don't support the :scope pseudo-class.
	// Support: Safari 6.0 only
	// Safari 6.0 supports :scope but it's an alias of :root there.
	support.scope = assert( function( el ) {
		docElem.appendChild( el ).appendChild( document.createElement( "div" ) );
		return typeof el.querySelectorAll !== "undefined" &&
			!el.querySelectorAll( ":scope fieldset div" ).length;
	} );

	/* Attributes
	---------------------------------------------------------------------- */

	// Support: IE<8
	// Verify that getAttribute really returns attributes and not properties
	// (excepting IE8 booleans)
	support.attributes = assert( function( el ) {
		el.className = "i";
		return !el.getAttribute( "className" );
	} );

	/* getElement(s)By*
	---------------------------------------------------------------------- */

	// Check if getElementsByTagName("*") returns only elements
	support.getElementsByTagName = assert( function( el ) {
		el.appendChild( document.createComment( "" ) );
		return !el.getElementsByTagName( "*" ).length;
	} );

	// Support: IE<9
	support.getElementsByClassName = rnative.test( document.getElementsByClassName );

	// Support: IE<10
	// Check if getElementById returns elements by name
	// The broken getElementById methods don't pick up programmatically-set names,
	// so use a roundabout getElementsByName test
	support.getById = assert( function( el ) {
		docElem.appendChild( el ).id = expando;
		return !document.getElementsByName || !document.getElementsByName( expando ).length;
	} );

	// ID filter and find
	if ( support.getById ) {
		Expr.filter[ "ID" ] = function( id ) {
			var attrId = id.replace( runescape, funescape );
			return function( elem ) {
				return elem.getAttribute( "id" ) === attrId;
			};
		};
		Expr.find[ "ID" ] = function( id, context ) {
			if ( typeof context.getElementById !== "undefined" && documentIsHTML ) {
				var elem = context.getElementById( id );
				return elem ? [ elem ] : [];
			}
		};
	} else {
		Expr.filter[ "ID" ] =  function( id ) {
			var attrId = id.replace( runescape, funescape );
			return function( elem ) {
				var node = typeof elem.getAttributeNode !== "undefined" &&
					elem.getAttributeNode( "id" );
				return node && node.value === attrId;
			};
		};

		// Support: IE 6 - 7 only
		// getElementById is not reliable as a find shortcut
		Expr.find[ "ID" ] = function( id, context ) {
			if ( typeof context.getElementById !== "undefined" && documentIsHTML ) {
				var node, i, elems,
					elem = context.getElementById( id );

				if ( elem ) {

					// Verify the id attribute
					node = elem.getAttributeNode( "id" );
					if ( node && node.value === id ) {
						return [ elem ];
					}

					// Fall back on getElementsByName
					elems = context.getElementsByName( id );
					i = 0;
					while ( ( elem = elems[ i++ ] ) ) {
						node = elem.getAttributeNode( "id" );
						if ( node && node.value === id ) {
							return [ elem ];
						}
					}
				}

				return [];
			}
		};
	}

	// Tag
	Expr.find[ "TAG" ] = support.getElementsByTagName ?
		function( tag, context ) {
			if ( typeof context.getElementsByTagName !== "undefined" ) {
				return context.getElementsByTagName( tag );

			// DocumentFragment nodes don't have gEBTN
			} else if ( support.qsa ) {
				return context.querySelectorAll( tag );
			}
		} :

		function( tag, context ) {
			var elem,
				tmp = [],
				i = 0,

				// By happy coincidence, a (broken) gEBTN appears on DocumentFragment nodes too
				results = context.getElementsByTagName( tag );

			// Filter out possible comments
			if ( tag === "*" ) {
				while ( ( elem = results[ i++ ] ) ) {
					if ( elem.nodeType === 1 ) {
						tmp.push( elem );
					}
				}

				return tmp;
			}
			return results;
		};

	// Class
	Expr.find[ "CLASS" ] = support.getElementsByClassName && function( className, context ) {
		if ( typeof context.getElementsByClassName !== "undefined" && documentIsHTML ) {
			return context.getElementsByClassName( className );
		}
	};

	/* QSA/matchesSelector
	---------------------------------------------------------------------- */

	// QSA and matchesSelector support

	// matchesSelector(:active) reports false when true (IE9/Opera 11.5)
	rbuggyMatches = [];

	// qSa(:focus) reports false when true (Chrome 21)
	// We allow this because of a bug in IE8/9 that throws an error
	// whenever `document.activeElement` is accessed on an iframe
	// So, we allow :focus to pass through QSA all the time to avoid the IE error
	// See https://bugs.jquery.com/ticket/13378
	rbuggyQSA = [];

	if ( ( support.qsa = rnative.test( document.querySelectorAll ) ) ) {

		// Build QSA regex
		// Regex strategy adopted from Diego Perini
		assert( function( el ) {

			var input;

			// Select is set to empty string on purpose
			// This is to test IE's treatment of not explicitly
			// setting a boolean content attribute,
			// since its presence should be enough
			// https://bugs.jquery.com/ticket/12359
			docElem.appendChild( el ).innerHTML = "<a id='" + expando + "'></a>" +
				"<select id='" + expando + "-\r\\' msallowcapture=''>" +
				"<option selected=''></option></select>";

			// Support: IE8, Opera 11-12.16
			// Nothing should be selected when empty strings follow ^= or $= or *=
			// The test attribute must be unknown in Opera but "safe" for WinRT
			// https://msdn.microsoft.com/en-us/library/ie/hh465388.aspx#attribute_section
			if ( el.querySelectorAll( "[msallowcapture^='']" ).length ) {
				rbuggyQSA.push( "[*^$]=" + whitespace + "*(?:''|\"\")" );
			}

			// Support: IE8
			// Boolean attributes and "value" are not treated correctly
			if ( !el.querySelectorAll( "[selected]" ).length ) {
				rbuggyQSA.push( "\\[" + whitespace + "*(?:value|" + booleans + ")" );
			}

			// Support: Chrome<29, Android<4.4, Safari<7.0+, iOS<7.0+, PhantomJS<1.9.8+
			if ( !el.querySelectorAll( "[id~=" + expando + "-]" ).length ) {
				rbuggyQSA.push( "~=" );
			}

			// Support: IE 11+, Edge 15 - 18+
			// IE 11/Edge don't find elements on a `[name='']` query in some cases.
			// Adding a temporary attribute to the document before the selection works
			// around the issue.
			// Interestingly, IE 10 & older don't seem to have the issue.
			input = document.createElement( "input" );
			input.setAttribute( "name", "" );
			el.appendChild( input );
			if ( !el.querySelectorAll( "[name='']" ).length ) {
				rbuggyQSA.push( "\\[" + whitespace + "*name" + whitespace + "*=" +
					whitespace + "*(?:''|\"\")" );
			}

			// Webkit/Opera - :checked should return selected option elements
			// http://www.w3.org/TR/2011/REC-css3-selectors-20110929/#checked
			// IE8 throws error here and will not see later tests
			if ( !el.querySelectorAll( ":checked" ).length ) {
				rbuggyQSA.push( ":checked" );
			}

			// Support: Safari 8+, iOS 8+
			// https://bugs.webkit.org/show_bug.cgi?id=136851
			// In-page `selector#id sibling-combinator selector` fails
			if ( !el.querySelectorAll( "a#" + expando + "+*" ).length ) {
				rbuggyQSA.push( ".#.+[+~]" );
			}

			// Support: Firefox <=3.6 - 5 only
			// Old Firefox doesn't throw on a badly-escaped identifier.
			el.querySelectorAll( "\\\f" );
			rbuggyQSA.push( "[\\r\\n\\f]" );
		} );

		assert( function( el ) {
			el.innerHTML = "<a href='' disabled='disabled'></a>" +
				"<select disabled='disabled'><option/></select>";

			// Support: Windows 8 Native Apps
			// The type and name attributes are restricted during .innerHTML assignment
			var input = document.createElement( "input" );
			input.setAttribute( "type", "hidden" );
			el.appendChild( input ).setAttribute( "name", "D" );

			// Support: IE8
			// Enforce case-sensitivity of name attribute
			if ( el.querySelectorAll( "[name=d]" ).length ) {
				rbuggyQSA.push( "name" + whitespace + "*[*^$|!~]?=" );
			}

			// FF 3.5 - :enabled/:disabled and hidden elements (hidden elements are still enabled)
			// IE8 throws error here and will not see later tests
			if ( el.querySelectorAll( ":enabled" ).length !== 2 ) {
				rbuggyQSA.push( ":enabled", ":disabled" );
			}

			// Support: IE9-11+
			// IE's :disabled selector does not pick up the children of disabled fieldsets
			docElem.appendChild( el ).disabled = true;
			if ( el.querySelectorAll( ":disabled" ).length !== 2 ) {
				rbuggyQSA.push( ":enabled", ":disabled" );
			}

			// Support: Opera 10 - 11 only
			// Opera 10-11 does not throw on post-comma invalid pseudos
			el.querySelectorAll( "*,:x" );
			rbuggyQSA.push( ",.*:" );
		} );
	}

	if ( ( support.matchesSelector = rnative.test( ( matches = docElem.matches ||
		docElem.webkitMatchesSelector ||
		docElem.mozMatchesSelector ||
		docElem.oMatchesSelector ||
		docElem.msMatchesSelector ) ) ) ) {

		assert( function( el ) {

			// Check to see if it's possible to do matchesSelector
			// on a disconnected node (IE 9)
			support.disconnectedMatch = matches.call( el, "*" );

			// This should fail with an exception
			// Gecko does not error, returns false instead
			matches.call( el, "[s!='']:x" );
			rbuggyMatches.push( "!=", pseudos );
		} );
	}

	rbuggyQSA = rbuggyQSA.length && new RegExp( rbuggyQSA.join( "|" ) );
	rbuggyMatches = rbuggyMatches.length && new RegExp( rbuggyMatches.join( "|" ) );

	/* Contains
	---------------------------------------------------------------------- */
	hasCompare = rnative.test( docElem.compareDocumentPosition );

	// Element contains another
	// Purposefully self-exclusive
	// As in, an element does not contain itself
	contains = hasCompare || rnative.test( docElem.contains ) ?
		function( a, b ) {
			var adown = a.nodeType === 9 ? a.documentElement : a,
				bup = b && b.parentNode;
			return a === bup || !!( bup && bup.nodeType === 1 && (
				adown.contains ?
					adown.contains( bup ) :
					a.compareDocumentPosition && a.compareDocumentPosition( bup ) & 16
			) );
		} :
		function( a, b ) {
			if ( b ) {
				while ( ( b = b.parentNode ) ) {
					if ( b === a ) {
						return true;
					}
				}
			}
			return false;
		};

	/* Sorting
	---------------------------------------------------------------------- */

	// Document order sorting
	sortOrder = hasCompare ?
	function( a, b ) {

		// Flag for duplicate removal
		if ( a === b ) {
			hasDuplicate = true;
			return 0;
		}

		// Sort on method existence if only one input has compareDocumentPosition
		var compare = !a.compareDocumentPosition - !b.compareDocumentPosition;
		if ( compare ) {
			return compare;
		}

		// Calculate position if both inputs belong to the same document
		// Support: IE 11+, Edge 17 - 18+
		// IE/Edge sometimes throw a "Permission denied" error when strict-comparing
		// two documents; shallow comparisons work.
		// eslint-disable-next-line eqeqeq
		compare = ( a.ownerDocument || a ) == ( b.ownerDocument || b ) ?
			a.compareDocumentPosition( b ) :

			// Otherwise we know they are disconnected
			1;

		// Disconnected nodes
		if ( compare & 1 ||
			( !support.sortDetached && b.compareDocumentPosition( a ) === compare ) ) {

			// Choose the first element that is related to our preferred document
			// Support: IE 11+, Edge 17 - 18+
			// IE/Edge sometimes throw a "Permission denied" error when strict-comparing
			// two documents; shallow comparisons work.
			// eslint-disable-next-line eqeqeq
			if ( a == document || a.ownerDocument == preferredDoc &&
				contains( preferredDoc, a ) ) {
				return -1;
			}

			// Support: IE 11+, Edge 17 - 18+
			// IE/Edge sometimes throw a "Permission denied" error when strict-comparing
			// two documents; shallow comparisons work.
			// eslint-disable-next-line eqeqeq
			if ( b == document || b.ownerDocument == preferredDoc &&
				contains( preferredDoc, b ) ) {
				return 1;
			}

			// Maintain original order
			return sortInput ?
				( indexOf( sortInput, a ) - indexOf( sortInput, b ) ) :
				0;
		}

		return compare & 4 ? -1 : 1;
	} :
	function( a, b ) {

		// Exit early if the nodes are identical
		if ( a === b ) {
			hasDuplicate = true;
			return 0;
		}

		var cur,
			i = 0,
			aup = a.parentNode,
			bup = b.parentNode,
			ap = [ a ],
			bp = [ b ];

		// Parentless nodes are either documents or disconnected
		if ( !aup || !bup ) {

			// Support: IE 11+, Edge 17 - 18+
			// IE/Edge sometimes throw a "Permission denied" error when strict-comparing
			// two documents; shallow comparisons work.
			/* eslint-disable eqeqeq */
			return a == document ? -1 :
				b == document ? 1 :
				/* eslint-enable eqeqeq */
				aup ? -1 :
				bup ? 1 :
				sortInput ?
				( indexOf( sortInput, a ) - indexOf( sortInput, b ) ) :
				0;

		// If the nodes are siblings, we can do a quick check
		} else if ( aup === bup ) {
			return siblingCheck( a, b );
		}

		// Otherwise we need full lists of their ancestors for comparison
		cur = a;
		while ( ( cur = cur.parentNode ) ) {
			ap.unshift( cur );
		}
		cur = b;
		while ( ( cur = cur.parentNode ) ) {
			bp.unshift( cur );
		}

		// Walk down the tree looking for a discrepancy
		while ( ap[ i ] === bp[ i ] ) {
			i++;
		}

		return i ?

			// Do a sibling check if the nodes have a common ancestor
			siblingCheck( ap[ i ], bp[ i ] ) :

			// Otherwise nodes in our document sort first
			// Support: IE 11+, Edge 17 - 18+
			// IE/Edge sometimes throw a "Permission denied" error when strict-comparing
			// two documents; shallow comparisons work.
			/* eslint-disable eqeqeq */
			ap[ i ] == preferredDoc ? -1 :
			bp[ i ] == preferredDoc ? 1 :
			/* eslint-enable eqeqeq */
			0;
	};

	return document;
};

Sizzle.matches = function( expr, elements ) {
	return Sizzle( expr, null, null, elements );
};

Sizzle.matchesSelector = function( elem, expr ) {
	setDocument( elem );

	if ( support.matchesSelector && documentIsHTML &&
		!nonnativeSelectorCache[ expr + " " ] &&
		( !rbuggyMatches || !rbuggyMatches.test( expr ) ) &&
		( !rbuggyQSA     || !rbuggyQSA.test( expr ) ) ) {

		try {
			var ret = matches.call( elem, expr );

			// IE 9's matchesSelector returns false on disconnected nodes
			if ( ret || support.disconnectedMatch ||

				// As well, disconnected nodes are said to be in a document
				// fragment in IE 9
				elem.document && elem.document.nodeType !== 11 ) {
				return ret;
			}
		} catch ( e ) {
			nonnativeSelectorCache( expr, true );
		}
	}

	return Sizzle( expr, document, null, [ elem ] ).length > 0;
};

Sizzle.contains = function( context, elem ) {

	// Set document vars if needed
	// Support: IE 11+, Edge 17 - 18+
	// IE/Edge sometimes throw a "Permission denied" error when strict-comparing
	// two documents; shallow comparisons work.
	// eslint-disable-next-line eqeqeq
	if ( ( context.ownerDocument || context ) != document ) {
		setDocument( context );
	}
	return contains( context, elem );
};

Sizzle.attr = function( elem, name ) {

	// Set document vars if needed
	// Support: IE 11+, Edge 17 - 18+
	// IE/Edge sometimes throw a "Permission denied" error when strict-comparing
	// two documents; shallow comparisons work.
	// eslint-disable-next-line eqeqeq
	if ( ( elem.ownerDocument || elem ) != document ) {
		setDocument( elem );
	}

	var fn = Expr.attrHandle[ name.toLowerCase() ],

		// Don't get fooled by Object.prototype properties (jQuery #13807)
		val = fn && hasOwn.call( Expr.attrHandle, name.toLowerCase() ) ?
			fn( elem, name, !documentIsHTML ) :
			undefined;

	return val !== undefined ?
		val :
		support.attributes || !documentIsHTML ?
			elem.getAttribute( name ) :
			( val = elem.getAttributeNode( name ) ) && val.specified ?
				val.value :
				null;
};

Sizzle.escape = function( sel ) {
	return ( sel + "" ).replace( rcssescape, fcssescape );
};

Sizzle.error = function( msg ) {
	throw new Error( "Syntax error, unrecognized expression: " + msg );
};

/**
 * Document sorting and removing duplicates
 * @param {ArrayLike} results
 */
Sizzle.uniqueSort = function( results ) {
	var elem,
		duplicates = [],
		j = 0,
		i = 0;

	// Unless we *know* we can detect duplicates, assume their presence
	hasDuplicate = !support.detectDuplicates;
	sortInput = !support.sortStable && results.slice( 0 );
	results.sort( sortOrder );

	if ( hasDuplicate ) {
		while ( ( elem = results[ i++ ] ) ) {
			if ( elem === results[ i ] ) {
				j = duplicates.push( i );
			}
		}
		while ( j-- ) {
			results.splice( duplicates[ j ], 1 );
		}
	}

	// Clear input after sorting to release objects
	// See https://github.com/jquery/sizzle/pull/225
	sortInput = null;

	return results;
};

/**
 * Utility function for retrieving the text value of an array of DOM nodes
 * @param {Array|Element} elem
 */
getText = Sizzle.getText = function( elem ) {
	var node,
		ret = "",
		i = 0,
		nodeType = elem.nodeType;

	if ( !nodeType ) {

		// If no nodeType, this is expected to be an array
		while ( ( node = elem[ i++ ] ) ) {

			// Do not traverse comment nodes
			ret += getText( node );
		}
	} else if ( nodeType === 1 || nodeType === 9 || nodeType === 11 ) {

		// Use textContent for elements
		// innerText usage removed for consistency of new lines (jQuery #11153)
		if ( typeof elem.textContent === "string" ) {
			return elem.textContent;
		} else {

			// Traverse its children
			for ( elem = elem.firstChild; elem; elem = elem.nextSibling ) {
				ret += getText( elem );
			}
		}
	} else if ( nodeType === 3 || nodeType === 4 ) {
		return elem.nodeValue;
	}

	// Do not include comment or processing instruction nodes

	return ret;
};

Expr = Sizzle.selectors = {

	// Can be adjusted by the user
	cacheLength: 50,

	createPseudo: markFunction,

	match: matchExpr,

	attrHandle: {},

	find: {},

	relative: {
		">": { dir: "parentNode", first: true },
		" ": { dir: "parentNode" },
		"+": { dir: "previousSibling", first: true },
		"~": { dir: "previousSibling" }
	},

	preFilter: {
		"ATTR": function( match ) {
			match[ 1 ] = match[ 1 ].replace( runescape, funescape );

			// Move the given value to match[3] whether quoted or unquoted
			match[ 3 ] = ( match[ 3 ] || match[ 4 ] ||
				match[ 5 ] || "" ).replace( runescape, funescape );

			if ( match[ 2 ] === "~=" ) {
				match[ 3 ] = " " + match[ 3 ] + " ";
			}

			return match.slice( 0, 4 );
		},

		"CHILD": function( match ) {

			/* matches from matchExpr["CHILD"]
				1 type (only|nth|...)
				2 what (child|of-type)
				3 argument (even|odd|\d*|\d*n([+-]\d+)?|...)
				4 xn-component of xn+y argument ([+-]?\d*n|)
				5 sign of xn-component
				6 x of xn-component
				7 sign of y-component
				8 y of y-component
			*/
			match[ 1 ] = match[ 1 ].toLowerCase();

			if ( match[ 1 ].slice( 0, 3 ) === "nth" ) {

				// nth-* requires argument
				if ( !match[ 3 ] ) {
					Sizzle.error( match[ 0 ] );
				}

				// numeric x and y parameters for Expr.filter.CHILD
				// remember that false/true cast respectively to 0/1
				match[ 4 ] = +( match[ 4 ] ?
					match[ 5 ] + ( match[ 6 ] || 1 ) :
					2 * ( match[ 3 ] === "even" || match[ 3 ] === "odd" ) );
				match[ 5 ] = +( ( match[ 7 ] + match[ 8 ] ) || match[ 3 ] === "odd" );

				// other types prohibit arguments
			} else if ( match[ 3 ] ) {
				Sizzle.error( match[ 0 ] );
			}

			return match;
		},

		"PSEUDO": function( match ) {
			var excess,
				unquoted = !match[ 6 ] && match[ 2 ];

			if ( matchExpr[ "CHILD" ].test( match[ 0 ] ) ) {
				return null;
			}

			// Accept quoted arguments as-is
			if ( match[ 3 ] ) {
				match[ 2 ] = match[ 4 ] || match[ 5 ] || "";

			// Strip excess characters from unquoted arguments
			} else if ( unquoted && rpseudo.test( unquoted ) &&

				// Get excess from tokenize (recursively)
				( excess = tokenize( unquoted, true ) ) &&

				// advance to the next closing parenthesis
				( excess = unquoted.indexOf( ")", unquoted.length - excess ) - unquoted.length ) ) {

				// excess is a negative index
				match[ 0 ] = match[ 0 ].slice( 0, excess );
				match[ 2 ] = unquoted.slice( 0, excess );
			}

			// Return only captures needed by the pseudo filter method (type and argument)
			return match.slice( 0, 3 );
		}
	},

	filter: {

		"TAG": function( nodeNameSelector ) {
			var nodeName = nodeNameSelector.replace( runescape, funescape ).toLowerCase();
			return nodeNameSelector === "*" ?
				function() {
					return true;
				} :
				function( elem ) {
					return elem.nodeName && elem.nodeName.toLowerCase() === nodeName;
				};
		},

		"CLASS": function( className ) {
			var pattern = classCache[ className + " " ];

			return pattern ||
				( pattern = new RegExp( "(^|" + whitespace +
					")" + className + "(" + whitespace + "|$)" ) ) && classCache(
						className, function( elem ) {
							return pattern.test(
								typeof elem.className === "string" && elem.className ||
								typeof elem.getAttribute !== "undefined" &&
									elem.getAttribute( "class" ) ||
								""
							);
				} );
		},

		"ATTR": function( name, operator, check ) {
			return function( elem ) {
				var result = Sizzle.attr( elem, name );

				if ( result == null ) {
					return operator === "!=";
				}
				if ( !operator ) {
					return true;
				}

				result += "";

				/* eslint-disable max-len */

				return operator === "=" ? result === check :
					operator === "!=" ? result !== check :
					operator === "^=" ? check && result.indexOf( check ) === 0 :
					operator === "*=" ? check && result.indexOf( check ) > -1 :
					operator === "$=" ? check && result.slice( -check.length ) === check :
					operator === "~=" ? ( " " + result.replace( rwhitespace, " " ) + " " ).indexOf( check ) > -1 :
					operator === "|=" ? result === check || result.slice( 0, check.length + 1 ) === check + "-" :
					false;
				/* eslint-enable max-len */

			};
		},

		"CHILD": function( type, what, _argument, first, last ) {
			var simple = type.slice( 0, 3 ) !== "nth",
				forward = type.slice( -4 ) !== "last",
				ofType = what === "of-type";

			return first === 1 && last === 0 ?

				// Shortcut for :nth-*(n)
				function( elem ) {
					return !!elem.parentNode;
				} :

				function( elem, _context, xml ) {
					var cache, uniqueCache, outerCache, node, nodeIndex, start,
						dir = simple !== forward ? "nextSibling" : "previousSibling",
						parent = elem.parentNode,
						name = ofType && elem.nodeName.toLowerCase(),
						useCache = !xml && !ofType,
						diff = false;

					if ( parent ) {

						// :(first|last|only)-(child|of-type)
						if ( simple ) {
							while ( dir ) {
								node = elem;
								while ( ( node = node[ dir ] ) ) {
									if ( ofType ?
										node.nodeName.toLowerCase() === name :
										node.nodeType === 1 ) {

										return false;
									}
								}

								// Reverse direction for :only-* (if we haven't yet done so)
								start = dir = type === "only" && !start && "nextSibling";
							}
							return true;
						}

						start = [ forward ? parent.firstChild : parent.lastChild ];

						// non-xml :nth-child(...) stores cache data on `parent`
						if ( forward && useCache ) {

							// Seek `elem` from a previously-cached index

							// ...in a gzip-friendly way
							node = parent;
							outerCache = node[ expando ] || ( node[ expando ] = {} );

							// Support: IE <9 only
							// Defend against cloned attroperties (jQuery gh-1709)
							uniqueCache = outerCache[ node.uniqueID ] ||
								( outerCache[ node.uniqueID ] = {} );

							cache = uniqueCache[ type ] || [];
							nodeIndex = cache[ 0 ] === dirruns && cache[ 1 ];
							diff = nodeIndex && cache[ 2 ];
							node = nodeIndex && parent.childNodes[ nodeIndex ];

							while ( ( node = ++nodeIndex && node && node[ dir ] ||

								// Fallback to seeking `elem` from the start
								( diff = nodeIndex = 0 ) || start.pop() ) ) {

								// When found, cache indexes on `parent` and break
								if ( node.nodeType === 1 && ++diff && node === elem ) {
									uniqueCache[ type ] = [ dirruns, nodeIndex, diff ];
									break;
								}
							}

						} else {

							// Use previously-cached element index if available
							if ( useCache ) {

								// ...in a gzip-friendly way
								node = elem;
								outerCache = node[ expando ] || ( node[ expando ] = {} );

								// Support: IE <9 only
								// Defend against cloned attroperties (jQuery gh-1709)
								uniqueCache = outerCache[ node.uniqueID ] ||
									( outerCache[ node.uniqueID ] = {} );

								cache = uniqueCache[ type ] || [];
								nodeIndex = cache[ 0 ] === dirruns && cache[ 1 ];
								diff = nodeIndex;
							}

							// xml :nth-child(...)
							// or :nth-last-child(...) or :nth(-last)?-of-type(...)
							if ( diff === false ) {

								// Use the same loop as above to seek `elem` from the start
								while ( ( node = ++nodeIndex && node && node[ dir ] ||
									( diff = nodeIndex = 0 ) || start.pop() ) ) {

									if ( ( ofType ?
										node.nodeName.toLowerCase() === name :
										node.nodeType === 1 ) &&
										++diff ) {

										// Cache the index of each encountered element
										if ( useCache ) {
											outerCache = node[ expando ] ||
												( node[ expando ] = {} );

											// Support: IE <9 only
											// Defend against cloned attroperties (jQuery gh-1709)
											uniqueCache = outerCache[ node.uniqueID ] ||
												( outerCache[ node.uniqueID ] = {} );

											uniqueCache[ type ] = [ dirruns, diff ];
										}

										if ( node === elem ) {
											break;
										}
									}
								}
							}
						}

						// Incorporate the offset, then check against cycle size
						diff -= last;
						return diff === first || ( diff % first === 0 && diff / first >= 0 );
					}
				};
		},

		"PSEUDO": function( pseudo, argument ) {

			// pseudo-class names are case-insensitive
			// http://www.w3.org/TR/selectors/#pseudo-classes
			// Prioritize by case sensitivity in case custom pseudos are added with uppercase letters
			// Remember that setFilters inherits from pseudos
			var args,
				fn = Expr.pseudos[ pseudo ] || Expr.setFilters[ pseudo.toLowerCase() ] ||
					Sizzle.error( "unsupported pseudo: " + pseudo );

			// The user may use createPseudo to indicate that
			// arguments are needed to create the filter function
			// just as Sizzle does
			if ( fn[ expando ] ) {
				return fn( argument );
			}

			// But maintain support for old signatures
			if ( fn.length > 1 ) {
				args = [ pseudo, pseudo, "", argument ];
				return Expr.setFilters.hasOwnProperty( pseudo.toLowerCase() ) ?
					markFunction( function( seed, matches ) {
						var idx,
							matched = fn( seed, argument ),
							i = matched.length;
						while ( i-- ) {
							idx = indexOf( seed, matched[ i ] );
							seed[ idx ] = !( matches[ idx ] = matched[ i ] );
						}
					} ) :
					function( elem ) {
						return fn( elem, 0, args );
					};
			}

			return fn;
		}
	},

	pseudos: {

		// Potentially complex pseudos
		"not": markFunction( function( selector ) {

			// Trim the selector passed to compile
			// to avoid treating leading and trailing
			// spaces as combinators
			var input = [],
				results = [],
				matcher = compile( selector.replace( rtrim, "$1" ) );

			return matcher[ expando ] ?
				markFunction( function( seed, matches, _context, xml ) {
					var elem,
						unmatched = matcher( seed, null, xml, [] ),
						i = seed.length;

					// Match elements unmatched by `matcher`
					while ( i-- ) {
						if ( ( elem = unmatched[ i ] ) ) {
							seed[ i ] = !( matches[ i ] = elem );
						}
					}
				} ) :
				function( elem, _context, xml ) {
					input[ 0 ] = elem;
					matcher( input, null, xml, results );

					// Don't keep the element (issue #299)
					input[ 0 ] = null;
					return !results.pop();
				};
		} ),

		"has": markFunction( function( selector ) {
			return function( elem ) {
				return Sizzle( selector, elem ).length > 0;
			};
		} ),

		"contains": markFunction( function( text ) {
			text = text.replace( runescape, funescape );
			return function( elem ) {
				return ( elem.textContent || getText( elem ) ).indexOf( text ) > -1;
			};
		} ),

		// "Whether an element is represented by a :lang() selector
		// is based solely on the element's language value
		// being equal to the identifier C,
		// or beginning with the identifier C immediately followed by "-".
		// The matching of C against the element's language value is performed case-insensitively.
		// The identifier C does not have to be a valid language name."
		// http://www.w3.org/TR/selectors/#lang-pseudo
		"lang": markFunction( function( lang ) {

			// lang value must be a valid identifier
			if ( !ridentifier.test( lang || "" ) ) {
				Sizzle.error( "unsupported lang: " + lang );
			}
			lang = lang.replace( runescape, funescape ).toLowerCase();
			return function( elem ) {
				var elemLang;
				do {
					if ( ( elemLang = documentIsHTML ?
						elem.lang :
						elem.getAttribute( "xml:lang" ) || elem.getAttribute( "lang" ) ) ) {

						elemLang = elemLang.toLowerCase();
						return elemLang === lang || elemLang.indexOf( lang + "-" ) === 0;
					}
				} while ( ( elem = elem.parentNode ) && elem.nodeType === 1 );
				return false;
			};
		} ),

		// Miscellaneous
		"target": function( elem ) {
			var hash = window.location && window.location.hash;
			return hash && hash.slice( 1 ) === elem.id;
		},

		"root": function( elem ) {
			return elem === docElem;
		},

		"focus": function( elem ) {
			return elem === document.activeElement &&
				( !document.hasFocus || document.hasFocus() ) &&
				!!( elem.type || elem.href || ~elem.tabIndex );
		},

		// Boolean properties
		"enabled": createDisabledPseudo( false ),
		"disabled": createDisabledPseudo( true ),

		"checked": function( elem ) {

			// In CSS3, :checked should return both checked and selected elements
			// http://www.w3.org/TR/2011/REC-css3-selectors-20110929/#checked
			var nodeName = elem.nodeName.toLowerCase();
			return ( nodeName === "input" && !!elem.checked ) ||
				( nodeName === "option" && !!elem.selected );
		},

		"selected": function( elem ) {

			// Accessing this property makes selected-by-default
			// options in Safari work properly
			if ( elem.parentNode ) {
				// eslint-disable-next-line no-unused-expressions
				elem.parentNode.selectedIndex;
			}

			return elem.selected === true;
		},

		// Contents
		"empty": function( elem ) {

			// http://www.w3.org/TR/selectors/#empty-pseudo
			// :empty is negated by element (1) or content nodes (text: 3; cdata: 4; entity ref: 5),
			//   but not by others (comment: 8; processing instruction: 7; etc.)
			// nodeType < 6 works because attributes (2) do not appear as children
			for ( elem = elem.firstChild; elem; elem = elem.nextSibling ) {
				if ( elem.nodeType < 6 ) {
					return false;
				}
			}
			return true;
		},

		"parent": function( elem ) {
			return !Expr.pseudos[ "empty" ]( elem );
		},

		// Element/input types
		"header": function( elem ) {
			return rheader.test( elem.nodeName );
		},

		"input": function( elem ) {
			return rinputs.test( elem.nodeName );
		},

		"button": function( elem ) {
			var name = elem.nodeName.toLowerCase();
			return name === "input" && elem.type === "button" || name === "button";
		},

		"text": function( elem ) {
			var attr;
			return elem.nodeName.toLowerCase() === "input" &&
				elem.type === "text" &&

				// Support: IE<8
				// New HTML5 attribute values (e.g., "search") appear with elem.type === "text"
				( ( attr = elem.getAttribute( "type" ) ) == null ||
					attr.toLowerCase() === "text" );
		},

		// Position-in-collection
		"first": createPositionalPseudo( function() {
			return [ 0 ];
		} ),

		"last": createPositionalPseudo( function( _matchIndexes, length ) {
			return [ length - 1 ];
		} ),

		"eq": createPositionalPseudo( function( _matchIndexes, length, argument ) {
			return [ argument < 0 ? argument + length : argument ];
		} ),

		"even": createPositionalPseudo( function( matchIndexes, length ) {
			var i = 0;
			for ( ; i < length; i += 2 ) {
				matchIndexes.push( i );
			}
			return matchIndexes;
		} ),

		"odd": createPositionalPseudo( function( matchIndexes, length ) {
			var i = 1;
			for ( ; i < length; i += 2 ) {
				matchIndexes.push( i );
			}
			return matchIndexes;
		} ),

		"lt": createPositionalPseudo( function( matchIndexes, length, argument ) {
			var i = argument < 0 ?
				argument + length :
				argument > length ?
					length :
					argument;
			for ( ; --i >= 0; ) {
				matchIndexes.push( i );
			}
			return matchIndexes;
		} ),

		"gt": createPositionalPseudo( function( matchIndexes, length, argument ) {
			var i = argument < 0 ? argument + length : argument;
			for ( ; ++i < length; ) {
				matchIndexes.push( i );
			}
			return matchIndexes;
		} )
	}
};

Expr.pseudos[ "nth" ] = Expr.pseudos[ "eq" ];

// Add button/input type pseudos
for ( i in { radio: true, checkbox: true, file: true, password: true, image: true } ) {
	Expr.pseudos[ i ] = createInputPseudo( i );
}
for ( i in { submit: true, reset: true } ) {
	Expr.pseudos[ i ] = createButtonPseudo( i );
}

// Easy API for creating new setFilters
function setFilters() {}
setFilters.prototype = Expr.filters = Expr.pseudos;
Expr.setFilters = new setFilters();

tokenize = Sizzle.tokenize = function( selector, parseOnly ) {
	var matched, match, tokens, type,
		soFar, groups, preFilters,
		cached = tokenCache[ selector + " " ];

	if ( cached ) {
		return parseOnly ? 0 : cached.slice( 0 );
	}

	soFar = selector;
	groups = [];
	preFilters = Expr.preFilter;

	while ( soFar ) {

		// Comma and first run
		if ( !matched || ( match = rcomma.exec( soFar ) ) ) {
			if ( match ) {

				// Don't consume trailing commas as valid
				soFar = soFar.slice( match[ 0 ].length ) || soFar;
			}
			groups.push( ( tokens = [] ) );
		}

		matched = false;

		// Combinators
		if ( ( match = rcombinators.exec( soFar ) ) ) {
			matched = match.shift();
			tokens.push( {
				value: matched,

				// Cast descendant combinators to space
				type: match[ 0 ].replace( rtrim, " " )
			} );
			soFar = soFar.slice( matched.length );
		}

		// Filters
		for ( type in Expr.filter ) {
			if ( ( match = matchExpr[ type ].exec( soFar ) ) && ( !preFilters[ type ] ||
				( match = preFilters[ type ]( match ) ) ) ) {
				matched = match.shift();
				tokens.push( {
					value: matched,
					type: type,
					matches: match
				} );
				soFar = soFar.slice( matched.length );
			}
		}

		if ( !matched ) {
			break;
		}
	}

	// Return the length of the invalid excess
	// if we're just parsing
	// Otherwise, throw an error or return tokens
	return parseOnly ?
		soFar.length :
		soFar ?
			Sizzle.error( selector ) :

			// Cache the tokens
			tokenCache( selector, groups ).slice( 0 );
};

function toSelector( tokens ) {
	var i = 0,
		len = tokens.length,
		selector = "";
	for ( ; i < len; i++ ) {
		selector += tokens[ i ].value;
	}
	return selector;
}

function addCombinator( matcher, combinator, base ) {
	var dir = combinator.dir,
		skip = combinator.next,
		key = skip || dir,
		checkNonElements = base && key === "parentNode",
		doneName = done++;

	return combinator.first ?

		// Check against closest ancestor/preceding element
		function( elem, context, xml ) {
			while ( ( elem = elem[ dir ] ) ) {
				if ( elem.nodeType === 1 || checkNonElements ) {
					return matcher( elem, context, xml );
				}
			}
			return false;
		} :

		// Check against all ancestor/preceding elements
		function( elem, context, xml ) {
			var oldCache, uniqueCache, outerCache,
				newCache = [ dirruns, doneName ];

			// We can't set arbitrary data on XML nodes, so they don't benefit from combinator caching
			if ( xml ) {
				while ( ( elem = elem[ dir ] ) ) {
					if ( elem.nodeType === 1 || checkNonElements ) {
						if ( matcher( elem, context, xml ) ) {
							return true;
						}
					}
				}
			} else {
				while ( ( elem = elem[ dir ] ) ) {
					if ( elem.nodeType === 1 || checkNonElements ) {
						outerCache = elem[ expando ] || ( elem[ expando ] = {} );

						// Support: IE <9 only
						// Defend against cloned attroperties (jQuery gh-1709)
						uniqueCache = outerCache[ elem.uniqueID ] ||
							( outerCache[ elem.uniqueID ] = {} );

						if ( skip && skip === elem.nodeName.toLowerCase() ) {
							elem = elem[ dir ] || elem;
						} else if ( ( oldCache = uniqueCache[ key ] ) &&
							oldCache[ 0 ] === dirruns && oldCache[ 1 ] === doneName ) {

							// Assign to newCache so results back-propagate to previous elements
							return ( newCache[ 2 ] = oldCache[ 2 ] );
						} else {

							// Reuse newcache so results back-propagate to previous elements
							uniqueCache[ key ] = newCache;

							// A match means we're done; a fail means we have to keep checking
							if ( ( newCache[ 2 ] = matcher( elem, context, xml ) ) ) {
								return true;
							}
						}
					}
				}
			}
			return false;
		};
}

function elementMatcher( matchers ) {
	return matchers.length > 1 ?
		function( elem, context, xml ) {
			var i = matchers.length;
			while ( i-- ) {
				if ( !matchers[ i ]( elem, context, xml ) ) {
					return false;
				}
			}
			return true;
		} :
		matchers[ 0 ];
}

function multipleContexts( selector, contexts, results ) {
	var i = 0,
		len = contexts.length;
	for ( ; i < len; i++ ) {
		Sizzle( selector, contexts[ i ], results );
	}
	return results;
}

function condense( unmatched, map, filter, context, xml ) {
	var elem,
		newUnmatched = [],
		i = 0,
		len = unmatched.length,
		mapped = map != null;

	for ( ; i < len; i++ ) {
		if ( ( elem = unmatched[ i ] ) ) {
			if ( !filter || filter( elem, context, xml ) ) {
				newUnmatched.push( elem );
				if ( mapped ) {
					map.push( i );
				}
			}
		}
	}

	return newUnmatched;
}

function setMatcher( preFilter, selector, matcher, postFilter, postFinder, postSelector ) {
	if ( postFilter && !postFilter[ expando ] ) {
		postFilter = setMatcher( postFilter );
	}
	if ( postFinder && !postFinder[ expando ] ) {
		postFinder = setMatcher( postFinder, postSelector );
	}
	return markFunction( function( seed, results, context, xml ) {
		var temp, i, elem,
			preMap = [],
			postMap = [],
			preexisting = results.length,

			// Get initial elements from seed or context
			elems = seed || multipleContexts(
				selector || "*",
				context.nodeType ? [ context ] : context,
				[]
			),

			// Prefilter to get matcher input, preserving a map for seed-results synchronization
			matcherIn = preFilter && ( seed || !selector ) ?
				condense( elems, preMap, preFilter, context, xml ) :
				elems,

			matcherOut = matcher ?

				// If we have a postFinder, or filtered seed, or non-seed postFilter or preexisting results,
				postFinder || ( seed ? preFilter : preexisting || postFilter ) ?

					// ...intermediate processing is necessary
					[] :

					// ...otherwise use results directly
					results :
				matcherIn;

		// Find primary matches
		if ( matcher ) {
			matcher( matcherIn, matcherOut, context, xml );
		}

		// Apply postFilter
		if ( postFilter ) {
			temp = condense( matcherOut, postMap );
			postFilter( temp, [], context, xml );

			// Un-match failing elements by moving them back to matcherIn
			i = temp.length;
			while ( i-- ) {
				if ( ( elem = temp[ i ] ) ) {
					matcherOut[ postMap[ i ] ] = !( matcherIn[ postMap[ i ] ] = elem );
				}
			}
		}

		if ( seed ) {
			if ( postFinder || preFilter ) {
				if ( postFinder ) {

					// Get the final matcherOut by condensing this intermediate into postFinder contexts
					temp = [];
					i = matcherOut.length;
					while ( i-- ) {
						if ( ( elem = matcherOut[ i ] ) ) {

							// Restore matcherIn since elem is not yet a final match
							temp.push( ( matcherIn[ i ] = elem ) );
						}
					}
					postFinder( null, ( matcherOut = [] ), temp, xml );
				}

				// Move matched elements from seed to results to keep them synchronized
				i = matcherOut.length;
				while ( i-- ) {
					if ( ( elem = matcherOut[ i ] ) &&
						( temp = postFinder ? indexOf( seed, elem ) : preMap[ i ] ) > -1 ) {

						seed[ temp ] = !( results[ temp ] = elem );
					}
				}
			}

		// Add elements to results, through postFinder if defined
		} else {
			matcherOut = condense(
				matcherOut === results ?
					matcherOut.splice( preexisting, matcherOut.length ) :
					matcherOut
			);
			if ( postFinder ) {
				postFinder( null, results, matcherOut, xml );
			} else {
				push.apply( results, matcherOut );
			}
		}
	} );
}

function matcherFromTokens( tokens ) {
	var checkContext, matcher, j,
		len = tokens.length,
		leadingRelative = Expr.relative[ tokens[ 0 ].type ],
		implicitRelative = leadingRelative || Expr.relative[ " " ],
		i = leadingRelative ? 1 : 0,

		// The foundational matcher ensures that elements are reachable from top-level context(s)
		matchContext = addCombinator( function( elem ) {
			return elem === checkContext;
		}, implicitRelative, true ),
		matchAnyContext = addCombinator( function( elem ) {
			return indexOf( checkContext, elem ) > -1;
		}, implicitRelative, true ),
		matchers = [ function( elem, context, xml ) {
			var ret = ( !leadingRelative && ( xml || context !== outermostContext ) ) || (
				( checkContext = context ).nodeType ?
					matchContext( elem, context, xml ) :
					matchAnyContext( elem, context, xml ) );

			// Avoid hanging onto element (issue #299)
			checkContext = null;
			return ret;
		} ];

	for ( ; i < len; i++ ) {
		if ( ( matcher = Expr.relative[ tokens[ i ].type ] ) ) {
			matchers = [ addCombinator( elementMatcher( matchers ), matcher ) ];
		} else {
			matcher = Expr.filter[ tokens[ i ].type ].apply( null, tokens[ i ].matches );

			// Return special upon seeing a positional matcher
			if ( matcher[ expando ] ) {

				// Find the next relative operator (if any) for proper handling
				j = ++i;
				for ( ; j < len; j++ ) {
					if ( Expr.relative[ tokens[ j ].type ] ) {
						break;
					}
				}
				return setMatcher(
					i > 1 && elementMatcher( matchers ),
					i > 1 && toSelector(

					// If the preceding token was a descendant combinator, insert an implicit any-element `*`
					tokens
						.slice( 0, i - 1 )
						.concat( { value: tokens[ i - 2 ].type === " " ? "*" : "" } )
					).replace( rtrim, "$1" ),
					matcher,
					i < j && matcherFromTokens( tokens.slice( i, j ) ),
					j < len && matcherFromTokens( ( tokens = tokens.slice( j ) ) ),
					j < len && toSelector( tokens )
				);
			}
			matchers.push( matcher );
		}
	}

	return elementMatcher( matchers );
}

function matcherFromGroupMatchers( elementMatchers, setMatchers ) {
	var bySet = setMatchers.length > 0,
		byElement = elementMatchers.length > 0,
		superMatcher = function( seed, context, xml, results, outermost ) {
			var elem, j, matcher,
				matchedCount = 0,
				i = "0",
				unmatched = seed && [],
				setMatched = [],
				contextBackup = outermostContext,

				// We must always have either seed elements or outermost context
				elems = seed || byElement && Expr.find[ "TAG" ]( "*", outermost ),

				// Use integer dirruns iff this is the outermost matcher
				dirrunsUnique = ( dirruns += contextBackup == null ? 1 : Math.random() || 0.1 ),
				len = elems.length;

			if ( outermost ) {

				// Support: IE 11+, Edge 17 - 18+
				// IE/Edge sometimes throw a "Permission denied" error when strict-comparing
				// two documents; shallow comparisons work.
				// eslint-disable-next-line eqeqeq
				outermostContext = context == document || context || outermost;
			}

			// Add elements passing elementMatchers directly to results
			// Support: IE<9, Safari
			// Tolerate NodeList properties (IE: "length"; Safari: <number>) matching elements by id
			for ( ; i !== len && ( elem = elems[ i ] ) != null; i++ ) {
				if ( byElement && elem ) {
					j = 0;

					// Support: IE 11+, Edge 17 - 18+
					// IE/Edge sometimes throw a "Permission denied" error when strict-comparing
					// two documents; shallow comparisons work.
					// eslint-disable-next-line eqeqeq
					if ( !context && elem.ownerDocument != document ) {
						setDocument( elem );
						xml = !documentIsHTML;
					}
					while ( ( matcher = elementMatchers[ j++ ] ) ) {
						if ( matcher( elem, context || document, xml ) ) {
							results.push( elem );
							break;
						}
					}
					if ( outermost ) {
						dirruns = dirrunsUnique;
					}
				}

				// Track unmatched elements for set filters
				if ( bySet ) {

					// They will have gone through all possible matchers
					if ( ( elem = !matcher && elem ) ) {
						matchedCount--;
					}

					// Lengthen the array for every element, matched or not
					if ( seed ) {
						unmatched.push( elem );
					}
				}
			}

			// `i` is now the count of elements visited above, and adding it to `matchedCount`
			// makes the latter nonnegative.
			matchedCount += i;

			// Apply set filters to unmatched elements
			// NOTE: This can be skipped if there are no unmatched elements (i.e., `matchedCount`
			// equals `i`), unless we didn't visit _any_ elements in the above loop because we have
			// no element matchers and no seed.
			// Incrementing an initially-string "0" `i` allows `i` to remain a string only in that
			// case, which will result in a "00" `matchedCount` that differs from `i` but is also
			// numerically zero.
			if ( bySet && i !== matchedCount ) {
				j = 0;
				while ( ( matcher = setMatchers[ j++ ] ) ) {
					matcher( unmatched, setMatched, context, xml );
				}

				if ( seed ) {

					// Reintegrate element matches to eliminate the need for sorting
					if ( matchedCount > 0 ) {
						while ( i-- ) {
							if ( !( unmatched[ i ] || setMatched[ i ] ) ) {
								setMatched[ i ] = pop.call( results );
							}
						}
					}

					// Discard index placeholder values to get only actual matches
					setMatched = condense( setMatched );
				}

				// Add matches to results
				push.apply( results, setMatched );

				// Seedless set matches succeeding multiple successful matchers stipulate sorting
				if ( outermost && !seed && setMatched.length > 0 &&
					( matchedCount + setMatchers.length ) > 1 ) {

					Sizzle.uniqueSort( results );
				}
			}

			// Override manipulation of globals by nested matchers
			if ( outermost ) {
				dirruns = dirrunsUnique;
				outermostContext = contextBackup;
			}

			return unmatched;
		};

	return bySet ?
		markFunction( superMatcher ) :
		superMatcher;
}

compile = Sizzle.compile = function( selector, match /* Internal Use Only */ ) {
	var i,
		setMatchers = [],
		elementMatchers = [],
		cached = compilerCache[ selector + " " ];

	if ( !cached ) {

		// Generate a function of recursive functions that can be used to check each element
		if ( !match ) {
			match = tokenize( selector );
		}
		i = match.length;
		while ( i-- ) {
			cached = matcherFromTokens( match[ i ] );
			if ( cached[ expando ] ) {
				setMatchers.push( cached );
			} else {
				elementMatchers.push( cached );
			}
		}

		// Cache the compiled function
		cached = compilerCache(
			selector,
			matcherFromGroupMatchers( elementMatchers, setMatchers )
		);

		// Save selector and tokenization
		cached.selector = selector;
	}
	return cached;
};

/**
 * A low-level selection function that works with Sizzle's compiled
 *  selector functions
 * @param {String|Function} selector A selector or a pre-compiled
 *  selector function built with Sizzle.compile
 * @param {Element} context
 * @param {Array} [results]
 * @param {Array} [seed] A set of elements to match against
 */
select = Sizzle.select = function( selector, context, results, seed ) {
	var i, tokens, token, type, find,
		compiled = typeof selector === "function" && selector,
		match = !seed && tokenize( ( selector = compiled.selector || selector ) );

	results = results || [];

	// Try to minimize operations if there is only one selector in the list and no seed
	// (the latter of which guarantees us context)
	if ( match.length === 1 ) {

		// Reduce context if the leading compound selector is an ID
		tokens = match[ 0 ] = match[ 0 ].slice( 0 );
		if ( tokens.length > 2 && ( token = tokens[ 0 ] ).type === "ID" &&
			context.nodeType === 9 && documentIsHTML && Expr.relative[ tokens[ 1 ].type ] ) {

			context = ( Expr.find[ "ID" ]( token.matches[ 0 ]
				.replace( runescape, funescape ), context ) || [] )[ 0 ];
			if ( !context ) {
				return results;

			// Precompiled matchers will still verify ancestry, so step up a level
			} else if ( compiled ) {
				context = context.parentNode;
			}

			selector = selector.slice( tokens.shift().value.length );
		}

		// Fetch a seed set for right-to-left matching
		i = matchExpr[ "needsContext" ].test( selector ) ? 0 : tokens.length;
		while ( i-- ) {
			token = tokens[ i ];

			// Abort if we hit a combinator
			if ( Expr.relative[ ( type = token.type ) ] ) {
				break;
			}
			if ( ( find = Expr.find[ type ] ) ) {

				// Search, expanding context for leading sibling combinators
				if ( ( seed = find(
					token.matches[ 0 ].replace( runescape, funescape ),
					rsibling.test( tokens[ 0 ].type ) && testContext( context.parentNode ) ||
						context
				) ) ) {

					// If seed is empty or no tokens remain, we can return early
					tokens.splice( i, 1 );
					selector = seed.length && toSelector( tokens );
					if ( !selector ) {
						push.apply( results, seed );
						return results;
					}

					break;
				}
			}
		}
	}

	// Compile and execute a filtering function if one is not provided
	// Provide `match` to avoid retokenization if we modified the selector above
	( compiled || compile( selector, match ) )(
		seed,
		context,
		!documentIsHTML,
		results,
		!context || rsibling.test( selector ) && testContext( context.parentNode ) || context
	);
	return results;
};

// One-time assignments

// Sort stability
support.sortStable = expando.split( "" ).sort( sortOrder ).join( "" ) === expando;

// Support: Chrome 14-35+
// Always assume duplicates if they aren't passed to the comparison function
support.detectDuplicates = !!hasDuplicate;

// Initialize against the default document
setDocument();

// Support: Webkit<537.32 - Safari 6.0.3/Chrome 25 (fixed in Chrome 27)
// Detached nodes confoundingly follow *each other*
support.sortDetached = assert( function( el ) {

	// Should return 1, but returns 4 (following)
	return el.compareDocumentPosition( document.createElement( "fieldset" ) ) & 1;
} );

// Support: IE<8
// Prevent attribute/property "interpolation"
// https://msdn.microsoft.com/en-us/library/ms536429%28VS.85%29.aspx
if ( !assert( function( el ) {
	el.innerHTML = "<a href='#'></a>";
	return el.firstChild.getAttribute( "href" ) === "#";
} ) ) {
	addHandle( "type|href|height|width", function( elem, name, isXML ) {
		if ( !isXML ) {
			return elem.getAttribute( name, name.toLowerCase() === "type" ? 1 : 2 );
		}
	} );
}

// Support: IE<9
// Use defaultValue in place of getAttribute("value")
if ( !support.attributes || !assert( function( el ) {
	el.innerHTML = "<input/>";
	el.firstChild.setAttribute( "value", "" );
	return el.firstChild.getAttribute( "value" ) === "";
} ) ) {
	addHandle( "value", function( elem, _name, isXML ) {
		if ( !isXML && elem.nodeName.toLowerCase() === "input" ) {
			return elem.defaultValue;
		}
	} );
}

// Support: IE<9
// Use getAttributeNode to fetch booleans when getAttribute lies
if ( !assert( function( el ) {
	return el.getAttribute( "disabled" ) == null;
} ) ) {
	addHandle( booleans, function( elem, name, isXML ) {
		var val;
		if ( !isXML ) {
			return elem[ name ] === true ? name.toLowerCase() :
				( val = elem.getAttributeNode( name ) ) && val.specified ?
					val.value :
					null;
		}
	} );
}

return Sizzle;

} )( window );



jQuery.find = Sizzle;
jQuery.expr = Sizzle.selectors;

// Deprecated
jQuery.expr[ ":" ] = jQuery.expr.pseudos;
jQuery.uniqueSort = jQuery.unique = Sizzle.uniqueSort;
jQuery.text = Sizzle.getText;
jQuery.isXMLDoc = Sizzle.isXML;
jQuery.contains = Sizzle.contains;
jQuery.escapeSelector = Sizzle.escape;




var dir = function( elem, dir, until ) {
	var matched = [],
		truncate = until !== undefined;

	while ( ( elem = elem[ dir ] ) && elem.nodeType !== 9 ) {
		if ( elem.nodeType === 1 ) {
			if ( truncate && jQuery( elem ).is( until ) ) {
				break;
			}
			matched.push( elem );
		}
	}
	return matched;
};


var siblings = function( n, elem ) {
	var matched = [];

	for ( ; n; n = n.nextSibling ) {
		if ( n.nodeType === 1 && n !== elem ) {
			matched.push( n );
		}
	}

	return matched;
};


var rneedsContext = jQuery.expr.match.needsContext;



function nodeName( elem, name ) {

	return elem.nodeName && elem.nodeName.toLowerCase() === name.toLowerCase();

}
var rsingleTag = ( /^<([a-z][^\/\0>:\x20\t\r\n\f]*)[\x20\t\r\n\f]*\/?>(?:<\/\1>|)$/i );



// Implement the identical functionality for filter and not
function winnow( elements, qualifier, not ) {
	if ( isFunction( qualifier ) ) {
		return jQuery.grep( elements, function( elem, i ) {
			return !!qualifier.call( elem, i, elem ) !== not;
		} );
	}

	// Single element
	if ( qualifier.nodeType ) {
		return jQuery.grep( elements, function( elem ) {
			return ( elem === qualifier ) !== not;
		} );
	}

	// Arraylike of elements (jQuery, arguments, Array)
	if ( typeof qualifier !== "string" ) {
		return jQuery.grep( elements, function( elem ) {
			return ( indexOf.call( qualifier, elem ) > -1 ) !== not;
		} );
	}

	// Filtered directly for both simple and complex selectors
	return jQuery.filter( qualifier, elements, not );
}

jQuery.filter = function( expr, elems, not ) {
	var elem = elems[ 0 ];

	if ( not ) {
		expr = ":not(" + expr + ")";
	}

	if ( elems.length === 1 && elem.nodeType === 1 ) {
		return jQuery.find.matchesSelector( elem, expr ) ? [ elem ] : [];
	}

	return jQuery.find.matches( expr, jQuery.grep( elems, function( elem ) {
		return elem.nodeType === 1;
	} ) );
};

jQuery.fn.extend( {
	find: function( selector ) {
		var i, ret,
			len = this.length,
			self = this;

		if ( typeof selector !== "string" ) {
			return this.pushStack( jQuery( selector ).filter( function() {
				for ( i = 0; i < len; i++ ) {
					if ( jQuery.contains( self[ i ], this ) ) {
						return true;
					}
				}
			} ) );
		}

		ret = this.pushStack( [] );

		for ( i = 0; i < len; i++ ) {
			jQuery.find( selector, self[ i ], ret );
		}

		return len > 1 ? jQuery.uniqueSort( ret ) : ret;
	},
	filter: function( selector ) {
		return this.pushStack( winnow( this, selector || [], false ) );
	},
	not: function( selector ) {
		return this.pushStack( winnow( this, selector || [], true ) );
	},
	is: function( selector ) {
		return !!winnow(
			this,

			// If this is a positional/relative selector, check membership in the returned set
			// so $("p:first").is("p:last") won't return true for a doc with two "p".
			typeof selector === "string" && rneedsContext.test( selector ) ?
				jQuery( selector ) :
				selector || [],
			false
		).length;
	}
} );


// Initialize a jQuery object


// A central reference to the root jQuery(document)
var rootjQuery,

	// A simple way to check for HTML strings
	// Prioritize #id over <tag> to avoid XSS via location.hash (#9521)
	// Strict HTML recognition (#11290: must start with <)
	// Shortcut simple #id case for speed
	rquickExpr = /^(?:\s*(<[\w\W]+>)[^>]*|#([\w-]+))$/,

	init = jQuery.fn.init = function( selector, context, root ) {
		var match, elem;

		// HANDLE: $(""), $(null), $(undefined), $(false)
		if ( !selector ) {
			return this;
		}

		// Method init() accepts an alternate rootjQuery
		// so migrate can support jQuery.sub (gh-2101)
		root = root || rootjQuery;

		// Handle HTML strings
		if ( typeof selector === "string" ) {
			if ( selector[ 0 ] === "<" &&
				selector[ selector.length - 1 ] === ">" &&
				selector.length >= 3 ) {

				// Assume that strings that start and end with <> are HTML and skip the regex check
				match = [ null, selector, null ];

			} else {
				match = rquickExpr.exec( selector );
			}

			// Match html or make sure no context is specified for #id
			if ( match && ( match[ 1 ] || !context ) ) {

				// HANDLE: $(html) -> $(array)
				if ( match[ 1 ] ) {
					context = context instanceof jQuery ? context[ 0 ] : context;

					// Option to run scripts is true for back-compat
					// Intentionally let the error be thrown if parseHTML is not present
					jQuery.merge( this, jQuery.parseHTML(
						match[ 1 ],
						context && context.nodeType ? context.ownerDocument || context : document,
						true
					) );

					// HANDLE: $(html, props)
					if ( rsingleTag.test( match[ 1 ] ) && jQuery.isPlainObject( context ) ) {
						for ( match in context ) {

							// Properties of context are called as methods if possible
							if ( isFunction( this[ match ] ) ) {
								this[ match ]( context[ match ] );

							// ...and otherwise set as attributes
							} else {
								this.attr( match, context[ match ] );
							}
						}
					}

					return this;

				// HANDLE: $(#id)
				} else {
					elem = document.getElementById( match[ 2 ] );

					if ( elem ) {

						// Inject the element directly into the jQuery object
						this[ 0 ] = elem;
						this.length = 1;
					}
					return this;
				}

			// HANDLE: $(expr, $(...))
			} else if ( !context || context.jquery ) {
				return ( context || root ).find( selector );

			// HANDLE: $(expr, context)
			// (which is just equivalent to: $(context).find(expr)
			} else {
				return this.constructor( context ).find( selector );
			}

		// HANDLE: $(DOMElement)
		} else if ( selector.nodeType ) {
			this[ 0 ] = selector;
			this.length = 1;
			return this;

		// HANDLE: $(function)
		// Shortcut for document ready
		} else if ( isFunction( selector ) ) {
			return root.ready !== undefined ?
				root.ready( selector ) :

				// Execute immediately if ready is not present
				selector( jQuery );
		}

		return jQuery.makeArray( selector, this );
	};

// Give the init function the jQuery prototype for later instantiation
init.prototype = jQuery.fn;

// Initialize central reference
rootjQuery = jQuery( document );


var rparentsprev = /^(?:parents|prev(?:Until|All))/,

	// Methods guaranteed to produce a unique set when starting from a unique set
	guaranteedUnique = {
		children: true,
		contents: true,
		next: true,
		prev: true
	};

jQuery.fn.extend( {
	has: function( target ) {
		var targets = jQuery( target, this ),
			l = targets.length;

		return this.filter( function() {
			var i = 0;
			for ( ; i < l; i++ ) {
				if ( jQuery.contains( this, targets[ i ] ) ) {
					return true;
				}
			}
		} );
	},

	closest: function( selectors, context ) {
		var cur,
			i = 0,
			l = this.length,
			matched = [],
			targets = typeof selectors !== "string" && jQuery( selectors );

		// Positional selectors never match, since there's no _selection_ context
		if ( !rneedsContext.test( selectors ) ) {
			for ( ; i < l; i++ ) {
				for ( cur = this[ i ]; cur && cur !== context; cur = cur.parentNode ) {

					// Always skip document fragments
					if ( cur.nodeType < 11 && ( targets ?
						targets.index( cur ) > -1 :

						// Don't pass non-elements to Sizzle
						cur.nodeType === 1 &&
							jQuery.find.matchesSelector( cur, selectors ) ) ) {

						matched.push( cur );
						break;
					}
				}
			}
		}

		return this.pushStack( matched.length > 1 ? jQuery.uniqueSort( matched ) : matched );
	},

	// Determine the position of an element within the set
	index: function( elem ) {

		// No argument, return index in parent
		if ( !elem ) {
			return ( this[ 0 ] && this[ 0 ].parentNode ) ? this.first().prevAll().length : -1;
		}

		// Index in selector
		if ( typeof elem === "string" ) {
			return indexOf.call( jQuery( elem ), this[ 0 ] );
		}

		// Locate the position of the desired element
		return indexOf.call( this,

			// If it receives a jQuery object, the first element is used
			elem.jquery ? elem[ 0 ] : elem
		);
	},

	add: function( selector, context ) {
		return this.pushStack(
			jQuery.uniqueSort(
				jQuery.merge( this.get(), jQuery( selector, context ) )
			)
		);
	},

	addBack: function( selector ) {
		return this.add( selector == null ?
			this.prevObject : this.prevObject.filter( selector )
		);
	}
} );

function sibling( cur, dir ) {
	while ( ( cur = cur[ dir ] ) && cur.nodeType !== 1 ) {}
	return cur;
}

jQuery.each( {
	parent: function( elem ) {
		var parent = elem.parentNode;
		return parent && parent.nodeType !== 11 ? parent : null;
	},
	parents: function( elem ) {
		return dir( elem, "parentNode" );
	},
	parentsUntil: function( elem, _i, until ) {
		return dir( elem, "parentNode", until );
	},
	next: function( elem ) {
		return sibling( elem, "nextSibling" );
	},
	prev: function( elem ) {
		return sibling( elem, "previousSibling" );
	},
	nextAll: function( elem ) {
		return dir( elem, "nextSibling" );
	},
	prevAll: function( elem ) {
		return dir( elem, "previousSibling" );
	},
	nextUntil: function( elem, _i, until ) {
		return dir( elem, "nextSibling", until );
	},
	prevUntil: function( elem, _i, until ) {
		return dir( elem, "previousSibling", until );
	},
	siblings: function( elem ) {
		return siblings( ( elem.parentNode || {} ).firstChild, elem );
	},
	children: function( elem ) {
		return siblings( elem.firstChild );
	},
	contents: function( elem ) {
		if ( elem.contentDocument != null &&

			// Support: IE 11+
			// <object> elements with no `data` attribute has an object
			// `contentDocument` with a `null` prototype.
			getProto( elem.contentDocument ) ) {

			return elem.contentDocument;
		}

		// Support: IE 9 - 11 only, iOS 7 only, Android Browser <=4.3 only
		// Treat the template element as a regular one in browsers that
		// don't support it.
		if ( nodeName( elem, "template" ) ) {
			elem = elem.content || elem;
		}

		return jQuery.merge( [], elem.childNodes );
	}
}, function( name, fn ) {
	jQuery.fn[ name ] = function( until, selector ) {
		var matched = jQuery.map( this, fn, until );

		if ( name.slice( -5 ) !== "Until" ) {
			selector = until;
		}

		if ( selector && typeof selector === "string" ) {
			matched = jQuery.filter( selector, matched );
		}

		if ( this.length > 1 ) {

			// Remove duplicates
			if ( !guaranteedUnique[ name ] ) {
				jQuery.uniqueSort( matched );
			}

			// Reverse order for parents* and prev-derivatives
			if ( rparentsprev.test( name ) ) {
				matched.reverse();
			}
		}

		return this.pushStack( matched );
	};
} );
var rnothtmlwhite = ( /[^\x20\t\r\n\f]+/g );



// Convert String-formatted options into Object-formatted ones
function createOptions( options ) {
	var object = {};
	jQuery.each( options.match( rnothtmlwhite ) || [], function( _, flag ) {
		object[ flag ] = true;
	} );
	return object;
}

/*
 * Create a callback list using the following parameters:
 *
 *	options: an optional list of space-separated options that will change how
 *			the callback list behaves or a more traditional option object
 *
 * By default a callback list will act like an event callback list and can be
 * "fired" multiple times.
 *
 * Possible options:
 *
 *	once:			will ensure the callback list can only be fired once (like a Deferred)
 *
 *	memory:			will keep track of previous values and will call any callback added
 *					after the list has been fired right away with the latest "memorized"
 *					values (like a Deferred)
 *
 *	unique:			will ensure a callback can only be added once (no duplicate in the list)
 *
 *	stopOnFalse:	interrupt callings when a callback returns false
 *
 */
jQuery.Callbacks = function( options ) {

	// Convert options from String-formatted to Object-formatted if needed
	// (we check in cache first)
	options = typeof options === "string" ?
		createOptions( options ) :
		jQuery.extend( {}, options );

	var // Flag to know if list is currently firing
		firing,

		// Last fire value for non-forgettable lists
		memory,

		// Flag to know if list was already fired
		fired,

		// Flag to prevent firing
		locked,

		// Actual callback list
		list = [],

		// Queue of execution data for repeatable lists
		queue = [],

		// Index of currently firing callback (modified by add/remove as needed)
		firingIndex = -1,

		// Fire callbacks
		fire = function() {

			// Enforce single-firing
			locked = locked || options.once;

			// Execute callbacks for all pending executions,
			// respecting firingIndex overrides and runtime changes
			fired = firing = true;
			for ( ; queue.length; firingIndex = -1 ) {
				memory = queue.shift();
				while ( ++firingIndex < list.length ) {

					// Run callback and check for early termination
					if ( list[ firingIndex ].apply( memory[ 0 ], memory[ 1 ] ) === false &&
						options.stopOnFalse ) {

						// Jump to end and forget the data so .add doesn't re-fire
						firingIndex = list.length;
						memory = false;
					}
				}
			}

			// Forget the data if we're done with it
			if ( !options.memory ) {
				memory = false;
			}

			firing = false;

			// Clean up if we're done firing for good
			if ( locked ) {

				// Keep an empty list if we have data for future add calls
				if ( memory ) {
					list = [];

				// Otherwise, this object is spent
				} else {
					list = "";
				}
			}
		},

		// Actual Callbacks object
		self = {

			// Add a callback or a collection of callbacks to the list
			add: function() {
				if ( list ) {

					// If we have memory from a past run, we should fire after adding
					if ( memory && !firing ) {
						firingIndex = list.length - 1;
						queue.push( memory );
					}

					( function add( args ) {
						jQuery.each( args, function( _, arg ) {
							if ( isFunction( arg ) ) {
								if ( !options.unique || !self.has( arg ) ) {
									list.push( arg );
								}
							} else if ( arg && arg.length && toType( arg ) !== "string" ) {

								// Inspect recursively
								add( arg );
							}
						} );
					} )( arguments );

					if ( memory && !firing ) {
						fire();
					}
				}
				return this;
			},

			// Remove a callback from the list
			remove: function() {
				jQuery.each( arguments, function( _, arg ) {
					var index;
					while ( ( index = jQuery.inArray( arg, list, index ) ) > -1 ) {
						list.splice( index, 1 );

						// Handle firing indexes
						if ( index <= firingIndex ) {
							firingIndex--;
						}
					}
				} );
				return this;
			},

			// Check if a given callback is in the list.
			// If no argument is given, return whether or not list has callbacks attached.
			has: function( fn ) {
				return fn ?
					jQuery.inArray( fn, list ) > -1 :
					list.length > 0;
			},

			// Remove all callbacks from the list
			empty: function() {
				if ( list ) {
					list = [];
				}
				return this;
			},

			// Disable .fire and .add
			// Abort any current/pending executions
			// Clear all callbacks and values
			disable: function() {
				locked = queue = [];
				list = memory = "";
				return this;
			},
			disabled: function() {
				return !list;
			},

			// Disable .fire
			// Also disable .add unless we have memory (since it would have no effect)
			// Abort any pending executions
			lock: function() {
				locked = queue = [];
				if ( !memory && !firing ) {
					list = memory = "";
				}
				return this;
			},
			locked: function() {
				return !!locked;
			},

			// Call all callbacks with the given context and arguments
			fireWith: function( context, args ) {
				if ( !locked ) {
					args = args || [];
					args = [ context, args.slice ? args.slice() : args ];
					queue.push( args );
					if ( !firing ) {
						fire();
					}
				}
				return this;
			},

			// Call all the callbacks with the given arguments
			fire: function() {
				self.fireWith( this, arguments );
				return this;
			},

			// To know if the callbacks have already been called at least once
			fired: function() {
				return !!fired;
			}
		};

	return self;
};


function Identity( v ) {
	return v;
}
function Thrower( ex ) {
	throw ex;
}

function adoptValue( value, resolve, reject, noValue ) {
	var method;

	try {

		// Check for promise aspect first to privilege synchronous behavior
		if ( value && isFunction( ( method = value.promise ) ) ) {
			method.call( value ).done( resolve ).fail( reject );

		// Other thenables
		} else if ( value && isFunction( ( method = value.then ) ) ) {
			method.call( value, resolve, reject );

		// Other non-thenables
		} else {

			// Control `resolve` arguments by letting Array#slice cast boolean `noValue` to integer:
			// * false: [ value ].slice( 0 ) => resolve( value )
			// * true: [ value ].slice( 1 ) => resolve()
			resolve.apply( undefined, [ value ].slice( noValue ) );
		}

	// For Promises/A+, convert exceptions into rejections
	// Since jQuery.when doesn't unwrap thenables, we can skip the extra checks appearing in
	// Deferred#then to conditionally suppress rejection.
	} catch ( value ) {

		// Support: Android 4.0 only
		// Strict mode functions invoked without .call/.apply get global-object context
		reject.apply( undefined, [ value ] );
	}
}

jQuery.extend( {

	Deferred: function( func ) {
		var tuples = [

				// action, add listener, callbacks,
				// ... .then handlers, argument index, [final state]
				[ "notify", "progress", jQuery.Callbacks( "memory" ),
					jQuery.Callbacks( "memory" ), 2 ],
				[ "resolve", "done", jQuery.Callbacks( "once memory" ),
					jQuery.Callbacks( "once memory" ), 0, "resolved" ],
				[ "reject", "fail", jQuery.Callbacks( "once memory" ),
					jQuery.Callbacks( "once memory" ), 1, "rejected" ]
			],
			state = "pending",
			promise = {
				state: function() {
					return state;
				},
				always: function() {
					deferred.done( arguments ).fail( arguments );
					return this;
				},
				"catch": function( fn ) {
					return promise.then( null, fn );
				},

				// Keep pipe for back-compat
				pipe: function( /* fnDone, fnFail, fnProgress */ ) {
					var fns = arguments;

					return jQuery.Deferred( function( newDefer ) {
						jQuery.each( tuples, function( _i, tuple ) {

							// Map tuples (progress, done, fail) to arguments (done, fail, progress)
							var fn = isFunction( fns[ tuple[ 4 ] ] ) && fns[ tuple[ 4 ] ];

							// deferred.progress(function() { bind to newDefer or newDefer.notify })
							// deferred.done(function() { bind to newDefer or newDefer.resolve })
							// deferred.fail(function() { bind to newDefer or newDefer.reject })
							deferred[ tuple[ 1 ] ]( function() {
								var returned = fn && fn.apply( this, arguments );
								if ( returned && isFunction( returned.promise ) ) {
									returned.promise()
										.progress( newDefer.notify )
										.done( newDefer.resolve )
										.fail( newDefer.reject );
								} else {
									newDefer[ tuple[ 0 ] + "With" ](
										this,
										fn ? [ returned ] : arguments
									);
								}
							} );
						} );
						fns = null;
					} ).promise();
				},
				then: function( onFulfilled, onRejected, onProgress ) {
					var maxDepth = 0;
					function resolve( depth, deferred, handler, special ) {
						return function() {
							var that = this,
								args = arguments,
								mightThrow = function() {
									var returned, then;

									// Support: Promises/A+ section 2.3.3.3.3
									// https://promisesaplus.com/#point-59
									// Ignore double-resolution attempts
									if ( depth < maxDepth ) {
										return;
									}

									returned = handler.apply( that, args );

									// Support: Promises/A+ section 2.3.1
									// https://promisesaplus.com/#point-48
									if ( returned === deferred.promise() ) {
										throw new TypeError( "Thenable self-resolution" );
									}

									// Support: Promises/A+ sections 2.3.3.1, 3.5
									// https://promisesaplus.com/#point-54
									// https://promisesaplus.com/#point-75
									// Retrieve `then` only once
									then = returned &&

										// Support: Promises/A+ section 2.3.4
										// https://promisesaplus.com/#point-64
										// Only check objects and functions for thenability
										( typeof returned === "object" ||
											typeof returned === "function" ) &&
										returned.then;

									// Handle a returned thenable
									if ( isFunction( then ) ) {

										// Special processors (notify) just wait for resolution
										if ( special ) {
											then.call(
												returned,
												resolve( maxDepth, deferred, Identity, special ),
												resolve( maxDepth, deferred, Thrower, special )
											);

										// Normal processors (resolve) also hook into progress
										} else {

											// ...and disregard older resolution values
											maxDepth++;

											then.call(
												returned,
												resolve( maxDepth, deferred, Identity, special ),
												resolve( maxDepth, deferred, Thrower, special ),
												resolve( maxDepth, deferred, Identity,
													deferred.notifyWith )
											);
										}

									// Handle all other returned values
									} else {

										// Only substitute handlers pass on context
										// and multiple values (non-spec behavior)
										if ( handler !== Identity ) {
											that = undefined;
											args = [ returned ];
										}

										// Process the value(s)
										// Default process is resolve
										( special || deferred.resolveWith )( that, args );
									}
								},

								// Only normal processors (resolve) catch and reject exceptions
								process = special ?
									mightThrow :
									function() {
										try {
											mightThrow();
										} catch ( e ) {

											if ( jQuery.Deferred.exceptionHook ) {
												jQuery.Deferred.exceptionHook( e,
													process.stackTrace );
											}

											// Support: Promises/A+ section 2.3.3.3.4.1
											// https://promisesaplus.com/#point-61
											// Ignore post-resolution exceptions
											if ( depth + 1 >= maxDepth ) {

												// Only substitute handlers pass on context
												// and multiple values (non-spec behavior)
												if ( handler !== Thrower ) {
													that = undefined;
													args = [ e ];
												}

												deferred.rejectWith( that, args );
											}
										}
									};

							// Support: Promises/A+ section 2.3.3.3.1
							// https://promisesaplus.com/#point-57
							// Re-resolve promises immediately to dodge false rejection from
							// subsequent errors
							if ( depth ) {
								process();
							} else {

								// Call an optional hook to record the stack, in case of exception
								// since it's otherwise lost when execution goes async
								if ( jQuery.Deferred.getStackHook ) {
									process.stackTrace = jQuery.Deferred.getStackHook();
								}
								window.setTimeout( process );
							}
						};
					}

					return jQuery.Deferred( function( newDefer ) {

						// progress_handlers.add( ... )
						tuples[ 0 ][ 3 ].add(
							resolve(
								0,
								newDefer,
								isFunction( onProgress ) ?
									onProgress :
									Identity,
								newDefer.notifyWith
							)
						);

						// fulfilled_handlers.add( ... )
						tuples[ 1 ][ 3 ].add(
							resolve(
								0,
								newDefer,
								isFunction( onFulfilled ) ?
									onFulfilled :
									Identity
							)
						);

						// rejected_handlers.add( ... )
						tuples[ 2 ][ 3 ].add(
							resolve(
								0,
								newDefer,
								isFunction( onRejected ) ?
									onRejected :
									Thrower
							)
						);
					} ).promise();
				},

				// Get a promise for this deferred
				// If obj is provided, the promise aspect is added to the object
				promise: function( obj ) {
					return obj != null ? jQuery.extend( obj, promise ) : promise;
				}
			},
			deferred = {};

		// Add list-specific methods
		jQuery.each( tuples, function( i, tuple ) {
			var list = tuple[ 2 ],
				stateString = tuple[ 5 ];

			// promise.progress = list.add
			// promise.done = list.add
			// promise.fail = list.add
			promise[ tuple[ 1 ] ] = list.add;

			// Handle state
			if ( stateString ) {
				list.add(
					function() {

						// state = "resolved" (i.e., fulfilled)
						// state = "rejected"
						state = stateString;
					},

					// rejected_callbacks.disable
					// fulfilled_callbacks.disable
					tuples[ 3 - i ][ 2 ].disable,

					// rejected_handlers.disable
					// fulfilled_handlers.disable
					tuples[ 3 - i ][ 3 ].disable,

					// progress_callbacks.lock
					tuples[ 0 ][ 2 ].lock,

					// progress_handlers.lock
					tuples[ 0 ][ 3 ].lock
				);
			}

			// progress_handlers.fire
			// fulfilled_handlers.fire
			// rejected_handlers.fire
			list.add( tuple[ 3 ].fire );

			// deferred.notify = function() { deferred.notifyWith(...) }
			// deferred.resolve = function() { deferred.resolveWith(...) }
			// deferred.reject = function() { deferred.rejectWith(...) }
			deferred[ tuple[ 0 ] ] = function() {
				deferred[ tuple[ 0 ] + "With" ]( this === deferred ? undefined : this, arguments );
				return this;
			};

			// deferred.notifyWith = list.fireWith
			// deferred.resolveWith = list.fireWith
			// deferred.rejectWith = list.fireWith
			deferred[ tuple[ 0 ] + "With" ] = list.fireWith;
		} );

		// Make the deferred a promise
		promise.promise( deferred );

		// Call given func if any
		if ( func ) {
			func.call( deferred, deferred );
		}

		// All done!
		return deferred;
	},

	// Deferred helper
	when: function( singleValue ) {
		var

			// count of uncompleted subordinates
			remaining = arguments.length,

			// count of unprocessed arguments
			i = remaining,

			// subordinate fulfillment data
			resolveContexts = Array( i ),
			resolveValues = slice.call( arguments ),

			// the primary Deferred
			primary = jQuery.Deferred(),

			// subordinate callback factory
			updateFunc = function( i ) {
				return function( value ) {
					resolveContexts[ i ] = this;
					resolveValues[ i ] = arguments.length > 1 ? slice.call( arguments ) : value;
					if ( !( --remaining ) ) {
						primary.resolveWith( resolveContexts, resolveValues );
					}
				};
			};

		// Single- and empty arguments are adopted like Promise.resolve
		if ( remaining <= 1 ) {
			adoptValue( singleValue, primary.done( updateFunc( i ) ).resolve, primary.reject,
				!remaining );

			// Use .then() to unwrap secondary thenables (cf. gh-3000)
			if ( primary.state() === "pending" ||
				isFunction( resolveValues[ i ] && resolveValues[ i ].then ) ) {

				return primary.then();
			}
		}

		// Multiple arguments are aggregated like Promise.all array elements
		while ( i-- ) {
			adoptValue( resolveValues[ i ], updateFunc( i ), primary.reject );
		}

		return primary.promise();
	}
} );


// These usually indicate a programmer mistake during development,
// warn about them ASAP rather than swallowing them by default.
var rerrorNames = /^(Eval|Internal|Range|Reference|Syntax|Type|URI)Error$/;

jQuery.Deferred.exceptionHook = function( error, stack ) {

	// Support: IE 8 - 9 only
	// Console exists when dev tools are open, which can happen at any time
	if ( window.console && window.console.warn && error && rerrorNames.test( error.name ) ) {
		window.console.warn( "jQuery.Deferred exception: " + error.message, error.stack, stack );
	}
};




jQuery.readyException = function( error ) {
	window.setTimeout( function() {
		throw error;
	} );
};




// The deferred used on DOM ready
var readyList = jQuery.Deferred();

jQuery.fn.ready = function( fn ) {

	readyList
		.then( fn )

		// Wrap jQuery.readyException in a function so that the lookup
		// happens at the time of error handling instead of callback
		// registration.
		.catch( function( error ) {
			jQuery.readyException( error );
		} );

	return this;
};

jQuery.extend( {

	// Is the DOM ready to be used? Set to true once it occurs.
	isReady: false,

	// A counter to track how many items to wait for before
	// the ready event fires. See #6781
	readyWait: 1,

	// Handle when the DOM is ready
	ready: function( wait ) {

		// Abort if there are pending holds or we're already ready
		if ( wait === true ? --jQuery.readyWait : jQuery.isReady ) {
			return;
		}

		// Remember that the DOM is ready
		jQuery.isReady = true;

		// If a normal DOM Ready event fired, decrement, and wait if need be
		if ( wait !== true && --jQuery.readyWait > 0 ) {
			return;
		}

		// If there are functions bound, to execute
		readyList.resolveWith( document, [ jQuery ] );
	}
} );

jQuery.ready.then = readyList.then;

// The ready event handler and self cleanup method
function completed() {
	document.removeEventListener( "DOMContentLoaded", completed );
	window.removeEventListener( "load", completed );
	jQuery.ready();
}

// Catch cases where $(document).ready() is called
// after the browser event has already occurred.
// Support: IE <=9 - 10 only
// Older IE sometimes signals "interactive" too soon
if ( document.readyState === "complete" ||
	( document.readyState !== "loading" && !document.documentElement.doScroll ) ) {

	// Handle it asynchronously to allow scripts the opportunity to delay ready
	window.setTimeout( jQuery.ready );

} else {

	// Use the handy event callback
	document.addEventListener( "DOMContentLoaded", completed );

	// A fallback to window.onload, that will always work
	window.addEventListener( "load", completed );
}




// Multifunctional method to get and set values of a collection
// The value/s can optionally be executed if it's a function
var access = function( elems, fn, key, value, chainable, emptyGet, raw ) {
	var i = 0,
		len = elems.length,
		bulk = key == null;

	// Sets many values
	if ( toType( key ) === "object" ) {
		chainable = true;
		for ( i in key ) {
			access( elems, fn, i, key[ i ], true, emptyGet, raw );
		}

	// Sets one value
	} else if ( value !== undefined ) {
		chainable = true;

		if ( !isFunction( value ) ) {
			raw = true;
		}

		if ( bulk ) {

			// Bulk operations run against the entire set
			if ( raw ) {
				fn.call( elems, value );
				fn = null;

			// ...except when executing function values
			} else {
				bulk = fn;
				fn = function( elem, _key, value ) {
					return bulk.call( jQuery( elem ), value );
				};
			}
		}

		if ( fn ) {
			for ( ; i < len; i++ ) {
				fn(
					elems[ i ], key, raw ?
						value :
						value.call( elems[ i ], i, fn( elems[ i ], key ) )
				);
			}
		}
	}

	if ( chainable ) {
		return elems;
	}

	// Gets
	if ( bulk ) {
		return fn.call( elems );
	}

	return len ? fn( elems[ 0 ], key ) : emptyGet;
};


// Matches dashed string for camelizing
var rmsPrefix = /^-ms-/,
	rdashAlpha = /-([a-z])/g;

// Used by camelCase as callback to replace()
function fcamelCase( _all, letter ) {
	return letter.toUpperCase();
}

// Convert dashed to camelCase; used by the css and data modules
// Support: IE <=9 - 11, Edge 12 - 15
// Microsoft forgot to hump their vendor prefix (#9572)
function camelCase( string ) {
	return string.replace( rmsPrefix, "ms-" ).replace( rdashAlpha, fcamelCase );
}
var acceptData = function( owner ) {

	// Accepts only:
	//  - Node
	//    - Node.ELEMENT_NODE
	//    - Node.DOCUMENT_NODE
	//  - Object
	//    - Any
	return owner.nodeType === 1 || owner.nodeType === 9 || !( +owner.nodeType );
};




function Data() {
	this.expando = jQuery.expando + Data.uid++;
}

Data.uid = 1;

Data.prototype = {

	cache: function( owner ) {

		// Check if the owner object already has a cache
		var value = owner[ this.expando ];

		// If not, create one
		if ( !value ) {
			value = {};

			// We can accept data for non-element nodes in modern browsers,
			// but we should not, see #8335.
			// Always return an empty object.
			if ( acceptData( owner ) ) {

				// If it is a node unlikely to be stringify-ed or looped over
				// use plain assignment
				if ( owner.nodeType ) {
					owner[ this.expando ] = value;

				// Otherwise secure it in a non-enumerable property
				// configurable must be true to allow the property to be
				// deleted when data is removed
				} else {
					Object.defineProperty( owner, this.expando, {
						value: value,
						configurable: true
					} );
				}
			}
		}

		return value;
	},
	set: function( owner, data, value ) {
		var prop,
			cache = this.cache( owner );

		// Handle: [ owner, key, value ] args
		// Always use camelCase key (gh-2257)
		if ( typeof data === "string" ) {
			cache[ camelCase( data ) ] = value;

		// Handle: [ owner, { properties } ] args
		} else {

			// Copy the properties one-by-one to the cache object
			for ( prop in data ) {
				cache[ camelCase( prop ) ] = data[ prop ];
			}
		}
		return cache;
	},
	get: function( owner, key ) {
		return key === undefined ?
			this.cache( owner ) :

			// Always use camelCase key (gh-2257)
			owner[ this.expando ] && owner[ this.expando ][ camelCase( key ) ];
	},
	access: function( owner, key, value ) {

		// In cases where either:
		//
		//   1. No key was specified
		//   2. A string key was specified, but no value provided
		//
		// Take the "read" path and allow the get method to determine
		// which value to return, respectively either:
		//
		//   1. The entire cache object
		//   2. The data stored at the key
		//
		if ( key === undefined ||
				( ( key && typeof key === "string" ) && value === undefined ) ) {

			return this.get( owner, key );
		}

		// When the key is not a string, or both a key and value
		// are specified, set or extend (existing objects) with either:
		//
		//   1. An object of properties
		//   2. A key and value
		//
		this.set( owner, key, value );

		// Since the "set" path can have two possible entry points
		// return the expected data based on which path was taken[*]
		return value !== undefined ? value : key;
	},
	remove: function( owner, key ) {
		var i,
			cache = owner[ this.expando ];

		if ( cache === undefined ) {
			return;
		}

		if ( key !== undefined ) {

			// Support array or space separated string of keys
			if ( Array.isArray( key ) ) {

				// If key is an array of keys...
				// We always set camelCase keys, so remove that.
				key = key.map( camelCase );
			} else {
				key = camelCase( key );

				// If a key with the spaces exists, use it.
				// Otherwise, create an array by matching non-whitespace
				key = key in cache ?
					[ key ] :
					( key.match( rnothtmlwhite ) || [] );
			}

			i = key.length;

			while ( i-- ) {
				delete cache[ key[ i ] ];
			}
		}

		// Remove the expando if there's no more data
		if ( key === undefined || jQuery.isEmptyObject( cache ) ) {

			// Support: Chrome <=35 - 45
			// Webkit & Blink performance suffers when deleting properties
			// from DOM nodes, so set to undefined instead
			// https://bugs.chromium.org/p/chromium/issues/detail?id=378607 (bug restricted)
			if ( owner.nodeType ) {
				owner[ this.expando ] = undefined;
			} else {
				delete owner[ this.expando ];
			}
		}
	},
	hasData: function( owner ) {
		var cache = owner[ this.expando ];
		return cache !== undefined && !jQuery.isEmptyObject( cache );
	}
};
var dataPriv = new Data();

var dataUser = new Data();



//	Implementation Summary
//
//	1. Enforce API surface and semantic compatibility with 1.9.x branch
//	2. Improve the module's maintainability by reducing the storage
//		paths to a single mechanism.
//	3. Use the same single mechanism to support "private" and "user" data.
//	4. _Never_ expose "private" data to user code (TODO: Drop _data, _removeData)
//	5. Avoid exposing implementation details on user objects (eg. expando properties)
//	6. Provide a clear path for implementation upgrade to WeakMap in 2014

var rbrace = /^(?:\{[\w\W]*\}|\[[\w\W]*\])$/,
	rmultiDash = /[A-Z]/g;

function getData( data ) {
	if ( data === "true" ) {
		return true;
	}

	if ( data === "false" ) {
		return false;
	}

	if ( data === "null" ) {
		return null;
	}

	// Only convert to a number if it doesn't change the string
	if ( data === +data + "" ) {
		return +data;
	}

	if ( rbrace.test( data ) ) {
		return JSON.parse( data );
	}

	return data;
}

function dataAttr( elem, key, data ) {
	var name;

	// If nothing was found internally, try to fetch any
	// data from the HTML5 data-* attribute
	if ( data === undefined && elem.nodeType === 1 ) {
		name = "data-" + key.replace( rmultiDash, "-$&" ).toLowerCase();
		data = elem.getAttribute( name );

		if ( typeof data === "string" ) {
			try {
				data = getData( data );
			} catch ( e ) {}

			// Make sure we set the data so it isn't changed later
			dataUser.set( elem, key, data );
		} else {
			data = undefined;
		}
	}
	return data;
}

jQuery.extend( {
	hasData: function( elem ) {
		return dataUser.hasData( elem ) || dataPriv.hasData( elem );
	},

	data: function( elem, name, data ) {
		return dataUser.access( elem, name, data );
	},

	removeData: function( elem, name ) {
		dataUser.remove( elem, name );
	},

	// TODO: Now that all calls to _data and _removeData have been replaced
	// with direct calls to dataPriv methods, these can be deprecated.
	_data: function( elem, name, data ) {
		return dataPriv.access( elem, name, data );
	},

	_removeData: function( elem, name ) {
		dataPriv.remove( elem, name );
	}
} );

jQuery.fn.extend( {
	data: function( key, value ) {
		var i, name, data,
			elem = this[ 0 ],
			attrs = elem && elem.attributes;

		// Gets all values
		if ( key === undefined ) {
			if ( this.length ) {
				data = dataUser.get( elem );

				if ( elem.nodeType === 1 && !dataPriv.get( elem, "hasDataAttrs" ) ) {
					i = attrs.length;
					while ( i-- ) {

						// Support: IE 11 only
						// The attrs elements can be null (#14894)
						if ( attrs[ i ] ) {
							name = attrs[ i ].name;
							if ( name.indexOf( "data-" ) === 0 ) {
								name = camelCase( name.slice( 5 ) );
								dataAttr( elem, name, data[ name ] );
							}
						}
					}
					dataPriv.set( elem, "hasDataAttrs", true );
				}
			}

			return data;
		}

		// Sets multiple values
		if ( typeof key === "object" ) {
			return this.each( function() {
				dataUser.set( this, key );
			} );
		}

		return access( this, function( value ) {
			var data;

			// The calling jQuery object (element matches) is not empty
			// (and therefore has an element appears at this[ 0 ]) and the
			// `value` parameter was not undefined. An empty jQuery object
			// will result in `undefined` for elem = this[ 0 ] which will
			// throw an exception if an attempt to read a data cache is made.
			if ( elem && value === undefined ) {

				// Attempt to get data from the cache
				// The key will always be camelCased in Data
				data = dataUser.get( elem, key );
				if ( data !== undefined ) {
					return data;
				}

				// Attempt to "discover" the data in
				// HTML5 custom data-* attrs
				data = dataAttr( elem, key );
				if ( data !== undefined ) {
					return data;
				}

				// We tried really hard, but the data doesn't exist.
				return;
			}

			// Set the data...
			this.each( function() {

				// We always store the camelCased key
				dataUser.set( this, key, value );
			} );
		}, null, value, arguments.length > 1, null, true );
	},

	removeData: function( key ) {
		return this.each( function() {
			dataUser.remove( this, key );
		} );
	}
} );


jQuery.extend( {
	queue: function( elem, type, data ) {
		var queue;

		if ( elem ) {
			type = ( type || "fx" ) + "queue";
			queue = dataPriv.get( elem, type );

			// Speed up dequeue by getting out quickly if this is just a lookup
			if ( data ) {
				if ( !queue || Array.isArray( data ) ) {
					queue = dataPriv.access( elem, type, jQuery.makeArray( data ) );
				} else {
					queue.push( data );
				}
			}
			return queue || [];
		}
	},

	dequeue: function( elem, type ) {
		type = type || "fx";

		var queue = jQuery.queue( elem, type ),
			startLength = queue.length,
			fn = queue.shift(),
			hooks = jQuery._queueHooks( elem, type ),
			next = function() {
				jQuery.dequeue( elem, type );
			};

		// If the fx queue is dequeued, always remove the progress sentinel
		if ( fn === "inprogress" ) {
			fn = queue.shift();
			startLength--;
		}

		if ( fn ) {

			// Add a progress sentinel to prevent the fx queue from being
			// automatically dequeued
			if ( type === "fx" ) {
				queue.unshift( "inprogress" );
			}

			// Clear up the last queue stop function
			delete hooks.stop;
			fn.call( elem, next, hooks );
		}

		if ( !startLength && hooks ) {
			hooks.empty.fire();
		}
	},

	// Not public - generate a queueHooks object, or return the current one
	_queueHooks: function( elem, type ) {
		var key = type + "queueHooks";
		return dataPriv.get( elem, key ) || dataPriv.access( elem, key, {
			empty: jQuery.Callbacks( "once memory" ).add( function() {
				dataPriv.remove( elem, [ type + "queue", key ] );
			} )
		} );
	}
} );

jQuery.fn.extend( {
	queue: function( type, data ) {
		var setter = 2;

		if ( typeof type !== "string" ) {
			data = type;
			type = "fx";
			setter--;
		}

		if ( arguments.length < setter ) {
			return jQuery.queue( this[ 0 ], type );
		}

		return data === undefined ?
			this :
			this.each( function() {
				var queue = jQuery.queue( this, type, data );

				// Ensure a hooks for this queue
				jQuery._queueHooks( this, type );

				if ( type === "fx" && queue[ 0 ] !== "inprogress" ) {
					jQuery.dequeue( this, type );
				}
			} );
	},
	dequeue: function( type ) {
		return this.each( function() {
			jQuery.dequeue( this, type );
		} );
	},
	clearQueue: function( type ) {
		return this.queue( type || "fx", [] );
	},

	// Get a promise resolved when queues of a certain type
	// are emptied (fx is the type by default)
	promise: function( type, obj ) {
		var tmp,
			count = 1,
			defer = jQuery.Deferred(),
			elements = this,
			i = this.length,
			resolve = function() {
				if ( !( --count ) ) {
					defer.resolveWith( elements, [ elements ] );
				}
			};

		if ( typeof type !== "string" ) {
			obj = type;
			type = undefined;
		}
		type = type || "fx";

		while ( i-- ) {
			tmp = dataPriv.get( elements[ i ], type + "queueHooks" );
			if ( tmp && tmp.empty ) {
				count++;
				tmp.empty.add( resolve );
			}
		}
		resolve();
		return defer.promise( obj );
	}
} );
var pnum = ( /[+-]?(?:\d*\.|)\d+(?:[eE][+-]?\d+|)/ ).source;

var rcssNum = new RegExp( "^(?:([+-])=|)(" + pnum + ")([a-z%]*)$", "i" );


var cssExpand = [ "Top", "Right", "Bottom", "Left" ];

var documentElement = document.documentElement;



	var isAttached = function( elem ) {
			return jQuery.contains( elem.ownerDocument, elem );
		},
		composed = { composed: true };

	// Support: IE 9 - 11+, Edge 12 - 18+, iOS 10.0 - 10.2 only
	// Check attachment across shadow DOM boundaries when possible (gh-3504)
	// Support: iOS 10.0-10.2 only
	// Early iOS 10 versions support `attachShadow` but not `getRootNode`,
	// leading to errors. We need to check for `getRootNode`.
	if ( documentElement.getRootNode ) {
		isAttached = function( elem ) {
			return jQuery.contains( elem.ownerDocument, elem ) ||
				elem.getRootNode( composed ) === elem.ownerDocument;
		};
	}
var isHiddenWithinTree = function( elem, el ) {

		// isHiddenWithinTree might be called from jQuery#filter function;
		// in that case, element will be second argument
		elem = el || elem;

		// Inline style trumps all
		return elem.style.display === "none" ||
			elem.style.display === "" &&

			// Otherwise, check computed style
			// Support: Firefox <=43 - 45
			// Disconnected elements can have computed display: none, so first confirm that elem is
			// in the document.
			isAttached( elem ) &&

			jQuery.css( elem, "display" ) === "none";
	};



function adjustCSS( elem, prop, valueParts, tween ) {
	var adjusted, scale,
		maxIterations = 20,
		currentValue = tween ?
			function() {
				return tween.cur();
			} :
			function() {
				return jQuery.css( elem, prop, "" );
			},
		initial = currentValue(),
		unit = valueParts && valueParts[ 3 ] || ( jQuery.cssNumber[ prop ] ? "" : "px" ),

		// Starting value computation is required for potential unit mismatches
		initialInUnit = elem.nodeType &&
			( jQuery.cssNumber[ prop ] || unit !== "px" && +initial ) &&
			rcssNum.exec( jQuery.css( elem, prop ) );

	if ( initialInUnit && initialInUnit[ 3 ] !== unit ) {

		// Support: Firefox <=54
		// Halve the iteration target value to prevent interference from CSS upper bounds (gh-2144)
		initial = initial / 2;

		// Trust units reported by jQuery.css
		unit = unit || initialInUnit[ 3 ];

		// Iteratively approximate from a nonzero starting point
		initialInUnit = +initial || 1;

		while ( maxIterations-- ) {

			// Evaluate and update our best guess (doubling guesses that zero out).
			// Finish if the scale equals or crosses 1 (making the old*new product non-positive).
			jQuery.style( elem, prop, initialInUnit + unit );
			if ( ( 1 - scale ) * ( 1 - ( scale = currentValue() / initial || 0.5 ) ) <= 0 ) {
				maxIterations = 0;
			}
			initialInUnit = initialInUnit / scale;

		}

		initialInUnit = initialInUnit * 2;
		jQuery.style( elem, prop, initialInUnit + unit );

		// Make sure we update the tween properties later on
		valueParts = valueParts || [];
	}

	if ( valueParts ) {
		initialInUnit = +initialInUnit || +initial || 0;

		// Apply relative offset (+=/-=) if specified
		adjusted = valueParts[ 1 ] ?
			initialInUnit + ( valueParts[ 1 ] + 1 ) * valueParts[ 2 ] :
			+valueParts[ 2 ];
		if ( tween ) {
			tween.unit = unit;
			tween.start = initialInUnit;
			tween.end = adjusted;
		}
	}
	return adjusted;
}


var defaultDisplayMap = {};

function getDefaultDisplay( elem ) {
	var temp,
		doc = elem.ownerDocument,
		nodeName = elem.nodeName,
		display = defaultDisplayMap[ nodeName ];

	if ( display ) {
		return display;
	}

	temp = doc.body.appendChild( doc.createElement( nodeName ) );
	display = jQuery.css( temp, "display" );

	temp.parentNode.removeChild( temp );

	if ( display === "none" ) {
		display = "block";
	}
	defaultDisplayMap[ nodeName ] = display;

	return display;
}

function showHide( elements, show ) {
	var display, elem,
		values = [],
		index = 0,
		length = elements.length;

	// Determine new display value for elements that need to change
	for ( ; index < length; index++ ) {
		elem = elements[ index ];
		if ( !elem.style ) {
			continue;
		}

		display = elem.style.display;
		if ( show ) {

			// Since we force visibility upon cascade-hidden elements, an immediate (and slow)
			// check is required in this first loop unless we have a nonempty display value (either
			// inline or about-to-be-restored)
			if ( display === "none" ) {
				values[ index ] = dataPriv.get( elem, "display" ) || null;
				if ( !values[ index ] ) {
					elem.style.display = "";
				}
			}
			if ( elem.style.display === "" && isHiddenWithinTree( elem ) ) {
				values[ index ] = getDefaultDisplay( elem );
			}
		} else {
			if ( display !== "none" ) {
				values[ index ] = "none";

				// Remember what we're overwriting
				dataPriv.set( elem, "display", display );
			}
		}
	}

	// Set the display of the elements in a second loop to avoid constant reflow
	for ( index = 0; index < length; index++ ) {
		if ( values[ index ] != null ) {
			elements[ index ].style.display = values[ index ];
		}
	}

	return elements;
}

jQuery.fn.extend( {
	show: function() {
		return showHide( this, true );
	},
	hide: function() {
		return showHide( this );
	},
	toggle: function( state ) {
		if ( typeof state === "boolean" ) {
			return state ? this.show() : this.hide();
		}

		return this.each( function() {
			if ( isHiddenWithinTree( this ) ) {
				jQuery( this ).show();
			} else {
				jQuery( this ).hide();
			}
		} );
	}
} );
var rcheckableType = ( /^(?:checkbox|radio)$/i );

var rtagName = ( /<([a-z][^\/\0>\x20\t\r\n\f]*)/i );

var rscriptType = ( /^$|^module$|\/(?:java|ecma)script/i );



( function() {
	var fragment = document.createDocumentFragment(),
		div = fragment.appendChild( document.createElement( "div" ) ),
		input = document.createElement( "input" );

	// Support: Android 4.0 - 4.3 only
	// Check state lost if the name is set (#11217)
	// Support: Windows Web Apps (WWA)
	// `name` and `type` must use .setAttribute for WWA (#14901)
	input.setAttribute( "type", "radio" );
	input.setAttribute( "checked", "checked" );
	input.setAttribute( "name", "t" );

	div.appendChild( input );

	// Support: Android <=4.1 only
	// Older WebKit doesn't clone checked state correctly in fragments
	support.checkClone = div.cloneNode( true ).cloneNode( true ).lastChild.checked;

	// Support: IE <=11 only
	// Make sure textarea (and checkbox) defaultValue is properly cloned
	div.innerHTML = "<textarea>x</textarea>";
	support.noCloneChecked = !!div.cloneNode( true ).lastChild.defaultValue;

	// Support: IE <=9 only
	// IE <=9 replaces <option> tags with their contents when inserted outside of
	// the select element.
	div.innerHTML = "<option></option>";
	support.option = !!div.lastChild;
} )();


// We have to close these tags to support XHTML (#13200)
var wrapMap = {

	// XHTML parsers do not magically insert elements in the
	// same way that tag soup parsers do. So we cannot shorten
	// this by omitting <tbody> or other required elements.
	thead: [ 1, "<table>", "</table>" ],
	col: [ 2, "<table><colgroup>", "</colgroup></table>" ],
	tr: [ 2, "<table><tbody>", "</tbody></table>" ],
	td: [ 3, "<table><tbody><tr>", "</tr></tbody></table>" ],

	_default: [ 0, "", "" ]
};

wrapMap.tbody = wrapMap.tfoot = wrapMap.colgroup = wrapMap.caption = wrapMap.thead;
wrapMap.th = wrapMap.td;

// Support: IE <=9 only
if ( !support.option ) {
	wrapMap.optgroup = wrapMap.option = [ 1, "<select multiple='multiple'>", "</select>" ];
}


function getAll( context, tag ) {

	// Support: IE <=9 - 11 only
	// Use typeof to avoid zero-argument method invocation on host objects (#15151)
	var ret;

	if ( typeof context.getElementsByTagName !== "undefined" ) {
		ret = context.getElementsByTagName( tag || "*" );

	} else if ( typeof context.querySelectorAll !== "undefined" ) {
		ret = context.querySelectorAll( tag || "*" );

	} else {
		ret = [];
	}

	if ( tag === undefined || tag && nodeName( context, tag ) ) {
		return jQuery.merge( [ context ], ret );
	}

	return ret;
}


// Mark scripts as having already been evaluated
function setGlobalEval( elems, refElements ) {
	var i = 0,
		l = elems.length;

	for ( ; i < l; i++ ) {
		dataPriv.set(
			elems[ i ],
			"globalEval",
			!refElements || dataPriv.get( refElements[ i ], "globalEval" )
		);
	}
}


var rhtml = /<|&#?\w+;/;

function buildFragment( elems, context, scripts, selection, ignored ) {
	var elem, tmp, tag, wrap, attached, j,
		fragment = context.createDocumentFragment(),
		nodes = [],
		i = 0,
		l = elems.length;

	for ( ; i < l; i++ ) {
		elem = elems[ i ];

		if ( elem || elem === 0 ) {

			// Add nodes directly
			if ( toType( elem ) === "object" ) {

				// Support: Android <=4.0 only, PhantomJS 1 only
				// push.apply(_, arraylike) throws on ancient WebKit
				jQuery.merge( nodes, elem.nodeType ? [ elem ] : elem );

			// Convert non-html into a text node
			} else if ( !rhtml.test( elem ) ) {
				nodes.push( context.createTextNode( elem ) );

			// Convert html into DOM nodes
			} else {
				tmp = tmp || fragment.appendChild( context.createElement( "div" ) );

				// Deserialize a standard representation
				tag = ( rtagName.exec( elem ) || [ "", "" ] )[ 1 ].toLowerCase();
				wrap = wrapMap[ tag ] || wrapMap._default;
				tmp.innerHTML = wrap[ 1 ] + jQuery.htmlPrefilter( elem ) + wrap[ 2 ];

				// Descend through wrappers to the right content
				j = wrap[ 0 ];
				while ( j-- ) {
					tmp = tmp.lastChild;
				}

				// Support: Android <=4.0 only, PhantomJS 1 only
				// push.apply(_, arraylike) throws on ancient WebKit
				jQuery.merge( nodes, tmp.childNodes );

				// Remember the top-level container
				tmp = fragment.firstChild;

				// Ensure the created nodes are orphaned (#12392)
				tmp.textContent = "";
			}
		}
	}

	// Remove wrapper from fragment
	fragment.textContent = "";

	i = 0;
	while ( ( elem = nodes[ i++ ] ) ) {

		// Skip elements already in the context collection (trac-4087)
		if ( selection && jQuery.inArray( elem, selection ) > -1 ) {
			if ( ignored ) {
				ignored.push( elem );
			}
			continue;
		}

		attached = isAttached( elem );

		// Append to fragment
		tmp = getAll( fragment.appendChild( elem ), "script" );

		// Preserve script evaluation history
		if ( attached ) {
			setGlobalEval( tmp );
		}

		// Capture executables
		if ( scripts ) {
			j = 0;
			while ( ( elem = tmp[ j++ ] ) ) {
				if ( rscriptType.test( elem.type || "" ) ) {
					scripts.push( elem );
				}
			}
		}
	}

	return fragment;
}


var rtypenamespace = /^([^.]*)(?:\.(.+)|)/;

function returnTrue() {
	return true;
}

function returnFalse() {
	return false;
}

// Support: IE <=9 - 11+
// focus() and blur() are asynchronous, except when they are no-op.
// So expect focus to be synchronous when the element is already active,
// and blur to be synchronous when the element is not already active.
// (focus and blur are always synchronous in other supported browsers,
// this just defines when we can count on it).
function expectSync( elem, type ) {
	return ( elem === safeActiveElement() ) === ( type === "focus" );
}

// Support: IE <=9 only
// Accessing document.activeElement can throw unexpectedly
// https://bugs.jquery.com/ticket/13393
function safeActiveElement() {
	try {
		return document.activeElement;
	} catch ( err ) { }
}

function on( elem, types, selector, data, fn, one ) {
	var origFn, type;

	// Types can be a map of types/handlers
	if ( typeof types === "object" ) {

		// ( types-Object, selector, data )
		if ( typeof selector !== "string" ) {

			// ( types-Object, data )
			data = data || selector;
			selector = undefined;
		}
		for ( type in types ) {
			on( elem, type, selector, data, types[ type ], one );
		}
		return elem;
	}

	if ( data == null && fn == null ) {

		// ( types, fn )
		fn = selector;
		data = selector = undefined;
	} else if ( fn == null ) {
		if ( typeof selector === "string" ) {

			// ( types, selector, fn )
			fn = data;
			data = undefined;
		} else {

			// ( types, data, fn )
			fn = data;
			data = selector;
			selector = undefined;
		}
	}
	if ( fn === false ) {
		fn = returnFalse;
	} else if ( !fn ) {
		return elem;
	}

	if ( one === 1 ) {
		origFn = fn;
		fn = function( event ) {

			// Can use an empty set, since event contains the info
			jQuery().off( event );
			return origFn.apply( this, arguments );
		};

		// Use same guid so caller can remove using origFn
		fn.guid = origFn.guid || ( origFn.guid = jQuery.guid++ );
	}
	return elem.each( function() {
		jQuery.event.add( this, types, fn, data, selector );
	} );
}

/*
 * Helper functions for managing events -- not part of the public interface.
 * Props to Dean Edwards' addEvent library for many of the ideas.
 */
jQuery.event = {

	global: {},

	add: function( elem, types, handler, data, selector ) {

		var handleObjIn, eventHandle, tmp,
			events, t, handleObj,
			special, handlers, type, namespaces, origType,
			elemData = dataPriv.get( elem );

		// Only attach events to objects that accept data
		if ( !acceptData( elem ) ) {
			return;
		}

		// Caller can pass in an object of custom data in lieu of the handler
		if ( handler.handler ) {
			handleObjIn = handler;
			handler = handleObjIn.handler;
			selector = handleObjIn.selector;
		}

		// Ensure that invalid selectors throw exceptions at attach time
		// Evaluate against documentElement in case elem is a non-element node (e.g., document)
		if ( selector ) {
			jQuery.find.matchesSelector( documentElement, selector );
		}

		// Make sure that the handler has a unique ID, used to find/remove it later
		if ( !handler.guid ) {
			handler.guid = jQuery.guid++;
		}

		// Init the element's event structure and main handler, if this is the first
		if ( !( events = elemData.events ) ) {
			events = elemData.events = Object.create( null );
		}
		if ( !( eventHandle = elemData.handle ) ) {
			eventHandle = elemData.handle = function( e ) {

				// Discard the second event of a jQuery.event.trigger() and
				// when an event is called after a page has unloaded
				return typeof jQuery !== "undefined" && jQuery.event.triggered !== e.type ?
					jQuery.event.dispatch.apply( elem, arguments ) : undefined;
			};
		}

		// Handle multiple events separated by a space
		types = ( types || "" ).match( rnothtmlwhite ) || [ "" ];
		t = types.length;
		while ( t-- ) {
			tmp = rtypenamespace.exec( types[ t ] ) || [];
			type = origType = tmp[ 1 ];
			namespaces = ( tmp[ 2 ] || "" ).split( "." ).sort();

			// There *must* be a type, no attaching namespace-only handlers
			if ( !type ) {
				continue;
			}

			// If event changes its type, use the special event handlers for the changed type
			special = jQuery.event.special[ type ] || {};

			// If selector defined, determine special event api type, otherwise given type
			type = ( selector ? special.delegateType : special.bindType ) || type;

			// Update special based on newly reset type
			special = jQuery.event.special[ type ] || {};

			// handleObj is passed to all event handlers
			handleObj = jQuery.extend( {
				type: type,
				origType: origType,
				data: data,
				handler: handler,
				guid: handler.guid,
				selector: selector,
				needsContext: selector && jQuery.expr.match.needsContext.test( selector ),
				namespace: namespaces.join( "." )
			}, handleObjIn );

			// Init the event handler queue if we're the first
			if ( !( handlers = events[ type ] ) ) {
				handlers = events[ type ] = [];
				handlers.delegateCount = 0;

				// Only use addEventListener if the special events handler returns false
				if ( !special.setup ||
					special.setup.call( elem, data, namespaces, eventHandle ) === false ) {

					if ( elem.addEventListener ) {
						elem.addEventListener( type, eventHandle );
					}
				}
			}

			if ( special.add ) {
				special.add.call( elem, handleObj );

				if ( !handleObj.handler.guid ) {
					handleObj.handler.guid = handler.guid;
				}
			}

			// Add to the element's handler list, delegates in front
			if ( selector ) {
				handlers.splice( handlers.delegateCount++, 0, handleObj );
			} else {
				handlers.push( handleObj );
			}

			// Keep track of which events have ever been used, for event optimization
			jQuery.event.global[ type ] = true;
		}

	},

	// Detach an event or set of events from an element
	remove: function( elem, types, handler, selector, mappedTypes ) {

		var j, origCount, tmp,
			events, t, handleObj,
			special, handlers, type, namespaces, origType,
			elemData = dataPriv.hasData( elem ) && dataPriv.get( elem );

		if ( !elemData || !( events = elemData.events ) ) {
			return;
		}

		// Once for each type.namespace in types; type may be omitted
		types = ( types || "" ).match( rnothtmlwhite ) || [ "" ];
		t = types.length;
		while ( t-- ) {
			tmp = rtypenamespace.exec( types[ t ] ) || [];
			type = origType = tmp[ 1 ];
			namespaces = ( tmp[ 2 ] || "" ).split( "." ).sort();

			// Unbind all events (on this namespace, if provided) for the element
			if ( !type ) {
				for ( type in events ) {
					jQuery.event.remove( elem, type + types[ t ], handler, selector, true );
				}
				continue;
			}

			special = jQuery.event.special[ type ] || {};
			type = ( selector ? special.delegateType : special.bindType ) || type;
			handlers = events[ type ] || [];
			tmp = tmp[ 2 ] &&
				new RegExp( "(^|\\.)" + namespaces.join( "\\.(?:.*\\.|)" ) + "(\\.|$)" );

			// Remove matching events
			origCount = j = handlers.length;
			while ( j-- ) {
				handleObj = handlers[ j ];

				if ( ( mappedTypes || origType === handleObj.origType ) &&
					( !handler || handler.guid === handleObj.guid ) &&
					( !tmp || tmp.test( handleObj.namespace ) ) &&
					( !selector || selector === handleObj.selector ||
						selector === "**" && handleObj.selector ) ) {
					handlers.splice( j, 1 );

					if ( handleObj.selector ) {
						handlers.delegateCount--;
					}
					if ( special.remove ) {
						special.remove.call( elem, handleObj );
					}
				}
			}

			// Remove generic event handler if we removed something and no more handlers exist
			// (avoids potential for endless recursion during removal of special event handlers)
			if ( origCount && !handlers.length ) {
				if ( !special.teardown ||
					special.teardown.call( elem, namespaces, elemData.handle ) === false ) {

					jQuery.removeEvent( elem, type, elemData.handle );
				}

				delete events[ type ];
			}
		}

		// Remove data and the expando if it's no longer used
		if ( jQuery.isEmptyObject( events ) ) {
			dataPriv.remove( elem, "handle events" );
		}
	},

	dispatch: function( nativeEvent ) {

		var i, j, ret, matched, handleObj, handlerQueue,
			args = new Array( arguments.length ),

			// Make a writable jQuery.Event from the native event object
			event = jQuery.event.fix( nativeEvent ),

			handlers = (
				dataPriv.get( this, "events" ) || Object.create( null )
			)[ event.type ] || [],
			special = jQuery.event.special[ event.type ] || {};

		// Use the fix-ed jQuery.Event rather than the (read-only) native event
		args[ 0 ] = event;

		for ( i = 1; i < arguments.length; i++ ) {
			args[ i ] = arguments[ i ];
		}

		event.delegateTarget = this;

		// Call the preDispatch hook for the mapped type, and let it bail if desired
		if ( special.preDispatch && special.preDispatch.call( this, event ) === false ) {
			return;
		}

		// Determine handlers
		handlerQueue = jQuery.event.handlers.call( this, event, handlers );

		// Run delegates first; they may want to stop propagation beneath us
		i = 0;
		while ( ( matched = handlerQueue[ i++ ] ) && !event.isPropagationStopped() ) {
			event.currentTarget = matched.elem;

			j = 0;
			while ( ( handleObj = matched.handlers[ j++ ] ) &&
				!event.isImmediatePropagationStopped() ) {

				// If the event is namespaced, then each handler is only invoked if it is
				// specially universal or its namespaces are a superset of the event's.
				if ( !event.rnamespace || handleObj.namespace === false ||
					event.rnamespace.test( handleObj.namespace ) ) {

					event.handleObj = handleObj;
					event.data = handleObj.data;

					ret = ( ( jQuery.event.special[ handleObj.origType ] || {} ).handle ||
						handleObj.handler ).apply( matched.elem, args );

					if ( ret !== undefined ) {
						if ( ( event.result = ret ) === false ) {
							event.preventDefault();
							event.stopPropagation();
						}
					}
				}
			}
		}

		// Call the postDispatch hook for the mapped type
		if ( special.postDispatch ) {
			special.postDispatch.call( this, event );
		}

		return event.result;
	},

	handlers: function( event, handlers ) {
		var i, handleObj, sel, matchedHandlers, matchedSelectors,
			handlerQueue = [],
			delegateCount = handlers.delegateCount,
			cur = event.target;

		// Find delegate handlers
		if ( delegateCount &&

			// Support: IE <=9
			// Black-hole SVG <use> instance trees (trac-13180)
			cur.nodeType &&

			// Support: Firefox <=42
			// Suppress spec-violating clicks indicating a non-primary pointer button (trac-3861)
			// https://www.w3.org/TR/DOM-Level-3-Events/#event-type-click
			// Support: IE 11 only
			// ...but not arrow key "clicks" of radio inputs, which can have `button` -1 (gh-2343)
			!( event.type === "click" && event.button >= 1 ) ) {

			for ( ; cur !== this; cur = cur.parentNode || this ) {

				// Don't check non-elements (#13208)
				// Don't process clicks on disabled elements (#6911, #8165, #11382, #11764)
				if ( cur.nodeType === 1 && !( event.type === "click" && cur.disabled === true ) ) {
					matchedHandlers = [];
					matchedSelectors = {};
					for ( i = 0; i < delegateCount; i++ ) {
						handleObj = handlers[ i ];

						// Don't conflict with Object.prototype properties (#13203)
						sel = handleObj.selector + " ";

						if ( matchedSelectors[ sel ] === undefined ) {
							matchedSelectors[ sel ] = handleObj.needsContext ?
								jQuery( sel, this ).index( cur ) > -1 :
								jQuery.find( sel, this, null, [ cur ] ).length;
						}
						if ( matchedSelectors[ sel ] ) {
							matchedHandlers.push( handleObj );
						}
					}
					if ( matchedHandlers.length ) {
						handlerQueue.push( { elem: cur, handlers: matchedHandlers } );
					}
				}
			}
		}

		// Add the remaining (directly-bound) handlers
		cur = this;
		if ( delegateCount < handlers.length ) {
			handlerQueue.push( { elem: cur, handlers: handlers.slice( delegateCount ) } );
		}

		return handlerQueue;
	},

	addProp: function( name, hook ) {
		Object.defineProperty( jQuery.Event.prototype, name, {
			enumerable: true,
			configurable: true,

			get: isFunction( hook ) ?
				function() {
					if ( this.originalEvent ) {
						return hook( this.originalEvent );
					}
				} :
				function() {
					if ( this.originalEvent ) {
						return this.originalEvent[ name ];
					}
				},

			set: function( value ) {
				Object.defineProperty( this, name, {
					enumerable: true,
					configurable: true,
					writable: true,
					value: value
				} );
			}
		} );
	},

	fix: function( originalEvent ) {
		return originalEvent[ jQuery.expando ] ?
			originalEvent :
			new jQuery.Event( originalEvent );
	},

	special: {
		load: {

			// Prevent triggered image.load events from bubbling to window.load
			noBubble: true
		},
		click: {

			// Utilize native event to ensure correct state for checkable inputs
			setup: function( data ) {

				// For mutual compressibility with _default, replace `this` access with a local var.
				// `|| data` is dead code meant only to preserve the variable through minification.
				var el = this || data;

				// Claim the first handler
				if ( rcheckableType.test( el.type ) &&
					el.click && nodeName( el, "input" ) ) {

					// dataPriv.set( el, "click", ... )
					leverageNative( el, "click", returnTrue );
				}

				// Return false to allow normal processing in the caller
				return false;
			},
			trigger: function( data ) {

				// For mutual compressibility with _default, replace `this` access with a local var.
				// `|| data` is dead code meant only to preserve the variable through minification.
				var el = this || data;

				// Force setup before triggering a click
				if ( rcheckableType.test( el.type ) &&
					el.click && nodeName( el, "input" ) ) {

					leverageNative( el, "click" );
				}

				// Return non-false to allow normal event-path propagation
				return true;
			},

			// For cross-browser consistency, suppress native .click() on links
			// Also prevent it if we're currently inside a leveraged native-event stack
			_default: function( event ) {
				var target = event.target;
				return rcheckableType.test( target.type ) &&
					target.click && nodeName( target, "input" ) &&
					dataPriv.get( target, "click" ) ||
					nodeName( target, "a" );
			}
		},

		beforeunload: {
			postDispatch: function( event ) {

				// Support: Firefox 20+
				// Firefox doesn't alert if the returnValue field is not set.
				if ( event.result !== undefined && event.originalEvent ) {
					event.originalEvent.returnValue = event.result;
				}
			}
		}
	}
};

// Ensure the presence of an event listener that handles manually-triggered
// synthetic events by interrupting progress until reinvoked in response to
// *native* events that it fires directly, ensuring that state changes have
// already occurred before other listeners are invoked.
function leverageNative( el, type, expectSync ) {

	// Missing expectSync indicates a trigger call, which must force setup through jQuery.event.add
	if ( !expectSync ) {
		if ( dataPriv.get( el, type ) === undefined ) {
			jQuery.event.add( el, type, returnTrue );
		}
		return;
	}

	// Register the controller as a special universal handler for all event namespaces
	dataPriv.set( el, type, false );
	jQuery.event.add( el, type, {
		namespace: false,
		handler: function( event ) {
			var notAsync, result,
				saved = dataPriv.get( this, type );

			if ( ( event.isTrigger & 1 ) && this[ type ] ) {

				// Interrupt processing of the outer synthetic .trigger()ed event
				// Saved data should be false in such cases, but might be a leftover capture object
				// from an async native handler (gh-4350)
				if ( !saved.length ) {

					// Store arguments for use when handling the inner native event
					// There will always be at least one argument (an event object), so this array
					// will not be confused with a leftover capture object.
					saved = slice.call( arguments );
					dataPriv.set( this, type, saved );

					// Trigger the native event and capture its result
					// Support: IE <=9 - 11+
					// focus() and blur() are asynchronous
					notAsync = expectSync( this, type );
					this[ type ]();
					result = dataPriv.get( this, type );
					if ( saved !== result || notAsync ) {
						dataPriv.set( this, type, false );
					} else {
						result = {};
					}
					if ( saved !== result ) {

						// Cancel the outer synthetic event
						event.stopImmediatePropagation();
						event.preventDefault();

						// Support: Chrome 86+
						// In Chrome, if an element having a focusout handler is blurred by
						// clicking outside of it, it invokes the handler synchronously. If
						// that handler calls `.remove()` on the element, the data is cleared,
						// leaving `result` undefined. We need to guard against this.
						return result && result.value;
					}

				// If this is an inner synthetic event for an event with a bubbling surrogate
				// (focus or blur), assume that the surrogate already propagated from triggering the
				// native event and prevent that from happening again here.
				// This technically gets the ordering wrong w.r.t. to `.trigger()` (in which the
				// bubbling surrogate propagates *after* the non-bubbling base), but that seems
				// less bad than duplication.
				} else if ( ( jQuery.event.special[ type ] || {} ).delegateType ) {
					event.stopPropagation();
				}

			// If this is a native event triggered above, everything is now in order
			// Fire an inner synthetic event with the original arguments
			} else if ( saved.length ) {

				// ...and capture the result
				dataPriv.set( this, type, {
					value: jQuery.event.trigger(

						// Support: IE <=9 - 11+
						// Extend with the prototype to reset the above stopImmediatePropagation()
						jQuery.extend( saved[ 0 ], jQuery.Event.prototype ),
						saved.slice( 1 ),
						this
					)
				} );

				// Abort handling of the native event
				event.stopImmediatePropagation();
			}
		}
	} );
}

jQuery.removeEvent = function( elem, type, handle ) {

	// This "if" is needed for plain objects
	if ( elem.removeEventListener ) {
		elem.removeEventListener( type, handle );
	}
};

jQuery.Event = function( src, props ) {

	// Allow instantiation without the 'new' keyword
	if ( !( this instanceof jQuery.Event ) ) {
		return new jQuery.Event( src, props );
	}

	// Event object
	if ( src && src.type ) {
		this.originalEvent = src;
		this.type = src.type;

		// Events bubbling up the document may have been marked as prevented
		// by a handler lower down the tree; reflect the correct value.
		this.isDefaultPrevented = src.defaultPrevented ||
				src.defaultPrevented === undefined &&

				// Support: Android <=2.3 only
				src.returnValue === false ?
			returnTrue :
			returnFalse;

		// Create target properties
		// Support: Safari <=6 - 7 only
		// Target should not be a text node (#504, #13143)
		this.target = ( src.target && src.target.nodeType === 3 ) ?
			src.target.parentNode :
			src.target;

		this.currentTarget = src.currentTarget;
		this.relatedTarget = src.relatedTarget;

	// Event type
	} else {
		this.type = src;
	}

	// Put explicitly provided properties onto the event object
	if ( props ) {
		jQuery.extend( this, props );
	}

	// Create a timestamp if incoming event doesn't have one
	this.timeStamp = src && src.timeStamp || Date.now();

	// Mark it as fixed
	this[ jQuery.expando ] = true;
};

// jQuery.Event is based on DOM3 Events as specified by the ECMAScript Language Binding
// https://www.w3.org/TR/2003/WD-DOM-Level-3-Events-20030331/ecma-script-binding.html
jQuery.Event.prototype = {
	constructor: jQuery.Event,
	isDefaultPrevented: returnFalse,
	isPropagationStopped: returnFalse,
	isImmediatePropagationStopped: returnFalse,
	isSimulated: false,

	preventDefault: function() {
		var e = this.originalEvent;

		this.isDefaultPrevented = returnTrue;

		if ( e && !this.isSimulated ) {
			e.preventDefault();
		}
	},
	stopPropagation: function() {
		var e = this.originalEvent;

		this.isPropagationStopped = returnTrue;

		if ( e && !this.isSimulated ) {
			e.stopPropagation();
		}
	},
	stopImmediatePropagation: function() {
		var e = this.originalEvent;

		this.isImmediatePropagationStopped = returnTrue;

		if ( e && !this.isSimulated ) {
			e.stopImmediatePropagation();
		}

		this.stopPropagation();
	}
};

// Includes all common event props including KeyEvent and MouseEvent specific props
jQuery.each( {
	altKey: true,
	bubbles: true,
	cancelable: true,
	changedTouches: true,
	ctrlKey: true,
	detail: true,
	eventPhase: true,
	metaKey: true,
	pageX: true,
	pageY: true,
	shiftKey: true,
	view: true,
	"char": true,
	code: true,
	charCode: true,
	key: true,
	keyCode: true,
	button: true,
	buttons: true,
	clientX: true,
	clientY: true,
	offsetX: true,
	offsetY: true,
	pointerId: true,
	pointerType: true,
	screenX: true,
	screenY: true,
	targetTouches: true,
	toElement: true,
	touches: true,
	which: true
}, jQuery.event.addProp );

jQuery.each( { focus: "focusin", blur: "focusout" }, function( type, delegateType ) {
	jQuery.event.special[ type ] = {

		// Utilize native event if possible so blur/focus sequence is correct
		setup: function() {

			// Claim the first handler
			// dataPriv.set( this, "focus", ... )
			// dataPriv.set( this, "blur", ... )
			leverageNative( this, type, expectSync );

			// Return false to allow normal processing in the caller
			return false;
		},
		trigger: function() {

			// Force setup before trigger
			leverageNative( this, type );

			// Return non-false to allow normal event-path propagation
			return true;
		},

		// Suppress native focus or blur as it's already being fired
		// in leverageNative.
		_default: function() {
			return true;
		},

		delegateType: delegateType
	};
} );

// Create mouseenter/leave events using mouseover/out and event-time checks
// so that event delegation works in jQuery.
// Do the same for pointerenter/pointerleave and pointerover/pointerout
//
// Support: Safari 7 only
// Safari sends mouseenter too often; see:
// https://bugs.chromium.org/p/chromium/issues/detail?id=470258
// for the description of the bug (it existed in older Chrome versions as well).
jQuery.each( {
	mouseenter: "mouseover",
	mouseleave: "mouseout",
	pointerenter: "pointerover",
	pointerleave: "pointerout"
}, function( orig, fix ) {
	jQuery.event.special[ orig ] = {
		delegateType: fix,
		bindType: fix,

		handle: function( event ) {
			var ret,
				target = this,
				related = event.relatedTarget,
				handleObj = event.handleObj;

			// For mouseenter/leave call the handler if related is outside the target.
			// NB: No relatedTarget if the mouse left/entered the browser window
			if ( !related || ( related !== target && !jQuery.contains( target, related ) ) ) {
				event.type = handleObj.origType;
				ret = handleObj.handler.apply( this, arguments );
				event.type = fix;
			}
			return ret;
		}
	};
} );

jQuery.fn.extend( {

	on: function( types, selector, data, fn ) {
		return on( this, types, selector, data, fn );
	},
	one: function( types, selector, data, fn ) {
		return on( this, types, selector, data, fn, 1 );
	},
	off: function( types, selector, fn ) {
		var handleObj, type;
		if ( types && types.preventDefault && types.handleObj ) {

			// ( event )  dispatched jQuery.Event
			handleObj = types.handleObj;
			jQuery( types.delegateTarget ).off(
				handleObj.namespace ?
					handleObj.origType + "." + handleObj.namespace :
					handleObj.origType,
				handleObj.selector,
				handleObj.handler
			);
			return this;
		}
		if ( typeof types === "object" ) {

			// ( types-object [, selector] )
			for ( type in types ) {
				this.off( type, selector, types[ type ] );
			}
			return this;
		}
		if ( selector === false || typeof selector === "function" ) {

			// ( types [, fn] )
			fn = selector;
			selector = undefined;
		}
		if ( fn === false ) {
			fn = returnFalse;
		}
		return this.each( function() {
			jQuery.event.remove( this, types, fn, selector );
		} );
	}
} );


var

	// Support: IE <=10 - 11, Edge 12 - 13 only
	// In IE/Edge using regex groups here causes severe slowdowns.
	// See https://connect.microsoft.com/IE/feedback/details/1736512/
	rnoInnerhtml = /<script|<style|<link/i,

	// checked="checked" or checked
	rchecked = /checked\s*(?:[^=]|=\s*.checked.)/i,
	rcleanScript = /^\s*<!(?:\[CDATA\[|--)|(?:\]\]|--)>\s*$/g;

// Prefer a tbody over its parent table for containing new rows
function manipulationTarget( elem, content ) {
	if ( nodeName( elem, "table" ) &&
		nodeName( content.nodeType !== 11 ? content : content.firstChild, "tr" ) ) {

		return jQuery( elem ).children( "tbody" )[ 0 ] || elem;
	}

	return elem;
}

// Replace/restore the type attribute of script elements for safe DOM manipulation
function disableScript( elem ) {
	elem.type = ( elem.getAttribute( "type" ) !== null ) + "/" + elem.type;
	return elem;
}
function restoreScript( elem ) {
	if ( ( elem.type || "" ).slice( 0, 5 ) === "true/" ) {
		elem.type = elem.type.slice( 5 );
	} else {
		elem.removeAttribute( "type" );
	}

	return elem;
}

function cloneCopyEvent( src, dest ) {
	var i, l, type, pdataOld, udataOld, udataCur, events;

	if ( dest.nodeType !== 1 ) {
		return;
	}

	// 1. Copy private data: events, handlers, etc.
	if ( dataPriv.hasData( src ) ) {
		pdataOld = dataPriv.get( src );
		events = pdataOld.events;

		if ( events ) {
			dataPriv.remove( dest, "handle events" );

			for ( type in events ) {
				for ( i = 0, l = events[ type ].length; i < l; i++ ) {
					jQuery.event.add( dest, type, events[ type ][ i ] );
				}
			}
		}
	}

	// 2. Copy user data
	if ( dataUser.hasData( src ) ) {
		udataOld = dataUser.access( src );
		udataCur = jQuery.extend( {}, udataOld );

		dataUser.set( dest, udataCur );
	}
}

// Fix IE bugs, see support tests
function fixInput( src, dest ) {
	var nodeName = dest.nodeName.toLowerCase();

	// Fails to persist the checked state of a cloned checkbox or radio button.
	if ( nodeName === "input" && rcheckableType.test( src.type ) ) {
		dest.checked = src.checked;

	// Fails to return the selected option to the default selected state when cloning options
	} else if ( nodeName === "input" || nodeName === "textarea" ) {
		dest.defaultValue = src.defaultValue;
	}
}

function domManip( collection, args, callback, ignored ) {

	// Flatten any nested arrays
	args = flat( args );

	var fragment, first, scripts, hasScripts, node, doc,
		i = 0,
		l = collection.length,
		iNoClone = l - 1,
		value = args[ 0 ],
		valueIsFunction = isFunction( value );

	// We can't cloneNode fragments that contain checked, in WebKit
	if ( valueIsFunction ||
			( l > 1 && typeof value === "string" &&
				!support.checkClone && rchecked.test( value ) ) ) {
		return collection.each( function( index ) {
			var self = collection.eq( index );
			if ( valueIsFunction ) {
				args[ 0 ] = value.call( this, index, self.html() );
			}
			domManip( self, args, callback, ignored );
		} );
	}

	if ( l ) {
		fragment = buildFragment( args, collection[ 0 ].ownerDocument, false, collection, ignored );
		first = fragment.firstChild;

		if ( fragment.childNodes.length === 1 ) {
			fragment = first;
		}

		// Require either new content or an interest in ignored elements to invoke the callback
		if ( first || ignored ) {
			scripts = jQuery.map( getAll( fragment, "script" ), disableScript );
			hasScripts = scripts.length;

			// Use the original fragment for the last item
			// instead of the first because it can end up
			// being emptied incorrectly in certain situations (#8070).
			for ( ; i < l; i++ ) {
				node = fragment;

				if ( i !== iNoClone ) {
					node = jQuery.clone( node, true, true );

					// Keep references to cloned scripts for later restoration
					if ( hasScripts ) {

						// Support: Android <=4.0 only, PhantomJS 1 only
						// push.apply(_, arraylike) throws on ancient WebKit
						jQuery.merge( scripts, getAll( node, "script" ) );
					}
				}

				callback.call( collection[ i ], node, i );
			}

			if ( hasScripts ) {
				doc = scripts[ scripts.length - 1 ].ownerDocument;

				// Reenable scripts
				jQuery.map( scripts, restoreScript );

				// Evaluate executable scripts on first document insertion
				for ( i = 0; i < hasScripts; i++ ) {
					node = scripts[ i ];
					if ( rscriptType.test( node.type || "" ) &&
						!dataPriv.access( node, "globalEval" ) &&
						jQuery.contains( doc, node ) ) {

						if ( node.src && ( node.type || "" ).toLowerCase()  !== "module" ) {

							// Optional AJAX dependency, but won't run scripts if not present
							if ( jQuery._evalUrl && !node.noModule ) {
								jQuery._evalUrl( node.src, {
									nonce: node.nonce || node.getAttribute( "nonce" )
								}, doc );
							}
						} else {
							DOMEval( node.textContent.replace( rcleanScript, "" ), node, doc );
						}
					}
				}
			}
		}
	}

	return collection;
}

function remove( elem, selector, keepData ) {
	var node,
		nodes = selector ? jQuery.filter( selector, elem ) : elem,
		i = 0;

	for ( ; ( node = nodes[ i ] ) != null; i++ ) {
		if ( !keepData && node.nodeType === 1 ) {
			jQuery.cleanData( getAll( node ) );
		}

		if ( node.parentNode ) {
			if ( keepData && isAttached( node ) ) {
				setGlobalEval( getAll( node, "script" ) );
			}
			node.parentNode.removeChild( node );
		}
	}

	return elem;
}

jQuery.extend( {
	htmlPrefilter: function( html ) {
		return html;
	},

	clone: function( elem, dataAndEvents, deepDataAndEvents ) {
		var i, l, srcElements, destElements,
			clone = elem.cloneNode( true ),
			inPage = isAttached( elem );

		// Fix IE cloning issues
		if ( !support.noCloneChecked && ( elem.nodeType === 1 || elem.nodeType === 11 ) &&
				!jQuery.isXMLDoc( elem ) ) {

			// We eschew Sizzle here for performance reasons: https://jsperf.com/getall-vs-sizzle/2
			destElements = getAll( clone );
			srcElements = getAll( elem );

			for ( i = 0, l = srcElements.length; i < l; i++ ) {
				fixInput( srcElements[ i ], destElements[ i ] );
			}
		}

		// Copy the events from the original to the clone
		if ( dataAndEvents ) {
			if ( deepDataAndEvents ) {
				srcElements = srcElements || getAll( elem );
				destElements = destElements || getAll( clone );

				for ( i = 0, l = srcElements.length; i < l; i++ ) {
					cloneCopyEvent( srcElements[ i ], destElements[ i ] );
				}
			} else {
				cloneCopyEvent( elem, clone );
			}
		}

		// Preserve script evaluation history
		destElements = getAll( clone, "script" );
		if ( destElements.length > 0 ) {
			setGlobalEval( destElements, !inPage && getAll( elem, "script" ) );
		}

		// Return the cloned set
		return clone;
	},

	cleanData: function( elems ) {
		var data, elem, type,
			special = jQuery.event.special,
			i = 0;

		for ( ; ( elem = elems[ i ] ) !== undefined; i++ ) {
			if ( acceptData( elem ) ) {
				if ( ( data = elem[ dataPriv.expando ] ) ) {
					if ( data.events ) {
						for ( type in data.events ) {
							if ( special[ type ] ) {
								jQuery.event.remove( elem, type );

							// This is a shortcut to avoid jQuery.event.remove's overhead
							} else {
								jQuery.removeEvent( elem, type, data.handle );
							}
						}
					}

					// Support: Chrome <=35 - 45+
					// Assign undefined instead of using delete, see Data#remove
					elem[ dataPriv.expando ] = undefined;
				}
				if ( elem[ dataUser.expando ] ) {

					// Support: Chrome <=35 - 45+
					// Assign undefined instead of using delete, see Data#remove
					elem[ dataUser.expando ] = undefined;
				}
			}
		}
	}
} );

jQuery.fn.extend( {
	detach: function( selector ) {
		return remove( this, selector, true );
	},

	remove: function( selector ) {
		return remove( this, selector );
	},

	text: function( value ) {
		return access( this, function( value ) {
			return value === undefined ?
				jQuery.text( this ) :
				this.empty().each( function() {
					if ( this.nodeType === 1 || this.nodeType === 11 || this.nodeType === 9 ) {
						this.textContent = value;
					}
				} );
		}, null, value, arguments.length );
	},

	append: function() {
		return domManip( this, arguments, function( elem ) {
			if ( this.nodeType === 1 || this.nodeType === 11 || this.nodeType === 9 ) {
				var target = manipulationTarget( this, elem );
				target.appendChild( elem );
			}
		} );
	},

	prepend: function() {
		return domManip( this, arguments, function( elem ) {
			if ( this.nodeType === 1 || this.nodeType === 11 || this.nodeType === 9 ) {
				var target = manipulationTarget( this, elem );
				target.insertBefore( elem, target.firstChild );
			}
		} );
	},

	before: function() {
		return domManip( this, arguments, function( elem ) {
			if ( this.parentNode ) {
				this.parentNode.insertBefore( elem, this );
			}
		} );
	},

	after: function() {
		return domManip( this, arguments, function( elem ) {
			if ( this.parentNode ) {
				this.parentNode.insertBefore( elem, this.nextSibling );
			}
		} );
	},

	empty: function() {
		var elem,
			i = 0;

		for ( ; ( elem = this[ i ] ) != null; i++ ) {
			if ( elem.nodeType === 1 ) {

				// Prevent memory leaks
				jQuery.cleanData( getAll( elem, false ) );

				// Remove any remaining nodes
				elem.textContent = "";
			}
		}

		return this;
	},

	clone: function( dataAndEvents, deepDataAndEvents ) {
		dataAndEvents = dataAndEvents == null ? false : dataAndEvents;
		deepDataAndEvents = deepDataAndEvents == null ? dataAndEvents : deepDataAndEvents;

		return this.map( function() {
			return jQuery.clone( this, dataAndEvents, deepDataAndEvents );
		} );
	},

	html: function( value ) {
		return access( this, function( value ) {
			var elem = this[ 0 ] || {},
				i = 0,
				l = this.length;

			if ( value === undefined && elem.nodeType === 1 ) {
				return elem.innerHTML;
			}

			// See if we can take a shortcut and just use innerHTML
			if ( typeof value === "string" && !rnoInnerhtml.test( value ) &&
				!wrapMap[ ( rtagName.exec( value ) || [ "", "" ] )[ 1 ].toLowerCase() ] ) {

				value = jQuery.htmlPrefilter( value );

				try {
					for ( ; i < l; i++ ) {
						elem = this[ i ] || {};

						// Remove element nodes and prevent memory leaks
						if ( elem.nodeType === 1 ) {
							jQuery.cleanData( getAll( elem, false ) );
							elem.innerHTML = value;
						}
					}

					elem = 0;

				// If using innerHTML throws an exception, use the fallback method
				} catch ( e ) {}
			}

			if ( elem ) {
				this.empty().append( value );
			}
		}, null, value, arguments.length );
	},

	replaceWith: function() {
		var ignored = [];

		// Make the changes, replacing each non-ignored context element with the new content
		return domManip( this, arguments, function( elem ) {
			var parent = this.parentNode;

			if ( jQuery.inArray( this, ignored ) < 0 ) {
				jQuery.cleanData( getAll( this ) );
				if ( parent ) {
					parent.replaceChild( elem, this );
				}
			}

		// Force callback invocation
		}, ignored );
	}
} );

jQuery.each( {
	appendTo: "append",
	prependTo: "prepend",
	insertBefore: "before",
	insertAfter: "after",
	replaceAll: "replaceWith"
}, function( name, original ) {
	jQuery.fn[ name ] = function( selector ) {
		var elems,
			ret = [],
			insert = jQuery( selector ),
			last = insert.length - 1,
			i = 0;

		for ( ; i <= last; i++ ) {
			elems = i === last ? this : this.clone( true );
			jQuery( insert[ i ] )[ original ]( elems );

			// Support: Android <=4.0 only, PhantomJS 1 only
			// .get() because push.apply(_, arraylike) throws on ancient WebKit
			push.apply( ret, elems.get() );
		}

		return this.pushStack( ret );
	};
} );
var rnumnonpx = new RegExp( "^(" + pnum + ")(?!px)[a-z%]+$", "i" );

var getStyles = function( elem ) {

		// Support: IE <=11 only, Firefox <=30 (#15098, #14150)
		// IE throws on elements created in popups
		// FF meanwhile throws on frame elements through "defaultView.getComputedStyle"
		var view = elem.ownerDocument.defaultView;

		if ( !view || !view.opener ) {
			view = window;
		}

		return view.getComputedStyle( elem );
	};

var swap = function( elem, options, callback ) {
	var ret, name,
		old = {};

	// Remember the old values, and insert the new ones
	for ( name in options ) {
		old[ name ] = elem.style[ name ];
		elem.style[ name ] = options[ name ];
	}

	ret = callback.call( elem );

	// Revert the old values
	for ( name in options ) {
		elem.style[ name ] = old[ name ];
	}

	return ret;
};


var rboxStyle = new RegExp( cssExpand.join( "|" ), "i" );



( function() {

	// Executing both pixelPosition & boxSizingReliable tests require only one layout
	// so they're executed at the same time to save the second computation.
	function computeStyleTests() {

		// This is a singleton, we need to execute it only once
		if ( !div ) {
			return;
		}

		container.style.cssText = "position:absolute;left:-11111px;width:60px;" +
			"margin-top:1px;padding:0;border:0";
		div.style.cssText =
			"position:relative;display:block;box-sizing:border-box;overflow:scroll;" +
			"margin:auto;border:1px;padding:1px;" +
			"width:60%;top:1%";
		documentElement.appendChild( container ).appendChild( div );

		var divStyle = window.getComputedStyle( div );
		pixelPositionVal = divStyle.top !== "1%";

		// Support: Android 4.0 - 4.3 only, Firefox <=3 - 44
		reliableMarginLeftVal = roundPixelMeasures( divStyle.marginLeft ) === 12;

		// Support: Android 4.0 - 4.3 only, Safari <=9.1 - 10.1, iOS <=7.0 - 9.3
		// Some styles come back with percentage values, even though they shouldn't
		div.style.right = "60%";
		pixelBoxStylesVal = roundPixelMeasures( divStyle.right ) === 36;

		// Support: IE 9 - 11 only
		// Detect misreporting of content dimensions for box-sizing:border-box elements
		boxSizingReliableVal = roundPixelMeasures( divStyle.width ) === 36;

		// Support: IE 9 only
		// Detect overflow:scroll screwiness (gh-3699)
		// Support: Chrome <=64
		// Don't get tricked when zoom affects offsetWidth (gh-4029)
		div.style.position = "absolute";
		scrollboxSizeVal = roundPixelMeasures( div.offsetWidth / 3 ) === 12;

		documentElement.removeChild( container );

		// Nullify the div so it wouldn't be stored in the memory and
		// it will also be a sign that checks already performed
		div = null;
	}

	function roundPixelMeasures( measure ) {
		return Math.round( parseFloat( measure ) );
	}

	var pixelPositionVal, boxSizingReliableVal, scrollboxSizeVal, pixelBoxStylesVal,
		reliableTrDimensionsVal, reliableMarginLeftVal,
		container = document.createElement( "div" ),
		div = document.createElement( "div" );

	// Finish early in limited (non-browser) environments
	if ( !div.style ) {
		return;
	}

	// Support: IE <=9 - 11 only
	// Style of cloned element affects source element cloned (#8908)
	div.style.backgroundClip = "content-box";
	div.cloneNode( true ).style.backgroundClip = "";
	support.clearCloneStyle = div.style.backgroundClip === "content-box";

	jQuery.extend( support, {
		boxSizingReliable: function() {
			computeStyleTests();
			return boxSizingReliableVal;
		},
		pixelBoxStyles: function() {
			computeStyleTests();
			return pixelBoxStylesVal;
		},
		pixelPosition: function() {
			computeStyleTests();
			return pixelPositionVal;
		},
		reliableMarginLeft: function() {
			computeStyleTests();
			return reliableMarginLeftVal;
		},
		scrollboxSize: function() {
			computeStyleTests();
			return scrollboxSizeVal;
		},

		// Support: IE 9 - 11+, Edge 15 - 18+
		// IE/Edge misreport `getComputedStyle` of table rows with width/height
		// set in CSS while `offset*` properties report correct values.
		// Behavior in IE 9 is more subtle than in newer versions & it passes
		// some versions of this test; make sure not to make it pass there!
		//
		// Support: Firefox 70+
		// Only Firefox includes border widths
		// in computed dimensions. (gh-4529)
		reliableTrDimensions: function() {
			var table, tr, trChild, trStyle;
			if ( reliableTrDimensionsVal == null ) {
				table = document.createElement( "table" );
				tr = document.createElement( "tr" );
				trChild = document.createElement( "div" );

				table.style.cssText = "position:absolute;left:-11111px;border-collapse:separate";
				tr.style.cssText = "border:1px solid";

				// Support: Chrome 86+
				// Height set through cssText does not get applied.
				// Computed height then comes back as 0.
				tr.style.height = "1px";
				trChild.style.height = "9px";

				// Support: Android 8 Chrome 86+
				// In our bodyBackground.html iframe,
				// display for all div elements is set to "inline",
				// which causes a problem only in Android 8 Chrome 86.
				// Ensuring the div is display: block
				// gets around this issue.
				trChild.style.display = "block";

				documentElement
					.appendChild( table )
					.appendChild( tr )
					.appendChild( trChild );

				trStyle = window.getComputedStyle( tr );
				reliableTrDimensionsVal = ( parseInt( trStyle.height, 10 ) +
					parseInt( trStyle.borderTopWidth, 10 ) +
					parseInt( trStyle.borderBottomWidth, 10 ) ) === tr.offsetHeight;

				documentElement.removeChild( table );
			}
			return reliableTrDimensionsVal;
		}
	} );
} )();


function curCSS( elem, name, computed ) {
	var width, minWidth, maxWidth, ret,

		// Support: Firefox 51+
		// Retrieving style before computed somehow
		// fixes an issue with getting wrong values
		// on detached elements
		style = elem.style;

	computed = computed || getStyles( elem );

	// getPropertyValue is needed for:
	//   .css('filter') (IE 9 only, #12537)
	//   .css('--customProperty) (#3144)
	if ( computed ) {
		ret = computed.getPropertyValue( name ) || computed[ name ];

		if ( ret === "" && !isAttached( elem ) ) {
			ret = jQuery.style( elem, name );
		}

		// A tribute to the "awesome hack by Dean Edwards"
		// Android Browser returns percentage for some values,
		// but width seems to be reliably pixels.
		// This is against the CSSOM draft spec:
		// https://drafts.csswg.org/cssom/#resolved-values
		if ( !support.pixelBoxStyles() && rnumnonpx.test( ret ) && rboxStyle.test( name ) ) {

			// Remember the original values
			width = style.width;
			minWidth = style.minWidth;
			maxWidth = style.maxWidth;

			// Put in the new values to get a computed value out
			style.minWidth = style.maxWidth = style.width = ret;
			ret = computed.width;

			// Revert the changed values
			style.width = width;
			style.minWidth = minWidth;
			style.maxWidth = maxWidth;
		}
	}

	return ret !== undefined ?

		// Support: IE <=9 - 11 only
		// IE returns zIndex value as an integer.
		ret + "" :
		ret;
}


function addGetHookIf( conditionFn, hookFn ) {

	// Define the hook, we'll check on the first run if it's really needed.
	return {
		get: function() {
			if ( conditionFn() ) {

				// Hook not needed (or it's not possible to use it due
				// to missing dependency), remove it.
				delete this.get;
				return;
			}

			// Hook needed; redefine it so that the support test is not executed again.
			return ( this.get = hookFn ).apply( this, arguments );
		}
	};
}


var cssPrefixes = [ "Webkit", "Moz", "ms" ],
	emptyStyle = document.createElement( "div" ).style,
	vendorProps = {};

// Return a vendor-prefixed property or undefined
function vendorPropName( name ) {

	// Check for vendor prefixed names
	var capName = name[ 0 ].toUpperCase() + name.slice( 1 ),
		i = cssPrefixes.length;

	while ( i-- ) {
		name = cssPrefixes[ i ] + capName;
		if ( name in emptyStyle ) {
			return name;
		}
	}
}

// Return a potentially-mapped jQuery.cssProps or vendor prefixed property
function finalPropName( name ) {
	var final = jQuery.cssProps[ name ] || vendorProps[ name ];

	if ( final ) {
		return final;
	}
	if ( name in emptyStyle ) {
		return name;
	}
	return vendorProps[ name ] = vendorPropName( name ) || name;
}


var

	// Swappable if display is none or starts with table
	// except "table", "table-cell", or "table-caption"
	// See here for display values: https://developer.mozilla.org/en-US/docs/CSS/display
	rdisplayswap = /^(none|table(?!-c[ea]).+)/,
	rcustomProp = /^--/,
	cssShow = { position: "absolute", visibility: "hidden", display: "block" },
	cssNormalTransform = {
		letterSpacing: "0",
		fontWeight: "400"
	};

function setPositiveNumber( _elem, value, subtract ) {

	// Any relative (+/-) values have already been
	// normalized at this point
	var matches = rcssNum.exec( value );
	return matches ?

		// Guard against undefined "subtract", e.g., when used as in cssHooks
		Math.max( 0, matches[ 2 ] - ( subtract || 0 ) ) + ( matches[ 3 ] || "px" ) :
		value;
}

function boxModelAdjustment( elem, dimension, box, isBorderBox, styles, computedVal ) {
	var i = dimension === "width" ? 1 : 0,
		extra = 0,
		delta = 0;

	// Adjustment may not be necessary
	if ( box === ( isBorderBox ? "border" : "content" ) ) {
		return 0;
	}

	for ( ; i < 4; i += 2 ) {

		// Both box models exclude margin
		if ( box === "margin" ) {
			delta += jQuery.css( elem, box + cssExpand[ i ], true, styles );
		}

		// If we get here with a content-box, we're seeking "padding" or "border" or "margin"
		if ( !isBorderBox ) {

			// Add padding
			delta += jQuery.css( elem, "padding" + cssExpand[ i ], true, styles );

			// For "border" or "margin", add border
			if ( box !== "padding" ) {
				delta += jQuery.css( elem, "border" + cssExpand[ i ] + "Width", true, styles );

			// But still keep track of it otherwise
			} else {
				extra += jQuery.css( elem, "border" + cssExpand[ i ] + "Width", true, styles );
			}

		// If we get here with a border-box (content + padding + border), we're seeking "content" or
		// "padding" or "margin"
		} else {

			// For "content", subtract padding
			if ( box === "content" ) {
				delta -= jQuery.css( elem, "padding" + cssExpand[ i ], true, styles );
			}

			// For "content" or "padding", subtract border
			if ( box !== "margin" ) {
				delta -= jQuery.css( elem, "border" + cssExpand[ i ] + "Width", true, styles );
			}
		}
	}

	// Account for positive content-box scroll gutter when requested by providing computedVal
	if ( !isBorderBox && computedVal >= 0 ) {

		// offsetWidth/offsetHeight is a rounded sum of content, padding, scroll gutter, and border
		// Assuming integer scroll gutter, subtract the rest and round down
		delta += Math.max( 0, Math.ceil(
			elem[ "offset" + dimension[ 0 ].toUpperCase() + dimension.slice( 1 ) ] -
			computedVal -
			delta -
			extra -
			0.5

		// If offsetWidth/offsetHeight is unknown, then we can't determine content-box scroll gutter
		// Use an explicit zero to avoid NaN (gh-3964)
		) ) || 0;
	}

	return delta;
}

function getWidthOrHeight( elem, dimension, extra ) {

	// Start with computed style
	var styles = getStyles( elem ),

		// To avoid forcing a reflow, only fetch boxSizing if we need it (gh-4322).
		// Fake content-box until we know it's needed to know the true value.
		boxSizingNeeded = !support.boxSizingReliable() || extra,
		isBorderBox = boxSizingNeeded &&
			jQuery.css( elem, "boxSizing", false, styles ) === "border-box",
		valueIsBorderBox = isBorderBox,

		val = curCSS( elem, dimension, styles ),
		offsetProp = "offset" + dimension[ 0 ].toUpperCase() + dimension.slice( 1 );

	// Support: Firefox <=54
	// Return a confounding non-pixel value or feign ignorance, as appropriate.
	if ( rnumnonpx.test( val ) ) {
		if ( !extra ) {
			return val;
		}
		val = "auto";
	}


	// Support: IE 9 - 11 only
	// Use offsetWidth/offsetHeight for when box sizing is unreliable.
	// In those cases, the computed value can be trusted to be border-box.
	if ( ( !support.boxSizingReliable() && isBorderBox ||

		// Support: IE 10 - 11+, Edge 15 - 18+
		// IE/Edge misreport `getComputedStyle` of table rows with width/height
		// set in CSS while `offset*` properties report correct values.
		// Interestingly, in some cases IE 9 doesn't suffer from this issue.
		!support.reliableTrDimensions() && nodeName( elem, "tr" ) ||

		// Fall back to offsetWidth/offsetHeight when value is "auto"
		// This happens for inline elements with no explicit setting (gh-3571)
		val === "auto" ||

		// Support: Android <=4.1 - 4.3 only
		// Also use offsetWidth/offsetHeight for misreported inline dimensions (gh-3602)
		!parseFloat( val ) && jQuery.css( elem, "display", false, styles ) === "inline" ) &&

		// Make sure the element is visible & connected
		elem.getClientRects().length ) {

		isBorderBox = jQuery.css( elem, "boxSizing", false, styles ) === "border-box";

		// Where available, offsetWidth/offsetHeight approximate border box dimensions.
		// Where not available (e.g., SVG), assume unreliable box-sizing and interpret the
		// retrieved value as a content box dimension.
		valueIsBorderBox = offsetProp in elem;
		if ( valueIsBorderBox ) {
			val = elem[ offsetProp ];
		}
	}

	// Normalize "" and auto
	val = parseFloat( val ) || 0;

	// Adjust for the element's box model
	return ( val +
		boxModelAdjustment(
			elem,
			dimension,
			extra || ( isBorderBox ? "border" : "content" ),
			valueIsBorderBox,
			styles,

			// Provide the current computed size to request scroll gutter calculation (gh-3589)
			val
		)
	) + "px";
}

jQuery.extend( {

	// Add in style property hooks for overriding the default
	// behavior of getting and setting a style property
	cssHooks: {
		opacity: {
			get: function( elem, computed ) {
				if ( computed ) {

					// We should always get a number back from opacity
					var ret = curCSS( elem, "opacity" );
					return ret === "" ? "1" : ret;
				}
			}
		}
	},

	// Don't automatically add "px" to these possibly-unitless properties
	cssNumber: {
		"animationIterationCount": true,
		"columnCount": true,
		"fillOpacity": true,
		"flexGrow": true,
		"flexShrink": true,
		"fontWeight": true,
		"gridArea": true,
		"gridColumn": true,
		"gridColumnEnd": true,
		"gridColumnStart": true,
		"gridRow": true,
		"gridRowEnd": true,
		"gridRowStart": true,
		"lineHeight": true,
		"opacity": true,
		"order": true,
		"orphans": true,
		"widows": true,
		"zIndex": true,
		"zoom": true
	},

	// Add in properties whose names you wish to fix before
	// setting or getting the value
	cssProps: {},

	// Get and set the style property on a DOM Node
	style: function( elem, name, value, extra ) {

		// Don't set styles on text and comment nodes
		if ( !elem || elem.nodeType === 3 || elem.nodeType === 8 || !elem.style ) {
			return;
		}

		// Make sure that we're working with the right name
		var ret, type, hooks,
			origName = camelCase( name ),
			isCustomProp = rcustomProp.test( name ),
			style = elem.style;

		// Make sure that we're working with the right name. We don't
		// want to query the value if it is a CSS custom property
		// since they are user-defined.
		if ( !isCustomProp ) {
			name = finalPropName( origName );
		}

		// Gets hook for the prefixed version, then unprefixed version
		hooks = jQuery.cssHooks[ name ] || jQuery.cssHooks[ origName ];

		// Check if we're setting a value
		if ( value !== undefined ) {
			type = typeof value;

			// Convert "+=" or "-=" to relative numbers (#7345)
			if ( type === "string" && ( ret = rcssNum.exec( value ) ) && ret[ 1 ] ) {
				value = adjustCSS( elem, name, ret );

				// Fixes bug #9237
				type = "number";
			}

			// Make sure that null and NaN values aren't set (#7116)
			if ( value == null || value !== value ) {
				return;
			}

			// If a number was passed in, add the unit (except for certain CSS properties)
			// The isCustomProp check can be removed in jQuery 4.0 when we only auto-append
			// "px" to a few hardcoded values.
			if ( type === "number" && !isCustomProp ) {
				value += ret && ret[ 3 ] || ( jQuery.cssNumber[ origName ] ? "" : "px" );
			}

			// background-* props affect original clone's values
			if ( !support.clearCloneStyle && value === "" && name.indexOf( "background" ) === 0 ) {
				style[ name ] = "inherit";
			}

			// If a hook was provided, use that value, otherwise just set the specified value
			if ( !hooks || !( "set" in hooks ) ||
				( value = hooks.set( elem, value, extra ) ) !== undefined ) {

				if ( isCustomProp ) {
					style.setProperty( name, value );
				} else {
					style[ name ] = value;
				}
			}

		} else {

			// If a hook was provided get the non-computed value from there
			if ( hooks && "get" in hooks &&
				( ret = hooks.get( elem, false, extra ) ) !== undefined ) {

				return ret;
			}

			// Otherwise just get the value from the style object
			return style[ name ];
		}
	},

	css: function( elem, name, extra, styles ) {
		var val, num, hooks,
			origName = camelCase( name ),
			isCustomProp = rcustomProp.test( name );

		// Make sure that we're working with the right name. We don't
		// want to modify the value if it is a CSS custom property
		// since they are user-defined.
		if ( !isCustomProp ) {
			name = finalPropName( origName );
		}

		// Try prefixed name followed by the unprefixed name
		hooks = jQuery.cssHooks[ name ] || jQuery.cssHooks[ origName ];

		// If a hook was provided get the computed value from there
		if ( hooks && "get" in hooks ) {
			val = hooks.get( elem, true, extra );
		}

		// Otherwise, if a way to get the computed value exists, use that
		if ( val === undefined ) {
			val = curCSS( elem, name, styles );
		}

		// Convert "normal" to computed value
		if ( val === "normal" && name in cssNormalTransform ) {
			val = cssNormalTransform[ name ];
		}

		// Make numeric if forced or a qualifier was provided and val looks numeric
		if ( extra === "" || extra ) {
			num = parseFloat( val );
			return extra === true || isFinite( num ) ? num || 0 : val;
		}

		return val;
	}
} );

jQuery.each( [ "height", "width" ], function( _i, dimension ) {
	jQuery.cssHooks[ dimension ] = {
		get: function( elem, computed, extra ) {
			if ( computed ) {

				// Certain elements can have dimension info if we invisibly show them
				// but it must have a current display style that would benefit
				return rdisplayswap.test( jQuery.css( elem, "display" ) ) &&

					// Support: Safari 8+
					// Table columns in Safari have non-zero offsetWidth & zero
					// getBoundingClientRect().width unless display is changed.
					// Support: IE <=11 only
					// Running getBoundingClientRect on a disconnected node
					// in IE throws an error.
					( !elem.getClientRects().length || !elem.getBoundingClientRect().width ) ?
					swap( elem, cssShow, function() {
						return getWidthOrHeight( elem, dimension, extra );
					} ) :
					getWidthOrHeight( elem, dimension, extra );
			}
		},

		set: function( elem, value, extra ) {
			var matches,
				styles = getStyles( elem ),

				// Only read styles.position if the test has a chance to fail
				// to avoid forcing a reflow.
				scrollboxSizeBuggy = !support.scrollboxSize() &&
					styles.position === "absolute",

				// To avoid forcing a reflow, only fetch boxSizing if we need it (gh-3991)
				boxSizingNeeded = scrollboxSizeBuggy || extra,
				isBorderBox = boxSizingNeeded &&
					jQuery.css( elem, "boxSizing", false, styles ) === "border-box",
				subtract = extra ?
					boxModelAdjustment(
						elem,
						dimension,
						extra,
						isBorderBox,
						styles
					) :
					0;

			// Account for unreliable border-box dimensions by comparing offset* to computed and
			// faking a content-box to get border and padding (gh-3699)
			if ( isBorderBox && scrollboxSizeBuggy ) {
				subtract -= Math.ceil(
					elem[ "offset" + dimension[ 0 ].toUpperCase() + dimension.slice( 1 ) ] -
					parseFloat( styles[ dimension ] ) -
					boxModelAdjustment( elem, dimension, "border", false, styles ) -
					0.5
				);
			}

			// Convert to pixels if value adjustment is needed
			if ( subtract && ( matches = rcssNum.exec( value ) ) &&
				( matches[ 3 ] || "px" ) !== "px" ) {

				elem.style[ dimension ] = value;
				value = jQuery.css( elem, dimension );
			}

			return setPositiveNumber( elem, value, subtract );
		}
	};
} );

jQuery.cssHooks.marginLeft = addGetHookIf( support.reliableMarginLeft,
	function( elem, computed ) {
		if ( computed ) {
			return ( parseFloat( curCSS( elem, "marginLeft" ) ) ||
				elem.getBoundingClientRect().left -
					swap( elem, { marginLeft: 0 }, function() {
						return elem.getBoundingClientRect().left;
					} )
			) + "px";
		}
	}
);

// These hooks are used by animate to expand properties
jQuery.each( {
	margin: "",
	padding: "",
	border: "Width"
}, function( prefix, suffix ) {
	jQuery.cssHooks[ prefix + suffix ] = {
		expand: function( value ) {
			var i = 0,
				expanded = {},

				// Assumes a single number if not a string
				parts = typeof value === "string" ? value.split( " " ) : [ value ];

			for ( ; i < 4; i++ ) {
				expanded[ prefix + cssExpand[ i ] + suffix ] =
					parts[ i ] || parts[ i - 2 ] || parts[ 0 ];
			}

			return expanded;
		}
	};

	if ( prefix !== "margin" ) {
		jQuery.cssHooks[ prefix + suffix ].set = setPositiveNumber;
	}
} );

jQuery.fn.extend( {
	css: function( name, value ) {
		return access( this, function( elem, name, value ) {
			var styles, len,
				map = {},
				i = 0;

			if ( Array.isArray( name ) ) {
				styles = getStyles( elem );
				len = name.length;

				for ( ; i < len; i++ ) {
					map[ name[ i ] ] = jQuery.css( elem, name[ i ], false, styles );
				}

				return map;
			}

			return value !== undefined ?
				jQuery.style( elem, name, value ) :
				jQuery.css( elem, name );
		}, name, value, arguments.length > 1 );
	}
} );


function Tween( elem, options, prop, end, easing ) {
	return new Tween.prototype.init( elem, options, prop, end, easing );
}
jQuery.Tween = Tween;

Tween.prototype = {
	constructor: Tween,
	init: function( elem, options, prop, end, easing, unit ) {
		this.elem = elem;
		this.prop = prop;
		this.easing = easing || jQuery.easing._default;
		this.options = options;
		this.start = this.now = this.cur();
		this.end = end;
		this.unit = unit || ( jQuery.cssNumber[ prop ] ? "" : "px" );
	},
	cur: function() {
		var hooks = Tween.propHooks[ this.prop ];

		return hooks && hooks.get ?
			hooks.get( this ) :
			Tween.propHooks._default.get( this );
	},
	run: function( percent ) {
		var eased,
			hooks = Tween.propHooks[ this.prop ];

		if ( this.options.duration ) {
			this.pos = eased = jQuery.easing[ this.easing ](
				percent, this.options.duration * percent, 0, 1, this.options.duration
			);
		} else {
			this.pos = eased = percent;
		}
		this.now = ( this.end - this.start ) * eased + this.start;

		if ( this.options.step ) {
			this.options.step.call( this.elem, this.now, this );
		}

		if ( hooks && hooks.set ) {
			hooks.set( this );
		} else {
			Tween.propHooks._default.set( this );
		}
		return this;
	}
};

Tween.prototype.init.prototype = Tween.prototype;

Tween.propHooks = {
	_default: {
		get: function( tween ) {
			var result;

			// Use a property on the element directly when it is not a DOM element,
			// or when there is no matching style property that exists.
			if ( tween.elem.nodeType !== 1 ||
				tween.elem[ tween.prop ] != null && tween.elem.style[ tween.prop ] == null ) {
				return tween.elem[ tween.prop ];
			}

			// Passing an empty string as a 3rd parameter to .css will automatically
			// attempt a parseFloat and fallback to a string if the parse fails.
			// Simple values such as "10px" are parsed to Float;
			// complex values such as "rotate(1rad)" are returned as-is.
			result = jQuery.css( tween.elem, tween.prop, "" );

			// Empty strings, null, undefined and "auto" are converted to 0.
			return !result || result === "auto" ? 0 : result;
		},
		set: function( tween ) {

			// Use step hook for back compat.
			// Use cssHook if its there.
			// Use .style if available and use plain properties where available.
			if ( jQuery.fx.step[ tween.prop ] ) {
				jQuery.fx.step[ tween.prop ]( tween );
			} else if ( tween.elem.nodeType === 1 && (
				jQuery.cssHooks[ tween.prop ] ||
					tween.elem.style[ finalPropName( tween.prop ) ] != null ) ) {
				jQuery.style( tween.elem, tween.prop, tween.now + tween.unit );
			} else {
				tween.elem[ tween.prop ] = tween.now;
			}
		}
	}
};

// Support: IE <=9 only
// Panic based approach to setting things on disconnected nodes
Tween.propHooks.scrollTop = Tween.propHooks.scrollLeft = {
	set: function( tween ) {
		if ( tween.elem.nodeType && tween.elem.parentNode ) {
			tween.elem[ tween.prop ] = tween.now;
		}
	}
};

jQuery.easing = {
	linear: function( p ) {
		return p;
	},
	swing: function( p ) {
		return 0.5 - Math.cos( p * Math.PI ) / 2;
	},
	_default: "swing"
};

jQuery.fx = Tween.prototype.init;

// Back compat <1.8 extension point
jQuery.fx.step = {};




var
	fxNow, inProgress,
	rfxtypes = /^(?:toggle|show|hide)$/,
	rrun = /queueHooks$/;

function schedule() {
	if ( inProgress ) {
		if ( document.hidden === false && window.requestAnimationFrame ) {
			window.requestAnimationFrame( schedule );
		} else {
			window.setTimeout( schedule, jQuery.fx.interval );
		}

		jQuery.fx.tick();
	}
}

// Animations created synchronously will run synchronously
function createFxNow() {
	window.setTimeout( function() {
		fxNow = undefined;
	} );
	return ( fxNow = Date.now() );
}

// Generate parameters to create a standard animation
function genFx( type, includeWidth ) {
	var which,
		i = 0,
		attrs = { height: type };

	// If we include width, step value is 1 to do all cssExpand values,
	// otherwise step value is 2 to skip over Left and Right
	includeWidth = includeWidth ? 1 : 0;
	for ( ; i < 4; i += 2 - includeWidth ) {
		which = cssExpand[ i ];
		attrs[ "margin" + which ] = attrs[ "padding" + which ] = type;
	}

	if ( includeWidth ) {
		attrs.opacity = attrs.width = type;
	}

	return attrs;
}

function createTween( value, prop, animation ) {
	var tween,
		collection = ( Animation.tweeners[ prop ] || [] ).concat( Animation.tweeners[ "*" ] ),
		index = 0,
		length = collection.length;
	for ( ; index < length; index++ ) {
		if ( ( tween = collection[ index ].call( animation, prop, value ) ) ) {

			// We're done with this property
			return tween;
		}
	}
}

function defaultPrefilter( elem, props, opts ) {
	var prop, value, toggle, hooks, oldfire, propTween, restoreDisplay, display,
		isBox = "width" in props || "height" in props,
		anim = this,
		orig = {},
		style = elem.style,
		hidden = elem.nodeType && isHiddenWithinTree( elem ),
		dataShow = dataPriv.get( elem, "fxshow" );

	// Queue-skipping animations hijack the fx hooks
	if ( !opts.queue ) {
		hooks = jQuery._queueHooks( elem, "fx" );
		if ( hooks.unqueued == null ) {
			hooks.unqueued = 0;
			oldfire = hooks.empty.fire;
			hooks.empty.fire = function() {
				if ( !hooks.unqueued ) {
					oldfire();
				}
			};
		}
		hooks.unqueued++;

		anim.always( function() {

			// Ensure the complete handler is called before this completes
			anim.always( function() {
				hooks.unqueued--;
				if ( !jQuery.queue( elem, "fx" ).length ) {
					hooks.empty.fire();
				}
			} );
		} );
	}

	// Detect show/hide animations
	for ( prop in props ) {
		value = props[ prop ];
		if ( rfxtypes.test( value ) ) {
			delete props[ prop ];
			toggle = toggle || value === "toggle";
			if ( value === ( hidden ? "hide" : "show" ) ) {

				// Pretend to be hidden if this is a "show" and
				// there is still data from a stopped show/hide
				if ( value === "show" && dataShow && dataShow[ prop ] !== undefined ) {
					hidden = true;

				// Ignore all other no-op show/hide data
				} else {
					continue;
				}
			}
			orig[ prop ] = dataShow && dataShow[ prop ] || jQuery.style( elem, prop );
		}
	}

	// Bail out if this is a no-op like .hide().hide()
	propTween = !jQuery.isEmptyObject( props );
	if ( !propTween && jQuery.isEmptyObject( orig ) ) {
		return;
	}

	// Restrict "overflow" and "display" styles during box animations
	if ( isBox && elem.nodeType === 1 ) {

		// Support: IE <=9 - 11, Edge 12 - 15
		// Record all 3 overflow attributes because IE does not infer the shorthand
		// from identically-valued overflowX and overflowY and Edge just mirrors
		// the overflowX value there.
		opts.overflow = [ style.overflow, style.overflowX, style.overflowY ];

		// Identify a display type, preferring old show/hide data over the CSS cascade
		restoreDisplay = dataShow && dataShow.display;
		if ( restoreDisplay == null ) {
			restoreDisplay = dataPriv.get( elem, "display" );
		}
		display = jQuery.css( elem, "display" );
		if ( display === "none" ) {
			if ( restoreDisplay ) {
				display = restoreDisplay;
			} else {

				// Get nonempty value(s) by temporarily forcing visibility
				showHide( [ elem ], true );
				restoreDisplay = elem.style.display || restoreDisplay;
				display = jQuery.css( elem, "display" );
				showHide( [ elem ] );
			}
		}

		// Animate inline elements as inline-block
		if ( display === "inline" || display === "inline-block" && restoreDisplay != null ) {
			if ( jQuery.css( elem, "float" ) === "none" ) {

				// Restore the original display value at the end of pure show/hide animations
				if ( !propTween ) {
					anim.done( function() {
						style.display = restoreDisplay;
					} );
					if ( restoreDisplay == null ) {
						display = style.display;
						restoreDisplay = display === "none" ? "" : display;
					}
				}
				style.display = "inline-block";
			}
		}
	}

	if ( opts.overflow ) {
		style.overflow = "hidden";
		anim.always( function() {
			style.overflow = opts.overflow[ 0 ];
			style.overflowX = opts.overflow[ 1 ];
			style.overflowY = opts.overflow[ 2 ];
		} );
	}

	// Implement show/hide animations
	propTween = false;
	for ( prop in orig ) {

		// General show/hide setup for this element animation
		if ( !propTween ) {
			if ( dataShow ) {
				if ( "hidden" in dataShow ) {
					hidden = dataShow.hidden;
				}
			} else {
				dataShow = dataPriv.access( elem, "fxshow", { display: restoreDisplay } );
			}

			// Store hidden/visible for toggle so `.stop().toggle()` "reverses"
			if ( toggle ) {
				dataShow.hidden = !hidden;
			}

			// Show elements before animating them
			if ( hidden ) {
				showHide( [ elem ], true );
			}

			/* eslint-disable no-loop-func */

			anim.done( function() {

				/* eslint-enable no-loop-func */

				// The final step of a "hide" animation is actually hiding the element
				if ( !hidden ) {
					showHide( [ elem ] );
				}
				dataPriv.remove( elem, "fxshow" );
				for ( prop in orig ) {
					jQuery.style( elem, prop, orig[ prop ] );
				}
			} );
		}

		// Per-property setup
		propTween = createTween( hidden ? dataShow[ prop ] : 0, prop, anim );
		if ( !( prop in dataShow ) ) {
			dataShow[ prop ] = propTween.start;
			if ( hidden ) {
				propTween.end = propTween.start;
				propTween.start = 0;
			}
		}
	}
}

function propFilter( props, specialEasing ) {
	var index, name, easing, value, hooks;

	// camelCase, specialEasing and expand cssHook pass
	for ( index in props ) {
		name = camelCase( index );
		easing = specialEasing[ name ];
		value = props[ index ];
		if ( Array.isArray( value ) ) {
			easing = value[ 1 ];
			value = props[ index ] = value[ 0 ];
		}

		if ( index !== name ) {
			props[ name ] = value;
			delete props[ index ];
		}

		hooks = jQuery.cssHooks[ name ];
		if ( hooks && "expand" in hooks ) {
			value = hooks.expand( value );
			delete props[ name ];

			// Not quite $.extend, this won't overwrite existing keys.
			// Reusing 'index' because we have the correct "name"
			for ( index in value ) {
				if ( !( index in props ) ) {
					props[ index ] = value[ index ];
					specialEasing[ index ] = easing;
				}
			}
		} else {
			specialEasing[ name ] = easing;
		}
	}
}

function Animation( elem, properties, options ) {
	var result,
		stopped,
		index = 0,
		length = Animation.prefilters.length,
		deferred = jQuery.Deferred().always( function() {

			// Don't match elem in the :animated selector
			delete tick.elem;
		} ),
		tick = function() {
			if ( stopped ) {
				return false;
			}
			var currentTime = fxNow || createFxNow(),
				remaining = Math.max( 0, animation.startTime + animation.duration - currentTime ),

				// Support: Android 2.3 only
				// Archaic crash bug won't allow us to use `1 - ( 0.5 || 0 )` (#12497)
				temp = remaining / animation.duration || 0,
				percent = 1 - temp,
				index = 0,
				length = animation.tweens.length;

			for ( ; index < length; index++ ) {
				animation.tweens[ index ].run( percent );
			}

			deferred.notifyWith( elem, [ animation, percent, remaining ] );

			// If there's more to do, yield
			if ( percent < 1 && length ) {
				return remaining;
			}

			// If this was an empty animation, synthesize a final progress notification
			if ( !length ) {
				deferred.notifyWith( elem, [ animation, 1, 0 ] );
			}

			// Resolve the animation and report its conclusion
			deferred.resolveWith( elem, [ animation ] );
			return false;
		},
		animation = deferred.promise( {
			elem: elem,
			props: jQuery.extend( {}, properties ),
			opts: jQuery.extend( true, {
				specialEasing: {},
				easing: jQuery.easing._default
			}, options ),
			originalProperties: properties,
			originalOptions: options,
			startTime: fxNow || createFxNow(),
			duration: options.duration,
			tweens: [],
			createTween: function( prop, end ) {
				var tween = jQuery.Tween( elem, animation.opts, prop, end,
					animation.opts.specialEasing[ prop ] || animation.opts.easing );
				animation.tweens.push( tween );
				return tween;
			},
			stop: function( gotoEnd ) {
				var index = 0,

					// If we are going to the end, we want to run all the tweens
					// otherwise we skip this part
					length = gotoEnd ? animation.tweens.length : 0;
				if ( stopped ) {
					return this;
				}
				stopped = true;
				for ( ; index < length; index++ ) {
					animation.tweens[ index ].run( 1 );
				}

				// Resolve when we played the last frame; otherwise, reject
				if ( gotoEnd ) {
					deferred.notifyWith( elem, [ animation, 1, 0 ] );
					deferred.resolveWith( elem, [ animation, gotoEnd ] );
				} else {
					deferred.rejectWith( elem, [ animation, gotoEnd ] );
				}
				return this;
			}
		} ),
		props = animation.props;

	propFilter( props, animation.opts.specialEasing );

	for ( ; index < length; index++ ) {
		result = Animation.prefilters[ index ].call( animation, elem, props, animation.opts );
		if ( result ) {
			if ( isFunction( result.stop ) ) {
				jQuery._queueHooks( animation.elem, animation.opts.queue ).stop =
					result.stop.bind( result );
			}
			return result;
		}
	}

	jQuery.map( props, createTween, animation );

	if ( isFunction( animation.opts.start ) ) {
		animation.opts.start.call( elem, animation );
	}

	// Attach callbacks from options
	animation
		.progress( animation.opts.progress )
		.done( animation.opts.done, animation.opts.complete )
		.fail( animation.opts.fail )
		.always( animation.opts.always );

	jQuery.fx.timer(
		jQuery.extend( tick, {
			elem: elem,
			anim: animation,
			queue: animation.opts.queue
		} )
	);

	return animation;
}

jQuery.Animation = jQuery.extend( Animation, {

	tweeners: {
		"*": [ function( prop, value ) {
			var tween = this.createTween( prop, value );
			adjustCSS( tween.elem, prop, rcssNum.exec( value ), tween );
			return tween;
		} ]
	},

	tweener: function( props, callback ) {
		if ( isFunction( props ) ) {
			callback = props;
			props = [ "*" ];
		} else {
			props = props.match( rnothtmlwhite );
		}

		var prop,
			index = 0,
			length = props.length;

		for ( ; index < length; index++ ) {
			prop = props[ index ];
			Animation.tweeners[ prop ] = Animation.tweeners[ prop ] || [];
			Animation.tweeners[ prop ].unshift( callback );
		}
	},

	prefilters: [ defaultPrefilter ],

	prefilter: function( callback, prepend ) {
		if ( prepend ) {
			Animation.prefilters.unshift( callback );
		} else {
			Animation.prefilters.push( callback );
		}
	}
} );

jQuery.speed = function( speed, easing, fn ) {
	var opt = speed && typeof speed === "object" ? jQuery.extend( {}, speed ) : {
		complete: fn || !fn && easing ||
			isFunction( speed ) && speed,
		duration: speed,
		easing: fn && easing || easing && !isFunction( easing ) && easing
	};

	// Go to the end state if fx are off
	if ( jQuery.fx.off ) {
		opt.duration = 0;

	} else {
		if ( typeof opt.duration !== "number" ) {
			if ( opt.duration in jQuery.fx.speeds ) {
				opt.duration = jQuery.fx.speeds[ opt.duration ];

			} else {
				opt.duration = jQuery.fx.speeds._default;
			}
		}
	}

	// Normalize opt.queue - true/undefined/null -> "fx"
	if ( opt.queue == null || opt.queue === true ) {
		opt.queue = "fx";
	}

	// Queueing
	opt.old = opt.complete;

	opt.complete = function() {
		if ( isFunction( opt.old ) ) {
			opt.old.call( this );
		}

		if ( opt.queue ) {
			jQuery.dequeue( this, opt.queue );
		}
	};

	return opt;
};

jQuery.fn.extend( {
	fadeTo: function( speed, to, easing, callback ) {

		// Show any hidden elements after setting opacity to 0
		return this.filter( isHiddenWithinTree ).css( "opacity", 0 ).show()

			// Animate to the value specified
			.end().animate( { opacity: to }, speed, easing, callback );
	},
	animate: function( prop, speed, easing, callback ) {
		var empty = jQuery.isEmptyObject( prop ),
			optall = jQuery.speed( speed, easing, callback ),
			doAnimation = function() {

				// Operate on a copy of prop so per-property easing won't be lost
				var anim = Animation( this, jQuery.extend( {}, prop ), optall );

				// Empty animations, or finishing resolves immediately
				if ( empty || dataPriv.get( this, "finish" ) ) {
					anim.stop( true );
				}
			};

		doAnimation.finish = doAnimation;

		return empty || optall.queue === false ?
			this.each( doAnimation ) :
			this.queue( optall.queue, doAnimation );
	},
	stop: function( type, clearQueue, gotoEnd ) {
		var stopQueue = function( hooks ) {
			var stop = hooks.stop;
			delete hooks.stop;
			stop( gotoEnd );
		};

		if ( typeof type !== "string" ) {
			gotoEnd = clearQueue;
			clearQueue = type;
			type = undefined;
		}
		if ( clearQueue ) {
			this.queue( type || "fx", [] );
		}

		return this.each( function() {
			var dequeue = true,
				index = type != null && type + "queueHooks",
				timers = jQuery.timers,
				data = dataPriv.get( this );

			if ( index ) {
				if ( data[ index ] && data[ index ].stop ) {
					stopQueue( data[ index ] );
				}
			} else {
				for ( index in data ) {
					if ( data[ index ] && data[ index ].stop && rrun.test( index ) ) {
						stopQueue( data[ index ] );
					}
				}
			}

			for ( index = timers.length; index--; ) {
				if ( timers[ index ].elem === this &&
					( type == null || timers[ index ].queue === type ) ) {

					timers[ index ].anim.stop( gotoEnd );
					dequeue = false;
					timers.splice( index, 1 );
				}
			}

			// Start the next in the queue if the last step wasn't forced.
			// Timers currently will call their complete callbacks, which
			// will dequeue but only if they were gotoEnd.
			if ( dequeue || !gotoEnd ) {
				jQuery.dequeue( this, type );
			}
		} );
	},
	finish: function( type ) {
		if ( type !== false ) {
			type = type || "fx";
		}
		return this.each( function() {
			var index,
				data = dataPriv.get( this ),
				queue = data[ type + "queue" ],
				hooks = data[ type + "queueHooks" ],
				timers = jQuery.timers,
				length = queue ? queue.length : 0;

			// Enable finishing flag on private data
			data.finish = true;

			// Empty the queue first
			jQuery.queue( this, type, [] );

			if ( hooks && hooks.stop ) {
				hooks.stop.call( this, true );
			}

			// Look for any active animations, and finish them
			for ( index = timers.length; index--; ) {
				if ( timers[ index ].elem === this && timers[ index ].queue === type ) {
					timers[ index ].anim.stop( true );
					timers.splice( index, 1 );
				}
			}

			// Look for any animations in the old queue and finish them
			for ( index = 0; index < length; index++ ) {
				if ( queue[ index ] && queue[ index ].finish ) {
					queue[ index ].finish.call( this );
				}
			}

			// Turn off finishing flag
			delete data.finish;
		} );
	}
} );

jQuery.each( [ "toggle", "show", "hide" ], function( _i, name ) {
	var cssFn = jQuery.fn[ name ];
	jQuery.fn[ name ] = function( speed, easing, callback ) {
		return speed == null || typeof speed === "boolean" ?
			cssFn.apply( this, arguments ) :
			this.animate( genFx( name, true ), speed, easing, callback );
	};
} );

// Generate shortcuts for custom animations
jQuery.each( {
	slideDown: genFx( "show" ),
	slideUp: genFx( "hide" ),
	slideToggle: genFx( "toggle" ),
	fadeIn: { opacity: "show" },
	fadeOut: { opacity: "hide" },
	fadeToggle: { opacity: "toggle" }
}, function( name, props ) {
	jQuery.fn[ name ] = function( speed, easing, callback ) {
		return this.animate( props, speed, easing, callback );
	};
} );

jQuery.timers = [];
jQuery.fx.tick = function() {
	var timer,
		i = 0,
		timers = jQuery.timers;

	fxNow = Date.now();

	for ( ; i < timers.length; i++ ) {
		timer = timers[ i ];

		// Run the timer and safely remove it when done (allowing for external removal)
		if ( !timer() && timers[ i ] === timer ) {
			timers.splice( i--, 1 );
		}
	}

	if ( !timers.length ) {
		jQuery.fx.stop();
	}
	fxNow = undefined;
};

jQuery.fx.timer = function( timer ) {
	jQuery.timers.push( timer );
	jQuery.fx.start();
};

jQuery.fx.interval = 13;
jQuery.fx.start = function() {
	if ( inProgress ) {
		return;
	}

	inProgress = true;
	schedule();
};

jQuery.fx.stop = function() {
	inProgress = null;
};

jQuery.fx.speeds = {
	slow: 600,
	fast: 200,

	// Default speed
	_default: 400
};


// Based off of the plugin by Clint Helfers, with permission.
// https://web.archive.org/web/20100324014747/http://blindsignals.com/index.php/2009/07/jquery-delay/
jQuery.fn.delay = function( time, type ) {
	time = jQuery.fx ? jQuery.fx.speeds[ time ] || time : time;
	type = type || "fx";

	return this.queue( type, function( next, hooks ) {
		var timeout = window.setTimeout( next, time );
		hooks.stop = function() {
			window.clearTimeout( timeout );
		};
	} );
};


( function() {
	var input = document.createElement( "input" ),
		select = document.createElement( "select" ),
		opt = select.appendChild( document.createElement( "option" ) );

	input.type = "checkbox";

	// Support: Android <=4.3 only
	// Default value for a checkbox should be "on"
	support.checkOn = input.value !== "";

	// Support: IE <=11 only
	// Must access selectedIndex to make default options select
	support.optSelected = opt.selected;

	// Support: IE <=11 only
	// An input loses its value after becoming a radio
	input = document.createElement( "input" );
	input.value = "t";
	input.type = "radio";
	support.radioValue = input.value === "t";
} )();


var boolHook,
	attrHandle = jQuery.expr.attrHandle;

jQuery.fn.extend( {
	attr: function( name, value ) {
		return access( this, jQuery.attr, name, value, arguments.length > 1 );
	},

	removeAttr: function( name ) {
		return this.each( function() {
			jQuery.removeAttr( this, name );
		} );
	}
} );

jQuery.extend( {
	attr: function( elem, name, value ) {
		var ret, hooks,
			nType = elem.nodeType;

		// Don't get/set attributes on text, comment and attribute nodes
		if ( nType === 3 || nType === 8 || nType === 2 ) {
			return;
		}

		// Fallback to prop when attributes are not supported
		if ( typeof elem.getAttribute === "undefined" ) {
			return jQuery.prop( elem, name, value );
		}

		// Attribute hooks are determined by the lowercase version
		// Grab necessary hook if one is defined
		if ( nType !== 1 || !jQuery.isXMLDoc( elem ) ) {
			hooks = jQuery.attrHooks[ name.toLowerCase() ] ||
				( jQuery.expr.match.bool.test( name ) ? boolHook : undefined );
		}

		if ( value !== undefined ) {
			if ( value === null ) {
				jQuery.removeAttr( elem, name );
				return;
			}

			if ( hooks && "set" in hooks &&
				( ret = hooks.set( elem, value, name ) ) !== undefined ) {
				return ret;
			}

			elem.setAttribute( name, value + "" );
			return value;
		}

		if ( hooks && "get" in hooks && ( ret = hooks.get( elem, name ) ) !== null ) {
			return ret;
		}

		ret = jQuery.find.attr( elem, name );

		// Non-existent attributes return null, we normalize to undefined
		return ret == null ? undefined : ret;
	},

	attrHooks: {
		type: {
			set: function( elem, value ) {
				if ( !support.radioValue && value === "radio" &&
					nodeName( elem, "input" ) ) {
					var val = elem.value;
					elem.setAttribute( "type", value );
					if ( val ) {
						elem.value = val;
					}
					return value;
				}
			}
		}
	},

	removeAttr: function( elem, value ) {
		var name,
			i = 0,

			// Attribute names can contain non-HTML whitespace characters
			// https://html.spec.whatwg.org/multipage/syntax.html#attributes-2
			attrNames = value && value.match( rnothtmlwhite );

		if ( attrNames && elem.nodeType === 1 ) {
			while ( ( name = attrNames[ i++ ] ) ) {
				elem.removeAttribute( name );
			}
		}
	}
} );

// Hooks for boolean attributes
boolHook = {
	set: function( elem, value, name ) {
		if ( value === false ) {

			// Remove boolean attributes when set to false
			jQuery.removeAttr( elem, name );
		} else {
			elem.setAttribute( name, name );
		}
		return name;
	}
};

jQuery.each( jQuery.expr.match.bool.source.match( /\w+/g ), function( _i, name ) {
	var getter = attrHandle[ name ] || jQuery.find.attr;

	attrHandle[ name ] = function( elem, name, isXML ) {
		var ret, handle,
			lowercaseName = name.toLowerCase();

		if ( !isXML ) {

			// Avoid an infinite loop by temporarily removing this function from the getter
			handle = attrHandle[ lowercaseName ];
			attrHandle[ lowercaseName ] = ret;
			ret = getter( elem, name, isXML ) != null ?
				lowercaseName :
				null;
			attrHandle[ lowercaseName ] = handle;
		}
		return ret;
	};
} );




var rfocusable = /^(?:input|select|textarea|button)$/i,
	rclickable = /^(?:a|area)$/i;

jQuery.fn.extend( {
	prop: function( name, value ) {
		return access( this, jQuery.prop, name, value, arguments.length > 1 );
	},

	removeProp: function( name ) {
		return this.each( function() {
			delete this[ jQuery.propFix[ name ] || name ];
		} );
	}
} );

jQuery.extend( {
	prop: function( elem, name, value ) {
		var ret, hooks,
			nType = elem.nodeType;

		// Don't get/set properties on text, comment and attribute nodes
		if ( nType === 3 || nType === 8 || nType === 2 ) {
			return;
		}

		if ( nType !== 1 || !jQuery.isXMLDoc( elem ) ) {

			// Fix name and attach hooks
			name = jQuery.propFix[ name ] || name;
			hooks = jQuery.propHooks[ name ];
		}

		if ( value !== undefined ) {
			if ( hooks && "set" in hooks &&
				( ret = hooks.set( elem, value, name ) ) !== undefined ) {
				return ret;
			}

			return ( elem[ name ] = value );
		}

		if ( hooks && "get" in hooks && ( ret = hooks.get( elem, name ) ) !== null ) {
			return ret;
		}

		return elem[ name ];
	},

	propHooks: {
		tabIndex: {
			get: function( elem ) {

				// Support: IE <=9 - 11 only
				// elem.tabIndex doesn't always return the
				// correct value when it hasn't been explicitly set
				// https://web.archive.org/web/20141116233347/http://fluidproject.org/blog/2008/01/09/getting-setting-and-removing-tabindex-values-with-javascript/
				// Use proper attribute retrieval(#12072)
				var tabindex = jQuery.find.attr( elem, "tabindex" );

				if ( tabindex ) {
					return parseInt( tabindex, 10 );
				}

				if (
					rfocusable.test( elem.nodeName ) ||
					rclickable.test( elem.nodeName ) &&
					elem.href
				) {
					return 0;
				}

				return -1;
			}
		}
	},

	propFix: {
		"for": "htmlFor",
		"class": "className"
	}
} );

// Support: IE <=11 only
// Accessing the selectedIndex property
// forces the browser to respect setting selected
// on the option
// The getter ensures a default option is selected
// when in an optgroup
// eslint rule "no-unused-expressions" is disabled for this code
// since it considers such accessions noop
if ( !support.optSelected ) {
	jQuery.propHooks.selected = {
		get: function( elem ) {

			/* eslint no-unused-expressions: "off" */

			var parent = elem.parentNode;
			if ( parent && parent.parentNode ) {
				parent.parentNode.selectedIndex;
			}
			return null;
		},
		set: function( elem ) {

			/* eslint no-unused-expressions: "off" */

			var parent = elem.parentNode;
			if ( parent ) {
				parent.selectedIndex;

				if ( parent.parentNode ) {
					parent.parentNode.selectedIndex;
				}
			}
		}
	};
}

jQuery.each( [
	"tabIndex",
	"readOnly",
	"maxLength",
	"cellSpacing",
	"cellPadding",
	"rowSpan",
	"colSpan",
	"useMap",
	"frameBorder",
	"contentEditable"
], function() {
	jQuery.propFix[ this.toLowerCase() ] = this;
} );




	// Strip and collapse whitespace according to HTML spec
	// https://infra.spec.whatwg.org/#strip-and-collapse-ascii-whitespace
	function stripAndCollapse( value ) {
		var tokens = value.match( rnothtmlwhite ) || [];
		return tokens.join( " " );
	}


function getClass( elem ) {
	return elem.getAttribute && elem.getAttribute( "class" ) || "";
}

function classesToArray( value ) {
	if ( Array.isArray( value ) ) {
		return value;
	}
	if ( typeof value === "string" ) {
		return value.match( rnothtmlwhite ) || [];
	}
	return [];
}

jQuery.fn.extend( {
	addClass: function( value ) {
		var classes, elem, cur, curValue, clazz, j, finalValue,
			i = 0;

		if ( isFunction( value ) ) {
			return this.each( function( j ) {
				jQuery( this ).addClass( value.call( this, j, getClass( this ) ) );
			} );
		}

		classes = classesToArray( value );

		if ( classes.length ) {
			while ( ( elem = this[ i++ ] ) ) {
				curValue = getClass( elem );
				cur = elem.nodeType === 1 && ( " " + stripAndCollapse( curValue ) + " " );

				if ( cur ) {
					j = 0;
					while ( ( clazz = classes[ j++ ] ) ) {
						if ( cur.indexOf( " " + clazz + " " ) < 0 ) {
							cur += clazz + " ";
						}
					}

					// Only assign if different to avoid unneeded rendering.
					finalValue = stripAndCollapse( cur );
					if ( curValue !== finalValue ) {
						elem.setAttribute( "class", finalValue );
					}
				}
			}
		}

		return this;
	},

	removeClass: function( value ) {
		var classes, elem, cur, curValue, clazz, j, finalValue,
			i = 0;

		if ( isFunction( value ) ) {
			return this.each( function( j ) {
				jQuery( this ).removeClass( value.call( this, j, getClass( this ) ) );
			} );
		}

		if ( !arguments.length ) {
			return this.attr( "class", "" );
		}

		classes = classesToArray( value );

		if ( classes.length ) {
			while ( ( elem = this[ i++ ] ) ) {
				curValue = getClass( elem );

				// This expression is here for better compressibility (see addClass)
				cur = elem.nodeType === 1 && ( " " + stripAndCollapse( curValue ) + " " );

				if ( cur ) {
					j = 0;
					while ( ( clazz = classes[ j++ ] ) ) {

						// Remove *all* instances
						while ( cur.indexOf( " " + clazz + " " ) > -1 ) {
							cur = cur.replace( " " + clazz + " ", " " );
						}
					}

					// Only assign if different to avoid unneeded rendering.
					finalValue = stripAndCollapse( cur );
					if ( curValue !== finalValue ) {
						elem.setAttribute( "class", finalValue );
					}
				}
			}
		}

		return this;
	},

	toggleClass: function( value, stateVal ) {
		var type = typeof value,
			isValidValue = type === "string" || Array.isArray( value );

		if ( typeof stateVal === "boolean" && isValidValue ) {
			return stateVal ? this.addClass( value ) : this.removeClass( value );
		}

		if ( isFunction( value ) ) {
			return this.each( function( i ) {
				jQuery( this ).toggleClass(
					value.call( this, i, getClass( this ), stateVal ),
					stateVal
				);
			} );
		}

		return this.each( function() {
			var className, i, self, classNames;

			if ( isValidValue ) {

				// Toggle individual class names
				i = 0;
				self = jQuery( this );
				classNames = classesToArray( value );

				while ( ( className = classNames[ i++ ] ) ) {

					// Check each className given, space separated list
					if ( self.hasClass( className ) ) {
						self.removeClass( className );
					} else {
						self.addClass( className );
					}
				}

			// Toggle whole class name
			} else if ( value === undefined || type === "boolean" ) {
				className = getClass( this );
				if ( className ) {

					// Store className if set
					dataPriv.set( this, "__className__", className );
				}

				// If the element has a class name or if we're passed `false`,
				// then remove the whole classname (if there was one, the above saved it).
				// Otherwise bring back whatever was previously saved (if anything),
				// falling back to the empty string if nothing was stored.
				if ( this.setAttribute ) {
					this.setAttribute( "class",
						className || value === false ?
							"" :
							dataPriv.get( this, "__className__" ) || ""
					);
				}
			}
		} );
	},

	hasClass: function( selector ) {
		var className, elem,
			i = 0;

		className = " " + selector + " ";
		while ( ( elem = this[ i++ ] ) ) {
			if ( elem.nodeType === 1 &&
				( " " + stripAndCollapse( getClass( elem ) ) + " " ).indexOf( className ) > -1 ) {
				return true;
			}
		}

		return false;
	}
} );




var rreturn = /\r/g;

jQuery.fn.extend( {
	val: function( value ) {
		var hooks, ret, valueIsFunction,
			elem = this[ 0 ];

		if ( !arguments.length ) {
			if ( elem ) {
				hooks = jQuery.valHooks[ elem.type ] ||
					jQuery.valHooks[ elem.nodeName.toLowerCase() ];

				if ( hooks &&
					"get" in hooks &&
					( ret = hooks.get( elem, "value" ) ) !== undefined
				) {
					return ret;
				}

				ret = elem.value;

				// Handle most common string cases
				if ( typeof ret === "string" ) {
					return ret.replace( rreturn, "" );
				}

				// Handle cases where value is null/undef or number
				return ret == null ? "" : ret;
			}

			return;
		}

		valueIsFunction = isFunction( value );

		return this.each( function( i ) {
			var val;

			if ( this.nodeType !== 1 ) {
				return;
			}

			if ( valueIsFunction ) {
				val = value.call( this, i, jQuery( this ).val() );
			} else {
				val = value;
			}

			// Treat null/undefined as ""; convert numbers to string
			if ( val == null ) {
				val = "";

			} else if ( typeof val === "number" ) {
				val += "";

			} else if ( Array.isArray( val ) ) {
				val = jQuery.map( val, function( value ) {
					return value == null ? "" : value + "";
				} );
			}

			hooks = jQuery.valHooks[ this.type ] || jQuery.valHooks[ this.nodeName.toLowerCase() ];

			// If set returns undefined, fall back to normal setting
			if ( !hooks || !( "set" in hooks ) || hooks.set( this, val, "value" ) === undefined ) {
				this.value = val;
			}
		} );
	}
} );

jQuery.extend( {
	valHooks: {
		option: {
			get: function( elem ) {

				var val = jQuery.find.attr( elem, "value" );
				return val != null ?
					val :

					// Support: IE <=10 - 11 only
					// option.text throws exceptions (#14686, #14858)
					// Strip and collapse whitespace
					// https://html.spec.whatwg.org/#strip-and-collapse-whitespace
					stripAndCollapse( jQuery.text( elem ) );
			}
		},
		select: {
			get: function( elem ) {
				var value, option, i,
					options = elem.options,
					index = elem.selectedIndex,
					one = elem.type === "select-one",
					values = one ? null : [],
					max = one ? index + 1 : options.length;

				if ( index < 0 ) {
					i = max;

				} else {
					i = one ? index : 0;
				}

				// Loop through all the selected options
				for ( ; i < max; i++ ) {
					option = options[ i ];

					// Support: IE <=9 only
					// IE8-9 doesn't update selected after form reset (#2551)
					if ( ( option.selected || i === index ) &&

							// Don't return options that are disabled or in a disabled optgroup
							!option.disabled &&
							( !option.parentNode.disabled ||
								!nodeName( option.parentNode, "optgroup" ) ) ) {

						// Get the specific value for the option
						value = jQuery( option ).val();

						// We don't need an array for one selects
						if ( one ) {
							return value;
						}

						// Multi-Selects return an array
						values.push( value );
					}
				}

				return values;
			},

			set: function( elem, value ) {
				var optionSet, option,
					options = elem.options,
					values = jQuery.makeArray( value ),
					i = options.length;

				while ( i-- ) {
					option = options[ i ];

					/* eslint-disable no-cond-assign */

					if ( option.selected =
						jQuery.inArray( jQuery.valHooks.option.get( option ), values ) > -1
					) {
						optionSet = true;
					}

					/* eslint-enable no-cond-assign */
				}

				// Force browsers to behave consistently when non-matching value is set
				if ( !optionSet ) {
					elem.selectedIndex = -1;
				}
				return values;
			}
		}
	}
} );

// Radios and checkboxes getter/setter
jQuery.each( [ "radio", "checkbox" ], function() {
	jQuery.valHooks[ this ] = {
		set: function( elem, value ) {
			if ( Array.isArray( value ) ) {
				return ( elem.checked = jQuery.inArray( jQuery( elem ).val(), value ) > -1 );
			}
		}
	};
	if ( !support.checkOn ) {
		jQuery.valHooks[ this ].get = function( elem ) {
			return elem.getAttribute( "value" ) === null ? "on" : elem.value;
		};
	}
} );




// Return jQuery for attributes-only inclusion


support.focusin = "onfocusin" in window;


var rfocusMorph = /^(?:focusinfocus|focusoutblur)$/,
	stopPropagationCallback = function( e ) {
		e.stopPropagation();
	};

jQuery.extend( jQuery.event, {

	trigger: function( event, data, elem, onlyHandlers ) {

		var i, cur, tmp, bubbleType, ontype, handle, special, lastElement,
			eventPath = [ elem || document ],
			type = hasOwn.call( event, "type" ) ? event.type : event,
			namespaces = hasOwn.call( event, "namespace" ) ? event.namespace.split( "." ) : [];

		cur = lastElement = tmp = elem = elem || document;

		// Don't do events on text and comment nodes
		if ( elem.nodeType === 3 || elem.nodeType === 8 ) {
			return;
		}

		// focus/blur morphs to focusin/out; ensure we're not firing them right now
		if ( rfocusMorph.test( type + jQuery.event.triggered ) ) {
			return;
		}

		if ( type.indexOf( "." ) > -1 ) {

			// Namespaced trigger; create a regexp to match event type in handle()
			namespaces = type.split( "." );
			type = namespaces.shift();
			namespaces.sort();
		}
		ontype = type.indexOf( ":" ) < 0 && "on" + type;

		// Caller can pass in a jQuery.Event object, Object, or just an event type string
		event = event[ jQuery.expando ] ?
			event :
			new jQuery.Event( type, typeof event === "object" && event );

		// Trigger bitmask: & 1 for native handlers; & 2 for jQuery (always true)
		event.isTrigger = onlyHandlers ? 2 : 3;
		event.namespace = namespaces.join( "." );
		event.rnamespace = event.namespace ?
			new RegExp( "(^|\\.)" + namespaces.join( "\\.(?:.*\\.|)" ) + "(\\.|$)" ) :
			null;

		// Clean up the event in case it is being reused
		event.result = undefined;
		if ( !event.target ) {
			event.target = elem;
		}

		// Clone any incoming data and prepend the event, creating the handler arg list
		data = data == null ?
			[ event ] :
			jQuery.makeArray( data, [ event ] );

		// Allow special events to draw outside the lines
		special = jQuery.event.special[ type ] || {};
		if ( !onlyHandlers && special.trigger && special.trigger.apply( elem, data ) === false ) {
			return;
		}

		// Determine event propagation path in advance, per W3C events spec (#9951)
		// Bubble up to document, then to window; watch for a global ownerDocument var (#9724)
		if ( !onlyHandlers && !special.noBubble && !isWindow( elem ) ) {

			bubbleType = special.delegateType || type;
			if ( !rfocusMorph.test( bubbleType + type ) ) {
				cur = cur.parentNode;
			}
			for ( ; cur; cur = cur.parentNode ) {
				eventPath.push( cur );
				tmp = cur;
			}

			// Only add window if we got to document (e.g., not plain obj or detached DOM)
			if ( tmp === ( elem.ownerDocument || document ) ) {
				eventPath.push( tmp.defaultView || tmp.parentWindow || window );
			}
		}

		// Fire handlers on the event path
		i = 0;
		while ( ( cur = eventPath[ i++ ] ) && !event.isPropagationStopped() ) {
			lastElement = cur;
			event.type = i > 1 ?
				bubbleType :
				special.bindType || type;

			// jQuery handler
			handle = ( dataPriv.get( cur, "events" ) || Object.create( null ) )[ event.type ] &&
				dataPriv.get( cur, "handle" );
			if ( handle ) {
				handle.apply( cur, data );
			}

			// Native handler
			handle = ontype && cur[ ontype ];
			if ( handle && handle.apply && acceptData( cur ) ) {
				event.result = handle.apply( cur, data );
				if ( event.result === false ) {
					event.preventDefault();
				}
			}
		}
		event.type = type;

		// If nobody prevented the default action, do it now
		if ( !onlyHandlers && !event.isDefaultPrevented() ) {

			if ( ( !special._default ||
				special._default.apply( eventPath.pop(), data ) === false ) &&
				acceptData( elem ) ) {

				// Call a native DOM method on the target with the same name as the event.
				// Don't do default actions on window, that's where global variables be (#6170)
				if ( ontype && isFunction( elem[ type ] ) && !isWindow( elem ) ) {

					// Don't re-trigger an onFOO event when we call its FOO() method
					tmp = elem[ ontype ];

					if ( tmp ) {
						elem[ ontype ] = null;
					}

					// Prevent re-triggering of the same event, since we already bubbled it above
					jQuery.event.triggered = type;

					if ( event.isPropagationStopped() ) {
						lastElement.addEventListener( type, stopPropagationCallback );
					}

					elem[ type ]();

					if ( event.isPropagationStopped() ) {
						lastElement.removeEventListener( type, stopPropagationCallback );
					}

					jQuery.event.triggered = undefined;

					if ( tmp ) {
						elem[ ontype ] = tmp;
					}
				}
			}
		}

		return event.result;
	},

	// Piggyback on a donor event to simulate a different one
	// Used only for `focus(in | out)` events
	simulate: function( type, elem, event ) {
		var e = jQuery.extend(
			new jQuery.Event(),
			event,
			{
				type: type,
				isSimulated: true
			}
		);

		jQuery.event.trigger( e, null, elem );
	}

} );

jQuery.fn.extend( {

	trigger: function( type, data ) {
		return this.each( function() {
			jQuery.event.trigger( type, data, this );
		} );
	},
	triggerHandler: function( type, data ) {
		var elem = this[ 0 ];
		if ( elem ) {
			return jQuery.event.trigger( type, data, elem, true );
		}
	}
} );


// Support: Firefox <=44
// Firefox doesn't have focus(in | out) events
// Related ticket - https://bugzilla.mozilla.org/show_bug.cgi?id=687787
//
// Support: Chrome <=48 - 49, Safari <=9.0 - 9.1
// focus(in | out) events fire after focus & blur events,
// which is spec violation - http://www.w3.org/TR/DOM-Level-3-Events/#events-focusevent-event-order
// Related ticket - https://bugs.chromium.org/p/chromium/issues/detail?id=449857
if ( !support.focusin ) {
	jQuery.each( { focus: "focusin", blur: "focusout" }, function( orig, fix ) {

		// Attach a single capturing handler on the document while someone wants focusin/focusout
		var handler = function( event ) {
			jQuery.event.simulate( fix, event.target, jQuery.event.fix( event ) );
		};

		jQuery.event.special[ fix ] = {
			setup: function() {

				// Handle: regular nodes (via `this.ownerDocument`), window
				// (via `this.document`) & document (via `this`).
				var doc = this.ownerDocument || this.document || this,
					attaches = dataPriv.access( doc, fix );

				if ( !attaches ) {
					doc.addEventListener( orig, handler, true );
				}
				dataPriv.access( doc, fix, ( attaches || 0 ) + 1 );
			},
			teardown: function() {
				var doc = this.ownerDocument || this.document || this,
					attaches = dataPriv.access( doc, fix ) - 1;

				if ( !attaches ) {
					doc.removeEventListener( orig, handler, true );
					dataPriv.remove( doc, fix );

				} else {
					dataPriv.access( doc, fix, attaches );
				}
			}
		};
	} );
}
var location = window.location;

var nonce = { guid: Date.now() };

var rquery = ( /\?/ );



// Cross-browser xml parsing
jQuery.parseXML = function( data ) {
	var xml, parserErrorElem;
	if ( !data || typeof data !== "string" ) {
		return null;
	}

	// Support: IE 9 - 11 only
	// IE throws on parseFromString with invalid input.
	try {
		xml = ( new window.DOMParser() ).parseFromString( data, "text/xml" );
	} catch ( e ) {}

	parserErrorElem = xml && xml.getElementsByTagName( "parsererror" )[ 0 ];
	if ( !xml || parserErrorElem ) {
		jQuery.error( "Invalid XML: " + (
			parserErrorElem ?
				jQuery.map( parserErrorElem.childNodes, function( el ) {
					return el.textContent;
				} ).join( "\n" ) :
				data
		) );
	}
	return xml;
};


var
	rbracket = /\[\]$/,
	rCRLF = /\r?\n/g,
	rsubmitterTypes = /^(?:submit|button|image|reset|file)$/i,
	rsubmittable = /^(?:input|select|textarea|keygen)/i;

function buildParams( prefix, obj, traditional, add ) {
	var name;

	if ( Array.isArray( obj ) ) {

		// Serialize array item.
		jQuery.each( obj, function( i, v ) {
			if ( traditional || rbracket.test( prefix ) ) {

				// Treat each array item as a scalar.
				add( prefix, v );

			} else {

				// Item is non-scalar (array or object), encode its numeric index.
				buildParams(
					prefix + "[" + ( typeof v === "object" && v != null ? i : "" ) + "]",
					v,
					traditional,
					add
				);
			}
		} );

	} else if ( !traditional && toType( obj ) === "object" ) {

		// Serialize object item.
		for ( name in obj ) {
			buildParams( prefix + "[" + name + "]", obj[ name ], traditional, add );
		}

	} else {

		// Serialize scalar item.
		add( prefix, obj );
	}
}

// Serialize an array of form elements or a set of
// key/values into a query string
jQuery.param = function( a, traditional ) {
	var prefix,
		s = [],
		add = function( key, valueOrFunction ) {

			// If value is a function, invoke it and use its return value
			var value = isFunction( valueOrFunction ) ?
				valueOrFunction() :
				valueOrFunction;

			s[ s.length ] = encodeURIComponent( key ) + "=" +
				encodeURIComponent( value == null ? "" : value );
		};

	if ( a == null ) {
		return "";
	}

	// If an array was passed in, assume that it is an array of form elements.
	if ( Array.isArray( a ) || ( a.jquery && !jQuery.isPlainObject( a ) ) ) {

		// Serialize the form elements
		jQuery.each( a, function() {
			add( this.name, this.value );
		} );

	} else {

		// If traditional, encode the "old" way (the way 1.3.2 or older
		// did it), otherwise encode params recursively.
		for ( prefix in a ) {
			buildParams( prefix, a[ prefix ], traditional, add );
		}
	}

	// Return the resulting serialization
	return s.join( "&" );
};

jQuery.fn.extend( {
	serialize: function() {
		return jQuery.param( this.serializeArray() );
	},
	serializeArray: function() {
		return this.map( function() {

			// Can add propHook for "elements" to filter or add form elements
			var elements = jQuery.prop( this, "elements" );
			return elements ? jQuery.makeArray( elements ) : this;
		} ).filter( function() {
			var type = this.type;

			// Use .is( ":disabled" ) so that fieldset[disabled] works
			return this.name && !jQuery( this ).is( ":disabled" ) &&
				rsubmittable.test( this.nodeName ) && !rsubmitterTypes.test( type ) &&
				( this.checked || !rcheckableType.test( type ) );
		} ).map( function( _i, elem ) {
			var val = jQuery( this ).val();

			if ( val == null ) {
				return null;
			}

			if ( Array.isArray( val ) ) {
				return jQuery.map( val, function( val ) {
					return { name: elem.name, value: val.replace( rCRLF, "\r\n" ) };
				} );
			}

			return { name: elem.name, value: val.replace( rCRLF, "\r\n" ) };
		} ).get();
	}
} );


var
	r20 = /%20/g,
	rhash = /#.*$/,
	rantiCache = /([?&])_=[^&]*/,
	rheaders = /^(.*?):[ \t]*([^\r\n]*)$/mg,

	// #7653, #8125, #8152: local protocol detection
	rlocalProtocol = /^(?:about|app|app-storage|.+-extension|file|res|widget):$/,
	rnoContent = /^(?:GET|HEAD)$/,
	rprotocol = /^\/\//,

	/* Prefilters
	 * 1) They are useful to introduce custom dataTypes (see ajax/jsonp.js for an example)
	 * 2) These are called:
	 *    - BEFORE asking for a transport
	 *    - AFTER param serialization (s.data is a string if s.processData is true)
	 * 3) key is the dataType
	 * 4) the catchall symbol "*" can be used
	 * 5) execution will start with transport dataType and THEN continue down to "*" if needed
	 */
	prefilters = {},

	/* Transports bindings
	 * 1) key is the dataType
	 * 2) the catchall symbol "*" can be used
	 * 3) selection will start with transport dataType and THEN go to "*" if needed
	 */
	transports = {},

	// Avoid comment-prolog char sequence (#10098); must appease lint and evade compression
	allTypes = "*/".concat( "*" ),

	// Anchor tag for parsing the document origin
	originAnchor = document.createElement( "a" );

originAnchor.href = location.href;

// Base "constructor" for jQuery.ajaxPrefilter and jQuery.ajaxTransport
function addToPrefiltersOrTransports( structure ) {

	// dataTypeExpression is optional and defaults to "*"
	return function( dataTypeExpression, func ) {

		if ( typeof dataTypeExpression !== "string" ) {
			func = dataTypeExpression;
			dataTypeExpression = "*";
		}

		var dataType,
			i = 0,
			dataTypes = dataTypeExpression.toLowerCase().match( rnothtmlwhite ) || [];

		if ( isFunction( func ) ) {

			// For each dataType in the dataTypeExpression
			while ( ( dataType = dataTypes[ i++ ] ) ) {

				// Prepend if requested
				if ( dataType[ 0 ] === "+" ) {
					dataType = dataType.slice( 1 ) || "*";
					( structure[ dataType ] = structure[ dataType ] || [] ).unshift( func );

				// Otherwise append
				} else {
					( structure[ dataType ] = structure[ dataType ] || [] ).push( func );
				}
			}
		}
	};
}

// Base inspection function for prefilters and transports
function inspectPrefiltersOrTransports( structure, options, originalOptions, jqXHR ) {

	var inspected = {},
		seekingTransport = ( structure === transports );

	function inspect( dataType ) {
		var selected;
		inspected[ dataType ] = true;
		jQuery.each( structure[ dataType ] || [], function( _, prefilterOrFactory ) {
			var dataTypeOrTransport = prefilterOrFactory( options, originalOptions, jqXHR );
			if ( typeof dataTypeOrTransport === "string" &&
				!seekingTransport && !inspected[ dataTypeOrTransport ] ) {

				options.dataTypes.unshift( dataTypeOrTransport );
				inspect( dataTypeOrTransport );
				return false;
			} else if ( seekingTransport ) {
				return !( selected = dataTypeOrTransport );
			}
		} );
		return selected;
	}

	return inspect( options.dataTypes[ 0 ] ) || !inspected[ "*" ] && inspect( "*" );
}

// A special extend for ajax options
// that takes "flat" options (not to be deep extended)
// Fixes #9887
function ajaxExtend( target, src ) {
	var key, deep,
		flatOptions = jQuery.ajaxSettings.flatOptions || {};

	for ( key in src ) {
		if ( src[ key ] !== undefined ) {
			( flatOptions[ key ] ? target : ( deep || ( deep = {} ) ) )[ key ] = src[ key ];
		}
	}
	if ( deep ) {
		jQuery.extend( true, target, deep );
	}

	return target;
}

/* Handles responses to an ajax request:
 * - finds the right dataType (mediates between content-type and expected dataType)
 * - returns the corresponding response
 */
function ajaxHandleResponses( s, jqXHR, responses ) {

	var ct, type, finalDataType, firstDataType,
		contents = s.contents,
		dataTypes = s.dataTypes;

	// Remove auto dataType and get content-type in the process
	while ( dataTypes[ 0 ] === "*" ) {
		dataTypes.shift();
		if ( ct === undefined ) {
			ct = s.mimeType || jqXHR.getResponseHeader( "Content-Type" );
		}
	}

	// Check if we're dealing with a known content-type
	if ( ct ) {
		for ( type in contents ) {
			if ( contents[ type ] && contents[ type ].test( ct ) ) {
				dataTypes.unshift( type );
				break;
			}
		}
	}

	// Check to see if we have a response for the expected dataType
	if ( dataTypes[ 0 ] in responses ) {
		finalDataType = dataTypes[ 0 ];
	} else {

		// Try convertible dataTypes
		for ( type in responses ) {
			if ( !dataTypes[ 0 ] || s.converters[ type + " " + dataTypes[ 0 ] ] ) {
				finalDataType = type;
				break;
			}
			if ( !firstDataType ) {
				firstDataType = type;
			}
		}

		// Or just use first one
		finalDataType = finalDataType || firstDataType;
	}

	// If we found a dataType
	// We add the dataType to the list if needed
	// and return the corresponding response
	if ( finalDataType ) {
		if ( finalDataType !== dataTypes[ 0 ] ) {
			dataTypes.unshift( finalDataType );
		}
		return responses[ finalDataType ];
	}
}

/* Chain conversions given the request and the original response
 * Also sets the responseXXX fields on the jqXHR instance
 */
function ajaxConvert( s, response, jqXHR, isSuccess ) {
	var conv2, current, conv, tmp, prev,
		converters = {},

		// Work with a copy of dataTypes in case we need to modify it for conversion
		dataTypes = s.dataTypes.slice();

	// Create converters map with lowercased keys
	if ( dataTypes[ 1 ] ) {
		for ( conv in s.converters ) {
			converters[ conv.toLowerCase() ] = s.converters[ conv ];
		}
	}

	current = dataTypes.shift();

	// Convert to each sequential dataType
	while ( current ) {

		if ( s.responseFields[ current ] ) {
			jqXHR[ s.responseFields[ current ] ] = response;
		}

		// Apply the dataFilter if provided
		if ( !prev && isSuccess && s.dataFilter ) {
			response = s.dataFilter( response, s.dataType );
		}

		prev = current;
		current = dataTypes.shift();

		if ( current ) {

			// There's only work to do if current dataType is non-auto
			if ( current === "*" ) {

				current = prev;

			// Convert response if prev dataType is non-auto and differs from current
			} else if ( prev !== "*" && prev !== current ) {

				// Seek a direct converter
				conv = converters[ prev + " " + current ] || converters[ "* " + current ];

				// If none found, seek a pair
				if ( !conv ) {
					for ( conv2 in converters ) {

						// If conv2 outputs current
						tmp = conv2.split( " " );
						if ( tmp[ 1 ] === current ) {

							// If prev can be converted to accepted input
							conv = converters[ prev + " " + tmp[ 0 ] ] ||
								converters[ "* " + tmp[ 0 ] ];
							if ( conv ) {

								// Condense equivalence converters
								if ( conv === true ) {
									conv = converters[ conv2 ];

								// Otherwise, insert the intermediate dataType
								} else if ( converters[ conv2 ] !== true ) {
									current = tmp[ 0 ];
									dataTypes.unshift( tmp[ 1 ] );
								}
								break;
							}
						}
					}
				}

				// Apply converter (if not an equivalence)
				if ( conv !== true ) {

					// Unless errors are allowed to bubble, catch and return them
					if ( conv && s.throws ) {
						response = conv( response );
					} else {
						try {
							response = conv( response );
						} catch ( e ) {
							return {
								state: "parsererror",
								error: conv ? e : "No conversion from " + prev + " to " + current
							};
						}
					}
				}
			}
		}
	}

	return { state: "success", data: response };
}

jQuery.extend( {

	// Counter for holding the number of active queries
	active: 0,

	// Last-Modified header cache for next request
	lastModified: {},
	etag: {},

	ajaxSettings: {
		url: location.href,
		type: "GET",
		isLocal: rlocalProtocol.test( location.protocol ),
		global: true,
		processData: true,
		async: true,
		contentType: "application/x-www-form-urlencoded; charset=UTF-8",

		/*
		timeout: 0,
		data: null,
		dataType: null,
		username: null,
		password: null,
		cache: null,
		throws: false,
		traditional: false,
		headers: {},
		*/

		accepts: {
			"*": allTypes,
			text: "text/plain",
			html: "text/html",
			xml: "application/xml, text/xml",
			json: "application/json, text/javascript"
		},

		contents: {
			xml: /\bxml\b/,
			html: /\bhtml/,
			json: /\bjson\b/
		},

		responseFields: {
			xml: "responseXML",
			text: "responseText",
			json: "responseJSON"
		},

		// Data converters
		// Keys separate source (or catchall "*") and destination types with a single space
		converters: {

			// Convert anything to text
			"* text": String,

			// Text to html (true = no transformation)
			"text html": true,

			// Evaluate text as a json expression
			"text json": JSON.parse,

			// Parse text as xml
			"text xml": jQuery.parseXML
		},

		// For options that shouldn't be deep extended:
		// you can add your own custom options here if
		// and when you create one that shouldn't be
		// deep extended (see ajaxExtend)
		flatOptions: {
			url: true,
			context: true
		}
	},

	// Creates a full fledged settings object into target
	// with both ajaxSettings and settings fields.
	// If target is omitted, writes into ajaxSettings.
	ajaxSetup: function( target, settings ) {
		return settings ?

			// Building a settings object
			ajaxExtend( ajaxExtend( target, jQuery.ajaxSettings ), settings ) :

			// Extending ajaxSettings
			ajaxExtend( jQuery.ajaxSettings, target );
	},

	ajaxPrefilter: addToPrefiltersOrTransports( prefilters ),
	ajaxTransport: addToPrefiltersOrTransports( transports ),

	// Main method
	ajax: function( url, options ) {

		// If url is an object, simulate pre-1.5 signature
		if ( typeof url === "object" ) {
			options = url;
			url = undefined;
		}

		// Force options to be an object
		options = options || {};

		var transport,

			// URL without anti-cache param
			cacheURL,

			// Response headers
			responseHeadersString,
			responseHeaders,

			// timeout handle
			timeoutTimer,

			// Url cleanup var
			urlAnchor,

			// Request state (becomes false upon send and true upon completion)
			completed,

			// To know if global events are to be dispatched
			fireGlobals,

			// Loop variable
			i,

			// uncached part of the url
			uncached,

			// Create the final options object
			s = jQuery.ajaxSetup( {}, options ),

			// Callbacks context
			callbackContext = s.context || s,

			// Context for global events is callbackContext if it is a DOM node or jQuery collection
			globalEventContext = s.context &&
				( callbackContext.nodeType || callbackContext.jquery ) ?
				jQuery( callbackContext ) :
				jQuery.event,

			// Deferreds
			deferred = jQuery.Deferred(),
			completeDeferred = jQuery.Callbacks( "once memory" ),

			// Status-dependent callbacks
			statusCode = s.statusCode || {},

			// Headers (they are sent all at once)
			requestHeaders = {},
			requestHeadersNames = {},

			// Default abort message
			strAbort = "canceled",

			// Fake xhr
			jqXHR = {
				readyState: 0,

				// Builds headers hashtable if needed
				getResponseHeader: function( key ) {
					var match;
					if ( completed ) {
						if ( !responseHeaders ) {
							responseHeaders = {};
							while ( ( match = rheaders.exec( responseHeadersString ) ) ) {
								responseHeaders[ match[ 1 ].toLowerCase() + " " ] =
									( responseHeaders[ match[ 1 ].toLowerCase() + " " ] || [] )
										.concat( match[ 2 ] );
							}
						}
						match = responseHeaders[ key.toLowerCase() + " " ];
					}
					return match == null ? null : match.join( ", " );
				},

				// Raw string
				getAllResponseHeaders: function() {
					return completed ? responseHeadersString : null;
				},

				// Caches the header
				setRequestHeader: function( name, value ) {
					if ( completed == null ) {
						name = requestHeadersNames[ name.toLowerCase() ] =
							requestHeadersNames[ name.toLowerCase() ] || name;
						requestHeaders[ name ] = value;
					}
					return this;
				},

				// Overrides response content-type header
				overrideMimeType: function( type ) {
					if ( completed == null ) {
						s.mimeType = type;
					}
					return this;
				},

				// Status-dependent callbacks
				statusCode: function( map ) {
					var code;
					if ( map ) {
						if ( completed ) {

							// Execute the appropriate callbacks
							jqXHR.always( map[ jqXHR.status ] );
						} else {

							// Lazy-add the new callbacks in a way that preserves old ones
							for ( code in map ) {
								statusCode[ code ] = [ statusCode[ code ], map[ code ] ];
							}
						}
					}
					return this;
				},

				// Cancel the request
				abort: function( statusText ) {
					var finalText = statusText || strAbort;
					if ( transport ) {
						transport.abort( finalText );
					}
					done( 0, finalText );
					return this;
				}
			};

		// Attach deferreds
		deferred.promise( jqXHR );

		// Add protocol if not provided (prefilters might expect it)
		// Handle falsy url in the settings object (#10093: consistency with old signature)
		// We also use the url parameter if available
		s.url = ( ( url || s.url || location.href ) + "" )
			.replace( rprotocol, location.protocol + "//" );

		// Alias method option to type as per ticket #12004
		s.type = options.method || options.type || s.method || s.type;

		// Extract dataTypes list
		s.dataTypes = ( s.dataType || "*" ).toLowerCase().match( rnothtmlwhite ) || [ "" ];

		// A cross-domain request is in order when the origin doesn't match the current origin.
		if ( s.crossDomain == null ) {
			urlAnchor = document.createElement( "a" );

			// Support: IE <=8 - 11, Edge 12 - 15
			// IE throws exception on accessing the href property if url is malformed,
			// e.g. http://example.com:80x/
			try {
				urlAnchor.href = s.url;

				// Support: IE <=8 - 11 only
				// Anchor's host property isn't correctly set when s.url is relative
				urlAnchor.href = urlAnchor.href;
				s.crossDomain = originAnchor.protocol + "//" + originAnchor.host !==
					urlAnchor.protocol + "//" + urlAnchor.host;
			} catch ( e ) {

				// If there is an error parsing the URL, assume it is crossDomain,
				// it can be rejected by the transport if it is invalid
				s.crossDomain = true;
			}
		}

		// Convert data if not already a string
		if ( s.data && s.processData && typeof s.data !== "string" ) {
			s.data = jQuery.param( s.data, s.traditional );
		}

		// Apply prefilters
		inspectPrefiltersOrTransports( prefilters, s, options, jqXHR );

		// If request was aborted inside a prefilter, stop there
		if ( completed ) {
			return jqXHR;
		}

		// We can fire global events as of now if asked to
		// Don't fire events if jQuery.event is undefined in an AMD-usage scenario (#15118)
		fireGlobals = jQuery.event && s.global;

		// Watch for a new set of requests
		if ( fireGlobals && jQuery.active++ === 0 ) {
			jQuery.event.trigger( "ajaxStart" );
		}

		// Uppercase the type
		s.type = s.type.toUpperCase();

		// Determine if request has content
		s.hasContent = !rnoContent.test( s.type );

		// Save the URL in case we're toying with the If-Modified-Since
		// and/or If-None-Match header later on
		// Remove hash to simplify url manipulation
		cacheURL = s.url.replace( rhash, "" );

		// More options handling for requests with no content
		if ( !s.hasContent ) {

			// Remember the hash so we can put it back
			uncached = s.url.slice( cacheURL.length );

			// If data is available and should be processed, append data to url
			if ( s.data && ( s.processData || typeof s.data === "string" ) ) {
				cacheURL += ( rquery.test( cacheURL ) ? "&" : "?" ) + s.data;

				// #9682: remove data so that it's not used in an eventual retry
				delete s.data;
			}

			// Add or update anti-cache param if needed
			if ( s.cache === false ) {
				cacheURL = cacheURL.replace( rantiCache, "$1" );
				uncached = ( rquery.test( cacheURL ) ? "&" : "?" ) + "_=" + ( nonce.guid++ ) +
					uncached;
			}

			// Put hash and anti-cache on the URL that will be requested (gh-1732)
			s.url = cacheURL + uncached;

		// Change '%20' to '+' if this is encoded form body content (gh-2658)
		} else if ( s.data && s.processData &&
			( s.contentType || "" ).indexOf( "application/x-www-form-urlencoded" ) === 0 ) {
			s.data = s.data.replace( r20, "+" );
		}

		// Set the If-Modified-Since and/or If-None-Match header, if in ifModified mode.
		if ( s.ifModified ) {
			if ( jQuery.lastModified[ cacheURL ] ) {
				jqXHR.setRequestHeader( "If-Modified-Since", jQuery.lastModified[ cacheURL ] );
			}
			if ( jQuery.etag[ cacheURL ] ) {
				jqXHR.setRequestHeader( "If-None-Match", jQuery.etag[ cacheURL ] );
			}
		}

		// Set the correct header, if data is being sent
		if ( s.data && s.hasContent && s.contentType !== false || options.contentType ) {
			jqXHR.setRequestHeader( "Content-Type", s.contentType );
		}

		// Set the Accepts header for the server, depending on the dataType
		jqXHR.setRequestHeader(
			"Accept",
			s.dataTypes[ 0 ] && s.accepts[ s.dataTypes[ 0 ] ] ?
				s.accepts[ s.dataTypes[ 0 ] ] +
					( s.dataTypes[ 0 ] !== "*" ? ", " + allTypes + "; q=0.01" : "" ) :
				s.accepts[ "*" ]
		);

		// Check for headers option
		for ( i in s.headers ) {
			jqXHR.setRequestHeader( i, s.headers[ i ] );
		}

		// Allow custom headers/mimetypes and early abort
		if ( s.beforeSend &&
			( s.beforeSend.call( callbackContext, jqXHR, s ) === false || completed ) ) {

			// Abort if not done already and return
			return jqXHR.abort();
		}

		// Aborting is no longer a cancellation
		strAbort = "abort";

		// Install callbacks on deferreds
		completeDeferred.add( s.complete );
		jqXHR.done( s.success );
		jqXHR.fail( s.error );

		// Get transport
		transport = inspectPrefiltersOrTransports( transports, s, options, jqXHR );

		// If no transport, we auto-abort
		if ( !transport ) {
			done( -1, "No Transport" );
		} else {
			jqXHR.readyState = 1;

			// Send global event
			if ( fireGlobals ) {
				globalEventContext.trigger( "ajaxSend", [ jqXHR, s ] );
			}

			// If request was aborted inside ajaxSend, stop there
			if ( completed ) {
				return jqXHR;
			}

			// Timeout
			if ( s.async && s.timeout > 0 ) {
				timeoutTimer = window.setTimeout( function() {
					jqXHR.abort( "timeout" );
				}, s.timeout );
			}

			try {
				completed = false;
				transport.send( requestHeaders, done );
			} catch ( e ) {

				// Rethrow post-completion exceptions
				if ( completed ) {
					throw e;
				}

				// Propagate others as results
				done( -1, e );
			}
		}

		// Callback for when everything is done
		function done( status, nativeStatusText, responses, headers ) {
			var isSuccess, success, error, response, modified,
				statusText = nativeStatusText;

			// Ignore repeat invocations
			if ( completed ) {
				return;
			}

			completed = true;

			// Clear timeout if it exists
			if ( timeoutTimer ) {
				window.clearTimeout( timeoutTimer );
			}

			// Dereference transport for early garbage collection
			// (no matter how long the jqXHR object will be used)
			transport = undefined;

			// Cache response headers
			responseHeadersString = headers || "";

			// Set readyState
			jqXHR.readyState = status > 0 ? 4 : 0;

			// Determine if successful
			isSuccess = status >= 200 && status < 300 || status === 304;

			// Get response data
			if ( responses ) {
				response = ajaxHandleResponses( s, jqXHR, responses );
			}

			// Use a noop converter for missing script but not if jsonp
			if ( !isSuccess &&
				jQuery.inArray( "script", s.dataTypes ) > -1 &&
				jQuery.inArray( "json", s.dataTypes ) < 0 ) {
				s.converters[ "text script" ] = function() {};
			}

			// Convert no matter what (that way responseXXX fields are always set)
			response = ajaxConvert( s, response, jqXHR, isSuccess );

			// If successful, handle type chaining
			if ( isSuccess ) {

				// Set the If-Modified-Since and/or If-None-Match header, if in ifModified mode.
				if ( s.ifModified ) {
					modified = jqXHR.getResponseHeader( "Last-Modified" );
					if ( modified ) {
						jQuery.lastModified[ cacheURL ] = modified;
					}
					modified = jqXHR.getResponseHeader( "etag" );
					if ( modified ) {
						jQuery.etag[ cacheURL ] = modified;
					}
				}

				// if no content
				if ( status === 204 || s.type === "HEAD" ) {
					statusText = "nocontent";

				// if not modified
				} else if ( status === 304 ) {
					statusText = "notmodified";

				// If we have data, let's convert it
				} else {
					statusText = response.state;
					success = response.data;
					error = response.error;
					isSuccess = !error;
				}
			} else {

				// Extract error from statusText and normalize for non-aborts
				error = statusText;
				if ( status || !statusText ) {
					statusText = "error";
					if ( status < 0 ) {
						status = 0;
					}
				}
			}

			// Set data for the fake xhr object
			jqXHR.status = status;
			jqXHR.statusText = ( nativeStatusText || statusText ) + "";

			// Success/Error
			if ( isSuccess ) {
				deferred.resolveWith( callbackContext, [ success, statusText, jqXHR ] );
			} else {
				deferred.rejectWith( callbackContext, [ jqXHR, statusText, error ] );
			}

			// Status-dependent callbacks
			jqXHR.statusCode( statusCode );
			statusCode = undefined;

			if ( fireGlobals ) {
				globalEventContext.trigger( isSuccess ? "ajaxSuccess" : "ajaxError",
					[ jqXHR, s, isSuccess ? success : error ] );
			}

			// Complete
			completeDeferred.fireWith( callbackContext, [ jqXHR, statusText ] );

			if ( fireGlobals ) {
				globalEventContext.trigger( "ajaxComplete", [ jqXHR, s ] );

				// Handle the global AJAX counter
				if ( !( --jQuery.active ) ) {
					jQuery.event.trigger( "ajaxStop" );
				}
			}
		}

		return jqXHR;
	},

	getJSON: function( url, data, callback ) {
		return jQuery.get( url, data, callback, "json" );
	},

	getScript: function( url, callback ) {
		return jQuery.get( url, undefined, callback, "script" );
	}
} );

jQuery.each( [ "get", "post" ], function( _i, method ) {
	jQuery[ method ] = function( url, data, callback, type ) {

		// Shift arguments if data argument was omitted
		if ( isFunction( data ) ) {
			type = type || callback;
			callback = data;
			data = undefined;
		}

		// The url can be an options object (which then must have .url)
		return jQuery.ajax( jQuery.extend( {
			url: url,
			type: method,
			dataType: type,
			data: data,
			success: callback
		}, jQuery.isPlainObject( url ) && url ) );
	};
} );

jQuery.ajaxPrefilter( function( s ) {
	var i;
	for ( i in s.headers ) {
		if ( i.toLowerCase() === "content-type" ) {
			s.contentType = s.headers[ i ] || "";
		}
	}
} );


jQuery._evalUrl = function( url, options, doc ) {
	return jQuery.ajax( {
		url: url,

		// Make this explicit, since user can override this through ajaxSetup (#11264)
		type: "GET",
		dataType: "script",
		cache: true,
		async: false,
		global: false,

		// Only evaluate the response if it is successful (gh-4126)
		// dataFilter is not invoked for failure responses, so using it instead
		// of the default converter is kludgy but it works.
		converters: {
			"text script": function() {}
		},
		dataFilter: function( response ) {
			jQuery.globalEval( response, options, doc );
		}
	} );
};


jQuery.fn.extend( {
	wrapAll: function( html ) {
		var wrap;

		if ( this[ 0 ] ) {
			if ( isFunction( html ) ) {
				html = html.call( this[ 0 ] );
			}

			// The elements to wrap the target around
			wrap = jQuery( html, this[ 0 ].ownerDocument ).eq( 0 ).clone( true );

			if ( this[ 0 ].parentNode ) {
				wrap.insertBefore( this[ 0 ] );
			}

			wrap.map( function() {
				var elem = this;

				while ( elem.firstElementChild ) {
					elem = elem.firstElementChild;
				}

				return elem;
			} ).append( this );
		}

		return this;
	},

	wrapInner: function( html ) {
		if ( isFunction( html ) ) {
			return this.each( function( i ) {
				jQuery( this ).wrapInner( html.call( this, i ) );
			} );
		}

		return this.each( function() {
			var self = jQuery( this ),
				contents = self.contents();

			if ( contents.length ) {
				contents.wrapAll( html );

			} else {
				self.append( html );
			}
		} );
	},

	wrap: function( html ) {
		var htmlIsFunction = isFunction( html );

		return this.each( function( i ) {
			jQuery( this ).wrapAll( htmlIsFunction ? html.call( this, i ) : html );
		} );
	},

	unwrap: function( selector ) {
		this.parent( selector ).not( "body" ).each( function() {
			jQuery( this ).replaceWith( this.childNodes );
		} );
		return this;
	}
} );


jQuery.expr.pseudos.hidden = function( elem ) {
	return !jQuery.expr.pseudos.visible( elem );
};
jQuery.expr.pseudos.visible = function( elem ) {
	return !!( elem.offsetWidth || elem.offsetHeight || elem.getClientRects().length );
};




jQuery.ajaxSettings.xhr = function() {
	try {
		return new window.XMLHttpRequest();
	} catch ( e ) {}
};

var xhrSuccessStatus = {

		// File protocol always yields status code 0, assume 200
		0: 200,

		// Support: IE <=9 only
		// #1450: sometimes IE returns 1223 when it should be 204
		1223: 204
	},
	xhrSupported = jQuery.ajaxSettings.xhr();

support.cors = !!xhrSupported && ( "withCredentials" in xhrSupported );
support.ajax = xhrSupported = !!xhrSupported;

jQuery.ajaxTransport( function( options ) {
	var callback, errorCallback;

	// Cross domain only allowed if supported through XMLHttpRequest
	if ( support.cors || xhrSupported && !options.crossDomain ) {
		return {
			send: function( headers, complete ) {
				var i,
					xhr = options.xhr();

				xhr.open(
					options.type,
					options.url,
					options.async,
					options.username,
					options.password
				);

				// Apply custom fields if provided
				if ( options.xhrFields ) {
					for ( i in options.xhrFields ) {
						xhr[ i ] = options.xhrFields[ i ];
					}
				}

				// Override mime type if needed
				if ( options.mimeType && xhr.overrideMimeType ) {
					xhr.overrideMimeType( options.mimeType );
				}

				// X-Requested-With header
				// For cross-domain requests, seeing as conditions for a preflight are
				// akin to a jigsaw puzzle, we simply never set it to be sure.
				// (it can always be set on a per-request basis or even using ajaxSetup)
				// For same-domain requests, won't change header if already provided.
				if ( !options.crossDomain && !headers[ "X-Requested-With" ] ) {
					headers[ "X-Requested-With" ] = "XMLHttpRequest";
				}

				// Set headers
				for ( i in headers ) {
					xhr.setRequestHeader( i, headers[ i ] );
				}

				// Callback
				callback = function( type ) {
					return function() {
						if ( callback ) {
							callback = errorCallback = xhr.onload =
								xhr.onerror = xhr.onabort = xhr.ontimeout =
									xhr.onreadystatechange = null;

							if ( type === "abort" ) {
								xhr.abort();
							} else if ( type === "error" ) {

								// Support: IE <=9 only
								// On a manual native abort, IE9 throws
								// errors on any property access that is not readyState
								if ( typeof xhr.status !== "number" ) {
									complete( 0, "error" );
								} else {
									complete(

										// File: protocol always yields status 0; see #8605, #14207
										xhr.status,
										xhr.statusText
									);
								}
							} else {
								complete(
									xhrSuccessStatus[ xhr.status ] || xhr.status,
									xhr.statusText,

									// Support: IE <=9 only
									// IE9 has no XHR2 but throws on binary (trac-11426)
									// For XHR2 non-text, let the caller handle it (gh-2498)
									( xhr.responseType || "text" ) !== "text"  ||
									typeof xhr.responseText !== "string" ?
										{ binary: xhr.response } :
										{ text: xhr.responseText },
									xhr.getAllResponseHeaders()
								);
							}
						}
					};
				};

				// Listen to events
				xhr.onload = callback();
				errorCallback = xhr.onerror = xhr.ontimeout = callback( "error" );

				// Support: IE 9 only
				// Use onreadystatechange to replace onabort
				// to handle uncaught aborts
				if ( xhr.onabort !== undefined ) {
					xhr.onabort = errorCallback;
				} else {
					xhr.onreadystatechange = function() {

						// Check readyState before timeout as it changes
						if ( xhr.readyState === 4 ) {

							// Allow onerror to be called first,
							// but that will not handle a native abort
							// Also, save errorCallback to a variable
							// as xhr.onerror cannot be accessed
							window.setTimeout( function() {
								if ( callback ) {
									errorCallback();
								}
							} );
						}
					};
				}

				// Create the abort callback
				callback = callback( "abort" );

				try {

					// Do send the request (this may raise an exception)
					xhr.send( options.hasContent && options.data || null );
				} catch ( e ) {

					// #14683: Only rethrow if this hasn't been notified as an error yet
					if ( callback ) {
						throw e;
					}
				}
			},

			abort: function() {
				if ( callback ) {
					callback();
				}
			}
		};
	}
} );




// Prevent auto-execution of scripts when no explicit dataType was provided (See gh-2432)
jQuery.ajaxPrefilter( function( s ) {
	if ( s.crossDomain ) {
		s.contents.script = false;
	}
} );

// Install script dataType
jQuery.ajaxSetup( {
	accepts: {
		script: "text/javascript, application/javascript, " +
			"application/ecmascript, application/x-ecmascript"
	},
	contents: {
		script: /\b(?:java|ecma)script\b/
	},
	converters: {
		"text script": function( text ) {
			jQuery.globalEval( text );
			return text;
		}
	}
} );

// Handle cache's special case and crossDomain
jQuery.ajaxPrefilter( "script", function( s ) {
	if ( s.cache === undefined ) {
		s.cache = false;
	}
	if ( s.crossDomain ) {
		s.type = "GET";
	}
} );

// Bind script tag hack transport
jQuery.ajaxTransport( "script", function( s ) {

	// This transport only deals with cross domain or forced-by-attrs requests
	if ( s.crossDomain || s.scriptAttrs ) {
		var script, callback;
		return {
			send: function( _, complete ) {
				script = jQuery( "<script>" )
					.attr( s.scriptAttrs || {} )
					.prop( { charset: s.scriptCharset, src: s.url } )
					.on( "load error", callback = function( evt ) {
						script.remove();
						callback = null;
						if ( evt ) {
							complete( evt.type === "error" ? 404 : 200, evt.type );
						}
					} );

				// Use native DOM manipulation to avoid our domManip AJAX trickery
				document.head.appendChild( script[ 0 ] );
			},
			abort: function() {
				if ( callback ) {
					callback();
				}
			}
		};
	}
} );




var oldCallbacks = [],
	rjsonp = /(=)\?(?=&|$)|\?\?/;

// Default jsonp settings
jQuery.ajaxSetup( {
	jsonp: "callback",
	jsonpCallback: function() {
		var callback = oldCallbacks.pop() || ( jQuery.expando + "_" + ( nonce.guid++ ) );
		this[ callback ] = true;
		return callback;
	}
} );

// Detect, normalize options and install callbacks for jsonp requests
jQuery.ajaxPrefilter( "json jsonp", function( s, originalSettings, jqXHR ) {

	var callbackName, overwritten, responseContainer,
		jsonProp = s.jsonp !== false && ( rjsonp.test( s.url ) ?
			"url" :
			typeof s.data === "string" &&
				( s.contentType || "" )
					.indexOf( "application/x-www-form-urlencoded" ) === 0 &&
				rjsonp.test( s.data ) && "data"
		);

	// Handle iff the expected data type is "jsonp" or we have a parameter to set
	if ( jsonProp || s.dataTypes[ 0 ] === "jsonp" ) {

		// Get callback name, remembering preexisting value associated with it
		callbackName = s.jsonpCallback = isFunction( s.jsonpCallback ) ?
			s.jsonpCallback() :
			s.jsonpCallback;

		// Insert callback into url or form data
		if ( jsonProp ) {
			s[ jsonProp ] = s[ jsonProp ].replace( rjsonp, "$1" + callbackName );
		} else if ( s.jsonp !== false ) {
			s.url += ( rquery.test( s.url ) ? "&" : "?" ) + s.jsonp + "=" + callbackName;
		}

		// Use data converter to retrieve json after script execution
		s.converters[ "script json" ] = function() {
			if ( !responseContainer ) {
				jQuery.error( callbackName + " was not called" );
			}
			return responseContainer[ 0 ];
		};

		// Force json dataType
		s.dataTypes[ 0 ] = "json";

		// Install callback
		overwritten = window[ callbackName ];
		window[ callbackName ] = function() {
			responseContainer = arguments;
		};

		// Clean-up function (fires after converters)
		jqXHR.always( function() {

			// If previous value didn't exist - remove it
			if ( overwritten === undefined ) {
				jQuery( window ).removeProp( callbackName );

			// Otherwise restore preexisting value
			} else {
				window[ callbackName ] = overwritten;
			}

			// Save back as free
			if ( s[ callbackName ] ) {

				// Make sure that re-using the options doesn't screw things around
				s.jsonpCallback = originalSettings.jsonpCallback;

				// Save the callback name for future use
				oldCallbacks.push( callbackName );
			}

			// Call if it was a function and we have a response
			if ( responseContainer && isFunction( overwritten ) ) {
				overwritten( responseContainer[ 0 ] );
			}

			responseContainer = overwritten = undefined;
		} );

		// Delegate to script
		return "script";
	}
} );




// Support: Safari 8 only
// In Safari 8 documents created via document.implementation.createHTMLDocument
// collapse sibling forms: the second one becomes a child of the first one.
// Because of that, this security measure has to be disabled in Safari 8.
// https://bugs.webkit.org/show_bug.cgi?id=137337
support.createHTMLDocument = ( function() {
	var body = document.implementation.createHTMLDocument( "" ).body;
	body.innerHTML = "<form></form><form></form>";
	return body.childNodes.length === 2;
} )();


// Argument "data" should be string of html
// context (optional): If specified, the fragment will be created in this context,
// defaults to document
// keepScripts (optional): If true, will include scripts passed in the html string
jQuery.parseHTML = function( data, context, keepScripts ) {
	if ( typeof data !== "string" ) {
		return [];
	}
	if ( typeof context === "boolean" ) {
		keepScripts = context;
		context = false;
	}

	var base, parsed, scripts;

	if ( !context ) {

		// Stop scripts or inline event handlers from being executed immediately
		// by using document.implementation
		if ( support.createHTMLDocument ) {
			context = document.implementation.createHTMLDocument( "" );

			// Set the base href for the created document
			// so any parsed elements with URLs
			// are based on the document's URL (gh-2965)
			base = context.createElement( "base" );
			base.href = document.location.href;
			context.head.appendChild( base );
		} else {
			context = document;
		}
	}

	parsed = rsingleTag.exec( data );
	scripts = !keepScripts && [];

	// Single tag
	if ( parsed ) {
		return [ context.createElement( parsed[ 1 ] ) ];
	}

	parsed = buildFragment( [ data ], context, scripts );

	if ( scripts && scripts.length ) {
		jQuery( scripts ).remove();
	}

	return jQuery.merge( [], parsed.childNodes );
};


/**
 * Load a url into a page
 */
jQuery.fn.load = function( url, params, callback ) {
	var selector, type, response,
		self = this,
		off = url.indexOf( " " );

	if ( off > -1 ) {
		selector = stripAndCollapse( url.slice( off ) );
		url = url.slice( 0, off );
	}

	// If it's a function
	if ( isFunction( params ) ) {

		// We assume that it's the callback
		callback = params;
		params = undefined;

	// Otherwise, build a param string
	} else if ( params && typeof params === "object" ) {
		type = "POST";
	}

	// If we have elements to modify, make the request
	if ( self.length > 0 ) {
		jQuery.ajax( {
			url: url,

			// If "type" variable is undefined, then "GET" method will be used.
			// Make value of this field explicit since
			// user can override it through ajaxSetup method
			type: type || "GET",
			dataType: "html",
			data: params
		} ).done( function( responseText ) {

			// Save response for use in complete callback
			response = arguments;

			self.html( selector ?

				// If a selector was specified, locate the right elements in a dummy div
				// Exclude scripts to avoid IE 'Permission Denied' errors
				jQuery( "<div>" ).append( jQuery.parseHTML( responseText ) ).find( selector ) :

				// Otherwise use the full result
				responseText );

		// If the request succeeds, this function gets "data", "status", "jqXHR"
		// but they are ignored because response was set above.
		// If it fails, this function gets "jqXHR", "status", "error"
		} ).always( callback && function( jqXHR, status ) {
			self.each( function() {
				callback.apply( this, response || [ jqXHR.responseText, status, jqXHR ] );
			} );
		} );
	}

	return this;
};




jQuery.expr.pseudos.animated = function( elem ) {
	return jQuery.grep( jQuery.timers, function( fn ) {
		return elem === fn.elem;
	} ).length;
};




jQuery.offset = {
	setOffset: function( elem, options, i ) {
		var curPosition, curLeft, curCSSTop, curTop, curOffset, curCSSLeft, calculatePosition,
			position = jQuery.css( elem, "position" ),
			curElem = jQuery( elem ),
			props = {};

		// Set position first, in-case top/left are set even on static elem
		if ( position === "static" ) {
			elem.style.position = "relative";
		}

		curOffset = curElem.offset();
		curCSSTop = jQuery.css( elem, "top" );
		curCSSLeft = jQuery.css( elem, "left" );
		calculatePosition = ( position === "absolute" || position === "fixed" ) &&
			( curCSSTop + curCSSLeft ).indexOf( "auto" ) > -1;

		// Need to be able to calculate position if either
		// top or left is auto and position is either absolute or fixed
		if ( calculatePosition ) {
			curPosition = curElem.position();
			curTop = curPosition.top;
			curLeft = curPosition.left;

		} else {
			curTop = parseFloat( curCSSTop ) || 0;
			curLeft = parseFloat( curCSSLeft ) || 0;
		}

		if ( isFunction( options ) ) {

			// Use jQuery.extend here to allow modification of coordinates argument (gh-1848)
			options = options.call( elem, i, jQuery.extend( {}, curOffset ) );
		}

		if ( options.top != null ) {
			props.top = ( options.top - curOffset.top ) + curTop;
		}
		if ( options.left != null ) {
			props.left = ( options.left - curOffset.left ) + curLeft;
		}

		if ( "using" in options ) {
			options.using.call( elem, props );

		} else {
			curElem.css( props );
		}
	}
};

jQuery.fn.extend( {

	// offset() relates an element's border box to the document origin
	offset: function( options ) {

		// Preserve chaining for setter
		if ( arguments.length ) {
			return options === undefined ?
				this :
				this.each( function( i ) {
					jQuery.offset.setOffset( this, options, i );
				} );
		}

		var rect, win,
			elem = this[ 0 ];

		if ( !elem ) {
			return;
		}

		// Return zeros for disconnected and hidden (display: none) elements (gh-2310)
		// Support: IE <=11 only
		// Running getBoundingClientRect on a
		// disconnected node in IE throws an error
		if ( !elem.getClientRects().length ) {
			return { top: 0, left: 0 };
		}

		// Get document-relative position by adding viewport scroll to viewport-relative gBCR
		rect = elem.getBoundingClientRect();
		win = elem.ownerDocument.defaultView;
		return {
			top: rect.top + win.pageYOffset,
			left: rect.left + win.pageXOffset
		};
	},

	// position() relates an element's margin box to its offset parent's padding box
	// This corresponds to the behavior of CSS absolute positioning
	position: function() {
		if ( !this[ 0 ] ) {
			return;
		}

		var offsetParent, offset, doc,
			elem = this[ 0 ],
			parentOffset = { top: 0, left: 0 };

		// position:fixed elements are offset from the viewport, which itself always has zero offset
		if ( jQuery.css( elem, "position" ) === "fixed" ) {

			// Assume position:fixed implies availability of getBoundingClientRect
			offset = elem.getBoundingClientRect();

		} else {
			offset = this.offset();

			// Account for the *real* offset parent, which can be the document or its root element
			// when a statically positioned element is identified
			doc = elem.ownerDocument;
			offsetParent = elem.offsetParent || doc.documentElement;
			while ( offsetParent &&
				( offsetParent === doc.body || offsetParent === doc.documentElement ) &&
				jQuery.css( offsetParent, "position" ) === "static" ) {

				offsetParent = offsetParent.parentNode;
			}
			if ( offsetParent && offsetParent !== elem && offsetParent.nodeType === 1 ) {

				// Incorporate borders into its offset, since they are outside its content origin
				parentOffset = jQuery( offsetParent ).offset();
				parentOffset.top += jQuery.css( offsetParent, "borderTopWidth", true );
				parentOffset.left += jQuery.css( offsetParent, "borderLeftWidth", true );
			}
		}

		// Subtract parent offsets and element margins
		return {
			top: offset.top - parentOffset.top - jQuery.css( elem, "marginTop", true ),
			left: offset.left - parentOffset.left - jQuery.css( elem, "marginLeft", true )
		};
	},

	// This method will return documentElement in the following cases:
	// 1) For the element inside the iframe without offsetParent, this method will return
	//    documentElement of the parent window
	// 2) For the hidden or detached element
	// 3) For body or html element, i.e. in case of the html node - it will return itself
	//
	// but those exceptions were never presented as a real life use-cases
	// and might be considered as more preferable results.
	//
	// This logic, however, is not guaranteed and can change at any point in the future
	offsetParent: function() {
		return this.map( function() {
			var offsetParent = this.offsetParent;

			while ( offsetParent && jQuery.css( offsetParent, "position" ) === "static" ) {
				offsetParent = offsetParent.offsetParent;
			}

			return offsetParent || documentElement;
		} );
	}
} );

// Create scrollLeft and scrollTop methods
jQuery.each( { scrollLeft: "pageXOffset", scrollTop: "pageYOffset" }, function( method, prop ) {
	var top = "pageYOffset" === prop;

	jQuery.fn[ method ] = function( val ) {
		return access( this, function( elem, method, val ) {

			// Coalesce documents and windows
			var win;
			if ( isWindow( elem ) ) {
				win = elem;
			} else if ( elem.nodeType === 9 ) {
				win = elem.defaultView;
			}

			if ( val === undefined ) {
				return win ? win[ prop ] : elem[ method ];
			}

			if ( win ) {
				win.scrollTo(
					!top ? val : win.pageXOffset,
					top ? val : win.pageYOffset
				);

			} else {
				elem[ method ] = val;
			}
		}, method, val, arguments.length );
	};
} );

// Support: Safari <=7 - 9.1, Chrome <=37 - 49
// Add the top/left cssHooks using jQuery.fn.position
// Webkit bug: https://bugs.webkit.org/show_bug.cgi?id=29084
// Blink bug: https://bugs.chromium.org/p/chromium/issues/detail?id=589347
// getComputedStyle returns percent when specified for top/left/bottom/right;
// rather than make the css module depend on the offset module, just check for it here
jQuery.each( [ "top", "left" ], function( _i, prop ) {
	jQuery.cssHooks[ prop ] = addGetHookIf( support.pixelPosition,
		function( elem, computed ) {
			if ( computed ) {
				computed = curCSS( elem, prop );

				// If curCSS returns percentage, fallback to offset
				return rnumnonpx.test( computed ) ?
					jQuery( elem ).position()[ prop ] + "px" :
					computed;
			}
		}
	);
} );


// Create innerHeight, innerWidth, height, width, outerHeight and outerWidth methods
jQuery.each( { Height: "height", Width: "width" }, function( name, type ) {
	jQuery.each( {
		padding: "inner" + name,
		content: type,
		"": "outer" + name
	}, function( defaultExtra, funcName ) {

		// Margin is only for outerHeight, outerWidth
		jQuery.fn[ funcName ] = function( margin, value ) {
			var chainable = arguments.length && ( defaultExtra || typeof margin !== "boolean" ),
				extra = defaultExtra || ( margin === true || value === true ? "margin" : "border" );

			return access( this, function( elem, type, value ) {
				var doc;

				if ( isWindow( elem ) ) {

					// $( window ).outerWidth/Height return w/h including scrollbars (gh-1729)
					return funcName.indexOf( "outer" ) === 0 ?
						elem[ "inner" + name ] :
						elem.document.documentElement[ "client" + name ];
				}

				// Get document width or height
				if ( elem.nodeType === 9 ) {
					doc = elem.documentElement;

					// Either scroll[Width/Height] or offset[Width/Height] or client[Width/Height],
					// whichever is greatest
					return Math.max(
						elem.body[ "scroll" + name ], doc[ "scroll" + name ],
						elem.body[ "offset" + name ], doc[ "offset" + name ],
						doc[ "client" + name ]
					);
				}

				return value === undefined ?

					// Get width or height on the element, requesting but not forcing parseFloat
					jQuery.css( elem, type, extra ) :

					// Set width or height on the element
					jQuery.style( elem, type, value, extra );
			}, type, chainable ? margin : undefined, chainable );
		};
	} );
} );


jQuery.each( [
	"ajaxStart",
	"ajaxStop",
	"ajaxComplete",
	"ajaxError",
	"ajaxSuccess",
	"ajaxSend"
], function( _i, type ) {
	jQuery.fn[ type ] = function( fn ) {
		return this.on( type, fn );
	};
} );




jQuery.fn.extend( {

	bind: function( types, data, fn ) {
		return this.on( types, null, data, fn );
	},
	unbind: function( types, fn ) {
		return this.off( types, null, fn );
	},

	delegate: function( selector, types, data, fn ) {
		return this.on( types, selector, data, fn );
	},
	undelegate: function( selector, types, fn ) {

		// ( namespace ) or ( selector, types [, fn] )
		return arguments.length === 1 ?
			this.off( selector, "**" ) :
			this.off( types, selector || "**", fn );
	},

	hover: function( fnOver, fnOut ) {
		return this.mouseenter( fnOver ).mouseleave( fnOut || fnOver );
	}
} );

jQuery.each(
	( "blur focus focusin focusout resize scroll click dblclick " +
	"mousedown mouseup mousemove mouseover mouseout mouseenter mouseleave " +
	"change select submit keydown keypress keyup contextmenu" ).split( " " ),
	function( _i, name ) {

		// Handle event binding
		jQuery.fn[ name ] = function( data, fn ) {
			return arguments.length > 0 ?
				this.on( name, null, data, fn ) :
				this.trigger( name );
		};
	}
);




// Support: Android <=4.0 only
// Make sure we trim BOM and NBSP
var rtrim = /^[\s\uFEFF\xA0]+|[\s\uFEFF\xA0]+$/g;

// Bind a function to a context, optionally partially applying any
// arguments.
// jQuery.proxy is deprecated to promote standards (specifically Function#bind)
// However, it is not slated for removal any time soon
jQuery.proxy = function( fn, context ) {
	var tmp, args, proxy;

	if ( typeof context === "string" ) {
		tmp = fn[ context ];
		context = fn;
		fn = tmp;
	}

	// Quick check to determine if target is callable, in the spec
	// this throws a TypeError, but we will just return undefined.
	if ( !isFunction( fn ) ) {
		return undefined;
	}

	// Simulated bind
	args = slice.call( arguments, 2 );
	proxy = function() {
		return fn.apply( context || this, args.concat( slice.call( arguments ) ) );
	};

	// Set the guid of unique handler to the same of original handler, so it can be removed
	proxy.guid = fn.guid = fn.guid || jQuery.guid++;

	return proxy;
};

jQuery.holdReady = function( hold ) {
	if ( hold ) {
		jQuery.readyWait++;
	} else {
		jQuery.ready( true );
	}
};
jQuery.isArray = Array.isArray;
jQuery.parseJSON = JSON.parse;
jQuery.nodeName = nodeName;
jQuery.isFunction = isFunction;
jQuery.isWindow = isWindow;
jQuery.camelCase = camelCase;
jQuery.type = toType;

jQuery.now = Date.now;

jQuery.isNumeric = function( obj ) {

	// As of jQuery 3.0, isNumeric is limited to
	// strings and numbers (primitives or objects)
	// that can be coerced to finite numbers (gh-2662)
	var type = jQuery.type( obj );
	return ( type === "number" || type === "string" ) &&

		// parseFloat NaNs numeric-cast false positives ("")
		// ...but misinterprets leading-number strings, particularly hex literals ("0x...")
		// subtraction forces infinities to NaN
		!isNaN( obj - parseFloat( obj ) );
};

jQuery.trim = function( text ) {
	return text == null ?
		"" :
		( text + "" ).replace( rtrim, "" );
};



// Register as a named AMD module, since jQuery can be concatenated with other
// files that may use define, but not via a proper concatenation script that
// understands anonymous AMD modules. A named AMD is safest and most robust
// way to register. Lowercase jquery is used because AMD module names are
// derived from file names, and jQuery is normally delivered in a lowercase
// file name. Do this after creating the global so that if an AMD module wants
// to call noConflict to hide this version of jQuery, it will work.

// Note that for maximum portability, libraries that are not jQuery should
// declare themselves as anonymous modules, and avoid setting a global if an
// AMD loader is present. jQuery is a special case. For more information, see
// https://github.com/jrburke/requirejs/wiki/Updating-existing-libraries#wiki-anon

if ( typeof define === "function" && define.amd ) {
	define( "jquery", [], function() {
		return jQuery;
	} );
}




var

	// Map over jQuery in case of overwrite
	_jQuery = window.jQuery,

	// Map over the $ in case of overwrite
	_$ = window.$;

jQuery.noConflict = function( deep ) {
	if ( window.$ === jQuery ) {
		window.$ = _$;
	}

	if ( deep && window.jQuery === jQuery ) {
		window.jQuery = _jQuery;
	}

	return jQuery;
};

// Expose jQuery and $ identifiers, even in AMD
// (#7102#comment:10, https://github.com/jquery/jquery/pull/557)
// and CommonJS for browser emulators (#13566)
if ( typeof noGlobal === "undefined" ) {
	window.jQuery = window.$ = jQuery;
}




return jQuery;
} );
/*!
* TableSorter 2.17.8 min - Client-side table sorting with ease!
* Copyright (c) 2007 Christian Bach
*/
!function(h){h.extend({tablesorter:new function(){function d(){var b=arguments[0],a=1<arguments.length?Array.prototype.slice.call(arguments):b;if("undefined"!==typeof console&&"undefined"!==typeof console.log)console[/error/i.test(b)?"error":/warn/i.test(b)?"warn":"log"](a);else alert(a)}function q(b,a){d(b+" ("+((new Date).getTime()-a.getTime())+"ms)")}function p(b){for(var a in b)return!1;return!0}function r(b,a,c){if(!a)return"";var f,e=b.config,l=e.textExtraction||"",d="",d="basic"===l?h(a).attr(e.textAttribute)|| a.textContent||a.innerText||h(a).text()||"":"function"===typeof l?l(a,b,c):"function"===typeof(f=g.getColumnData(b,l,c))?f(a,b,c):a.textContent||a.innerText||h(a).text()||"";return h.trim(d)}function v(b){var a,c,f=b.config,e=f.$tbodies=f.$table.children("tbody:not(."+f.cssInfoBlock+")"),l,x,k,h,m,B,u,s,t,p=0,v="",w=e.length;if(0===w)return f.debug?d("Warning: *Empty table!* Not building a parser cache"):"";f.debug&&(t=new Date,d("Detecting parsers for each column"));a=[];for(c=[];p<w;){l=e[p].rows; if(l[p])for(x=f.columns,k=0;k<x;k++){h=f.$headers.filter('[data-column="'+k+'"]:last');m=g.getColumnData(b,f.headers,k);s=g.getParserById(g.getData(h,m,"extractor"));u=g.getParserById(g.getData(h,m,"sorter"));B="false"===g.getData(h,m,"parser");f.empties[k]=(g.getData(h,m,"empty")||f.emptyTo||(f.emptyToBottom?"bottom":"top")).toLowerCase();f.strings[k]=(g.getData(h,m,"string")||f.stringTo||"max").toLowerCase();B&&(u=g.getParserById("no-parser"));s||(s=!1);if(!u)a:{h=b;m=l;B=-1;u=k;for(var A=void 0, K=g.parsers.length,G=!1,z="",A=!0;""===z&&A;)B++,m[B]?(G=m[B].cells[u],z=r(h,G,u),h.config.debug&&d("Checking if value was empty on row "+B+", column: "+u+': "'+z+'"')):A=!1;for(;0<=--K;)if((A=g.parsers[K])&&"text"!==A.id&&A.is&&A.is(z,h,G)){u=A;break a}u=g.getParserById("text")}f.debug&&(v+="column:"+k+"; extractor:"+s.id+"; parser:"+u.id+"; string:"+f.strings[k]+"; empty: "+f.empties[k]+"\n");c[k]=u;a[k]=s}p+=c.length?w:1}f.debug&&(d(v?v:"No parsers detected"),q("Completed detecting parsers",t)); f.parsers=c;f.extractors=a}function w(b){var a,c,f,e,l,x,k,n,m,p,u,s=b.config,t=s.$table.children("tbody"),v=s.extractors,w=s.parsers;s.cache={};s.totalRows=0;if(!w)return s.debug?d("Warning: *Empty table!* Not building a cache"):"";s.debug&&(n=new Date);s.showProcessing&&g.isProcessing(b,!0);for(l=0;l<t.length;l++)if(u=[],a=s.cache[l]={normalized:[]},!t.eq(l).hasClass(s.cssInfoBlock)){m=t[l]&&t[l].rows.length||0;for(f=0;f<m;++f)if(p={child:[]},x=h(t[l].rows[f]),k=[],x.hasClass(s.cssChildRow)&&0!== f)c=a.normalized.length-1,a.normalized[c][s.columns].$row=a.normalized[c][s.columns].$row.add(x),x.prev().hasClass(s.cssChildRow)||x.prev().addClass(g.css.cssHasChild),p.child[c]=h.trim(x[0].textContent||x[0].innerText||x.text()||"");else{p.$row=x;p.order=f;for(e=0;e<s.columns;++e)"undefined"===typeof w[e]?s.debug&&d("No parser found for cell:",x[0].cells[e],"does it have a header?"):(c=r(b,x[0].cells[e],e),c="undefined"===typeof v[e].id?c:v[e].format(c,b,x[0].cells[e],e),c="no-parser"===w[e].id? "":w[e].format(c,b,x[0].cells[e],e),k.push(s.ignoreCase&&"string"===typeof c?c.toLowerCase():c),"numeric"===(w[e].type||"").toLowerCase()&&(u[e]=Math.max(Math.abs(c)||0,u[e]||0)));k[s.columns]=p;a.normalized.push(k)}a.colMax=u;s.totalRows+=a.normalized.length}s.showProcessing&&g.isProcessing(b);s.debug&&q("Building cache for "+m+" rows",n)}function z(b,a){var c=b.config,f=c.widgetOptions,e=b.tBodies,l=[],d=c.cache,k,n,m,r,u,s;if(p(d))return c.appender?c.appender(b,l):b.isUpdating?c.$table.trigger("updateComplete", b):"";c.debug&&(s=new Date);for(u=0;u<e.length;u++)if(k=h(e[u]),k.length&&!k.hasClass(c.cssInfoBlock)){m=g.processTbody(b,k,!0);k=d[u].normalized;n=k.length;for(r=0;r<n;r++)l.push(k[r][c.columns].$row),c.appender&&(!c.pager||c.pager.removeRows&&f.pager_removeRows||c.pager.ajax)||m.append(k[r][c.columns].$row);g.processTbody(b,m,!1)}c.appender&&c.appender(b,l);c.debug&&q("Rebuilt table",s);a||c.appender||g.applyWidget(b);b.isUpdating&&c.$table.trigger("updateComplete",b)}function D(b){return/^d/i.test(b)|| 1===b}function E(b){var a,c,f,e,l,x,k,n=b.config;n.headerList=[];n.headerContent=[];n.debug&&(k=new Date);n.columns=g.computeColumnIndex(n.$table.children("thead, tfoot").children("tr"));e=n.cssIcon?'<i class="'+(n.cssIcon===g.css.icon?g.css.icon:n.cssIcon+" "+g.css.icon)+'"></i>':"";n.$headers=h(b).find(n.selectorHeaders).each(function(k){c=h(this);a=g.getColumnData(b,n.headers,k,!0);n.headerContent[k]=h(this).html();""!==n.headerTemplate&&(l=n.headerTemplate.replace(/\{content\}/g,h(this).html()).replace(/\{icon\}/g, e),n.onRenderTemplate&&(f=n.onRenderTemplate.apply(c,[k,l]))&&"string"===typeof f&&(l=f),h(this).html('<div class="'+g.css.headerIn+'">'+l+"</div>"));n.onRenderHeader&&n.onRenderHeader.apply(c,[k]);this.column=parseInt(h(this).attr("data-column"),10);this.order=D(g.getData(c,a,"sortInitialOrder")||n.sortInitialOrder)?[1,0,2]:[0,1,2];this.count=-1;this.lockedOrder=!1;x=g.getData(c,a,"lockedOrder")||!1;"undefined"!==typeof x&&!1!==x&&(this.order=this.lockedOrder=D(x)?[1,1,1]:[0,0,0]);c.addClass(g.css.header+ " "+n.cssHeader);n.headerList[k]=this;c.parent().addClass(g.css.headerRow+" "+n.cssHeaderRow).attr("role","row");n.tabIndex&&c.attr("tabindex",0)}).attr({scope:"col",role:"columnheader"});H(b);n.debug&&(q("Built headers:",k),d(n.$headers))}function C(b,a,c){var f=b.config;f.$table.find(f.selectorRemove).remove();v(b);w(b);I(f.$table,a,c)}function H(b){var a,c,f,e=b.config;e.$headers.each(function(l,d){c=h(d);f=g.getColumnData(b,e.headers,l,!0);a="false"===g.getData(d,f,"sorter")||"false"===g.getData(d, f,"parser");d.sortDisabled=a;c[a?"addClass":"removeClass"]("sorter-false").attr("aria-disabled",""+a);b.id&&(a?c.removeAttr("aria-controls"):c.attr("aria-controls",b.id))})}function F(b){var a,c,f=b.config,e=f.sortList,l=e.length,d=g.css.sortNone+" "+f.cssNone,k=[g.css.sortAsc+" "+f.cssAsc,g.css.sortDesc+" "+f.cssDesc],n=["ascending","descending"],m=h(b).find("tfoot tr").children().add(f.$extraHeaders).removeClass(k.join(" "));f.$headers.removeClass(k.join(" ")).addClass(d).attr("aria-sort","none"); for(a=0;a<l;a++)if(2!==e[a][1]&&(b=f.$headers.not(".sorter-false").filter('[data-column="'+e[a][0]+'"]'+(1===l?":last":"")),b.length)){for(c=0;c<b.length;c++)b[c].sortDisabled||b.eq(c).removeClass(d).addClass(k[e[a][1]]).attr("aria-sort",n[e[a][1]]);m.length&&m.filter('[data-column="'+e[a][0]+'"]').removeClass(d).addClass(k[e[a][1]])}f.$headers.not(".sorter-false").each(function(){var b=h(this),a=this.order[(this.count+1)%(f.sortReset?3:2)],a=b.text()+": "+g.language[b.hasClass(g.css.sortAsc)?"sortAsc": b.hasClass(g.css.sortDesc)?"sortDesc":"sortNone"]+g.language[0===a?"nextAsc":1===a?"nextDesc":"nextNone"];b.attr("aria-label",a)})}function O(b){var a,c,f=b.config;f.widthFixed&&0===f.$table.find("colgroup").length&&(a=h("<colgroup>"),c=h(b).width(),h(b.tBodies).not("."+f.cssInfoBlock).find("tr:first").children(":visible").each(function(){a.append(h("<col>").css("width",parseInt(h(this).width()/c*1E3,10)/10+"%"))}),f.$table.prepend(a))}function P(b,a){var c,f,e,l,g,k=b.config,d=a||k.sortList;k.sortList= [];h.each(d,function(b,a){l=parseInt(a[0],10);if(e=k.$headers.filter('[data-column="'+l+'"]:last')[0]){f=(f=(""+a[1]).match(/^(1|d|s|o|n)/))?f[0]:"";switch(f){case "1":case "d":f=1;break;case "s":f=g||0;break;case "o":c=e.order[(g||0)%(k.sortReset?3:2)];f=0===c?1:1===c?0:2;break;case "n":e.count+=1;f=e.order[e.count%(k.sortReset?3:2)];break;default:f=0}g=0===b?f:g;c=[l,parseInt(f,10)||0];k.sortList.push(c);f=h.inArray(c[1],e.order);e.count=0<=f?f:c[1]%(k.sortReset?3:2)}})}function Q(b,a){return b&& b[a]?b[a].type||"":""}function L(b,a,c){if(b.isUpdating)return setTimeout(function(){L(b,a,c)},50);var f,e,l,d,k=b.config,n=!c[k.sortMultiSortKey],m=k.$table;m.trigger("sortStart",b);a.count=c[k.sortResetKey]?2:(a.count+1)%(k.sortReset?3:2);k.sortRestart&&(e=a,k.$headers.each(function(){this===e||!n&&h(this).is("."+g.css.sortDesc+",."+g.css.sortAsc)||(this.count=-1)}));e=a.column;if(n){k.sortList=[];if(null!==k.sortForce)for(f=k.sortForce,l=0;l<f.length;l++)f[l][0]!==e&&k.sortList.push(f[l]);f=a.order[a.count]; if(2>f&&(k.sortList.push([e,f]),1<a.colSpan))for(l=1;l<a.colSpan;l++)k.sortList.push([e+l,f])}else{if(k.sortAppend&&1<k.sortList.length)for(l=0;l<k.sortAppend.length;l++)d=g.isValueInArray(k.sortAppend[l][0],k.sortList),0<=d&&k.sortList.splice(d,1);if(0<=g.isValueInArray(e,k.sortList))for(l=0;l<k.sortList.length;l++)d=k.sortList[l],f=k.$headers.filter('[data-column="'+d[0]+'"]:last')[0],d[0]===e&&(d[1]=f.order[a.count],2===d[1]&&(k.sortList.splice(l,1),f.count=-1));else if(f=a.order[a.count],2>f&& (k.sortList.push([e,f]),1<a.colSpan))for(l=1;l<a.colSpan;l++)k.sortList.push([e+l,f])}if(null!==k.sortAppend)for(f=k.sortAppend,l=0;l<f.length;l++)f[l][0]!==e&&k.sortList.push(f[l]);m.trigger("sortBegin",b);setTimeout(function(){F(b);J(b);z(b);m.trigger("sortEnd",b)},1)}function J(b){var a,c,f,e,l,d,k,h,m,r,u,s=0,t=b.config,v=t.textSorter||"",w=t.sortList,y=w.length,z=b.tBodies.length;if(!t.serverSideSorting&&!p(t.cache)){t.debug&&(l=new Date);for(c=0;c<z;c++)d=t.cache[c].colMax,k=t.cache[c].normalized, k.sort(function(c,l){for(a=0;a<y;a++){e=w[a][0];h=w[a][1];s=0===h;if(t.sortStable&&c[e]===l[e]&&1===y)break;(f=/n/i.test(Q(t.parsers,e)))&&t.strings[e]?(f="boolean"===typeof t.string[t.strings[e]]?(s?1:-1)*(t.string[t.strings[e]]?-1:1):t.strings[e]?t.string[t.strings[e]]||0:0,m=t.numberSorter?t.numberSorter(c[e],l[e],s,d[e],b):g["sortNumeric"+(s?"Asc":"Desc")](c[e],l[e],f,d[e],e,b)):(r=s?c:l,u=s?l:c,m="function"===typeof v?v(r[e],u[e],s,e,b):"object"===typeof v&&v.hasOwnProperty(e)?v[e](r[e],u[e], s,e,b):g["sortNatural"+(s?"Asc":"Desc")](c[e],l[e],e,b,t));if(m)return m}return c[t.columns].order-l[t.columns].order});t.debug&&q("Sorting on "+w.toString()+" and dir "+h+" time",l)}}function M(b,a){var c=b[0];c.isUpdating&&b.trigger("updateComplete",c);h.isFunction(a)&&a(b[0])}function I(b,a,c){var f=b[0].config.sortList;!1!==a&&!b[0].isProcessing&&f.length?b.trigger("sorton",[f,function(){M(b,c)},!0]):(M(b,c),g.applyWidget(b[0],!1))}function N(b){var a=b.config,c=a.$table;c.unbind("sortReset update updateRows updateCell updateAll addRows updateComplete sorton appendCache updateCache applyWidgetId applyWidgets refreshWidgets destroy mouseup mouseleave ".split(" ").join(a.namespace+ " ")).bind("sortReset"+a.namespace,function(f,e){f.stopPropagation();a.sortList=[];F(b);J(b);z(b);h.isFunction(e)&&e(b)}).bind("updateAll"+a.namespace,function(f,e,c){f.stopPropagation();b.isUpdating=!0;g.refreshWidgets(b,!0,!0);g.restoreHeaders(b);E(b);g.bindEvents(b,a.$headers,!0);N(b);C(b,e,c)}).bind("update"+a.namespace+" updateRows"+a.namespace,function(a,e,c){a.stopPropagation();b.isUpdating=!0;H(b);C(b,e,c)}).bind("updateCell"+a.namespace,function(f,e,l,g){f.stopPropagation();b.isUpdating= !0;c.find(a.selectorRemove).remove();var d,n,m;n=c.find("tbody");m=h(e);f=n.index(h.fn.closest?m.closest("tbody"):m.parents("tbody").filter(":first"));d=h.fn.closest?m.closest("tr"):m.parents("tr").filter(":first");e=m[0];n.length&&0<=f&&(n=n.eq(f).find("tr").index(d),m=m.index(),a.cache[f].normalized[n][a.columns].$row=d,d="undefined"===typeof a.extractors[m].id?r(b,e,m):a.extractors[m].format(r(b,e,m),b,e,m),e="no-parser"===a.parsers[m].id?"":a.parsers[m].format(d,b,e,m),a.cache[f].normalized[n][m]= a.ignoreCase&&"string"===typeof e?e.toLowerCase():e,"numeric"===(a.parsers[m].type||"").toLowerCase()&&(a.cache[f].colMax[m]=Math.max(Math.abs(e)||0,a.cache[f].colMax[m]||0)),I(c,l,g))}).bind("addRows"+a.namespace,function(f,e,l,g){f.stopPropagation();b.isUpdating=!0;if(p(a.cache))H(b),C(b,l,g);else{e=h(e).attr("role","row");var d,n,m,q,u,s=e.filter("tr").length,t=c.find("tbody").index(e.parents("tbody").filter(":first"));a.parsers&&a.parsers.length||v(b);for(f=0;f<s;f++){n=e[f].cells.length;u=[]; q={child:[],$row:e.eq(f),order:a.cache[t].normalized.length};for(d=0;d<n;d++)m="undefined"===typeof a.extractors[d].id?r(b,e[f].cells[d],d):a.extractors[d].format(r(b,e[f].cells[d],d),b,e[f].cells[d],d),m="no-parser"===a.parsers[d].id?"":a.parsers[d].format(m,b,e[f].cells[d],d),u[d]=a.ignoreCase&&"string"===typeof m?m.toLowerCase():m,"numeric"===(a.parsers[d].type||"").toLowerCase()&&(a.cache[t].colMax[d]=Math.max(Math.abs(u[d])||0,a.cache[t].colMax[d]||0));u.push(q);a.cache[t].normalized.push(u)}I(c, l,g)}}).bind("updateComplete"+a.namespace,function(){b.isUpdating=!1}).bind("sorton"+a.namespace,function(a,e,d,x){var k=b.config;a.stopPropagation();c.trigger("sortStart",this);P(b,e);F(b);k.delayInit&&p(k.cache)&&w(b);c.trigger("sortBegin",this);J(b);z(b,x);c.trigger("sortEnd",this);g.applyWidget(b);h.isFunction(d)&&d(b)}).bind("appendCache"+a.namespace,function(a,e,c){a.stopPropagation();z(b,c);h.isFunction(e)&&e(b)}).bind("updateCache"+a.namespace,function(c,e){a.parsers&&a.parsers.length||v(b); w(b);h.isFunction(e)&&e(b)}).bind("applyWidgetId"+a.namespace,function(c,e){c.stopPropagation();g.getWidgetById(e).format(b,a,a.widgetOptions)}).bind("applyWidgets"+a.namespace,function(a,c){a.stopPropagation();g.applyWidget(b,c)}).bind("refreshWidgets"+a.namespace,function(a,c,d){a.stopPropagation();g.refreshWidgets(b,c,d)}).bind("destroy"+a.namespace,function(a,c,d){a.stopPropagation();g.destroy(b,c,d)}).bind("resetToLoadState"+a.namespace,function(){g.refreshWidgets(b,!0,!0);a=h.extend(!0,g.defaults, a.originalSettings);b.hasInitialized=!1;g.setup(b,a)})}var g=this;g.version="2.17.8";g.parsers=[];g.widgets=[];g.defaults={theme:"default",widthFixed:!1,showProcessing:!1,headerTemplate:"{content}",onRenderTemplate:null,onRenderHeader:null,cancelSelection:!0,tabIndex:!0,dateFormat:"mmddyyyy",sortMultiSortKey:"shiftKey",sortResetKey:"ctrlKey",usNumberFormat:!0,delayInit:!1,serverSideSorting:!1,headers:{},ignoreCase:!0,sortForce:null,sortList:[],sortAppend:null,sortStable:!1,sortInitialOrder:"asc", sortLocaleCompare:!1,sortReset:!1,sortRestart:!1,emptyTo:"bottom",stringTo:"max",textExtraction:"basic",textAttribute:"data-text",textSorter:null,numberSorter:null,widgets:[],widgetOptions:{zebra:["even","odd"]},initWidgets:!0,initialized:null,tableClass:"",cssAsc:"",cssDesc:"",cssNone:"",cssHeader:"",cssHeaderRow:"",cssProcessing:"",cssChildRow:"tablesorter-childRow",cssIcon:"tablesorter-icon",cssInfoBlock:"tablesorter-infoOnly",selectorHeaders:"> thead th, > thead td",selectorSort:"th, td",selectorRemove:".remove-me", debug:!1,headerList:[],empties:{},strings:{},parsers:[]};g.css={table:"tablesorter",cssHasChild:"tablesorter-hasChildRow",childRow:"tablesorter-childRow",header:"tablesorter-header",headerRow:"tablesorter-headerRow",headerIn:"tablesorter-header-inner",icon:"tablesorter-icon",info:"tablesorter-infoOnly",processing:"tablesorter-processing",sortAsc:"tablesorter-headerAsc",sortDesc:"tablesorter-headerDesc",sortNone:"tablesorter-headerUnSorted"};g.language={sortAsc:"Ascending sort applied, ",sortDesc:"Descending sort applied, ", sortNone:"No sort applied, ",nextAsc:"activate to apply an ascending sort",nextDesc:"activate to apply a descending sort",nextNone:"activate to remove the sort"};g.log=d;g.benchmark=q;g.construct=function(b){return this.each(function(){var a=h.extend(!0,{},g.defaults,b);a.originalSettings=b;!this.hasInitialized&&g.buildTable&&"TABLE"!==this.tagName?g.buildTable(this,a):g.setup(this,a)})};g.setup=function(b,a){if(!b||!b.tHead||0===b.tBodies.length||!0===b.hasInitialized)return a.debug?d("ERROR: stopping initialization! No table, thead, tbody or tablesorter has already been initialized"): "";var c="",f=h(b),e=h.metadata;b.hasInitialized=!1;b.isProcessing=!0;b.config=a;h.data(b,"tablesorter",a);a.debug&&h.data(b,"startoveralltimer",new Date);a.supportsDataObject=function(a){a[0]=parseInt(a[0],10);return 1<a[0]||1===a[0]&&4<=parseInt(a[1],10)}(h.fn.jquery.split("."));a.string={max:1,min:-1,emptymin:1,emptymax:-1,zero:0,none:0,"null":0,top:!0,bottom:!1};a.emptyTo=a.emptyTo.toLowerCase();a.stringTo=a.stringTo.toLowerCase();/tablesorter\-/.test(f.attr("class"))||(c=""!==a.theme?" tablesorter-"+ a.theme:"");a.table=b;a.$table=f.addClass(g.css.table+" "+a.tableClass+c).attr("role","grid");a.$headers=f.find(a.selectorHeaders);a.namespace=a.namespace?"."+a.namespace.replace(/\W/g,""):".tablesorter"+Math.random().toString(16).slice(2);a.$table.children().children("tr").attr("role","row");a.$tbodies=f.children("tbody:not(."+a.cssInfoBlock+")").attr({"aria-live":"polite","aria-relevant":"all"});a.$table.find("caption").length&&a.$table.attr("aria-labelledby","theCaption");a.widgetInit={};a.textExtraction= a.$table.attr("data-text-extraction")||a.textExtraction||"basic";E(b);O(b);v(b);a.totalRows=0;a.delayInit||w(b);g.bindEvents(b,a.$headers,!0);N(b);a.supportsDataObject&&"undefined"!==typeof f.data().sortlist?a.sortList=f.data().sortlist:e&&f.metadata()&&f.metadata().sortlist&&(a.sortList=f.metadata().sortlist);g.applyWidget(b,!0);0<a.sortList.length?f.trigger("sorton",[a.sortList,{},!a.initWidgets,!0]):(F(b),a.initWidgets&&g.applyWidget(b,!1));a.showProcessing&&f.unbind("sortBegin"+a.namespace+" sortEnd"+ a.namespace).bind("sortBegin"+a.namespace+" sortEnd"+a.namespace,function(c){clearTimeout(a.processTimer);g.isProcessing(b);"sortBegin"===c.type&&(a.processTimer=setTimeout(function(){g.isProcessing(b,!0)},500))});b.hasInitialized=!0;b.isProcessing=!1;a.debug&&g.benchmark("Overall initialization time",h.data(b,"startoveralltimer"));f.trigger("tablesorter-initialized",b);"function"===typeof a.initialized&&a.initialized(b)};g.getColumnData=function(b,a,c,f){if("undefined"!==typeof a&&null!==a){b=h(b)[0]; var e,d=b.config;if(a[c])return f?a[c]:a[d.$headers.index(d.$headers.filter('[data-column="'+c+'"]:last'))];for(e in a)if("string"===typeof e&&(b=f?d.$headers.eq(c).filter(e):d.$headers.filter('[data-column="'+c+'"]:last').filter(e),b.length))return a[e]}};g.computeColumnIndex=function(b){var a=[],c=0,f,e,d,g,k,n,m,p,q,s;for(f=0;f<b.length;f++)for(k=b[f].cells,e=0;e<k.length;e++){d=k[e];g=h(d);n=d.parentNode.rowIndex;g.index();m=d.rowSpan||1;p=d.colSpan||1;"undefined"===typeof a[n]&&(a[n]=[]);for(d= 0;d<a[n].length+1;d++)if("undefined"===typeof a[n][d]){q=d;break}c=Math.max(q,c);g.attr({"data-column":q});for(d=n;d<n+m;d++)for("undefined"===typeof a[d]&&(a[d]=[]),s=a[d],g=q;g<q+p;g++)s[g]="x"}return c+1};g.isProcessing=function(b,a,c){b=h(b);var f=b[0].config,e=c||b.find("."+g.css.header);a?("undefined"!==typeof c&&0<f.sortList.length&&(e=e.filter(function(){return this.sortDisabled?!1:0<=g.isValueInArray(parseFloat(h(this).attr("data-column")),f.sortList)})),b.add(e).addClass(g.css.processing+ " "+f.cssProcessing)):b.add(e).removeClass(g.css.processing+" "+f.cssProcessing)};g.processTbody=function(b,a,c){b=h(b)[0];if(c)return b.isProcessing=!0,a.before('<span class="tablesorter-savemyplace"/>'),c=h.fn.detach?a.detach():a.remove();c=h(b).find("span.tablesorter-savemyplace");a.insertAfter(c);c.remove();b.isProcessing=!1};g.clearTableBody=function(b){h(b)[0].config.$tbodies.children().detach()};g.bindEvents=function(b,a,c){b=h(b)[0];var f,e=b.config;!0!==c&&(e.$extraHeaders=e.$extraHeaders? e.$extraHeaders.add(a):a);a.find(e.selectorSort).add(a.filter(e.selectorSort)).unbind(["mousedown","mouseup","sort","keyup",""].join(e.namespace+" ")).bind(["mousedown","mouseup","sort","keyup",""].join(e.namespace+" "),function(c,d){var g;g=c.type;if(!(1!==(c.which||c.button)&&!/sort|keyup/.test(g)||"keyup"===g&&13!==c.which||"mouseup"===g&&!0!==d&&250<(new Date).getTime()-f)){if("mousedown"===g)return f=(new Date).getTime(),/(input|select|button|textarea)/i.test(c.target.tagName)?"":!e.cancelSelection; e.delayInit&&p(e.cache)&&w(b);g=h.fn.closest?h(this).closest("th, td")[0]:/TH|TD/.test(this.tagName)?this:h(this).parents("th, td")[0];g=e.$headers[a.index(g)];g.sortDisabled||L(b,g,c)}});e.cancelSelection&&a.attr("unselectable","on").bind("selectstart",!1).css({"user-select":"none",MozUserSelect:"none"})};g.restoreHeaders=function(b){var a=h(b)[0].config;a.$table.find(a.selectorHeaders).each(function(b){h(this).find("."+g.css.headerIn).length&&h(this).html(a.headerContent[b])})};g.destroy=function(b, a,c){b=h(b)[0];if(b.hasInitialized){g.refreshWidgets(b,!0,!0);var f=h(b),e=b.config,d=f.find("thead:first"),q=d.find("tr."+g.css.headerRow).removeClass(g.css.headerRow+" "+e.cssHeaderRow),k=f.find("tfoot:first > tr").children("th, td");!1===a&&0<=h.inArray("uitheme",e.widgets)&&(f.trigger("applyWidgetId",["uitheme"]),f.trigger("applyWidgetId",["zebra"]));d.find("tr").not(q).remove();f.removeData("tablesorter").unbind("sortReset update updateAll updateRows updateCell addRows updateComplete sorton appendCache updateCache applyWidgetId applyWidgets refreshWidgets destroy mouseup mouseleave keypress sortBegin sortEnd resetToLoadState ".split(" ").join(e.namespace+ " "));e.$headers.add(k).removeClass([g.css.header,e.cssHeader,e.cssAsc,e.cssDesc,g.css.sortAsc,g.css.sortDesc,g.css.sortNone].join(" ")).removeAttr("data-column").removeAttr("aria-label").attr("aria-disabled","true");q.find(e.selectorSort).unbind(["mousedown","mouseup","keypress",""].join(e.namespace+" "));g.restoreHeaders(b);f.toggleClass(g.css.table+" "+e.tableClass+" tablesorter-"+e.theme,!1===a);b.hasInitialized=!1;delete b.config.cache;"function"===typeof c&&c(b)}};g.regex={chunk:/(^([+\-]?(?:0|[1-9]\d*)(?:\.\d*)?(?:[eE][+\-]?\d+)?)?$|^0x[0-9a-f]+$|\d+)/gi, chunks:/(^\\0|\\0$)/,hex:/^0x[0-9a-f]+$/i};g.sortNatural=function(b,a){if(b===a)return 0;var c,f,e,d,h,k;f=g.regex;if(f.hex.test(a)){c=parseInt(b.match(f.hex),16);e=parseInt(a.match(f.hex),16);if(c<e)return-1;if(c>e)return 1}c=b.replace(f.chunk,"\\0$1\\0").replace(f.chunks,"").split("\\0");f=a.replace(f.chunk,"\\0$1\\0").replace(f.chunks,"").split("\\0");k=Math.max(c.length,f.length);for(h=0;h<k;h++){e=isNaN(c[h])?c[h]||0:parseFloat(c[h])||0;d=isNaN(f[h])?f[h]||0:parseFloat(f[h])||0;if(isNaN(e)!== isNaN(d))return isNaN(e)?1:-1;typeof e!==typeof d&&(e+="",d+="");if(e<d)return-1;if(e>d)return 1}return 0};g.sortNaturalAsc=function(b,a,c,f,e){if(b===a)return 0;c=e.string[e.empties[c]||e.emptyTo];return""===b&&0!==c?"boolean"===typeof c?c?-1:1:-c||-1:""===a&&0!==c?"boolean"===typeof c?c?1:-1:c||1:g.sortNatural(b,a)};g.sortNaturalDesc=function(b,a,c,f,e){if(b===a)return 0;c=e.string[e.empties[c]||e.emptyTo];return""===b&&0!==c?"boolean"===typeof c?c?-1:1:c||1:""===a&&0!==c?"boolean"===typeof c?c? 1:-1:-c||-1:g.sortNatural(a,b)};g.sortText=function(b,a){return b>a?1:b<a?-1:0};g.getTextValue=function(b,a,c){if(c){var f=b?b.length:0,e=c+a;for(c=0;c<f;c++)e+=b.charCodeAt(c);return a*e}return 0};g.sortNumericAsc=function(b,a,c,f,e,d){if(b===a)return 0;d=d.config;e=d.string[d.empties[e]||d.emptyTo];if(""===b&&0!==e)return"boolean"===typeof e?e?-1:1:-e||-1;if(""===a&&0!==e)return"boolean"===typeof e?e?1:-1:e||1;isNaN(b)&&(b=g.getTextValue(b,c,f));isNaN(a)&&(a=g.getTextValue(a,c,f));return b-a};g.sortNumericDesc= function(b,a,c,f,e,d){if(b===a)return 0;d=d.config;e=d.string[d.empties[e]||d.emptyTo];if(""===b&&0!==e)return"boolean"===typeof e?e?-1:1:e||1;if(""===a&&0!==e)return"boolean"===typeof e?e?1:-1:-e||-1;isNaN(b)&&(b=g.getTextValue(b,c,f));isNaN(a)&&(a=g.getTextValue(a,c,f));return a-b};g.sortNumeric=function(b,a){return b-a};g.characterEquivalents={a:"\u00e1\u00e0\u00e2\u00e3\u00e4\u0105\u00e5",A:"\u00c1\u00c0\u00c2\u00c3\u00c4\u0104\u00c5",c:"\u00e7\u0107\u010d",C:"\u00c7\u0106\u010c",e:"\u00e9\u00e8\u00ea\u00eb\u011b\u0119", E:"\u00c9\u00c8\u00ca\u00cb\u011a\u0118",i:"\u00ed\u00ec\u0130\u00ee\u00ef\u0131",I:"\u00cd\u00cc\u0130\u00ce\u00cf",o:"\u00f3\u00f2\u00f4\u00f5\u00f6",O:"\u00d3\u00d2\u00d4\u00d5\u00d6",ss:"\u00df",SS:"\u1e9e",u:"\u00fa\u00f9\u00fb\u00fc\u016f",U:"\u00da\u00d9\u00db\u00dc\u016e"};g.replaceAccents=function(b){var a,c="[",d=g.characterEquivalents;if(!g.characterRegex){g.characterRegexArray={};for(a in d)"string"===typeof a&&(c+=d[a],g.characterRegexArray[a]=new RegExp("["+d[a]+"]","g"));g.characterRegex= new RegExp(c+"]")}if(g.characterRegex.test(b))for(a in d)"string"===typeof a&&(b=b.replace(g.characterRegexArray[a],a));return b};g.isValueInArray=function(b,a){var c,d=a.length;for(c=0;c<d;c++)if(a[c][0]===b)return c;return-1};g.addParser=function(b){var a,c=g.parsers.length,d=!0;for(a=0;a<c;a++)g.parsers[a].id.toLowerCase()===b.id.toLowerCase()&&(d=!1);d&&g.parsers.push(b)};g.getParserById=function(b){if("false"==b)return!1;var a,c=g.parsers.length;for(a=0;a<c;a++)if(g.parsers[a].id.toLowerCase()=== b.toString().toLowerCase())return g.parsers[a];return!1};g.addWidget=function(b){g.widgets.push(b)};g.hasWidget=function(b,a){b=h(b);return b.length&&b[0].config&&b[0].config.widgetInit[a]||!1};g.getWidgetById=function(b){var a,c,d=g.widgets.length;for(a=0;a<d;a++)if((c=g.widgets[a])&&c.hasOwnProperty("id")&&c.id.toLowerCase()===b.toLowerCase())return c};g.applyWidget=function(b,a){b=h(b)[0];var c=b.config,d=c.widgetOptions,e=[],l,p,k;!1!==a&&b.hasInitialized&&(b.isApplyingWidgets||b.isUpdating)|| (c.debug&&(l=new Date),c.widgets.length&&(b.isApplyingWidgets=!0,c.widgets=h.grep(c.widgets,function(a,b){return h.inArray(a,c.widgets)===b}),h.each(c.widgets||[],function(a,b){(k=g.getWidgetById(b))&&k.id&&(k.priority||(k.priority=10),e[a]=k)}),e.sort(function(a,b){return a.priority<b.priority?-1:a.priority===b.priority?0:1}),h.each(e,function(e,g){if(g){if(a||!c.widgetInit[g.id])c.widgetInit[g.id]=!0,g.hasOwnProperty("options")&&(d=b.config.widgetOptions=h.extend(!0,{},g.options,d)),g.hasOwnProperty("init")&& g.init(b,g,c,d);!a&&g.hasOwnProperty("format")&&g.format(b,c,d,!1)}})),setTimeout(function(){b.isApplyingWidgets=!1},0),c.debug&&(p=c.widgets.length,q("Completed "+(!0===a?"initializing ":"applying ")+p+" widget"+(1!==p?"s":""),l)))};g.refreshWidgets=function(b,a,c){b=h(b)[0];var f,e=b.config,l=e.widgets,q=g.widgets,k=q.length;for(f=0;f<k;f++)q[f]&&q[f].id&&(a||0>h.inArray(q[f].id,l))&&(e.debug&&d('Refeshing widgets: Removing "'+q[f].id+'"'),q[f].hasOwnProperty("remove")&&e.widgetInit[q[f].id]&&(q[f].remove(b, e,e.widgetOptions),e.widgetInit[q[f].id]=!1));!0!==c&&g.applyWidget(b,a)};g.getData=function(b,a,c){var d="";b=h(b);var e,g;if(!b.length)return"";e=h.metadata?b.metadata():!1;g=" "+(b.attr("class")||"");"undefined"!==typeof b.data(c)||"undefined"!==typeof b.data(c.toLowerCase())?d+=b.data(c)||b.data(c.toLowerCase()):e&&"undefined"!==typeof e[c]?d+=e[c]:a&&"undefined"!==typeof a[c]?d+=a[c]:" "!==g&&g.match(" "+c+"-")&&(d=g.match(new RegExp("\\s"+c+"-([\\w-]+)"))[1]||"");return h.trim(d)};g.formatFloat= function(b,a){if("string"!==typeof b||""===b)return b;var c;b=(a&&a.config?!1!==a.config.usNumberFormat:"undefined"!==typeof a?a:1)?b.replace(/,/g,""):b.replace(/[\s|\.]/g,"").replace(/,/g,".");/^\s*\([.\d]+\)/.test(b)&&(b=b.replace(/^\s*\(([.\d]+)\)/,"-$1"));c=parseFloat(b);return isNaN(c)?h.trim(b):c};g.isDigit=function(b){return isNaN(b)?/^[\-+(]?\d+[)]?$/.test(b.toString().replace(/[,.'"\s]/g,"")):!0}}});var r=h.tablesorter;h.fn.extend({tablesorter:r.construct});r.addParser({id:"no-parser",is:function(){return!1}, format:function(){return""},type:"text"});r.addParser({id:"text",is:function(){return!0},format:function(d,q){var p=q.config;d&&(d=h.trim(p.ignoreCase?d.toLocaleLowerCase():d),d=p.sortLocaleCompare?r.replaceAccents(d):d);return d},type:"text"});r.addParser({id:"digit",is:function(d){return r.isDigit(d)},format:function(d,q){var p=r.formatFloat((d||"").replace(/[^\w,. \-()]/g,""),q);return d&&"number"===typeof p?p:d?h.trim(d&&q.config.ignoreCase?d.toLocaleLowerCase():d):d},type:"numeric"});r.addParser({id:"currency", is:function(d){return/^\(?\d+[\u00a3$\u20ac\u00a4\u00a5\u00a2?.]|[\u00a3$\u20ac\u00a4\u00a5\u00a2?.]\d+\)?$/.test((d||"").replace(/[+\-,. ]/g,""))},format:function(d,q){var p=r.formatFloat((d||"").replace(/[^\w,. \-()]/g,""),q);return d&&"number"===typeof p?p:d?h.trim(d&&q.config.ignoreCase?d.toLocaleLowerCase():d):d},type:"numeric"});r.addParser({id:"ipAddress",is:function(d){return/^\d{1,3}[\.]\d{1,3}[\.]\d{1,3}[\.]\d{1,3}$/.test(d)},format:function(d,h){var p,y=d?d.split("."):"",v="",w=y.length; for(p=0;p<w;p++)v+=("00"+y[p]).slice(-3);return d?r.formatFloat(v,h):d},type:"numeric"});r.addParser({id:"url",is:function(d){return/^(https?|ftp|file):\/\//.test(d)},format:function(d){return d?h.trim(d.replace(/(https?|ftp|file):\/\//,"")):d},parsed:!0,type:"text"});r.addParser({id:"isoDate",is:function(d){return/^\d{4}[\/\-]\d{1,2}[\/\-]\d{1,2}/.test(d)},format:function(d,h){return d?r.formatFloat(""!==d?(new Date(d.replace(/-/g,"/"))).getTime()||d:"",h):d},type:"numeric"});r.addParser({id:"percent", is:function(d){return/(\d\s*?%|%\s*?\d)/.test(d)&&15>d.length},format:function(d,h){return d?r.formatFloat(d.replace(/%/g,""),h):d},type:"numeric"});r.addParser({id:"usLongDate",is:function(d){return/^[A-Z]{3,10}\.?\s+\d{1,2},?\s+(\d{4})(\s+\d{1,2}:\d{2}(:\d{2})?(\s+[AP]M)?)?$/i.test(d)||/^\d{1,2}\s+[A-Z]{3,10}\s+\d{4}/i.test(d)},format:function(d,h){return d?r.formatFloat((new Date(d.replace(/(\S)([AP]M)$/i,"$1 $2"))).getTime()||d,h):d},type:"numeric"});r.addParser({id:"shortDate",is:function(d){return/(^\d{1,2}[\/\s]\d{1,2}[\/\s]\d{4})|(^\d{4}[\/\s]\d{1,2}[\/\s]\d{1,2})/.test((d|| "").replace(/\s+/g," ").replace(/[\-.,]/g,"/"))},format:function(d,h,p,y){if(d){p=h.config;var v=p.$headers.filter("[data-column="+y+"]:last");y=v.length&&v[0].dateFormat||r.getData(v,r.getColumnData(h,p.headers,y),"dateFormat")||p.dateFormat;d=d.replace(/\s+/g," ").replace(/[\-.,]/g,"/");"mmddyyyy"===y?d=d.replace(/(\d{1,2})[\/\s](\d{1,2})[\/\s](\d{4})/,"$3/$1/$2"):"ddmmyyyy"===y?d=d.replace(/(\d{1,2})[\/\s](\d{1,2})[\/\s](\d{4})/,"$3/$2/$1"):"yyyymmdd"===y&&(d=d.replace(/(\d{4})[\/\s](\d{1,2})[\/\s](\d{1,2})/, "$1/$2/$3"))}return d?r.formatFloat((new Date(d)).getTime()||d,h):d},type:"numeric"});r.addParser({id:"time",is:function(d){return/^(([0-2]?\d:[0-5]\d)|([0-1]?\d:[0-5]\d\s?([AP]M)))$/i.test(d)},format:function(d,h){return d?r.formatFloat((new Date("2000/01/01 "+d.replace(/(\S)([AP]M)$/i,"$1 $2"))).getTime()||d,h):d},type:"numeric"});r.addParser({id:"metadata",is:function(){return!1},format:function(d,q,p){d=q.config;d=d.parserMetadataName?d.parserMetadataName:"sortValue";return h(p).metadata()[d]}, type:"numeric"});r.addWidget({id:"zebra",priority:90,format:function(d,q,p){var y,v,w,z,D,E=new RegExp(q.cssChildRow,"i"),C=q.$tbodies;q.debug&&(D=new Date);for(d=0;d<C.length;d++)w=0,y=C.eq(d),y=y.children("tr:visible").not(q.selectorRemove),y.each(function(){v=h(this);E.test(this.className)||w++;z=0===w%2;v.removeClass(p.zebra[z?1:0]).addClass(p.zebra[z?0:1])});q.debug&&r.benchmark("Applying Zebra widget",D)},remove:function(d,q,p){var r;q=q.$tbodies;var v=(p.zebra||["even","odd"]).join(" ");for(p= 0;p<q.length;p++)r=h.tablesorter.processTbody(d,q.eq(p),!0),r.children().removeClass(v),h.tablesorter.processTbody(d,r,!1)}})}(jQuery);
/*!
 * ZUI: Standard edition - v1.9.2 - 2020-07-09
 * http://openzui.com
 * GitHub: https://github.com/easysoft/zui.git 
 * Copyright (c) 2020 cnezsoft.com; Licensed MIT
 */
/*! Some code copy from Bootstrap v3.0.0 by @fat and @mdo. (Copyright 2013 Twitter, Inc. Licensed under http://www.apache.org/licenses/)*/
!function(t,e,i){"use strict";if("undefined"==typeof t)throw new Error("ZUI requires jQuery");t.zui||(t.zui=function(e){t.isPlainObject(e)&&t.extend(t.zui,e)});var n={all:-1,left:0,middle:1,right:2},o=0;t.zui({uuid:function(t){var e=1e8*(Date.now()-1580890015292)+1e3*Math.floor(1e5*Math.random())+o++%1e3;return t?e:e.toString(36)},callEvent:function(e,n,o){if(t.isFunction(e)){o!==i&&(e=t.proxy(e,o));var a=e(n);return n&&(n.result=a),!(a!==i&&!a)}return 1},strCode:function(t){var e=0;if(t&&t.length)for(var i=0;i<t.length;++i)e+=i*t.charCodeAt(i);return e},getMouseButtonCode:function(t){return"number"!=typeof t&&(t=n[t]),t!==i&&null!==t||(t=-1),t},defaultLang:"en",clientLang:function(){var i,n=e.config;if("undefined"!=typeof n&&n.clientLang&&(i=n.clientLang),!i){var o=t("html").attr("lang");i=o?o:navigator.userLanguage||navigator.userLanguage||t.zui.defaultLang}return i.replace("-","_").toLowerCase()},langDataMap:{},addLangData:function(e,i,n){var o={};n&&i&&e?(o[i]={},o[i][e]=n):e&&i&&!n?(n=i,t.each(n,function(t){o[t]={},o[t][e]=n[t]})):!e||i||n||t.each(n,function(e){var i=n[e];t.each(i,function(t){o[t]||(o[t]={}),o[t][e]=i[t]})}),t.extend(!0,t.zui.langDataMap,o)},getLangData:function(e,i,n){if(!arguments.length)return t.extend({},t.zui.langDataMap);if(1===arguments.length)return t.extend({},t.zui.langDataMap[e]);if(2===arguments.length){var o=t.zui.langDataMap[e];return o?i?o[i]:o:{}}if(3===arguments.length){i=i||t.zui.clientLang();var o=t.zui.langDataMap[e],a=o?o[i]:{};return t.extend(!0,{},n[i]||n.en||n.zh_cn,a)}return null},lang:function(){return arguments.length&&t.isPlainObject(arguments[arguments.length-1])?t.zui.addLangData.apply(null,arguments):t.zui.getLangData.apply(null,arguments)},_scrollbarWidth:0,checkBodyScrollbar:function(){if(document.body.clientWidth>=e.innerWidth)return 0;if(!t.zui._scrollbarWidth){var i=document.createElement("div");i.className="scrollbar-measure",document.body.appendChild(i),t.zui._scrollbarWidth=i.offsetWidth-i.clientWidth,document.body.removeChild(i)}return t.zui._scrollbarWidth},fixBodyScrollbar:function(){if(t.zui.checkBodyScrollbar()){var e=t("body"),i=parseInt(e.css("padding-right")||0,10);return t.zui._scrollbarWidth&&e.css({paddingRight:i+t.zui._scrollbarWidth,overflowY:"hidden"}),!0}},resetBodyScrollbar:function(){t("body").css({paddingRight:"",overflowY:""})}}),t.fn.callEvent=function(e,n,o){var a=t(this),s=e.indexOf(".zui."),r=s<0?e:e.substring(0,s),l=t.Event(r,n);if(o===i&&s>0&&(o=a.data(e.substring(s+1))),o&&o.options){var d=o.options[r];t.isFunction(d)&&(l.result=t.zui.callEvent(d,l,o))}return a.trigger(l),l},t.fn.callComEvent=function(e,n,o){o===i||t.isArray(o)||(o=[o]);var a,s=this;s.trigger(n,o);var r=e.options[n];return r&&(a=r.apply(e,o)),a}}(jQuery,window,void 0),function(t){"use strict";t.fn.fixOlPd=function(e){return e=e||10,this.each(function(){var i=t(this);i.css("paddingLeft",Math.ceil(Math.log10(i.children().length))*e+10)})},t(function(){t(".ol-pd-fix,.article ol").fixOlPd()})}(jQuery),+function(t){"use strict";var e=function(i,n){this.$element=t(i),this.options=t.extend({},e.DEFAULTS,n),this.isLoading=!1};e.DEFAULTS={loadingText:"loading..."},e.prototype.setState=function(e){var i="disabled",n=this.$element,o=n.is("input")?"val":"html",a=n.data();e+="Text",a.resetText||n.data("resetText",n[o]()),n[o](a[e]||this.options[e]),setTimeout(t.proxy(function(){"loadingText"==e?(this.isLoading=!0,n.addClass(i).attr(i,i)):this.isLoading&&(this.isLoading=!1,n.removeClass(i).removeAttr(i))},this),0)},e.prototype.toggle=function(){var t=!0,e=this.$element.closest('[data-toggle="buttons"]');if(e.length){var i=this.$element.find("input");"radio"==i.prop("type")&&(i.prop("checked")&&this.$element.hasClass("active")?t=!1:e.find(".active").removeClass("active")),t&&i.prop("checked",!this.$element.hasClass("active")).trigger("change")}t&&this.$element.toggleClass("active")};var i=t.fn.button;t.fn.button=function(i){return this.each(function(){var n=t(this),o=n.data("zui.button"),a="object"==typeof i&&i;o||n.data("zui.button",o=new e(this,a)),"toggle"==i?o.toggle():i&&o.setState(i)})},t.fn.button.Constructor=e,t.fn.button.noConflict=function(){return t.fn.button=i,this},t(document).on("click.zui.button.data-api","[data-toggle^=button]",function(e){var i=t(e.target);i.hasClass("btn")||(i=i.closest(".btn")),i.button("toggle"),e.preventDefault()})}(jQuery),+function(t){"use strict";var e='[data-dismiss="alert"]',i="zui.alert",n=function(i){t(i).on("click",e,this.close)};n.prototype.close=function(e){function n(){s.trigger("closed."+i).remove()}var o=t(this),a=o.attr("data-target");a||(a=o.attr("href"),a=a&&a.replace(/.*(?=#[^\s]*$)/,""));var s=t(a);e&&e.preventDefault(),s.length||(s=o.hasClass("alert")?o:o.parent()),s.trigger(e=t.Event("close."+i)),e.isDefaultPrevented()||(s.removeClass("in"),t.support.transition&&s.hasClass("fade")?s.one(t.support.transition.end,n).emulateTransitionEnd(150):n())};var o=t.fn.alert;t.fn.alert=function(e){return this.each(function(){var o=t(this),a=o.data(i);a||o.data(i,a=new n(this)),"string"==typeof e&&a[e].call(o)})},t.fn.alert.Constructor=n,t.fn.alert.noConflict=function(){return t.fn.alert=o,this},t(document).on("click."+i+".data-api",e,n.prototype.close)}(window.jQuery),function(t,e){"use strict";var i="zui.pager",n={page:1,recTotal:0,recPerPage:10},o={zh_cn:{pageOfText:"第 {0} 页",prev:"上一页",next:"下一页",first:"第一页",last:"最后一页","goto":"跳转",pageOf:"第 <strong>{page}</strong> 页",totalPage:"共 <strong>{totalPage}</strong> 页",totalCount:"共 <strong>{recTotal}</strong> 项",pageSize:"每页 <strong>{recPerPage}</strong> 项",itemsRange:"第 <strong>{start}</strong> ~ <strong>{end}</strong> 项",pageOfTotal:"第 <strong>{page}</strong>/<strong>{totalPage}</strong> 页"},zh_tw:{pageOfText:"第 {0} 頁",prev:"上一頁",next:"下一頁",first:"第一頁",last:"最後一頁","goto":"跳轉",pageOf:"第 <strong>{page}</strong> 頁",totalPage:"共 <strong>{totalPage}</strong> 頁",totalCount:"共 <strong>{recTotal}</strong> 項",pageSize:"每頁 <strong>{recPerPage}</strong> 項",itemsRange:"第 <strong>{start}</strong> ~ <strong>{end}</strong> 項",pageOfTotal:"第 <strong>{page}</strong>/<strong>{totalPage}</strong> 頁"},en:{pageOfText:"Page {0}",prev:"Prev",next:"Next",first:"First",last:"Last","goto":"Goto",pageOf:"Page <strong>{page}</strong>",totalPage:"<strong>{totalPage}</strong> pages",totalCount:"Total: <strong>{recTotal}</strong> items",pageSize:"<strong>{recPerPage}</strong> per page",itemsRange:"From <strong>{start}</strong> to <strong>{end}</strong>",pageOfTotal:"Page <strong>{page}</strong> of <strong>{totalPage}</strong>"}},a=function(e,n){var s=this;s.name=i,s.$=t(e),n=s.options=t.extend({},a.DEFAULTS,this.$.data(),n),s.langName=n.lang||t.zui.clientLang(),s.lang=t.zui.getLangData(i,s.langName,o),s.state={},s.set(n.page,n.recTotal,n.recPerPage,!0),s.$.on("click",".pager-goto-btn",function(){var e=t(this).closest(".pager-goto"),i=parseInt(e.find(".pager-goto-input").val());NaN!==i&&s.set(i)}).on("click",".pager-item",function(){var e=t(this).data("page");"number"==typeof e&&e>0&&s.set(e)}).on("click",".pager-size-menu [data-size]",function(){var e=t(this).data("size");"number"==typeof e&&e>0&&s.set(-1,-1,e)})};a.prototype.set=function(e,i,o,a){var s=this;"object"==typeof e&&null!==e&&(o=e.recPerPage,i=e.recTotal,e=e.page);var r=s.state;r||(r=t.extend({},n));var l=t.extend({},r);return"number"==typeof o&&o>0&&(r.recPerPage=o),"number"==typeof i&&i>=0&&(r.recTotal=i),"number"==typeof e&&e>=0&&(r.page=e),r.totalPage=r.recTotal&&r.recPerPage?Math.ceil(r.recTotal/r.recPerPage):1,r.page=Math.max(0,Math.min(r.page,r.totalPage)),r.pageRecCount=r.recTotal,r.page&&r.recTotal&&(r.page<r.totalPage?r.pageRecCount=r.recPerPage:r.page>1&&(r.pageRecCount=r.recTotal-r.recPerPage*(r.page-1))),r.skip=r.page>1?(r.page-1)*r.recPerPage:0,r.start=r.skip+1,r.end=r.skip+r.pageRecCount,r.prev=r.page>1?r.page-1:0,r.next=r.page<r.totalPage?r.page+1:0,s.state=r,a||l.page===r.page&&l.recTotal===r.recTotal&&l.recPerPage===r.recPerPage||s.$.callComEvent(s,"onPageChange",[r,l]),s.render()},a.prototype.createLinkItem=function(i,n,o){var a=this;n===e&&(n=i);var s=t('<a title="'+a.lang.pageOfText.format(i)+'" class="pager-item" data-page="'+i+'"/>').attr("href",i?a.createLink(i,a.state):"###").html(n);return o||(s=t("<li />").append(s).toggleClass("active",i===a.state.page).toggleClass("disabled",!i||i===a.state.page)),s},a.prototype.createNavItems=function(t){var i=this,n=i.$,o=i.state,a=o.totalPage,s=o.page,r=function(t,o){if(t===!1)return void n.append(i.createLinkItem(0,o||i.options.navEllipsisItem));o===e&&(o=t);for(var a=t;a<=o;++a)n.append(i.createLinkItem(a))};t===e&&(t=i.options.maxNavCount||10),r(1),a>1&&(a<=t?r(2,a):s<t-2?(r(2,t-2),r(!1),r(a)):s>a-t+2?(r(!1),r(a-t+2,a)):(r(!1),r(s-Math.ceil((t-4)/2),s+Math.floor((t-4)/2)),r(!1),r(a)))},a.prototype.createGoto=function(){var e=this,i=this.state,n=t('<div class="input-group pager-goto" style="width: '+(35+9*(i.page+"").length+25+12*e.lang["goto"].length)+'px"><input value="'+i.page+'" type="number" min="1" max="'+i.totalPage+'" placeholder="'+i.page+'" class="form-control pager-goto-input"><span class="input-group-btn"><button class="btn pager-goto-btn" type="button">'+e.lang["goto"]+"</button></span></div>");return n},a.prototype.createSizeMenu=function(){var e=this,i=this.state,n=t('<ul class="dropdown-menu"></ul>'),o=e.options.pageSizeOptions;"string"==typeof o&&(o=o.split(","));for(var a=0;a<o.length;++a){var s=o[a];"string"==typeof s&&(s=parseInt(s));var r=t('<li><a href="###" data-size="'+s+'">'+s+"</a></li>").toggleClass("active",s===i.recPerPage);n.append(r)}return t('<div class="btn-group pager-size-menu"><button type="button" class="btn dropdown-toggle" data-toggle="dropdown">'+e.lang.pageSize.format(i)+' <span class="caret"></span></button></div>').addClass(e.options.menuDirection).append(n)},a.prototype.createElement=function(e,i,n){var o=this,a=t.proxy(o.createLinkItem,o),s=o.lang;switch(e){case"prev":return a(n.prev,s.prev);case"prev_icon":return a(n.prev,'<i class="icon '+o.options.prevIcon+'"></i>');case"next":return a(n.next,s.next);case"next_icon":return a(n.next,'<i class="icon '+o.options.nextIcon+'"></i>');case"first":return a(1,s.first);case"first_icon":return a(1,'<i class="icon '+o.options.firstIcon+'"></i>');case"last":return a(n.totalPage,s.last);case"last_icon":return a(n.totalPage,'<i class="icon '+o.options.lastIcon+'"></i>');case"space":case"|":return t('<li class="space" />');case"nav":case"pages":return void o.createNavItems();case"total_text":return t(('<div class="pager-label">'+s.totalCount+"</div>").format(n));case"page_text":return t(('<div class="pager-label">'+s.pageOf+"</div>").format(n));case"total_page_text":return t(('<div class="pager-label">'+s.totalPage+"</div>").format(n));case"page_of_total_text":return t(('<div class="pager-label">'+s.pageOfTotal+"</div>").format(n));case"page_size_text":return t(('<div class="pager-label">'+s.pageSize+"</div>").format(n));case"items_range_text":return t(('<div class="pager-label">'+s.itemsRange+"</div>").format(n));case"goto":return o.createGoto();case"size_menu":return o.createSizeMenu();default:return t("<li/>").html(e.format(n))}},a.prototype.createLink=function(i,n){i===e&&(i=this.state.page),n===e&&(n=this.state);var o=this.options.linkCreator;return"string"==typeof o?o.format(t.extend({},n,{page:i})):t.isFunction(o)?o(i,n):"#page="+i},a.prototype.render=function(e){var i=this,n=i.state,o=i.options.elementCreator||i.createElement,a=t.isPlainObject(o);e=e||i.elements||i.options.elements,"string"==typeof e&&(e=e.split(",")),i.elements=e,i.$.empty();for(var s=0;s<e.length;++s){var r=t.trim(e[s]),l=a?o[r]||o:o,d=l.call(i,r,i.$,n);d===!1&&(d=i.createElement(r,i.$,n)),d instanceof t&&("LI"!==d[0].tagName&&(d=t("<li/>").append(d)),i.$.append(d))}var c=null;return i.$.children("li").each(function(){var e=t(this),i=!!e.children(".pager-item").length;c?c.toggleClass("pager-item-right",!i):i&&e.addClass("pager-item-left"),c=i?e:null}),c&&c.addClass("pager-item-right"),i.$.callComEvent(i,"onRender",[n]),i},a.DEFAULTS=t.extend({elements:["first_icon","prev_icon","pages","next_icon","last_icon","page_of_total_text","items_range_text","total_text"],prevIcon:"icon-double-angle-left",nextIcon:"icon-double-angle-right",firstIcon:"icon-step-backward",lastIcon:"icon-step-forward",navEllipsisItem:'<i class="icon icon-ellipsis-h"></i>',maxNavCount:10,menuDirection:"dropdown",pageSizeOptions:[10,20,30,50,100]},n),t.fn.pager=function(e){return this.each(function(){var n=t(this),o=n.data(i),s="object"==typeof e&&e;o||n.data(i,o=new a(this,s)),"string"==typeof e&&o[e]()})},a.NAME=i,a.LANG=o,t.fn.pager.Constructor=a,t(function(){t('[data-ride="pager"]').pager()})}(jQuery,void 0),+function(t){"use strict";var e="zui.tab",i=function(e){this.element=t(e)};i.prototype.show=function(){var i=this.element,n=i.closest("ul:not(.dropdown-menu)"),o=i.attr("data-target")||i.attr("data-tab");if(o||(o=i.attr("href"),o=o&&o.replace(/.*(?=#[^\s]*$)/,"")),!i.parent("li").hasClass("active")){var a=n.find(".active:last a")[0],s=t.Event("show."+e,{relatedTarget:a});if(i.trigger(s),!s.isDefaultPrevented()){var r=t(o);this.activate(i.parent("li"),n),this.activate(r,r.parent(),function(){i.trigger({type:"shown."+e,relatedTarget:a})})}}},i.prototype.activate=function(e,i,n){function o(){a.removeClass("active").find("> .dropdown-menu > .active").removeClass("active"),e.addClass("active"),s?(e[0].offsetWidth,e.addClass("in")):e.removeClass("fade"),e.parent(".dropdown-menu")&&e.closest("li.dropdown").addClass("active"),n&&n()}var a=i.find("> .active"),s=n&&t.support.transition&&a.hasClass("fade");s?a.one(t.support.transition.end,o).emulateTransitionEnd(150):o(),a.removeClass("in")};var n=t.fn.tab;t.fn.tab=function(n){return this.each(function(){var o=t(this),a=o.data(e);a||o.data(e,a=new i(this)),"string"==typeof n&&a[n]()})},t.fn.tab.Constructor=i,t.fn.tab.noConflict=function(){return t.fn.tab=n,this},t(document).on("click.zui.tab.data-api",'[data-toggle="tab"], [data-tab]',function(e){e.preventDefault(),t(this).tab("show")})}(window.jQuery),+function(t){"use strict";function e(){var t=document.createElement("bootstrap"),e={WebkitTransition:"webkitTransitionEnd",MozTransition:"transitionend",OTransition:"oTransitionEnd otransitionend",transition:"transitionend"};for(var i in e)if(void 0!==t.style[i])return{end:e[i]};return!1}t.fn.emulateTransitionEnd=function(e){var i=!1,n=this;t(this).one("bsTransitionEnd",function(){i=!0});var o=function(){i||t(n).trigger(t.support.transition.end)};return setTimeout(o,e),this},t(function(){t.support.transition=e(),t.support.transition&&(t.event.special.bsTransitionEnd={bindType:t.support.transition.end,delegateType:t.support.transition.end,handle:function(e){if(t(e.target).is(this))return e.handleObj.handler.apply(this,arguments)}})})}(jQuery),+function(t){"use strict";var e="zui.collapse",i=function(e,n){this.$element=t(e),this.options=t.extend({},i.DEFAULTS,n),this.transitioning=null,this.options.parent&&(this.$parent=t(this.options.parent)),this.options.toggle&&this.toggle()};i.DEFAULTS={toggle:!0},i.prototype.dimension=function(){var t=this.$element.hasClass("width");return t?"width":"height"},i.prototype.show=function(){if(!this.transitioning&&!this.$element.hasClass("in")){var i=t.Event("show."+e);if(this.$element.trigger(i),!i.isDefaultPrevented()){var n=this.$parent&&this.$parent.find(".in");if(n&&n.length){var o=n.data(e);if(o&&o.transitioning)return;n.collapse("hide"),o||n.data(e,null)}var a=this.dimension();this.$element.removeClass("collapse").addClass("collapsing")[a](0),this.transitioning=1;var s=function(){this.$element.removeClass("collapsing").addClass("in")[a]("auto"),this.transitioning=0,this.$element.trigger("shown."+e)};if(!t.support.transition)return s.call(this);var r=t.camelCase(["scroll",a].join("-"));this.$element.one(t.support.transition.end,t.proxy(s,this)).emulateTransitionEnd(350)[a](this.$element[0][r])}}},i.prototype.hide=function(){if(!this.transitioning&&this.$element.hasClass("in")){var i=t.Event("hide."+e);if(this.$element.trigger(i),!i.isDefaultPrevented()){var n=this.dimension();this.$element[n](this.$element[n]())[0].offsetHeight,this.$element.addClass("collapsing").removeClass("collapse").removeClass("in"),this.transitioning=1;var o=function(){this.transitioning=0,this.$element.trigger("hidden."+e).removeClass("collapsing").addClass("collapse")};return t.support.transition?void this.$element[n](0).one(t.support.transition.end,t.proxy(o,this)).emulateTransitionEnd(350):o.call(this)}}},i.prototype.toggle=function(){this[this.$element.hasClass("in")?"hide":"show"]()};var n=t.fn.collapse;t.fn.collapse=function(n){return this.each(function(){var o=t(this),a=o.data(e),s=t.extend({},i.DEFAULTS,o.data(),"object"==typeof n&&n);a||o.data(e,a=new i(this,s)),"string"==typeof n&&a[n]()})},t.fn.collapse.Constructor=i,t.fn.collapse.noConflict=function(){return t.fn.collapse=n,this},t(document).on("click."+e+".data-api","[data-toggle=collapse]",function(i){var n,o=t(this),a=o.attr("data-target")||i.preventDefault()||(n=o.attr("href"))&&n.replace(/.*(?=#[^\s]+$)/,""),s=t(a),r=s.data(e),l=r?"toggle":o.data(),d=o.attr("data-parent"),c=d&&t(d);r&&r.transitioning||(c&&c.find('[data-toggle=collapse][data-parent="'+d+'"]').not(o).addClass("collapsed"),o[s.hasClass("in")?"addClass":"removeClass"]("collapsed")),s.collapse(l)})}(window.jQuery),function(t,e){"use strict";var i=1200,n=992,o=768,a=e(t),s=function(){var t=a.width();e("html").toggleClass("screen-desktop",t>=n&&t<i).toggleClass("screen-desktop-wide",t>=i).toggleClass("screen-tablet",t>=o&&t<n).toggleClass("screen-phone",t<o).toggleClass("device-mobile",t<n).toggleClass("device-desktop",t>=n)},r="",l=navigator.userAgent;l.match(/(iPad|iPhone|iPod)/i)?r+=" os-ios":l.match(/android/i)?r+=" os-android":l.match(/Win/i)?r+=" os-windows":l.match(/Mac/i)?r+=" os-mac":l.match(/Linux/i)?r+=" os-linux":l.match(/X11/i)&&(r+=" os-unix"),"ontouchstart"in document.documentElement&&(r+=" is-touchable"),e("html").addClass(r),a.resize(s),s()}(window,jQuery),function(t){"use strict";var e={zh_cn:'您的浏览器版本过低，无法体验所有功能，建议升级或者更换浏览器。 <a href="https://browsehappy.com/" target="_blank" class="alert-link">了解更多...</a>',zh_tw:'您的瀏覽器版本過低，無法體驗所有功能，建議升級或者更换瀏覽器。<a href="https://browsehappy.com/" target="_blank" class="alert-link">了解更多...</a>',en:'Your browser is too old, it has been unable to experience the colorful internet. We strongly recommend that you upgrade a better one. <a href="https://browsehappy.com/" target="_blank" class="alert-link">Learn more...</a>'},i=function(){for(var t=!1,e=11;e>5;e--)if(this.isIE(e)){t=e;break}this.ie=t,this.cssHelper()};i.prototype.cssHelper=function(){var e=this.ie,i=t("html");i.toggleClass("ie",e).removeClass("ie-6 ie-7 ie-8 ie-9 ie-10"),e&&i.addClass("ie-"+e).toggleClass("gt-ie-7 gte-ie-8 support-ie",e>=8).toggleClass("lte-ie-7 lt-ie-8 outdated-ie",e<8).toggleClass("gt-ie-8 gte-ie-9",e>=9).toggleClass("lte-ie-8 lt-ie-9",e<9).toggleClass("gt-ie-9 gte-ie-10",e>=10).toggleClass("lte-ie-9 lt-ie-10",e<10).toggleClass("gt-ie-10 gte-ie-11",e>=11).toggleClass("lte-ie-10 lt-ie-11",e<11)},i.prototype.tip=function(i){var n=t("#browseHappyTip");n.length||(n=t('<div id="browseHappyTip" class="alert alert-dismissable alert-danger-inverse alert-block" style="position: relative; z-index: 99999"><button type="button" class="close" data-dismiss="alert" aria-hidden="true">×</button><div class="container"><div class="content text-center"></div></div></div>'),n.prependTo("body")),i||(i=t.zui.getLangData("zui.browser",t.zui.clientLang(),e),"object"==typeof i&&(i=i.tip)),n.find(".content").html(i)},i.prototype.isIE=function(t){if(11===t)return this.isIE11();if(10===t)return this.isIE10();if(!t&&(this.isIE11()||this.isIE10()))return!0;var e=document.createElement("b");return e.innerHTML="<!--[if IE "+(t||"")+"]><i></i><![endif]-->",1===e.getElementsByTagName("i").length},i.prototype.isIE10=function(){return navigator.appVersion.indexOf("MSIE 10")!==-1},i.prototype.isIE11=function(){var t=navigator.userAgent;return t.indexOf("Trident")!==-1&&t.indexOf("rv:11")!==-1},t.zui({browser:new i}),t(function(){t("body").hasClass("disabled-browser-tip")||t.zui.browser.ie&&t.zui.browser.ie<8&&t.zui.browser.tip()})}(jQuery),function(t){"use strict";const e=864e5,i=function(t){return t instanceof Date||("number"==typeof t&&t<1e10&&(t*=1e3),t=new Date(t)),t},n=function(t){return i(t).getTime()},o=function(t,e){t=i(t),void 0===e&&(e="yyyy-MM-dd hh:mm:ss");var n={"M+":t.getMonth()+1,"d+":t.getDate(),"h+":t.getHours(),"m+":t.getMinutes(),"s+":t.getSeconds(),"q+":Math.floor((t.getMonth()+3)/3),"S+":t.getMilliseconds()};/(y+)/i.test(e)&&(e=e.replace(RegExp.$1,(t.getFullYear()+"").substr(4-RegExp.$1.length)));for(var o in n)new RegExp("("+o+")").test(e)&&(e=e.replace(RegExp.$1,1==RegExp.$1.length?n[o]:("00"+n[o]).substr((""+n[o]).length)));return e},a=function(t,e){return t.setTime(t.getTime()+e),t},s=function(t,i){return a(t,i*e)},r=function(t){return new Date(i(t).getTime())},l=function(t){return t%4===0&&t%100!==0||t%400===0},d=function(t,e){return[31,l(t)?29:28,31,30,31,30,31,31,30,31,30,31][e]},c=function(t){return d(t.getFullYear(),t.getMonth())},p=function(t){return t.setHours(0),t.setMinutes(0),t.setSeconds(0),t.setMilliseconds(0),t},u=function(t,e){var i=t.getDate();return t.setDate(1),t.setMonth(t.getMonth()+e),t.setDate(Math.min(i,c(t))),t},h=function(t,e){e=e||1;for(var i=new Date(t.getTime());i.getDay()!=e;)i=s(i,-1);return p(i)},f=function(t,e){return t.toDateString()===e.toDateString()},g=function(t,e){var i=h(t),n=s(r(i),7);return e>=i&&e<n},m=function(t,e){return t.getFullYear()===e.getFullYear()},v={formatDate:o,createDate:i,date:{ONEDAY_TICKS:e,create:i,getTimestamp:n,format:o,addMilliseconds:a,addDays:s,cloneDate:r,isLeapYear:l,getDaysInMonth:d,getDaysOfThisMonth:c,clearTime:p,addMonths:u,getLastWeekday:h,isSameDay:f,isSameWeek:g,isSameYear:m}};t.$&&t.$.zui?$.zui(v):t.dateHelper=v.date,t.noDatePrototypeHelper||(Date.ONEDAY_TICKS=e,Date.prototype.format||(Date.prototype.format=function(t){return o(this,t)}),Date.prototype.addMilliseconds||(Date.prototype.addMilliseconds=function(t){return a(this,t)}),Date.prototype.addDays||(Date.prototype.addDays=function(t){return s(this,t)}),Date.prototype.clone||(Date.prototype.clone=function(){return r(this)}),Date.isLeapYear||(Date.isLeapYear=function(t){return l(t)}),Date.getDaysInMonth||(Date.getDaysInMonth=function(t,e){return d(t,e)}),Date.prototype.isLeapYear||(Date.prototype.isLeapYear=function(){return l(this.getFullYear())}),Date.prototype.clearTime||(Date.prototype.clearTime=function(){return p(this)}),Date.prototype.getDaysInMonth||(Date.prototype.getDaysInMonth=function(){return c(this)}),Date.prototype.addMonths||(Date.prototype.addMonths=function(t){return u(this,t)}),Date.prototype.getLastWeekday||(Date.prototype.getLastWeekday=function(t){return h(this,t)}),Date.prototype.isSameDay||(Date.prototype.isSameDay=function(t){return f(t,this)}),Date.prototype.isSameWeek||(Date.prototype.isSameWeek=function(t){return g(t,this)}),Date.prototype.isSameYear||(Date.prototype.isSameYear=function(t){return m(this,t)}),Date.create||(Date.create=function(t){return i(t)}),Date.timestamp||(Date.timestamp=function(t){return n(t)}))}(window),function(){"use strict";const t=function(t,e){if(arguments.length>1){var i;if(2==arguments.length&&"object"==typeof e)for(var n in e)void 0!==e[n]&&(i=new RegExp("({"+n+"})","g"),t=t.replace(i,e[n]));else for(var o=1;o<arguments.length;o++)void 0!==arguments[o]&&(i=new RegExp("({["+(o-1)+"]})","g"),t=t.replace(i,arguments[o]))}return t},e=function(t){if(null!==t){var e,i;return i=/\d*/i,e=t.match(i),e==t}return!1},i={formatString:t,string:{format:t,isNum:e}};window.$&&window.$.zui?$.zui(i):window.stringHelper=i.string,window.noStringPrototypeHelper||(String.prototype.format||(String.prototype.format=function(){var e=[].slice.call(arguments);return e.unshift(this),t.apply(this,e)}),String.prototype.isNum||(String.prototype.isNum=function(){return e(this)}),String.prototype.endsWith||(String.prototype.endsWith=function(t,e){return(void 0===e||e>this.length)&&(e=this.length),this.substring(e-t.length,e)===t}),String.prototype.startsWith||Object.defineProperty(String.prototype,"startsWith",{value:function(t,e){return e=!e||e<0?0:+e,this.substring(e,e+t.length)===t}}),String.prototype.includes||(String.prototype.includes=function(){return String.prototype.indexOf.apply(this,arguments)!==-1}))}(),/*!
 * jQuery resize event - v1.1
 * http://benalman.com/projects/jquery-resize-plugin/
 * Copyright (c) 2010 "Cowboy" Ben Alman
 * MIT & GPL http://benalman.com/about/license/
 */
function(t,e,i){"$:nomunge";function n(){o=e[r](function(){a.each(function(){var e=t(this),i=e.width(),n=e.height(),o=t.data(this,d);i===o.w&&n===o.h||e.trigger(l,[o.w=i,o.h=n])}),n()},s[c])}var o,a=t([]),s=t.resize=t.extend(t.resize,{}),r="setTimeout",l="resize",d=l+"-special-event",c="delay",p="throttleWindow";s[c]=250,s[p]=!0,t.event.special[l]={setup:function(){if(!s[p]&&this[r])return!1;var e=t(this);a=a.add(e),t.data(this,d,{w:e.width(),h:e.height()}),1===a.length&&n()},teardown:function(){if(!s[p]&&this[r])return!1;var e=t(this);a=a.not(e),e.removeData(d),a.length||clearTimeout(o)},add:function(e){function n(e,n,a){var s=t(this),r=t.data(this,d)||{};r.w=n!==i?n:s.width(),r.h=a!==i?a:s.height(),o.apply(this,arguments)}if(!s[p]&&this[r])return!1;var o;return t.isFunction(e)?(o=e,n):(o=e.handler,void(e.handler=n))}}}(jQuery,this),+function(t){"use strict";function e(n,o){var a,s=t.proxy(this.process,this);this.$element=t(t(n).is("body")?window:n),this.$body=t("body"),this.$scrollElement=this.$element.on("scroll."+i+".data-api",s),this.options=t.extend({},e.DEFAULTS,o),this.selector||(this.selector=(this.options.target||(a=t(n).attr("href"))&&a.replace(/.*(?=#[^\s]+$)/,"")||"")+" .nav li > a"),this.offsets=t([]),this.targets=t([]),this.activeTarget=null,this.refresh(),this.process()}var i="zui.scrollspy";e.DEFAULTS={offset:10},e.prototype.refresh=function(){var e=this.$element[0]==window?"offset":"position";this.offsets=t([]),this.targets=t([]);var i=this;this.$body.find(this.selector).map(function(){var n=t(this),o=n.data("target")||n.attr("href"),a=/^#./.test(o)&&t(o);return a&&a.length&&a.is(":visible")&&[[a[e]().top+(!t.isWindow(i.$scrollElement.get(0))&&i.$scrollElement.scrollTop()),o]]||null}).sort(function(t,e){return t[0]-e[0]}).each(function(){i.offsets.push(this[0]),i.targets.push(this[1])})},e.prototype.process=function(){var t,e=this.$scrollElement.scrollTop()+this.options.offset,i=this.$scrollElement[0].scrollHeight||this.$body[0].scrollHeight,n=i-this.$scrollElement.height(),o=this.offsets,a=this.targets,s=this.activeTarget;if(e>=n)return s!=(t=a.last()[0])&&this.activate(t);if(s&&e<=o[0])return s!=(t=a[0])&&this.activate(t);for(t=o.length;t--;)s!=a[t]&&e>=o[t]&&(!o[t+1]||e<=o[t+1])&&this.activate(a[t])},e.prototype.activate=function(e){this.activeTarget=e,t(this.selector).parentsUntil(this.options.target,".active").removeClass("active");var n=this.selector+'[data-target="'+e+'"],'+this.selector+'[href="'+e+'"]',o=t(n).parents("li").addClass("active");o.parent(".dropdown-menu").length&&(o=o.closest("li.dropdown").addClass("active")),o.trigger("activate."+i)};var n=t.fn.scrollspy;t.fn.scrollspy=function(n){return this.each(function(){var o=t(this),a=o.data(i),s="object"==typeof n&&n;a||o.data(i,a=new e(this,s)),"string"==typeof n&&a[n]()})},t.fn.scrollspy.Constructor=e,t.fn.scrollspy.noConflict=function(){return t.fn.scrollspy=n,this},t(window).on("load",function(){t('[data-spy="scroll"]').each(function(){var e=t(this);e.scrollspy(e.data())})})}(jQuery),function(t,e){"use strict";var i,n,o="localStorage",a="page_"+t.location.pathname+t.location.search,s=function(){this.silence=!0;try{o in t&&t[o]&&t[o].setItem&&(this.enable=!0,i=t[o])}catch(s){}this.enable||(n={},i={getLength:function(){var t=0;return e.each(n,function(){t++}),t},key:function(t){var i,o=0;return e.each(n,function(e){return o===t?(i=e,!1):void o++}),i},removeItem:function(t){delete n[t]},getItem:function(t){return n[t]},setItem:function(t,e){n[t]=e},clear:function(){n={}}}),this.storage=i,this.page=this.get(a,{})};s.prototype.pageSave=function(){if(e.isEmptyObject(this.page))this.remove(a);else{var t,i=[];for(t in this.page){var n=this.page[t];null===n&&i.push(t)}for(t=i.length-1;t>=0;t--)delete this.page[i[t]];this.set(a,this.page)}},s.prototype.pageRemove=function(t){"undefined"!=typeof this.page[t]&&(this.page[t]=null,this.pageSave())},s.prototype.pageClear=function(){this.page={},this.pageSave()},s.prototype.pageGet=function(t,e){var i=this.page[t];return void 0===e||null!==i&&void 0!==i?i:e},s.prototype.pageSet=function(t,i){e.isPlainObject(t)?e.extend(!0,this.page,t):this.page[this.serialize(t)]=i,this.pageSave()},s.prototype.check=function(){if(!this.enable&&!this.silence)throw new Error("Browser not support localStorage or enable status been set true.");return this.enable},s.prototype.length=function(){return this.check()?i.getLength?i.getLength():i.length:0},s.prototype.removeItem=function(t){return i.removeItem(t),this},s.prototype.remove=function(t){return this.removeItem(t)},s.prototype.getItem=function(t){return i.getItem(t)},s.prototype.get=function(t,e){var i=this.deserialize(this.getItem(t));return"undefined"!=typeof i&&null!==i||"undefined"==typeof e?i:e},s.prototype.key=function(t){return i.key(t)},s.prototype.setItem=function(t,e){return i.setItem(t,e),this},s.prototype.set=function(t,e){return void 0===e?this.remove(t):(this.setItem(t,this.serialize(e)),this)},s.prototype.clear=function(){return i.clear(),this},s.prototype.forEach=function(t){for(var e=this.length(),n=e-1;n>=0;n--){var o=i.key(n);t(o,this.get(o))}return this},s.prototype.getAll=function(){var t={};return this.forEach(function(e,i){t[e]=i}),t},s.prototype.serialize=function(t){return"string"==typeof t?t:JSON.stringify(t)},s.prototype.deserialize=function(t){if("string"==typeof t)try{return JSON.parse(t)}catch(e){return t||void 0}},e.zui({store:new s})}(window,jQuery),function(t){"use strict";var e="zui.searchBox",i=function(e,n){var o=this;o.name=name,o.$=t(e),o.options=n=t.extend({},i.DEFAULTS,o.$.data(),n);var a=o.$.is(n.inputSelector)?o.$:o.$.find(n.inputSelector);if(a.length){var s=function(){o.changeTimer&&(clearTimeout(o.changeTimer),o.changeTimer=null)},r=function(){s();var t=o.getSearch();if(t!==o.lastValue){var e=""===t;a.toggleClass("empty",e),o.$.callComEvent(o,"onSearchChange",[t,e]),o.lastValue=t}};o.$input=a=a.first(),a.on(n.listenEvent,function(t){o.changeTimer=setTimeout(function(){r()},n.changeDelay)}).on("focus",function(t){a.addClass("focus"),o.$.callComEvent(o,"onFocus",[t])}).on("blur",function(t){a.removeClass("focus"),o.$.callComEvent(o,"onBlur",[t])}).on("keydown",function(t){var e=0,i=t.which;27===i&&n.escToClear?(this.setSearch("",!0),r(),e=1):13===i&&n.onPressEnter&&(r(),o.$.callComEvent(o,"onPressEnter",[t]));var a=o.$.callComEvent(o,"onKeyDown",[t]);a===!1&&(e=1),e&&t.preventDefault()}),o.$.on("click",".search-clear-btn",function(t){o.setSearch("",!0),r(),o.focus(),t.preventDefault()}),r()}else console.error("ZUI: search box init error, cannot find search box input element.")};i.DEFAULTS={inputSelector:'input[type="search"],input[type="text"]',listenEvent:"change input paste",changeDelay:500},i.prototype.getSearch=function(){return this.$input&&t.trim(this.$input.val())},i.prototype.setSearch=function(t,e){var i=this.$input;i&&(i.val(t),e||i.trigger("change"))},i.prototype.focus=function(){this.$input&&this.$input.focus()},t.fn.searchBox=function(n){return this.each(function(){var o=t(this),a=o.data(e),s="object"==typeof n&&n;a||o.data(e,a=new i(this,s)),"string"==typeof n&&a[n]()})},i.NAME=e,t.fn.searchBox.Constructor=i}(jQuery),function(t,e){"use strict";var i="zui.draggable",n={container:"body",move:!0},o=0,a=function(e,i){var a=this;a.$=t(e),a.id=o++,a.options=t.extend({},n,a.$.data(),i),a.init()};a.DEFAULTS=n,a.NAME=i,a.prototype.init=function(){var n,o,a,s,r,l=this,d=l.$,c="before",p="drag",u="finish",h="."+i+"."+l.id,f="mousedown"+h,g="mouseup"+h,m="mousemove"+h,v=l.options,y=v.selector,b=v.handle,w=d,C=t.isFunction(v.move),x=function(t){var e=t.pageX,i=t.pageY;r=!0;var o={left:e-a.x,top:i-a.y};w.removeClass("drag-ready").addClass("dragging"),v.move&&(C?v.move(o,w):w.css(o)),v[p]&&v[p]({event:t,element:w,startOffset:a,pos:o,offset:{x:e-n.x,y:i-n.y},smallOffset:{x:e-s.x,y:i-s.y}}),s.x=e,s.y=i,v.stopPropagation&&t.stopPropagation()},$=function(i){if(t(e).off(h),!r)return void w.removeClass("drag-ready");var o={left:i.pageX-a.x,top:i.pageY-a.y};w.removeClass("drag-ready dragging"),v.move&&(C?v.move(o,w):w.css(o)),v[u]&&v[u]({event:i,element:w,startOffset:a,pos:o,offset:{x:i.pageX-n.x,y:i.pageY-n.y},smallOffset:{x:i.pageX-s.x,y:i.pageY-s.y}}),i.preventDefault(),v.stopPropagation&&i.stopPropagation()},T=function(i){var l=t.zui.getMouseButtonCode(v.mouseButton);if(!(l>-1&&i.button!==l)){var d=t(this);if(y&&(w=b?d.closest(y):d),v[c]){var p=v[c]({event:i,element:w});if(p===!1)return}var u=t(v.container),h=w.offset();o=u.offset(),n={x:i.pageX,y:i.pageY},a={x:i.pageX-h.left+o.left,y:i.pageY-h.top+o.top},s=t.extend({},n),r=!1,w.addClass("drag-ready"),i.preventDefault(),v.stopPropagation&&i.stopPropagation(),t(e).on(m,x).on(g,$)}};b?d.on(f,b,T):y?d.on(f,y,T):d.on(f,T)},a.prototype.destroy=function(){var n="."+i+"."+this.id;this.$.off(n),t(e).off(n),this.$.data(i,null)},t.fn.draggable=function(e){return this.each(function(){var n=t(this),o=n.data(i),s="object"==typeof e&&e;o||n.data(i,o=new a(this,s)),"string"==typeof e&&o[e]()})},t.fn.draggable.Constructor=a}(jQuery,document),function(t,e,i){"use strict";var n="zui.droppable",o={target:".droppable-target",deviation:5,sensorOffsetX:0,sensorOffsetY:0,dropToClass:"drop-to"},a=0,s=function(e,i){var n=this;n.id=a++,n.$=t(e),n.options=t.extend({},o,n.$.data(),i),n.init()};s.DEFAULTS=o,s.NAME=n,s.prototype.trigger=function(e,i){return t.zui.callEvent(this.options[e],i,this)},s.prototype.init=function(){var o,a,s,r,l,d,c,p,u,h,f,g,m,v=this,y=v.$,b=v.options,w=b.deviation,C="."+n+"."+v.id,x="mousedown"+C,$="mouseup"+C,T="mousemove"+C,D=b.selector,S=b.handle,k=b.flex,z=b.container,E=b.canMoveHere,P=b.dropToClass,I=y,M=!1,O=z?t(b.container).first():D?y:t("body"),L=function(e){if(M&&(f={left:e.pageX,top:e.pageY},!(i.abs(f.left-p.left)<w&&i.abs(f.top-p.top)<w))){if(null===s){var n=O.css("position");"absolute"!=n&&"relative"!=n&&"fixed"!=n&&(d=n,O.css("position","relative")),s=I.clone().removeClass("drag-from").addClass("drag-shadow").css({position:"absolute",width:I.outerWidth(),transition:"none"}).appendTo(O),I.addClass("dragging"),v.trigger("start",{event:e,element:I,shadowElement:s,targets:o})}var c={left:f.left-h.left,top:f.top-h.top},m={left:c.left-u.left,top:c.top-u.top};s.css(m),t.extend(g,f);var y=!1;r=!1,k||o.removeClass(P);var C=null;if(o.each(function(){var e=t(this),i=e.offset(),n=e.outerWidth(),o=e.outerHeight(),a=i.left+b.sensorOffsetX,s=i.top+b.sensorOffsetY;if(f.left>a&&f.top>s&&f.left<a+n&&f.top<s+o&&(C&&C.removeClass(P),C=e,!b.nested))return!1}),C){r=!0;var x=C.data("id");I.data("id")!=x&&(l=!1),(null===a||a.data("id")!==x&&!l)&&(y=!0),a=C,k&&o.removeClass(P),a.addClass(P)}k?null!==a&&a.length&&(r=!0):(I.toggleClass("drop-in",r),s.toggleClass("drop-in",r)),E&&E(I,a)===!1||v.trigger("drag",{event:e,isIn:r,target:a,element:I,isNew:y,selfTarget:l,clickOffset:h,offset:c,position:{left:c.left-u.left,top:c.top-u.top},mouseOffset:f}),e.preventDefault()}},j=function(i){if(t(e).off(C),clearTimeout(m),M){if(M=!1,d&&O.css("position",d),null===s)return I.removeClass("drag-from"),void v.trigger("always",{event:i,cancel:!0});r||(a=null);var n=!0;f=i?{left:i.pageX,top:i.pageY}:g;var c={left:f.left-h.left,top:f.top-h.top},p={left:f.left-g.left,top:f.top-g.top};g.left=f.left,g.top=f.top;var y={event:i,isIn:r,target:a,element:I,isNew:!l&&null!==a,selfTarget:l,offset:c,mouseOffset:f,position:{left:c.left-u.left,top:c.top-u.top},lastMouseOffset:g,moveOffset:p};n=v.trigger("beforeDrop",y),n&&r&&v.trigger("drop",y),o.removeClass(P),I.removeClass("dragging").removeClass("drag-from"),s.remove(),s=null,v.trigger("finish",y),v.trigger("always",y),i&&i.preventDefault()}},A=function(i){var n=t.zui.getMouseButtonCode(b.mouseButton);if(!(n>-1&&i.button!==n)){var f=t(this);D&&(I=S?f.closest(D):f),I.hasClass("drag-shadow")||b.before&&b.before({event:i,element:I})===!1||(M=!0,o=t.isFunction(b.target)?b.target(I,y):O.find(b.target),a=null,s=null,r=!1,l=!0,d=null,c=I.offset(),u=O.offset(),u.top=u.top-O.scrollTop(),u.left=u.left-O.scrollLeft(),p={left:i.pageX,top:i.pageY},g=t.extend({},p),h={left:p.left-c.left,top:p.top-c.top},I.addClass("drag-from"),t(e).on(T,L).on($,j),m=setTimeout(function(){t(e).on(x,j)},10),i.preventDefault(),b.stopPropagation&&i.stopPropagation())}};S?y.on(x,S,A):D?y.on(x,D,A):y.on(x,A)},s.prototype.destroy=function(){var i="."+n+"."+this.id;this.$.off(i),t(e).off(i),this.$.data(n,null)},s.prototype.reset=function(){this.destroy(),this.init()},t.fn.droppable=function(e){return this.each(function(){var i=t(this),o=i.data(n),a="object"==typeof e&&e;o||i.data(n,o=new s(this,a)),"string"==typeof e&&o[e]()})},t.fn.droppable.Constructor=s}(jQuery,document,Math),+function(t,e){"use strict";function i(e,i,a){return this.each(function(){var s=t(this),r=s.data(n),l=t.extend({},o.DEFAULTS,s.data(),"object"==typeof e&&e);r||s.data(n,r=new o(this,l)),"string"==typeof e?r[e](i,a):l.show&&r.show(i,a)})}var n="zui.modal",o=function(i,o){var a=this;a.options=o,a.$body=t(document.body),a.$element=t(i),a.$backdrop=a.isShown=null,a.scrollbarWidth=0,o.moveable===e&&(a.options.moveable=a.$element.hasClass("modal-moveable")),o.remote&&a.$element.find(".modal-content").load(o.remote,function(){a.$element.trigger("loaded."+n)}),o.scrollInside&&t(window).on("resize."+n,function(){a.isShown&&a.adjustPosition()})};o.VERSION="3.2.0",o.TRANSITION_DURATION=300,o.BACKDROP_TRANSITION_DURATION=150,o.DEFAULTS={backdrop:!0,keyboard:!0,show:!0,position:"fit"};var a=function(e,i){var n=t(window);i.left=Math.max(0,Math.min(i.left,n.width()-e.outerWidth())),i.top=Math.max(0,Math.min(i.top,n.height()-e.outerHeight())),e.css(i)};o.prototype.toggle=function(t,e){return this.isShown?this.hide():this.show(t,e)},o.prototype.adjustPosition=function(i){var o=this,s=o.options;if(i===e&&(i=s.position),i!==e&&null!==i){t.isFunction(i)&&(i=i(o));var r=o.$element.find(".modal-dialog"),l=t(window).height(),d={maxHeight:"initial",overflow:"visible"},c=r.find(".modal-body").css(d);if(s.scrollInside&&c.length){var p=s.headerHeight,u=s.footerHeight,h=r.find(".modal-header"),f=r.find(".modal-footer");p="number"!=typeof p&&h.length?h.outerHeight():t.isFunction(p)?p(h):0,u="number"!=typeof u&&f.length?f.outerHeight():t.isFunction(u)?u(f):0,d.maxHeight=l-p-u,d.overflow=c[0].scrollHeight>d.maxHeight?"auto":"visible",c.css(d)}var g=Math.max(0,(l-r.outerHeight())/2);if("fit"===i?i={top:g>50?Math.floor(2*g/3):g}:"center"===i?i={top:g}:t.isPlainObject(i)||(i={top:i}),r.hasClass("modal-moveable")){var m=null,v=s.rememberPos;v&&(v===!0?m=o.$element.data("modal-pos"):t.zui.store&&(m=t.zui.store.pageGet(n+".rememberPos."+v))),i=t.extend(i,{left:Math.max(0,(t(window).width()-r.outerWidth())/2)},m),"inside"===s.moveable?a(r,i):r.css(i)}else r.css(i)}},o.prototype.setMoveable=function(){t.fn.draggable||console.error("Moveable modal requires draggable.js.");var e=this,i=e.options,o=e.$element.find(".modal-dialog").removeClass("modal-dragged");o.toggleClass("modal-moveable",!!i.moveable),e.$element.data("modal-moveable-setup")||o.draggable({container:e.$element,handle:".modal-header",before:function(){var t=o.css("margin-top");t&&"0px"!==t&&o.css("top",t).css("margin-top","").addClass("modal-dragged")},finish:function(o){var a=i.rememberPos;a&&(e.$element.data("modal-pos",o.pos),t.zui.store&&a!==!0&&t.zui.store.pageSet(n+".rememberPos."+a,o.pos))},move:"inside"!==i.moveable||function(t){a(o,t)}})},o.prototype.show=function(e,i){var a=this,s=t.Event("show."+n,{relatedTarget:e});a.$element.trigger(s),a.$element.toggleClass("modal-scroll-inside",!!a.options.scrollInside),a.isShown||s.isDefaultPrevented()||(a.isShown=!0,a.options.moveable&&a.setMoveable(),a.options.backdrop!==!1&&(a.$body.addClass("modal-open"),a.setScrollbar()),a.escape(),a.$element.on("click.dismiss."+n,'[data-dismiss="modal"]',function(t){a.hide(),t.stopPropagation()}),a.backdrop(function(){var s=t.support.transition&&a.$element.hasClass("fade");a.$element.parent().length||a.$element.appendTo(a.$body),a.$element.show().scrollTop(0),s&&a.$element[0].offsetWidth,a.$element.addClass("in").attr("aria-hidden",!1),a.adjustPosition(i),a.enforceFocus();var r=t.Event("shown."+n,{relatedTarget:e});s?a.$element.find(".modal-dialog").one("bsTransitionEnd",function(){a.$element.trigger("focus").trigger(r)}).emulateTransitionEnd(o.TRANSITION_DURATION):a.$element.trigger("focus").trigger(r)}))},o.prototype.hide=function(e){e&&e.preventDefault&&e.preventDefault();var i=this;e=t.Event("hide."+n),i.$element.trigger(e),i.isShown&&!e.isDefaultPrevented()&&(i.isShown=!1,i.options.backdrop!==!1&&(i.$body.removeClass("modal-open"),i.resetScrollbar()),i.escape(),t(document).off("focusin."+n),i.$element.removeClass("in").attr("aria-hidden",!0).off("click.dismiss."+n),t.support.transition&&i.$element.hasClass("fade")?i.$element.one("bsTransitionEnd",t.proxy(i.hideModal,i)).emulateTransitionEnd(o.TRANSITION_DURATION):i.hideModal())},o.prototype.enforceFocus=function(){t(document).off("focusin."+n).on("focusin."+n,t.proxy(function(t){this.$element[0]===t.target||this.$element.has(t.target).length||this.$element.trigger("focus")},this))},o.prototype.escape=function(){this.isShown&&this.options.keyboard?t(document).on("keydown.dismiss."+n,t.proxy(function(i){if(27==i.which){var o=t.Event("escaping."+n),a=this.$element.triggerHandler(o,"esc");if(a!=e&&!a)return;this.hide()}},this)):this.isShown||t(document).off("keydown.dismiss."+n)},o.prototype.hideModal=function(){var t=this;this.$element.hide(),this.backdrop(function(){t.$element.trigger("hidden."+n)})},o.prototype.removeBackdrop=function(){this.$backdrop&&this.$backdrop.remove(),this.$backdrop=null},o.prototype.backdrop=function(e){var i=this,a=this.$element.hasClass("fade")?"fade":"";if(this.isShown&&this.options.backdrop){var s=t.support.transition&&a;if(this.$backdrop=t('<div class="modal-backdrop '+a+'" />').appendTo(this.$body),this.$element.on("mousedown.dismiss."+n,t.proxy(function(t){t.target===t.currentTarget&&("static"==this.options.backdrop?this.$element[0].focus.call(this.$element[0]):this.hide.call(this))},this)),s&&this.$backdrop[0].offsetWidth,this.$backdrop.addClass("in"),!e)return;s?this.$backdrop.one("bsTransitionEnd",e).emulateTransitionEnd(o.BACKDROP_TRANSITION_DURATION):e()}else if(!this.isShown&&this.$backdrop){this.$backdrop.removeClass("in");var r=function(){i.removeBackdrop(),e&&e()};t.support.transition&&this.$element.hasClass("fade")?this.$backdrop.one("bsTransitionEnd",r).emulateTransitionEnd(o.BACKDROP_TRANSITION_DURATION):r()}else e&&e()},o.prototype.setScrollbar=function(){t.zui.fixBodyScrollbar()&&this.options.onSetScrollbar&&this.options.onSetScrollbar(paddingRight)},o.prototype.resetScrollbar=function(){t.zui.resetBodyScrollbar(),this.options.onSetScrollbar&&this.options.onSetScrollbar("")},o.prototype.measureScrollbar=function(){var t=document.createElement("div");t.className="modal-scrollbar-measure",this.$body.append(t);var e=t.offsetWidth-t.clientWidth;return this.$body[0].removeChild(t),e};var s=t.fn.modal;t.fn.modal=i,t.fn.modal.Constructor=o,t.fn.modal.noConflict=function(){return t.fn.modal=s,this},t(document).on("click."+n+".data-api",'[data-toggle="modal"]',function(e){var o=t(this),a=o.attr("href"),s=null;try{s=t(o.attr("data-target")||a&&a.replace(/.*(?=#[^\s]+$)/,""))}catch(r){return}if(s.length){var l=s.data(n)?"toggle":t.extend({remote:!/#/.test(a)&&a},s.data(),o.data());o.is("a")&&e.preventDefault(),s.one("show."+n,function(t){t.isDefaultPrevented()||s.one("hidden."+n,function(){o.is(":visible")&&o.trigger("focus")})}),i.call(s,l,this,o.data("position"))}})}(jQuery,void 0),function(t,e,i){"use strict";if(!t.fn.modal)throw new Error("Modal trigger requires modal.js");var n="zui.modaltrigger",o="ajax",a=".zui.modal",s="string",r=function(e,i){e=t.extend({},r.DEFAULTS,t.ModalTriggerDefaults,i?i.data():null,e),this.isShown,this.$trigger=i,this.options=e,this.id=t.zui.uuid()};r.DEFAULTS={type:"custom",height:"auto",name:"triggerModal",fade:!0,position:"fit",showHeader:!0,delay:0,backdrop:!0,keyboard:!0,waittime:0,loadingIcon:"icon-spinner-indicator",scrollInside:!1},r.prototype.initOptions=function(i){if(i.url&&(!i.type||i.type!=o&&"iframe"!=i.type)&&(i.type=o),i.remote)i.type=o,typeof i.remote===s&&(i.url=i.remote);else if(i.iframe)i.type="iframe",typeof i.iframe===s&&(i.url=i.iframe);else if(i.custom&&(i.type="custom",typeof i.custom===s)){var n;try{n=t(i.custom)}catch(a){}n&&n.length?i.custom=n:t.isFunction(e[i.custom])&&(i.custom=e[i.custom])}return i},r.prototype.init=function(e){var i=this,o=t("#"+e.name);o.length&&(i.isShown||o.off(a),o.remove()),o=t('<div id="'+e.name+'" class="modal modal-trigger '+(e.className||"")+'">'+("string"==typeof e.loadingIcon&&0===e.loadingIcon.indexOf("icon-")?'<div class="icon icon-spin loader '+e.loadingIcon+'"></div>':e.loadingIcon)+'<div class="modal-dialog"><div class="modal-content"><div class="modal-header"><button class="close" data-dismiss="modal">×</button><h4 class="modal-title"><i class="modal-icon"></i> <span class="modal-title-name"></span></h4></div><div class="modal-body"></div></div></div></div>').appendTo("body").data(n,i);var s=function(i,n){var s=e[i];t.isFunction(s)&&o.on(n+a,s)};s("onShow","show"),s("shown","shown"),s("onHide","hide"),s("hidden","hidden"),s("loaded","loaded"),o.on("shown"+a,function(){i.isShown=!0}).on("hidden"+a,function(){i.isShown=!1}),this.$modal=o,this.$dialog=o.find(".modal-dialog"),e.mergeOptions&&(this.options=e)},r.prototype.show=function(i){var a=this,l=t.extend({},r.DEFAULTS,a.options,{url:a.$trigger?a.$trigger.attr("href")||a.$trigger.attr("data-url")||a.$trigger.data("url"):a.options.url},i),d=a.isShown;l=a.initOptions(l),d||a.init(l);var c=a.$modal,p=c.find(".modal-dialog"),u=l.custom,h=p.find(".modal-body").css("padding","").toggleClass("load-indicator loading",!!d),f=p.find(".modal-header"),g=p.find(".modal-content");c.toggleClass("fade",l.fade).addClass(l.className).toggleClass("modal-loading",!d).toggleClass("modal-scroll-inside",!!l.scrollInside),p.toggleClass("modal-md","md"===l.size).toggleClass("modal-sm","sm"===l.size).toggleClass("modal-lg","lg"===l.size).toggleClass("modal-fullscreen","fullscreen"===l.size),f.toggle(l.showHeader),f.find(".modal-icon").attr("class","modal-icon icon-"+l.icon),f.find(".modal-title-name").text(l.title||""),l.size&&"fullscreen"===l.size&&(l.width="",l.height="");var m=function(){clearTimeout(this.resizeTask),this.resizeTask=setTimeout(function(){a.adjustPosition(l.position)},100)},v=function(t,e){return"undefined"==typeof t&&(t=l.delay),setTimeout(function(){p=c.find(".modal-dialog"),l.width&&"auto"!=l.width&&p.css("width",l.width),l.height&&"auto"!=l.height&&(p.css("height",l.height),"iframe"===l.type&&h.css("height",p.height()-f.outerHeight())),a.adjustPosition(l.position),c.removeClass("modal-loading").removeClass("modal-updating"),d&&h.removeClass("loading"),"iframe"!=l.type&&(h=p.off("resize."+n).find(".modal-body").off("resize."+n),(h.length?h:p).on("resize."+n,m)),e&&e()},t)};if("custom"===l.type&&u)if(t.isFunction(u)){var y=u({modal:c,options:l,modalTrigger:a,ready:v});typeof y===s&&(h.html(y),v())}else u instanceof t?(h.html(t("<div>").append(u.clone()).html()),v()):(h.html(u),v());else if(l.url){var b=function(){var t=c.callComEvent(a,"broken");"string"==typeof t&&h.html(t),v()};if(c.attr("ref",l.url),"iframe"===l.type){c.addClass("modal-iframe"),this.firstLoad=!0;var w="iframe-"+l.name;f.detach(),h.detach(),g.empty().append(f).append(h),h.css("padding",0).html('<iframe id="'+w+'" name="'+w+'" src="'+l.url+'" frameborder="no"  allowfullscreen="true" mozallowfullscreen="true" webkitallowfullscreen="true"  allowtransparency="true" scrolling="auto" style="width: 100%; height: 100%; left: 0px;"></iframe>'),l.waittime>0&&(a.waitTimeout=v(l.waittime,b));var C=document.getElementById(w);C.onload=C.onreadystatechange=function(i){var o=!!l.scrollInside;if(a.firstLoad&&c.addClass("modal-loading"),!this.readyState||"complete"==this.readyState){a.firstLoad=!1,l.waittime>0&&clearTimeout(a.waitTimeout);try{c.attr("ref",C.contentWindow.location.href);var s=e.frames[w].$;if(s&&"auto"===l.height&&"fullscreen"!=l.size){var r=s("body").addClass("body-modal").toggleClass("body-modal-scroll-inside",o);l.iframeBodyClass&&r.addClass(l.iframeBodyClass);var d=[],p=function(i){c.removeClass("fade");var n=r.outerHeight();if(i===!0&&l.onlyIncreaseHeight&&(n=Math.max(n,h.data("minModalHeight")||0),h.data("minModalHeight",n)),o){var a=l.headerHeight;"number"!=typeof a?a=f.outerHeight():t.isFunction(a)&&(a=a(f));var s=t(e).height();n=Math.min(n,s-a)}for(d.length>1&&n===d[0]&&(n=Math.max(n,d[1])),d.push(n);d.length>2;)d.shift();h.css("height",n),l.fade&&c.addClass("fade"),v()};c.callComEvent(a,"loaded",{modalType:"iframe",jQuery:s}),setTimeout(p,100),r.off("resize."+n).on("resize."+n,p),o&&t(e).off("resize."+n).on("resize."+n,p)}else v();var u=l.handleLinkInIframe;u&&s("body").on("click","string"==typeof u?u:"a[href]",function(){t(this).is('[data-toggle="modal"]')||c.addClass("modal-updating")}),l.iframeStyle&&s("head").append("<style>"+l.iframeStyle+"</style>")}catch(i){v()}}}}else t.ajax(t.extend({url:l.url,success:function(i){try{var s=t(i);s.filter(".modal-dialog").length?p.parent().empty().append(s):s.filter(".modal-content").length?p.find(".modal-content").replaceWith(s):h.wrapInner(s)}catch(r){e.console&&e.console.warn&&console.warn("ZUI: Cannot recogernize remote content.",{error:r,data:i}),c.html(i)}c.callComEvent(a,"loaded",{modalType:o}),v(),l.scrollInside&&t(e).off("resize."+n).on("resize."+n,m)},error:b},l.ajaxOptions))}d||c.modal({show:"show",backdrop:l.backdrop,moveable:l.moveable,rememberPos:l.rememberPos,keyboard:l.keyboard,scrollInside:l.scrollInside})},r.prototype.close=function(i,n){var o=this;(i||n)&&o.$modal.on("hidden"+a,function(){t.isFunction(i)&&i(),typeof n===s&&n.length&&!o.$modal.data("cancel-reload")&&("this"===n?e.location.reload():e.location=n)}),o.$modal.modal("hide")},r.prototype.toggle=function(t){this.isShown?this.close():this.show(t)},r.prototype.adjustPosition=function(e){e=e===i?this.options.position:e,t.isFunction(e)&&(e=e(this)),this.$modal.modal("adjustPosition",e)},t.zui({ModalTrigger:r,modalTrigger:new r}),t.fn.modalTrigger=function(e,i){return t(this).each(function(){var o=t(this),a=o.data(n),l=t.extend({title:o.attr("title")||o.text(),url:o.attr("href"),type:o.hasClass("iframe")?"iframe":""},o.data(),t.isPlainObject(e)&&e);a||o.data(n,a=new r(l,o)),typeof e==s?a[e](i):l.show&&a.show(i),o.on((l.trigger||"click")+".toggle."+n,function(e){l=t.extend(l,{url:o.attr("href")||o.attr("data-url")||o.data("url")||l.url}),a.toggle(l),o.is("a")&&e.preventDefault()})})};var l=t.fn.modal;t.fn.modal=function(e,i){return t(this).each(function(){var n=t(this);n.hasClass("modal")?l.call(n,e,i):n.modalTrigger(e,i)})},t.fn.modal.bs=l;var d=function(e){return e?e=t(e):(e=t(".modal.modal-trigger"),!e.length),e&&e instanceof t?e:null},c=function(i,o,a){var s=i;if(t.isFunction(i)){var r=a;a=o,o=i,i=r}i=d(i),i&&i.length?i.each(function(){t(this).data(n).close(o,a)}):t("body").hasClass("modal-open")||t(".modal.in").length||t("body").hasClass("body-modal")&&e.parent.$.zui.closeModal(s,o,a)},p=function(t,e){e=d(e),e&&e.length&&e.modal("adjustPosition",t)},u=function(e,i){"string"==typeof e&&(e={url:e});var o=d(i);o&&o.length&&o.each(function(){t(this).data(n).show(e)})};t.zui({reloadModal:u,closeModal:c,ajustModalPosition:p,adjustModalPosition:p}),t(document).on("click."+n+".data-api",'[data-toggle="modal"]',function(e){var i=t(this),o=i.attr("href"),a=null;try{a=t(i.attr("data-target")||o&&o.replace(/.*(?=#[^\s]+$)/,""))}catch(s){}a&&a.length||(i.data(n)?i.trigger(".toggle."+n):i.modalTrigger({show:!0})),i.is("a")&&e.preventDefault()}).on("click."+n+".data-api",'[data-dismiss="modal"]',function(){t.zui.closeModal()})}(window.jQuery,window,void 0),+function(t){"use strict";var e=function(t,e){this.type=null,this.options=null,this.enabled=null,this.timeout=null,this.hoverState=null,this.$element=null,this.init("tooltip",t,e)};e.DEFAULTS={animation:!0,placement:"top",selector:!1,template:'<div class="tooltip"><div class="tooltip-arrow"></div><div class="tooltip-inner"></div></div>',trigger:"hover focus",title:"",delay:0,html:!1,container:!1},e.prototype.init=function(e,i,n){this.enabled=!0,this.type=e,this.$element=t(i),this.options=this.getOptions(n);for(var o=this.options.trigger.split(" "),a=o.length;a--;){var s=o[a];if("click"==s)this.$element.on("click."+this.type,this.options.selector,t.proxy(this.toggle,this));else if("manual"!=s){var r="hover"==s?"mouseenter":"focus",l="hover"==s?"mouseleave":"blur";this.$element.on(r+"."+this.type,this.options.selector,t.proxy(this.enter,this)),this.$element.on(l+"."+this.type,this.options.selector,t.proxy(this.leave,this))}}this.options.selector?this._options=t.extend({},this.options,{trigger:"manual",selector:""}):this.fixTitle()},e.prototype.getDefaults=function(){return e.DEFAULTS},e.prototype.getOptions=function(e){return e=t.extend({},this.getDefaults(),this.$element.data(),e),e.delay&&"number"==typeof e.delay&&(e.delay={show:e.delay,hide:e.delay}),e},e.prototype.getDelegateOptions=function(){var e={},i=this.getDefaults();return this._options&&t.each(this._options,function(t,n){i[t]!=n&&(e[t]=n)}),e},e.prototype.enter=function(e){var i=e instanceof this.constructor?e:t(e.currentTarget)[this.type](this.getDelegateOptions()).data("zui."+this.type);return clearTimeout(i.timeout),i.hoverState="in",i.options.delay&&i.options.delay.show?void(i.timeout=setTimeout(function(){"in"==i.hoverState&&i.show()},i.options.delay.show)):i.show()},e.prototype.leave=function(e){var i=e instanceof this.constructor?e:t(e.currentTarget)[this.type](this.getDelegateOptions()).data("zui."+this.type);return clearTimeout(i.timeout),i.hoverState="out",i.options.delay&&i.options.delay.hide?void(i.timeout=setTimeout(function(){"out"==i.hoverState&&i.hide()},i.options.delay.hide)):i.hide()},e.prototype.show=function(e){var i=t.Event("show.zui."+this.type);if((e||this.hasContent())&&this.enabled){var n=this;if(n.$element.trigger(i),i.isDefaultPrevented())return;var o=n.tip();n.setContent(e),n.options.animation&&o.addClass("fade");var a="function"==typeof n.options.placement?n.options.placement.call(n,o[0],n.$element[0]):n.options.placement,s=/\s?auto?\s?/i,r=s.test(a);r&&(a=a.replace(s,"")||"top"),o.detach().css({top:0,left:0,display:"block"}).addClass(a),n.options.container?o.appendTo(n.options.container):o.insertAfter(n.$element);var l=n.getPosition(),d=o[0].offsetWidth,c=o[0].offsetHeight;if(r){var p=n.$element.parent(),u=a,h=document.documentElement.scrollTop||document.body.scrollTop,f="body"==n.options.container?window.innerWidth:p.outerWidth(),g="body"==n.options.container?window.innerHeight:p.outerHeight(),m="body"==n.options.container?0:p.offset().left;a="bottom"==a&&l.top+l.height+c-h>g?"top":"top"==a&&l.top-h-c<0?"bottom":"right"==a&&l.right+d>f?"left":"left"==a&&l.left-d<m?"right":a,o.removeClass(u).addClass(a)}var v=n.getCalculatedOffset(a,l,d,c);n.applyPlacement(v,a);var y=function(){var t=n.hoverState;n.$element.trigger("shown.zui."+n.type),n.hoverState=null,"out"==t&&n.leave(n)};t.support.transition&&n.$tip.hasClass("fade")?o.one("bsTransitionEnd",y).emulateTransitionEnd(150):y()}},e.prototype.applyPlacement=function(t,e){var i,n=this.tip(),o=n[0].offsetWidth,a=n[0].offsetHeight,s=parseInt(n.css("margin-top"),10),r=parseInt(n.css("margin-left"),10);isNaN(s)&&(s=0),isNaN(r)&&(r=0),t.top=t.top+s,t.left=t.left+r,n.offset(t).addClass("in");var l=n[0].offsetWidth,d=n[0].offsetHeight;if("top"==e&&d!=a&&(i=!0,t.top=t.top+a-d),/bottom|top/.test(e)){var c=0;t.left<0&&(c=t.left*-2,t.left=0,n.offset(t),l=n[0].offsetWidth,d=n[0].offsetHeight),this.replaceArrow(c-o+l,l,"left")}else this.replaceArrow(d-a,d,"top");i&&n.offset(t)},e.prototype.replaceArrow=function(t,e,i){this.arrow().css(i,t?50*(1-t/e)+"%":"")},e.prototype.setContent=function(t){var e=this.tip(),i=t||this.getTitle();this.options.tipId&&e.attr("id",this.options.tipId),this.options.tipClass&&e.addClass(this.options.tipClass),
e.find(".tooltip-inner")[this.options.html?"html":"text"](i),e.removeClass("fade in top bottom left right")},e.prototype.hide=function(){function e(){"in"!=i.hoverState&&n.detach()}var i=this,n=this.tip(),o=t.Event("hide.zui."+this.type);if(this.$element.trigger(o),!o.isDefaultPrevented())return n.removeClass("in"),t.support.transition&&this.$tip.hasClass("fade")?n.one(t.support.transition.end,e).emulateTransitionEnd(150):e(),this.$element.trigger("hidden.zui."+this.type),this},e.prototype.fixTitle=function(){var t=this.$element;(t.attr("title")||"string"!=typeof t.attr("data-original-title"))&&t.attr("data-original-title",t.attr("title")||"").attr("title","")},e.prototype.hasContent=function(){return this.getTitle()},e.prototype.getPosition=function(){var e=this.$element[0];return t.extend({},"function"==typeof e.getBoundingClientRect?e.getBoundingClientRect():{width:e.offsetWidth,height:e.offsetHeight},this.$element.offset())},e.prototype.getCalculatedOffset=function(t,e,i,n){return"bottom"==t?{top:e.top+e.height,left:e.left+e.width/2-i/2}:"top"==t?{top:e.top-n,left:e.left+e.width/2-i/2}:"left"==t?{top:e.top+e.height/2-n/2,left:e.left-i}:{top:e.top+e.height/2-n/2,left:e.left+e.width}},e.prototype.getTitle=function(){var t,e=this.$element,i=this.options;return t=e.attr("data-original-title")||("function"==typeof i.title?i.title.call(e[0]):i.title)},e.prototype.tip=function(){return this.$tip=this.$tip||t(this.options.template)},e.prototype.arrow=function(){return this.$arrow=this.$arrow||this.tip().find(".tooltip-arrow")},e.prototype.validate=function(){this.$element[0].parentNode||(this.hide(),this.$element=null,this.options=null)},e.prototype.enable=function(){this.enabled=!0},e.prototype.disable=function(){this.enabled=!1},e.prototype.toggleEnabled=function(){this.enabled=!this.enabled},e.prototype.toggle=function(e){var i=e?t(e.currentTarget)[this.type](this.getDelegateOptions()).data("zui."+this.type):this;i.tip().hasClass("in")?i.leave(i):i.enter(i)},e.prototype.destroy=function(){this.hide().$element.off("."+this.type).removeData("zui."+this.type)};var i=t.fn.tooltip;t.fn.tooltip=function(i,n){return this.each(function(){var o=t(this),a=o.data("zui.tooltip"),s="object"==typeof i&&i;a||o.data("zui.tooltip",a=new e(this,s)),"string"==typeof i&&a[i](n)})},t.fn.tooltip.Constructor=e,t.fn.tooltip.noConflict=function(){return t.fn.tooltip=i,this}}(window.jQuery),+function(t){"use strict";var e=function(t,e){this.init("popover",t,e)};if(!t.fn.tooltip)throw new Error("Popover requires tooltip.js");e.DEFAULTS=t.extend({},t.fn.tooltip.Constructor.DEFAULTS,{placement:"right",trigger:"click",content:"",template:'<div class="popover"><div class="arrow"></div><h3 class="popover-title"></h3><div class="popover-content"></div></div>'}),e.prototype=t.extend({},t.fn.tooltip.Constructor.prototype),e.prototype.constructor=e,e.prototype.getDefaults=function(){return e.DEFAULTS},e.prototype.setContent=function(){var t=this.tip(),e=this.getTarget();if(e)return e.find(".arrow").length<1&&t.addClass("no-arrow"),void t.html(e.html());var i=this.getTitle(),n=this.getContent();t.find(".popover-title")[this.options.html?"html":"text"](i),t.find(".popover-content")[this.options.html?"html":"text"](n),t.removeClass("fade top bottom left right in"),this.options.tipId&&t.attr("id",this.options.tipId),this.options.tipClass&&t.addClass(this.options.tipClass),t.find(".popover-title").html()||t.find(".popover-title").hide()},e.prototype.hasContent=function(){return this.getTarget()||this.getTitle()||this.getContent()},e.prototype.getContent=function(){var t=this.$element,e=this.options;return t.attr("data-content")||("function"==typeof e.content?e.content.call(t[0]):e.content)},e.prototype.getTarget=function(){var e=this.$element,i=this.options,n=e.attr("data-target")||("function"==typeof i.target?i.target.call(e[0]):i.target);return!!n&&("$next"==n?e.next(".popover"):t(n))},e.prototype.arrow=function(){return this.$arrow=this.$arrow||this.tip().find(".arrow")},e.prototype.tip=function(){return this.$tip||(this.$tip=t(this.options.template)),this.$tip};var i=t.fn.popover;t.fn.popover=function(i){return this.each(function(){var n=t(this),o=n.data("zui.popover"),a="object"==typeof i&&i;o||n.data("zui.popover",o=new e(this,a)),"string"==typeof i&&o[i]()})},t.fn.popover.Constructor=e,t.fn.popover.noConflict=function(){return t.fn.popover=i,this}}(window.jQuery),+function(t){"use strict";function e(){t(o).remove(),t(a).each(function(e){var o=i(t(this));o.hasClass("open")&&(o.trigger(e=t.Event("hide."+n)),e.isDefaultPrevented()||o.removeClass("open").trigger("hidden."+n))})}function i(e){var i=e.attr("data-target");i||(i=e.attr("href"),i=i&&/#/.test(i)&&i.replace(/.*(?=#[^\s]*$)/,""));var n;try{n=i&&t(i)}catch(o){}return n&&n.length?n:e.parent()}var n="zui.dropdown",o=".dropdown-backdrop",a="[data-toggle=dropdown]",s=function(e){t(e).on("click."+n,this.toggle)};s.prototype.toggle=function(o){var a=t(this);if(!a.is(".disabled, :disabled")){var s=i(a),r=s.hasClass("open");if(e(),!r){if("ontouchstart"in document.documentElement&&!s.closest(".navbar-nav").length&&t('<div class="dropdown-backdrop"/>').insertAfter(t(this)).on("click",e),s.trigger(o=t.Event("show."+n)),o.isDefaultPrevented())return;s.toggleClass("open").trigger("shown."+n),a.focus()}return!1}},s.prototype.keydown=function(e){if(/(38|40|27)/.test(e.keyCode)){var n=t(this);if(e.preventDefault(),e.stopPropagation(),!n.is(".disabled, :disabled")){var o=i(n),s=o.hasClass("open");if(!s||s&&27==e.keyCode)return 27==e.which&&o.find(a).focus(),n.click();var r=t("[role=menu] li:not(.divider):visible a",o);if(r.length){var l=r.index(r.filter(":focus"));38==e.keyCode&&l>0&&l--,40==e.keyCode&&l<r.length-1&&l++,~l||(l=0),r.eq(l).focus()}}}};var r=t.fn.dropdown;t.fn.dropdown=function(e){return this.each(function(){var i=t(this),n=i.data("dropdown");n||i.data("dropdown",n=new s(this)),"string"==typeof e&&n[e].call(i)})},t.fn.dropdown.Constructor=s,t.fn.dropdown.noConflict=function(){return t.fn.dropdown=r,this};var l=n+".data-api";t(document).on("click."+l,e).on("click."+l,".dropdown form",function(t){t.stopPropagation()}).on("click."+l,a,s.prototype.toggle).on("keydown."+l,a+", [role=menu]",s.prototype.keydown)}(window.jQuery),function(t,e){"use strict";var i="zui.contextmenu",n={duration:200},o=!1,a={},s="zui-contextmenu-"+t.zui.uuid(),r=0,l=0,d=function(){return t(document).off("mousemove."+i).on("mousemove."+i,function(t){r=t.clientX,l=t.clientY}),a},c=function(e,i){if("string"==typeof e&&(e="seperator"===e||"divider"===e||"-"===e||"|"===e?{type:"seperator"}:{label:e,id:i}),"seperator"===e.type||"divider"===e.type)return t('<li class="divider"></li>');var n=t("<a/>").attr({href:e.url||"###","class":e.className,style:e.style}).data("item",e);return e.html?e.html===!0?n.html(e.label||e.text):n=t(e.html):n.text(e.label||e.text),e.onClick&&n.on("click",e.onClick),t("<li />").toggleClass("disabled",e.disabled===!0).append(n)},p=null,u=function(e,i){"function"==typeof e&&(i=e,e=null),p&&(clearTimeout(p),p=null);var n=t("#"+s);if(n.length){var o=n.data("options");if(!e||o.id===e){var r=function(){n.hide(),o.onHidden&&o.onHidden(),i&&i()};o.onHide&&o.onHide();var l=o.animation;n.removeClass("in"),l?p=setTimeout(r,o.duration):r()}}return a},h=function(d,h,f){t.isPlainObject(d)&&(f=h,h=d,d=h.items),o=!0,h=t.extend({},n,h);var g=h.x,m=h.y;g===e&&(g=(h.event||h).clientX),g===e&&(g=r),m===e&&(m=(h.event||h).clientY),m===e&&(m=l);var v=t("#"+s);v.length||(v=t('<div style="display: none; position: fixed; z-index: 2000;" class="contextmenu" id="'+s+'"><ul class="dropdown-menu contextmenu-menu"></ul></div>').appendTo("body"));var y=v.find(".contextmenu-menu").off("click."+i).on("click."+i,"a",function(e){var i=t(this),n=h.onClickItem&&h.onClickItem(i.data("item"),i,e);n!==!1&&u()}).empty();y.attr("class","dropdown-menu contextmenu-menu"+(h.className?" "+h.className:"")),v.hide().attr("class","contextmenu");var b=h.itemCreator||c,w=typeof d;"string"===w?d=d.split(","):"function"===w&&(d=d(h)),t.each(d,function(t,e){y.append(b(e,t,h))});var C=h.animation,x=h.duration;C===!0&&(h.animation=C="fade"),p&&(clearTimeout(p),p=null);var $=function(){v.addClass("in"),h.onShown&&h.onShown(),f&&f()};h.onShow&&h.onShow(),v.data("options",{animation:C,onHide:h.onHide,onHidden:h.onHidden,id:h.id,duration:x});var T=t(window);return g=Math.max(0,Math.min(g,T.width()-y.outerWidth())),m=Math.max(0,Math.min(m,T.height()-y.outerHeight())),v.css({left:g,top:m}),C?(v.addClass("open").addClass(C).show(),p=setTimeout(function(){$(),o=!1},h.duration)):(v.addClass("open").show(),$(),p=setTimeout(function(){o=!1},200)),a};t(document).on("click",function(e){o||t(e.target).closest(".contextmenu").length||u()}),t.extend(a,{NAME:i,DEFAULTS:n,show:h,hide:u,listenMouse:d}),t.zui({ContextMenu:a});var f=function(e,n){var o=this;o.name=i,o.$=t(e),n=o.options=t.extend({trigger:"contextmenu"},a.DEFAULTS,this.$.data(),n);var s=n.trigger;o.id=t.zui.uuid();var r=function(t){if("mousedown"!==t.type||2===t.button){var e={x:t.clientX,y:t.clientY,event:t};return n.itemsCreator&&(e.items=n.itemsCreator.call(this,t)),o.show(e),t.preventDefault(),t.returnValue=!1,!1}},l=s+"."+i;n.selector?o.$.on(l,n.selector,r):o.$.on(l,r)};f.prototype.destory=function(){that.$.off("."+i)},f.prototype.hide=function(t){a.hide(this.id,t)},f.prototype.show=function(e,i){e=t.extend({},this.options,e),a.show(e,i)},t.fn.contextmenu=function(e){return this.each(function(){var n=t(this),o=n.data(i),a="object"==typeof e&&e;o||n.data(i,o=new f(this,a)),"string"==typeof e&&o[e]()})},t.fn.contextmenu.Constructor=f}(jQuery,void 0),+function(t){"use strict";var e=function(e,i){this.$element=t(e),this.$indicators=this.$element.find(".carousel-indicators"),this.options=i,this.paused=this.sliding=this.interval=this.$active=this.$items=null,"hover"==this.options.pause&&this.$element.on("mouseenter",t.proxy(this.pause,this)).on("mouseleave",t.proxy(this.cycle,this))};e.DEFAULTS={interval:5e3,pause:"hover",wrap:!0,touchable:!0},e.prototype.touchable=function(){function e(e){var e=e||window.event;e.originalEvent&&(e=e.originalEvent);var a=t(this);switch(e.type){case"touchstart":n=e.touches[0].pageX,o=e.touches[0].pageY;break;case"touchend":var s=e.changedTouches[0].pageX-n,r=e.changedTouches[0].pageY-o;if(Math.abs(s)>Math.abs(r))i(a,s),Math.abs(s)>10&&e.preventDefault();else{var l=t(window);t("body,html").animate({scrollTop:l.scrollTop()-r},400)}}}function i(t,e){e>10?a.prev():e<-10&&a.next()}if(this.options.touchable){this.$element.on("touchstart touchmove touchend",e);var n,o,a=this}},e.prototype.cycle=function(e){return e||(this.paused=!1),this.interval&&clearInterval(this.interval),this.options.interval&&!this.paused&&(this.interval=setInterval(t.proxy(this.next,this),this.options.interval)),this},e.prototype.getActiveIndex=function(){return this.$active=this.$element.find(".item.active"),this.$items=this.$active.parent().children(),this.$items.index(this.$active)},e.prototype.to=function(e){var i=this,n=this.getActiveIndex();if(!(e>this.$items.length-1||e<0))return this.sliding?this.$element.one("slid",function(){i.to(e)}):n==e?this.pause().cycle():this.slide(e>n?"next":"prev",t(this.$items[e]))},e.prototype.pause=function(e){return e||(this.paused=!0),this.$element.find(".next, .prev").length&&t.support.transition.end&&(this.$element.trigger(t.support.transition.end),this.cycle(!0)),this.interval=clearInterval(this.interval),this},e.prototype.next=function(){if(!this.sliding)return this.slide("next")},e.prototype.prev=function(){if(!this.sliding)return this.slide("prev")},e.prototype.slide=function(e,i){var n=this.$element.find(".item.active"),o=i||n[e](),a=this.interval,s="next"==e?"left":"right",r="next"==e?"first":"last",l=this;if(!o.length){if(!this.options.wrap)return;o=this.$element.find(".item")[r]()}this.sliding=!0,a&&this.pause();var d=t.Event("slide.zui.carousel",{relatedTarget:o[0],direction:s});if(!o.hasClass("active")){if(this.$indicators.length&&(this.$indicators.find(".active").removeClass("active"),this.$element.one("slid",function(){var e=t(l.$indicators.children()[l.getActiveIndex()]);e&&e.addClass("active")})),t.support.transition&&this.$element.hasClass("slide")){if(this.$element.trigger(d),d.isDefaultPrevented())return;o.addClass(e),o[0].offsetWidth,n.addClass(s),o.addClass(s),n.one(t.support.transition.end,function(){o.removeClass([e,s].join(" ")).addClass("active"),n.removeClass(["active",s].join(" ")),l.sliding=!1,setTimeout(function(){l.$element.trigger("slid")},0)}).emulateTransitionEnd(600)}else{if(this.$element.trigger(d),d.isDefaultPrevented())return;n.removeClass("active"),o.addClass("active"),this.sliding=!1,this.$element.trigger("slid")}return a&&this.cycle(),this}};var i=t.fn.carousel;t.fn.carousel=function(i){return this.each(function(){var n=t(this),o=n.data("zui.carousel"),a=t.extend({},e.DEFAULTS,n.data(),"object"==typeof i&&i),s="string"==typeof i?i:a.slide;o||n.data("zui.carousel",o=new e(this,a)),"number"==typeof i?o.to(i):s?o[s]():a.interval&&o.pause().cycle(),a.touchable&&o.touchable()})},t.fn.carousel.Constructor=e,t.fn.carousel.noConflict=function(){return t.fn.carousel=i,this},t(document).on("click.zui.carousel.data-api","[data-slide], [data-slide-to]",function(e){var i,n=t(this),o=t(n.attr("data-target")||(i=n.attr("href"))&&i.replace(/.*(?=#[^\s]+$)/,"")),a=t.extend({},o.data(),n.data()),s=n.attr("data-slide-to");s&&(a.interval=!1),o.carousel(a),(s=n.attr("data-slide-to"))&&o.data("zui.carousel").to(s),e.preventDefault()}),t(window).on("load",function(){t('[data-ride="carousel"]').each(function(){var e=t(this);e.carousel(e.data())})})}(window.jQuery),/*! TangBin: image.ready.js http://www.planeart.cn/?p=1121 */
function(t){"use strict";t.zui.imgReady=function(){var t=[],e=null,i=function(){for(var e=0;e<t.length;e++)t[e].end?t.splice(e--,1):t[e]();!t.length&&n()},n=function(){clearInterval(e),e=null};return function(n,o,a,s){var r,l,d,c,p,u=new Image;return u.src=n,u.complete?(o.call(u),void(a&&a.call(u))):(l=u.width,d=u.height,u.onerror=function(){s&&s.call(u),r.end=!0,u=u.onload=u.onerror=null},r=function(){c=u.width,p=u.height,(c!==l||p!==d||c*p>1024)&&(o.call(u),r.end=!0)},r(),u.onload=function(){!r.end&&r(),a&&a.call(u),u=u.onload=u.onerror=null},void(r.end||(t.push(r),null===e&&(e=setInterval(i,40)))))}}()}(jQuery),function(t,e,i){"use strict";if(!t.fn.modalTrigger)throw new Error("modal & modalTrigger requires for lightbox");if(!t.zui.imgReady)throw new Error("imgReady requires for lightbox");var n=function(e,i){this.$=t(e),this.options=this.getOptions(i),this.init()};n.DEFAULTS={modalTeamplate:'<div class="icon-spinner icon-spin loader"></div><div class="modal-dialog"><button class="close" data-dismiss="modal" aria-hidden="true"><i class="icon-remove"></i></button><button class="controller prev"><i class="icon icon-chevron-left"></i></button><button class="controller next"><i class="icon icon-chevron-right"></i></button><img class="lightbox-img" src="{image}" alt="" data-dismiss="modal" /><div class="caption"><div class="content">{caption}<div></div></div>'},n.prototype.getOptions=function(e){var i="image";return e=t.extend({},n.DEFAULTS,this.$.data(),e),e[i]||(e[i]=this.$.attr("src")||this.$.attr("href")||this.$.find("img").attr("src"),this.$.data(i,e[i])),e},n.prototype.init=function(){this.bindEvents()},n.prototype.initGroups=function(){var e=this.$.data("groups");e||(e=t('[data-toggle="lightbox"][data-group="'+this.options.group+'"], [data-lightbox-group="'+this.options.group+'"]'),this.$.data("groups",e),e.each(function(e){t(this).attr("data-group-index",e)})),this.groups=e,this.groupIndex=parseInt(this.$.data("group-index"))},n.prototype.setImage=function(t,e){void 0!==t&&(this.options.image=t),void 0!==e&&(this.options.caption=e)},n.prototype.show=function(t,e){this.setImage(t,e),this.$.triggerHandler("click")},n.prototype.bindEvents=function(){var n=this.$,o=this,a=this.options;return!!a.image&&void n.modalTrigger({type:"custom",name:"lightboxModal",position:"center",custom:function(n){o.initGroups();var s=n.modal,r=o.groups,l=o.groupIndex;s.addClass("modal-lightbox").html(a.modalTeamplate.format(a)).toggleClass("lightbox-with-caption","string"==typeof a.caption).removeClass("lightbox-full").data("group-index",l);var d=s.find(".modal-dialog"),c=t(e).width();t.zui.imgReady(a.image,function(){d.css({width:i.min(c,this.width)}),c<this.width+30&&s.addClass("lightbox-full"),n.ready(200)}),s.find(".prev").toggleClass("show",r.filter('[data-group-index="'+(l-1)+'"]').length>0),s.find(".next").toggleClass("show",r.filter('[data-group-index="'+(l+1)+'"]').length>0),s.find(".controller").click(function(){var o=t(this),a=s.data("group-index")+(o.hasClass("prev")?-1:1),l=r.filter('[data-group-index="'+a+'"]');if(l.length){var p=l.data("image"),u=l.data("caption");s.addClass("modal-loading").data("group-index",a).toggleClass("lightbox-with-caption","string"==typeof u).removeClass("lightbox-full"),s.find(".lightbox-img").attr("src",p),s.find(".caption > .content").text(u),c=t(e).width(),t.zui.imgReady(p,function(){d.css({width:i.min(c,this.width)}),c<this.width+30&&s.addClass("lightbox-full"),n.ready()})}return s.find(".prev").toggleClass("show",r.filter('[data-group-index="'+(a-1)+'"]').length>0),s.find(".next").toggleClass("show",r.filter('[data-group-index="'+(a+1)+'"]').length>0),!1})}})},t.fn.lightbox=function(e){var i="group"+(new Date).getTime();return this.each(function(){var o=t(this),a="object"==typeof e&&e;"object"==typeof a&&a.group?o.attr("data-lightbox-group",a.group):o.data("group")?o.attr("data-lightbox-group",o.data("group")):o.attr("data-lightbox-group",i),o.data("group",o.data("lightbox-group"));var s=o.data("zui.lightbox");s||o.data("zui.lightbox",s=new n(this,a)),"string"==typeof e&&s[e]()})},t.fn.lightbox.Constructor=n,t(function(){t('[data-toggle="lightbox"]').lightbox()})}(jQuery,window,Math),function(t,e,i){"use strict";var n=0,o='<div class="messager messager-{type} {placement}" style="display: none"><div class="messager-content"></div><div class="messager-actions"></div></div>',a={icons:{},type:"default",placement:"top",time:4e3,parent:"body",close:!0,fade:!0,scale:!0},s={},r=function(e,r){t.isPlainObject(e)?r=t.extend({},r,e):e&&(r?r.content=e:r={content:e});var l=this;r=l.options=t.extend({},a,r),l.id=r.id||n++;var d=s[l.id];d&&d.destroy(),s[l.id]=l,l.$=t(o.format(r)).toggleClass("fade",r.fade).toggleClass("scale",r.scale).attr("id","messager-"+l.id),r.cssClass&&l.$.addClass(r.cssClass);var c=!1,p=l.$.find(".messager-actions"),u=function(e){var n=t('<button type="button" class="action action-'+e.name+'"/>');"close"===e.name&&n.addClass("close"),e.html!==i&&n.html(e.html),e.icon!==i&&n.append('<i class="action-icon icon-'+e.icon+'"/>'),e.text!==i&&n.append('<span class="action-text">'+e.text+"</span>"),e.tooltip!==i&&n.attr("title",e.tooltip).tooltip(),n.data("action",e),p.append(n)};r.actions&&t.each(r.actions,function(t,e){e.name===i&&(e.name=t),"close"==e.name&&(c=!0),u(e)}),!c&&r.close&&u({name:"close",html:"&times;"}),l.$.on("click",".action",function(e){var i,n=t(this).data("action");r.onAction&&(i=r.onAction.call(this,n.name,n,l),i===!1)||t.isFunction(n.action)&&(i=n.action.call(this,l),i===!1)||(l.hide(),e.stopPropagation())}),l.$.on("click",function(t){if(r.onAction){var e=r.onAction.call(this,"content",null,l);e===!0&&l.hide()}}),l.$.data("zui.messager",l),r.show&&l.message!==i&&l.show()};r.prototype.update=function(e,i){t.isPlainObject(e)?i=e:e&&(i?i.content=e:i={content:e});var n=this,o=n.options;n.$.removeClass("messager-"+o.type);var a=n.$.find(".messager-content");o.contentClass&&a.removeClass(o.contentClass),i&&(o=t.extend(o,i)),n.$.addClass("messager-"+o.type).toggleClass("messager-notification",!!o.notification),o.contentClass&&a.addClass(o.contentClass);var s=o.title,r=o.icon;if(e=o.content,a.empty(),s){var l=t('<div class="messager-title"></div>');l[o.html?"html":"text"](s),a.append(l)}if(e){var d=t('<div class="messager-text"></div>');d[o.html?"html":"text"](e),a.append(d)}var c=n.$.find(".messager-icon");if(r){var p=t.isPlainObject(r)?r.html:'<i class="icon-'+r+' icon"></i>';c.length?c.html(p):a.before('<div class="messager-icon">'+p+"<div>")}else c.remove();n.$.toggleClass("messager-has-icon",!!r),n.updateTime||o.onUpdate&&o.onUpdate.call(n,o),n.updateTime=Date.now()},r.prototype.show=function(n,o){var a=this,s=this.options;if(t.isFunction(n)){var r=o;o=n,r!==i&&(n=r)}if(a.isShow)return void a.hide(function(){a.show(n,o)});a.hiding&&(clearTimeout(a.hiding),a.hiding=null),a.update(n);var l=s.placement,d=t(s.parent),c=d.children(".messagers-holder."+l);if(c.length||(c=t("<div/>").attr("class","messagers-holder "+l).appendTo(d)),c.append(a.$),"center"===l){var p=t(e).height()-c.height();c.css("top",Math.max(-p,p/2))}return a.$.show().addClass("in"),s.time&&(a.hiding=setTimeout(function(){a.hide()},s.time)),a.isShow=!0,o&&o(),s.onShow&&s.onShow.call(a,s),a},r.prototype.hide=function(t,e){t===!0&&(e=!0,t=null);var i=this,n=i.options;if(i.$.hasClass("in")){i.$.removeClass("in");var o=function(){var o=i.$.parent();i.$.detach(),o.children().length||o.remove(),t&&t(!0),n.onHide&&n.onHide.call(i,e)};e?o():setTimeout(o,200)}else t&&t(!1),n.onHide&&n.onHide.call(i,e);i.isShow=!1},r.prototype.destroy=function(){var t=this;t.hide(function(){t.$.remove(),t.$=null},!0),delete s[t.id]};var l=function(e){if(e===i)t(".messager").each(function(){var e=t(this).data("zui.messager");e&&e.hide&&e.hide(!0)});else{var n=t("#messager-"+e).data("zui.messager");n&&n.hide&&n.hide()}},d=function(e,n){"string"==typeof n&&(n={type:n}),t.isPlainObject(e)&&(n=t.extend({},n,e),e=null),n=t.extend({},n),n.id===i&&l();var o=s[n.id]||new r(e,n);return o.show(),o},c={notification:!0,placement:"bottom-right",time:0,icon:"bell icon-2x"},p=function(e,i,n){var o=t.extend({id:t.zui.uuid()},c),a="string"==typeof e,s="string"==typeof i;return a&&s?n=t.extend(o,n,{title:e,content:i}):a&&t.isPlainObject(i)?n=t.extend(o,n,i,{title:e}):t.isPlainObject(e)?n=t.extend(o,n,i,e):a&&(n=t.extend(o,n,{title:e})),d(n)},u=function(t){return"string"==typeof t?{placement:t}:t},h={show:d,hide:l};r.all=s,r.DEFAULTS=a,r.NOTIFICATION_DEFAULTS=c,t.each({primary:0,success:"ok-sign",info:"info-sign",warning:"warning-sign",danger:"exclamation-sign",important:0,special:0},function(e,i){h[e]=function(n,o){return d(n,t.extend({type:e,icon:r.DEFAULTS.icons[e]||i||null},u(o)))}}),t.zui({Messager:r,showMessager:d,showNotification:p,messager:h})}(jQuery,window,void 0),function(t,e,i,n){"use strict";function o(t){if(t=t.toLowerCase(),t&&c.test(t)){var e;if(4===t.length){var i="#";for(e=1;e<4;e+=1)i+=t.slice(e,e+1).concat(t.slice(e,e+1));t=i}var n=[];for(e=1;e<7;e+=2)n.push(b("0x"+t.slice(e,e+2)));return{r:n[0],g:n[1],b:n[2],a:1}}throw new Error("Wrong hex string! (hex: "+t+")")}function a(e){return typeof e===f&&("transparent"===e.toLowerCase()||m[e.toLowerCase()]||c.test(t.trim(e.toLowerCase())))}function s(t){function e(t){return t=t<0?t+1:t>1?t-1:t,6*t<1?r+(s-r)*t*6:2*t<1?s:3*t<2?r+(s-r)*(2/3-t)*6:r}var i=t.h,n=t.s,o=t.l,a=t.a;i=d(i)%u/u,n=l(d(n)),o=l(d(o)),a=l(d(a));var s=o<=.5?o*(n+1):o+n-o*n,r=2*o-s,c={r:e(i+1/3)*p,g:e(i)*p,b:e(i-1/3)*p,a:a};return c}function r(t,i,n){return v(n)&&(n=0),v(i)&&(i=p),e.min(e.max(t,n),i)}function l(t,e){return r(t,e)}function d(t){return"number"==typeof t?t:parseFloat(t)}var c=/^#([0-9a-fA-f]{3}|[0-9a-fA-f]{6})$/,p=255,u=360,h=100,f="string",g="object",m={aliceblue:"#f0f8ff",antiquewhite:"#faebd7",aqua:"#00ffff",aquamarine:"#7fffd4",azure:"#f0ffff",beige:"#f5f5dc",bisque:"#ffe4c4",black:"#000000",blanchedalmond:"#ffebcd",blue:"#0000ff",blueviolet:"#8a2be2",brown:"#a52a2a",burlywood:"#deb887",cadetblue:"#5f9ea0",chartreuse:"#7fff00",chocolate:"#d2691e",coral:"#ff7f50",cornflowerblue:"#6495ed",cornsilk:"#fff8dc",crimson:"#dc143c",cyan:"#00ffff",darkblue:"#00008b",darkcyan:"#008b8b",darkgoldenrod:"#b8860b",darkgray:"#a9a9a9",darkgreen:"#006400",darkkhaki:"#bdb76b",darkmagenta:"#8b008b",darkolivegreen:"#556b2f",darkorange:"#ff8c00",darkorchid:"#9932cc",darkred:"#8b0000",darksalmon:"#e9967a",darkseagreen:"#8fbc8f",darkslateblue:"#483d8b",darkslategray:"#2f4f4f",darkturquoise:"#00ced1",darkviolet:"#9400d3",deeppink:"#ff1493",deepskyblue:"#00bfff",dimgray:"#696969",dodgerblue:"#1e90ff",firebrick:"#b22222",floralwhite:"#fffaf0",forestgreen:"#228b22",fuchsia:"#ff00ff",gainsboro:"#dcdcdc",ghostwhite:"#f8f8ff",gold:"#ffd700",goldenrod:"#daa520",gray:"#808080",green:"#008000",greenyellow:"#adff2f",honeydew:"#f0fff0",hotpink:"#ff69b4",indianred:"#cd5c5c",indigo:"#4b0082",ivory:"#fffff0",khaki:"#f0e68c",lavender:"#e6e6fa",lavenderblush:"#fff0f5",lawngreen:"#7cfc00",lemonchiffon:"#fffacd",lightblue:"#add8e6",lightcoral:"#f08080",lightcyan:"#e0ffff",lightgoldenrodyellow:"#fafad2",lightgray:"#d3d3d3",lightgreen:"#90ee90",lightpink:"#ffb6c1",lightsalmon:"#ffa07a",lightseagreen:"#20b2aa",lightskyblue:"#87cefa",lightslategray:"#778899",lightsteelblue:"#b0c4de",lightyellow:"#ffffe0",lime:"#00ff00",limegreen:"#32cd32",linen:"#faf0e6",magenta:"#ff00ff",maroon:"#800000",mediumaquamarine:"#66cdaa",mediumblue:"#0000cd",mediumorchid:"#ba55d3",mediumpurple:"#9370db",mediumseagreen:"#3cb371",mediumslateblue:"#7b68ee",mediumspringgreen:"#00fa9a",mediumturquoise:"#48d1cc",mediumvioletred:"#c71585",midnightblue:"#191970",mintcream:"#f5fffa",mistyrose:"#ffe4e1",moccasin:"#ffe4b5",navajowhite:"#ffdead",navy:"#000080",oldlace:"#fdf5e6",olive:"#808000",olivedrab:"#6b8e23",orange:"#ffa500",orangered:"#ff4500",orchid:"#da70d6",palegoldenrod:"#eee8aa",palegreen:"#98fb98",paleturquoise:"#afeeee",palevioletred:"#db7093",papayawhip:"#ffefd5",peachpuff:"#ffdab9",peru:"#cd853f",pink:"#ffc0cb",plum:"#dda0dd",powderblue:"#b0e0e6",purple:"#800080",red:"#ff0000",rosybrown:"#bc8f8f",royalblue:"#4169e1",saddlebrown:"#8b4513",salmon:"#fa8072",sandybrown:"#f4a460",seagreen:"#2e8b57",seashell:"#fff5ee",sienna:"#a0522d",silver:"#c0c0c0",skyblue:"#87ceeb",slateblue:"#6a5acd",slategray:"#708090",snow:"#fffafa",springgreen:"#00ff7f",steelblue:"#4682b4",tan:"#d2b48c",teal:"#008080",thistle:"#d8bfd8",tomato:"#ff6347",turquoise:"#40e0d0",violet:"#ee82ee",wheat:"#f5deb3",white:"#ffffff",whitesmoke:"#f5f5f5",yellow:"#ffff00",yellowgreen:"#9acd32"},v=function(t){return t===n},y=function(t){return!v(t)},b=function(t){return parseInt(t)},w=function(t){return b(l(d(t),p))},C=function(t,e,i,n){var a=this;if(a.r=a.g=a.b=0,a.a=1,y(n)&&(a.a=l(d(n),1)),y(t)&&y(e)&&y(i))a.r=w(t),a.g=w(e),a.b=w(i);else if(y(t)){var r=typeof t;if(r==f)if(t=t.toLowerCase(),"transparent"===t)a.a=0;else if(m[t])a.rgb(o(m[t]));else if(0===t.indexOf("rgb")){var c=t.substring(t.indexOf("(")+1,t.lastIndexOf(")")).split(",",4);a.rgb({r:c[0],g:c[1],b:c[2],a:c[3]})}else a.rgb(o(t));else if("number"==r&&v(e))a.r=a.g=a.b=w(t);else if(r==g&&y(t.r))a.r=w(t.r),y(t.g)&&(a.g=w(t.g)),y(t.b)&&(a.b=w(t.b)),y(t.a)&&(a.a=l(d(t.a),1));else if(r==g&&y(t.h)){var p={h:l(d(t.h),u),s:1,l:1,a:1};y(t.s)&&(p.s=l(d(t.s),1)),y(t.l)&&(p.l=l(d(t.l),1)),y(t.a)&&(p.a=l(d(t.a),1)),a.rgb(s(p))}}};C.prototype.rgb=function(t){var e=this;if(y(t)){if(typeof t==g)y(t.r)&&(e.r=w(t.r)),y(t.g)&&(e.g=w(t.g)),y(t.b)&&(e.b=w(t.b)),y(t.a)&&(e.a=l(d(t.a),1));else{var i=b(d(t));e.r=i,e.g=i,e.b=i}return e}return{r:e.r,g:e.g,b:e.b,a:e.a}},C.prototype.hue=function(t){var e=this,i=e.toHsl();return v(t)?i.h:(i.h=l(d(t),u),e.rgb(s(i)),e)},C.prototype.darken=function(t){var e=this,i=e.toHsl();return i.l-=t/h,i.l=l(i.l,1),e.rgb(s(i)),e},C.prototype.clone=function(){var t=this;return new C(t.r,t.g,t.b,t.a)},C.prototype.lighten=function(t){return this.darken(-t)},C.prototype.fade=function(t){return this.a=l(t/h,1),this},C.prototype.spin=function(t){var e=this.toHsl(),i=(e.h+t)%u;return e.h=i<0?u+i:i,this.rgb(s(e))},C.prototype.toHsl=function(){var t,i,n=this,o=n.r/p,a=n.g/p,s=n.b/p,r=n.a,l=e.max(o,a,s),d=e.min(o,a,s),c=(l+d)/2,h=l-d;if(l===d)t=i=0;else{switch(i=c>.5?h/(2-l-d):h/(l+d),l){case o:t=(a-s)/h+(a<s?6:0);break;case a:t=(s-o)/h+2;break;case s:t=(o-a)/h+4}t/=6}return{h:t*u,s:i,l:c,a:r}},C.prototype.luma=function(){var t=this.r/p,i=this.g/p,n=this.b/p;return t=t<=.03928?t/12.92:e.pow((t+.055)/1.055,2.4),i=i<=.03928?i/12.92:e.pow((i+.055)/1.055,2.4),n=n<=.03928?n/12.92:e.pow((n+.055)/1.055,2.4),.2126*t+.7152*i+.0722*n},C.prototype.saturate=function(t){var e=this.toHsl();return e.s+=t/h,e.s=l(e.s),this.rgb(s(e))},C.prototype.desaturate=function(t){return this.saturate(-t)},C.prototype.contrast=function(t,e,i){if(e=v(e)?new C(p,p,p,1):new C(e),t=v(t)?new C(0,0,0,1):new C(t),t.luma()>e.luma()){var n=e;e=t,t=n}return this.a<.5?t:(i=v(i)?.43:d(i),this.luma()<i?e:t)},C.prototype.hexStr=function(){var t=this.r.toString(16),e=this.g.toString(16),i=this.b.toString(16);return 1==t.length&&(t="0"+t),1==e.length&&(e="0"+e),1==i.length&&(i="0"+i),"#"+t+e+i},C.prototype.toCssStr=function(){var t=this;return t.a>0?t.a<1?"rgba("+t.r+","+t.g+","+t.b+","+t.a+")":t.hexStr():"transparent"},C.isColor=a,C.names=m,C.get=function(t){return new C(t)},t.zui({Color:C})}(jQuery,Math,window,void 0),function(t){"use strict";function e(e,i){if(e===!1)return e;if(!e)return i;e===!0?e={add:!0,"delete":!0,edit:!0,sort:!0}:"string"==typeof e&&(e=e.split(","));var n;return t.isArray(e)&&(n={},t.each(e,function(e,i){t.isPlainObject(i)?n[i.action]=i:n[i]=!0}),e=n),t.isPlainObject(e)&&(n={},t.each(e,function(e,i){i?n[e]=t.extend({type:e},s[e],t.isPlainObject(i)?i:null):n[e]=!1}),e=n),i?t.extend(!0,{},i,e):e}function i(e,i,n){return i=i||e.type,t(n||e.template).addClass("tree-action").attr(t.extend({"data-type":i,title:e.title||""},e.attr)).data("action",e)}var n="zui.tree",o=0,a=function(e,i){this.name=n,this.$=t(e),this.getOptions(i),this._init()},s={sort:{template:'<a class="sort-handler" href="javascript:;"><i class="icon icon-move"></i></a>'},add:{template:'<a href="javascript:;"><i class="icon icon-plus"></i></a>'},edit:{template:'<a href="javascript:;"><i class="icon icon-pencil"></i></a>'},"delete":{template:'<a href="javascript:;"><i class="icon icon-trash"></i></a>'}};a.DEFAULTS={animate:null,initialState:"normal",toggleTemplate:'<i class="list-toggle icon"></i>'},a.prototype.add=function(e,i,n,o,a){var s,r=t(e),l=this.options;if(r.is("li")?(s=r.children("ul"),s.length||(s=t("<ul/>"),r.append(s),this._initList(s,r))):s=r,s){var d=this;t.isArray(i)||(i=[i]),t.each(i,function(e,i){var n=t("<li/>").data(i).appendTo(s);void 0!==i.id&&n.attr("data-id",i.id);var o=l.itemWrapper?t(l.itemWrapper===!0?'<div class="tree-item-wrapper"/>':l.itemWrapper).appendTo(n):n;if(i.html)o.html(i.html);else if(t.isFunction(d.options.itemCreator)){var a=d.options.itemCreator(n,i);a!==!0&&a!==!1&&o.html(a)}else i.url?o.append(t("<a/>",{href:i.url}).text(i.title||i.name)):o.append(t("<span/>").text(i.title||i.name));d._initItem(n,i.idx||e,s,i),i.children&&i.children.length&&d.add(n,i.children)}),this._initList(s),n&&!s.hasClass("tree")&&d.expand(s.parent("li"),o,a)}},a.prototype.reload=function(e){var i=this;e&&(i.$.empty(),i.add(i.$,e)),i.isPreserve&&i.store.time&&i.$.find("li:not(.tree-action-item)").each(function(){var e=t(this);i[i.store[e.data("id")]?"expand":"collapse"](e,!0,!0)})},a.prototype._initList=function(n,o,a,s){var r=this;n.hasClass("tree")?(a=0,o=null):(o=(o||n.closest("li")).addClass("has-list"),o.find(".list-toggle").length||o.prepend(this.options.toggleTemplate),a=a||o.data("idx")),n.removeClass("has-active-item");var l=n.attr("data-idx",a||0).children("li:not(.tree-action-item)").each(function(e){r._initItem(t(this),e+1,n)});1!==l.length||l.find("ul").length||l.addClass("tree-single-item"),s=s||(o?o.data():null);var d=e(s?s.actions:null,this.actions);if(d){if(d.add&&d.add.templateInList!==!1){var c=n.children("li.tree-action-item");c.length?c.detach().appendTo(n):t('<li class="tree-action-item"/>').append(i(d.add,"add",d.add.templateInList)).appendTo(n)}d.sort&&n.sortable(t.extend({dragCssClass:"tree-drag-holder",trigger:".sort-handler",selector:"li:not(.tree-action-item)",finish:function(t){r.callEvent("action",{action:d.sort,$list:n,target:t.target,item:s})}},d.sort.options,t.isPlainObject(this.options.sortable)?this.options.sortable:null))}o&&(o.hasClass("open")||s&&s.open)&&o.addClass("open in")},a.prototype._initItem=function(n,o,a,s){if(void 0===o){var r=n.prev("li");o=r.length?r.data("idx")+1:1}if(a=a||n.closest("ul"),n.attr("data-idx",o).removeClass("tree-single-item"),!n.data("id")){var l=o;a.hasClass("tree")||(l=a.parent("li").data("id")+"-"+l),n.attr("data-id",l)}n.hasClass("active")&&a.parent("li").addClass("has-active-item"),s=s||n.data();var d=e(s.actions,this.actions);if(d){var c=n.find(".tree-actions");c.length||(c=t('<div class="tree-actions"/>').appendTo(this.options.itemWrapper?n.find(".tree-item-wrapper"):n),t.each(d,function(t,e){e&&c.append(i(e,t))}))}var p=n.children("ul");p.length&&this._initList(p,n,o,s)},a.prototype._init=function(){var i=this.options,a=this;this.actions=e(i.actions),this.$.addClass("tree"),i.animate&&this.$.addClass("tree-animate"),this._initList(this.$);var s=i.initialState,r=t.zui&&t.zui.store&&t.zui.store.enable;r&&(this.selector=n+"::"+(i.name||"")+"#"+(this.$.attr("id")||o++),this.store=t.zui.store[i.name?"get":"pageGet"](this.selector,{})),"preserve"===s&&(r?this.isPreserve=!0:this.options.initialState=s="normal"),this.reload(i.data),r&&(this.isPreserve=!0),"expand"===s?this.expand():"collapse"===s?this.collapse():"active"===s&&this.expandSelect(".active"),this.$.on("click",'.list-toggle,a[href="#"],.tree-toggle',function(e){var i=t(this),n=i.parent("li");a.callEvent("hit",{target:n,item:n.data()}),a.toggle(n),i.is("a")&&e.preventDefault()}).on("click",".tree-action",function(){var e=t(this),i=e.data();if(i.action&&(i=i.action),"sort"!==i.type){var n=e.closest("li:not(.tree-action-item)");a.callEvent("action",{action:i,target:this,$item:n,item:n.data()})}})},a.prototype.preserve=function(e,i,n){if(this.isPreserve)if(e)i=i||e.data("id"),n=void 0===n&&e.hasClass("open"),n?this.store[i]=n:delete this.store[i],this.store.time=(new Date).getTime(),t.zui.store[this.options.name?"set":"pageSet"](this.selector,this.store);else{var o=this;this.store={},this.$.find("li").each(function(){o.preserve(t(this))})}},a.prototype.expandSelect=function(t){this.show(t,!0)},a.prototype.expand=function(t,e,i){t?(t.addClass("open"),!e&&this.options.animate?setTimeout(function(){t.addClass("in")},10):t.addClass("in")):t=this.$.find("li.has-list").addClass("open in"),i||this.preserve(t),this.callEvent("expand",t,this)},a.prototype.show=function(e,i,n){var o=this;e instanceof t||(e=o.$.find("li").filter(e)),e.each(function(){var e=t(this);if(o.expand(e,i,n),e)for(var a=e.parent("ul");a&&a.length&&!a.hasClass("tree");){var s=a.parent("li");s.length?(o.expand(s,i,n),a=s.parent("ul")):a=!1}})},a.prototype.collapse=function(t,e,i){t?!e&&this.options.animate?(t.removeClass("in"),setTimeout(function(){t.removeClass("open")},300)):t.removeClass("open in"):t=this.$.find("li.has-list").removeClass("open in"),i||this.preserve(t),this.callEvent("collapse",t,this)},a.prototype.toggle=function(t){var e=t&&t.hasClass("open")||t===!1||void 0===t&&this.$.find("li.has-list.open").length;this[e?"collapse":"expand"](t)},a.prototype.getOptions=function(e){this.options=t.extend({},a.DEFAULTS,this.$.data(),e),null===this.options.animate&&this.$.hasClass("tree-animate")&&(this.options.animate=!0)},a.prototype.toData=function(e,i){t.isFunction(e)&&(i=e,e=null),e=e||this.$;var n=this;return e.children("li:not(.tree-action-item)").map(function(){var e=t(this),o=e.data();delete o["zui.droppable"];var a=e.children("ul");return a.length&&(o.children=n.toData(a)),t.isFunction(i)?i(o,e):o}).get()},a.prototype.callEvent=function(e,i){var n;return t.isFunction(this.options[e])&&(n=this.options[e](i,this)),this.$.trigger(t.Event(e+"."+this.name,i)),n},t.fn.tree=function(e,i){return this.each(function(){var o=t(this),s=o.data(n),r="object"==typeof e&&e;s||o.data(n,s=new a(this,r)),"string"==typeof e&&s[e](i)})},t.fn.tree.Constructor=a,t(function(){t('[data-ride="tree"]').tree()})}(jQuery),
/*!
 * Chart.js 1.0.2
 * Copyright 2015 Nick Downie
 * Released under the MIT license
 * http://chartjs.org/
 */

function(V) {
	var ak = V && V.zui ? V.zui : this,
		ag = (ak.Chart, function(b) {
			this.canvas = b.canvas, this.ctx = b;
			var c = function(f, g) {
					return f["offset" + g] ? f["offset" + g] : document.defaultView.getComputedStyle(f).getPropertyValue(g)
				},
				a = this.width = c(b.canvas, "Width"),
				d = this.height = c(b.canvas, "Height");
			b.canvas.width = a, b.canvas.height = d;
			var a = this.width = b.canvas.width,
				d = this.height = b.canvas.height;
			return this.aspectRatio = this.width / this.height, ac.retinaScale(this), this
		});
	ag.defaults = {
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
	}, ag.types = {};
	var ac = ag.helpers = {},
		ab = ac.each = function(b, d, a) {
			var g = Array.prototype.slice.call(arguments, 3);
			if (b) {
				if (b.length === +b.length) {
					var f;
					for (f = 0; f < b.length; f++) {
						d.apply(a, [b[f], f].concat(g))
					}
				} else {
					for (var c in b) {
						d.apply(a, [b[c], c].concat(g))
					}
				}
			}
		},
		X = ac.clone = function(a) {
			var b = {};
			return ab(a, function(c, d) {
				a.hasOwnProperty(d) && (b[d] = c)
			}), b
		},
		ao = ac.extend = function(a) {
			return ab(Array.prototype.slice.call(arguments, 1), function(b) {
				ab(b, function(c, d) {
					b.hasOwnProperty(d) && (a[d] = c)
				})
			}), a
		},
		Z = ac.merge = function(b, c) {
			var a = Array.prototype.slice.call(arguments, 0);
			return a.unshift({}), ao.apply(null, a)
		},
		ae = ac.indexOf = function(b, c) {
			if (Array.prototype.indexOf) {
				return b.indexOf(c)
			}
			for (var a = 0; a < b.length; a++) {
				if (b[a] === c) {
					return a
				}
			}
			return -1
		},
		ah = (ac.where = function(b, c) {
			var a = [];
			return ac.each(b, function(d) {
				c(d) && a.push(d)
			}), a
		}, ac.findNextWhere = function(b, c, a) {
			a || (a = -1);
			for (var f = a + 1; f < b.length; f++) {
				var d = b[f];
				if (c(d)) {
					return d
				}
			}
		}, ac.findPreviousWhere = function(b, c, a) {
			a || (a = b.length);
			for (var f = a - 1; f >= 0; f--) {
				var d = b[f];
				if (c(d)) {
					return d
				}
			}
		}, ac.inherits = function(b) {
			var c = this,
				a = b && b.hasOwnProperty("constructor") ? b.constructor : function() {
					return c.apply(this, arguments)
				},
				d = function() {
					this.constructor = a
				};
			return d.prototype = c.prototype, a.prototype = new d, a.extend = ah, b && ao(a.prototype, b), a.__super__ = c.prototype, a
		}),
		am = ac.noop = function() {},
		al = ac.uid = function() {
			var a = 0;
			return function() {
				return "chart-" + a++
			}
		}(),
		U = ac.warn = function(a) {
			window.console && "function" == typeof window.console.warn && console.warn(a)
		},
		aa = ac.amd = "function" == typeof define && define.amd,
		aj = ac.isNumber = function(a) {
			return !isNaN(parseFloat(a)) && isFinite(a)
		},
		ai = ac.max = function(a) {
			return Math.max.apply(Math, a)
		},
		ad = ac.min = function(a) {
			return Math.min.apply(Math, a)
		},
		Q = (ac.cap = function(b, c, a) {
			if (aj(c)) {
				if (b > c) {
					return c
				}
			} else {
				if (aj(a) && b < a) {
					return a
				}
			}
			return b
		}, ac.getDecimalPlaces = function(a) {
			return a % 1 !== 0 && aj(a) ? a.toString().split(".")[1].length : 0
		}),
		J = ac.radians = function(a) {
			return a * (Math.PI / 180)
		},
		an = (ac.getAngleFromPoint = function(b, d) {
			var a = d.x - b.x,
				g = d.y - b.y,
				f = Math.sqrt(a * a + g * g),
				c = 2 * Math.PI + Math.atan2(g, a);
			return a < 0 && g < 0 && (c += 2 * Math.PI), {
				angle: c,
				distance: f
			}
		}, ac.aliasPixel = function(a) {
			return a % 2 === 0 ? 0 : 0.5
		}),
		O = (ac.splineCurve = function(d, h, c, l) {
			var k = Math.sqrt(Math.pow(h.x - d.x, 2) + Math.pow(h.y - d.y, 2)),
				f = Math.sqrt(Math.pow(c.x - h.x, 2) + Math.pow(c.y - h.y, 2)),
				b = l * k / (k + f),
				g = l * f / (k + f);
			return {
				inner: {
					x: h.x - b * (c.x - d.x),
					y: h.y - b * (c.y - d.y)
				},
				outer: {
					x: h.x + g * (c.x - d.x),
					y: h.y + g * (c.y - d.y)
				}
			}
		}, ac.calculateOrderOfMagnitude = function(a) {
			return Math.floor(Math.log(a) / Math.LN10)
		}),
		N = (ac.calculateScaleRange = function(T, D, x, m, k) {
			var aq = 2,
				M = Math.floor(D / (1.5 * x)),
				b = aq >= M,
				w = ai(T),
				z = ad(T);
			w === z && (w += 0.5, z >= 0.5 && !m ? z -= 0.5 : w += 0.5);
			for (var I = Math.abs(w - z), F = O(I), S = Math.ceil(w / (1 * Math.pow(10, F))) * Math.pow(10, F), g = m ? 0 : Math.floor(z / (1 * Math.pow(10, F))) * Math.pow(10, F), C = S - g, P = Math.pow(10, F), L = Math.round(C / P);
			(L > M || 2 * L < M) && !b;) {
				if (L > M) {
					P *= 2, L = Math.round(C / P), L % 1 !== 0 && (b = !0)
				} else {
					if (k && F >= 0) {
						if (P / 2 % 1 !== 0) {
							break
						}
						P /= 2, L = Math.round(C / P)
					} else {
						P /= 2, L = Math.round(C / P)
					}
				}
			}
			return b && (L = aq, P = C / L), {
				steps: L,
				stepValue: P,
				min: g,
				max: g + L * P
			}
		}, ac.template = function(b, c) {
			function a(g, h) {
				var f = /\W/.test(g) ? new Function("obj", "var p=[],print=function(){p.push.apply(p,arguments);};with(obj){p.push('" + g.replace(/[\r\t\n]/g, " ").split("<%").join("\t").replace(/((^|%>)[^\t]*)'/g, "$1\r").replace(/\t=(.*?)%>/g, "',$1,'").split("\t").join("');").split("%>").join("p.push('").split("\r").join("\\'") + "');}return p.join('');") : d[g] = d[g];
				return h ? f(h) : f
			}
			if (b instanceof Function) {
				return b(c)
			}
			var d = {};
			return a(b, c)
		}),
		Y = (ac.generateLabels = function(b, d, a, f) {
			var c = new Array(d);
			return labelTemplateString && ab(c, function(g, h) {
				c[h] = N(b, {
					value: a + f * (h + 1)
				})
			}), c
		}, ac.easingEffects = {
			linear: function(a) {
				return a
			},
			easeInQuad: function(a) {
				return a * a
			},
			easeOutQuad: function(a) {
				return -1 * a * (a - 2)
			},
			easeInOutQuad: function(a) {
				return (a /= 0.5) < 1 ? 0.5 * a * a : -0.5 * (--a * (a - 2) - 1)
			},
			easeInCubic: function(a) {
				return a * a * a
			},
			easeOutCubic: function(a) {
				return 1 * ((a = a / 1 - 1) * a * a + 1)
			},
			easeInOutCubic: function(a) {
				return (a /= 0.5) < 1 ? 0.5 * a * a * a : 0.5 * ((a -= 2) * a * a + 2)
			},
			easeInQuart: function(a) {
				return a * a * a * a
			},
			easeOutQuart: function(a) {
				return -1 * ((a = a / 1 - 1) * a * a * a - 1)
			},
			easeInOutQuart: function(a) {
				return (a /= 0.5) < 1 ? 0.5 * a * a * a * a : -0.5 * ((a -= 2) * a * a * a - 2)
			},
			easeInQuint: function(a) {
				return 1 * (a /= 1) * a * a * a * a
			},
			easeOutQuint: function(a) {
				return 1 * ((a = a / 1 - 1) * a * a * a * a + 1)
			},
			easeInOutQuint: function(a) {
				return (a /= 0.5) < 1 ? 0.5 * a * a * a * a * a : 0.5 * ((a -= 2) * a * a * a * a + 2)
			},
			easeInSine: function(a) {
				return -1 * Math.cos(a / 1 * (Math.PI / 2)) + 1
			},
			easeOutSine: function(a) {
				return 1 * Math.sin(a / 1 * (Math.PI / 2))
			},
			easeInOutSine: function(a) {
				return -0.5 * (Math.cos(Math.PI * a / 1) - 1)
			},
			easeInExpo: function(a) {
				return 0 === a ? 1 : 1 * Math.pow(2, 10 * (a / 1 - 1))
			},
			easeOutExpo: function(a) {
				return 1 === a ? 1 : 1 * (-Math.pow(2, -10 * a / 1) + 1)
			},
			easeInOutExpo: function(a) {
				return 0 === a ? 0 : 1 === a ? 1 : (a /= 0.5) < 1 ? 0.5 * Math.pow(2, 10 * (a - 1)) : 0.5 * (-Math.pow(2, -10 * --a) + 2)
			},
			easeInCirc: function(a) {
				return a >= 1 ? a : -1 * (Math.sqrt(1 - (a /= 1) * a) - 1)
			},
			easeOutCirc: function(a) {
				return 1 * Math.sqrt(1 - (a = a / 1 - 1) * a)
			},
			easeInOutCirc: function(a) {
				return (a /= 0.5) < 1 ? -0.5 * (Math.sqrt(1 - a * a) - 1) : 0.5 * (Math.sqrt(1 - (a -= 2) * a) + 1)
			},
			easeInElastic: function(b) {
				var c = 1.70158,
					a = 0,
					d = 1;
				return 0 === b ? 0 : 1 == (b /= 1) ? 1 : (a || (a = 0.3), d < Math.abs(1) ? (d = 1, c = a / 4) : c = a / (2 * Math.PI) * Math.asin(1 / d), -(d * Math.pow(2, 10 * (b -= 1)) * Math.sin((1 * b - c) * (2 * Math.PI) / a)))
			},
			easeOutElastic: function(b) {
				var c = 1.70158,
					a = 0,
					d = 1;
				return 0 === b ? 0 : 1 == (b /= 1) ? 1 : (a || (a = 0.3), d < Math.abs(1) ? (d = 1, c = a / 4) : c = a / (2 * Math.PI) * Math.asin(1 / d), d * Math.pow(2, -10 * b) * Math.sin((1 * b - c) * (2 * Math.PI) / a) + 1)
			},
			easeInOutElastic: function(b) {
				var c = 1.70158,
					a = 0,
					d = 1;
				return 0 === b ? 0 : 2 == (b /= 0.5) ? 1 : (a || (a = 1 * (0.3 * 1.5)), d < Math.abs(1) ? (d = 1, c = a / 4) : c = a / (2 * Math.PI) * Math.asin(1 / d), b < 1 ? -0.5 * (d * Math.pow(2, 10 * (b -= 1)) * Math.sin((1 * b - c) * (2 * Math.PI) / a)) : d * Math.pow(2, -10 * (b -= 1)) * Math.sin((1 * b - c) * (2 * Math.PI) / a) * 0.5 + 1)
			},
			easeInBack: function(a) {
				var b = 1.70158;
				return 1 * (a /= 1) * a * ((b + 1) * a - b)
			},
			easeOutBack: function(a) {
				var b = 1.70158;
				return 1 * ((a = a / 1 - 1) * a * ((b + 1) * a + b) + 1)
			},
			easeInOutBack: function(a) {
				var b = 1.70158;
				return (a /= 0.5) < 1 ? 0.5 * (a * a * (((b *= 1.525) + 1) * a - b)) : 0.5 * ((a -= 2) * a * (((b *= 1.525) + 1) * a + b) + 2)
			},
			easeInBounce: function(a) {
				return 1 - Y.easeOutBounce(1 - a)
			},
			easeOutBounce: function(a) {
				return (a /= 1) < 1 / 2.75 ? 1 * (7.5625 * a * a) : a < 2 / 2.75 ? 1 * (7.5625 * (a -= 1.5 / 2.75) * a + 0.75) : a < 2.5 / 2.75 ? 1 * (7.5625 * (a -= 2.25 / 2.75) * a + 0.9375) : 1 * (7.5625 * (a -= 2.625 / 2.75) * a + 0.984375)
			},
			easeInOutBounce: function(a) {
				return a < 0.5 ? 0.5 * Y.easeInBounce(2 * a) : 0.5 * Y.easeOutBounce(2 * a - 1) + 0.5
			}
		}),
		ap = ac.requestAnimFrame = function() {
			return window.requestAnimationFrame || window.webkitRequestAnimationFrame || window.mozRequestAnimationFrame || window.oRequestAnimationFrame || window.msRequestAnimationFrame ||
			function(a) {
				return window.setTimeout(a, 1000 / 60)
			}
		}(),
		af = ac.cancelAnimFrame = function() {
			return window.cancelAnimationFrame || window.webkitCancelAnimationFrame || window.mozCancelAnimationFrame || window.oCancelAnimationFrame || window.msCancelAnimationFrame ||
			function(a) {
				return window.clearTimeout(a, 1000 / 60)
			}
		}(),
		j = (ac.animationLoop = function(m, h, g, d, c, p) {
			var k = 0,
				b = Y[g] || Y.linear,
				f = function() {
					k++;
					var a = k / h,
						e = b(a);
					m.call(p, e, a, k), d.call(p, e, a), k < h ? p.animationFrame = ap(f) : c.apply(p)
				};
			ap(f)
		}, ac.getRelativePosition = function(b) {
			var d, a, g = b.originalEvent || b,
				f = b.currentTarget || b.srcElement,
				c = f.getBoundingClientRect();
			return g.touches ? (d = g.touches[0].clientX - c.left, a = g.touches[0].clientY - c.top) : (d = g.clientX - c.left, a = g.clientY - c.top), {
				x: d,
				y: a
			}
		}, ac.addEvent = function(b, c, a) {
			b.addEventListener ? b.addEventListener(c, a) : b.attachEvent ? b.attachEvent("on" + c, a) : b["on" + c] = a
		}),
		A = ac.removeEvent = function(b, c, a) {
			b.removeEventListener ? b.removeEventListener(c, a, !1) : b.detachEvent ? b.detachEvent("on" + c, a) : b["on" + c] = am
		},
		W = (ac.bindEvents = function(b, c, a) {
			b.events || (b.events = {}), ab(c, function(d) {
				b.events[d] = function() {
					a.apply(b, arguments)
				}, j(b.chart.canvas, d, b.events[d])
			})
		}, ac.unbindEvents = function(a, b) {
			ab(b, function(d, c) {
				A(a.chart.canvas, c, d)
			})
		}),
		E = ac.getMaximumWidth = function(a) {
			var b = a.parentNode;
			return b.clientWidth
		},
		B = ac.getMaximumHeight = function(a) {
			var b = a.parentNode;
			return b.clientHeight
		},
		R = (ac.getMaximumSize = ac.getMaximumWidth, ac.retinaScale = function(b) {
			var c = b.ctx,
				a = b.canvas.width,
				d = b.canvas.height;
			window.devicePixelRatio && (c.canvas.style.width = a + "px", c.canvas.style.height = d + "px", c.canvas.height = d * window.devicePixelRatio, c.canvas.width = a * window.devicePixelRatio, c.scale(window.devicePixelRatio, window.devicePixelRatio))
		}),
		G = ac.clear = function(a) {
			a.ctx.clearRect(0, 0, a.width, a.height)
		},
		H = ac.fontString = function(b, c, a) {
			return c + " " + b + "px " + a
		},
		K = ac.longestText = function(b, c, a) {
			b.font = c;
			var d = 0;
			return ab(a, function(g) {
				var f = b.measureText(g).width;
				d = f > d ? f : d
			}), d
		},
		q = ac.drawRoundedRectangle = function(b, d, a, g, f, c) {
			b.beginPath(), b.moveTo(d + c, a), b.lineTo(d + g - c, a), b.quadraticCurveTo(d + g, a, d + g, a + c), b.lineTo(d + g, a + f - c), b.quadraticCurveTo(d + g, a + f, d + g - c, a + f), b.lineTo(d + c, a + f), b.quadraticCurveTo(d, a + f, d, a + f - c), b.lineTo(d, a + c), b.quadraticCurveTo(d, a, d + c, a), b.closePath()
		};
	ag.instances = {}, ag.Type = function(a, b, c) {
		this.options = b, this.chart = c, this.id = al(), ag.instances[this.id] = this, b.responsive && this.resize(), this.initialize.call(this, a)
	}, ao(ag.Type.prototype, {
		initialize: function() {
			return this
		},
		clear: function() {
			return G(this.chart), this
		},
		stop: function() {
			return af(this.animationFrame), this
		},
		resize: function(b) {
			this.stop();
			var c = this.chart.canvas,
				a = E(this.chart.canvas),
				d = this.options.maintainAspectRatio ? a / this.chart.aspectRatio : B(this.chart.canvas);
			return c.width = this.chart.width = a, c.height = this.chart.height = d, R(this.chart), "function" == typeof b && b.apply(this, Array.prototype.slice.call(arguments, 1)), this
		},
		reflow: am,
		render: function(a) {
			return a && this.reflow(), this.options.animation && !a ? ac.animationLoop(this.draw, this.options.animationSteps, this.options.animationEasing, this.options.onAnimationProgress, this.options.onAnimationComplete, this) : (this.draw(), this.options.onAnimationComplete.call(this)), this
		},
		generateLegend: function() {
			return N(this.options.legendTemplate, this)
		},
		destroy: function() {
			this.clear(), W(this, this.events);
			var a = this.chart.canvas;
			a.width = this.chart.width, a.height = this.chart.height, a.style.removeProperty ? (a.style.removeProperty("width"), a.style.removeProperty("height")) : (a.style.removeAttribute("width"), a.style.removeAttribute("height")), delete ag.instances[this.id]
		},
		showTooltip: function(o, g) {
			"undefined" == typeof this.activeElements && (this.activeElements = []);
			var p = function(a) {
					var c = !1;
					return a.length !== this.activeElements.length ? c = !0 : (ab(a, function(e, d) {
						e !== this.activeElements[d] && (c = !0)
					}, this), c)
				}.call(this, o);
			if (p || g) {
				if (this.activeElements = o, this.draw(), this.options.customTooltips && this.options.customTooltips(!1), o.length > 0) {
					if (this.datasets && this.datasets.length > 1) {
						for (var m, b, f = this.datasets.length - 1; f >= 0 && (m = this.datasets[f].points || this.datasets[f].bars || this.datasets[f].segments, b = ae(m, o[0]), b === -1); f--) {}
						var l = [],
							k = [],
							n = function(z) {
								var w, r, c, C, x, d = [],
									v = [],
									y = [];
								return ac.each(this.datasets, function(a) {
									a.showTooltips !== !1 && (w = a.points || a.bars || a.segments, w[b] && w[b].hasValue() && d.push(w[b]))
								}), ac.each(d, function(a) {
									v.push(a.x), y.push(a.y), l.push(ac.template(this.options.multiTooltipTemplate, a)), k.push({
										fill: a._saved.fillColor || a.fillColor,
										stroke: a._saved.strokeColor || a.strokeColor
									})
								}, this), x = ad(y), c = ai(y), C = ad(v), r = ai(v), {
									x: C > this.chart.width / 2 ? C : r,
									y: (x + c) / 2
								}
							}.call(this, b);
						new ag.MultiTooltip({
							x: n.x,
							y: n.y,
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
							labels: l,
							legendColors: k,
							legendColorBackground: this.options.multiTooltipKeyBackground,
							title: N(this.options.multiTooltipTitleTemplate, o[0]),
							chart: this.chart,
							ctx: this.chart.ctx,
							custom: this.options.customTooltips
						}).draw()
					} else {
						ab(o, function(a) {
							var c = a.tooltipPosition();
							new ag.Tooltip({
								x: Math.round(c.x),
								y: Math.round(c.y),
								xPadding: this.options.tooltipXPadding,
								yPadding: this.options.tooltipYPadding,
								fillColor: this.options.tooltipFillColor,
								textColor: this.options.tooltipFontColor,
								fontFamily: this.options.tooltipFontFamily,
								fontStyle: this.options.tooltipFontStyle,
								fontSize: this.options.tooltipFontSize,
								caretHeight: this.options.tooltipCaretSize,
								cornerRadius: this.options.tooltipCornerRadius,
								text: N(this.options.tooltipTemplate, a),
								chart: this.chart,
								custom: this.options.customTooltips
							}).draw()
						}, this)
					}
				}
				return this
			}
		},
		toBase64Image: function() {
			return this.chart.canvas.toDataURL.apply(this.chart.canvas, arguments)
		}
	}), ag.Type.extend = function(b) {
		var c = this,
			f = function() {
				return c.apply(this, arguments)
			};
		if (f.prototype = X(c.prototype), ao(f.prototype, b), f.extend = ag.Type.extend, b.name || c.prototype.name) {
			var d = b.name || c.prototype.name,
				a = ag.defaults[c.prototype.name] ? X(ag.defaults[c.prototype.name]) : {};
			ag.defaults[d] = ao(a, b.defaults), ag.types[d] = f, ag.prototype[d] = function(g, k) {
				var h = Z(ag.defaults.global, ag.defaults[d], k || {});
				return new f(g, h, this)
			}
		} else {
			U("Name not provided for this chart, so it hasn't been registered")
		}
		return c
	}, ag.Element = function(a) {
		ao(this, a), this.initialize.apply(this, arguments), this.save()
	}, ao(ag.Element.prototype, {
		initialize: function() {},
		restore: function(a) {
			return a ? ab(a, function(b) {
				this[b] = this._saved[b]
			}, this) : ao(this, this._saved), this
		},
		save: function() {
			return this._saved = X(this), delete this._saved._saved, this
		},
		update: function(a) {
			return ab(a, function(b, c) {
				this._saved[c] = this[c], this[c] = b
			}, this), this
		},
		transition: function(a, b) {
			return ab(a, function(d, c) {
				this[c] = (d - this._saved[c]) * b + this._saved[c]
			}, this), this
		},
		tooltipPosition: function() {
			return {
				x: this.x,
				y: this.y
			}
		},
		hasValue: function() {
			return aj(this.value)
		}
	}), ag.Element.extend = ah, ag.Point = ag.Element.extend({
		display: !0,
		inRange: function(b, c) {
			var a = this.hitDetectionRadius + this.radius;
			return Math.pow(b - this.x, 2) + Math.pow(c - this.y, 2) < Math.pow(a, 2)
		},
		draw: function() {
			if (this.display) {
				var a = this.ctx;
				a.beginPath(), a.arc(this.x, this.y, this.radius, 0, 2 * Math.PI), a.closePath(), a.strokeStyle = this.strokeColor, a.lineWidth = this.strokeWidth, a.fillStyle = this.fillColor, a.fill(), a.stroke()
			}
		}
	}), ag.Arc = ag.Element.extend({
		inRange: function(b, d) {
			var a = ac.getAngleFromPoint(this, {
				x: b,
				y: d
			}),
				f = a.angle >= this.startAngle && a.angle <= this.endAngle,
				c = a.distance >= this.innerRadius && a.distance <= this.outerRadius;
			return f && c
		},
		tooltipPosition: function() {
			var a = this.startAngle + (this.endAngle - this.startAngle) / 2,
				b = (this.outerRadius - this.innerRadius) / 2 + this.innerRadius;
			return {
				x: this.x + Math.cos(a) * b,
				y: this.y + Math.sin(a) * b
			}
		},
		draw: function(b) {
			var c = this.ctx;
			if (c.beginPath(), c.arc(this.x, this.y, this.outerRadius, this.startAngle, this.endAngle), c.arc(this.x, this.y, this.innerRadius, this.endAngle, this.startAngle, !0), c.closePath(), c.strokeStyle = this.strokeColor, c.lineWidth = this.strokeWidth, c.fillStyle = this.fillColor, c.fill(), c.lineJoin = "bevel", this.showStroke && c.stroke(), this.circleBeginEnd) {
				var a = (this.outerRadius + this.innerRadius) / 2,
					d = (this.outerRadius - this.innerRadius) / 2;
				c.beginPath(), c.arc(this.x + Math.cos(this.startAngle) * a, this.y + Math.sin(this.startAngle) * a, d, 0, 2 * Math.PI), c.closePath(), c.fill(), c.beginPath(), c.arc(this.x + Math.cos(this.endAngle) * a, this.y + Math.sin(this.endAngle) * a, d, 0, 2 * Math.PI), c.closePath(), c.fill()
			}
		}
	}), ag.Rectangle = ag.Element.extend({
		draw: function() {
			var b = this.ctx,
				d = this.width / 2,
				a = this.x - d,
				g = this.x + d,
				f = this.base - (this.base - this.y),
				c = this.strokeWidth / 2;
			this.showStroke && (a += c, g -= c, f += c), b.beginPath(), b.fillStyle = this.fillColor, b.strokeStyle = this.strokeColor, b.lineWidth = this.strokeWidth, b.moveTo(a, this.base), b.lineTo(a, f), b.lineTo(g, f), b.lineTo(g, this.base), b.fill(), this.showStroke && b.stroke()
		},
		height: function() {
			return this.base - this.y
		},
		inRange: function(a, b) {
			return a >= this.x - this.width / 2 && a <= this.x + this.width / 2 && b >= this.y && b <= this.base
		}
	}), ag.Tooltip = ag.Element.extend({
		draw: function() {
			var d = this.chart.ctx;
			d.font = H(this.fontSize, this.fontStyle, this.fontFamily), this.xAlign = "center", this.yAlign = "above";
			var g = this.caretPadding = 2,
				c = d.measureText(this.text).width + 2 * this.xPadding,
				k = this.fontSize + 2 * this.yPadding,
				h = k + this.caretHeight + g;
			this.x + c / 2 > this.chart.width ? this.xAlign = "left" : this.x - c / 2 < 0 && (this.xAlign = "right"), this.y - h < 0 && (this.yAlign = "below");
			var f = this.x - c / 2,
				b = this.y - h;
			if (d.fillStyle = this.fillColor, this.custom) {
				this.custom(this)
			} else {
				switch (this.yAlign) {
				case "above":
					d.beginPath(), d.moveTo(this.x, this.y - g), d.lineTo(this.x + this.caretHeight, this.y - (g + this.caretHeight)), d.lineTo(this.x - this.caretHeight, this.y - (g + this.caretHeight)), d.closePath(), d.fill();
					break;
				case "below":
					b = this.y + g + this.caretHeight, d.beginPath(), d.moveTo(this.x, this.y + g), d.lineTo(this.x + this.caretHeight, this.y + g + this.caretHeight), d.lineTo(this.x - this.caretHeight, this.y + g + this.caretHeight), d.closePath(), d.fill()
				}
				switch (this.xAlign) {
				case "left":
					f = this.x - c + (this.cornerRadius + this.caretHeight);
					break;
				case "right":
					f = this.x - (this.cornerRadius + this.caretHeight)
				}
				q(d, f, b, c, k, this.cornerRadius), d.fill(), d.fillStyle = this.textColor, d.textAlign = "center", d.textBaseline = "middle", d.fillText(this.text, f + c / 2, b + k / 2)
			}
		}
	}), ag.MultiTooltip = ag.Element.extend({
		initialize: function() {
			this.font = H(this.fontSize, this.fontStyle, this.fontFamily), this.titleFont = H(this.titleFontSize, this.titleFontStyle, this.titleFontFamily), this.height = this.labels.length * this.fontSize + (this.labels.length - 1) * (this.fontSize / 2) + 2 * this.yPadding + 1.5 * this.titleFontSize, this.ctx.font = this.titleFont;
			var b = this.ctx.measureText(this.title).width,
				c = K(this.ctx, this.font, this.labels) + this.fontSize + 3,
				a = ai([c, b]);
			this.width = a + 2 * this.xPadding;
			var d = this.height / 2;
			this.y - d < 0 ? this.y = d : this.y + d > this.chart.height && (this.y = this.chart.height - d), this.x > this.chart.width / 2 ? this.x -= this.xOffset + this.width : this.x += this.xOffset
		},
		getLineHeight: function(b) {
			var c = this.y - this.height / 2 + this.yPadding,
				a = b - 1;
			return 0 === b ? c + this.titleFontSize / 2 : c + (1.5 * this.fontSize * a + this.fontSize / 2) + 1.5 * this.titleFontSize
		},
		draw: function() {
			if (this.custom) {
				this.custom(this)
			} else {
				q(this.ctx, this.x, this.y - this.height / 2, this.width, this.height, this.cornerRadius);
				var a = this.ctx;
				a.fillStyle = this.fillColor, a.fill(), a.closePath(), a.textAlign = "left", a.textBaseline = "middle", a.fillStyle = this.titleTextColor, a.font = this.titleFont, a.fillText(this.title, this.x + this.xPadding, this.getLineHeight(0)), a.font = this.font, ac.each(this.labels, function(c, b) {
					a.fillStyle = this.textColor, a.fillText(c, this.x + this.xPadding + this.fontSize + 3, this.getLineHeight(b + 1)), a.fillStyle = this.legendColorBackground, a.fillRect(this.x + this.xPadding, this.getLineHeight(b + 1) - this.fontSize / 2, this.fontSize, this.fontSize), a.fillStyle = this.legendColors[b].fill, a.fillRect(this.x + this.xPadding, this.getLineHeight(b + 1) - this.fontSize / 2, this.fontSize, this.fontSize)
				}, this)
			}
		}
	}), ag.Scale = ag.Element.extend({
		initialize: function() {
			this.fit()
		},
		buildYLabels: function() {
			this.yLabels = [];
			for (var a = Q(this.stepValue), b = 0; b <= this.steps; b++) {
				this.yLabels.push(N(this.templateString, {
					value: (this.min + b * this.stepValue).toFixed(a)
				}))
			}
			this.yLabelWidth = this.display && this.showLabels ? K(this.ctx, this.font, this.yLabels) : 0
		},
		addXLabel: function(a) {
			this.xLabels.push(a), this.valuesCount++, this.fit()
		},
		removeXLabel: function() {
			this.xLabels.shift(), this.valuesCount--, this.fit()
		},
		fit: function() {
			this.startPoint = this.display ? this.fontSize : 0, this.endPoint = this.display ? this.height - 1.5 * this.fontSize - 5 : this.height, this.startPoint += this.padding, this.endPoint -= this.padding;
			var a, b = this.endPoint - this.startPoint;
			for (this.calculateYRange(b), this.buildYLabels(), this.calculateXLabelRotation(); b > this.endPoint - this.startPoint;) {
				b = this.endPoint - this.startPoint, a = this.yLabelWidth, this.calculateYRange(b), this.buildYLabels(), a < this.yLabelWidth && this.calculateXLabelRotation()
			}
		},
		calculateXLabelRotation: function() {
			this.ctx.font = this.font;
			var d, g, c = this.ctx.measureText(this.xLabels[0]).width,
				k = this.ctx.measureText(this.xLabels[this.xLabels.length - 1]).width;
			if (this.xScalePaddingRight = k / 2 + 3, this.xScalePaddingLeft = c / 2 > this.yLabelWidth + 10 ? c / 2 : this.yLabelWidth + 10, this.xLabelRotation = 0, this.display) {
				var h, f = K(this.ctx, this.font, this.xLabels);
				this.xLabelWidth = f;
				for (var b = Math.floor(this.calculateX(1) - this.calculateX(0)) - 6; this.xLabelWidth > b && 0 === this.xLabelRotation || this.xLabelWidth > b && this.xLabelRotation <= 90 && this.xLabelRotation > 0;) {
					h = Math.cos(J(this.xLabelRotation)), d = h * c, g = h * k, d + this.fontSize / 2 > this.yLabelWidth + 8 && (this.xScalePaddingLeft = d + this.fontSize / 2), this.xScalePaddingRight = this.fontSize / 2, this.xLabelRotation++, this.xLabelWidth = h * f
				}
				this.xLabelRotation > 0 && (this.endPoint -= Math.sin(J(this.xLabelRotation)) * f + 3)
			} else {
				this.xLabelWidth = 0, this.xScalePaddingRight = this.padding, this.xScalePaddingLeft = this.padding
			}
		},
		calculateYRange: am,
		drawingArea: function() {
			return this.startPoint - this.endPoint
		},
		calculateY: function(a) {
			var b = this.drawingArea() / (this.min - this.max);
			return this.endPoint - b * (a - this.min)
		},
		calculateX: function(b) {
			var c = (this.xLabelRotation > 0, this.width - (this.xScalePaddingLeft + this.xScalePaddingRight)),
				a = c / Math.max(this.valuesCount - (this.offsetGridLines ? 0 : 1), 1),
				d = a * b + this.xScalePaddingLeft;
			return this.offsetGridLines && (d += a / 2), Math.round(d)
		},
		update: function(a) {
			ac.extend(this, a), this.fit()
		},
		draw: function() {
			var b = this.ctx,
				d = (this.endPoint - this.startPoint) / this.steps,
				a = Math.round(this.xScalePaddingLeft);
			if (this.display) {
				b.fillStyle = this.textColor, b.font = this.font;
				var c = this.showBeyondLine ? 5 : 0;
				ab(this.yLabels, function(m, f) {
					var k = this.endPoint - d * f,
						e = Math.round(k),
						g = this.showHorizontalLines;
					b.textAlign = "right", b.textBaseline = "middle", this.showLabels && b.fillText(m, a - 10, k), 0 !== f || g || (g = !0), g && b.beginPath(), f > 0 ? (b.lineWidth = this.gridLineWidth, b.strokeStyle = this.gridLineColor) : (b.lineWidth = this.lineWidth, b.strokeStyle = this.lineColor), e += ac.aliasPixel(b.lineWidth), g && (b.moveTo(a, e), b.lineTo(this.width, e), b.stroke(), b.closePath()), b.lineWidth = this.lineWidth, b.strokeStyle = this.lineColor, b.beginPath(), b.moveTo(a - c, e), b.lineTo(a, e), b.stroke(), b.closePath()
				}, this), ab(this.xLabels, function(k, g) {
					var m = this.calculateX(g) + an(this.lineWidth),
						l = this.calculateX(g - (this.offsetGridLines ? 0.5 : 0)) + an(this.lineWidth),
						f = this.xLabelRotation > 0,
						h = this.showVerticalLines;
					0 !== g || h || (h = !0), h && b.beginPath(), g > 0 ? (b.lineWidth = this.gridLineWidth, b.strokeStyle = this.gridLineColor) : (b.lineWidth = this.lineWidth, b.strokeStyle = this.lineColor), h && (b.moveTo(l, this.endPoint), b.lineTo(l, this.startPoint - 3), b.stroke(), b.closePath()), b.lineWidth = this.lineWidth, b.strokeStyle = this.lineColor, b.beginPath(), b.moveTo(l, this.endPoint), b.lineTo(l, this.endPoint + c), b.stroke(), b.closePath(), b.save(), b.translate(m, f ? this.endPoint + 12 : this.endPoint + 8), b.rotate(J(this.xLabelRotation) * -1), b.font = this.font, b.textAlign = f ? "right" : "center", b.textBaseline = f ? "middle" : "top", b.fillText(k, 0, 0), b.restore()
				}, this)
			}
		}
	}), ag.RadialScale = ag.Element.extend({
		initialize: function() {
			this.size = ad([this.height, this.width]), this.drawingArea = this.display ? this.size / 2 - (this.fontSize / 2 + this.backdropPaddingY) : this.size / 2
		},
		calculateCenterOffset: function(a) {
			var b = this.drawingArea / (this.max - this.min);
			return (a - this.min) * b
		},
		update: function() {
			this.lineArc ? this.drawingArea = this.display ? this.size / 2 - (this.fontSize / 2 + this.backdropPaddingY) : this.size / 2 : this.setScaleSize(), this.buildYLabels()
		},
		buildYLabels: function() {
			this.yLabels = [];
			for (var a = Q(this.stepValue), b = 0; b <= this.steps; b++) {
				this.yLabels.push(N(this.templateString, {
					value: (this.min + b * this.stepValue).toFixed(a)
				}))
			}
		},
		getCircumference: function() {
			return 2 * Math.PI / this.valuesCount
		},
		setScaleSize: function() {
			var L, z, w, m, k, M, F, b, v, x, D, C, I = ad([this.height / 2 - this.pointLabelFontSize - 5, this.width / 2]),
				f = this.width,
				y = 0;
			for (this.ctx.font = H(this.pointLabelFontSize, this.pointLabelFontStyle, this.pointLabelFontFamily), z = 0; z < this.valuesCount; z++) {
				L = this.getPointPosition(z, I), w = this.ctx.measureText(N(this.templateString, {
					value: this.labels[z]
				})).width + 5, 0 === z || z === this.valuesCount / 2 ? (m = w / 2, L.x + m > f && (f = L.x + m, k = z), L.x - m < y && (y = L.x - m, F = z)) : z < this.valuesCount / 2 ? L.x + w > f && (f = L.x + w, k = z) : z > this.valuesCount / 2 && L.x - w < y && (y = L.x - w, F = z)
			}
			v = y, x = Math.ceil(f - this.width), M = this.getIndexAngle(k), b = this.getIndexAngle(F), D = x / Math.sin(M + Math.PI / 2), C = v / Math.sin(b + Math.PI / 2), D = aj(D) ? D : 0, C = aj(C) ? C : 0, this.drawingArea = I - (C + D) / 2, this.setCenterPoint(C, D)
		},
		setCenterPoint: function(b, c) {
			var a = this.width - c - this.drawingArea,
				d = b + this.drawingArea;
			this.xCenter = (d + a) / 2, this.yCenter = this.height / 2
		},
		getIndexAngle: function(a) {
			var b = 2 * Math.PI / this.valuesCount;
			return a * b - Math.PI / 2
		},
		getPointPosition: function(b, c) {
			var a = this.getIndexAngle(b);
			return {
				x: Math.cos(a) * c + this.xCenter,
				y: Math.sin(a) * c + this.yCenter
			}
		},
		draw: function() {
			if (this.display) {
				var o = this.ctx;
				if (ab(this.yLabels, function(w, l) {
					if (l > 0) {
						var y, x = l * (this.drawingArea / this.steps),
							u = this.yCenter - x;
						if (this.lineWidth > 0) {
							if (o.strokeStyle = this.lineColor, o.lineWidth = this.lineWidth, this.lineArc) {
								o.beginPath(), o.arc(this.xCenter, this.yCenter, x, 0, 2 * Math.PI), o.closePath(), o.stroke()
							} else {
								o.beginPath();
								for (var h = 0; h < this.valuesCount; h++) {
									y = this.getPointPosition(h, this.calculateCenterOffset(this.min + l * this.stepValue)), 0 === h ? o.moveTo(y.x, y.y) : o.lineTo(y.x, y.y)
								}
								o.closePath(), o.stroke()
							}
						}
						if (this.showLabels) {
							if (o.font = H(this.fontSize, this.fontStyle, this.fontFamily), this.showLabelBackdrop) {
								var v = o.measureText(w).width;
								o.fillStyle = this.backdropColor, o.fillRect(this.xCenter - v / 2 - this.backdropPaddingX, u - this.fontSize / 2 - this.backdropPaddingY, v + 2 * this.backdropPaddingX, this.fontSize + 2 * this.backdropPaddingY)
							}
							o.textAlign = "center", o.textBaseline = "middle", o.fillStyle = this.fontColor, o.fillText(w, this.xCenter, u)
						}
					}
				}, this), !this.lineArc) {
					o.lineWidth = this.angleLineWidth, o.strokeStyle = this.angleLineColor;
					for (var k = this.valuesCount - 1; k >= 0; k--) {
						if (this.angleLineWidth > 0) {
							var f = this.getPointPosition(k, this.calculateCenterOffset(this.max));
							o.beginPath(), o.moveTo(this.xCenter, this.yCenter), o.lineTo(f.x, f.y), o.stroke(), o.closePath()
						}
						var c = this.getPointPosition(k, this.calculateCenterOffset(this.max) + 5);
						o.font = H(this.pointLabelFontSize, this.pointLabelFontStyle, this.pointLabelFontFamily), o.fillStyle = this.pointLabelFontColor;
						var p = this.labels.length,
							m = this.labels.length / 2,
							b = m / 2,
							d = k < b || k > p - b,
							g = k === b || k === p - b;
						0 === k ? o.textAlign = "center" : k === m ? o.textAlign = "center" : k < m ? o.textAlign = "left" : o.textAlign = "right", g ? o.textBaseline = "middle" : d ? o.textBaseline = "bottom" : o.textBaseline = "top", o.fillText(this.labels[k], c.x, c.y)
					}
				}
			}
		}
	}), ac.addEvent(window, "resize", function() {
		var a;
		return function() {
			clearTimeout(a), a = setTimeout(function() {
				ab(ag.instances, function(b) {
					b.options.responsive && b.resize(b.render, !0)
				})
			}, 50)
		}
	}()), aa ? define(function() {
		return ag
	}) : "object" == typeof module && module.exports && (module.exports = ag), ak.Chart = ag, V.fn.chart = function() {
		var a = [];
		return this.each(function() {
			a.push(new ag(this.getContext("2d")))
		}), 1 === a.length ? a[0] : a
	}
}.call(this, jQuery), function(b) {
	var c = b && b.zui ? b.zui : this,
		a = c.Chart,
		f = a.helpers,
		d = {
			scaleShowGridLines: !0,
			scaleGridLineColor: "rgba(0,0,0,.05)",
			scaleGridLineWidth: 1,
			scaleShowHorizontalLines: !0,
			scaleShowBeyondLine: !0,
			scaleShowVerticalLines: !0,
			bezierCurve: !0,
			bezierCurveTension: 0.4,
			pointDot: !0,
			pointDotRadius: 4,
			pointDotStrokeWidth: 1,
			pointHitDetectionRadius: 20,
			datasetStroke: !0,
			datasetStrokeWidth: 2,
			datasetFill: !0,
			legendTemplate: '<ul class="<%=name.toLowerCase()%>-legend"><% for (var i=0; i<datasets.length; i++){%><li><span style="background-color:<%=datasets[i].strokeColor%>"></span><%if(datasets[i].label){%><%=datasets[i].label%><%}%></li><%}%></ul>'
		};
	a.Type.extend({
		name: "Line",
		defaults: d,
		initialize: function(g) {
			this.PointClass = a.Point.extend({
				strokeWidth: this.options.pointDotStrokeWidth,
				radius: this.options.pointDotRadius,
				display: this.options.pointDot,
				hitDetectionRadius: this.options.pointHitDetectionRadius,
				ctx: this.chart.ctx,
				inRange: function(e) {
					return Math.pow(e - this.x, 2) < Math.pow(this.radius + this.hitDetectionRadius, 2)
				}
			}), this.datasets = [], this.options.showTooltips && f.bindEvents(this, this.options.tooltipEvents, function(h) {
				var j = "mouseout" !== h.type ? this.getPointsAtEvent(h) : [];
				this.eachPoints(function(e) {
					e.restore(["fillColor", "strokeColor"])
				}), f.each(j, function(e) {
					e.fillColor = e.highlightFill, e.strokeColor = e.highlightStroke
				}), this.showTooltip(j)
			}), f.each(g.datasets, function(h) {
				if (b.zui && b.zui.Color && b.zui.Color.get) {
					var k = b.zui.Color.get(h.color),
						j = k.toCssStr();
					h.fillColor || (h.fillColor = k.clone().fade(20).toCssStr()), h.strokeColor || (h.strokeColor = j), h.pointColor || (h.pointColor = j), h.pointStrokeColor || (h.pointStrokeColor = "#fff"), h.pointHighlightFill || (h.pointHighlightFill = "#fff"), h.pointHighlightStroke || (h.pointHighlightStroke = j)
				}
				var e = {
					label: h.label || null,
					fillColor: h.fillColor,
					strokeColor: h.strokeColor,
					pointColor: h.pointColor,
					pointStrokeColor: h.pointStrokeColor,
					showTooltips: h.showTooltips !== !1,
					points: []
				};
				this.datasets.push(e), f.each(h.data, function(l, m) {
					e.points.push(new this.PointClass({
						value: l,
						label: g.labels[m],
						datasetLabel: h.label,
						strokeColor: h.pointStrokeColor,
						fillColor: h.pointColor,
						highlightFill: h.pointHighlightFill || h.pointColor,
						highlightStroke: h.pointHighlightStroke || h.pointStrokeColor
					}))
				}, this), this.buildScale(g.labels), this.eachPoints(function(l, m) {
					f.extend(l, {
						x: this.scale.calculateX(m),
						y: this.scale.endPoint
					}), l.save()
				}, this)
			}, this), this.render()
		},
		update: function() {
			this.scale.update(), f.each(this.activeElements, function(e) {
				e.restore(["fillColor", "strokeColor"])
			}), this.eachPoints(function(e) {
				e.save()
			}), this.render()
		},
		eachPoints: function(e) {
			f.each(this.datasets, function(g) {
				f.each(g.points, e, this)
			}, this)
		},
		getPointsAtEvent: function(h) {
			var j = [],
				g = f.getRelativePosition(h);
			return f.each(this.datasets, function(e) {
				f.each(e.points, function(k) {
					k.inRange(g.x, g.y) && j.push(k)
				})
			}, this), j
		},
		buildScale: function(g) {
			var j = this,
				k = function() {
					var e = [];
					return j.eachPoints(function(l) {
						e.push(l.value)
					}), e
				},
				h = {
					templateString: this.options.scaleLabel,
					height: this.chart.height,
					width: this.chart.width,
					ctx: this.chart.ctx,
					textColor: this.options.scaleFontColor,
					fontSize: this.options.scaleFontSize,
					fontStyle: this.options.scaleFontStyle,
					fontFamily: this.options.scaleFontFamily,
					valuesCount: g.length,
					beginAtZero: this.options.scaleBeginAtZero,
					integersOnly: this.options.scaleIntegersOnly,
					calculateYRange: function(l) {
						var m = f.calculateScaleRange(k(), l, this.fontSize, this.beginAtZero, this.integersOnly);
						f.extend(this, m)
					},
					xLabels: g,
					font: f.fontString(this.options.scaleFontSize, this.options.scaleFontStyle, this.options.scaleFontFamily),
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
			this.options.scaleOverride && f.extend(h, {
				calculateYRange: f.noop,
				steps: this.options.scaleSteps,
				stepValue: this.options.scaleStepWidth,
				min: this.options.scaleStartValue,
				max: this.options.scaleStartValue + this.options.scaleSteps * this.options.scaleStepWidth
			}), this.scale = new a.Scale(h)
		},
		addData: function(g, h) {
			f.each(g, function(j, e) {
				this.datasets[e].points.push(new this.PointClass({
					value: j,
					label: h,
					datasetLabel: this.datasets[e].label,
					x: this.scale.calculateX(this.scale.valuesCount + 1),
					y: this.scale.endPoint,
					strokeColor: this.datasets[e].pointStrokeColor,
					fillColor: this.datasets[e].pointColor
				}))
			}, this), this.scale.addXLabel(h), this.update()
		},
		removeData: function() {
			this.scale.removeXLabel(), f.each(this.datasets, function(e) {
				e.points.shift()
			}, this), this.update()
		},
		reflow: function() {
			var e = f.extend({
				height: this.chart.height,
				width: this.chart.width
			});
			this.scale.update(e)
		},
		draw: function(j) {
			var l = j || 1;
			this.clear();
			var h = this.chart.ctx,
				m = function(e) {
					return null !== e.value
				},
				k = function(o, p, n) {
					return f.findNextWhere(p, m, n) || o
				},
				g = function(o, p, n) {
					return f.findPreviousWhere(p, m, n) || o
				};
			this.scale.draw(l), f.each(this.datasets, function(e) {
				var n = f.where(e.points, m);
				f.each(e.points, function(p, o) {
					p.hasValue() && p.transition({
						y: this.scale.calculateY(p.value),
						x: this.scale.calculateX(o)
					}, l)
				}, this), this.options.bezierCurve && f.each(n, function(p, q) {
					var o = q > 0 && q < n.length - 1 ? this.options.bezierCurveTension : 0;
					p.controlPoints = f.splineCurve(g(p, n, q), p, k(p, n, q), o), p.controlPoints.outer.y > this.scale.endPoint ? p.controlPoints.outer.y = this.scale.endPoint : p.controlPoints.outer.y < this.scale.startPoint && (p.controlPoints.outer.y = this.scale.startPoint), p.controlPoints.inner.y > this.scale.endPoint ? p.controlPoints.inner.y = this.scale.endPoint : p.controlPoints.inner.y < this.scale.startPoint && (p.controlPoints.inner.y = this.scale.startPoint)
				}, this), h.lineWidth = this.options.datasetStrokeWidth, h.strokeStyle = e.strokeColor, h.beginPath(), f.each(n, function(o, p) {
					if (0 === p) {
						h.moveTo(o.x, o.y)
					} else {
						if (this.options.bezierCurve) {
							var q = g(o, n, p);
							h.bezierCurveTo(q.controlPoints.outer.x, q.controlPoints.outer.y, o.controlPoints.inner.x, o.controlPoints.inner.y, o.x, o.y)
						} else {
							h.lineTo(o.x, o.y)
						}
					}
				}, this), h.stroke(), this.options.datasetFill && n.length > 0 && (h.lineTo(n[n.length - 1].x, this.scale.endPoint), h.lineTo(n[0].x, this.scale.endPoint), h.fillStyle = e.fillColor, h.closePath(), h.fill()), f.each(n, function(o) {
					o.draw()
				})
			}, this)
		}
	}), b.fn.lineChart = function(g, j) {
		var h = [];
		return this.each(function() {
			var e = b(this);
			h.push(new a(this.getContext("2d")).Line(g, b.extend(e.data(), j)))
		}), 1 === h.length ? h[0] : h
	}
}.call(this, jQuery), function(b) {
	var c = b && b.zui ? b.zui : this,
		a = c.Chart,
		f = a.helpers,
		d = {
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
	a.Type.extend({
		name: "Doughnut",
		defaults: d,
		initialize: function(e) {
			this.segments = [], this.outerRadius = (f.min([this.chart.width, this.chart.height]) - this.options.segmentStrokeWidth / 2) / 2, this.SegmentArc = a.Arc.extend({
				ctx: this.chart.ctx,
				x: this.chart.width / 2,
				y: this.chart.height / 2
			}), this.options.showTooltips && f.bindEvents(this, this.options.tooltipEvents, function(g) {
				var h = "mouseout" !== g.type ? this.getSegmentsAtEvent(g) : [];
				f.each(this.segments, function(j) {
					j.restore(["fillColor"])
				}), f.each(h, function(j) {
					j.fillColor = j.highlightColor
				}), this.showTooltip(h)
			}), this.calculateTotal(e), f.each(e, function(g, h) {
				this.addData(g, h, !0)
			}, this), this.render()
		},
		getSegmentsAtEvent: function(h) {
			var j = [],
				g = f.getRelativePosition(h);
			return f.each(this.segments, function(e) {
				e.inRange(g.x, g.y) && j.push(e)
			}, this), j
		},
		addData: function(j, g, l) {
			if (b.zui && b.zui.Color && b.zui.Color.get) {
				var k = new b.zui.Color.get(j.color);
				j.color = k.toCssStr(), j.highlight || (j.highlight = k.lighten(5).toCssStr())
			}
			var h = g || this.segments.length;
			this.segments.splice(h, 0, new this.SegmentArc({
				id: "undefined" == typeof j.id ? h : j.id,
				value: j.value,
				outerRadius: this.options.animateScale ? 0 : this.outerRadius,
				innerRadius: this.options.animateScale ? 0 : this.outerRadius / 100 * this.options.percentageInnerCutout,
				fillColor: j.color,
				highlightColor: j.highlight || j.color,
				showStroke: this.options.segmentShowStroke,
				strokeWidth: this.options.segmentStrokeWidth,
				strokeColor: this.options.segmentStrokeColor,
				startAngle: 1.5 * Math.PI,
				circumference: this.options.animateRotate ? 0 : this.calculateCircumference(j.value),
				showLabel: j.showLabel !== !1,
				circleBeginEnd: j.circleBeginEnd,
				label: j.label
			})), l || (this.reflow(), this.update())
		},
		calculateCircumference: function(e) {
			return 2 * Math.PI * (Math.abs(e) / this.total)
		},
		calculateTotal: function(e) {
			this.total = 0, f.each(e, function(g) {
				this.total += Math.abs(g.value)
			}, this)
		},
		update: function() {
			this.calculateTotal(this.segments), f.each(this.activeElements, function(e) {
				e.restore(["fillColor"])
			}), f.each(this.segments, function(e) {
				e.save()
			}), this.render()
		},
		removeData: function(g) {
			var h = f.isNumber(g) ? g : this.segments.length - 1;
			this.segments.splice(h, 1), this.reflow(), this.update()
		},
		reflow: function() {
			f.extend(this.SegmentArc.prototype, {
				x: this.chart.width / 2,
				y: this.chart.height / 2
			}), this.outerRadius = (f.min([this.chart.width, this.chart.height]) - this.options.segmentStrokeWidth / 2) / 2, f.each(this.segments, function(e) {
				e.update({
					outerRadius: this.outerRadius,
					innerRadius: this.outerRadius / 100 * this.options.percentageInnerCutout
				})
			}, this)
		},
		drawLabel: function(G, C, z) {
			var q = this.options,
				K = (G.endAngle + G.startAngle) / 2,
				w = q.scaleLabelPlacement;
			"inside" !== w && "outside" !== w && this.chart.width - this.chart.height > 50 && G.circumference < Math.PI / 18 && (w = "outside");
			var B = Math.cos(K) * G.outerRadius,
				D = Math.sin(K) * G.outerRadius,
				I = f.template(q.scaleLabel, {
					value: "undefined" == typeof C ? G.value : Math.round(C * G.value),
					label: G.label
				}),
				H = this.chart.ctx;
			H.font = f.fontString(q.scaleFontSize, q.scaleFontStyle, q.scaleFontFamily), H.textBaseline = "middle", H.textAlign = "center";
			var n = (H.measureText(I).width, this.chart.width / 2),
				x = this.chart.height / 2;
			if ("outside" === w) {
				var F = B >= 0,
					E = B + n,
					A = D + x;
				H.textAlign = F ? "left" : "right", H.measureText(I).width, B = F ? Math.max(n + G.outerRadius + 10, B + 30 + n) : Math.min(n - G.outerRadius - 10, B - 30 + n);
				var k = q.scaleFontSize * (q.scaleLineHeight || 1),
					j = Math.round((0.8 * D + x) / k) + 1,
					J = (Math.floor(this.chart.width / k) + 1, F ? 1 : -1);
				if (z[j * J] && (j > 1 ? j-- : j++), z[j * J]) {
					return
				}
				D = (j - 1) * k + q.scaleFontSize / 2, z[j * J] = !0, H.beginPath(), H.moveTo(E, A), H.lineTo(B, D), B = F ? B + 5 : B - 5, H.lineTo(B, D), H.strokeStyle = b.zui && b.zui.Color ? new b.zui.Color(G.fillColor).fade(40).toCssStr() : G.fillColor, H.strokeWidth = q.scaleLineWidth, H.stroke(), H.fillStyle = G.fillColor
			} else {
				B = 0.7 * B + n, D = 0.7 * D + x, H.fillStyle = b.zui && b.zui.Color ? new b.zui.Color(G.fillColor).contrast().toCssStr() : "#fff"
			}
			H.fillText(I, B, D)
		},
		draw: function(h) {
			var j = h ? h : 1;
			this.clear();
			var g;
			if (f.each(this.segments, function(l, e) {
				l.transition({
					circumference: this.calculateCircumference(l.value),
					outerRadius: this.outerRadius,
					innerRadius: this.outerRadius / 100 * this.options.percentageInnerCutout
				}, j), l.endAngle = l.startAngle + l.circumference, this.options.reverseDrawOrder || l.draw(), 0 === e && (l.startAngle = 1.5 * Math.PI), e < this.segments.length - 1 && (this.segments[e + 1].startAngle = l.endAngle)
			}, this), this.options.reverseDrawOrder && f.each(this.segments.slice().reverse(), function(l, m) {
				l.draw()
			}, this), this.options.scaleShowLabels) {
				var k = this.segments.slice().sort(function(l, m) {
					return m.value - l.value
				}),
					g = {};
				f.each(k, function(l, m) {
					l.showLabel && this.drawLabel(l, h, g)
				}, this)
			}
		}
	}), a.types.Doughnut.extend({
		name: "Pie",
		defaults: f.merge(d, {
			percentageInnerCutout: 0
		})
	}), b.fn.pieChart = function(g, j) {
		var h = [];
		return this.each(function() {
			var e = b(this);
			h.push(new a(this.getContext("2d")).Pie(g, b.extend(e.data(), j)))
		}), 1 === h.length ? h[0] : h
	}, b.fn.doughnutChart = function(g, j) {
		var h = [];
		return this.each(function() {
			var e = b(this);
			h.push(new a(this.getContext("2d")).Doughnut(g, b.extend(e.data(), j)))
		}), 1 === h.length ? h[0] : h
	}
}.call(this, jQuery), function(b) {
	var c = b && b.zui ? b.zui : this,
		a = c.Chart,
		f = a.helpers,
		d = {
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
	a.Type.extend({
		name: "Bar",
		defaults: d,
		initialize: function(g) {
			var h = this.options;
			this.ScaleClass = a.Scale.extend({
				offsetGridLines: !0,
				calculateBarX: function(l, o, k) {
					var p = this.calculateBaseWidth(),
						m = this.calculateX(k) - p / 2,
						j = this.calculateBarWidth(l);
					return m + j * o + o * h.barDatasetSpacing + j / 2
				},
				calculateBaseWidth: function() {
					return this.calculateX(1) - this.calculateX(0) - 2 * h.barValueSpacing
				},
				calculateBarWidth: function(j) {
					var k = this.calculateBaseWidth() - (j - 1) * h.barDatasetSpacing;
					return k / j
				}
			}), this.datasets = [], this.options.showTooltips && f.bindEvents(this, this.options.tooltipEvents, function(j) {
				var k = "mouseout" !== j.type ? this.getBarsAtEvent(j) : [];
				this.eachBars(function(e) {
					e.restore(["fillColor", "strokeColor"])
				}), f.each(k, function(e) {
					e.fillColor = e.highlightFill, e.strokeColor = e.highlightStroke
				}), this.showTooltip(k)
			}), this.BarClass = a.Rectangle.extend({
				strokeWidth: this.options.barStrokeWidth,
				showStroke: this.options.barShowStroke,
				ctx: this.chart.ctx
			}), f.each(g.datasets, function(j, m) {
				if (b.zui && b.zui.Color && b.zui.Color.get) {
					var k = b.zui.Color.get(j.color),
						e = k.toCssStr();
					j.fillColor || (j.fillColor = k.clone().fade(50).toCssStr()), j.strokeColor || (j.strokeColor = e)
				}
				var l = {
					label: j.label || null,
					fillColor: j.fillColor,
					strokeColor: j.strokeColor,
					bars: []
				};
				this.datasets.push(l), f.each(j.data, function(o, p) {
					l.bars.push(new this.BarClass({
						value: o,
						label: g.labels[p],
						datasetLabel: j.label,
						strokeColor: j.strokeColor,
						fillColor: j.fillColor,
						highlightFill: j.highlightFill || j.fillColor,
						highlightStroke: j.highlightStroke || j.strokeColor
					}))
				}, this)
			}, this), this.buildScale(g.labels), this.BarClass.prototype.base = this.scale.endPoint, this.eachBars(function(k, l, j) {
				f.extend(k, {
					width: this.scale.calculateBarWidth(this.datasets.length),
					x: this.scale.calculateBarX(this.datasets.length, j, l),
					y: this.scale.endPoint
				}), k.save()
			}, this), this.render()
		},
		update: function() {
			this.scale.update(), f.each(this.activeElements, function(e) {
				e.restore(["fillColor", "strokeColor"])
			}), this.eachBars(function(e) {
				e.save()
			}), this.render()
		},
		eachBars: function(e) {
			f.each(this.datasets, function(h, g) {
				f.each(h.bars, e, this, g)
			}, this)
		},
		getBarsAtEvent: function(j) {
			for (var l, h = [], m = f.getRelativePosition(j), k = function(e) {
					h.push(e.bars[l])
				}, g = 0; g < this.datasets.length; g++) {
				for (l = 0; l < this.datasets[g].bars.length; l++) {
					if (this.datasets[g].bars[l].inRange(m.x, m.y)) {
						return f.each(this.datasets, k), h
					}
				}
			}
			return h
		},
		buildScale: function(h) {
			var j = this,
				g = function() {
					var e = [];
					return j.eachBars(function(l) {
						e.push(l.value)
					}), e
				},
				k = {
					templateString: this.options.scaleLabel,
					height: this.chart.height,
					width: this.chart.width,
					ctx: this.chart.ctx,
					textColor: this.options.scaleFontColor,
					fontSize: this.options.scaleFontSize,
					fontStyle: this.options.scaleFontStyle,
					fontFamily: this.options.scaleFontFamily,
					valuesCount: h.length,
					beginAtZero: this.options.scaleBeginAtZero,
					integersOnly: this.options.scaleIntegersOnly,
					calculateYRange: function(l) {
						var m = f.calculateScaleRange(g(), l, this.fontSize, this.beginAtZero, this.integersOnly);
						f.extend(this, m)
					},
					xLabels: h,
					font: f.fontString(this.options.scaleFontSize, this.options.scaleFontStyle, this.options.scaleFontFamily),
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
			this.options.scaleOverride && f.extend(k, {
				calculateYRange: f.noop,
				steps: this.options.scaleSteps,
				stepValue: this.options.scaleStepWidth,
				min: this.options.scaleStartValue,
				max: this.options.scaleStartValue + this.options.scaleSteps * this.options.scaleStepWidth
			}), this.scale = new this.ScaleClass(k)
		},
		addData: function(g, h) {
			f.each(g, function(j, e) {
				this.datasets[e].bars.push(new this.BarClass({
					value: j,
					label: h,
					x: this.scale.calculateBarX(this.datasets.length, e, this.scale.valuesCount + 1),
					y: this.scale.endPoint,
					width: this.scale.calculateBarWidth(this.datasets.length),
					base: this.scale.endPoint,
					strokeColor: this.datasets[e].strokeColor,
					fillColor: this.datasets[e].fillColor
				}))
			}, this), this.scale.addXLabel(h), this.update()
		},
		removeData: function() {
			this.scale.removeXLabel(), f.each(this.datasets, function(e) {
				e.bars.shift()
			}, this), this.update()
		},
		reflow: function() {
			f.extend(this.BarClass.prototype, {
				y: this.scale.endPoint,
				base: this.scale.endPoint
			});
			var e = f.extend({
				height: this.chart.height,
				width: this.chart.width
			});
			this.scale.update(e)
		},
		drawLabel: function(h, k) {
			var g = this.options;
			k = k || g.scaleValuePlacement, k = k ? k.toLowerCase() : "auto", "auto" === k && (k = h.y < 15 ? "insdie" : "outside");
			var l = "insdie" === k ? h.y + 10 : h.y - 10,
				j = this.chart.ctx;
			j.font = f.fontString(g.scaleFontSize, g.scaleFontStyle, g.scaleFontFamily), j.textBaseline = "middle", j.textAlign = "center", j.fillStyle = g.scaleFontColor, j.fillText(h.value, h.x, l)
		},
		draw: function(h) {
			var j = h || 1;
			this.clear();
			this.chart.ctx;
			this.scale.draw(j);
			var g = this.options.scaleShowLabels && this.options.scaleValuePlacement;
			f.each(this.datasets, function(e, k) {
				f.each(e.bars, function(l, m) {
					l.hasValue() && (l.base = this.scale.endPoint, l.transition({
						x: this.scale.calculateBarX(this.datasets.length, k, m),
						y: this.scale.calculateY(l.value),
						width: this.scale.calculateBarWidth(this.datasets.length)
					}, j).draw()), g && this.drawLabel(l)
				}, this)
			}, this)
		}
	}), b.fn.barChart = function(g, j) {
		var h = [];
		return this.each(function() {
			var e = b(this);
			h.push(new a(this.getContext("2d")).Bar(g, b.extend(e.data(), j)))
		}), 1 === h.length ? h[0] : h
	}
}.call(this, jQuery),
/*!
 * Datetimepicker for Bootstrap
 * Copyright 2012 Stefan Petre
 * Licensed under the Apache License v2.0
 */
!
function(b) {
	function c() {
		return new Date(Date.UTC.apply(Date, arguments))
	}
	var a = function(j, g) {
			var h = this;
			this.element = b(j), this.language = (g.language || this.element.data("date-language") || (b.zui && b.zui.clientLang ? b.zui.clientLang().replace("_", "-") : "zh-cn")).toLowerCase(), this.language = this.language in f ? this.language : "en", this.isRTL = f[this.language].rtl || !1, this.formatType = g.formatType || this.element.data("format-type") || "standard", this.format = d.parseFormat(g.format || this.element.data("date-format") || f[this.language].format || d.getDefaultFormat(this.formatType, "input"), this.formatType), this.isInline = !1, this.isVisible = !1, this.isInput = this.element.is("input"), this.component = !! this.element.is(".date") && this.element.find(".input-group-addon .icon-th, .input-group-addon .icon-time, .input-group-addon .icon-calendar").parent(), this.componentReset = !! this.element.is(".date") && this.element.find(".input-group-addon .icon-remove").parent(), this.hasInput = this.component && this.element.find("input").length, this.component && 0 === this.component.length && (this.component = !1), this.linkField = g.linkField || this.element.data("link-field") || !1, this.linkFormat = d.parseFormat(g.linkFormat || this.element.data("link-format") || d.getDefaultFormat(this.formatType, "link"), this.formatType), this.minuteStep = g.minuteStep || this.element.data("minute-step") || 5, this.pickerPosition = g.pickerPosition || this.element.data("picker-position") || "bottom-right", this.showMeridian = g.showMeridian || this.element.data("show-meridian") || !1, this.initialDate = g.initialDate || new Date, this.pickerClass = g.eleClass, this.pickerId = g.eleId, this._attachEvents(), this.formatViewType = "datetime", "formatViewType" in g ? this.formatViewType = g.formatViewType : "formatViewType" in this.element.data() && (this.formatViewType = this.element.data("formatViewType")), this.minView = 0, "minView" in g ? this.minView = g.minView : "minView" in this.element.data() && (this.minView = this.element.data("min-view")), this.minView = d.convertViewMode(this.minView), this.maxView = d.modes.length - 1, "maxView" in g ? this.maxView = g.maxView : "maxView" in this.element.data() && (this.maxView = this.element.data("max-view")), this.maxView = d.convertViewMode(this.maxView), this.wheelViewModeNavigation = !1, "wheelViewModeNavigation" in g ? this.wheelViewModeNavigation = g.wheelViewModeNavigation : "wheelViewModeNavigation" in this.element.data() && (this.wheelViewModeNavigation = this.element.data("view-mode-wheel-navigation")), this.wheelViewModeNavigationInverseDirection = !1, "wheelViewModeNavigationInverseDirection" in g ? this.wheelViewModeNavigationInverseDirection = g.wheelViewModeNavigationInverseDirection : "wheelViewModeNavigationInverseDirection" in this.element.data() && (this.wheelViewModeNavigationInverseDirection = this.element.data("view-mode-wheel-navigation-inverse-dir")), this.wheelViewModeNavigationDelay = 100, "wheelViewModeNavigationDelay" in g ? this.wheelViewModeNavigationDelay = g.wheelViewModeNavigationDelay : "wheelViewModeNavigationDelay" in this.element.data() && (this.wheelViewModeNavigationDelay = this.element.data("view-mode-wheel-navigation-delay")), this.startViewMode = 2, "startView" in g ? this.startViewMode = g.startView : "startView" in this.element.data() && (this.startViewMode = this.element.data("start-view")), this.startViewMode = d.convertViewMode(this.startViewMode), this.viewMode = this.startViewMode, this.viewSelect = this.minView, "viewSelect" in g ? this.viewSelect = g.viewSelect : "viewSelect" in this.element.data() && (this.viewSelect = this.element.data("view-select")), this.viewSelect = d.convertViewMode(this.viewSelect), this.forceParse = !0, "forceParse" in g ? this.forceParse = g.forceParse : "dateForceParse" in this.element.data() && (this.forceParse = this.element.data("date-force-parse")), this.picker = b(d.template).appendTo(this.isInline ? this.element : "body").on({
				click: b.proxy(this.click, this),
				mousedown: b.proxy(this.mousedown, this)
			}), this.wheelViewModeNavigation && (b.fn.mousewheel ? this.picker.on({
				mousewheel: b.proxy(this.mousewheel, this)
			}) : console.log("Mouse Wheel event is not supported. Please include the jQuery Mouse Wheel plugin before enabling this option")), this.isInline ? this.picker.addClass("datetimepicker-inline") : this.picker.addClass("datetimepicker-dropdown-" + this.pickerPosition + " dropdown-menu"), this.isRTL && (this.picker.addClass("datetimepicker-rtl"), this.picker.find(".prev span, .next span").toggleClass("icon-arrow-left icon-arrow-right")), b(document).on("mousedown", function(k) {
				0 === b(k.target).closest(".datetimepicker").length && h.hide()
			}), this.autoclose = !1, "autoclose" in g ? this.autoclose = g.autoclose : "dateAutoclose" in this.element.data() && (this.autoclose = this.element.data("date-autoclose")), this.keyboardNavigation = !0, "keyboardNavigation" in g ? this.keyboardNavigation = g.keyboardNavigation : "dateKeyboardNavigation" in this.element.data() && (this.keyboardNavigation = this.element.data("date-keyboard-navigation")), this.todayBtn = g.todayBtn || this.element.data("date-today-btn") || !1, this.todayHighlight = g.todayHighlight || this.element.data("date-today-highlight") || !1, this.weekStart = (g.weekStart || this.element.data("date-weekstart") || f[this.language].weekStart || 0) % 7, this.weekEnd = (this.weekStart + 6) % 7, this.startDate = -(1 / 0), this.endDate = 1 / 0, this.daysOfWeekDisabled = [], this.setStartDate(g.startDate || this.element.data("date-startdate")), this.setEndDate(g.endDate || this.element.data("date-enddate")), this.setDaysOfWeekDisabled(g.daysOfWeekDisabled || this.element.data("date-days-of-week-disabled")), this.fillDow(), this.fillMonths(), this.update(), this.showMode(), this.isInline && this.show()
		};
	a.prototype = {
		constructor: a,
		_events: [],
		_attachEvents: function() {
			this._detachEvents(), this.isInput ? this._events = [
				[this.element,
				{
					focus: b.proxy(this.show, this),
					keyup: b.proxy(this.update, this),
					keydown: b.proxy(this.keydown, this)
				}]
			] : this.component && this.hasInput ? (this._events = [
				[this.element.find("input"),
				{
					focus: b.proxy(this.show, this),
					keyup: b.proxy(this.update, this),
					keydown: b.proxy(this.keydown, this)
				}],
				[this.component,
				{
					click: b.proxy(this.show, this)
				}]
			], this.componentReset && this._events.push([this.componentReset,
			{
				click: b.proxy(this.reset, this)
			}])) : this.element.is("div") ? this.isInline = !0 : this._events = [
				[this.element,
				{
					click: b.proxy(this.show, this)
				}]
			];
			for (var h, g, j = 0; j < this._events.length; j++) {
				h = this._events[j][0], g = this._events[j][1], h.on(g)
			}
		},
		_detachEvents: function() {
			for (var h, j, g = 0; g < this._events.length; g++) {
				h = this._events[g][0], j = this._events[g][1], h.off(j)
			}
			this._events = []
		},
		show: function(g) {
			this.picker.show(), this.height = this.component ? this.component.outerHeight() : this.element.outerHeight(), this.forceParse && this.update(), this.place(), b(window).on("resize", b.proxy(this.place, this)), g && (g.stopPropagation(), g.preventDefault()), this.isVisible = !0, this.element.trigger({
				type: "show",
				date: this.date
			})
		},
		hide: function(g) {
			this.isVisible && (this.isInline || (this.picker.hide(), b(window).off("resize", this.place), this.viewMode = this.startViewMode, this.showMode(), this.isInput || b(document).off("mousedown", this.hide), this.forceParse && (this.isInput && this.element.val() || this.hasInput && this.element.find("input").val()) && this.setValue(), this.isVisible = !1, this.element.trigger({
				type: "hide",
				date: this.date
			})))
		},
		remove: function() {
			this._detachEvents(), this.picker.remove(), delete this.picker, delete this.element.data().datetimepicker
		},
		getDate: function() {
			var e = this.getUTCDate();
			return new Date(e.getTime() + 60000 * e.getTimezoneOffset())
		},
		getUTCDate: function() {
			return this.date
		},
		setDate: function(e) {
			this.setUTCDate(new Date(e.getTime() - 60000 * e.getTimezoneOffset()))
		},
		setUTCDate: function(e) {
			e >= this.startDate && e <= this.endDate ? (this.date = e, this.setValue(), this.viewDate = this.date, this.fill()) : this.element.trigger({
				type: "outOfRange",
				date: e,
				startDate: this.startDate,
				endDate: this.endDate
			})
		},
		setFormat: function(g) {
			this.format = d.parseFormat(g, this.formatType);
			var h;
			this.isInput ? h = this.element : this.component && (h = this.element.find("input")), h && h.val() && this.setValue()
		},
		setValue: function() {
			var g = this.getFormattedDate();
			this.isInput ? this.element.val(g) : (this.component && this.element.find("input").val(g), this.element.data("date", g)), this.linkField && b("#" + this.linkField).val(this.getFormattedDate(this.linkFormat))
		},
		getFormattedDate: function(e) {
			return void 0 == e && (e = this.format), d.formatDate(this.date, e, this.language, this.formatType)
		},
		setStartDate: function(e) {
			this.startDate = e || -(1 / 0), this.startDate !== -(1 / 0) && (this.startDate = d.parseDate(this.startDate, this.format, this.language, this.formatType)), this.update(), this.updateNavArrows()
		},
		setEndDate: function(e) {
			this.endDate = e || 1 / 0, this.endDate !== 1 / 0 && (this.endDate = d.parseDate(this.endDate, this.format, this.language, this.formatType)), this.update(), this.updateNavArrows()
		},
		setDaysOfWeekDisabled: function(g) {
			this.daysOfWeekDisabled = g || [], b.isArray(this.daysOfWeekDisabled) || (this.daysOfWeekDisabled = this.daysOfWeekDisabled.split(/,\s*/)), this.daysOfWeekDisabled = b.map(this.daysOfWeekDisabled, function(e) {
				return parseInt(e, 10)
			}), this.update(), this.updateNavArrows()
		},
		place: function() {
			if (!this.isInline) {
				var l = 0;
				b("div").each(function() {
					var e = parseInt(b(this).css("zIndex"), 10);
					e > l && (l = e)
				});
				var h, p, m, j = l + 10;
				this.component ? (h = this.component.offset(), m = h.left, "bottom-left" !== this.pickerPosition && "top-left" !== this.pickerPosition && "auto-left" !== this.pickerPosition || (m += this.component.outerWidth() - this.picker.outerWidth())) : (h = this.element.offset(), m = h.left);
				var g = 0 === this.pickerPosition.indexOf("auto-"),
					k = g ? (h.top + this.picker.outerHeight() > b(window).height() + b(window).scrollTop() ? "top" : "bottom") + (0 === this.pickerPosition.lastIndexOf("-left") ? "-left" : "-right") : this.pickerPosition;
				p = "top-left" === k || "top-right" === k ? h.top - this.picker.outerHeight() : h.top + this.height, this.picker.css({
					top: p,
					left: m,
					zIndex: j
				}).attr("class", "datetimepicker dropdown-menu datetimepicker-dropdown-" + k), this.pickerClass && this.picker.addClass(this.pickerClass), this.pickerId && this.picker.attr("id", this.pickerId)
			}
		},
		update: function() {
			var g, h = !1;
			arguments && arguments.length && ("string" == typeof arguments[0] || arguments[0] instanceof Date) ? (g = arguments[0], h = !0) : (g = this.element.data("date") || (this.isInput ? this.element.val() : this.element.find("input").val()) || this.initialDate, ("string" == typeof g || g instanceof String) && (g = g.replace(/^\s+|\s+$/g, ""))), g || (g = new Date, h = !1), this.date = d.parseDate(g, this.format, this.language, this.formatType), h && this.setValue(), this.date < this.startDate ? this.viewDate = new Date(this.startDate) : this.date > this.endDate ? this.viewDate = new Date(this.endDate) : this.viewDate = new Date(this.date), this.fill()
		},
		fillDow: function() {
			for (var g = this.weekStart, h = "<tr>"; g < this.weekStart + 7;) {
				h += '<th class="dow">' + f[this.language].daysMin[g++ % 7] + "</th>"
			}
			h += "</tr>", this.picker.find(".datetimepicker-days thead").append(h)
		},
		fillMonths: function() {
			for (var g = "", h = 0; h < 12;) {
				g += '<span class="month">' + f[this.language].monthsShort[h++] + "</span>"
			}
			this.picker.find(".datetimepicker-months td").html(g)
		},
		fill: function() {
			if (null != this.date && null != this.viewDate) {
				var X = new Date(this.viewDate),
					N = X.getUTCFullYear(),
					ae = X.getUTCMonth(),
					Q = X.getUTCDate(),
					V = X.getUTCHours(),
					Y = X.getUTCMinutes(),
					ac = this.startDate !== -(1 / 0) ? this.startDate.getUTCFullYear() : -(1 / 0),
					ab = this.startDate !== -(1 / 0) ? this.startDate.getUTCMonth() : -(1 / 0),
					J = this.endDate !== 1 / 0 ? this.endDate.getUTCFullYear() : 1 / 0,
					R = this.endDate !== 1 / 0 ? this.endDate.getUTCMonth() : 1 / 0,
					aa = new c(this.date.getUTCFullYear(), this.date.getUTCMonth(), this.date.getUTCDate()).valueOf(),
					Z = new Date;
				if (this.picker.find(".datetimepicker-days thead th:eq(1)").text(f[this.language].months[ae] + " " + N), "time" == this.formatViewType) {
					var U = V % 12 ? V % 12 : 12,
						H = (U < 10 ? "0" : "") + U,
						B = (Y < 10 ? "0" : "") + Y,
						ad = f[this.language].meridiem[V < 12 ? 0 : 1];
					this.picker.find(".datetimepicker-hours thead th:eq(1)").text(H + ":" + B + " " + ad.toUpperCase()), this.picker.find(".datetimepicker-minutes thead th:eq(1)").text(H + ":" + B + " " + ad.toUpperCase())
				} else {
					this.picker.find(".datetimepicker-hours thead th:eq(1)").text(Q + " " + f[this.language].months[ae] + " " + N), this.picker.find(".datetimepicker-minutes thead th:eq(1)").text(Q + " " + f[this.language].months[ae] + " " + N)
				}
				this.picker.find("tfoot th.today").text(f[this.language].today).toggle(this.todayBtn !== !1), this.updateNavArrows(), this.fillMonths();
				var G = c(N, ae - 1, 28, 0, 0, 0, 0),
					E = d.getDaysInMonth(G.getUTCFullYear(), G.getUTCMonth());
				G.setUTCDate(E), G.setUTCDate(E - (G.getUTCDay() - this.weekStart + 7) % 7);
				var O = new Date(G);
				O.setUTCDate(O.getUTCDate() + 42), O = O.valueOf();
				for (var af, W = []; G.valueOf() < O;) {
					G.getUTCDay() == this.weekStart && W.push("<tr>"), af = "", G.getUTCFullYear() < N || G.getUTCFullYear() == N && G.getUTCMonth() < ae ? af += " old" : (G.getUTCFullYear() > N || G.getUTCFullYear() == N && G.getUTCMonth() > ae) && (af += " new"), this.todayHighlight && G.getUTCFullYear() == Z.getFullYear() && G.getUTCMonth() == Z.getMonth() && G.getUTCDate() == Z.getDate() && (af += " today"), G.valueOf() == aa && (af += " active"), (G.valueOf() + 86400000 <= this.startDate || G.valueOf() > this.endDate || b.inArray(G.getUTCDay(), this.daysOfWeekDisabled) !== -1) && (af += " disabled"), W.push('<td class="day' + af + '">' + G.getUTCDate() + "</td>"), G.getUTCDay() == this.weekEnd && W.push("</tr>"), G.setUTCDate(G.getUTCDate() + 1)
				}
				this.picker.find(".datetimepicker-days tbody").empty().append(W.join("")), W = [];
				for (var e = "", j = "", K = "", o = 0; o < 24; o++) {
					var n = c(N, ae, Q, o);
					af = "", n.valueOf() + 3600000 <= this.startDate || n.valueOf() > this.endDate ? af += " disabled" : V == o && (af += " active"), this.showMeridian && 2 == f[this.language].meridiem.length ? (j = o < 12 ? f[this.language].meridiem[0] : f[this.language].meridiem[1], j != K && ("" != K && W.push("</fieldset>"), W.push('<fieldset class="hour"><legend>' + j.toUpperCase() + "</legend>")), K = j, e = o % 12 ? o % 12 : 12, W.push('<span class="hour' + af + " hour_" + (o < 12 ? "am" : "pm") + '">' + e + "</span>"), 23 == o && W.push("</fieldset>")) : (e = o + ":00", W.push('<span class="hour' + af + '">' + e + "</span>"))
				}
				this.picker.find(".datetimepicker-hours td").html(W.join("")), W = [], e = "", j = "", K = "";
				for (var o = 0; o < 60; o += this.minuteStep) {
					var n = c(N, ae, Q, V, o, 0);
					af = "", n.valueOf() < this.startDate || n.valueOf() > this.endDate ? af += " disabled" : Math.floor(Y / this.minuteStep) == Math.floor(o / this.minuteStep) && (af += " active"), this.showMeridian && 2 == f[this.language].meridiem.length ? (j = V < 12 ? f[this.language].meridiem[0] : f[this.language].meridiem[1], j != K && ("" != K && W.push("</fieldset>"), W.push('<fieldset class="minute"><legend>' + j.toUpperCase() + "</legend>")), K = j, e = V % 12 ? V % 12 : 12, W.push('<span class="minute' + af + '">' + e + ":" + (o < 10 ? "0" + o : o) + "</span>"), 59 == o && W.push("</fieldset>")) : (e = o + ":00", W.push('<span class="minute' + af + '">' + V + ":" + (o < 10 ? "0" + o : o) + "</span>"))
				}
				this.picker.find(".datetimepicker-minutes td").html(W.join(""));
				var I = this.date.getUTCFullYear(),
					q = this.picker.find(".datetimepicker-months").find("th:eq(1)").text(N).end().find("span").removeClass("active");
				I == N && q.eq(this.date.getUTCMonth()).addClass("active"), (N < ac || N > J) && q.addClass("disabled"), N == ac && q.slice(0, ab).addClass("disabled"), N == J && q.slice(R + 1).addClass("disabled"), W = "", N = 10 * parseInt(N / 10, 10);
				var A = this.picker.find(".datetimepicker-years").find("th:eq(1)").text(N + "-" + (N + 9)).end().find("td");
				N -= 1;
				for (var o = -1; o < 11; o++) {
					W += '<span class="year' + (o == -1 || 10 == o ? " old" : "") + (I == N ? " active" : "") + (N < ac || N > J ? " disabled" : "") + '">' + N + "</span>", N += 1
				}
				A.html(W), this.place()
			}
		},
		updateNavArrows: function() {
			var h = new Date(this.viewDate),
				j = h.getUTCFullYear(),
				g = h.getUTCMonth(),
				l = h.getUTCDate(),
				k = h.getUTCHours();
			switch (this.viewMode) {
			case 0:
				this.startDate !== -(1 / 0) && j <= this.startDate.getUTCFullYear() && g <= this.startDate.getUTCMonth() && l <= this.startDate.getUTCDate() && k <= this.startDate.getUTCHours() ? this.picker.find(".prev").css({
					visibility: "hidden"
				}) : this.picker.find(".prev").css({
					visibility: "visible"
				}), this.endDate !== 1 / 0 && j >= this.endDate.getUTCFullYear() && g >= this.endDate.getUTCMonth() && l >= this.endDate.getUTCDate() && k >= this.endDate.getUTCHours() ? this.picker.find(".next").css({
					visibility: "hidden"
				}) : this.picker.find(".next").css({
					visibility: "visible"
				});
				break;
			case 1:
				this.startDate !== -(1 / 0) && j <= this.startDate.getUTCFullYear() && g <= this.startDate.getUTCMonth() && l <= this.startDate.getUTCDate() ? this.picker.find(".prev").css({
					visibility: "hidden"
				}) : this.picker.find(".prev").css({
					visibility: "visible"
				}), this.endDate !== 1 / 0 && j >= this.endDate.getUTCFullYear() && g >= this.endDate.getUTCMonth() && l >= this.endDate.getUTCDate() ? this.picker.find(".next").css({
					visibility: "hidden"
				}) : this.picker.find(".next").css({
					visibility: "visible"
				});
				break;
			case 2:
				this.startDate !== -(1 / 0) && j <= this.startDate.getUTCFullYear() && g <= this.startDate.getUTCMonth() ? this.picker.find(".prev").css({
					visibility: "hidden"
				}) : this.picker.find(".prev").css({
					visibility: "visible"
				}), this.endDate !== 1 / 0 && j >= this.endDate.getUTCFullYear() && g >= this.endDate.getUTCMonth() ? this.picker.find(".next").css({
					visibility: "hidden"
				}) : this.picker.find(".next").css({
					visibility: "visible"
				});
				break;
			case 3:
			case 4:
				this.startDate !== -(1 / 0) && j <= this.startDate.getUTCFullYear() ? this.picker.find(".prev").css({
					visibility: "hidden"
				}) : this.picker.find(".prev").css({
					visibility: "visible"
				}), this.endDate !== 1 / 0 && j >= this.endDate.getUTCFullYear() ? this.picker.find(".next").css({
					visibility: "hidden"
				}) : this.picker.find(".next").css({
					visibility: "visible"
				})
			}
		},
		mousewheel: function(h) {
			if (h.preventDefault(), h.stopPropagation(), !this.wheelPause) {
				this.wheelPause = !0;
				var g = h.originalEvent,
					k = g.wheelDelta,
					j = k > 0 ? 1 : 0 === k ? 0 : -1;
				this.wheelViewModeNavigationInverseDirection && (j = -j), this.showMode(j), setTimeout(b.proxy(function() {
					this.wheelPause = !1
				}, this), this.wheelViewModeNavigationDelay)
			}
		},
		click: function(m) {
			m.stopPropagation(), m.preventDefault();
			var j = b(m.target).closest("span, td, th, legend");
			if (1 == j.length) {
				if (j.is(".disabled")) {
					return void this.element.trigger({
						type: "outOfRange",
						date: this.viewDate,
						startDate: this.startDate,
						endDate: this.endDate
					})
				}
				switch (j[0].nodeName.toLowerCase()) {
				case "th":
					switch (j[0].className) {
					case "switch":
						this.showMode(1);
						break;
					case "prev":
					case "next":
						var y = d.modes[this.viewMode].navStep * ("prev" == j[0].className ? -1 : 1);
						switch (this.viewMode) {
						case 0:
							this.viewDate = this.moveHour(this.viewDate, y);
							break;
						case 1:
							this.viewDate = this.moveDate(this.viewDate, y);
							break;
						case 2:
							this.viewDate = this.moveMonth(this.viewDate, y);
							break;
						case 3:
						case 4:
							this.viewDate = this.moveYear(this.viewDate, y)
						}
						this.fill();
						break;
					case "today":
						var w = new Date;
						w = c(w.getFullYear(), w.getMonth(), w.getDate(), w.getHours(), w.getMinutes(), w.getSeconds(), 0), w < this.startDate ? w = this.startDate : w > this.endDate && (w = this.endDate), this.viewMode = this.startViewMode, this.showMode(0), this._setDate(w), this.fill(), this.autoclose && this.hide()
					}
					break;
				case "span":
					if (!j.is(".disabled")) {
						var e = this.viewDate.getUTCFullYear(),
							k = this.viewDate.getUTCMonth(),
							o = this.viewDate.getUTCDate(),
							v = this.viewDate.getUTCHours(),
							q = this.viewDate.getUTCMinutes(),
							x = this.viewDate.getUTCSeconds();
						if (j.is(".month") ? (this.viewDate.setUTCDate(1), k = j.parent().find("span").index(j), o = this.viewDate.getUTCDate(), this.viewDate.setUTCMonth(k), this.element.trigger({
							type: "changeMonth",
							date: this.viewDate
						}), this.viewSelect >= 3 && this._setDate(c(e, k, o, v, q, x, 0))) : j.is(".year") ? (this.viewDate.setUTCDate(1), e = parseInt(j.text(), 10) || 0, this.viewDate.setUTCFullYear(e), this.element.trigger({
							type: "changeYear",
							date: this.viewDate
						}), this.viewSelect >= 4 && this._setDate(c(e, k, o, v, q, x, 0))) : j.is(".hour") ? (v = parseInt(j.text(), 10) || 0, (j.hasClass("hour_am") || j.hasClass("hour_pm")) && (12 == v && j.hasClass("hour_am") ? v = 0 : 12 != v && j.hasClass("hour_pm") && (v += 12)), this.viewDate.setUTCHours(v), this.element.trigger({
							type: "changeHour",
							date: this.viewDate
						}), this.viewSelect >= 1 && this._setDate(c(e, k, o, v, q, x, 0))) : j.is(".minute") && (q = parseInt(j.text().substr(j.text().indexOf(":") + 1), 10) || 0, this.viewDate.setUTCMinutes(q), this.element.trigger({
							type: "changeMinute",
							date: this.viewDate
						}), this.viewSelect >= 0 && this._setDate(c(e, k, o, v, q, x, 0))), 0 != this.viewMode) {
							var g = this.viewMode;
							this.showMode(-1), this.fill(), g == this.viewMode && this.autoclose && this.hide()
						} else {
							this.fill(), this.autoclose && this.hide()
						}
					}
					break;
				case "td":
					if (j.is(".day") && !j.is(".disabled")) {
						var o = parseInt(j.text(), 10) || 1,
							e = this.viewDate.getUTCFullYear(),
							k = this.viewDate.getUTCMonth(),
							v = this.viewDate.getUTCHours(),
							q = this.viewDate.getUTCMinutes(),
							x = this.viewDate.getUTCSeconds();
						j.is(".old") ? 0 === k ? (k = 11, e -= 1) : k -= 1 : j.is(".new") && (11 == k ? (k = 0, e += 1) : k += 1), this.viewDate.setUTCFullYear(e), this.viewDate.setUTCMonth(k, o), this.element.trigger({
							type: "changeDay",
							date: this.viewDate
						}), this.viewSelect >= 2 && this._setDate(c(e, k, o, v, q, x, 0))
					}
					var g = this.viewMode;
					this.showMode(-1), this.fill(), g == this.viewMode && this.autoclose && this.hide()
				}
			}
		},
		_setDate: function(h, j) {
			j && "date" != j || (this.date = h), j && "view" != j || (this.viewDate = h), this.fill(), this.setValue();
			var g;
			this.isInput ? g = this.element : this.component && (g = this.element.find("input")), g && (g.change(), this.autoclose && (!j || "date" == j)), this.element.trigger({
				type: "changeDate",
				date: this.date
			}), null === h && (this.date = this.viewDate)
		},
		moveMinute: function(h, j) {
			if (!j) {
				return h
			}
			var g = new Date(h.valueOf());
			return g.setUTCMinutes(g.getUTCMinutes() + j * this.minuteStep), g
		},
		moveHour: function(h, j) {
			if (!j) {
				return h
			}
			var g = new Date(h.valueOf());
			return g.setUTCHours(g.getUTCHours() + j), g
		},
		moveDate: function(h, j) {
			if (!j) {
				return h
			}
			var g = new Date(h.valueOf());
			return g.setUTCDate(g.getUTCDate() + j), g
		},
		moveMonth: function(u, p) {
			if (!p) {
				return u
			}
			var m, j, h = new Date(u.valueOf()),
				v = h.getUTCDate(),
				q = h.getUTCMonth(),
				g = Math.abs(p);
			if (p = p > 0 ? 1 : -1, 1 == g) {
				j = p == -1 ?
				function() {
					return h.getUTCMonth() == q
				} : function() {
					return h.getUTCMonth() != m
				}, m = q + p, h.setUTCMonth(m), (m < 0 || m > 11) && (m = (m + 12) % 12)
			} else {
				for (var k = 0; k < g; k++) {
					h = this.moveMonth(h, p)
				}
				m = h.getUTCMonth(), h.setUTCDate(v), j = function() {
					return m != h.getUTCMonth()
				}
			}
			for (; j();) {
				h.setUTCDate(--v), h.setUTCMonth(m)
			}
			return h
		},
		moveYear: function(g, h) {
			return this.moveMonth(g, 12 * h)
		},
		dateWithinRange: function(e) {
			return e >= this.startDate && e <= this.endDate
		},
		keydown: function(j) {
			if (this.picker.is(":not(:visible)")) {
				return void(27 == j.keyCode && this.show())
			}
			var l, h, p, m = !1;
			switch (j.keyCode) {
			case 27:
				this.hide(), j.preventDefault();
				break;
			case 37:
			case 39:
				if (!this.keyboardNavigation) {
					break
				}
				l = 37 == j.keyCode ? -1 : 1, viewMode = this.viewMode, j.ctrlKey ? viewMode += 2 : j.shiftKey && (viewMode += 1), 4 == viewMode ? (h = this.moveYear(this.date, l), p = this.moveYear(this.viewDate, l)) : 3 == viewMode ? (h = this.moveMonth(this.date, l), p = this.moveMonth(this.viewDate, l)) : 2 == viewMode ? (h = this.moveDate(this.date, l), p = this.moveDate(this.viewDate, l)) : 1 == viewMode ? (h = this.moveHour(this.date, l), p = this.moveHour(this.viewDate, l)) : 0 == viewMode && (h = this.moveMinute(this.date, l), p = this.moveMinute(this.viewDate, l)), this.dateWithinRange(h) && (this.date = h, this.viewDate = p, this.setValue(), this.update(), j.preventDefault(), m = !0);
				break;
			case 38:
			case 40:
				if (!this.keyboardNavigation) {
					break
				}
				l = 38 == j.keyCode ? -1 : 1, viewMode = this.viewMode, j.ctrlKey ? viewMode += 2 : j.shiftKey && (viewMode += 1), 4 == viewMode ? (h = this.moveYear(this.date, l), p = this.moveYear(this.viewDate, l)) : 3 == viewMode ? (h = this.moveMonth(this.date, l), p = this.moveMonth(this.viewDate, l)) : 2 == viewMode ? (h = this.moveDate(this.date, 7 * l), p = this.moveDate(this.viewDate, 7 * l)) : 1 == viewMode ? this.showMeridian ? (h = this.moveHour(this.date, 6 * l), p = this.moveHour(this.viewDate, 6 * l)) : (h = this.moveHour(this.date, 4 * l), p = this.moveHour(this.viewDate, 4 * l)) : 0 == viewMode && (h = this.moveMinute(this.date, 4 * l), p = this.moveMinute(this.viewDate, 4 * l)), this.dateWithinRange(h) && (this.date = h, this.viewDate = p, this.setValue(), this.update(), j.preventDefault(), m = !0);
				break;
			case 13:
				if (0 != this.viewMode) {
					var k = this.viewMode;
					this.showMode(-1), this.fill(), k == this.viewMode && this.autoclose && this.hide()
				} else {
					this.fill(), this.autoclose && this.hide()
				}
				j.preventDefault();
				break;
			case 9:
				this.hide()
			}
			if (m) {
				var g;
				this.isInput ? g = this.element : this.component && (g = this.element.find("input")), g && g.change(), this.element.trigger({
					type: "changeDate",
					date: this.date
				})
			}
		},
		showMode: function(g) {
			if (g) {
				var h = Math.max(0, Math.min(d.modes.length - 1, this.viewMode + g));
				h >= this.minView && h <= this.maxView && (this.element.trigger({
					type: "changeMode",
					date: this.viewDate,
					oldViewMode: this.viewMode,
					newViewMode: h
				}), this.viewMode = h)
			}
			this.picker.find(">div").hide().filter(".datetimepicker-" + d.modes[this.viewMode].clsName).css("display", "block"), this.updateNavArrows()
		},
		reset: function(e) {
			this._setDate(null, "date")
		}
	}, b.fn.datetimepicker = function(g) {
		var h = Array.apply(null, arguments);
		return h.shift(), this.each(function() {
			var k = b(this),
				j = k.data("datetimepicker"),
				e = "object" == typeof g && g;
			j || k.data("datetimepicker", j = new a(this, b.extend({}, b.fn.datetimepicker.defaults, k.data(), e))), "string" == typeof g && "function" == typeof j[g] && j[g].apply(j, h)
		})
	}, b.fn.datetimepicker.defaults = {
		pickerPosition: "auto-right"
	}, b.fn.datetimepicker.Constructor = a;
	var f = b.fn.datetimepicker.dates = {
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
	f["zh-cn"] = {
		days: ["星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"],
		daysShort: ["周日", "周一", "周二", "周三", "周四", "周五", "周六", "周日"],
		daysMin: ["日", "一", "二", "三", "四", "五", "六", "日"],
		months: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
		monthsShort: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
		today: "今日",
		suffix: [],
		meridiem: []
	}, f["zh-tw"] = {
		days: ["星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"],
		daysShort: ["周日", "周一", "周二", "周三", "周四", "周五", "周六", "周日"],
		daysMin: ["日", "一", "二", "三", "四", "五", "六", "日"],
		months: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
		monthsShort: ["一月", "二月", "三月", "四月", "五月", "六月", "七月", "八月", "九月", "十月", "十一月", "十二月"],
		today: "今天",
		suffix: [],
		meridiem: ["上午", "下午"]
	};
	var d = {
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
		isLeapYear: function(e) {
			return e % 4 === 0 && e % 100 !== 0 || e % 400 === 0
		},
		getDaysInMonth: function(g, h) {
			return [31, d.isLeapYear(g) ? 29 : 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31][h]
		},
		getDefaultFormat: function(g, h) {
			if ("standard" == g) {
				return "input" == h ? "yyyy-mm-dd hh:ii" : "yyyy-mm-dd hh:ii:ss"
			}
			if ("php" == g) {
				return "input" == h ? "Y-m-d H:i" : "Y-m-d H:i:s"
			}
			throw new Error("Invalid format type.")
		},
		validParts: function(e) {
			if ("standard" == e) {
				return /hh?|HH?|p|P|ii?|ss?|dd?|DD?|mm?|MM?|yy(?:yy)?/g
			}
			if ("php" == e) {
				return /[dDjlNwzFmMnStyYaABgGhHis]/g
			}
			throw new Error("Invalid format type.")
		},
		nonpunctuation: /[^ -\/:-@\[-`{-~\t\n\rTZ]+/g,
		parseFormat: function(h, j) {
			var g = h.replace(this.validParts(j), "\0").split("\0"),
				k = h.match(this.validParts(j));
			if (!g || !g.length || !k || 0 == k.length) {
				throw new Error("Invalid date format.")
			}
			return {
				separators: g,
				parts: k
			}
		},
		parseDate: function(k, J, F, e) {
			if (k instanceof Date) {
				var q = new Date(k.valueOf() - 60000 * k.getTimezoneOffset());
				return q.setMilliseconds(0), q
			}
			if (/^\d{4}\-\d{1,2}\-\d{1,2}$/.test(k) && (J = this.parseFormat("yyyy-mm-dd", e)), /^\d{4}\-\d{1,2}\-\d{1,2}[T ]\d{1,2}\:\d{1,2}$/.test(k) && (J = this.parseFormat("yyyy-mm-dd hh:ii", e)), /^\d{4}\-\d{1,2}\-\d{1,2}[T ]\d{1,2}\:\d{1,2}\:\d{1,2}[Z]{0,1}$/.test(k) && (J = this.parseFormat("yyyy-mm-dd hh:ii:ss", e)), /^[-+]\d+[dmwy]([\s,]+[-+]\d+[dmwy])*$/.test(k)) {
				var x, C, B = /([-+]\d+)([dmwy])/,
					I = k.match(/([-+]\d+)([dmwy])/g);
				k = new Date;
				for (var j = 0; j < I.length; j++) {
					switch (x = B.exec(I[j]), C = parseInt(x[1]), x[2]) {
					case "d":
						k.setUTCDate(k.getUTCDate() + C);
						break;
					case "m":
						k = a.prototype.moveMonth.call(a.prototype, k, C);
						break;
					case "w":
						k.setUTCDate(k.getUTCDate() + 7 * C);
						break;
					case "y":
						k = a.prototype.moveYear.call(a.prototype, k, C)
					}
				}
				return c(k.getUTCFullYear(), k.getUTCMonth(), k.getUTCDate(), k.getUTCHours(), k.getUTCMinutes(), k.getUTCSeconds(), 0)
			}
			var A, z, x, I = k && k.match(this.nonpunctuation) || [],
				k = new Date(0, 0, 0, 0, 0, 0, 0),
				n = {},
				H = ["hh", "h", "ii", "i", "ss", "s", "yyyy", "yy", "M", "MM", "m", "mm", "D", "DD", "d", "dd", "H", "HH", "p", "P"],
				E = {
					hh: function(g, h) {
						return g.setUTCHours(h)
					},
					h: function(g, h) {
						return g.setUTCHours(h)
					},
					HH: function(g, h) {
						return g.setUTCHours(12 == h ? 0 : h)
					},
					H: function(g, h) {
						return g.setUTCHours(12 == h ? 0 : h)
					},
					ii: function(g, h) {
						return g.setUTCMinutes(h)
					},
					i: function(g, h) {
						return g.setUTCMinutes(h)
					},
					ss: function(g, h) {
						return g.setUTCSeconds(h)
					},
					s: function(g, h) {
						return g.setUTCSeconds(h)
					},
					yyyy: function(g, h) {
						return g.setUTCFullYear(h)
					},
					yy: function(g, h) {
						return g.setUTCFullYear(2000 + h)
					},
					m: function(g, h) {
						for (h -= 1; h < 0;) {
							h += 12
						}
						for (h %= 12, g.setUTCMonth(h); g.getUTCMonth() != h;) {
							g.setUTCDate(g.getUTCDate() - 1)
						}
						return g
					},
					d: function(g, h) {
						return g.setUTCDate(h)
					},
					p: function(g, h) {
						return g.setUTCHours(1 == h ? g.getUTCHours() + 12 : g.getUTCHours())
					}
				};
			if (E.M = E.MM = E.mm = E.m, E.dd = E.d, E.P = E.p, k = c(k.getFullYear(), k.getMonth(), k.getDate(), k.getHours(), k.getMinutes(), k.getSeconds()), I.length == J.parts.length) {
				for (var j = 0, D = J.parts.length; j < D; j++) {
					if (A = parseInt(I[j], 10), x = J.parts[j], isNaN(A)) {
						switch (x) {
						case "MM":
							z = b(f[F].months).filter(function() {
								var g = this.slice(0, I[j].length),
									h = I[j].slice(0, g.length);
								return g == h
							}), A = b.inArray(z[0], f[F].months) + 1;
							break;
						case "M":
							z = b(f[F].monthsShort).filter(function() {
								var g = this.slice(0, I[j].length),
									h = I[j].slice(0, g.length);
								return g == h
							}), A = b.inArray(z[0], f[F].monthsShort) + 1;
							break;
						case "p":
						case "P":
							A = b.inArray(I[j].toLowerCase(), f[F].meridiem)
						}
					}
					n[x] = A
				}
				for (var G, j = 0; j < H.length; j++) {
					G = H[j], G in n && !isNaN(n[G]) && E[G](k, n[G])
				}
			}
			return k
		},
		formatDate: function(p, k, n, j) {
			if (null == p) {
				return ""
			}
			var o;
			if ("standard" == j) {
				o = {
					yy: p.getUTCFullYear().toString().substring(2),
					yyyy: p.getUTCFullYear(),
					m: p.getUTCMonth() + 1,
					M: f[n].monthsShort[p.getUTCMonth()],
					MM: f[n].months[p.getUTCMonth()],
					d: p.getUTCDate(),
					D: f[n].daysShort[p.getUTCDay()],
					DD: f[n].days[p.getUTCDay()],
					p: 2 == f[n].meridiem.length ? f[n].meridiem[p.getUTCHours() < 12 ? 0 : 1] : "",
					h: p.getUTCHours(),
					i: p.getUTCMinutes(),
					s: p.getUTCSeconds()
				}, 2 == f[n].meridiem.length ? o.H = o.h % 12 == 0 ? 12 : o.h % 12 : o.H = o.h, o.HH = (o.H < 10 ? "0" : "") + o.H, o.P = o.p.toUpperCase(), o.hh = (o.h < 10 ? "0" : "") + o.h, o.ii = (o.i < 10 ? "0" : "") + o.i, o.ss = (o.s < 10 ? "0" : "") + o.s, o.dd = (o.d < 10 ? "0" : "") + o.d, o.mm = (o.m < 10 ? "0" : "") + o.m
			} else {
				if ("php" != j) {
					throw new Error("Invalid format type.")
				}
				o = {
					y: p.getUTCFullYear().toString().substring(2),
					Y: p.getUTCFullYear(),
					F: f[n].months[p.getUTCMonth()],
					M: f[n].monthsShort[p.getUTCMonth()],
					n: p.getUTCMonth() + 1,
					t: d.getDaysInMonth(p.getUTCFullYear(), p.getUTCMonth()),
					j: p.getUTCDate(),
					l: f[n].days[p.getUTCDay()],
					D: f[n].daysShort[p.getUTCDay()],
					w: p.getUTCDay(),
					N: 0 == p.getUTCDay() ? 7 : p.getUTCDay(),
					S: p.getUTCDate() % 10 <= f[n].suffix.length ? f[n].suffix[p.getUTCDate() % 10 - 1] : "",
					a: 2 == f[n].meridiem.length ? f[n].meridiem[p.getUTCHours() < 12 ? 0 : 1] : "",
					g: p.getUTCHours() % 12 == 0 ? 12 : p.getUTCHours() % 12,
					G: p.getUTCHours(),
					i: p.getUTCMinutes(),
					s: p.getUTCSeconds()
				}, o.m = (o.n < 10 ? "0" : "") + o.n, o.d = (o.j < 10 ? "0" : "") + o.j, o.A = o.a.toString().toUpperCase(), o.h = (o.g < 10 ? "0" : "") + o.g, o.H = (o.G < 10 ? "0" : "") + o.G, o.i = (o.i < 10 ? "0" : "") + o.i, o.s = (o.s < 10 ? "0" : "") + o.s
			}
			for (var p = [], g = b.extend([], k.separators), m = 0, q = k.parts.length; m < q; m++) {
				g.length && p.push(g.shift()), p.push(o[k.parts[m]])
			}
			return g.length && p.push(g.shift()), p.join("")
		},
		convertViewMode: function(e) {
			switch (e) {
			case 4:
			case "decade":
				e = 4;
				break;
			case 3:
			case "year":
				e = 3;
				break;
			case 2:
			case "month":
				e = 2;
				break;
			case 1:
			case "day":
				e = 1;
				break;
			case 0:
			case "hour":
				e = 0
			}
			return e
		},
		headTemplate: '<thead><tr><th class="prev"><i class="icon-arrow-left"/></th><th colspan="5" class="switch"></th><th class="next"><i class="icon-arrow-right"/></th></tr></thead>',
		contTemplate: '<tbody><tr><td colspan="7"></td></tr></tbody>',
		footTemplate: '<tfoot><tr><th colspan="7" class="today"></th></tr></tfoot>'
	};
	d.template = '<div class="datetimepicker"><div class="datetimepicker-minutes"><table class=" table-condensed">' + d.headTemplate + d.contTemplate + d.footTemplate + '</table></div><div class="datetimepicker-hours"><table class=" table-condensed">' + d.headTemplate + d.contTemplate + d.footTemplate + '</table></div><div class="datetimepicker-days"><table class=" table-condensed">' + d.headTemplate + "<tbody></tbody>" + d.footTemplate + '</table></div><div class="datetimepicker-months"><table class="table-condensed">' + d.headTemplate + d.contTemplate + d.footTemplate + '</table></div><div class="datetimepicker-years"><table class="table-condensed">' + d.headTemplate + d.contTemplate + d.footTemplate + "</table></div></div>", b.fn.datetimepicker.DPGlobal = d, b.fn.datetimepicker.noConflict = function() {
		return b.fn.datetimepicker = old, this
	}, b(document).on("focus.datetimepicker.data-api click.datetimepicker.data-api", '[data-provide="datetimepicker"]', function(h) {
		var g = b(this);
		g.data("datetimepicker") || (h.preventDefault(), g.datetimepicker("show"))
	}), b(function() {
		b('[data-provide="datetimepicker-inline"]').datetimepicker()
	})
}(window.jQuery), /*! bootbox.js v4.4.0 http://bootboxjs.com/license.txt */

function(a, b) {
	"function" == typeof define && define.amd ? define(["jquery"], b) : "object" == typeof exports ? module.exports = b(require("jquery")) : a.bootbox = b(a.jQuery)
}(this, function t(B, x) {
	function q(a) {
		var c = v[A.locale];
		return c ? c[a] : v.en[a]
	}
	function k(c, a, e) {
		c.stopPropagation(), c.preventDefault();
		var d = B.isFunction(e) && e.call(a, c) === !1;
		d || a.modal("hide")
	}
	function G(c) {
		var d, a = 0;
		for (d in c) {
			a++
		}
		return a
	}
	function E(c, a) {
		var d = 0;
		B.each(c, function(f, g) {
			a(f, g, d++)
		})
	}
	function b(c) {
		var a, d;
		if ("object" != typeof c) {
			throw new Error("Please supply an object of options")
		}
		if (!c.message) {
			throw new Error("Please specify a message")
		}
		return c = B.extend({}, A, c), c.buttons || (c.buttons = {}), a = c.buttons, d = G(a), E(a, function(e, g, f) {
			if (B.isFunction(g) && (g = a[e] = {
				callback: g
			}), "object" !== B.type(g)) {
				throw new Error("button with key " + e + " must be an object")
			}
			g.label || (g.label = e), g.className || (2 === d && ("ok" === e || "confirm" === e) || 1 === d ? g.className = "btn-primary" : g.className = "btn-default")
		}), c
	}
	function w(c, d) {
		var a = c.length,
			f = {};
		if (a < 1 || a > 2) {
			throw new Error("Invalid argument length")
		}
		return 2 === a || "string" == typeof c[0] ? (f[d[0]] = c[0], f[d[1]] = c[1]) : f = c[0], f
	}
	function y(c, a, d) {
		return B.extend(!0, {}, c, w(a, d))
	}
	function D(c, d, a, g) {
		var f = {
			className: "bootbox-" + c,
			buttons: C.apply(null, d)
		};
		return F(y(f, g, a), d)
	}
	function C() {
		for (var f = {}, h = 0, d = arguments.length; h < d; h++) {
			var l = arguments[h],
				g = l.toLowerCase(),
				c = l.toUpperCase();
			f[g] = {
				label: q(c)
			}
		}
		return f
	}
	function F(a, c) {
		var d = {};
		return E(c, function(f, g) {
			d[g] = !0
		}), E(a.buttons, function(e) {
			if (d[e] === x) {
				throw new Error("button key " + e + " is not allowed (options are " + c.join("\n") + ")")
			}
		}), a
	}
	var j = {
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
		A = {
			locale: B.zui && B.zui.clientLang ? B.zui.clientLang() : "zh_cn",
			backdrop: "static",
			animate: !0,
			className: null,
			closeButton: !0,
			show: !0,
			container: "body"
		},
		z = {};
	z.alert = function() {
		var a;
		if (a = D("alert", ["ok"], ["message", "callback"], arguments), a.callback && !B.isFunction(a.callback)) {
			throw new Error("alert requires callback property to be a function when provided")
		}
		return a.buttons.ok.callback = a.onEscape = function() {
			return !B.isFunction(a.callback) || a.callback.call(this)
		}, z.dialog(a)
	}, z.confirm = function() {
		var a;
		if (a = D("confirm", ["confirm", "cancel"], ["message", "callback"], arguments), a.buttons.cancel.callback = a.onEscape = function() {
			return a.callback.call(this, !1)
		}, a.buttons.confirm.callback = function() {
			return a.callback.call(this, !0)
		}, !B.isFunction(a.callback)) {
			throw new Error("confirm requires a callback")
		}
		return z.dialog(a)
	}, z.prompt = function() {
		var H, e, d, I, a, h, u;
		if (I = B(j.form), e = {
			className: "bootbox-prompt",
			buttons: C("cancel", "confirm"),
			value: "",
			inputType: "text"
		}, H = F(y(e, arguments, ["title", "callback"]), ["confirm", "cancel"]), h = H.show === x || H.show, H.message = I, H.buttons.cancel.callback = H.onEscape = function() {
			return H.callback.call(this, null)
		}, H.buttons.confirm.callback = function() {
			var c;
			switch (H.inputType) {
			case "text":
			case "textarea":
			case "email":
			case "select":
			case "date":
			case "time":
			case "number":
			case "password":
				c = a.val();
				break;
			case "checkbox":
				var f = a.find("input:checked");
				c = [], E(f, function(l, m) {
					c.push(B(m).val())
				})
			}
			return H.callback.call(this, c)
		}, H.show = !1, !H.title) {
			throw new Error("prompt requires a title")
		}
		if (!B.isFunction(H.callback)) {
			throw new Error("prompt requires a callback")
		}
		if (!j.inputs[H.inputType]) {
			throw new Error("invalid prompt type")
		}
		switch (a = B(j.inputs[H.inputType]), H.inputType) {
		case "text":
		case "textarea":
		case "email":
		case "date":
		case "time":
		case "number":
		case "password":
			a.val(H.value);
			break;
		case "select":
			var p = {};
			if (u = H.inputOptions || [], !B.isArray(u)) {
				throw new Error("Please pass an array of input options")
			}
			if (!u.length) {
				throw new Error("prompt with select requires options")
			}
			E(u, function(c, l) {
				var f = a;
				if (l.value === x || l.text === x) {
					throw new Error("given options in wrong format")
				}
				l.group && (p[l.group] || (p[l.group] = B("<optgroup/>").attr("label", l.group)), f = p[l.group]), f.append("<option value='" + l.value + "'>" + l.text + "</option>")
			}), E(p, function(c, f) {
				a.append(f)
			}), a.val(H.value);
			break;
		case "checkbox":
			var g = B.isArray(H.value) ? H.value : [H.value];
			if (u = H.inputOptions || [], !u.length) {
				throw new Error("prompt with checkbox requires options")
			}
			if (!u[0].value || !u[0].text) {
				throw new Error("given options in wrong format")
			}
			a = B("<div/>"), E(u, function(c, l) {
				var f = B(j.inputs[H.inputType]);
				f.find("input").attr("value", l.value), f.find("label").append(l.text), E(g, function(m, n) {
					n === l.value && f.find("input").prop("checked", !0)
				}), a.append(f)
			})
		}
		return H.placeholder && a.attr("placeholder", H.placeholder), H.pattern && a.attr("pattern", H.pattern), H.maxlength && a.attr("maxlength", H.maxlength), I.append(a), I.on("submit", function(c) {
			c.preventDefault(), c.stopPropagation(), d.find(".btn-primary").click()
		}), d = z.dialog(H), d.off("shown.zui.modal"), d.on("shown.zui.modal", function() {
			a.focus()
		}), h === !0 && d.modal("show"), d
	}, z.dialog = function(f) {
		f = b(f);
		var r = B(j.dialog),
			m = r.find(".modal-dialog"),
			a = r.find(".modal-body"),
			g = f.buttons,
			p = "",
			o = {
				onEscape: f.onEscape
			};
		if (B.fn.modal === x) {
			throw new Error("$.fn.modal is not defined; please double check you have included the Bootstrap JavaScript library. See http://getbootstrap.com/javascript/ for more details.")
		}
		if (E(g, function(c, d) {
			p += "<button data-bb-handler='" + c + "' type='button' class='btn " + d.className + "'>" + d.label + "</button>", o[c] = d.callback
		}), a.find(".bootbox-body").html(f.message), f.animate === !0 && r.addClass("fade"), f.className && r.addClass(f.className), "large" === f.size ? m.addClass("modal-lg") : "small" === f.size && m.addClass("modal-sm"), f.title && a.before(j.header), f.closeButton) {
			var e = B(j.closeButton);
			f.title ? r.find(".modal-header").prepend(e) : e.css("margin-top", "-10px").prependTo(a)
		}
		return f.title && r.find(".modal-title").html(f.title), p.length && (a.after(j.footer), r.find(".modal-footer").html(p)), r.on("hidden.zui.modal", function(c) {
			c.target === this && r.remove()
		}), r.on("shown.zui.modal", function() {
			r.find(".btn-primary:first").focus()
		}), "static" !== f.backdrop && r.on("click.dismiss.zui.modal", function(c) {
			r.children(".modal-backdrop").length && (c.currentTarget = r.children(".modal-backdrop").get(0)), c.target === c.currentTarget && r.trigger("escape.close.bb")
		}), r.on("escape.close.bb", function(c) {
			o.onEscape && k(c, r, o.onEscape)
		}), r.on("click", ".modal-footer button", function(d) {
			var c = B(this).data("bb-handler");
			k(d, r, o[c])
		}), r.on("click", ".bootbox-close-button", function(c) {
			k(c, r, o.onEscape)
		}), r.on("keyup", function(c) {
			27 === c.which && r.trigger("escape.close.bb")
		}), B(f.container).append(r), r.modal({
			backdrop: !! f.backdrop && "static",
			keyboard: !1,
			show: !1
		}), f.show && r.modal("show"), r
	}, z.setDefaults = function() {
		var a = {};
		2 === arguments.length ? a[arguments[0]] = arguments[1] : a = arguments[0], B.extend(A, a)
	}, z.hideAll = function() {
		return B(".bootbox").modal("hide"), z
	};
	var v = {
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
	return z.addLocale = function(c, a) {
		return B.each(["OK", "CANCEL", "CONFIRM"], function(d, f) {
			if (!a[f]) {
				throw new Error("Please supply a translation for '" + f + "'")
			}
		}), v[c] = {
			OK: a.OK,
			CANCEL: a.CANCEL,
			CONFIRM: a.CONFIRM
		}, z
	}, z.removeLocale = function(a) {
		return delete v[a], z
	}, z.setLocale = function(a) {
		return z.setDefaults("locale", a)
	}, z.init = function(a) {
		return t(a || B)
	}, z
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
	var d, h, c, k, j, f = {}.hasOwnProperty,
		b = function(l, m) {
			function a() {
				this.constructor = l
			}
			for (var o in m) {
				f.call(m, o) && (l[o] = m[o])
			}
			return a.prototype = m.prototype, l.prototype = new a, l.__super__ = m.prototype, l
		},
		g = {
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
	k = function() {
		function a() {
			this.options_index = 0, this.parsed = []
		}
		return a.prototype.add_node = function(e) {
			return "OPTGROUP" === e.nodeName.toUpperCase() ? this.add_group(e) : this.add_option(e)
		}, a.prototype.add_group = function(u) {
			var m, w, v, p, l, q;
			for (m = this.parsed.length, this.parsed.push({
				array_index: m,
				group: !0,
				label: this.escapeExpression(u.label),
				children: 0,
				disabled: u.disabled,
				title: u.title,
				search_keys: d.trim(u.getAttribute("data-keys") || "").replace(/,/g, " ")
			}), l = u.childNodes, q = [], v = 0, p = l.length; v < p; v++) {
				w = l[v], q.push(this.add_option(w, m, u.disabled))
			}
			return q
		}, a.prototype.add_option = function(m, l, o) {
			if ("OPTION" === m.nodeName.toUpperCase()) {
				return "" !== m.text ? (null != l && (this.parsed[l].children += 1), this.parsed.push({
					array_index: this.parsed.length,
					options_index: this.options_index,
					value: m.value,
					text: m.text,
					title: m.title,
					html: m.innerHTML,
					selected: m.selected,
					disabled: o === !0 ? o : m.disabled,
					group_array_index: l,
					classes: m.className,
					style: m.style.cssText,
					data: m.getAttribute("data-data"),
					search_keys: (d.trim(m.getAttribute("data-keys") || "") + m.value).replace(/,/, " ")
				})) : this.parsed.push({
					array_index: this.parsed.length,
					options_index: this.options_index,
					empty: !0
				}), this.options_index += 1
			}
		}, a.prototype.escapeExpression = function(m) {
			var n, l;
			return null == m || m === !1 ? "" : /[\&\<\>\"\'\`]/.test(m) ? (n = {
				"<": "&lt;",
				">": "&gt;",
				'"': "&quot;",
				"'": "&#x27;",
				"`": "&#x60;"
			}, l = /&(?!\w+;)|[\<\>\"\'\`]/g, m.replace(l, function(e) {
				return n[e] || "&amp;"
			})) : m
		}, a
	}(), k.select_to_array = function(n) {
		var q, m, r, p, l;
		for (m = new k, l = n.childNodes, r = 0, p = l.length; r < p; r++) {
			q = l[r], m.add_node(q)
		}
		return m.parsed
	}, h = function() {
		function a(e, l) {
			this.form_field = e, this.options = null != l ? l : {}, a.browser_is_supported() && (this.lang = g[this.options.lang || (d.zui.clientLang ? d.zui.clientLang() : "zh_cn")], this.is_multiple = this.form_field.multiple, this.set_default_text(), this.set_default_values(), this.setup(), this.set_up_html(), this.register_observers())
		}
		return a.prototype.set_default_values = function() {
			var e = this;
			return this.click_test_action = function(l) {
				return e.test_active_click(l)
			}, this.activate_action = function(l) {
				return e.activate_field(l)
			}, this.active_field = !1, this.mouse_on_container = !1, this.results_showing = !1, this.result_highlighted = null, this.allow_single_deselect = null != this.options.allow_single_deselect && null != this.form_field.options[0] && "" === this.form_field.options[0].text && this.options.allow_single_deselect, this.disable_search_threshold = this.options.disable_search_threshold || 0, this.disable_search = this.options.disable_search || !1, this.enable_split_word_search = null == this.options.enable_split_word_search || this.options.enable_split_word_search, this.group_search = null == this.options.group_search || this.options.group_search, this.search_contains = this.options.search_contains || !1, this.single_backstroke_delete = null == this.options.single_backstroke_delete || this.options.single_backstroke_delete, this.max_selected_options = this.options.max_selected_options || 1 / 0, this.drop_direction = this.options.drop_direction || "auto", this.middle_highlight = this.options.middle_highlight, this.compact_search = this.options.compact_search || !0, this.inherit_select_classes = this.options.inherit_select_classes || !1, this.display_selected_options = null == this.options.display_selected_options || this.options.display_selected_options, this.display_disabled_options = null == this.options.display_disabled_options || this.options.display_disabled_options
		}, a.prototype.set_default_text = function() {
			return this.form_field.getAttribute("data-placeholder") ? this.default_text = this.form_field.getAttribute("data-placeholder") : this.is_multiple ? this.default_text = this.options.placeholder_text_multiple || this.options.placeholder_text || a.default_multiple_text : this.default_text = this.options.placeholder_text_single || this.options.placeholder_text || a.default_single_text, this.results_none_found = this.form_field.getAttribute("data-no_results_text") || this.options.no_results_text || this.lang.no_results_text || a.default_no_result_text
		}, a.prototype.mouse_enter = function() {
			return this.mouse_on_container = !0
		}, a.prototype.mouse_leave = function() {
			return this.mouse_on_container = !1
		}, a.prototype.input_focus = function(l) {
			var m = this;
			if (this.is_multiple) {
				if (!this.active_field) {
					return setTimeout(function() {
						return m.container_mousedown()
					}, 50)
				}
			} else {
				if (!this.active_field) {
					return this.activate_field()
				}
			}
		}, a.prototype.input_blur = function(l) {
			var m = this;
			if (!this.mouse_on_container) {
				return this.active_field = !1, setTimeout(function() {
					return m.blur_test()
				}, 100)
			}
		}, a.prototype.results_option_build = function(m) {
			var q, l, u, r, p;
			for (q = "", p = this.results_data, u = 0, r = p.length; u < r; u++) {
				l = p[u], q += l.group ? this.result_add_group(l) : this.result_add_option(l), (null != m ? m.first : void 0) && (l.selected && this.is_multiple ? this.choice_build(l) : l.selected && !this.is_multiple && this.single_set_selected_text(l.text))
			}
			return q
		}, a.prototype.result_add_option = function(m) {
			var n, l;
			return m.search_match && this.include_option_in_results(m) ? (n = [], m.disabled || m.selected && this.is_multiple || n.push("active-result"), !m.disabled || m.selected && this.is_multiple || n.push("disabled-result"), m.selected && n.push("result-selected"), null != m.group_array_index && n.push("group-option"), "" !== m.classes && n.push(m.classes), l = document.createElement("li"), l.className = n.join(" "), l.style.cssText = m.style, l.title = m.title, l.setAttribute("data-option-array-index", m.array_index), l.setAttribute("data-data", m.data), l.innerHTML = m.search_text, this.outerHTML(l)) : ""
		}, a.prototype.result_add_group = function(l) {
			var m;
			return (l.search_match || l.group_match) && l.active_options > 0 ? (m = document.createElement("li"), m.className = "group-result", m.title = l.title, m.innerHTML = l.search_text, this.outerHTML(m)) : ""
		}, a.prototype.results_update_field = function() {
			if (this.set_default_text(), this.is_multiple || this.results_reset_cleanup(), this.result_clear_highlight(), this.results_build(), this.results_showing) {
				return this.winnow_results()
			}
		}, a.prototype.reset_single_select_options = function() {
			var m, p, l, r, q;
			for (r = this.results_data, q = [], p = 0, l = r.length; p < l; p++) {
				m = r[p], m.selected ? q.push(m.selected = !1) : q.push(void 0)
			}
			return q
		}, a.prototype.results_toggle = function() {
			return this.results_showing ? this.results_hide() : this.results_show()
		}, a.prototype.results_search = function(e) {
			return this.results_showing ? this.winnow_results(1) : this.results_show()
		}, a.prototype.winnow_results = function(F) {
			var A, y, w, v, G, D, m, x, z, C, B, E, q;
			for (this.no_results_clear(), G = 0, m = this.get_search_text(), A = m.replace(/[-[\]{}()*+?.,\\^$|#\s]/g, "\\$&"), v = this.search_contains ? "" : "^", w = new RegExp(v + A, "i"), C = new RegExp(A, "i"), q = this.results_data, B = 0, E = q.length; B < E; B++) {
				y = q[B], y.search_match = !1, D = null, this.include_option_in_results(y) && (y.group && (y.group_match = !1, y.active_options = 0), null != y.group_array_index && this.results_data[y.group_array_index] && (D = this.results_data[y.group_array_index], 0 === D.active_options && D.search_match && (G += 1), D.active_options += 1), y.group && !this.group_search || (y.search_text = y.group ? y.label : y.html, y.search_keys_match = this.search_string_match(y.search_keys, w), y.search_text_match = this.search_string_match(y.search_text, w), y.search_match = y.search_text_match || y.search_keys_match, y.search_match && !y.group && (G += 1), y.search_match ? (y.search_text_match && y.search_text.length ? (x = y.search_text.search(C), z = y.search_text.substr(0, x + m.length) + "</em>" + y.search_text.substr(x + m.length), y.search_text = z.substr(0, x) + "<em>" + z.substr(x)) : y.search_keys_match && y.search_keys.length && (x = y.search_keys.search(C), z = y.search_keys.substr(0, x + m.length) + "</em>" + y.search_keys.substr(x + m.length), y.search_text += '&nbsp; <small style="opacity: 0.7">' + z.substr(0, x) + "<em>" + z.substr(x) + "</small>"), null != D && (D.group_match = !0)) : null != y.group_array_index && this.results_data[y.group_array_index].search_match && (y.search_match = !0)))
			}
			return this.result_clear_highlight(), G < 1 && m.length ? (this.update_results_content(""), this.no_results(m)) : (this.update_results_content(this.results_option_build()), this.winnow_results_set_highlight(F))
		}, a.prototype.search_string_match = function(m, q) {
			var l, u, r, p;
			if (q.test(m)) {
				return !0
			}
			if (this.enable_split_word_search && (m.indexOf(" ") >= 0 || 0 === m.indexOf("[")) && (u = m.replace(/\[|\]/g, "").split(" "), u.length)) {
				for (r = 0, p = u.length; r < p; r++) {
					if (l = u[r], q.test(l)) {
						return !0
					}
				}
			}
		}, a.prototype.choices_count = function() {
			var m, o, l, p;
			if (null != this.selected_option_count) {
				return this.selected_option_count
			}
			for (this.selected_option_count = 0, p = this.form_field.options, o = 0, l = p.length; o < l; o++) {
				m = p[o], m.selected && "" != m.value && (this.selected_option_count += 1)
			}
			return this.selected_option_count
		}, a.prototype.choices_click = function(e) {
			if (e.preventDefault(), !this.results_showing && !this.is_disabled) {
				return this.results_show()
			}
		}, a.prototype.keyup_checker = function(m) {
			var n, l;
			switch (n = null != (l = m.which) ? l : m.keyCode, this.search_field_scale(), n) {
			case 8:
				if (this.is_multiple && this.backstroke_length < 1 && this.choices_count() > 0) {
					return this.keydown_backstroke()
				}
				if (!this.pending_backstroke) {
					return this.result_clear_highlight(), this.results_search()
				}
				break;
			case 13:
				if (m.preventDefault(), this.results_showing) {
					return this.result_select(m)
				}
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
		}, a.prototype.clipboard_event_checker = function(l) {
			var m = this;
			return setTimeout(function() {
				return m.results_search()
			}, 50)
		}, a.prototype.container_width = function() {
			return null != this.options.width ? this.options.width : this.form_field && this.form_field.classList && this.form_field.classList.contains("form-control") ? "100%" : "" + this.form_field.offsetWidth + "px"
		}, a.prototype.include_option_in_results = function(e) {
			return !(this.is_multiple && !this.display_selected_options && e.selected) && (!(!this.display_disabled_options && e.disabled) && !e.empty)
		}, a.prototype.search_results_touchstart = function(e) {
			return this.touch_started = !0, this.search_results_mouseover(e)
		}, a.prototype.search_results_touchmove = function(e) {
			return this.touch_started = !1, this.search_results_mouseout(e)
		}, a.prototype.search_results_touchend = function(e) {
			if (this.touch_started) {
				return this.search_results_mouseup(e)
			}
		}, a.prototype.outerHTML = function(l) {
			var m;
			return l.outerHTML ? l.outerHTML : (m = document.createElement("div"), m.appendChild(l), m.innerHTML)
		}, a.browser_is_supported = function() {
			return "Microsoft Internet Explorer" === window.navigator.appName ? document.documentMode >= 8 : !/iP(od|hone)/i.test(window.navigator.userAgent) && (!/Android/i.test(window.navigator.userAgent) || !/Mobile/i.test(window.navigator.userAgent))
		}, a.default_multiple_text = "", a.default_single_text = "", a.default_no_result_text = "No results match", a
	}(), d = jQuery, d.fn.extend({
		chosen: function(a) {
			return h.browser_is_supported() ? this.each(function(m) {
				var n, l;
				n = d(this), l = n.data("chosen"), "destroy" === a && l ? l.destroy() : l || n.data("chosen", new c(this, a))
			}) : this
		}
	}), c = function(l) {
		function a() {
			return j = a.__super__.constructor.apply(this, arguments)
		}
		return b(a, l), a.prototype.setup = function() {
			return this.form_field_jq = d(this.form_field), this.current_selectedIndex = this.form_field.selectedIndex, this.is_rtl = this.form_field_jq.hasClass("chosen-rtl")
		}, a.prototype.set_up_html = function() {
			var o, m;
			o = ["chosen-container"], o.push("chosen-container-" + (this.is_multiple ? "multi" : "single")), this.inherit_select_classes && this.form_field.className && o.push(this.form_field.className), this.is_rtl && o.push("chosen-rtl");
			var p = this.form_field.getAttribute("data-css-class");
			return p && o.push(p), m = {
				"class": o.join(" "),
				style: "width: " + this.container_width() + ";",
				title: this.form_field.title
			}, this.form_field.id.length && (m.id = this.form_field.id.replace(/[^\w]/g, "_") + "_chosen"), this.container = d("<div />", m), this.is_multiple ? this.container.html('<ul class="chosen-choices"><li class="search-field"><input type="text" value="' + this.default_text + '" class="default" autocomplete="off" style="width:25px;" /></li></ul><div class="chosen-drop"><ul class="chosen-results"></ul></div>') : (this.container.html('<a class="chosen-single chosen-default" tabindex="-1"><span>' + this.default_text + '</span><div><b></b></div><div class="chosen-search"><input type="text" autocomplete="off" /></div></a><div class="chosen-drop"><ul class="chosen-results"></ul></div>'), this.compact_search && this.container.find(".chosen-search").appendTo(this.container.find(".chosen-single"))), this.form_field_jq.hide().after(this.container), this.dropdown = this.container.find("div.chosen-drop").first(), this.search_field = this.container.find("input").first(), this.search_results = this.container.find("ul.chosen-results").first(), this.search_field_scale(), this.search_no_results = this.container.find("li.no-results").first(), this.is_multiple ? (this.search_choices = this.container.find("ul.chosen-choices").first(), this.search_container = this.container.find("li.search-field").first()) : (this.search_container = this.container.find("div.chosen-search").first(), this.selected_item = this.container.find(".chosen-single").first()), this.options.drop_width && this.dropdown.css("width", this.options.drop_width).addClass("chosen-drop-size-limited"), this.results_build(), this.set_tab_index(), this.set_label_behavior(), this.form_field_jq.trigger("chosen:ready", {
				chosen: this
			})
		}, a.prototype.register_observers = function() {
			var e = this;
			return this.container.bind("mousedown.chosen", function(m) {
				e.container_mousedown(m)
			}), this.container.bind("mouseup.chosen", function(m) {
				e.container_mouseup(m)
			}), this.container.bind("mouseenter.chosen", function(m) {
				e.mouse_enter(m)
			}), this.container.bind("mouseleave.chosen", function(m) {
				e.mouse_leave(m)
			}), this.search_results.bind("mouseup.chosen", function(m) {
				e.search_results_mouseup(m)
			}), this.search_results.bind("mouseover.chosen", function(m) {
				e.search_results_mouseover(m)
			}), this.search_results.bind("mouseout.chosen", function(m) {
				e.search_results_mouseout(m)
			}), this.search_results.bind("mousewheel.chosen DOMMouseScroll.chosen", function(m) {
				e.search_results_mousewheel(m)
			}), this.search_results.bind("touchstart.chosen", function(m) {
				e.search_results_touchstart(m)
			}), this.search_results.bind("touchmove.chosen", function(m) {
				e.search_results_touchmove(m)
			}), this.search_results.bind("touchend.chosen", function(m) {
				e.search_results_touchend(m)
			}), this.form_field_jq.bind("chosen:updated.chosen", function(m) {
				e.results_update_field(m)
			}), this.form_field_jq.bind("chosen:activate.chosen", function(m) {
				e.activate_field(m)
			}), this.form_field_jq.bind("chosen:open.chosen", function(m) {
				e.container_mousedown(m)
			}), this.form_field_jq.bind("chosen:close.chosen", function(m) {
				e.input_blur(m)
			}), this.search_field.bind("blur.chosen", function(m) {
				e.input_blur(m)
			}), this.search_field.bind("keyup.chosen", function(m) {
				e.keyup_checker(m)
			}), this.search_field.bind("keydown.chosen", function(m) {
				e.keydown_checker(m)
			}), this.search_field.bind("focus.chosen", function(m) {
				e.input_focus(m)
			}), this.search_field.bind("cut.chosen", function(m) {
				e.clipboard_event_checker(m)
			}), this.search_field.bind("paste.chosen", function(m) {
				e.clipboard_event_checker(m)
			}), this.is_multiple ? this.search_choices.bind("click.chosen", function(m) {
				e.choices_click(m)
			}) : this.container.bind("click.chosen", function(m) {
				m.preventDefault()
			})
		}, a.prototype.destroy = function() {
			return d(this.container[0].ownerDocument).unbind("click.chosen", this.click_test_action), this.search_field[0].tabIndex && (this.form_field_jq[0].tabIndex = this.search_field[0].tabIndex), this.container.remove(), this.form_field_jq.removeData("chosen"), this.form_field_jq.show()
		}, a.prototype.search_field_disabled = function() {
			return this.is_disabled = this.form_field_jq[0].disabled, this.is_disabled ? (this.container.addClass("chosen-disabled"), this.search_field[0].disabled = !0, this.is_multiple || this.selected_item.unbind("focus.chosen", this.activate_action), this.close_field()) : (this.container.removeClass("chosen-disabled"), this.search_field[0].disabled = !1, this.is_multiple ? void 0 : this.selected_item.bind("focus.chosen", this.activate_action))
		}, a.prototype.container_mousedown = function(m) {
			if (!this.is_disabled && (m && "mousedown" === m.type && !this.results_showing && m.preventDefault(), null == m || !d(m.target).hasClass("search-choice-close"))) {
				return this.active_field ? this.is_multiple || !m || d(m.target)[0] !== this.selected_item[0] && !d(m.target).parents("a.chosen-single").length || (m.preventDefault(), this.results_toggle()) : (this.is_multiple && this.search_field.val(""), d(this.container[0].ownerDocument).bind("click.chosen", this.click_test_action), this.results_show()), this.activate_field()
			}
		}, a.prototype.container_mouseup = function(e) {
			if ("ABBR" === e.target.nodeName && !this.is_disabled) {
				return this.results_reset(e)
			}
		}, a.prototype.search_results_mousewheel = function(m) {
			var n;
			if (m.originalEvent && (n = -m.originalEvent.wheelDelta || m.originalEvent.detail), null != n) {
				return m.preventDefault(), "DOMMouseScroll" === m.type && (n = 40 * n), this.search_results.scrollTop(n + this.search_results.scrollTop())
			}
		}, a.prototype.blur_test = function(e) {
			if (!this.active_field && this.container.hasClass("chosen-container-active")) {
				return this.close_field()
			}
		}, a.prototype.close_field = function() {
			return d(this.container[0].ownerDocument).unbind("click.chosen", this.click_test_action), this.active_field = !1, this.results_hide(), this.container.removeClass("chosen-container-active"), this.clear_backstroke(), this.show_search_field_default(), this.search_field_scale()
		}, a.prototype.activate_field = function() {
			return this.container.addClass("chosen-container-active"), this.active_field = !0, this.search_field.val(this.search_field.val()), this.search_field.focus()
		}, a.prototype.test_active_click = function(n) {
			var m;
			return m = d(n.target).closest(".chosen-container"), m.length && this.container[0] === m[0] ? this.active_field = !0 : this.close_field()
		}, a.prototype.results_build = function() {
			return this.parsing = !0, this.selected_option_count = null, this.results_data = k.select_to_array(this.form_field), this.is_multiple ? this.search_choices.find("li.search-choice").remove() : this.is_multiple || (this.single_set_selected_text(), this.disable_search || this.form_field.options.length <= this.disable_search_threshold ? (this.search_field[0].readOnly = !0, this.container.addClass("chosen-container-single-nosearch"), this.container.removeClass("chosen-with-search")) : (this.search_field[0].readOnly = !1, this.container.removeClass("chosen-container-single-nosearch"), this.container.addClass("chosen-with-search"))), this.update_results_content(this.results_option_build({
				first: !0
			})), this.search_field_disabled(), this.show_search_field_default(), this.search_field_scale(), this.parsing = !1
		}, a.prototype.result_do_highlight = function(y, w) {
			var v, q, p, z, x, m, u = -1;
			y.length && (this.result_clear_highlight(), this.result_highlight = y, this.result_highlight.addClass("highlighted"), p = parseInt(this.search_results.css("maxHeight"), 10), m = this.result_highlight.outerHeight(), x = this.search_results.scrollTop(), z = p + x, q = this.result_highlight.position().top + this.search_results.scrollTop(), v = q + m, this.middle_highlight && (w || "always" === this.middle_highlight || v >= z || q < x) ? u = Math.min(q - m, Math.max(0, q - (p - m) / 2)) : v >= z ? u = v - p > 0 ? v - p : 0 : q < x && (u = q), u > -1 && this.search_results.scrollTop(u))
		}, a.prototype.result_clear_highlight = function() {
			return this.result_highlight && this.result_highlight.removeClass("highlighted"), this.result_highlight = null
		}, a.prototype.results_show = function() {
			if (this.is_multiple && this.max_selected_options <= this.choices_count()) {
				return this.form_field_jq.trigger("chosen:maxselected", {
					chosen: this
				}), !1
			}
			this.results_showing = !0, this.search_field.focus(), this.search_field.val(this.search_field.val()), this.winnow_results(1);
			var o = this.drop_direction;
			if (d.isFunction(o) && (o = o.call(this)), "auto" === o) {
				if (this.drop_directionFixed) {
					o = this.drop_directionFixed
				} else {
					var m = this.container.find(".chosen-drop"),
						p = this.container.offset();
					p.top + m.outerHeight() + 30 > d(window).height() + d(window).scrollTop() && (o = "up"), this.drop_directionFixed = o
				}
			}
			return this.container.toggleClass("chosen-up", "up" === o).addClass("chosen-with-drop"), this.form_field_jq.trigger("chosen:showing_dropdown", {
				chosen: this
			})
		}, a.prototype.update_results_content = function(e) {
			return this.search_results.html(e)
		}, a.prototype.results_hide = function() {
			return this.results_showing && (this.result_clear_highlight(), this.container.removeClass("chosen-with-drop"), this.form_field_jq.trigger("chosen:hiding_dropdown", {
				chosen: this
			}), this.drop_directionFixed = 0), this.results_showing = !1
		}, a.prototype.set_tab_index = function(m) {
			var n;
			if (this.form_field.tabIndex) {
				return n = this.form_field.tabIndex, this.form_field.tabIndex = -1, this.search_field[0].tabIndex = n
			}
		}, a.prototype.set_label_behavior = function() {
			var m = this;
			if (this.form_field_label = this.form_field_jq.parents("label"), !this.form_field_label.length && this.form_field.id.length && (this.form_field_label = d("label[for='" + this.form_field.id + "']")), this.form_field_label.length > 0) {
				return this.form_field_label.bind("click.chosen", function(e) {
					return m.is_multiple ? m.container_mousedown(e) : m.activate_field()
				})
			}
		}, a.prototype.show_search_field_default = function() {
			return this.is_multiple && this.choices_count() < 1 && !this.active_field ? (this.search_field.val(this.default_text), this.search_field.addClass("default")) : (this.search_field.val(""), this.search_field.removeClass("default"))
		}, a.prototype.search_results_mouseup = function(n) {
			var m;
			if (m = d(n.target).hasClass("active-result") ? d(n.target) : d(n.target).parents(".active-result").first(), m.length) {
				return this.result_highlight = m, this.result_select(n), this.search_field.focus()
			}
		}, a.prototype.search_results_mouseover = function(n) {
			var m;
			if (m = d(n.target).hasClass("active-result") ? d(n.target) : d(n.target).parents(".active-result").first()) {
				return this.result_do_highlight(m)
			}
		}, a.prototype.search_results_mouseout = function(m) {
			if (d(m.target).hasClass("active-result")) {
				return this.result_clear_highlight()
			}
		}, a.prototype.choice_build = function(p) {
			var m, r, q = this;
			return m = d("<li />", {
				"class": "search-choice"
			}).html("<span title='" + p.html + "'>" + p.html + "</span>"), p.disabled ? m.addClass("search-choice-disabled") : (r = d("<a />", {
				"class": "search-choice-close",
				"data-option-array-index": p.array_index
			}), r.bind("click.chosen", function(e) {
				return q.choice_destroy_link_click(e)
			}), m.append(r)), this.search_container.before(m)
		}, a.prototype.choice_destroy_link_click = function(m) {
			if (m.preventDefault(), m.stopPropagation(), !this.is_disabled) {
				return this.choice_destroy(d(m.target))
			}
		}, a.prototype.choice_destroy = function(e) {
			if (this.result_deselect(e[0].getAttribute("data-option-array-index"))) {
				return this.show_search_field_default(), this.is_multiple && this.choices_count() > 0 && this.search_field.val().length < 1 && this.results_hide(), e.parents("li").first().remove(), this.search_field_scale()
			}
		}, a.prototype.results_reset = function() {
			if (this.reset_single_select_options(), this.form_field.options[0].selected = !0, this.single_set_selected_text(), this.show_search_field_default(), this.results_reset_cleanup(), this.form_field_jq.trigger("change"), this.active_field) {
				return this.results_hide()
			}
		}, a.prototype.results_reset_cleanup = function() {
			return this.current_selectedIndex = this.form_field.selectedIndex, this.selected_item.find("abbr").remove()
		}, a.prototype.result_select = function(n) {
			var o, m;
			if (this.result_highlight) {
				return o = this.result_highlight, this.result_clear_highlight(), this.is_multiple && this.max_selected_options <= this.choices_count() ? (this.form_field_jq.trigger("chosen:maxselected", {
					chosen: this
				}), !1) : (this.is_multiple ? o.removeClass("active-result") : this.reset_single_select_options(), m = this.results_data[o[0].getAttribute("data-option-array-index")], m.selected = !0, this.form_field.options[m.options_index].selected = !0, this.selected_option_count = null, this.is_multiple ? this.choice_build(m) : this.single_set_selected_text(m.text), (n.metaKey || n.ctrlKey) && this.is_multiple || this.results_hide(), this.search_field.val(""), (this.is_multiple || this.form_field.selectedIndex !== this.current_selectedIndex) && this.form_field_jq.trigger("change", {
					selected: this.form_field.options[m.options_index].value
				}), this.current_selectedIndex = this.form_field.selectedIndex, this.search_field_scale())
			}
		}, a.prototype.single_set_selected_text = function(e) {
			return null == e && (e = this.default_text), e === this.default_text ? this.selected_item.addClass("chosen-default") : (this.single_deselect_control_build(), this.selected_item.removeClass("chosen-default")), this.compact_search && this.search_field.attr("placeholder", e), this.selected_item.find("span").attr("title", e).text(e)
		}, a.prototype.result_deselect = function(m) {
			var n;
			return n = this.results_data[m], !this.form_field.options[n.options_index].disabled && (n.selected = !1, this.form_field.options[n.options_index].selected = !1, this.selected_option_count = null, this.result_clear_highlight(), this.results_showing && this.winnow_results(), this.form_field_jq.trigger("change", {
				deselected: this.form_field.options[n.options_index].value
			}), this.search_field_scale(), !0)
		}, a.prototype.single_deselect_control_build = function() {
			if (this.allow_single_deselect) {
				return this.selected_item.find("abbr").length || this.selected_item.find("span").first().after('<abbr class="search-choice-close"></abbr>'), this.selected_item.addClass("chosen-single-with-deselect")
			}
		}, a.prototype.get_search_text = function() {
			return this.search_field.val() === this.default_text ? "" : d("<div/>").text(d.trim(this.search_field.val())).html()
		}, a.prototype.winnow_results_set_highlight = function(n) {
			var o, m;
			if (m = this.is_multiple ? [] : this.search_results.find(".result-selected.active-result"), o = m.length ? m.first() : this.search_results.find(".active-result").first(), null != o) {
				return this.result_do_highlight(o, n)
			}
		}, a.prototype.no_results = function(n) {
			var m;
			return m = d('<li class="no-results">' + this.results_none_found + ' "<span></span>"</li>'), m.find("span").first().html(n), this.search_results.append(m), this.form_field_jq.trigger("chosen:no_results", {
				chosen: this
			})
		}, a.prototype.no_results_clear = function() {
			return this.search_results.find(".no-results").remove()
		}, a.prototype.keydown_arrow = function() {
			var e;
			return this.results_showing && this.result_highlight ? (e = this.result_highlight.nextAll("li.active-result").first()) ? this.result_do_highlight(e) : void 0 : this.results_show()
		}, a.prototype.keyup_arrow = function() {
			var e;
			return this.results_showing || this.is_multiple ? this.result_highlight ? (e = this.result_highlight.prevAll("li.active-result"), e.length ? this.result_do_highlight(e.first()) : (this.choices_count() > 0 && this.results_hide(), this.result_clear_highlight())) : void 0 : this.results_show()
		}, a.prototype.keydown_backstroke = function() {
			var e;
			return this.pending_backstroke ? (this.choice_destroy(this.pending_backstroke.find("a").first()), this.clear_backstroke()) : (e = this.search_container.siblings("li.search-choice").last(), e.length && !e.hasClass("search-choice-disabled") ? (this.pending_backstroke = e, this.single_backstroke_delete ? this.keydown_backstroke() : this.pending_backstroke.addClass("search-choice-focus")) : void 0)
		}, a.prototype.clear_backstroke = function() {
			return this.pending_backstroke && this.pending_backstroke.removeClass("search-choice-focus"), this.pending_backstroke = null
		}, a.prototype.keydown_checker = function(n) {
			var o, m;
			switch (o = null != (m = n.which) ? m : n.keyCode, this.search_field_scale(), 8 !== o && this.pending_backstroke && this.clear_backstroke(), o) {
			case 8:
				this.backstroke_length = this.search_field.val().length;
				break;
			case 9:
				this.results_showing && !this.is_multiple && this.result_select(n), this.mouse_on_container = !1;
				break;
			case 13:
				n.preventDefault();
				break;
			case 38:
				n.preventDefault(), this.keyup_arrow();
				break;
			case 40:
				n.preventDefault(), this.keydown_arrow()
			}
		}, a.prototype.search_field_scale = function() {
			var x, v, q, p, z, y, m, u, w;
			if (this.is_multiple) {
				for (q = 0, m = 0, z = "position:absolute; left: -1000px; top: -1000px; display:none;", y = ["font-size", "font-style", "font-weight", "font-family", "line-height", "text-transform", "letter-spacing"], u = 0, w = y.length; u < w; u++) {
					p = y[u], z += p + ":" + this.search_field.css(p) + ";"
				}
				return x = d("<div />", {
					style: z
				}), x.text(this.search_field.val()), d("body").append(x), m = x.width() + 25, x.remove(), v = this.container.outerWidth(), m > v - 10 && (m = v - 10), this.search_field.css({
					width: m + "px"
				})
			}
		}, a
	}(h)
}.call(this), function(b) {
	var c = "zui.selectable",
		a = function(e, g) {
			this.name = c, this.$ = b(e), this.id = b.zui.uuid(), this.selectOrder = 1, this.selections = {}, this.getOptions(g), this._init()
		},
		f = function(h, j, g) {
			return h >= g.left && h <= g.left + g.width && j >= g.top && j <= g.top + g.height
		},
		d = function(j, l) {
			var h = Math.max(j.left, l.left),
				m = Math.max(j.top, l.top),
				k = Math.min(j.left + j.width, l.left + l.width),
				g = Math.min(j.top + j.height, l.top + l.height);
			return f(h, m, j) && f(k, g, j) && f(h, m, l) && f(k, g, l)
		};
	a.DEFAULTS = {
		selector: "li,tr,div",
		trigger: "",
		selectClass: "active",
		rangeStyle: {
			border: "1px solid " + (b.zui.colorset ? b.zui.colorset.primary : "#3280fc"),
			backgroundColor: b.zui.colorset ? new b.zui.Color(b.zui.colorset.primary).fade(20).toCssStr() : "rgba(50, 128, 252, 0.2)"
		},
		clickBehavior: "toggle",
		ignoreVal: 3,
		listenClick: !0
	}, a.prototype.getOptions = function(g) {
		this.options = b.extend({}, a.DEFAULTS, this.$.data(), g)
	}, a.prototype.select = function(e) {
		this.toggle(e, !0)
	}, a.prototype.unselect = function(e) {
		this.toggle(e, !1)
	}, a.prototype.toggle = function(p, j, u) {
		var q, k, h = this.options.selector,
			m = this;
		if (void 0 === p) {
			return void this.$.find(h).each(function() {
				m.toggle(this, j)
			})
		}
		if ("object" == typeof p ? (q = b(p).closest(h), k = q.data("id")) : (k = p, q = m.$.find('.slectable-item[data-id="' + k + '"]')), q && q.length) {
			if (k || (k = b.zui.uuid(), q.attr("data-id", k)), void 0 !== j && null !== j || (j = !m.selections[k]), !! j != !! m.selections[k]) {
				var g;
				b.isFunction(u) && (g = u(j)), g !== !0 && (m.selections[k] = !! j && m.selectOrder++, m.callEvent(j ? "select" : "unselect", {
					id: k,
					selections: m.selections,
					target: q,
					selected: m.getSelectedArray()
				}, m))
			}
			m.options.selectClass && q.toggleClass(m.options.selectClass, j)
		}
	}, a.prototype.getSelectedArray = function() {
		var g = [];
		return b.each(this.selections, function(h, e) {
			e && g.push(h)
		}), g
	}, a.prototype.syncSelectionsFromClass = function() {
		var g = this;
		g.$children = g.$.find(g.options.selector);
		g.selections = {}, g.$children.each(function() {
			var e = b(this);
			g.selections[e.data("id")] = e.hasClass(g.options.selectClass)
		})
	}, a.prototype._init = function() {
		var M, I, F, A, Q, D, H, J = this.options,
			O = this,
			N = J.ignoreVal,
			z = !0,
			E = "." + this.name + "." + this.id,
			L = b.isFunction(J.checkFunc) ? J.checkFunc : null,
			K = b.isFunction(J.rangeFunc) ? J.rangeFunc : null,
			G = !1,
			q = null,
			j = "mousedown" + E,
			P = function() {
				A && O.$children.each(function() {
					var l = b(this),
						h = l.offset();
					h.width = l.outerWidth(), h.height = l.outerHeight();
					var m = K ? K.call(this, A, h) : d(A, h);
					if (L) {
						var g = L.call(O, {
							intersect: m,
							target: l,
							range: A,
							targetRange: h
						});
						g === !0 ? O.select(l) : g === !1 && O.unselect(l)
					} else {
						m ? O.select(l) : O.multiKey || O.unselect(l)
					}
				})
			},
			o = function(e) {
				G && (Q = e.pageX, D = e.pageY, A = {
					width: Math.abs(Q - M),
					height: Math.abs(D - I),
					left: Q > M ? M : Q,
					top: D > I ? I : D
				}, z && A.width < N && A.height < N || (F || (F = b('.selectable-range[data-id="' + O.id + '"]'), F.length || (F = b('<div class="selectable-range" data-id="' + O.id + '"></div>').css(b.extend({
					zIndex: 1060,
					position: "absolute",
					top: M,
					left: I,
					pointerEvents: "none"
				}, O.options.rangeStyle)).appendTo(b("body")))), F.css(A), clearTimeout(H), H = setTimeout(P, 10), z = !1))
			},
			k = function(g) {
				b(document).off(E), clearTimeout(q), G && (G = !1, F && F.remove(), z || A && (clearTimeout(H), P(), A = null), O.callEvent("finish", {
					selections: O.selections,
					selected: O.getSelectedArray()
				}), g.preventDefault())
			},
			B = function(l) {
				if (G) {
					return k(l)
				}
				var g = b.zui.getMouseButtonCode(J.mouseButton);
				if (!(g > -1 && l.button !== g || O.altKey || 3 === l.which || O.callEvent("start", l) === !1)) {
					var e = O.$children = O.$.find(J.selector);
					e.addClass("slectable-item");
					var h = O.multiKey ? "multi" : J.clickBehavior;
					if ("single" === h && O.unselect(), J.listenClick && ("multi" === h ? O.toggle(l.target) : "single" === h ? O.select(l.target) : "toggle" === h && O.toggle(l.target, null, function(m) {
						O.unselect()
					})), O.callEvent("startDrag", l) === !1) {
						return void O.callEvent("finish", {
							selections: O.selections,
							selected: O.getSelectedArray()
						})
					}
					M = l.pageX, I = l.pageY, F = null, z = !0, G = !0, b(document).on("mousemove" + E, o).on("mouseup" + E, k), q = setTimeout(function() {
						b(document).on(j, k)
					}, 10), l.preventDefault()
				}
			},
			R = J.container && "default" !== J.container ? b(J.container) : this.$;
		J.trigger ? R.on(j, J.trigger, B) : R.on(j, B), b(document).on("keydown", function(g) {
			var h = g.keyCode;
			17 === h || 91 == h ? O.multiKey = h : 18 === h && (O.altKey = !0)
		}).on("keyup", function(e) {
			O.multiKey = !1, O.altKey = !1
		})
	}, a.prototype.callEvent = function(j, g) {
		var l = b.Event(j + "." + this.name);
		this.$.trigger(l, g);
		var k = l.result,
			h = this.options[j];
		return b.isFunction(h) && (k = h.apply(this, b.isArray(g) ? g : [g])), k
	}, b.fn.selectable = function(e) {
		return this.each(function() {
			var j = b(this),
				h = j.data(c),
				g = "object" == typeof e && e;
			h || j.data(c, h = new a(this, g)), "string" == typeof e && h[e]()
		})
	}, b.fn.selectable.Constructor = a, b(function() {
		b('[data-ride="selectable"]').selectable()
	})
}(jQuery), +
function(d, g, c) {
	if (!d.fn.droppable) {
		return void console.error("Sortable requires droppable.js")
	}
	var j = "zui.sortable",
		h = {
			selector: "li,div",
			dragCssClass: "invisible",
			sortingClass: "sortable-sorting"
		},
		f = "order",
		b = function(k, a) {
			var l = this;
			l.$ = d(k), l.options = d.extend({}, h, l.$.data(), a), l.init()
		};
	b.DEFAULTS = h, b.NAME = j, b.prototype.init = function() {
		var w, s = this,
			p = s.$,
			m = s.options,
			z = m.selector,
			k = m.containerSelector,
			q = m.sortingClass,
			v = m.dragCssClass,
			y = m.targetSelector,
			x = m.reverse,
			A = function(a) {
				a = a || s.getItems(1);
				var l = a.length;
				l && a.each(function(o) {
					var n = x ? l - o : o;
					d(this).attr("data-" + f, n).data(f, n)
				})
			};
		A(), p.droppable({
			handle: m.trigger,
			target: y ? y : k ? z + "," + k : z,
			selector: z,
			container: p,
			always: m.always,
			flex: !0,
			lazy: m.lazy,
			canMoveHere: m.canMoveHere,
			dropToClass: m.dropToClass,
			before: m.before,
			nested: !! k,
			mouseButton: m.mouseButton,
			stopPropagation: m.stopPropagation,
			start: function(a) {
				v && a.element.addClass(v), w = !1, s.trigger("start", a)
			},
			drag: function(a) {
				if (p.addClass(q), a.isIn) {
					var u = a.element,
						e = a.target,
						B = k && e.is(k);
					if (B) {
						if (!e.children(z).filter(".dragging").length) {
							e.append(u);
							var r = s.getItems(1);
							A(r), s.trigger(f, {
								list: r,
								element: u
							})
						}
						return
					}
					var n = u.data(f),
						l = e.data(f);
					if (n === l) {
						return A(r)
					}
					n > l ? e[x ? "after" : "before"](u) : e[x ? "before" : "after"](u), w = !0;
					var r = s.getItems(1);
					A(r), s.trigger(f, {
						list: r,
						element: u
					})
				}
			},
			finish: function(a) {
				v && a.element && a.element.removeClass(v), p.removeClass(q), s.trigger("finish", {
					list: s.getItems(),
					element: a.element,
					changed: w
				})
			}
		})
	}, b.prototype.destroy = function() {
		this.$.droppable("destroy"), this.$.data(j, null)
	}, b.prototype.reset = function() {
		this.destroy(), this.init()
	}, b.prototype.getItems = function(k) {
		var a = this.$.find(this.options.selector).not(".drag-shadow");
		return k ? a : a.map(function() {
			var l = d(this);
			return {
				item: l,
				order: l.data("order")
			}
		})
	}, b.prototype.trigger = function(k, a) {
		return d.zui.callEvent(this.options[k], a, this)
	}, d.fn.sortable = function(a) {
		return this.each(function() {
			var e = d(this),
				l = e.data(j),
				k = "object" == typeof a && a;
			l ? "object" == typeof a && l.reset() : e.data(j, l = new b(this, k)), "string" == typeof a && l[a]()
		})
	}, d.fn.sortable.Constructor = b
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
(function (factory) {
	if (typeof define === 'function' && define.amd) {
		// AMD. Register as an anonymous module.
		define(['jquery'], factory);
	} else if (typeof module === 'object' && module.exports) {
		// Node/CommonJS
		module.exports = function( root, jQuery ) {
			if (typeof jQuery === 'undefined') {
				// require('jQuery') returns a factory that requires window to build a jQuery instance, we normalize how we use modules
				// that require this pattern but the window provided is a noop if it's defined (how jquery works)
				if (typeof window !== 'undefined') {
					jQuery = require('jquery');
				}
				else {
					jQuery = require('jquery')(root);
				}
			}
			factory(jQuery);
			return jQuery;
		};
	} else {
		// Browser globals
		factory(jQuery);
	}

}(function ($) {
/* eslint-enable */
	'use strict';

	/*
		Usage Note:
		-----------
		Do not use both ajaxSubmit and ajaxForm on the same form. These
		functions are mutually exclusive. Use ajaxSubmit if you want
		to bind your own submit handler to the form. For example,

		$(document).ready(function() {
			$('#myForm').on('submit', function(e) {
				e.preventDefault(); // <-- important
				$(this).ajaxSubmit({
					target: '#output'
				});
			});
		});

		Use ajaxForm when you want the plugin to manage all the event binding
		for you. For example,

		$(document).ready(function() {
			$('#myForm').ajaxForm({
				target: '#output'
			});
		});

		You can also use ajaxForm with delegation (requires jQuery v1.7+), so the
		form does not have to exist when you invoke ajaxForm:

		$('#myForm').ajaxForm({
			delegation: true,
			target: '#output'
		});

		When using ajaxForm, the ajaxSubmit function will be invoked for you
		at the appropriate time.
	*/

	var rCRLF = /\r?\n/g;

	/**
	 * Feature detection
	 */
	var feature = {};

	feature.fileapi = $('<input type="file">').get(0).files !== undefined;
	feature.formdata = (typeof window.FormData !== 'undefined');

	var hasProp = !!$.fn.prop;

	// attr2 uses prop when it can but checks the return type for
	// an expected string. This accounts for the case where a form
	// contains inputs with names like "action" or "method"; in those
	// cases "prop" returns the element
	$.fn.attr2 = function() {
		if (!hasProp) {
			return this.attr.apply(this, arguments);
		}

		var val = this.prop.apply(this, arguments);

		if ((val && val.jquery) || typeof val === 'string') {
			return val;
		}

		return this.attr.apply(this, arguments);
	};

	/**
	 * ajaxSubmit() provides a mechanism for immediately submitting
	 * an HTML form using AJAX.
	 *
	 * @param	{object|string}	options		jquery.form.js parameters or custom url for submission
	 * @param	{object}		data		extraData
	 * @param	{string}		dataType	ajax dataType
	 * @param	{function}		onSuccess	ajax success callback function
	 */
	$.fn.ajaxSubmit = function(options, data, dataType, onSuccess) {
		// fast fail if nothing selected (http://dev.jquery.com/ticket/2752)
		if (!this.length) {
			log('ajaxSubmit: skipping submit process - no element selected');

			return this;
		}

		/* eslint consistent-this: ["error", "$form"] */
		var method, action, url, isMsie, iframeSrc, $form = this;

		if (typeof options === 'function') {
			options = {success: options};

		} else if (typeof options === 'string' || (options === false && arguments.length > 0)) {
			options = {
				'url'      : options,
				'data'     : data,
				'dataType' : dataType
			};

			if (typeof onSuccess === 'function') {
				options.success = onSuccess;
			}

		} else if (typeof options === 'undefined') {
			options = {};
		}

		method = options.method || options.type || this.attr2('method');
		action = options.url || this.attr2('action');

		url = (typeof action === 'string') ? $.trim(action) : '';
		url = url || window.location.href || '';
		if (url) {
			// clean url (don't include hash vaue)
			url = (url.match(/^([^#]+)/) || [])[1];
		}
		// IE requires javascript:false in https, but this breaks chrome >83 and goes against spec.
		// Instead of using javascript:false always, let's only apply it for IE.
		isMsie = /(MSIE|Trident)/.test(navigator.userAgent || '');
		iframeSrc = (isMsie && /^https/i.test(window.location.href || '')) ? 'javascript:false' : 'about:blank'; // eslint-disable-line no-script-url

		options = $.extend(true, {
			url       : url,
			success   : $.ajaxSettings.success,
			type      : method || $.ajaxSettings.type,
			iframeSrc : iframeSrc
		}, options);

		// hook for manipulating the form data before it is extracted;
		// convenient for use with rich editors like tinyMCE or FCKEditor
		var veto = {};

		this.trigger('form-pre-serialize', [this, options, veto]);

		if (veto.veto) {
			log('ajaxSubmit: submit vetoed via form-pre-serialize trigger');

			return this;
		}

		// provide opportunity to alter form data before it is serialized
		if (options.beforeSerialize && options.beforeSerialize(this, options) === false) {
			log('ajaxSubmit: submit aborted via beforeSerialize callback');

			return this;
		}

		var traditional = options.traditional;

		if (typeof traditional === 'undefined') {
			traditional = $.ajaxSettings.traditional;
		}

		var elements = [];
		var qx, a = this.formToArray(options.semantic, elements, options.filtering);

		if (options.data) {
			var optionsData = $.isFunction(options.data) ? options.data(a) : options.data;

			options.extraData = optionsData;
			qx = $.param(optionsData, traditional);
		}

		// give pre-submit callback an opportunity to abort the submit
		if (options.beforeSubmit && options.beforeSubmit(a, this, options) === false) {
			log('ajaxSubmit: submit aborted via beforeSubmit callback');

			return this;
		}

		// fire vetoable 'validate' event
		this.trigger('form-submit-validate', [a, this, options, veto]);
		if (veto.veto) {
			log('ajaxSubmit: submit vetoed via form-submit-validate trigger');

			return this;
		}

		var q = $.param(a, traditional);

		if (qx) {
			q = (q ? (q + '&' + qx) : qx);
		}

		if (options.type.toUpperCase() === 'GET') {
			options.url += (options.url.indexOf('?') >= 0 ? '&' : '?') + q;
			options.data = null;	// data is null for 'get'
		} else {
			options.data = q;		// data is the query string for 'post'
		}

		var callbacks = [];

		if (options.resetForm) {
			callbacks.push(function() {
				$form.resetForm();
			});
		}

		if (options.clearForm) {
			callbacks.push(function() {
				$form.clearForm(options.includeHidden);
			});
		}

		// perform a load on the target only if dataType is not provided
		if (!options.dataType && options.target) {
			var oldSuccess = options.success || function(){};

			callbacks.push(function(data, textStatus, jqXHR) {
				var successArguments = arguments,
					fn = options.replaceTarget ? 'replaceWith' : 'html';

				$(options.target)[fn](data).each(function(){
					oldSuccess.apply(this, successArguments);
				});
			});

		} else if (options.success) {
			if ($.isArray(options.success)) {
				$.merge(callbacks, options.success);
			} else {
				callbacks.push(options.success);
			}
		}

		options.success = function(data, status, xhr) { // jQuery 1.4+ passes xhr as 3rd arg
			var context = options.context || this;		// jQuery 1.4+ supports scope context

			for (var i = 0, max = callbacks.length; i < max; i++) {
				callbacks[i].apply(context, [data, status, xhr || $form, $form]);
			}
		};

		if (options.error) {
			var oldError = options.error;

			options.error = function(xhr, status, error) {
				var context = options.context || this;

				oldError.apply(context, [xhr, status, error, $form]);
			};
		}

		if (options.complete) {
			var oldComplete = options.complete;

			options.complete = function(xhr, status) {
				var context = options.context || this;

				oldComplete.apply(context, [xhr, status, $form]);
			};
		}

		// are there files to upload?

		// [value] (issue #113), also see comment:
		// https://github.com/malsup/form/commit/588306aedba1de01388032d5f42a60159eea9228#commitcomment-2180219
		var fileInputs = $('input[type=file]:enabled', this).filter(function() {
			return $(this).val() !== '';
		});
		var hasFileInputs = fileInputs.length > 0;
		var mp = 'multipart/form-data';
		var multipart = ($form.attr('enctype') === mp || $form.attr('encoding') === mp);
		var fileAPI = feature.fileapi && feature.formdata;

		log('fileAPI :' + fileAPI);

		var shouldUseFrame = (hasFileInputs || multipart) && !fileAPI;
		var jqxhr;

		// options.iframe allows user to force iframe mode
		// 06-NOV-09: now defaulting to iframe mode if file input is detected
		if (options.iframe !== false && (options.iframe || shouldUseFrame)) {
			// hack to fix Safari hang (thanks to Tim Molendijk for this)
			// see: http://groups.google.com/group/jquery-dev/browse_thread/thread/36395b7ab510dd5d
			if (options.closeKeepAlive) {
				$.get(options.closeKeepAlive, function() {
					jqxhr = fileUploadIframe(a);
				});

			} else {
				jqxhr = fileUploadIframe(a);
			}

		} else if ((hasFileInputs || multipart) && fileAPI) {
			jqxhr = fileUploadXhr(a);

		} else {
			jqxhr = $.ajax(options);
		}

		$form.removeData('jqxhr').data('jqxhr', jqxhr);

		// clear element array
		for (var k = 0; k < elements.length; k++) {
			elements[k] = null;
		}

		// fire 'notify' event
		this.trigger('form-submit-notify', [this, options]);

		return this;

		// utility fn for deep serialization
		function deepSerialize(extraData) {
			var serialized = $.param(extraData, options.traditional).split('&');
			var len = serialized.length;
			var result = [];
			var i, part;

			for (i = 0; i < len; i++) {
				// #252; undo param space replacement
				serialized[i] = serialized[i].replace(/\+/g, ' ');
				part = serialized[i].split('=');
				// #278; use array instead of object storage, favoring array serializations
				result.push([decodeURIComponent(part[0]), decodeURIComponent(part[1])]);
			}

			return result;
		}

		// XMLHttpRequest Level 2 file uploads (big hat tip to francois2metz)
		function fileUploadXhr(a) {
			var formdata = new FormData();

			for (var i = 0; i < a.length; i++) {
				formdata.append(a[i].name, a[i].value);
			}

			if (options.extraData) {
				var serializedData = deepSerialize(options.extraData);

				for (i = 0; i < serializedData.length; i++) {
					if (serializedData[i]) {
						formdata.append(serializedData[i][0], serializedData[i][1]);
					}
				}
			}

			options.data = null;

			var s = $.extend(true, {}, $.ajaxSettings, options, {
				contentType : false,
				processData : false,
				cache       : false,
				type        : method || 'POST'
			});

			if (options.uploadProgress) {
				// workaround because jqXHR does not expose upload property
				s.xhr = function() {
					var xhr = $.ajaxSettings.xhr();

					if (xhr.upload) {
						xhr.upload.addEventListener('progress', function(event) {
							var percent = 0;
							var position = event.loaded || event.position;			/* event.position is deprecated */
							var total = event.total;

							if (event.lengthComputable) {
								percent = Math.ceil(position / total * 100);
							}

							options.uploadProgress(event, position, total, percent);
						}, false);
					}

					return xhr;
				};
			}

			s.data = null;

			var beforeSend = s.beforeSend;

			s.beforeSend = function(xhr, o) {
				// Send FormData() provided by user
				if (options.formData) {
					o.data = options.formData;
				} else {
					o.data = formdata;
				}

				if (beforeSend) {
					beforeSend.call(this, xhr, o);
				}
			};

			return $.ajax(s);
		}

		// private function for handling file uploads (hat tip to YAHOO!)
		function fileUploadIframe(a) {
			var form = $form[0], el, i, s, g, id, $io, io, xhr, sub, n, timedOut, timeoutHandle;
			var deferred = $.Deferred();

			// #341
			deferred.abort = function(status) {
				xhr.abort(status);
			};

			if (a) {
				// ensure that every serialized input is still enabled
				for (i = 0; i < elements.length; i++) {
					el = $(elements[i]);
					if (hasProp) {
						el.prop('disabled', false);
					} else {
						el.removeAttr('disabled');
					}
				}
			}

			s = $.extend(true, {}, $.ajaxSettings, options);
			s.context = s.context || s;
			id = 'jqFormIO' + new Date().getTime();
			var ownerDocument = form.ownerDocument;
			var $body = $form.closest('body');

			if (s.iframeTarget) {
				$io = $(s.iframeTarget, ownerDocument);
				n = $io.attr2('name');
				if (!n) {
					$io.attr2('name', id);
				} else {
					id = n;
				}

			} else {
				$io = $('<iframe name="' + id + '" src="' + s.iframeSrc + '" />', ownerDocument);
				$io.css({position: 'absolute', top: '-1000px', left: '-1000px'});
			}
			io = $io[0];


			xhr = { // mock object
				aborted               : 0,
				responseText          : null,
				responseXML           : null,
				status                : 0,
				statusText            : 'n/a',
				getAllResponseHeaders : function() {},
				getResponseHeader     : function() {},
				setRequestHeader      : function() {},
				abort                 : function(status) {
					var e = (status === 'timeout' ? 'timeout' : 'aborted');

					log('aborting upload... ' + e);
					this.aborted = 1;

					try { // #214, #257
						if (io.contentWindow.document.execCommand) {
							io.contentWindow.document.execCommand('Stop');
						}
					} catch (ignore) {}

					$io.attr('src', s.iframeSrc); // abort op in progress
					xhr.error = e;
					if (s.error) {
						s.error.call(s.context, xhr, e, status);
					}

					if (g) {
						$.event.trigger('ajaxError', [xhr, s, e]);
					}

					if (s.complete) {
						s.complete.call(s.context, xhr, e);
					}
				}
			};

			g = s.global;
			// trigger ajax global events so that activity/block indicators work like normal
			if (g && $.active++ === 0) {
				$.event.trigger('ajaxStart');
			}
			if (g) {
				$.event.trigger('ajaxSend', [xhr, s]);
			}

			if (s.beforeSend && s.beforeSend.call(s.context, xhr, s) === false) {
				if (s.global) {
					$.active--;
				}
				deferred.reject();

				return deferred;
			}

			if (xhr.aborted) {
				deferred.reject();

				return deferred;
			}

			// add submitting element to data if we know it
			sub = form.clk;
			if (sub) {
				n = sub.name;
				if (n && !sub.disabled) {
					s.extraData = s.extraData || {};
					s.extraData[n] = sub.value;
					if (sub.type === 'image') {
						s.extraData[n + '.x'] = form.clk_x;
						s.extraData[n + '.y'] = form.clk_y;
					}
				}
			}

			var CLIENT_TIMEOUT_ABORT = 1;
			var SERVER_ABORT = 2;

			function getDoc(frame) {
				/* it looks like contentWindow or contentDocument do not
				 * carry the protocol property in ie8, when running under ssl
				 * frame.document is the only valid response document, since
				 * the protocol is know but not on the other two objects. strange?
				 * "Same origin policy" http://en.wikipedia.org/wiki/Same_origin_policy
				 */

				var doc = null;

				// IE8 cascading access check
				try {
					if (frame.contentWindow) {
						doc = frame.contentWindow.document;
					}
				} catch (err) {
					// IE8 access denied under ssl & missing protocol
					log('cannot get iframe.contentWindow document: ' + err);
				}

				if (doc) { // successful getting content
					return doc;
				}

				try { // simply checking may throw in ie8 under ssl or mismatched protocol
					doc = frame.contentDocument ? frame.contentDocument : frame.document;
				} catch (err) {
					// last attempt
					log('cannot get iframe.contentDocument: ' + err);
					doc = frame.document;
				}

				return doc;
			}

			// Rails CSRF hack (thanks to Yvan Barthelemy)
			var csrf_token = $('meta[name=csrf-token]').attr('content');
			var csrf_param = $('meta[name=csrf-param]').attr('content');

			if (csrf_param && csrf_token) {
				s.extraData = s.extraData || {};
				s.extraData[csrf_param] = csrf_token;
			}

			// take a breath so that pending repaints get some cpu time before the upload starts
			function doSubmit() {
				// make sure form attrs are set
				var t = $form.attr2('target'),
					a = $form.attr2('action'),
					mp = 'multipart/form-data',
					et = $form.attr('enctype') || $form.attr('encoding') || mp;

				// update form attrs in IE friendly way
				form.setAttribute('target', id);
				if (!method || /post/i.test(method)) {
					form.setAttribute('method', 'POST');
				}
				if (a !== s.url) {
					form.setAttribute('action', s.url);
				}

				// ie borks in some cases when setting encoding
				if (!s.skipEncodingOverride && (!method || /post/i.test(method))) {
					$form.attr({
						encoding : 'multipart/form-data',
						enctype  : 'multipart/form-data'
					});
				}

				// support timout
				if (s.timeout) {
					timeoutHandle = setTimeout(function() {
						timedOut = true; cb(CLIENT_TIMEOUT_ABORT);
					}, s.timeout);
				}

				// look for server aborts
				function checkState() {
					try {
						var state = getDoc(io).readyState;

						log('state = ' + state);
						if (state && state.toLowerCase() === 'uninitialized') {
							setTimeout(checkState, 50);
						}

					} catch (e) {
						log('Server abort: ', e, ' (', e.name, ')');
						cb(SERVER_ABORT);				// eslint-disable-line callback-return
						if (timeoutHandle) {
							clearTimeout(timeoutHandle);
						}
						timeoutHandle = undefined;
					}
				}

				// add "extra" data to form if provided in options
				var extraInputs = [];

				try {
					if (s.extraData) {
						for (var n in s.extraData) {
							if (s.extraData.hasOwnProperty(n)) {
								// if using the $.param format that allows for multiple values with the same name
								if ($.isPlainObject(s.extraData[n]) && s.extraData[n].hasOwnProperty('name') && s.extraData[n].hasOwnProperty('value')) {
									extraInputs.push(
										$('<input type="hidden" name="' + s.extraData[n].name + '">', ownerDocument).val(s.extraData[n].value)
											.appendTo(form)[0]);
								} else {
									extraInputs.push(
										$('<input type="hidden" name="' + n + '">', ownerDocument).val(s.extraData[n])
											.appendTo(form)[0]);
								}
							}
						}
					}

					if (!s.iframeTarget) {
						// add iframe to doc and submit the form
						$io.appendTo($body);
					}

					if (io.attachEvent) {
						io.attachEvent('onload', cb);
					} else {
						io.addEventListener('load', cb, false);
					}

					setTimeout(checkState, 15);

					try {
						form.submit();

					} catch (err) {
						// just in case form has element with name/id of 'submit'
						var submitFn = document.createElement('form').submit;

						submitFn.apply(form);
					}

				} finally {
					// reset attrs and remove "extra" input elements
					form.setAttribute('action', a);
					form.setAttribute('enctype', et); // #380
					if (t) {
						form.setAttribute('target', t);
					} else {
						$form.removeAttr('target');
					}
					$(extraInputs).remove();
				}
			}

			if (s.forceSync) {
				doSubmit();
			} else {
				setTimeout(doSubmit, 10); // this lets dom updates render
			}

			var data, doc, domCheckCount = 50, callbackProcessed;

			function cb(e) {
				if (xhr.aborted || callbackProcessed) {
					return;
				}

				doc = getDoc(io);
				if (!doc) {
					log('cannot access response document');
					e = SERVER_ABORT;
				}
				if (e === CLIENT_TIMEOUT_ABORT && xhr) {
					xhr.abort('timeout');
					deferred.reject(xhr, 'timeout');

					return;

				}
				if (e === SERVER_ABORT && xhr) {
					xhr.abort('server abort');
					deferred.reject(xhr, 'error', 'server abort');

					return;
				}

				if (!doc || doc.location.href === s.iframeSrc) {
					// response not received yet
					if (!timedOut) {
						return;
					}
				}

				if (io.detachEvent) {
					io.detachEvent('onload', cb);
				} else {
					io.removeEventListener('load', cb, false);
				}

				var status = 'success', errMsg;

				try {
					if (timedOut) {
						throw 'timeout';
					}

					var isXml = s.dataType === 'xml' || doc.XMLDocument || $.isXMLDoc(doc);

					log('isXml=' + isXml);

					if (!isXml && window.opera && (doc.body === null || !doc.body.innerHTML)) {
						if (--domCheckCount) {
							// in some browsers (Opera) the iframe DOM is not always traversable when
							// the onload callback fires, so we loop a bit to accommodate
							log('requeing onLoad callback, DOM not available');
							setTimeout(cb, 250);

							return;
						}
						// let this fall through because server response could be an empty document
						// log('Could not access iframe DOM after mutiple tries.');
						// throw 'DOMException: not available';
					}

					// log('response detected');
					var docRoot = doc.body ? doc.body : doc.documentElement;

					xhr.responseText = docRoot ? docRoot.innerHTML : null;
					xhr.responseXML = doc.XMLDocument ? doc.XMLDocument : doc;
					if (isXml) {
						s.dataType = 'xml';
					}
					xhr.getResponseHeader = function(header){
						var headers = {'content-type': s.dataType};

						return headers[header.toLowerCase()];
					};
					// support for XHR 'status' & 'statusText' emulation :
					if (docRoot) {
						xhr.status = Number(docRoot.getAttribute('status')) || xhr.status;
						xhr.statusText = docRoot.getAttribute('statusText') || xhr.statusText;
					}

					var dt = (s.dataType || '').toLowerCase();
					var scr = /(json|script|text)/.test(dt);

					if (scr || s.textarea) {
						// see if user embedded response in textarea
						var ta = doc.getElementsByTagName('textarea')[0];

						if (ta) {
							xhr.responseText = ta.value;
							// support for XHR 'status' & 'statusText' emulation :
							xhr.status = Number(ta.getAttribute('status')) || xhr.status;
							xhr.statusText = ta.getAttribute('statusText') || xhr.statusText;

						} else if (scr) {
							// account for browsers injecting pre around json response
							var pre = doc.getElementsByTagName('pre')[0];
							var b = doc.getElementsByTagName('body')[0];

							if (pre) {
								xhr.responseText = pre.textContent ? pre.textContent : pre.innerText;
							} else if (b) {
								xhr.responseText = b.textContent ? b.textContent : b.innerText;
							}
						}

					} else if (dt === 'xml' && !xhr.responseXML && xhr.responseText) {
						xhr.responseXML = toXml(xhr.responseText);			// eslint-disable-line no-use-before-define
					}

					try {
						data = httpData(xhr, dt, s);						// eslint-disable-line no-use-before-define

					} catch (err) {
						status = 'parsererror';
						xhr.error = errMsg = (err || status);
					}

				} catch (err) {
					log('error caught: ', err);
					status = 'error';
					xhr.error = errMsg = (err || status);
				}

				if (xhr.aborted) {
					log('upload aborted');
					status = null;
				}

				if (xhr.status) { // we've set xhr.status
					status = ((xhr.status >= 200 && xhr.status < 300) || xhr.status === 304) ? 'success' : 'error';
				}

				// ordering of these callbacks/triggers is odd, but that's how $.ajax does it
				if (status === 'success') {
					if (s.success) {
						s.success.call(s.context, data, 'success', xhr);
					}

					deferred.resolve(xhr.responseText, 'success', xhr);

					if (g) {
						$.event.trigger('ajaxSuccess', [xhr, s]);
					}

				} else if (status) {
					if (typeof errMsg === 'undefined') {
						errMsg = xhr.statusText;
					}
					if (s.error) {
						s.error.call(s.context, xhr, status, errMsg);
					}
					deferred.reject(xhr, 'error', errMsg);
					if (g) {
						$.event.trigger('ajaxError', [xhr, s, errMsg]);
					}
				}

				if (g) {
					$.event.trigger('ajaxComplete', [xhr, s]);
				}

				if (g && !--$.active) {
					$.event.trigger('ajaxStop');
				}

				if (s.complete) {
					s.complete.call(s.context, xhr, status);
				}

				callbackProcessed = true;
				if (s.timeout) {
					clearTimeout(timeoutHandle);
				}

				// clean up
				setTimeout(function() {
					if (!s.iframeTarget) {
						$io.remove();
					} else { // adding else to clean up existing iframe response.
						$io.attr('src', s.iframeSrc);
					}
					xhr.responseXML = null;
				}, 100);
			}

			var toXml = $.parseXML || function(s, doc) { // use parseXML if available (jQuery 1.5+)
				if (window.ActiveXObject) {
					doc = new ActiveXObject('Microsoft.XMLDOM');
					doc.async = 'false';
					doc.loadXML(s);

				} else {
					doc = (new DOMParser()).parseFromString(s, 'text/xml');
				}

				return (doc && doc.documentElement && doc.documentElement.nodeName !== 'parsererror') ? doc : null;
			};
			var parseJSON = $.parseJSON || function(s) {
				/* jslint evil:true */
				return window['eval']('(' + s + ')');			// eslint-disable-line dot-notation
			};

			var httpData = function(xhr, type, s) { // mostly lifted from jq1.4.4

				var ct = xhr.getResponseHeader('content-type') || '',
					xml = ((type === 'xml' || !type) && ct.indexOf('xml') >= 0),
					data = xml ? xhr.responseXML : xhr.responseText;

				if (xml && data.documentElement.nodeName === 'parsererror') {
					if ($.error) {
						$.error('parsererror');
					}
				}
				if (s && s.dataFilter) {
					data = s.dataFilter(data, type);
				}
				if (typeof data === 'string') {
					if ((type === 'json' || !type) && ct.indexOf('json') >= 0) {
						data = parseJSON(data);
					} else if ((type === 'script' || !type) && ct.indexOf('javascript') >= 0) {
						$.globalEval(data);
					}
				}

				return data;
			};

			return deferred;
		}
	};

	/**
	 * ajaxForm() provides a mechanism for fully automating form submission.
	 *
	 * The advantages of using this method instead of ajaxSubmit() are:
	 *
	 * 1: This method will include coordinates for <input type="image"> elements (if the element
	 *	is used to submit the form).
	 * 2. This method will include the submit element's name/value data (for the element that was
	 *	used to submit the form).
	 * 3. This method binds the submit() method to the form for you.
	 *
	 * The options argument for ajaxForm works exactly as it does for ajaxSubmit. ajaxForm merely
	 * passes the options argument along after properly binding events for submit elements and
	 * the form itself.
	 */
	$.fn.ajaxForm = function(options, data, dataType, onSuccess) {
		if (typeof options === 'string' || (options === false && arguments.length > 0)) {
			options = {
				'url'      : options,
				'data'     : data,
				'dataType' : dataType
			};

			if (typeof onSuccess === 'function') {
				options.success = onSuccess;
			}
		}

		options = options || {};
		options.delegation = options.delegation && $.isFunction($.fn.on);

		// in jQuery 1.3+ we can fix mistakes with the ready state
		if (!options.delegation && this.length === 0) {
			var o = {s: this.selector, c: this.context};

			if (!$.isReady && o.s) {
				log('DOM not ready, queuing ajaxForm');
				$(function() {
					$(o.s, o.c).ajaxForm(options);
				});

				return this;
			}

			// is your DOM ready?  http://docs.jquery.com/Tutorials:Introducing_$(document).ready()
			log('terminating; zero elements found by selector' + ($.isReady ? '' : ' (DOM not ready)'));

			return this;
		}

		if (options.delegation) {
			$(document)
				.off('submit.form-plugin', this.selector, doAjaxSubmit)
				.off('click.form-plugin', this.selector, captureSubmittingElement)
				.on('submit.form-plugin', this.selector, options, doAjaxSubmit)
				.on('click.form-plugin', this.selector, options, captureSubmittingElement);

			return this;
		}

		if (options.beforeFormUnbind) {
			options.beforeFormUnbind(this, options);
		}

		return this.ajaxFormUnbind()
			.on('submit.form-plugin', options, doAjaxSubmit)
			.on('click.form-plugin', options, captureSubmittingElement);
	};

	// private event handlers
	function doAjaxSubmit(e) {
		/* jshint validthis:true */
		var options = e.data;

		if (!e.isDefaultPrevented()) { // if event has been canceled, don't proceed
			e.preventDefault();
			$(e.target).closest('form').ajaxSubmit(options); // #365
		}
	}

	function captureSubmittingElement(e) {
		/* jshint validthis:true */
		var target = e.target;
		var $el = $(target);

		if (!$el.is('[type=submit],[type=image]')) {
			// is this a child element of the submit el?  (ex: a span within a button)
			var t = $el.closest('[type=submit]');

			if (t.length === 0) {
				return;
			}
			target = t[0];
		}

		var form = target.form;

		form.clk = target;

		if (target.type === 'image') {
			if (typeof e.offsetX !== 'undefined') {
				form.clk_x = e.offsetX;
				form.clk_y = e.offsetY;

			} else if (typeof $.fn.offset === 'function') {
				var offset = $el.offset();

				form.clk_x = e.pageX - offset.left;
				form.clk_y = e.pageY - offset.top;

			} else {
				form.clk_x = e.pageX - target.offsetLeft;
				form.clk_y = e.pageY - target.offsetTop;
			}
		}
		// clear form vars
		setTimeout(function() {
			form.clk = form.clk_x = form.clk_y = null;
		}, 100);
	}


	// ajaxFormUnbind unbinds the event handlers that were bound by ajaxForm
	$.fn.ajaxFormUnbind = function() {
		return this.off('submit.form-plugin click.form-plugin');
	};

	/**
	 * formToArray() gathers form element data into an array of objects that can
	 * be passed to any of the following ajax functions: $.get, $.post, or load.
	 * Each object in the array has both a 'name' and 'value' property. An example of
	 * an array for a simple login form might be:
	 *
	 * [ { name: 'username', value: 'jresig' }, { name: 'password', value: 'secret' } ]
	 *
	 * It is this array that is passed to pre-submit callback functions provided to the
	 * ajaxSubmit() and ajaxForm() methods.
	 */
	$.fn.formToArray = function(semantic, elements, filtering) {
		var a = [];

		if (this.length === 0) {
			return a;
		}

		var form = this[0];
		var formId = this.attr('id');
		var els = (semantic || typeof form.elements === 'undefined') ? form.getElementsByTagName('*') : form.elements;
		var els2;

		if (els) {
			els = $.makeArray(els); // convert to standard array
		}

		// #386; account for inputs outside the form which use the 'form' attribute
		// FinesseRus: in non-IE browsers outside fields are already included in form.elements.
		if (formId && (semantic || /(Edge|Trident)\//.test(navigator.userAgent))) {
			els2 = $(':input[form="' + formId + '"]').get(); // hat tip @thet
			if (els2.length) {
				els = (els || []).concat(els2);
			}
		}

		if (!els || !els.length) {
			return a;
		}

		if ($.isFunction(filtering)) {
			els = $.map(els, filtering);
		}

		var i, j, n, v, el, max, jmax;

		for (i = 0, max = els.length; i < max; i++) {
			el = els[i];
			n = el.name;
			if (!n || el.disabled) {
				continue;
			}

			if (semantic && form.clk && el.type === 'image') {
				// handle image inputs on the fly when semantic == true
				if (form.clk === el) {
					a.push({name: n, value: $(el).val(), type: el.type});
					a.push({name: n + '.x', value: form.clk_x}, {name: n + '.y', value: form.clk_y});
				}
				continue;
			}

			v = $.fieldValue(el, true);
			if (v && v.constructor === Array) {
				if (elements) {
					elements.push(el);
				}
				for (j = 0, jmax = v.length; j < jmax; j++) {
					a.push({name: n, value: v[j]});
				}

			} else if (feature.fileapi && el.type === 'file') {
				if (elements) {
					elements.push(el);
				}

				var files = el.files;

				if (files.length) {
					for (j = 0; j < files.length; j++) {
						a.push({name: n, value: files[j], type: el.type});
					}
				} else {
					// #180
					a.push({name: n, value: '', type: el.type});
				}

			} else if (v !== null && typeof v !== 'undefined') {
				if (elements) {
					elements.push(el);
				}
				a.push({name: n, value: v, type: el.type, required: el.required});
			}
		}

		if (!semantic && form.clk) {
			// input type=='image' are not found in elements array! handle it here
			var $input = $(form.clk), input = $input[0];

			n = input.name;

			if (n && !input.disabled && input.type === 'image') {
				a.push({name: n, value: $input.val()});
				a.push({name: n + '.x', value: form.clk_x}, {name: n + '.y', value: form.clk_y});
			}
		}

		return a;
	};

	/**
	 * Serializes form data into a 'submittable' string. This method will return a string
	 * in the format: name1=value1&amp;name2=value2
	 */
	$.fn.formSerialize = function(semantic) {
		// hand off to jQuery.param for proper encoding
		return $.param(this.formToArray(semantic));
	};

	/**
	 * Serializes all field elements in the jQuery object into a query string.
	 * This method will return a string in the format: name1=value1&amp;name2=value2
	 */
	$.fn.fieldSerialize = function(successful) {
		var a = [];

		this.each(function() {
			var n = this.name;

			if (!n) {
				return;
			}

			var v = $.fieldValue(this, successful);

			if (v && v.constructor === Array) {
				for (var i = 0, max = v.length; i < max; i++) {
					a.push({name: n, value: v[i]});
				}

			} else if (v !== null && typeof v !== 'undefined') {
				a.push({name: this.name, value: v});
			}
		});

		// hand off to jQuery.param for proper encoding
		return $.param(a);
	};

	/**
	 * Returns the value(s) of the element in the matched set. For example, consider the following form:
	 *
	 *	<form><fieldset>
	 *		<input name="A" type="text">
	 *		<input name="A" type="text">
	 *		<input name="B" type="checkbox" value="B1">
	 *		<input name="B" type="checkbox" value="B2">
	 *		<input name="C" type="radio" value="C1">
	 *		<input name="C" type="radio" value="C2">
	 *	</fieldset></form>
	 *
	 *	var v = $('input[type=text]').fieldValue();
	 *	// if no values are entered into the text inputs
	 *	v === ['','']
	 *	// if values entered into the text inputs are 'foo' and 'bar'
	 *	v === ['foo','bar']
	 *
	 *	var v = $('input[type=checkbox]').fieldValue();
	 *	// if neither checkbox is checked
	 *	v === undefined
	 *	// if both checkboxes are checked
	 *	v === ['B1', 'B2']
	 *
	 *	var v = $('input[type=radio]').fieldValue();
	 *	// if neither radio is checked
	 *	v === undefined
	 *	// if first radio is checked
	 *	v === ['C1']
	 *
	 * The successful argument controls whether or not the field element must be 'successful'
	 * (per http://www.w3.org/TR/html4/interact/forms.html#successful-controls).
	 * The default value of the successful argument is true. If this value is false the value(s)
	 * for each element is returned.
	 *
	 * Note: This method *always* returns an array. If no valid value can be determined the
	 *	array will be empty, otherwise it will contain one or more values.
	 */
	$.fn.fieldValue = function(successful) {
		for (var val = [], i = 0, max = this.length; i < max; i++) {
			var el = this[i];
			var v = $.fieldValue(el, successful);

			if (v === null || typeof v === 'undefined' || (v.constructor === Array && !v.length)) {
				continue;
			}

			if (v.constructor === Array) {
				$.merge(val, v);
			} else {
				val.push(v);
			}
		}

		return val;
	};

	/**
	 * Returns the value of the field element.
	 */
	$.fieldValue = function(el, successful) {
		var n = el.name, t = el.type, tag = el.tagName.toLowerCase();

		if (typeof successful === 'undefined') {
			successful = true;
		}

		/* eslint-disable no-mixed-operators */
		if (successful && (!n || el.disabled || t === 'reset' || t === 'button' ||
			(t === 'checkbox' || t === 'radio') && !el.checked ||
			(t === 'submit' || t === 'image') && el.form && el.form.clk !== el ||
			tag === 'select' && el.selectedIndex === -1)) {
		/* eslint-enable no-mixed-operators */
			return null;
		}

		if (tag === 'select') {
			var index = el.selectedIndex;

			if (index < 0) {
				return null;
			}

			var a = [], ops = el.options;
			var one = (t === 'select-one');
			var max = (one ? index + 1 : ops.length);

			for (var i = (one ? index : 0); i < max; i++) {
				var op = ops[i];

				if (op.selected && !op.disabled) {
					var v = op.value;

					if (!v) { // extra pain for IE...
						v = (op.attributes && op.attributes.value && !(op.attributes.value.specified)) ? op.text : op.value;
					}

					if (one) {
						return v;
					}

					a.push(v);
				}
			}

			return a;
		}

		return $(el).val().replace(rCRLF, '\r\n');
	};

	/**
	 * Clears the form data. Takes the following actions on the form's input fields:
	 *  - input text fields will have their 'value' property set to the empty string
	 *  - select elements will have their 'selectedIndex' property set to -1
	 *  - checkbox and radio inputs will have their 'checked' property set to false
	 *  - inputs of type submit, button, reset, and hidden will *not* be effected
	 *  - button elements will *not* be effected
	 */
	$.fn.clearForm = function(includeHidden) {
		return this.each(function() {
			$('input,select,textarea', this).clearFields(includeHidden);
		});
	};

	/**
	 * Clears the selected form elements.
	 */
	$.fn.clearFields = $.fn.clearInputs = function(includeHidden) {
		var re = /^(?:color|date|datetime|email|month|number|password|range|search|tel|text|time|url|week)$/i; // 'hidden' is not in this list

		return this.each(function() {
			var t = this.type, tag = this.tagName.toLowerCase();

			if (re.test(t) || tag === 'textarea') {
				this.value = '';

			} else if (t === 'checkbox' || t === 'radio') {
				this.checked = false;

			} else if (tag === 'select') {
				this.selectedIndex = -1;

			} else if (t === 'file') {
				if (/MSIE/.test(navigator.userAgent)) {
					$(this).replaceWith($(this).clone(true));
				} else {
					$(this).val('');
				}

			} else if (includeHidden) {
				// includeHidden can be the value true, or it can be a selector string
				// indicating a special test; for example:
				// $('#myForm').clearForm('.special:hidden')
				// the above would clean hidden inputs that have the class of 'special'
				if ((includeHidden === true && /hidden/.test(t)) ||
					(typeof includeHidden === 'string' && $(this).is(includeHidden))) {
					this.value = '';
				}
			}
		});
	};


	/**
	 * Resets the form data or individual elements. Takes the following actions
	 * on the selected tags:
	 * - all fields within form elements will be reset to their original value
	 * - input / textarea / select fields will be reset to their original value
	 * - option / optgroup fields (for multi-selects) will defaulted individually
	 * - non-multiple options will find the right select to default
	 * - label elements will be searched against its 'for' attribute
	 * - all others will be searched for appropriate children to default
	 */
	$.fn.resetForm = function() {
		return this.each(function() {
			var el = $(this);
			var tag = this.tagName.toLowerCase();

			switch (tag) {
			case 'input':
				this.checked = this.defaultChecked;
				// fall through

			case 'textarea':
				this.value = this.defaultValue;

				return true;

			case 'option':
			case 'optgroup':
				var select = el.parents('select');

				if (select.length && select[0].multiple) {
					if (tag === 'option') {
						this.selected = this.defaultSelected;
					} else {
						el.find('option').resetForm();
					}
				} else {
					select.resetForm();
				}

				return true;

			case 'select':
				el.find('option').each(function(i) {				// eslint-disable-line consistent-return
					this.selected = this.defaultSelected;
					if (this.defaultSelected && !el[0].multiple) {
						el[0].selectedIndex = i;

						return false;
					}
				});

				return true;

			case 'label':
				var forEl = $(el.attr('for'));
				var list = el.find('input,select,textarea');

				if (forEl[0]) {
					list.unshift(forEl[0]);
				}

				list.resetForm();

				return true;

			case 'form':
				// guard against an input with the name of 'reset'
				// note that IE reports the reset function as an 'object'
				if (typeof this.reset === 'function' || (typeof this.reset === 'object' && !this.reset.nodeType)) {
					this.reset();
				}

				return true;

			default:
				el.find('form,input,label,select,textarea').resetForm();

				return true;
			}
		});
	};

	/**
	 * Enables or disables any matching elements.
	 */
	$.fn.enable = function(b) {
		if (typeof b === 'undefined') {
			b = true;
		}

		return this.each(function() {
			this.disabled = !b;
		});
	};

	/**
	 * Checks/unchecks any matching checkboxes or radio buttons and
	 * selects/deselects and matching option elements.
	 */
	$.fn.selected = function(select) {
		if (typeof select === 'undefined') {
			select = true;
		}

		return this.each(function() {
			var t = this.type;

			if (t === 'checkbox' || t === 'radio') {
				this.checked = select;

			} else if (this.tagName.toLowerCase() === 'option') {
				var $sel = $(this).parent('select');

				if (select && $sel[0] && $sel[0].type === 'select-one') {
					// deselect all other options
					$sel.find('option').selected(false);
				}

				this.selected = select;
			}
		});
	};

	// expose debug var
	$.fn.ajaxSubmit.debug = false;

	// helper fn for console logging
	function log() {
		if (!$.fn.ajaxSubmit.debug) {
			return;
		}

		var msg = '[jquery.form] ' + Array.prototype.join.call(arguments, '');

		if (window.console && window.console.log) {
			window.console.log(msg);

		} else if (window.opera && window.opera.postError) {
			window.opera.postError(msg);
		}
	}
})),

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

function(a) {
	function b(d) {
		if ("string" == typeof d.data) {
			var c = d.handler,
				f = d.data.toLowerCase().split(" ");
			d.handler = function(p) {
				if (this === p.target || !/textarea|select/i.test(p.target.nodeName) && "text" !== p.target.type) {
					var q = "keypress" !== p.type && a.hotkeys.specialKeys[p.which],
						m = String.fromCharCode(p.which).toLowerCase(),
						j = "",
						n = {};
					p.altKey && "alt" !== q && (j += "alt+"), p.ctrlKey && "ctrl" !== q && (j += "ctrl+"), p.metaKey && !p.ctrlKey && "meta" !== q && (j += "meta+"), p.shiftKey && "shift" !== q && (j += "shift+"), q ? n[j + q] = !0 : (n[j + m] = !0, n[j + a.hotkeys.shiftNums[m]] = !0, "shift+" === j && (n[a.hotkeys.shiftNums[m]] = !0));
					for (var g = 0, k = f.length; g < k; g++) {
						if (n[f[g]]) {
							return c.apply(this, arguments)
						}
					}
				}
			}
		}
	}
	a.hotkeys = {
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
	}, a.each(["keydown", "keyup", "keypress"], function() {
		a.event.special[this] = {
			add: b
		}
	})
}(jQuery), function(d) {
	function g(k, a) {
		if (k === !1) {
			return k
		}
		if (!k) {
			return a
		}
		k === !0 ? k = {
			add: !0,
			"delete": !0,
			edit: !0,
			sort: !0
		} : "string" == typeof k && (k = k.split(","));
		var l;
		return d.isArray(k) && (l = {}, d.each(k, function(n, m) {
			d.isPlainObject(m) ? l[m.action] = m : l[m] = !0
		}), k = l), d.isPlainObject(k) && (l = {}, d.each(k, function(n, m) {
			m ? l[n] = d.extend({
				type: n
			}, b[n], d.isPlainObject(m) ? m : null) : l[n] = !1
		}), k = l), a ? d.extend(!0, {}, a, k) : k
	}
	function c(k, a, l) {
		return a = a || k.type, d(l || k.template).addClass("tree-action").attr(d.extend({
			"data-type": a,
			title: k.title || ""
		}, k.attr)).data("action", k)
	}
	var j = "zui.tree",
		h = 0,
		f = function(k, a) {
			this.name = j, this.$ = d(k), this.getOptions(a), this._init()
		},
		b = {
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
	f.DEFAULTS = {
		animate: null,
		initialState: "normal",
		toggleTemplate: '<i class="list-toggle icon"></i>'
	}, f.prototype.add = function(w, u, p, m, y) {
		var x, k = d(w),
			q = this.options;
		if (k.is("li") ? (x = k.children("ul"), x.length || (x = d("<ul/>"), k.append(x), this._initList(x, k))) : x = k, x) {
			var v = this;
			d.isArray(u) || (u = [u]), d.each(u, function(r, a) {
				var A = d("<li/>").data(a).appendTo(x);
				void 0 !== a.id && A.attr("data-id", a.id);
				var z = q.itemWrapper ? d(q.itemWrapper === !0 ? '<div class="tree-item-wrapper"/>' : q.itemWrapper).appendTo(A) : A;
				if (a.html) {
					z.html(a.html)
				} else {
					if (d.isFunction(v.options.itemCreator)) {
						var l = v.options.itemCreator(A, a);
						l !== !0 && l !== !1 && z.html(l)
					} else {
						a.url ? z.append(d("<a/>", {
							href: a.url
						}).text(a.title || a.name)) : z.append(d("<span/>").text(a.title || a.name))
					}
				}
				v._initItem(A, a.idx || r, x, a), a.children && a.children.length && v.add(A, a.children)
			}), this._initList(x), p && !x.hasClass("tree") && v.expand(x.parent("li"), m, y)
		}
	}, f.prototype.reload = function(k) {
		var a = this;
		k && (a.$.empty(), a.add(a.$, k)), a.isPreserve && a.store.time && a.$.find("li:not(.tree-action-item)").each(function() {
			var l = d(this);
			a[a.store[l.data("id")] ? "expand" : "collapse"](l, !0, !0)
		})
	}, f.prototype._initList = function(w, u, p, k) {
		var q = this;
		w.hasClass("tree") ? (p = 0, u = null) : (u = (u || w.closest("li")).addClass("has-list"), u.find(".list-toggle").length || u.prepend(this.options.toggleTemplate), p = p || u.data("idx")), w.removeClass("has-active-item");
		var e = w.attr("data-idx", p || 0).children("li:not(.tree-action-item)").each(function(a) {
			q._initItem(d(this), a + 1, w)
		});
		1 !== e.length || e.find("ul").length || e.addClass("tree-single-item"), k = k || (u ? u.data() : null);
		var m = g(k ? k.actions : null, this.actions);
		if (m) {
			if (m.add && m.add.templateInList !== !1) {
				var v = w.children("li.tree-action-item");
				v.length ? v.detach().appendTo(w) : d('<li class="tree-action-item"/>').append(c(m.add, "add", m.add.templateInList)).appendTo(w)
			}
			m.sort && w.sortable(d.extend({
				dragCssClass: "tree-drag-holder",
				trigger: ".sort-handler",
				selector: "li:not(.tree-action-item)",
				finish: function(a) {
					q.callEvent("action", {
						action: m.sort,
						$list: w,
						target: a.target,
						item: k
					})
				}
			}, m.sort.options, d.isPlainObject(this.options.sortable) ? this.options.sortable : null))
		}
		u && (u.hasClass("open") || k && k.open) && u.addClass("open in")
	}, f.prototype._initItem = function(m, k, x, w) {
		if (void 0 === k) {
			var e = m.prev("li");
			k = e.length ? e.data("idx") + 1 : 1
		}
		if (x = x || m.closest("ul"), m.attr("data-idx", k).removeClass("tree-single-item"), !m.data("id")) {
			var p = k;
			x.hasClass("tree") || (p = x.parent("li").data("id") + "-" + p), m.attr("data-id", p)
		}
		m.hasClass("active") && x.parent("li").addClass("has-active-item"), w = w || m.data();
		var q = g(w.actions, this.actions);
		if (q) {
			var v = m.find(".tree-actions");
			v.length || (v = d('<div class="tree-actions"/>').appendTo(this.options.itemWrapper ? m.find(".tree-item-wrapper") : m), d.each(q, function(a, l) {
				l && v.append(c(l, a))
			}))
		}
		var u = m.children("ul");
		u.length && this._initList(u, m, k, w)
	}, f.prototype._init = function() {
		var k = this.options,
			l = this;
		this.actions = g(k.actions), this.$.addClass("tree"), k.animate && this.$.addClass("tree-animate"), this._initList(this.$);
		var e = k.initialState,
			m = d.zui && d.zui.store && d.zui.store.enable;
		m && (this.selector = j + "::" + (k.name || "") + "#" + (this.$.attr("id") || h++), this.store = d.zui.store[k.name ? "get" : "pageGet"](this.selector, {})), "preserve" === e && (m ? this.isPreserve = !0 : this.options.initialState = e = "normal"), this.reload(k.data), m && (this.isPreserve = !0), "expand" === e ? this.expand() : "collapse" === e && this.collapse(), this.$.on("click", '.list-toggle,a[href="#"],.tree-toggle', function(o) {
			var a = d(this),
				p = a.parent("li");
			l.callEvent("hit", {
				target: p,
				item: p.data()
			}), l.toggle(p), a.is("a") && o.preventDefault()
		}).on("click", ".tree-action", function() {
			var o = d(this),
				a = o.data();
			if (a.action && (a = a.action), "sort" !== a.type) {
				var p = o.closest("li:not(.tree-action-item)");
				l.callEvent("action", {
					action: a,
					target: this,
					$item: p,
					item: p.data()
				})
			}
		})
	}, f.prototype.preserve = function(k, a, m) {
		if (this.isPreserve) {
			if (k) {
				a = a || k.data("id"), m = void 0 === m && k.hasClass("open"), m ? this.store[a] = m : delete this.store[a], this.store.time = (new Date).getTime(), d.zui.store[this.options.name ? "set" : "pageSet"](this.selector, this.store)
			} else {
				var l = this;
				this.store = {}, this.$.find("li").each(function() {
					l.preserve(d(this))
				})
			}
		}
	}, f.prototype.expand = function(k, l, a) {
		k ? (k.addClass("open"), !l && this.options.animate ? setTimeout(function() {
			k.addClass("in")
		}, 10) : k.addClass("in")) : k = this.$.find("li.has-list").addClass("open in"), a || this.preserve(k), this.callEvent("expand", k, this)
	}, f.prototype.show = function(k, a, m) {
		var l = this;
		k.each(function() {
			var p = d(this);
			if (l.expand(p, a, m), p) {
				for (var o = p.parent("ul"); o && o.length && !o.hasClass("tree");) {
					var n = o.parent("li");
					n.length ? (l.expand(n, a, m), o = n.parent("ul")) : o = !1
				}
			}
		})
	}, f.prototype.collapse = function(k, l, a) {
		k ? !l && this.options.animate ? (k.removeClass("in"), setTimeout(function() {
			k.removeClass("open")
		}, 300)) : k.removeClass("open in") : k = this.$.find("li.has-list").removeClass("open in"), a || this.preserve(k), this.callEvent("collapse", k, this)
	}, f.prototype.toggle = function(a) {
		var k = a && a.hasClass("open") || a === !1 || void 0 === a && this.$.find("li.has-list.open").length;
		this[k ? "collapse" : "expand"](a)
	}, f.prototype.getOptions = function(a) {
		this.options = d.extend({}, f.DEFAULTS, this.$.data(), a), null === this.options.animate && this.$.hasClass("tree-animate") && (this.options.animate = !0)
	}, f.prototype.toData = function(k, a) {
		d.isFunction(k) && (a = k, k = null), k = k || this.$;
		var l = this;
		return k.children("li:not(.tree-action-item)").map(function() {
			var n = d(this),
				p = n.data();
			delete p["zui.droppable"];
			var m = n.children("ul");
			return m.length && (p.children = l.toData(m)), d.isFunction(a) ? a(p, n) : p
		}).get()
	}, f.prototype.callEvent = function(k, a) {
		var l;
		return d.isFunction(this.options[k]) && (l = this.options[k](a, this)), this.$.trigger(d.Event(k + "." + this.name, a)), l
	}, d.fn.tree = function(k, a) {
		return this.each(function() {
			var m = d(this),
				e = m.data(j),
				l = "object" == typeof k && k;
			e || m.data(j, e = new f(this, l)), "string" == typeof k && e[k](a)
		})
	}, d.fn.tree.Constructor = f, d(function() {
		d('[data-ride="tree"]').tree()
	})
}(jQuery), function(b) {
	var c = "zui.colorPicker",
		a = '<div class="colorpicker"><button type="button" class="btn dropdown-toggle" data-toggle="dropdown"><span class="cp-title"></span><i class="ic"></i></button><ul class="dropdown-menu clearfix"></ul></div>',
		f = {
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
		d = function(e, g) {
			this.name = c, this.$ = b(e), this.getOptions(g), this.init()
		};
	d.DEFAULTS = {
		colors: ["#00BCD4", "#388E3C", "#3280fc", "#3F51B5", "#9C27B0", "#795548", "#F57C00", "#F44336", "#E91E63"],
		pullMenuRight: !0,
		wrapper: "btn-wrapper",
		tileSize: 30,
		lineCount: 5,
		optional: !0,
		tooltip: "top",
		icon: "caret-down",
		updateBtn: "auto"
	}, d.prototype.init = function() {
		var k = this,
			m = k.options,
			l = k.$,
			h = l.parent(),
			g = !1;
		h.hasClass("colorpicker") ? k.$picker = h : (k.$picker = b(m.template || a), g = !0), k.$picker.addClass(m.wrapper).find(".cp-title").toggle(void 0 !== m.title).text(m.title), k.$menu = k.$picker.find(".dropdown-menu").toggleClass("pull-right", m.pullMenuRight), k.$btn = k.$picker.find(".btn.dropdown-toggle"), k.$btn.find(".ic").addClass("icon-" + m.icon), m.btnTip && k.$picker.attr("data-toggle", "tooltip").tooltip({
			title: m.btnTip,
			placement: m.tooltip,
			container: "body"
		}), l.attr("data-provide", null), g && l.after(k.$picker), k.colors = {}, b.each(m.colors, function(e, q) {
			if (b.zui.Color.isColor(q)) {
				var p = new b.zui.Color(q);
				k.colors[p.toCssStr()] = p
			}
		}), k.updateColors(), k.$picker.on("click", ".cp-tile", function() {
			k.setValue(b(this).data("color"))
		});
		var j = function() {
				var e = l.val(),
					n = b.zui.Color.isColor(e);
				l.parent().toggleClass("has-error", !(n || m.optional && "" === e)), n ? k.setValue(e, !0) : m.optional && "" === e ? l.tooltip("hide") : l.is(":focus") || l.tooltip("show", m.errorTip)
			};
		l.is("input:not([type=hidden])") ? (m.tooltip && l.attr("data-toggle", "tooltip").tooltip({
			trigger: "manual",
			placement: m.tooltip,
			tipClass: "tooltip-danger",
			container: "body"
		}), l.on("keyup paste input change", j)) : l.appendTo(k.$picker), j()
	}, d.prototype.addColor = function(h) {
		h instanceof b.zui.Color || (h = new b.zui.Color(h));
		var g = h.toCssStr(),
			k = this.options;
		this.colors[g] || (this.colors[g] = h);
		var j = b('<a href="###" class="cp-tile"></a>', {
			titile: h
		}).data("color", h).css({
			color: h.contrast().toCssStr(),
			background: g,
			"border-color": h.luma() > 0.43 ? "#ccc" : "transparent"
		}).attr("data-color", g);
		this.$menu.append(b("<li/>").css({
			width: k.tileSize,
			height: k.tileSize
		}).append(j)), k.optional && this.$menu.find(".cp-tile.empty").parent().detach().appendTo(this.$menu)
	}, d.prototype.updateColors = function(k) {
		var h = (this.$picker, this.$menu.children("li:not(.heading)").remove()),
			m = this.options,
			k = k || this.colors,
			l = this,
			j = 0;
		if (b.each(k, function(n, o) {
			l.addColor(o), j++
		}), m.optional) {
			var g = b('<li><a class="cp-tile empty" href="###"></a></li>').css({
				width: m.tileSize,
				height: m.tileSize
			});
			this.$menu.append(g), j++
		}
		h.css("width", Math.min(j, m.lineCount) * m.tileSize + 6)
	}, d.prototype.setValue = function(u, p) {
		var k = this,
			j = k.options,
			w = k.$btn;
		k.$menu.find(".cp-tile.active").removeClass("active");
		var v = "",
			g = j.updateBtn;
		if ("auto" === g) {
			var m = w.find(".color-bar");
			g = !m.length ||
			function(e) {
				m.css("background", e || "")
			}
		}
		if (u) {
			var q = new b.zui.Color(u);
			v = q.toCssStr().toLowerCase(), g && (b.isFunction(g) ? g(v, w, k) : w.css({
				background: v,
				color: q.contrast().toCssStr(),
				borderColor: q.luma() > 0.43 ? "#ccc" : v
			})), k.colors[v] || k.addColor(q), p || k.$.val().toLowerCase() === v || k.$.val(v).trigger("change"), k.$menu.find('.cp-tile[data-color="' + v + '"]').addClass("active"), k.$.tooltip("hide"), k.$.trigger("colorchange", q)
		} else {
			g && (b.isFunction(g) ? g(null, w, k) : w.attr("style", null)), p || "" === k.$.val() || k.$.val(v).trigger("change"), j.optional && k.$.tooltip("hide"), k.$menu.find(".cp-tile.empty").addClass("active"), k.$.trigger("colorchange", null)
		}
		j.updateBorder && b(j.updateBorder).css("border-color", v), j.updateBackground && b(j.updateBackground).css("background-color", v), j.updateColor && b(j.updateColor).css("color", v), j.updateText && b(j.updateText).text(v)
	}, d.prototype.getOptions = function(j) {
		var g = b.extend({}, d.DEFAULTS, this.$.data(), j);
		"string" == typeof g.colors && (g.colors = g.colors.split(","));
		var h = (g.lang || b.zui.clientLang()).toLowerCase();
		g.errorTip || (g.errorTip = f[h].errorTip), b.fn.tooltip || (g.btnTip = !1), this.options = g
	}, b.fn.colorPicker = function(e) {
		return this.each(function() {
			var j = b(this),
				h = j.data(c),
				g = "object" == typeof e && e;
			h || j.data(c, h = new d(this, g)), "string" == typeof e && h[e]()
		})
	}, b.fn.colorPicker.Constructor = d, b(function() {
		b('[data-provide="colorpicker"]').colorPicker()
	})
}(jQuery), function(b, d) {
	function a(e) {
		return e === d && (e = g += 1), f[e % f.length]
	}
	var g = 0,
		f = ["#00a9fc", "#ff5d5d", "#fdc137", "#00da88", "#7ec5ff", "#8666b8", "#bd7b46", "#ff9100", "#ff3d00", "#f57f17", "#00e5ff", "#00b0ff", "#2979ff", "#3d5afe", "#651fff", "#d500f9", "#f50057", "#ff1744"];
	jQuery.fn.tableChart = function() {
		b(this).each(function() {
			var u = b(this),
				m = u.data(),
				k = m.chart || "pie",
				y = b(m.target);
			if (y.length) {
				var x = null;
				if ("pie" === k) {
					m = b.extend({
						scaleShowLabels: !0,
						scaleLabel: "<%=label%>: <%=value%>"
					}, m);
					var j = [],
						p = u.find("tbody > tr").each(function(h) {
							var r = b(this),
								l = a();
							r.attr("data-id", h).find(".chart-color-dot").css("background", l), j.push({
								label: r.find(".chart-label").text(),
								value: parseInt(r.find(".chart-value").text()),
								color: l,
								id: h
							})
						});
					j.length > 1 ? m.scaleLabelPlacement = "outside" : 1 === j.length && (m.scaleLabelPlacement = "inside", j.push({
						label: "",
						value: j[0].value / 2000,
						color: "#fff",
						showLabel: !1
					})), x = y.pieChart(j, m), y.on("mousemove", function(h) {
						var l = x.getSegmentsAtEvent(h);
						p.removeClass("active"), l.length && p.filter('[data-id="' + l[0].id + '"]').addClass("active")
					})
				} else {
					if ("bar" === k) {
						var q = a(),
							w = [],
							v = {
								label: u.find("thead .chart-label").text(),
								color: q,
								data: []
							},
							p = u.find("tbody > tr").each(function(l) {
								var h = b(this);
								w.push(h.find(".chart-label").text()), v.data.push(parseInt(h.find(".chart-value").text())), h.find(".chart-color-dot").css("background", q)
							}),
							j = {
								labels: w,
								datasets: [v]
							};
						w.length && (m.barValueSpacing = 5), x = y.barChart(j, m)
					} else {
						if ("line" === k) {
							var q = a(),
								w = [],
								v = {
									label: u.find("thead .chart-label").text(),
									color: q,
									data: []
								},
								p = u.find("tbody > tr").each(function(l) {
									var h = b(this);
									w.push(h.find(".chart-label").text()), v.data.push(parseInt(h.find(".chart-value").text())), h.find(".chart-color-dot").css("background", q)
								}),
								j = {
									labels: w,
									datasets: [v]
								};
							w.length && (m.barValueSpacing = 5), x = y.lineChart(j, m)
						}
					}
				}
				null !== x && u.data("zui.chart", x)
			}
		})
	}, b(".table-chart").tableChart();
	var c = function(j, q) {
			var p = b(j);

			if (!p.data("initProgressPie")) {
				p.data("initProgressPie", 1);
				var k = p.is("canvas") ? p : p.find("canvas"),
					h = b.extend({
						value: 0,
						color: p.attr('data-color') || b.getThemeColor("primary") || "#006af1",
						backColor: b.getThemeColor("pale") || "#E9F2FB",
						doughnut: !0,
						doughnutSize: 85,
						width: 20,
						height: 20,
						showTip: !1,
						name: "",
						tipTemplate: "<%=value%>%",
						animation: "auto",
						realValue: parseFloat(p.find(".progress-value").text())
					}, q, p.data()),
					m = k.length;
				m || (k = b("<canvas>").appendTo(p)), k.attr("width") !== d ? h.width = k.attr("width") : k.attr("width", h.width), k.attr("height") !== d ? h.height = k.attr("height") : k.attr("height", h.height), m || 8 != b.zui.browser.ie || G_vmlCanvasManager.initElement(k[0]), "auto" === h.animation && (h.animation = h.width > 30), h.value = Math.max(0, Math.min(100, h.value)), p.addClass("progress-pie-" + h.width);
				var e = [{
					value: h.value,
					label: h.name,
					color: h.color,
					circleBeginEnd: !0
				}, {
					value: 100 - h.value,
					label: "",
					color: h.backColor
				}];
				k[h.doughnut ? "doughnutChart" : "pieChart"](e, b.extend({
					segmentShowStroke: !1,
					animation: h.animation,
					showTooltips: h.showTip,
					tooltipTemplate: h.tipTemplate,
					percentageInnerCutout: h.doughnutSize,
					reverseDrawOrder: !0,
					animationEasing: "easeInOutQuart",
					onAnimationProgress: h.realValue ?
					function(l) {
						p.find(".progress-value").text(Math.floor(h.realValue * l))
					} : d,
					onAnimationComplete: h.realValue ?
					function(l) {
						p.find(".progress-value").text(h.realValue)
					} : d
				}, h.chartOptions))
			}
		};
	jQuery.fn.progressPie = function(h) {
		b(this).each(function() {
			var e = b(this);
			if (!e.closest(".hidden").length) {
				var j = e.closest(".tab-pane");
				j.length && !j.hasClass("active") ? b('[data-toggle="tab"][data-target="#' + j.attr("id") + '"]').one("shown.zui.tab", function() {
					c(e, h)
				}) : c(this, h)
			}
		})
	}, b(function() {
		b(".table-chart").tableChart();
		var h = b(".progress-pie");
		h.length > 100 ? setTimeout(function() {
			h.progressPie()
		}, 1000) : h.progressPie()
	})
}(jQuery, void 0), function(a) {
	jQuery.fn.sparkline = function(b) {
		a(this).each(function() {
			var v = a(this),
				m = a.extend({
					values: v.attr("values"),
					width: v.width() - 4,
					height: v.height() - 4
				}, v.data(), b),
				k = m.height,
				D = [],
				B = m.width,
				e = m.values.split(","),
				q = 0;
			for (var w in e) {
				var A = parseFloat(e[w]);
				NaN != A && (D.push(A), q = Math.max(A, q))
			}
			var z = (Math.min(q, 30), Math.min(B, Math.max(10, D.length * B / 30))),
				C = v.children("canvas");
			C.length || (v.append('<canvas class="projectline-canvas"></canvas>'), C = v.children("canvas")), C.attr("width", z).attr("height", k);
			var j = {
				labels: D,
				datasets: [{
					fillColor: a.getThemeColor("pale") || "rgba(0,0,255,0.05)",
					strokeColor: a.getThemeColor("primary") || "#0054EC",
					pointColor: a.getThemeColor("secondary") || "rgba(255,136,0,1)",
					pointStrokeColor: "#fff",
					data: D
				}]
			},
				y = {
					animation: !0,
					scaleOverride: !0,
					scaleStepWidth: Math.ceil(q / 10),
					scaleSteps: 10,
					scaleStartValue: 0,
					showScale: !1,
					showTooltips: !1,
					pointDot: !1,
					scaleShowGridLines: !1,
					datasetStrokeWidth: 1
				},
				x = a(C).lineChart(j, y);
			v.data("sparklineChart", x)
		})
	}, a(function() {
		a(".sparkline").sparkline()
	})
}(jQuery), function(a) {
	a(function() {
		a.fn.fixedDate = function() {
			return a(this).each(function() {
				var f = a(this).attr("autocomplete", "off");
				"0000-00-00" == f.val() && f.focus(function() {
					"0000-00-00" == f.val() && f.val("").datetimepicker("update")
				}).blur(function() {
					"" == f.val() && f.val("0000-00-00")
				})
			})
		};
		var c = {
			language: a("html").attr("lang"),
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
			b = a.extend({}, c, {
				minView: 2,
				format: "yyyy-mm-dd"
			}),
			d = a.extend({}, c, {
				startView: 1,
				minView: 0,
				maxView: 1,
				format: "hh:ii"
			});
		a(".datepicker-wrapper").click(function() {
			a(this).find(".form-date, .form-datetime, .form-time").datetimepicker("show").focus()
		}), window.datepickerOptions = c, a.fn.datepicker = function(f) {
			return this.datetimepicker(a.extend({}, b, f))
		}, a.fn.timepicker = function(f) {
			return this.datetimepicker(a.extend({}, d, f))
		}, a.fn.datepickerAll = function() {
			return this.find(".form-datetime").fixedDate().datetimepicker(c), this.find(".form-date").fixedDate().datepicker(), this.find(".form-time").fixedDate().timepicker(), this
		}, a("body").datepickerAll()
	})
}(jQuery), function(a) {
	var b = function(w, q) {
			q = a.extend({
				idStart: 0,
				idEnd: 9,
				chosen: !0,
				datetimepicker: !0,
				colorPicker: !0,
				hotkeys: !0
			}, q, w.data());
			var k = w.find(".template");
			!k.length && q.template && (k = a(q.template));
			var j = 0,
				B = 0,
				z = function(c) {
					c.is("select.chosen") ? c.next(".chosen-container").find("input").focus() : c.focus()
				},
				f = function(d) {
					var c = w.find("[data-ctrl-index]:focus,.chosen-container-active").first();
					if (c.length) {
						if (c.is(".chosen-container-active")) {
							if (c.hasClass("chosen-with-drop") && ("down" === d || "up" === d)) {
								return
							}
							c = c.prev("select.chosen")
						}
						var h = c.data("ctrlIndex"),
							e = c.closest("tr").data("row");
						"down" === d ? e < B - 1 ? e += 1 : e = 0 : "up" === d ? e > 0 ? e -= 1 : e = B - 1 : "left" === d ? h > 0 ? h -= 1 : h = j - 1 : "right" === d && (h < j - 1 ? h += 1 : h = 0), z(w.find('tr[data-row="' + e + '"]').find('[data-ctrl-index="' + h + '"]'))
					}
				},
				m = {
					options: q,
					focusNext: f,
					focusControl: z
				},
				v = w.find("tbody,.batch-rows"),
				y = function(c) {
					a.fn.chosen && q.chosen && c.find(".chosen").chosen(a.isPlainObject(q.chosen) ? q.chosen : null), a.fn.datetimepicker && q.datetimepicker && c.datepickerAll(a.isPlainObject(q.datetimepicker) ? q.datetimepicker : null), a.fn.colorPicker && q.colorPicker && c.find("input.colorpicker").colorPicker(a.isPlainObject(q.colorPicker) ? q.colorPicker : null);
					var d = 0;
					c.find('input[type!="hidden"],textarea,select').each(function() {
						var h = a(this);
						h.parent().hasClass("chosen-search") || h.attr("data-ctrl-index", d++)
					}), j = Math.max(j, d)
				};
			if (k.length) {
				var x = k.remove().html(),
					A = function(h, l) {
						var c = x;
						"number" != typeof h && (h = B), B = Math.max(h + 1, B), c = c.replace(/\$idPlus/g, h + 1).replace(/\$id/g, h);
						var d = a("<" + k[0].tagName.toLowerCase() + " />").html(c);
						return d.attr("data-row", h).addClass(k.attr("class")).removeClass("template"), q.rowCreator && q.rowCreator(d, h, q), l ? l.after(d) : v.append(d), y(d), d
					};
				a.extend(m, {
					createRow: A,
					template: x
				});
				for (var g = q.idStart; g <= q.idEnd; ++g) {
					A(g)
				}
			} else {
				y(w)
			}
			w.on("click", ".btn-copy", function() {
				var d = a(this),
					c = a(d.data("copyFrom")).val(),
					h = a(d.data("copyTo")).val(c).addClass("highlight");
				setTimeout(function() {
					h.removeClass("highlight")
				}, 2000)
			}), q.hotkeys && a(document).on("keydown", function(d) {
				var h = {
					"Ctrl+#37": "left",
					"Ctrl+#39": "right",
					"#38": "up",
					"#40": "down",
					"Ctrl+#38": "up",
					"Ctrl+#40": "down"
				},
					c = [];
				d.ctrlKey && c.push("Ctrl"), c.push("#" + d.keyCode);
				var l = h[c.join("+")];
				l && (f(l), d.ctrlKey && (d.stopPropagation(), d.preventDefault()))
			}), w.data("zui.batchActionForm", m)
		};
	a.fn.batchActionForm = function(c) {
		return this.each(function() {
			b(a(this), c)
		})
	}
}(jQuery), function(b, d) {
	var a = "zui.table",
		g = {
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
		f = /^((?!chrome|android).)*safari/i.test(navigator.userAgent),
		c = function(p, q) {
			var k = this;
			k.name = a;
			var n = k.$ = b(p);
			q = k.options = b.extend({}, c.DEFAULTS, this.$.data(), q);
			var j = q.lang || "zh_cn";
			k.lang = b.isPlainObject(j) ? b.extend(!0, {}, g[j.lang || b.zui.clientLang()], j) : g[j], n.attr("id") || (n.attr("id", "table-" + b.zui.uuid()), q.hot && console.warn("ZUI: table hot replace id not defined, the element id attribute should be set.")), n.attr("data-ride") || n.attr("data-ride", "table"), k.getTable().find("thead>tr>th").each(function() {
				var l = b(this);
				if (!l.attr("title")) {
					var h = b.trim(l.find("a").text() || l.text() || "");
					h.length && l.attr("title", h)
				}
			}), q.checkable && (n.on("click", ".check-all", function() {
				k.checkAll(!b(this).hasClass("checked"))
			}).on("click", "tbody>tr", function(h) {
				b(h.target).closest('.btn,a,.not-check,.form-control,input[type="text"],.chosen-container').length || k.checkRow(b(this))
			}).on("click", 'tbody input[type="checkbox"],tbody label[for]', function(l) {
				l.stopPropagation();
				var h = b(this);
				h.is("label") && (h = h.closest(".checkbox-primary").find('input[type="checkbox"]')), k.checkRow(h.closest("tr"), h.is(":checked"))
			}), q.selectable && n.selectable(b.extend({}, {
				selector: k.isDataTable ? ".fixed-left tbody>tr" : "tbody>tr",
				selectClass: "",
				trigger: ".c-id",
				clickBehavior: "multi",
				listenClick: !1,
				select: function(h) {
					k.checkRow(h.target, !0), b.cookie("ajax_dragSelected") || (b.cookie("ajax_dragSelected", "on", {
						expires: config.cookieLife,
						path: config.webRoot
					}), b.ajaxSendScore("dragSelected"))
				},
				unselect: function(e) {
					k.checkRow(e.target, !1)
				},
				rangeStyle: {
					border: "1px solid #006af1",
					backgroundColor: "rgba(50,128,252,0.2)",
					borderRadius: "2px"
				}
			}, b.isPlainObject(q.selectable) ? q.selectable : null)));
			var m = k.$form = n.is("form") ? n : n.find("form");
			m.length && (q.ajaxForm ? m.ajaxForm(b.isPlainObject(q.ajaxForm) ? q.ajaxForm : null) : m.on("click", "[data-form-action]", function() {
				m.attr("action", b(this).data("formAction")).submit()
			})), (q.fixFooter || q.fixHeader) && (k.pageFooterHeight = b("#footer").outerHeight(), k.updateFixUI(), b(window).on("scroll resize", function() {
				k.updateFixUI()
			}).on("sidebar.toggle", function() {
				setTimeout(function() {
					k.updateFixUI()
				}, 200)
			})), q.group && (n.on("click", ".group-toggle", function() {
				k.toggleRowGroup(b(this).closest("tr").data("id"))
			}), b(document).on("click", ".group-collapse-all", function() {
				k.toggleGroups(!1)
			}).on("click", ".group-expand-all", function() {
				k.toggleGroups(!0)
			})), k.defaultStatistic = n.find(".table-statistic").html(), k.updateStatistic(), k.initModals(), k.checkItems = {}, k.updateCheckUI()
		};
	c.prototype.reload = function(k) {
		var h = this,
			m = h.options,
			l = m.replaceId;
		if (!l) {
			return k && k()
		}
		"self" === l && (l = h.$.attr("id"));
		var j = b("<div></div>");
		h.$.addClass("load-indicator loading"), j.load(window.location.href + " #" + l, function() {
			h.$.empty().html(j.children().html()).removeClass("load-indicator loading"), h.$.trigger("beforeTableReload"), h.updateStatistic(), h.initModals(), h.$.datepickerAll();
			var p = h.$.find("tbody>tr"),
				e = !1;
			b.each(h.checkItems, function(o, q) {
				q && (h.checkRow(p.filter('[data-id="' + o + '"]'), !0, !0), e = !0)
			}), e && h.updateCheckUI(), h.$.trigger("tableReload");
			var n = b("#mainMenu>.btn-toolbar>.btn-active-text>.label");
			n.length && n.text(h.getTable().find("tbody:first>tr:not(.table-children)").length), h.$.find('[data-ride="pager"]').pager(), k && k(), m.afterReload && m.afterReload()
		})
	}, c.prototype.initModals = function() {
		var j = this,
			h = j.options,
			l = j.$.find(h.iframeModalTrigger);
		if (l.length) {
			var k = {
				type: "iframe",
				onHide: h.replaceId ?
				function() {
					var e = b.cookie("selfClose");
					(1 == e || h.hot) && (b("#triggerModal").data("cancel-reload", 1), j.reload(function() {
						b.cookie("selfClose", 0)
					}))
				} : null
			};
			l.modalTrigger(k)
		}
	}, c.prototype.getTable = function() {
		var h = this.$;
		if (this.isDataTable) {
			return h.find("div.datatable")
		}
		var j = h.is("table") ? h : h.find("table:not(.fixed-header-copy)").first();
		return j.is(".datatable") && (this.isDataTable = !0, j = h.find("div.datatable")), j
	}, c.prototype.toggleGroups = function(j) {
		var h = this,
			k = {};
		h.$.find("tbody>tr").each(function() {
			var e = b(this).closest("tr").data("id");
			k[e] || h.toggleRowGroup(e, j)
		})
	}, c.prototype.toggleRowGroup = function(h, l) {
		var k = this.$.find('tbody>tr[data-id="' + h + '"]'),
			j = k.filter(".group-summary"),
			e = l === d ? !j.hasClass("hidden") : !! l;
		k.not(".group-summary").toggleClass("hidden", !e), j.toggleClass("hidden", e), b("body").toggleClass("table-group-collapsed", !this.$.find("tbody>tr.group-summary.hidden").length)
	}, c.prototype.updateStatistic = function() {
		var h = this,
			m = h.$.find(".table-statistic");
		if (m.length) {
			if (h.defaultStatistic === d && (h.defaultStatistic = m.html()), h.options.statisticCreator) {
				return void m.html(h.options.statisticCreator(h) || h.defaultStatistic)
			}
			var l = h.statisticCols;
			if (!l && l !== !1) {
				l = {};
				var j = !1;
				h.getTable().find("thead th").each(function(p) {
					var o = b(this),
						q = o.data("statistic");
					q && (j = !0, l[p] = {
						format: q,
						name: o.text()
					})
				}), h.statisticCols = !! j && l
			}
			var e = 0;
			l && b.each(l, function(n) {
				l[n].total = 0, l[n].checkedTotal = 0
			}), h.$.find(h.isDataTable ? ".fixed-left tbody>tr" : "tbody>tr").each(function() {
				var p = b(this),
					o = p.hasClass("checked"),
					q = p.children("td");
				o && e++, l && b.each(l, function(n) {
					var r = parseFloat(q.eq(n).text());
					isNaN(r) && (r = 0), l[n].total += r, o && (l[n].checkedTotal += r)
				})
			});
			var k = [];
			if (e) {
				k.push(h.lang.selectedItems.format(e))
			} else {
				if (h.defaultStatistic) {
					return void m.html(h.defaultStatistic)
				}
			}
			l && b.each(l, function(o) {
				var p = l[o],
					q = p[e ? "checkedTotal" : "total"];
				p.format && (q = p.format.format(q)), k.push(h.lang.attrTotal.format(p.name, q))
			}), m.html(k.join(", "))
		}
	}, c.prototype.updateFixUI = function(m) {
		var j = this,
			o = (new Date).getTime();
		if (!m && (j.lastUpdateCall && clearTimeout(j.lastUpdateCall), !j.lastUpdateTime || o - j.lastUpdateTime < 100)) {
			return void(j.lastUpdateCall = setTimeout(function() {
				j.updateFixUI(!0)
			}, 30))
		}
		if (j.lastUpdateTime = o, j.lastUpdateCall && (clearTimeout(j.lastUpdateCall), j.lastUpdateCall = null), f) {
			var k = j.getTable();
			if (k.parent().is(".table-responsive")) {
				var h = k.find("thead"),
					l = 0;
				h.find("th").each(function() {
					l += b(this).outerWidth()
				}), k.css("min-width", l)
			}
		}
		j.options.fixHeader && !j.isDataTable && j.fixHeader(), j.options.fixFooter && j.fixFooter()
	}, c.prototype.fixHeader = function() {
		var v = this,
			q = v.getTable(),
			m = q.find("thead"),
			k = m[0].getBoundingClientRect(),
			z = v.options.fixFooter,
			y = b.isFunction(z) ? z(k, m) : k.top < ("number" == typeof z ? z : -5),
			j = v.$.find(".fix-table-copy-wrapper"),
			p = q.parent(),
			u = p.is(".table-responsive");
		if (y) {
			if (j.length || (j = b('<div class="fix-table-copy-wrapper" style="overflow: hidden; position:fixed; z-index: 3; top: 0;"></div>').append(b('<table class="fixed-header-copy"></table>').addClass(q.attr("class")).append(m.clone())).insertAfter(q)), u) {
				var x = p[0].getBoundingClientRect();
				j.css({
					left: x.left,
					width: p.width()
				}), j.find(".fixed-header-copy").css({
					left: k.left - x.left,
					position: "relative",
					minWidth: q.width()
				}), p.data("fixHeaderScroll") || p.data("fixHeaderScroll", 1).on("scroll", function() {
					v.fixHeader()
				})
			} else {
				j.css({
					left: k.left,
					width: k.width
				})
			}
			var w = j.find("th");
			m.find("th").each(function(h) {
				w.eq(h).css("width", b(this).outerWidth())
			})
		} else {
			j.remove()
		}
	}, c.prototype.fixFooter = function() {
		var v, q = this,
			m = q.getTable(),
			k = q.$.find(".table-footer");
		if (q.isDataTable) {
			v = m[0].getBoundingClientRect()
		} else {
			var z = m.find("tbody");
			if (!z.length) {
				return
			}
			v = z[0].getBoundingClientRect(), v = z[0].getBoundingClientRect()
		}
		var y = q.options.fixFooter;
		k.toggleClass("fixed-footer", !! j);
		var j = b.isFunction(y) ? y(v, k) : v.bottom > window.innerHeight - 50 - ("number" == typeof y ? y : q.pageFooterHeight || 5);
		k.toggleClass("fixed-footer", !! j), m.toggleClass("with-footer-fixed", !! j), m.trigger("fixFooter", j);
		var p = b("body"),
			u = p.hasClass("body-modal");
		if (j) {
			var x = m.parent(),
				w = x.is(".table-responsive");
			k.css({
				bottom: q.pageFooterHeight || 0,
				left: w ? x[0].getBoundingClientRect().left : v.left,
				width: w ? x.width() : v.width
			}), u && p.css("padding-bottom", 40)
		} else {
			k.css({
				width: "",
				left: 0,
				bottom: 0
			}), u && p.css("padding-bottom", 0)
		}
	}, c.prototype.checkAll = function(j) {
		var h = this,
			k = h.$.find(h.isDataTable ? ".fixed-left tbody>tr" : "tbody>tr");
		k.each(function() {
			h.checkRow(b(this), j, !0)
		}), h.updateCheckUI()
	}, c.prototype.checkRow = function(h, e, l) {
		var k = this;
		k.isDataTable && !h.is(".datatable-row-left") && (h = k.getTable().find('.datatable-row-left[data-index="' + h.data("index") + '"]'));
		var j = h.find('input[type="checkbox"]');
		j.length && (e === d && (e = !j.is(":checked")), k.isDataTable ? k.getTable().find('.datatable-row[data-index="' + h.data("index") + '"]').toggleClass("checked", e) : h.toggleClass("checked", e), this.checkItems[h.data("id")] = e, j.prop("checked", e).trigger("change"), l || k.updateCheckUI())
	}, c.prototype.updateCheckUI = function() {
		var q = this,
			k = q.getTable(),
			v = k.find(q.isDataTable ? ".fixed-left tbody>tr" : "tbody>tr").not(".group-summary"),
			u = !1,
			m = null,
			j = 0,
			p = !1,
			h = v.length;
		v.each(function(r) {
			var e = b(this),
				o = e.find('input[type="checkbox"]');
			p = o.is(":checked");
			var l = q.isDataTable ? k.find('.datatable-row[data-index="' + e.data("index") + '"]') : e;
			l.toggleClass("checked", p), l.toggleClass("row-check-begin", p && !u), m && m.toggleClass("row-check-end", !p && u), p && (j += 1), m = l, u = p, h === r + 1 && l.toggleClass("row-check-end", p)
		}), q.$.toggleClass("has-row-checked", j > 0).find(".check-all").toggleClass("checked", !(!h || j !== h)), q.updateStatistic(), q.options.onCheckChange && q.options.onCheckChange()
	}, c.DEFAULTS = {
		checkable: !0,
		ajaxForm: !1,
		selectable: !0,
		fixHeader: !0,
		fixFooter: !0,
		iframeWidth: 900,
		replaceId: "self",
		hot: !1,
		iframeModalTrigger: ".iframe"
	}, b.fn.table = function(h) {
		return this.each(function() {
			var k = b(this),
				j = k.data(a),
				e = "object" == typeof h && h;
			j || k.data(a, j = new c(this, e)), "string" == typeof h && j[h]()
		})
	}, c.NAME = a, b.fn.table.Constructor = c, b(function() {
		b('[data-ride="table"]').table()
	})
}(jQuery, void 0), function(b, d, a) {
	b.fn._ajaxForm = b.fn.ajaxForm;
	var g = {
		timeout: d.config ? d.config.timeout : 0,
		dataType: "json",
		method: "post"
	},
		f = "";
	b.fn.enableForm = function(h, k, j) {
		return h === a && (h = !0), this.each(function() {
			var e = b(this);
			k || e.find('[type="submit"]').attr("disabled", h ? null : "disabled"), !j && e.hasClass("load-indicator") && e.toggleClass("loading", !h), e.toggleClass("form-disabled", !h)
		})
	}, b.enableForm = function(j, h, k) {
		j === !1 ? b("form").enableForm(j, h, k) : b("form.form-disabled").enableForm(!0, h, k)
	};
	var c = function(j, h, k) {
			"string" == typeof h && (k = h, h = null), k = k || "show", b.zui.messager ? b.zui.messager[k](j, h) : alert(j)
		};
	b.ajaxForm = function(j, n) {
		var e = b(j);
		if (e.length > 1) {
			return e.each(function() {
				b.ajaxForm(this, n)
			})
		}
		b.isFunction(n) && (n = {
			complete: n
		}), n = b.extend({}, g, e.data(), n);
		var m = n.beforeSubmit,
			p = n.error,
			o = n.success,
			k = n.finish;
		delete n.finish, delete n.success, delete n.onError, delete n.beforeSubmit, n = b.extend({
			beforeSubmit: function(q, C, A) {
				A.url=A.url.replace(/(\?|&)ajaxform=true/,"")
				A.url=A.url.indexOf("?")==-1?A.url+"?ajaxform=true" : A.url+"&ajaxform=true";
				C.toggleClass("loading",true);
				if (e.enableForm(!1), (m && m(q, C, A)) !== !1) {
					//修改单独上传临时文件
					var allSize=0,uploadSize=0;//总容量大小与已上传大小
					C.find('input[type="file"]:enabled').each(function(index, el) {
						if(this.files.length==1){
							allSize+=this.files[0].size;
						}
					});
					if(allSize>0){
						fileUpload(C,0,uploadSize,allSize,"");
						return false;
					}
				}				
			},
			success: function(z, s, r) {
				e.toggleClass("loading",false);
				if ((o && o(z, s, r, e)) !== !1) {
					try {
						"string" == typeof z && (z = JSON.parse(z))
					} catch (I) {}
					if (null === z || "object" != typeof z) {
						return z ? alert(z) : c("No response.", "danger")
					}
					var A = n.responser ? b(n.responser) : e.find(".form-responser");
					A.length || (A = b("#responser"));
					var E = z.message,
						q = function() {
							var M = z.callback;
							if (M) {
								var C = M.indexOf("("),
									x = (C > 0 ? M.substr(0, C) : M).split("."),
									v = d,
									y = x[0];
								x.length > 1 && (y = x[1], "top" === x[0] ? v = d.top : "parent" === x[0] && (v = d.parent));
								var w = v[y];
								for(var i=2;i<x.length;i++){
									w =w[x[i]]
								}
							
								if (b.isFunction(w)) {
									var L = [];
									return C > 0 && ")" == M[M.length - 1] && (L = b.parseJSON("[" + M.substring(C + 1, M.length - 1) + "]")), L.push(z), w.apply(e, L)
								}
							}
						};
					if ("success" === z.result) {
						if (e.enableForm(!0, 1), E) {
							var D = e.find('[type="submit"]'),
								B = !1;
							D.length && (D.popover({
								container: "body",
								trigger: "manual",
								content: E,
								tipClass: "popover-in-modal popover-success popover-form-result",
								placement: z.placement || n.popoverPlacement || "right"
							}).popover("show"), setTimeout(function() {
								D.popover("destroy")
							}, n.popoverTime || 2000), B = !0), A.length && (A.html('<span class="small text-success">' + E + "</span>").show().delay(3000).fadeOut(100), B = !0), B || c(E, "success")
						}
						if (k) {
							return k(z, !0, e)
						}
						if ((n.closeModal || z.closeModal) && setTimeout(b.zui.closeModal, n.closeModalTime || 2000), q() === !1) {
							return
						}
						var u = n.locate || z.locate;
						if (u) {
							if ("loadInModal" == u) {
								var K = b(".modal");
								setTimeout(function() {
									K.load(K.attr("ref"), function() {
										b(this).find(".modal-dialog").css("width", b(this).data("width")), b.zui.ajustModalPosition()
									})
								}, 1000)
							} else {
								var G = "reload" == u ? d.location.href : u;
								setTimeout(function() {
									d.location.href = G
								}, 1200)
							}
						}
						var F = n.ajaxReload || z.ajaxReload;
						if (F) {
							var J = b(F);
							J.length && J.load(d.location.href + " " + F, function() {
								J.find('[data-toggle="modal"]').modalTrigger()
							})
						}
					} else {
						if (e.enableForm(), "string" == typeof E) {
							A.length ? A.html('<span class="text-small text-red">' + E + "</span>").show().delay(3000).fadeOut(100) : c(E, "danger")
						} else {
							if ("object" == typeof E) {
								var H = !1,
									l = [];
								b.each(E, function(M, C) {
									var x = b.isArray(C) ? C.join(";") : C,
										w = b("#" + M);
									if (!w.length) {
										return void l.push(x)
									}
									var P = M + "Label",
										O = b("#" + P);
									if (!O.length) {
										var v = w.closest(".input-group").length,
											y = w.closest("td").length;
										O = b('<div id="' + P + '" class="text-danger help-text" />').appendTo(y ? w.closest("td") : v ? w.closest(".input-group").parent() : w.parent())
									}
									O.empty().append(x), w.addClass("has-error");
									var L = function() {
											var h = b("#" + P);
											if (h.length) {
												return h.remove(), w.removeClass("has-error"), !0
											}
										};
									w.on("change input mousedown", L);
									var N = b("#" + M + "_chosen");
									N.length && N.find(".chosen-single,.chosen-choices").addClass("has-error").on("mousedown", function() {
										L() === !0 && b(this).removeClass("has-error")
									}), H || (w.focus(), H = !0)
								}), l.length && c(l.join(";"), "danger")
							}
						}
						if (k) {
							return k(z, !1, e)
						}
						if (q() === !1) {
							return
						}
						var u = n.locate || z.locate;
						if (u) {
							var G = "reload" == u ? d.location.href : u;
							setTimeout(function() {
								d.location.href = G
							}, 2000)
						}
					}
				}
			},
			uploadProgress: function(event,position,total,percentComplete){
				
			},
			error: function(l, h, r) {
				e.toggleClass("loading",false);
				if ((p && p(l, h, r, e)) !== !1) {
					e.enableForm();
					var q = "timeout" == h || "error" == h ? d.lang ? d.lang.timeout : h : l.responseText + h + r;
					c(q, "danger")
				}
			}
		}, n), e._ajaxForm(n).data("zui.ajaxform", !0), e.on("click", "[data-form-action]", function() {
			e.attr("action", b(this).data("formAction")).submit()
		})
	}, b.fn.ajaxForm = function(h) {
		return this.each(function() {
			b.ajaxForm(this, h)
		})
	}, b.fn.setInputRequired = function() {
		return this.each(function() {
			var j = b(this),
				h = j.parent();
			h.is(".input-control,td") ? h.addClass("required") : j.is(".chosen") ? j.attr("required", null).next(".chosen-container").addClass("required") : h.addClass("required"), j.attr("required", null);
			var k = h.closest(".input-group");
			k.length && 1 === k.find(".required,input[required],select[required]").length && k.addClass("required")
		})
	}, b(function() {
		b('.form-ajax,form[data-type="ajax"]').ajaxForm(), setTimeout(function() {
			var e = d.config.requiredFields,
				h = b("form");
			e && (e = e.split(",")), e && e.length && b.each(e, function(j, k) {
				h.find("#" + k).attr("required", "required")
			}), h.find("input[required],select[required],textarea[required]").setInputRequired()
		}, 400), b("#hiddenwin"), b('form[target="hiddenwin"]').on("submit", function() {
			var h = b(this);
			h.data("zui.ajaxform") || h.enableForm(!1).data("disabledTime", (new Date).getTime())
		}).on("click", function() {
			var j = b(this),
				h = j.data("disabledTime");
			h && (new Date).getTime() - h > 10000 && j.enableForm(!0).data("disabledTime", null)
		})
	})
}(jQuery, window, void 0), function(b) {
	var c = "zui.searchList",
		a = function(g, h) {
			if (g && g.length) {
				for (var f = 0; f < g.length; ++f) {
					if (h.indexOf(g[f]) < 0) {
						return !1
					}
				}
			}
			return !0
		},
		d = function(g, n) {
			var k = this;
			k.name = c;
			var f = k.$ = b(g);
			n = k.options = b.extend({}, d.DEFAULTS, this.$.data(), n);
			var m = f.find(n.searchBox);
			m.length && (m.searchBox({
				onSearchChange: function(h) {
					k.search(h)
				},
				onKeyDown: function(o) {
					var p = o.which;
					if (13 === p) {
						var l = k.getActiveItem();
						l.length && (n.onSelectItem ? n.onSelectItem(l) : window.location.href = l.attr("href")), o.preventDefault()
					} else {
						if (38 === p) {
							var l = k.getActiveItem();
							l.removeClass("active");
							for (var q = l.prev(); q.length && !q.is(".search-list-item:not(.hidden)");) {
								q = q.prev()
							}
							q.length || (q = k.getItems().not(".hidden").last()), k.scrollTo(q.addClass("active")), o.preventDefault()
						} else {
							if (40 === p) {
								var l = k.getActiveItem();
								l.removeClass("active");
								for (var h = l.next(); h.length && !h.is(".search-list-item:not(.hidden)");) {
									h = h.next()
								}
								h.length || (h = k.getItems().not(".hidden").first()), k.scrollTo(h.addClass("active")), o.preventDefault()
							}
						}
					}
				}
			}), k.searchBox = m.data("zui.searchBox"), k.search(k.searchBox.getSearch()));
			var e = k.$menu = f.closest(".dropdown-menu");
			if (e.length) {
				k.isDropdown = !0, f.on("click", function(h) {
					b(h.target).closest(n.selector).length || h.stopPropagation()
				});
				var j = e.parent();
				j.on(j.hasClass("dropdown-hover") ? "mouseenter" : "shown.zui.dropdown", function() {
					k.tryLoadRemote(function() {
						setTimeout(function() {
							k.searchBox && k.searchBox.focus()
						}, 50)
					})
				})
			}
			f.on("mouseenter", n.selector, function() {
				f.find(k.options.selector).not(".hidden").removeClass("active"), b(this).addClass("active")
			})
		};
	d.prototype.tryLoadRemote = function(g) {
		var h = this,
			f = h.options;
		f.url || f.ajax ? h.isLoaded ? g() : h.loadRemote(g) : g()
	}, d.prototype.loadRemote = function(g) {
		var f = this,
			h = f.options;
		f.$menu.addClass("load-indicator loading").find(".list-group").remove(), f.isLoaded = !1, b.ajax(b.extend({
			url: h.url,
			type: "GET",
			dataType: "html",
			success: function(l, k, j) {
				var e = b(l);
				e.hasClass("list-group") || (e = b('<div class="list-group"></div>').append(e)), f.$menu.append(e), f.$menu.removeClass("loading"), f.isLoaded = !0, g && g(!0)
			},
			error: function() {
				f.$menu.removeClass("loading").append('<div class="list-group"><div class="text-error has-padding">' + (h.errorText || window.lang && window.lang.timeout) + "</div></div>"), g && g(!1)
			}
		}, h.ajax))
	}, d.prototype.scrollTo = function(e) {
		e.length && e[0].scrollIntoView({
			behavior: "smooth"
		})
	}, d.prototype.getItems = function() {
		return this.$.find(this.options.selector).addClass("search-list-item")
	}, d.prototype.getActiveItem = function() {
		return this.getItems().filter(".active:first")
	}, d.prototype.search = function(h) {
		var k = this,
			j = void 0 === h || null === h || "" === h,
			g = k.getItems().removeClass("active");
		if (j) {
			g.removeClass("hidden")
		} else {
			var f = b.trim(h).split(" ");
			g.each(function() {
				var l = b(this),
					m = l.text() + " " + (l.data("key") || l.data("filter"));
				l.toggleClass("hidden", !a(f, m))
			})
		}
		k.scrollTo(g.not(".hidden").first().addClass("active"))
	}, d.DEFAULTS = {
		selector: ".list-group a",
		searchBox: ".search-box",
		onSelectItem: null
	}, b.fn.searchList = function(e) {
		return this.each(function() {
			var h = b(this),
				g = h.data(c),
				f = "object" == typeof e && e;
			g || h.data(c, g = new d(this, f)), "string" == typeof e && g[e]()
		})
	}, d.NAME = c, b.fn.searchList.Constructor = d, b(function() {
		b('[data-ride="searchList"]').searchList()
	})
}(jQuery), function(b) {
	var c = "zui.labelSelector",
		a = function(f, e) {
			var d = this;
			d.name = c, d.$ = b(f), e = d.options = b.extend({}, a.DEFAULTS, this.$.data(), e), d.$.hide(), d.update()
		};
	a.prototype.select = function(d) {
		d += "", this.$wrapper.find(".label.active").removeClass("active"), this.$wrapper.find('.label[data-value="' + d + '"]').addClass("active"), this.$.val(d).trigger("change")
	}, a.prototype.update = function() {
		var g = this,
			d = g.options,
			j = g.$wrapper;
		if (!j) {
			if (d.wrapper) {
				j = b(d.wrapper)
			} else {
				var h = g.$.next();
				j = h.hasClass(".label-selector") ? h : b('<div class="label-selector"></div>')
			}
			j.parent().length || g.$.after(j), g.$wrapper = j, j.on("click", ".label", function(e) {
				var l = g.$.val(),
					k = b(this).data("value");
				g.hasEmptyValue !== !1 && k == l && (k = g.hasEmptyValue), g.select(k), e.preventDefault()
			})
		}
		j.empty();
		var f = g.$.val();
		g.hasEmptyValue = !1, g.$.children("option").each(function() {
			var m = b(this),
				n = {
					label: m.text(),
					value: m.val()
				},
				k = "" === n.value || "0" === n.value,
				l = b(d.labelTemplate || '<span class="label"></span>');
			d.labelClass && !k && l.addClass(d.labelClass), d.labelCreator ? l = d.labelCreator(l) : (l.data("option", n).attr("data-value", n.value), k && !n.label ? l.addClass("empty").append('<i class="icon icon-close"></i>') : l.text(n.label).toggleClass("active", f === n.value)), j.append(l)
		})
	}, a.DEFAULTS = {}, b.fn.labelSelector = function(d) {
		return this.each(function() {
			var g = b(this),
				f = g.data(c),
				e = "object" == typeof d && d;
			f || g.data(c, f = new a(this, e)), "string" == typeof d && f[d]()
		})
	}, a.NAME = c, b.fn.labelSelector.Constructor = a, b(function() {
		b('[data-provide="labelSelector"]').labelSelector()
	})
}(jQuery), function(d) {
	var h = "zui.fileInput",
		c = d.BYTE_UNITS = {
			B: 1,
			KB: 1024,
			MB: 1048576,
			GB: 1073741824,
			TB: 1099511627776
		},
		k = d.formatBytes = function(a, l, m) {
			return void 0 === l && (l = 2), m || (m = a < c.KB ? "B" : a < c.MB ? "KB" : a < c.GB ? "MB" : a < c.TB ? "GB" : "TB"), (a / c[m]).toFixed(l) + m
		},
		j = function(a) {
			if ("string" == typeof a) {
				a = a.toUpperCase();
				var l = a.replace(/\d+/, "");
				a = parseFloat(a.replace(l, "")), a *= c[l] || c[l + "B"], a = Math.floor(a)
			}
			return a
		},
		f = function(n, m) {
			var p = this;
			p.name = h;
			var e = p.$ = d(n);
			m = p.options = d.extend({}, f.DEFAULTS, this.$.data(), m), m.fileMaxSize && "string" == typeof m.fileMaxSize && (m.fileMaxSize = j(m.fileMaxSize));
			var o = p.$input = e.find('input[type="file"]');
			e.on("click", ".file-input-btn", function() {
				o.trigger("click")
			}).on("click", ".file-input-rename", function() {
				p.oldName = e.addClass("edit").find(".file-editbox").focus().val()
			}).on("click", ".file-input-delete", function() {
				o.val(""), p.update(), m.onDelete && m.onDelete(p)
				d('input[value="'+d(this).attr("name")+'"]').remove();
			}).on("click", ".file-name-cancel", function() {
				e.removeClass("edit").find(".file-editbox").focus().val(p.oldName)
			}).on("click", ".file-name-confirm", function() {
				var l = e.find(".file-editbox"),
					a = d.trim(l.val());
				a.length ? e.removeClass("edit").find(".file-title").text(a) : l.focus()
			}).on("change input paste", ".file-editbox", function() {
				var a = d(this);
				oldName=a.parent().parent().find('.file-input-normal .file-input-rename').attr("name");
				if(oldName){
					s=oldName.split("_。。_");
					newname=a.val()+"_。。_"+s[1];
					a.parent().parent().find('.file-input-normal .file-input-rename').attr("name",newname);
					a.parent().parent().find('.file-input-normal .file-input-delete').attr("name",newname);
					d('input[value="'+oldName+'"]').val(newname);
				}
				
				a.attr("size", Math.max(5, a.val().length))
			}), o.on("change", function() {
				var a = p.getFile();
				a && m.fileMaxSize && a.size > m.fileMaxSize && (o.val(""), (window.bootbox || window).alert(m.fileSizeError.format(k(m.fileMaxSize)))), p.update()
			}), p.update()
		};
	f.prototype.getFile = function() {
		var a = this.$input.prop("files");
		return a && a[0]
	}, f.prototype.update = function(l) {
		var n = this,
			a = n.$,
			p = n.getFile(),
			m = !p;
		a.toggleClass("normal", !m).toggleClass("empty", m), p ? (n.oldName = p.name, a.find(".file-title").text(p.name).attr("title", p.name),a.find(".file-type").attr("name", p.name+'_type'), a.find(".file-size").text(k(p.size)), a.find(".file-editbox").val(p.name).attr("size", p.name.length), n.options.onSelect && n.options.onSelect(p, n)) : a.find(".file-editbox").val("")
	}, f.DEFAULTS = {
		fileMaxSize: 0,
		fileSizeError: "无法上传大于 {0} 的文件。"
	}, d.fn.fileInput = function(a) {
		return this.each(function() {
			var m = d(this),
				l = m.data(h),
				e = "object" == typeof a && a;
			l || m.data(h, l = new f(this, e)), "string" == typeof a && l[a]()
		})
	}, f.NAME = h, d.fn.fileInput.Constructor = f, d(function() {
		d('[data-provide="fileInput"]').fileInput()
	});
	var b = "zui.fileInputList",
		g = function(l, a) {
			var p = this;
			p.name = b;
			var m = p.$ = d(l);
			a = p.options = d.extend({}, g.DEFAULTS, this.$.data(), a), p.$template = m.find(".file-input").detach(), p.add()
		};
	g.prototype.add = function() {
		var l = this,
			m = l.options,
			a = l.$template.clone();
		"before" === m.appendWay ? l.$.prepend(a) : l.$.append(a), a.fileInput({
			fileMaxSize: m.eachFileMaxSize,
			fileSizeError: m.fileSizeError,
			onDelete: function(n) {
				n.$.remove(), l.options.onDelete && l.options.onDelete(n, l)
			},
			onSelect: function(o, n) {
				l.add(), l.options.onSelect && l.options.onSelect(o, n, l)
			}
		})
	}, g.DEFAULTS = {
		fileMaxSize: 0,
		eachFileMaxSize: 0,
		appendWay: "after",
		fileSizeError: "无法上传大于 {0} 的文件。"
	}, d.fn.fileInputList = function(a) {
		return this.each(function() {
			var e = d(this),
				m = e.data(b),
				l = "object" == typeof a && a;
			m || e.data(b, m = new g(this, l)), "string" == typeof a && m[a]()
		})
	}, g.NAME = b, d.fn.fileInputList.Constructor = g, d(function() {
		d('[data-provide="fileInputList"]').fileInputList()
	})
}(jQuery), function(G) {
	window.config || (window.config = {}), G.createLink = window.createLink = function(d, h, m) {
		var s="/"+d+"/"+h;
		if(typeof m=="object"){
			s+="?"
			var g=[];
			for(var i in m){
				g.push(i+"="+m[i])
			}
			s+=g.join("&")
		}
		if(typeof m=="string"){
			s+="?"+m
		}
		return s;
	}, G(function() {
		var a = G("#main,#mainContent,#mainRow,.auto-fade-in");
		a.length && a.hasClass("fade") && setTimeout(function() {
			a.addClass("in")
		}, a.data("fadeTime") || 200)
	}), G.ajaxSendScore = function(a) {
		G.get(G.createLink("score", "ajax", "method=" + a))
	};
	var W = function(b) {
			var c = 0;
			if (b) {
				var a = b.split(":");
				c += 60 * parseInt(a[0]), c += parseInt(a[1])
			}
			return c
		},
		N = function(b) {
			b %= 1440;
			var c = Math.floor(b / 60),
				a = b % 60;
			return c < 10 && (c = "0" + c), a < 10 && (a = "0" + a), c + ":" + a
		},
		M = function(b) {
			if ("string" == typeof b && (b = W(b)), "number" == typeof b) {
				if (b < 100000) {
					var a = new Date;
					a.setHours(Math.floor(b / 60) % 24), a.setMinutes(b % 60), b = a
				} else {
					b = new Date(b)
				}
			}
			return b
		},
		I = function(o, f) {
			for (var p = f ? M(f) : new Date, n = p.getHours(), b = 10 * Math.floor(p.getMinutes() / 10) + 10, e = 0; e < 24; ++e) {
				var g = (e + n) % 24;
				if (!(g < 5)) {
					for (var m = 0; m < 6; ++m) {
						var k = N(60 * g + 10 * m + b);
						o.append('<option value="' + k + '">' + k + "</option>")
					}
				}
			}
			o.val() || (time = W(p.format("hh:mm")), time = time - time % 10 + 10, o.val(N(time)))
		};
	G.fn.timeSpanControl = function(a) {
		return this.each(function() {
			var e = G(this),
				k = G.extend({}, a, e.data()),
				b = e.find('[name="begin"],.control-time-begin'),
				g = e.find('[name="end"],.control-time-end'),
				n = function() {
					var d = b.val();
					if (e.find(".hide-empty-begin").toggleClass("hide", !d), d) {
						var c = N(W(d) + 30);
						g.find('option[value="' + c + '"]').length && g.val(c), k.onChange && k.onChange(g, c)
					}
				};
			if (e.data("timeSpanControlInit")) {
				if (k.begin) {
					var m = M(k.begin).format("hh:mm");
					b.find('option[value="' + m + '"]').length && b.val(m), k.onChange && k.onChange(b, m)
				}
				if (k.end) {
					var f = M(k.end).format("hh:mm");
					g.find('option[value="' + f + '"]').length && g.val(f), k.onChange && k.onChange(g, f)
				}
			} else {
				b.on("change", n), I(b, k.begin), I(g, k.end), e.data("timeSpanControlInit", !0)
			}
			k.end || n()
		})
	}, G.timeSpanControl = {
		convertTimeToNum: W,
		convertNumToTime: N,
		initTimeSelect: I,
		createTime: M
	};
	var aa = G.setSearchType = function(f, c) {
			var h = G("#searchType");
			f || (f = h.val()), f = f || "bug", h.val(f);
			var g = G("#searchTypeMenu");
			g.find("li.selected").removeClass("selected");
			var d = g.find('a[data-value="' + f + '"]'),
				b = d.text();
			d.parent().addClass("selected"), G("#searchTypeName").text(b), c || G("#searchInput").focus()
		};
	G.gotoObject = function(b, a) {
		b || (b = G("#searchType").val()), a || (a = G("#searchInput").val()), a && b && (window.location.href = G.createLink(b, "testsuite" === b ? "library" : "view", "id=" + a))
	}, G(function() {
		aa(null, !0), G(document).on("keydown", function(a) {
			a.ctrlKey && 71 === a.keyCode && (G("#searchInput").val("").focus(), a.stopPropagation(), a.preventDefault())
		})
	}), G.removeAnchor = window.removeAnchor = function(a) {
		var b = a.lastIndexOf("#");
		return b > -1 ? a.substr(0, b) : a
	}, G.refreshPage = function() {
		location.href = removeAnchor(location.href)
	}, G.selectLang = window.selectLang = function(a) {
		G.cookie("lang", a, {
			expires: config.cookieLife,
			path: config.webRoot
		}), G.ajaxSendScore("selectLang"), G.refreshPage()
	}, G.selectTheme = window.selectTheme = function(a) {
		G.cookie("theme", a, {
			expires: config.cookieLife,
			path: config.webRoot
		}), G.ajaxSendScore("selectTheme"), G.refreshPage()
	}, G.chosenDefaultOptions = {
		disable_search_threshold: 1,
		compact_search: !0,
		allow_single_deselect: !0,
		placeholder_text_single: " ",
		placeholder_text_multiple: " ",
		search_contains: !0,
		drop_direction: function() {
			var c = G(this.container).closest(".table-responsive:not(.scroll-none)");
			if (c.length) {
				if (this.drop_directionFixed) {
					return this.drop_directionFixed
				}
				var a = "down",
					f = this.container.find(".chosen-drop"),
					d = this.container.position(),
					b = f.outerHeight();
				return d.top >= b && d.top + b < c.outerHeight() && (a = "up"), this.drop_directionFixed = a, a
			}
			return "auto"
		}
	}, G.chosenSimpleOptions = G.extend({}, G.chosenDefaultOptions, {
		disable_search_threshold: 6
	}), G.fn._chosen = G.fn.chosen, G.fn.chosen = function(a) {
		return "string" == typeof a ? this._chosen(a) : this.each(function() {
			var b = G(this).addClass("chosen-controled");
			return b._chosen(G.extend({}, b.hasClass("chosen-simple") ? G.chosenSimpleOptions : G.chosenDefaultOptions, b.data(), a))
		})
	}, G(function() {
		G(".chosen,.chosen-simple").each(function() {
			var a = G(this);
			a.closest(".template").length || a.chosen()
		})
	}), G.extend(G.fn.pager.Constructor.DEFAULTS, {
		maxNavCount: 8,
		prevIcon: "icon-angle-left",
		nextIcon: "icon-angle-right",
		firstIcon: "icon-first-page",
		lastIcon: "icon-last-page",
		navEllipsisItem: "…",
		menuDirection: "dropup",
		pageSizeOptions: [5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 100, 200, 500, 1000, 2000],
		elements: ["total_text", "size_menu", "first_icon", "prev_icon", '<div class="pager-label"><strong>{page}</strong>/<strong>{totalPage}</strong></div>', "next_icon", "last_icon"],
		onPageChange: function(b, a) {
			b.recPerPage !== a.recPerPage && G.cookie(this.options.pageCookie, b.recPerPage, {
				expires: config.cookieLife,
				path: config.webRoot
			}), b.recPerPage !== a.recPerPage && (window.location.href = this.createLink())
		}
	}), G.zui.Messager.DEFAULTS.cssClass = "messagger-zt", G.fn.reverseOrder = function() {
		return this.each(function() {
			var a = G(this);
			a.prependTo(a.parent())
		})
	};
	var K = function(h, d) {
			var m = G(h);
			d = G.extend({}, m.data(), d);
			var k = m.find(".histories-list"),
				f = !0,
				c = !1;
			m.on("click", ".btn-reverse", function() {
				k.children("li").reverseOrder(), f = !f, G(this).find(".icon").toggleClass("icon-arrow-up", f).toggleClass("icon-arrow-down", !f)
			}).on("click", ".btn-expand-all", function() {
				var a = G(this).find(".icon");
				c = !c, a.toggleClass("icon-plus", !c).toggleClass("icon-minus", c), k.children("li").toggleClass("show-changes", c)
			}).on("click", ".btn-expand", function() {
				G(this).closest("li").toggleClass("show-changes")
			}).on("click", ".btn-strip", function() {
				var a = G(this),
					p = a.find(".icon"),
					l = p.hasClass("icon-code");
				p.toggleClass("icon-code", !l).toggleClass("icon-text", l), a.attr("title", l ? d.original : d.textdiff), a.closest("li").toggleClass("show-original", l)
			}), k.find(".btn-strip").attr("title", d.original);
			var g = m.find(".modal-comment").modal({
				show: !1
			}).on("shown.zui.modal", function() {
				var a = g.find("#comment");
				a.length && (a.focus(), window.editor && window.editor.comment && window.editor.comment.focus())
			}).on("show.zui.modal", function() {
				var a = g.find("#comment");
				a.length && !a.data("keditor") && G.fn.kindeditor && a.kindeditor()
			});
			m.on("click", ".btn-comment", function(a) {
				g.modal("toggle"), a.preventDefault()
			}).on("click", ".btn-edit-comment,.btn-hide-form", function() {
				G(this).closest("li").toggleClass("show-form")
			});
			var b = m.find(".comment-edit-form");
			b.ajaxForm({
				success: function(l, o, a, p) {
					setTimeout(function() {
						b.closest("li").removeClass("show-form")
					}, 2000)
				}
			})
		};
	G.fn.histories = function(a) {
		return this.each(function() {
			K(this, a)
		})
	}, G(function() {
		G(".histories").histories()
	});
	var P = 0,
		R = 0;
	G.toggleSidebar = function(c) {
		var a = G("#sidebar");
		if (a.length) {
			var f = G("main");
			if (void 0 === c) {
				c = f.hasClass("hide-sidebar")
			} else {
				if (c && !f.hasClass("hide-sidebar")) {
					return
				}
			}
			f.toggleClass("hide-sidebar", !c), clearTimeout(P), G.zui.store.set(R, c);
			var d = a.children(".cell"),
				e = a.children(".cell2"),
				b = {
					overflow: "visible",
					maxHeight: "initial"
				};
			c ? (a.addClass("showing"), P = setTimeout(function() {
				a.removeClass("showing"), a.trigger("sidebar.toggle", c)
			}, 210)) : (a.trigger("sidebar.toggle", c), G(window).width() < 1900 && (b = {
				overflow: "hidden",
				maxHeight: G(window).height() - 45
			})), d.css(b),e.css(b)
		}
	};
	var Y = G.initSidebar = function() {
			var b = G("#sidebar");
			if (b.length) {
				if (b.data("init")) {
					return !0
				}
				R = "sidebar:" + (b.data("id") || config.currentModule + "/" + config.currentMethod);
				var a = G("main");
				a.on("click", ".sidebar-toggle", function() {
					G.toggleSidebar(a.hasClass("hide-sidebar"))
				});
				var d = G.zui.store.get(R, b.data("hide") !== !1);
				d === !1 && b.addClass("no-animate"), G.toggleSidebar(d), d === !1 && setTimeout(function() {
					b.removeClass("no-animate")
				}, 500);
				var c = function() {
						var e = b.find(".sidebar-toggle");
						if (e.length) {
							var h = e[0].getBoundingClientRect(),
								g = G(window).height(),
								f = Math.max(0, Math.floor(Math.min(g - 40, h.top + h.height) - Math.max(h.top, 0)) / 2) + (h.top < 0 ? 0 - h.top : 0);
							e.find(".icon").css("top", f)
						}
					};
				return c(), b.on("sidebar.toggle", c), G(window).on("resize", c).on("scroll", c), b.data("init", 1), !0
			}
		};
	Y() || G(Y), G.toggleQueryBox = function(b, a) {
		var c = G(a || "#queryBox");
		c.length && (void 0 === b && (b = !c.hasClass("show")), c.toggleClass("show", !! b), c.data("init") || (c.addClass("load-indicator loading").data("init", 1), G.get(G.createLink("search", "buildForm",{module:config.currentModule,method:config.currentMethod,queryID:(config.queryID?config.queryID:''),"onlybody":"yes"}), function(d) {
			c.html(d).removeClass("loading")
		})), G(".querybox-toggle").toggleClass("querybox-opened", b))
	}, G(function() {
		var a = G("#queryBox");
		a.length && (G(document).on("click", ".querybox-toggle", function() {
			G.toggleQueryBox()
		}), a.hasClass("show") && G.toggleQueryBox(!0))
	}), G.extend(G.fn.colorPicker.Constructor.DEFAULTS, {
		colors: ["#3DA7F5", "#75C941", "#2DBDB2", "#797EC9", "#FFAF38", "#FF4E3E"]
	}), window.setCheckedCookie = function() {
		var b = [],
			a = G('#mainContent .main-table tbody>tr input[type="checkbox"]:checked');
		a.each(function() {
			var c = parseInt(G(this).val(), 10);
			NaN !== c && b.push(c)
		}), G.cookie("checkedItem", b.join(","), {
			expires: config.cookieLife,
			path: config.webRoot
		})
	}, G.extend(G.fn.modal.bs.Constructor.DEFAULTS, {
		scrollInside: !0,
		backdrop: "static",
		headerHeight: 100
	}), G.extend(G.zui.ModalTrigger.DEFAULTS, {
		scrollInside: !0,
		backdrop: "static",
		headerHeight: 40
	}), G.fn.initIframeModal = function() {
		return this.each(function() {
			var b = G(this);
			if (!b.parents('[data-ride="table"],.skip-iframe-modal').length) {
				var a = {
					type: "iframe"
				};
				b.hasClass("export") && G.extend(a, {
					width: 800,
					shown: setCheckedCookie
				}, b.data()), b.modalTrigger(a)
			}
		})
	}, G(function() {
		G("a.iframe,.export").initIframeModal()
	});
	var X = function() {
			var f, c, h = G(this),
				g = G.extend({
					limitSize: 40,
					suffix: "…"
				}, h.data()),
				d = h.text();
			if (d.length > g.limitSize) {
				f = d, c = d.substr(0, g.limitSize) + g.suffix, h.text(c).addClass("limit-text-on");
				var b = g.toggleBtn ? G(g.toggleBtn) : h.next(".text-limit-toggle");
				b.text(b.data("textExpand")), b.on("click", function() {
					var a = h.toggleClass("limit-text-on").hasClass("limit-text-on");
					h.text(a ? c : f), b.text(b.data(a ? "textExpand" : "textCollapse"))
				})
			} else {
				(g.toggleBtn ? G(g.toggleBtn) : h.next(".text-limit-toggle")).hide()
			}
		};
	G.fn.textLimit = function() {
		return this.each(X)
	}, G(function() {
		G(".text-limit").textLimit()
	}), G.fixedTableHead = window.fixedTableHead = function(f, c) {
		var h = G(f);
		if (h.is("table") || (h = h.find("table")), h.length) {
			var g = G(c || window),
				d = null,
				b = function() {
					var s = h.children("thead"),
						n = s[0].getBoundingClientRect(),
						u = h.next(".fixed-head-table");
					if (n.top < 0) {
						var m = s.width();
						if (u.length) {
							if (d !== m) {
								d = m;
								var p = u.find("th");
								s.find("th").each(function(a) {
									p.eq(a).width(G(this).width())
								})
							}
						} else {
							var u = G("<table class='table fixed-head-table' style='position:fixed; top: 0;'></table>").addClass(h.attr("class")),
								k = s.clone(),
								p = k.find("th");
							s.find("th").each(function(a) {
								p.eq(a).width(G(this).width())
							}), u.append(k).insertAfter(h)
						}
						u.css({
							left: n.left,
							width: n.width
						}).show()
					} else {
						u.hide()
					}
				};
			g.on("scroll", b).on("resize", b), b()
		}
	}, G(document).on("click", "tr[data-url]", function() {
		var b = G(this),
			a = b.data("href") || b.data("url");
		a && (window.location.href = a)
	}), "yes" === config.onlybody && self === parent && (window.location.href = window.location.href.replace("?onlybody=yes", "").replace("&onlybody=yes", "")), G(function() {
		G("body").addClass("m-{currentModule}-{currentMethod}".format(config))
	});
	var F, L, V, U, O, E = function() {
			F || (F = G("#subNavbar"), L = G("#pageNav"), V = G("#pageActions"), U = F.children(".nav"), O = U.outerWidth());
			var g = F.outerWidth(),
				c = L.outerWidth() || 0,
				k = V.outerWidth() || 0;
			if (c = c ? c + 15 : 0, k = k ? k + 15 : 0, !c && !k) {
				return void U.css({
					maxWidth: null,
					left: null,
					position: "static"
				})
			}
			var h = Math.max(300, g - c - k),
				d = Math.min(h, O),
				b = (g - d) / 2,
				f = c && b < c ? c : k && b < k ? g - d - k : 0;
			U.css({
				maxWidth: h,
				left: f ? f - b : 0,
				position: "relative"
			})
		},
		z = function() {
			G.cookie("windowWidth", window.innerWidth), G.cookie("windowHeight", window.innerHeight), E()
		};
	G(z), G(window).on("resize", z);
	var Z = function() {
			var a = G("#back").attr("href");
			a && (window.location.href = a)
		},
		B = function() {
			G.cookie("ajax_lastNext") || (G.cookie("ajax_lastNext", "on", {
				expires: config.cookieLife,
				path: config.webRoot
			}), G.ajaxSendScore("lastNext"))
		},
		A = function() {
			var a = G("#prevPage").attr("href");
			a && (window.location.href = a), B()
		},
		J = function() {
			var a = G("#nextPage").attr("href");
			a && (window.location.href = a), B()
		};
	G(document).on("keydown", function(a) {
		a.altKey && 38 === a.keyCode ? Z() : 37 === a.keyCode ? A() : 39 === a.keyCode && J()
	}), G.fn.tree.Constructor.DEFAULTS.initialState = "preserve", G.closeModal = function(b, a, c) {
		G.zui.closeModal(c, b, a)
	}, G.getThemeColor = function(b) {
		if (!G.themeColor) {
			var a = G("#mainHeader");
			a.length && (G.themeColor = {
				primary: a.css("border-top-color"),
				pale: a.css("border-bottom-color"),
				secondary: a.css("background-color")
			})
		}
		return b ? G.themeColor && G.themeColor[b] : G.themeColor
	};
	var ab = function(c) {
			var a, f, d = G(c),
				b = d.children(".input-group-addon,.form-control:not(.chosen-controled),.chosen-container,.btn,.input-control,.input-group-btn,.datepicker-wrapper");
			b.each(function(p) {
				var s = G(this),
					k = s.is(".input-group-addon") ? "addon" : s.is(".chosen-container") ? "chosen" : s.is(".btn") ? "btn" : s.is(".input-control,.datepicker-wrapper") ? "insideInput" : s.is(".input-group-btn") ? "insideBtn" : "input",
					n = !a,
					g = p === b.length - 1,
					m = {};
				m.borderTopLeftRadius = 0, m.borderBottomLeftRadius = 0, m.borderTopRightRadius = 0, m.borderBottomRightRadius = 0, n && ("addon" === k && (m.borderLeftWidth = 1), m.borderTopLeftRadius = 2, m.borderBottomLeftRadius = 2), g && ("addon" === k && (m.borderRightWidth = 1), m.borderTopRightRadius = 2, m.borderBottomRightRadius = 2), f && ("chosen" !== f && "input" !== f && "btn" !== f && "insideInput" !== f && "insideBtn" !== f || "chosen" !== k && "input" !== f && "btn" !== k && "insideInput" !== k && "insideBtn" !== k || (m.borderLeftColor = "transparent")), ("insideBtn" === k ? s.find(".btn") : "insideInput" === k ? s.find(".form-control") : "chosen" === k ? s.find(".chosen-single,.chosen-choices") : s).css(m), a = s, f = k
			})
		};
	G.fn.fixInputGroup = function() {
		return this.each(function() {
			ab(this)
		})
	};
	var Q = function() {
			var g = G(".main-actions>.btn-toolbar");
			if (g.length) {
				var c, k, h = !1,
					d = null,
					b = g.children(),
					f = b.length;
				for (b.each(function(a) {
					c = G(this), k = c.is(".divider"), k && !d && c.hide(), h || k || (h = !0), d = k ? null : c, !k || a !== f - 1 && 0 !== a || c.hide()
				}); c.length && c.is(".divider");) {
					c = c.hide().prev()
				}
				h || g.hide()
			}
		};
	G(function() {
		G(".input-group,.btn-group").fixInputGroup(), Q()
	}), window.holders && G.each(window.holders, function(b) {
		var a = G("#" + b);
		a.length && a.is("input") && a.attr("placeholder", window.holders[b])
	});
	var j = function() {
			var g, c = "en" == config.clientLang ? "http://www.zentao.pm/book/zentaomanual/8.html?fullScreen=zentao" : "http://www.zentao.net/book/zentaopmshelp.html?fullScreen=zentao",
				k = G("#navbar > .nav").first(),
				h = 10000,
				d = function() {
					clearTimeout(g), G("#helpContent").removeClass("show-error")
				},
				b = G.openHelp = function() {
					d(), k.children("li.active:not(#helpMenuItem)").removeClass("active").addClass("close-help-tab"), G("#helpMenuItem").addClass("active");
					var e = G("#helpContent");
					if (e.length) {
						if (G("body").hasClass("show-help-tab")) {
							return void G("#helpIframe").get(0).contentWindow.location.replace(c)
						}
					} else {
						e = G('<div id="helpContent"><div class="load-error text-center"><h4 class="text-danger">' + lang.timeout + '</h4><p><a href="###" class="open-help-tab"><i class="icon icon-arrow-right"></i> ' + c + '</a></p></div><iframe id="helpIframe" name="helpIframe" src="' + c + '" frameborder="no" allowtransparency="true" scrolling="auto" hidefocus="" style="width: 100%; height: 100%; left: 0px;"></iframe></div>'), G("#header").after(e);
						var l = G("#helpIframe").get(0);
						g = setTimeout(function() {
							G("#helpContent").addClass("show-error")
						}, h), l.onload = l.onreadystatechange = function() {
							this.readyState && "complete" != this.readyState || d()
						}
					}
					G("body").addClass("show-help-tab")
				},
				f = G.closeHelp = function() {
					G("body").removeClass("show-help-tab"), G("#helpMenuItem").removeClass("active"), k.find("li.close-help-tab").removeClass("close-help-tab").addClass("active").find("a").focus()
				};
			G(document).on("click", ".open-help-tab", function(l) {
				var a = G("#helpMenuItem");
				a.length || (a = G('<li id="helpMenuItem"><a href="javascript:;" class="open-help-tab">' + G(this).text() + ' <i class="icon icon-close close-help-tab icon-sm"></i></a></li>'), k.append('<li class="divider"></li>').append(a)), b(), l.preventDefault()
			}).on("click", ".close-help-tab", function(a) {
				f(), a.stopPropagation(), a.preventDefault()
			})
		};
	G(j), G(function() {
		var b = G(".table-responsive"),
			a = function() {
				b.each(function() {
					this.scrollHeight - 3 <= this.clientHeight && this.scrollWidth - 3 <= this.clientWidth ? G(this).addClass("scroll-none").css("overflow", "visible") : G(this).removeClass("scroll-none").css("overflow", "auto")
				})
			};
		b.length && (a(), G(window).on("resize", a))
	});
	var q = function() {
			var a = this.value ? this.scrollHeight + 2 + "px" : "32px";
			this.style.height = "auto", this.style.height = a, G(this).closest("tr").find("textarea").each(function() {
				this.style.height = a
			})
		};
	G.autoResizeTextarea = function(a) {
		G(a).each(q)
	}, G(function() {
		G("textarea.autosize").each(q), G(document).on("input keyup paste change", "textarea.autosize", q)
	}), G(function() {
		var a = G("#dropMenu");
		a.length && a.on("click", ".toggle-right-col", function(b) {
			a.toggleClass("show-right-col"), b.stopPropagation(), b.preventDefault()
		})
	});
	var H = "undefined" != typeof InstallTrigger;
	G.zui.browser.firefox = H, G("html").toggleClass("is-firefox", H).toggleClass("not-firefox", !H), G(function() {
		var b = G("#mainContent>.main-col"),
			a = b.children(".main-actions"),
			d = a.prev();
		if (a.length && d.length) {
			G('<div class="main-actions-holder"></div>').css("height", a.outerHeight()).insertAfter(a);
			var c = function() {
					var f = d[0].getBoundingClientRect(),
						g = f.top + f.height + 120 > G(window).height();
					G("body").toggleClass("main-actions-fixed", g), g && a.width(d.width())
				};
			G.resetToolbarPosition = c, c(), G(window).on("resize scroll", c)
		}
	}), G(document).on("show.zui.modal", function() {
		G("body.body-modal").length && window.parent && window.parent !== window && window.parent.$("body").addClass("hide-modal-close")
	}).on("hidden.zui.modal", function() {
		G("body.body-modal").length && window.parent && window.parent !== window && window.parent.$("body").removeClass("hide-modal-close")
	})
}(jQuery);

function setPing() {
	$("#hiddenwin").attr("src", createLink("misc", "ping"))
}
function setForm() {
	var a = false;
	$("form").submit(function() {
		submitObj = $(this).find(":submit");
		if ($(submitObj).size() >= 1) {
			var c = submitObj.prop("tagName") == "BUTTON";
			submitLabel = c ? $(submitObj).html() : $(submitObj).attr("value");
			$(submitObj).attr("disabled", "disabled");
			var b = submitObj.attr("data-submitting") || lang.submitting;
			if (c) {
				submitObj.text(b)
			} else {
				$(submitObj).attr("value", b)
			}
			a = true
		}
	});
	$("body").click(function() {
		if (a) {
			$(submitObj).removeAttr("disabled");
			if (submitObj.prop("tagName") == "BUTTON") {
				submitObj.text(submitLabel)
			} else {
				$(submitObj).attr("value", submitLabel)
			}
			$(submitObj).removeClass("button-d")
		}
		a = false
	})
}
function setFormAction(b, c, f) {
	$form = typeof(f) == "undefined" ? $("form") : $(f).closest("form");
	if (c) {
		$form.attr("target", c)
	} else {
		$form.removeAttr("target")
	}
	$form.attr("action", b);
	var d = navigator.userAgent;
	var a = d.indexOf("AppleWebKit") > -1 && d.indexOf("Safari") > -1 && d.indexOf("Chrome") < 0;
	if (a) {
		var e = "checkbox-fix-" + $.zui.uuid();
		$form.find("[data-fix-checkbox]").remove();
		$form.find('input[type="checkbox"]:not(.rows-selector)').each(function() {
			var g = $(this);
			var h = e + g.val();
			g.clone().attr("data-fix-checkbox", h).css("display", "none").after('<div id="' + h + '"/>').appendTo($form)
		})
	}
	$form.submit()
}
function setImageSize(b, a) {
	if (!a) {
		bodyWidth = $("body").width();
		a = bodyWidth - 470
	}
	if ($(b).width() > a) {
		$(b).attr("width", a)
	}
	$(b).wrap('<a href="' + $(b).attr("src") + '" target="_blank"></a>')
}
function setMailto(a, b) {
	link = createLink("user", "ajaxGetContactUsers", "listID=" + b);
	$.get(link, function(c) {
		$("#" + a).replaceWith(c);
		$("#" + a + "_chosen").remove();
		$("#" + a).chosen()
	})
}
function ajaxGetContacts(a) {
	link = createLink("user", "ajaxGetContactList");
	$.get(link, function(b) {
		if (!b) {
			return false
		}
		$inputgroup = $(a).closest(".input-group");
		$inputgroup.find(".input-group-btn").remove();
		$inputgroup.append(b);
		$inputgroup.find("select:last").chosen().fixInputGroup()
	})
}
function addItem(d, b) {
	ItemList = document.getElementById(d);
	Target = document.getElementById(b);
	for (var a = 0; a < ItemList.length; a++) {
		var c = ItemList.options[a];
		if (c.selected) {
			flag = true;
			for (var f = 0; f < Target.length; f++) {
				var e = Target.options[f];
				if (e.value == c.value) {
					flag = false
				}
			}
			if (flag) {
				Target.options[Target.options.length] = new Option(c.text, c.value, 0, 0)
			}
		}
	}
}
function delItem(c) {
	ItemList = document.getElementById(c);
	for (var a = ItemList.length - 1; a >= 0; a--) {
		var b = ItemList.options[a];
		if (b.selected) {
			ItemList.options[a] = null
		}
	}
}
function upItem(c) {
	ItemList = document.getElementById(c);
	for (var a = 1; a < ItemList.length; a++) {
		var b = ItemList.options[a];
		if (b.selected) {
			tmpUpValue = ItemList.options[a - 1].value;
			tmpUpText = ItemList.options[a - 1].text;
			ItemList.options[a - 1].value = b.value;
			ItemList.options[a - 1].text = b.text;
			ItemList.options[a].value = tmpUpValue;
			ItemList.options[a].text = tmpUpText;
			ItemList.options[a - 1].selected = true;
			ItemList.options[a].selected = false;
			break
		}
	}
}
function downItem(c) {
	ItemList = document.getElementById(c);
	for (var a = 0; a < ItemList.length; a++) {
		var b = ItemList.options[a];
		if (b.selected) {
			tmpUpValue = ItemList.options[a + 1].value;
			tmpUpText = ItemList.options[a + 1].text;
			ItemList.options[a + 1].value = b.value;
			ItemList.options[a + 1].text = b.text;
			ItemList.options[a].value = tmpUpValue;
			ItemList.options[a].text = tmpUpText;
			ItemList.options[a + 1].selected = true;
			ItemList.options[a].selected = false;
			break
		}
	}
}
function selectItem(c) {
	ItemList = document.getElementById(c);
	for (var a = ItemList.length - 1; a >= 0; a--) {
		var b = ItemList.options[a];
		b.selected = true
	}
}
function ajaxDelete(a, c, b) {
	if (confirm(b)) {
		$.ajax({
			type: "GET",
			url: a,
			dataType: "json",
			success: function(d) {
				if (d.result == "success") {
					$.get(document.location.href, function(e) {
						if (!($(e).find("#" + c).length)) {
							location.reload()
						}
						$("#" + c).html($(e).find("#" + c).html());
						if (typeof sortTable == "function") {
							sortTable()
						}
						$("#" + c).find("[data-toggle=modal], a.iframe").modalTrigger();
						if ($("#" + c).find("table.datatable").length) {
							$("#" + c).find("table.datatable").datatable()
						}
					})
				}
			}
		})
	}
}
function isNum(b) {
	if (b != null) {
		var c, a;
		a = /\d*/i;
		c = b.match(a);
		return (c == b) ? true : false
	}
	return false
}
function startCron(a) {
	if (typeof(a) == "undefined") {
		a = 0
	}
	$.ajax({
		type: "GET",
		timeout: 100,
		url: createLink("cron", "ajaxExec", "restart=" + a)
	})
}
function computePasswordStrength(b) {
	if (b.length == 0) {
		return 0
	}
	var h = 0;
	var e = b.length;
	var c = "";
	var a = new Array();
	for (i = 0; i < e; i++) {
		letter = b.charAt(i);
		var d = letter.charCodeAt();
		if (d >= 48 && d <= 57) {
			a[2] = 2
		} else {
			if ((d >= 65 && d <= 90)) {
				a[1] = 2
			} else {
				if (d >= 97 && d <= 122) {
					a[0] = 1
				} else {
					a[3] = 3
				}
			}
		}
		if (c.indexOf(letter) == -1) {
			c += letter
		}
	}
	if (c.length > 4) {
		h += c.length - 4
	}
	var g = 0;
	var f = 0;
	for (i in a) {
		f += 1;
		g += a[i]
	}
	h += g + (2 * (f - 1));
	if (e < 6 && h >= 10) {
		h = 9
	}
	h = h > 29 ? 29 : h;
	h = Math.floor(h / 10);
	return h
}
function checkOnlybodyPage() {
	if (self == parent) {
		href = location.href.replace("?onlybody=yes", "");
		location.href = href.replace("&onlybody=yes", "")
	}
}
function fixedTheadOfList(d) {
	if ($(d).size() == 0) {
		return false
	}
	if ($(d).css("display") == "none") {
		return false
	}
	if ($(d).find("thead").size() == 0) {
		return false
	}
	e();
	$(window).scroll(g);
	$(".side-handle").click(function() {
		setTimeout(e, 300)
	});
	var b, f, c, a;

	function g() {
		f = $(d).find("thead").offset().top;
		a = $(d).parent().find(".fixedTheadOfList");
		if (a.size() <= 0 && f < $(window).scrollTop()) {
			b = $(d).width();
			c = "<table class='fixedTheadOfList'><thead>" + $(d).find("thead").html() + "</thead></table>";
			$(d).before(c);
			$(".fixedTheadOfList").addClass($(d).attr("class")).width(b)
		}
		if (a.size() > 0 && f >= $(window).scrollTop()) {
			a.remove()
		}
	}
	function e() {
		a = $(d).parent().find(".fixedTheadOfList");
		if (a.size() > 0) {
			a.remove()
		}
		g()
	}
}
function applyCssStyle(e, b) {
	b = b || "default";
	var d = "applyStyle-" + b;
	var c = $("style#" + d);
	if (!c.length) {
		c = $('<style id="' + d + '">').appendTo("body")
	}
	var a = c.get(0);
	if (a.styleSheet) {
		a.styleSheet.cssText = e
	} else {
		a.innerHTML = e
	}
}
function showBrowserNotice() {
	userAgent = navigator.userAgent.toLowerCase();
	$browser = new Object();
	$browser.msie = /msie/.test(userAgent);
	$browser.chrome = /chrome/.test(userAgent);
	var a = false;
	if (navigator.userAgent.indexOf("MetaSr") >= 0) {
		a = true
	} else {
		if (navigator.userAgent.indexOf("LBBROWSER") >= 0) {
			a = true
		} else {
			if (navigator.userAgent.indexOf("QQBrowser") >= 0) {
				a = true
			} else {
				if (navigator.userAgent.indexOf("TheWorld") >= 0) {
					a = true
				} else {
					if (navigator.userAgent.indexOf("BIDUBrowser") >= 0) {
						a = true
					} else {
						if (navigator.userAgent.indexOf("Maxthon") >= 0) {
							a = true
						}
					}
				}
			}
		}
	}
	if (a) {
		$("body").prepend('<div class="alert alert-info alert-dismissable" style="margin:0px;"><button type=button" onclick="ajaxIgnoreBrowser()" class="close" data-dismiss="alert" aria-hidden="true"><i class="icon-remove"></i></button><p>' + browserNotice + "</p></div>")
	}
}
function removeCookieByKey(a) {
	$.cookie(a, "", {
		expires: config.cookieLife,
		path: config.webRoot
	});
	location.href = location.href
}
function setHomepage(a, b) {
	$.get(createLink("custom", "ajaxSetHomepage", "module=" + a + "&page=" + b), function() {
		location.reload(true)
	})
}
function checkTutorial() {
	if (config.currentModule != "tutorial" && window.TUTORIAL && (!frameElement || frameElement.tagName != "IFRAME")) {
		if (confirm(window.TUTORIAL.tip)) {
			$.getJSON(createLink("tutorial", "ajaxQuit"), function() {
				window.location.reload()
			}).error(function() {
				alert(lang.timeout)
			})
		}
	}
}
function removeDitto() {
	$firstTr = $(".table-form").find("tbody tr:first");
	$firstTr.find("td select").each(function() {
		$(this).find("option[value='ditto']").remove();
		$(this).trigger("chosen:updated")
	})
}
function revertModuleCookie() {
	if ($('#mainmenu .nav li[data-id="project"]').hasClass("active")) {
		$('#modulemenu .nav li[data-id="task"] a').click(function() {
			$.cookie("moduleBrowseParam", 0, {
				expires: config.cookieLife,
				path: config.webRoot
			})
		})
	}
	if ($('#mainmenu .nav li[data-id="product"]').hasClass("active")) {
		$('#modulemenu .nav li[data-id="story"] a').click(function() {
			$.cookie("storyModule", 0, {
				expires: config.cookieLife,
				path: config.webRoot
			})
		})
	}
	if ($('#mainmenu .nav li[data-id="qa"]').hasClass("active")) {
		$('#modulemenu .nav li[data-id="bug"] a').click(function() {
			$.cookie("bugModule", 0, {
				expires: config.cookieLife,
				path: config.webRoot
			})
		});
		$('#modulemenu .nav li[data-id="testcase"] a').click(function() {
			$.cookie("caseModule", 0, {
				expires: config.cookieLife,
				path: config.webRoot
			})
		})
	}
}
function inputFocusJump(d) {
	var f = $("input").is(":focus");
	if (f) {
		var e = $("input:focus").attr("name").replace(/\[\d]/g, "");
		var g = $(":input[name^=" + e + "]:text:not(:disabled):not([name*='%'])");
		var c = g.length;
		var b = parseInt($("input:focus").attr("name").replace(/[^0-9]/g, ""));
		var a = d == "down" ? b + 1 : b - 1;
		if (a < c && a >= 0) {
			g[a].focus()
		}
	}
}
function selectFocusJump(e) {
	var g = $("select").is(":focus");
	if (g) {
		var f = $("select:focus").attr("name").replace(/\[\d]/g, "");
		var d = $("select[name^=" + f + "]:not([name*='%'])");
		var c = d.length;
		var b = parseInt($("select:focus").attr("name").replace(/[^0-9]/g, ""));
		var a = e == "down" ? b + 1 : b - 1;
		if (a < c && a >= 0) {
			d[a].focus()
		}
	}
}
function adjustNoticePosition() {
	var a = 25;
	$("#noticeBox").find(".alert").each(function() {
		$(this).css("bottom", a + "px");
		a += $(this).outerHeight(true) - 10
	})
}
function notifyMessage(a) {
	if (window.Notification) {
		if (Notification.permission == "granted") {
			new Notification("", {
				body: a
			})
		} else {
			if (Notification.permission != "denied") {
				Notification.requestPermission(function(b) {
					new Notification("", {
						body: a
					})
				})
			}
		}
	}
}
function getFingerprint() {
	if (typeof(Fingerprint) == "function") {
		return new Fingerprint().get()
	}
	fingerprint = "";
	$.each(navigator, function(a, b) {
		if (typeof(b) == "string") {
			fingerprint += b.length
		}
	});
	return fingerprint
}
needPing = true;
$(document).ready(function() {
	if (needPing) {
		setTimeout("setPing()", 1000 * 60 * 10)
	}
	checkTutorial();
	revertModuleCookie();
	$(document).on("click", "#helpMenuItem .close-help-tab", function() {
		$("#helpMenuItem").prev().remove();
		$("#helpMenuItem").remove()
	})
});
/*!
 * jQuery Cookie Plugin v1.4.1
 * https://github.com/carhartl/jquery-cookie
 * Copyright 2006, 2014 Klaus Hartl
 * Released under the MIT license
 */

!function(a) {
	"function" == typeof define && define.amd ? define(["jquery"], a) : a("object" == typeof exports ? require("jquery") : jQuery)
}(function(d) {
	function h(a) {
		return g.raw ? a : encodeURIComponent(a)
	}
	function c(a) {
		return g.raw ? a : decodeURIComponent(a)
	}
	function k(a) {
		return h(g.json ? JSON.stringify(a) : String(a))
	}
	function j(a) {
		0 === a.indexOf('"') && (a = a.slice(1, -1).replace(/\\"/g, '"').replace(/\\\\/g, "\\"));
		try {
			return a = decodeURIComponent(a.replace(b, " ")), g.json ? JSON.parse(a) : a
		} catch (l) {}
	}
	function f(l, a) {
		var m = g.raw ? l : j(l);
		return d.isFunction(a) ? a(m) : m
	}
	var b = /\+/g,
		g = d.cookie = function(n, A, r) {
			if (void 0 !== A && !d.isFunction(A)) {
				if (r = d.extend({}, g.defaults, r), "number" == typeof r.expires) {
					var s = r.expires,
						z = r.expires = new Date;
					z.setTime(+z + 86400000 * s)
				}
				return document.cookie = [h(n), "=", k(A), r.expires ? "; expires=" + r.expires.toUTCString() : "", r.path ? "; path=" + r.path : "", r.domain ? "; domain=" + r.domain : "", r.secure ? "; secure" : ""].join("")
			}
			for (var y = n ? void 0 : {}, C = document.cookie ? document.cookie.split("; ") : [], e = 0, x = C.length; e < x; e++) {
				var w = C[e].split("="),
					q = c(w.shift()),
					B = w.join("=");
				if (n && n === q) {
					y = f(B, A);
					break
				}
				n || void 0 === (B = f(B)) || (y[q] = B)
			}
			return y
		};
	g.defaults = {}, d.removeCookie = function(l, a) {
		return void 0 !== d.cookie(l) && (d.cookie(l, "", d.extend({}, a, {
			expires: -1
		})), !d.cookie(l))
	}
}), function(d, g) {
	var c, j, h = "localStorage",
		f = "page_" + d.location.pathname + d.location.search,
		b = function() {
			this.slience = !0;
			try {
				h in d && d[h] && d[h].setItem && (this.enable = !0, c = d[h])
			} catch (e) {}
			this.enable || (j = {}, c = {
				getLength: function() {
					var a = 0;
					return g.each(j, function() {
						a++
					}), a
				},
				key: function(k) {
					var a, l = 0;
					return g.each(j, function(m) {
						return l === k ? (a = m, !1) : void l++
					}), a
				},
				removeItem: function(a) {
					delete j[a]
				},
				getItem: function(a) {
					return j[a]
				},
				setItem: function(a, k) {
					j[a] = k
				},
				clear: function() {
					j = {}
				}
			}), this.storage = c, this.page = this.get(f, {})
		};
	b.prototype.pageSave = function() {
		if (g.isEmptyObject(this.page)) {
			this.remove(f)
		} else {
			var e, a = [];
			for (e in this.page) {
				var k = this.page[e];
				null === k && a.push(e)
			}
			for (e = a.length - 1; e >= 0; e--) {
				delete this.page[a[e]]
			}
			this.set(f, this.page)
		}
	}, b.prototype.pageRemove = function(a) {
		"undefined" != typeof this.page[a] && (this.page[a] = null, this.pageSave())
	}, b.prototype.pageClear = function() {
		this.page = {}, this.pageSave()
	}, b.prototype.pageGet = function(k, l) {
		var a = this.page[k];
		return void 0 === l || null !== a && void 0 !== a ? a : l
	}, b.prototype.pageSet = function(e, a) {
		g.isPlainObject(e) ? g.extend(!0, this.page, e) : this.page[this.serialize(e)] = a, this.pageSave()
	}, b.prototype.check = function() {
		if (!this.enable && !this.slience) {
			throw new Error("Browser not support localStorage or enable status been set true.")
		}
		return this.enable
	}, b.prototype.length = function() {
		return this.check() ? c.getLength ? c.getLength() : c.length : 0
	}, b.prototype.removeItem = function(a) {
		return c.removeItem(a), this
	}, b.prototype.remove = function(a) {
		return this.removeItem(a)
	}, b.prototype.getItem = function(a) {
		return c.getItem(a)
	}, b.prototype.get = function(k, l) {
		var a = this.deserialize(this.getItem(k));
		return "undefined" != typeof a && null !== a || "undefined" == typeof l ? a : l
	}, b.prototype.key = function(a) {
		return c.key(a)
	}, b.prototype.setItem = function(a, k) {
		return c.setItem(a, k), this
	}, b.prototype.set = function(a, k) {
		return void 0 === k ? this.remove(a) : (this.setItem(a, this.serialize(k)), this)
	}, b.prototype.clear = function() {
		return c.clear(), this
	}, b.prototype.forEach = function(a) {
		for (var k = this.length(), m = k - 1; m >= 0; m--) {
			var l = c.key(m);
			a(l, this.get(l))
		}
		return this
	}, b.prototype.getAll = function() {
		var a = {};
		return this.forEach(function(l, k) {
			a[l] = k
		}), a
	}, b.prototype.serialize = function(a) {
		return "string" == typeof a ? a : JSON.stringify(a)
	}, b.prototype.deserialize = function(a) {
		if ("string" == typeof a) {
			try {
				return JSON.parse(a)
			} catch (k) {
				return a || void 0
			}
		}
	}, g.zui({
		store: new b
	})
}(window, jQuery);