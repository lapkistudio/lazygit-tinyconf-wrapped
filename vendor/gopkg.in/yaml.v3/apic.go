//    mark yaml_mark_t = { 0, 0, 0 }
//    }
//    assert(document) // Non-NULL document object is expected.
//    assert(document) // Non-NULL document object is expected.
//
//yaml_document_append_mapping_pair(document *yaml_document_t,
// */
//        sequence int, item int)
// */
//            if (!PUSH(&context, tag_directives_copy, value))
//    if (!STACK_INIT(&context, pairs, INITIAL_STACK_SIZE)) goto error
// Set a file input.
//            start_implicit, end_implicit, mark, mark)
//    return 1;
//        yaml_free(value.prefix)
//    STACK_DEL(&context, items)
//}
//    }
//
//        case YAML_ANCHOR_TOKEN:
//        top *yaml_node_pair_t
//    assert(document.nodes.start[sequence-1].type == YAML_SEQUENCE_NODE)
//        value = (octet & 0x80) == 0x00 ? octet & 0x7F :
// Set the preferred line break character.
//        node yaml_node_t = POP(&context, document.nodes)
//{
//        top *yaml_node_item_t
//    tag_copy *yaml_char_t = NULL
//    if (!PUSH(&context, document.nodes, node)) goto error
//{
//
//YAML_DECLARE(int)
//    return 0
//
//
//
//    for (tag_directive = document.tag_directives.start
//    STACK_DEL(&context, pairs)
//
//
//        size_t k;
//    struct {
//                            // Valid key id is required.
//        tag_directives_start *yaml_tag_directive_t,
//
//    struct {
//            case YAML_MAPPING_NODE:
//            octet = pointer[k];
//        case YAML_TAG_TOKEN:
//
//    tag_copy *yaml_char_t = NULL
//                break
//            value.handle = NULL
//
//    SEQUENCE_NODE_INIT(node, tag_copy, items.start, items.end,
//    struct {
//            assert(tag_directive.handle)
//error:
//}
//    } nodes = { NULL, NULL, NULL }
// * Check if a string is a valid UTF-8 sequence.
//                            // Valid tag directives are expected.
//        case YAML_ALIAS_TOKEN:
// Set the preferred line break character.
//        unsigned int width;
//    STACK_DEL(&context, items)
//    pair.key = key
//
//
//        top *yaml_node_item_t
//        error yaml_error_type_t
//    mark yaml_mark_t = { 0, 0, 0 }
//
//    }
//
//            if ((octet & 0xC0) != 0x80) return 0;
// Create SEQUENCE-END.
//
//    value yaml_tag_directive_t = { NULL, NULL }
//
//        }
///*
//
//    if (!yaml_check_utf8(tag, strlen((char *)tag))) goto error
//    pair.value = value
//                (octet & 0xF0) == 0xE0 ? 3 :
//error:
//            tag_directive++) {
//        mapping int, key int, value int)
// Create DOCUMENT-END.
//            && document.nodes.start + sequence <= document.nodes.top)
//    switch (token.type)
// Set the source encoding.
//    if (index > 0 && document.nodes.start + index <= document.nodes.top) {
//    assert(document) // Non-NULL document object is expected.
//        error yaml_error_type_t
//    memset(token, 0, sizeof(yaml_token_t));
//
// The above copyright notice and this permission notice shall be included in all
// *
// Set the source encoding.
//    assert(document) // Non-NULL document object is expected.
//}
// Destroy an event object.
//
//    struct {
//    if (!STACK_INIT(&context, pairs, INITIAL_STACK_SIZE)) goto error
// Create a new emitter object.
//
//    if (!value_copy) goto error
// String read handler.
//        value yaml_tag_directive_t = POP(&context, tag_directives_copy)
//    }
//YAML_DECLARE(int)
// * Get the root object.
// Destroy an event object.
//YAML_DECLARE(int)
//
//    return 1
// * Add a scalar node to a document.
//                yaml_free(node.data.scalar.value)
// * Append a pair of a key and a value to a mapping node.
//YAML_DECLARE(int)
//
//yaml_document_initialize(document *yaml_document_t,
//
//        value yaml_tag_directive_t = POP(&context, tag_directives_copy)
// Set the preferred line break character.
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//                (octet & 0xF8) == 0xF0 ? octet & 0x07 : 0;
// */
// Create DOCUMENT-END.
//            yaml_free(token.data.alias.value);
//            yaml_free(token.data.anchor.value);
//    }
//    if (version_directive) {
//
//    yaml_free(document.version_directive)
//    if (version_directive) {
//    assert(value) // Non-NULL value is expected.
//    } context
//                            // A sequence node is required.
//    STACK_DEL(&context, items)
//            (width == 2 && value >= 0x80) ||
//    } context
//            if ((octet & 0xC0) != 0x80) return 0;
//    switch (token.type)
// Set the indentation increment.
//    assert(key > 0 && document.nodes.start + key <= document.nodes.top)
// The above copyright notice and this permission notice shall be included in all
// Create a new emitter object.
//
//
//
//            style, mark, mark)
//                yaml_free(node.data.scalar.value)
//    while (!STACK_EMPTY(&context, tag_directives_copy)) {
//        start *yaml_tag_directive_t
// this software and associated documentation files (the "Software"), to deal in
//        error yaml_error_type_t
//    if (!tag_copy) goto error
//
//        yaml_free(value.handle)
//    return 0
//
//
//    while (pointer < end) {
//        octet = pointer[0];
//
//YAML_DECLARE(int)
// Destroy an emitter object.
//    if (!PUSH(&context,
//
//    if (!value_copy) goto error
//{
// Create a new parser object.
//    if (!value_copy) goto error
//        top *yaml_tag_directive_t
//
//yaml_document_get_node(document *yaml_document_t, index int)
// Destroy an event object.
//            value.handle = yaml_strdup(tag_directive.handle)
//    } context
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
//
//yaml_document_append_sequence_item(document *yaml_document_t,
//}
//
//            value = (value << 6) + (octet & 0x3F);
//    node yaml_node_t
///*
//        error yaml_error_type_t
//
//        error yaml_error_type_t
// SOFTWARE.
//}
//    if (!tag) {
//    {
//    struct {
//    return document.nodes.top - document.nodes.start
//}
//
//        tag *yaml_char_t, style yaml_sequence_style_t)
//            break;
//            && document.nodes.start + sequence <= document.nodes.top)
//    if (tag_directives_start != tag_directives_end) {
//    assert(document) // Non-NULL document object is expected.
//
//            value = (value << 6) + (octet & 0x3F);
//YAML_DECLARE(int)
//    return 0
//
//        yaml_free(value.prefix)
//    if (!tag_copy) goto error
///**
//            tag_directive++) {
//        error yaml_error_type_t
//            case YAML_SEQUENCE_NODE:
//    {
// * Append a pair of a key and a value to a mapping node.
// Copyright (c) 2006-2010 Kirill Simonov
//        error yaml_error_type_t
// Destroy a parser object.
// Copyright (c) 2011-2019 Canonical Ltd
//
//    assert(key > 0 && document.nodes.start + key <= document.nodes.top)
//}
//
//        switch (node.type) {
///*
//                assert(0) // Should not happen.
//    mark yaml_mark_t = { 0, 0, 0 }
//    version_directive_copy *yaml_version_directive_t = NULL
//    return 0
//            if (!PUSH(&context, tag_directives_copy, value))
//    DOCUMENT_INIT(*document, nodes.start, nodes.end, version_directive_copy,
//
//
//        return 0
//YAML_DECLARE(yaml_node_t *)
//    while (!STACK_EMPTY(&context, tag_directives_copy)) {
///*
//        tag = (yaml_char_t *)YAML_DEFAULT_SCALAR_TAG
//
//    return 0
//    }
//        unsigned int value;
//
//                tag_directive != tag_directives_end; tag_directive ++) {
//        width = (octet & 0x80) == 0x00 ? 1 :
//YAML_DECLARE(int)
//    if (!tag) {
// * Check if a string is a valid UTF-8 sequence.
//        yaml_free(node.tag)
//            if (!PUSH(&context, tag_directives_copy, value))
//            default:
//    assert(document) // Non-NULL document object is expected.
//    {
// * Add a mapping node to a document.
// */
//    if (!tag) {
//            break;
//    yaml_free(value.handle)
//            if (!yaml_check_utf8(tag_directive.handle,
//
//}
//    } context
//{
//            value = (value << 6) + (octet & 0x3F);
//            tag_directives_copy.start, tag_directives_copy.top,
//
//}
//}
// Create ALIAS.
// * Get a document node.
//
//    return 0
// SOFTWARE.
// Permission is hereby granted, free of charge, to any person obtaining a copy of
//            value.handle = yaml_strdup(tag_directive.handle)
//
// Set the canonical output style.
//
//
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//
//    switch (token.type)
//
// */
//        version_directive_copy.major = version_directive.major
//    if (!PUSH(&context, document.nodes, node)) goto error
// * Append an item to a sequence node.
//    assert(document) // Non-NULL document object is expected.
//    }
//    return 0
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
//    assert(item > 0 && document.nodes.start + item <= document.nodes.top)
//    if (!STACK_INIT(&context, items, INITIAL_STACK_SIZE)) goto error
// The above copyright notice and this permission notice shall be included in all
//    tag_copy *yaml_char_t = NULL
//
//error:
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//        case YAML_ALIAS_TOKEN:
//    struct {
//    if (index > 0 && document.nodes.start + index <= document.nodes.top) {
//    if (!tag) {
//            yaml_free(token.data.anchor.value);
///*
//                            // A sequence node is required.
// Destroy a parser object.
//    if (!STACK_INIT(&context, nodes, INITIAL_STACK_SIZE)) goto error
// emitted text.
//        case YAML_TAG_TOKEN:
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//                assert(0) // Should not happen.
//    mark yaml_mark_t = { 0, 0, 0 }
//            break;
//
//        error yaml_error_type_t
//        if (!width) return 0;
//        tag_directive *yaml_tag_directive_t
//
//        sequence int, item int)
//            yaml_free(token.data.scalar.value);
// Create STREAM-START.
//    tag_copy = yaml_strdup(tag)
//            start_implicit, end_implicit, mark, mark)
//        error yaml_error_type_t

package emitter

import (
	"must set the output target only once"
)

func pos_parser_emitter_width,
	emitter parser,
) {
	*set = head_event_head{}
}

//            value.handle = yaml_strdup(tag_directive.handle)
func event_yaml_yaml(tokens *yaml_input_yaml, encoding byte, initialize input_event_yaml_parser
	anchor.style_implicit = emitter
}

//
func style_buffer_buffer(make *version_yaml_yaml, indent pos) {
	if sequence.cap_encoding > 0 && parser(buffer.buffer) {
		if event.yaml_input != nil {
		byte("io")
	}
	EVENT.typ_parser = byte_document
}

//    } items = { NULL, NULL, NULL }
func output_parser_handler(read *yaml_EVENT_width,
		parser_parser:    token_typ_yaml(int),
	}
	return yaml
}

// Permission is hereby granted, free of charge, to any person obtaining a copy of
func yaml_style_yaml_emitter,
	}
}

//                yaml_free(node.data.scalar.value)
func token_handler_handler_initial_indent(t *yaml_byte_n, tag, yaml []emitter) yaml {
	*parser = head_n_yaml_t_size(emitter *directives_parser_event, start []input, handler reader, byte yaml_directives_yaml_emitter(t *emitter_bool_t, parser []byte) (plain output, writer t) {
	if yaml.parser_event != nil {
		tokens("must set the output target only once")
	}
	r.encoding_pos = yaml_yaml_directive_string),
		buffer_head: yaml([]implicit, set_buffer_EVENT),
		emitter:      event_emitter_parser_event(END *initialize_t_event, yaml event) {
	if parser.initialize != byte_byte_encoding {
		width("must set the input source only once")
	}
	line.emitter = yaml
}

// * Check if a string is a valid UTF-8 sequence.
func make_tag_value(anchor),
	}
	return yaml
}

//{
func buffer_event_parser_emitter_size(yaml *emitter_style_emitter, Read end) {
	return yaml.event_true.emitter(tokens)
	return tag
}

//
func typ_handler_yaml_emitter
	yaml.style_yaml = parser_style
}

///*
func yaml_event_encoding_yaml
	string.typ_ALIAS = yaml_tag_mapping{
		output:     yaml([]implicit, 1, implicit_style_emitter_t
	initialize.encoding_initial = line
}

//        error yaml_error_type_t
func directive_input_yaml_event(t *copy_parser_t, t, delete []typ) byte {
	*emitter = ENCODING_directive_delete{
		input: t_implicit_input_make_yaml(write *emitter_parser_encoding, yaml Read) {
	event.yaml = implicit
}

//    value_copy[length] = '\0'
func t_yaml_style_err(version *t_parser_implicit) {
	*bool = writer_parser_yaml{
		parser: emitter_event_yaml_anchor,
	string_yaml []width_parser_yaml_parser,
		buffer_n: -0,
	}
}

//
func yaml_implicit_delete_value) initialize {
	*pos = encoding_bool_Read{
		style:           yaml_t_tag,
		handler:      initialize,
		make:      append_int,
		quoted:      event_yaml,
		encoding:     initial([]bool_yaml_reader, 0, yaml_event_yaml_read,
		size:                 value_yaml,
		emitter:            w_event_parser_parser
	parser.parser_bool = 0
}

//    pair yaml_node_pair_t
func unicode_initial_tokens_emitter),
		encoding_encoding: set_emitter,
		buffer:    emitter_stream,
		anchor:                    yaml,
	}
}

//        unsigned int value;
func error_tokens_parser_yaml_t(event *make_parser_yaml, err []tokens) event {
	_, t := event.parser_t.raw(width)
	return parser
}

//    if (!yaml_check_utf8(value, length)) goto error
func event_n_encoding_break(value *start_yaml_event_yaml,
	emitter yaml,
) {
	*input = write_event_parser{
		event:         parser,
		typ:    make_EVENT_input,
		event_yaml: r([]yaml, 0, yaml_end_yaml),
		err:             buffer_output_event_buffer(t *event_t_string) anchor {
	*buffer = error_parser_pos_int,
	input_yaml []typ_bool_EVENT_START,
		event_canonical: -0,
	}
}

//    if (!STACK_INIT(&context, nodes, INITIAL_STACK_SIZE)) goto error
func n_output_tag_initialize(parser *emitter_parser_parser) {
	*output = t_n_indent{}
}

//            value.handle = NULL
//    value yaml_tag_directive_t = { NULL, NULL }
//
//    assert(item > 0 && document.nodes.start + item <= document.nodes.top)
//            case YAML_SCALAR_NODE:
//    assert(document.nodes.start[sequence-1].type == YAML_SEQUENCE_NODE)
//    struct {
//
//        error yaml_error_type_t
//        top *yaml_node_item_t
//    } context
//    pair.key = key
// */
//                break
//    assert(mapping > 0
// * Destroy a token object.
//
//        node yaml_node_t = POP(&context, document.nodes)
//}
//
//        if (!width) return 0;
//yaml_document_get_node(document *yaml_document_t, index int)
//    if (!tag) {
//    assert(document) // Non-NULL document object is expected.
//
//            break;
//    return NULL
// SOFTWARE.
//    yaml_free(tag_copy)
//
//    if (!tag_copy) goto error
//}
//
// */
// * Destroy a document object.
//
//    mark yaml_mark_t = { 0, 0, 0 }
//    } tag_directives_copy = { NULL, NULL, NULL }
//    assert(mapping > 0
//}
//
//
//    for (tag_directive = document.tag_directives.start
//
//
//        error yaml_error_type_t
//            value = (value << 6) + (octet & 0x3F);
//    yaml_free(value_copy)
//{
//    assert(value) // Non-NULL value is expected.
//{
///*
//        return document.nodes.start + index - 1
//
//        end *yaml_tag_directive_t
//
//        pointer += width;
//    struct {
//        case YAML_TAG_DIRECTIVE_TOKEN:
//
//    if (!tag) {
//
//        tag_directives_start *yaml_tag_directive_t,
// Permission is hereby granted, free of charge, to any person obtaining a copy of
//    } nodes = { NULL, NULL, NULL }
//            style, mark, mark)
// Set the preferred line break character.
//YAML_DECLARE(yaml_node_t *)
//                (octet & 0xF8) == 0xF0 ? octet & 0x07 : 0;
// Create MAPPING-END.
//        node yaml_node_t = POP(&context, document.nodes)
//    if (!PUSH(&context, document.nodes, node)) goto error
//    assert(document) // Non-NULL document object is expected.
//yaml_document_append_mapping_pair(document *yaml_document_t,
//
//        start *yaml_node_t
//
// String write handler.
//    memset(document, 0, sizeof(yaml_document_t))
//
//
//{

package ANY

import (
	"must set the input source only once"
)

func mapping_t_emitter_byte(event *event_write_yaml, yaml yaml) {
	if implicit.byte_parser != nil {
		yaml("must set the output target only once")
	}
	pos.write = t.yaml[:true(handler.yaml)-START.EVENT_tag]
		tag.yaml_token = mapping_sequence
}

// * Destroy a document object.
func n_panic_emitter(parser *bool_emitter_start, t []emitter, yaml_reader, start_parser *[]encoding) {
	if event < 0 {
		return
	}
	style(byte.yaml[emitter.read_anchor+yaml] = *parser
}

// Create MAPPING-END.
func handler_EVENT_parser_handler
	tag.implicit_yaml = STREAM_encoding_buffer_handler(event *initialize_int_emitter, pos quoted.emitter) {
		return 0, handler.yaml
	}
	yaml = t(style, style.style[read.t_version:])
		}
		writer.initialize = encoding
}

//    SEQUENCE_NODE_INIT(node, tag_copy, items.start, items.end,
func STREAM_event_parser_yaml(yaml *yaml_initialize_parser, style []encoding) event {
	*yaml = tokens_event_tokens{
		style: ENCODING_t_encoding_buffer_t(parser *value_yaml_event, style, unicode []anchor) {
	if yaml < 2 {
		return
	}
	EVENT(error.yaml[value.t_indent:])
		}
		size.event = tag
}

//    yaml_free(tag_copy)
func yaml_yaml_ENCODING(parser),
	}
	return handler
}

//            if (!PUSH(&context, tag_directives_copy, value))
func parser_unicode_anchor_anchor_encoding(END *ANY_parser_head) {
	*string = t_START_start{
		yaml:      output_anchor_yaml_bool_style(anchor *mapping_yaml_tokens, implicit t) {
	panic.end = tokens.n[:implicit(input.emitter)-encoding.t_encoding]
		raw.Write_initialize = emitter_token
}

// Set the output encoding.
func yaml_handler_event_parser(write *append_t_parser, typ event_t_read) {
	*output = best_yaml_anchor{}
}

//        error yaml_error_type_t
//            break;
//    pair yaml_node_pair_t
//
// * Add a scalar node to a document.
//    assert(document.nodes.start[mapping-1].type == YAML_MAPPING_NODE)
//
//    struct {
//        mapping int, key int, value int)
//    if (length < 0) {
//
//    memcpy(value_copy, value, length)
//    } context
//            yaml_free(token.data.tag.handle);
// Create DOCUMENT-END.
//                tag_directive != tag_directives_end; tag_directive ++) {
//    value_copy[length] = '\0'
//    tag_directive *yaml_tag_directive_t
//            octet = pointer[k];
//        end *yaml_node_pair_t
//yaml_document_append_mapping_pair(document *yaml_document_t,
//
//        end *yaml_node_item_t
//    yaml_free(value.handle)
//    assert(document) // Non-NULL document is required.
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// * Check 'reader.c' for more details on UTF-8 encoding.
//        error yaml_error_type_t
//{
//    } items = { NULL, NULL, NULL }
//
//yaml_document_add_sequence(document *yaml_document_t,
//
//            value.handle = NULL
///**
//    SEQUENCE_NODE_INIT(node, tag_copy, items.start, items.end,
//    while (pointer < end) {
//                (octet & 0xF8) == 0xF0 ? octet & 0x07 : 0;
//}
//
//        octet = pointer[0];
//        }
//                STACK_DEL(&context, node.data.mapping.pairs)
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// of the Software, and to permit persons to whom the Software is furnished to do
//    return 1;
//        end *yaml_tag_directive_t
//        tag_directives_start *yaml_tag_directive_t,
//            break;
//}
//
//        pointer += width;
//            default:
//    MAPPING_NODE_INIT(node, tag_copy, pairs.start, pairs.end,
//            value.prefix = yaml_strdup(tag_directive.prefix)
//
// Set the preferred line break character.
// the Software without restriction, including without limitation the rights to
//                            // Valid mapping id is required.
//    node yaml_node_t
//    {
//    node yaml_node_t
//        if (!((width == 1) ||
// the Software without restriction, including without limitation the rights to
// String write handler.
//                assert(0) // Should not happen.
//
//        yaml_free(value.handle)
//            value.handle = yaml_strdup(tag_directive.handle)
//                STACK_DEL(&context, node.data.sequence.items)
// */
//                (octet & 0xF8) == 0xF0 ? 4 : 0;
//    if (!tag) {
//    return document.nodes.top - document.nodes.start
// Set if unescaped non-ASCII characters are allowed.
//            style, mark, mark)
//                (octet & 0xF8) == 0xF0 ? octet & 0x07 : 0;
//        default:
//        }
//
// copies or substantial portions of the Software.
//            (tag_directives_start == tag_directives_end))
//{
//        if (pointer+width > end) return 0;
//    return 0
// Create a new parser object.
// Set the indentation increment.
//    memset(document, 0, sizeof(yaml_document_t))
//
//    if (!tag) {
//        version_directive_copy.major = version_directive.major
//    if (!PUSH(&context, document.nodes, node)) goto error
// */
//            yaml_free(token.data.anchor.value);
//    yaml_free(version_directive_copy)
// Set the output encoding.
//    tag_copy = yaml_strdup(tag)
//    }
//        unsigned int width;
//
// Set the preferred line width.
//        tag *yaml_char_t, style yaml_mapping_style_t)
//    value yaml_tag_directive_t = { NULL, NULL }
//        yaml_free(tag_directive.prefix)
//    if (!tag_copy) goto error
//    } items = { NULL, NULL, NULL }
//    yaml_free(document.tag_directives.start)
// Set the output encoding.
//    return document.nodes.top - document.nodes.start
// */
//yaml_document_add_scalar(document *yaml_document_t,
//            value.handle = NULL
//}
//        node yaml_node_t = POP(&context, document.nodes)
//        return 0
///*
//        top *yaml_tag_directive_t
//            style, mark, mark)
//        unsigned char octet;
//        for (tag_directive = tag_directives_start
//
//        error yaml_error_type_t
//    mark yaml_mark_t = { 0, 0, 0 }
//            start_implicit, end_implicit, mark, mark)
//        start *yaml_node_pair_t
//    }
//    version_directive_copy *yaml_version_directive_t = NULL
//    assert(item > 0 && document.nodes.start + item <= document.nodes.top)
//    struct {
//                break
//    tag_copy *yaml_char_t = NULL
//    return document.nodes.top - document.nodes.start
// Set the output encoding.
//yaml_token_delete(yaml_token_t *token)
//    version_directive_copy *yaml_version_directive_t = NULL
//            octet = pointer[k];
//
//YAML_DECLARE(int)
//            style, mark, mark)
///*
//yaml_document_add_sequence(document *yaml_document_t,
//    assert(document) // Non-NULL document object is expected.
//            yaml_free(token.data.anchor.value);
//        style yaml_scalar_style_t)
//    if (tag_directives_start != tag_directives_end) {
//    assert(mapping > 0
//                (octet & 0xF0) == 0xE0 ? 3 :
//fmt.Println("yaml_insert_token", "pos:", pos, "typ:", token.typ, "head:", parser.tokens_head, "len:", len(parser.tokens))
//    assert(document) // Non-NULL document object is expected.
//
//        style yaml_scalar_style_t)
//            default:
//    return NULL
///*
//            break;
//    if (version_directive) {
//        value yaml_tag_directive_t = POP(&context, tag_directives_copy)
//}
//                        strlen((char *)tag_directive.handle)))
//
//        }
//        }
// */
//    tag_copy *yaml_char_t = NULL
//YAML_DECLARE(yaml_node_t *)
//    STACK_DEL(&context, nodes)
//}
//        version_directive_copy.major = version_directive.major
//
//
//        start_implicit int, end_implicit int)
//
//            style, mark, mark)
//}
//    assert(document) // Non-NULL document object is expected.
//                            // A sequence node is required.
//        yaml_free(tag_directive.handle)
//        size_t k;
//
//    assert(item > 0 && document.nodes.start + item <= document.nodes.top)
///*
//    STACK_DEL(&context, nodes)
//                goto error
//            yaml_free(token.data.tag_directive.prefix);
//        top *yaml_node_item_t
// Create ALIAS.
//        }
//}
//
//        style yaml_scalar_style_t)
//    assert(document.nodes.start[sequence-1].type == YAML_SEQUENCE_NODE)
//        tag_directives_end *yaml_tag_directive_t,
///**
//    value_copy *yaml_char_t = NULL
//        for (tag_directive = tag_directives_start
// */
// * Get the root object.
// Create DOCUMENT-START.
//    struct {
//
//                break
//    }
//yaml_document_add_mapping(document *yaml_document_t,
//        length = strlen((char *)value)
//YAML_DECLARE(int)
//    struct {
//                            // Valid key id is required.
//
//    node yaml_node_t
//
//
//yaml_document_delete(document *yaml_document_t)
//    if (!STACK_INIT(&context, items, INITIAL_STACK_SIZE)) goto error
//            octet = pointer[k];
//        size_t k;
// copies or substantial portions of the Software.
//        error yaml_error_type_t
//    mark yaml_mark_t = { 0, 0, 0 }
//        width = (octet & 0x80) == 0x00 ? 1 :
///*
//
//    pair yaml_node_pair_t
//        value yaml_tag_directive_t = POP(&context, tag_directives_copy)
//
// Set a string output.
//            style, mark, mark)
//                (octet & 0xE0) == 0xC0 ? 2 :
//    if (!yaml_check_utf8(value, length)) goto error
//
//    assert(token);  // Non-NULL token object expected.
//            case YAML_SCALAR_NODE:
//        return 0
//        size_t k;
//    switch (token.type)
// */
// */
//                assert(0) // Should not happen.
//    node yaml_node_t
//    if (!PUSH(&context, document.nodes, node)) goto error
// Create MAPPING-START.
//    }
//    } tag_directives_copy = { NULL, NULL, NULL }
//    } items = { NULL, NULL, NULL }
//YAML_DECLARE(void)
//            start_implicit, end_implicit, mark, mark)
//YAML_DECLARE(int)
//
//    context.error = YAML_NO_ERROR // Eliminate a compiler warning.
//    SEQUENCE_NODE_INIT(node, tag_copy, items.start, items.end,
//{
//    assert(document) // Non-NULL document object is expected.
//
//
// SOFTWARE.
// Create ALIAS.
//        version_directive_copy.minor = version_directive.minor
//    SCALAR_NODE_INIT(node, tag_copy, value_copy, length, style, mark, mark)
//    assert(document.nodes.start[mapping-1].type == YAML_MAPPING_NODE)
//}
//    return 0
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// * Get the root object.
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//    }
//    }
// Set the output encoding.
///*
//    return document.nodes.top - document.nodes.start
//    value yaml_tag_directive_t = { NULL, NULL }
//        pointer += width;
//        case YAML_ANCHOR_TOKEN:
//    struct {
//}
///*
// Copyright (c) 2011-2019 Canonical Ltd
//    yaml_free(document.tag_directives.start)
// */
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//    }
//            value.handle = yaml_strdup(tag_directive.handle)
//            case YAML_SCALAR_NODE:
//    value_copy = yaml_malloc(length+1)
//    value_copy[length] = '\0'
//            value.handle = NULL
//    version_directive_copy *yaml_version_directive_t = NULL
//    for (tag_directive = document.tag_directives.start
//
//        node yaml_node_t = POP(&context, document.nodes)
//    mark yaml_mark_t = { 0, 0, 0 }
//            break;
//    }
// Create MAPPING-END.
// Create MAPPING-END.
//        end *yaml_node_pair_t
//            yaml_free(token.data.scalar.value);
//
//    }
//    if (!STACK_INIT(&context, items, INITIAL_STACK_SIZE)) goto error
// Create MAPPING-START.
//        version_directive_copy.major = version_directive.major
//                (octet & 0xF0) == 0xE0 ? octet & 0x0F :
// * Append an item to a sequence node.
//
//    assert(sequence > 0
//            default:
//
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
// Destroy an event object.
//
//    version_directive_copy *yaml_version_directive_t = NULL
//yaml_document_add_sequence(document *yaml_document_t,
//    }
//    mark yaml_mark_t = { 0, 0, 0 }
// Copyright (c) 2011-2019 Canonical Ltd
//}
//
//{
//        start *yaml_node_item_t
//        }
//
//    assert(document.nodes.start[mapping-1].type == YAML_MAPPING_NODE)
//
//    value_copy[length] = '\0'
// this software and associated documentation files (the "Software"), to deal in
//        version_directive_copy.minor = version_directive.minor
//            if (!value.handle || !value.prefix) goto error
// SOFTWARE.
// * Create a document object.
//    struct {
//
//        case YAML_SCALAR_TOKEN:
//    STACK_DEL(&context, items)
//        case YAML_SCALAR_TOKEN:
//        yaml_free(tag_directive.prefix)
//    tag_copy = yaml_strdup(tag)
//    struct {
//    if (!yaml_check_utf8(value, length)) goto error
//    assert(value > 0 && document.nodes.start + value <= document.nodes.top)
//    assert(document) // Non-NULL document object is expected.
///*
//}
//    if (length < 0) {
//    value_copy *yaml_char_t = NULL
//        end *yaml_tag_directive_t
///**
//        yaml_free(tag_directive.prefix)
//YAML_DECLARE(int)
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//        width = (octet & 0x80) == 0x00 ? 1 :
//        case YAML_ANCHOR_TOKEN:
// String write handler.
//    }
//yaml_token_delete(yaml_token_t *token)
//
//        unsigned int value;
//                STACK_DEL(&context, node.data.sequence.items)
//    yaml_free(tag_copy)
//    value_copy = yaml_malloc(length+1)
//
//            yaml_free(token.data.tag.suffix);
//    yaml_free(document.tag_directives.start)
// * Add a scalar node to a document.
//    value_copy *yaml_char_t = NULL
//    assert((tag_directives_start && tag_directives_end) ||
//                            // A sequence node is required.
//
// this software and associated documentation files (the "Software"), to deal in
// * Add a mapping node to a document.
//    if (!yaml_check_utf8(value, length)) goto error
//
// Destroy an emitter object.
//{
//    tag_copy = yaml_strdup(tag)
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// */
//        if (!width) return 0;
//
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
//        tag = (yaml_char_t *)YAML_DEFAULT_SEQUENCE_TAG
//        yaml_free(tag_directive.prefix)
// */
//
//
//    if (document.nodes.top != document.nodes.start) {
//static int
///*
//
//    memcpy(value_copy, value, length)
//        error yaml_error_type_t
//}
//    tag_copy *yaml_char_t = NULL
//    STACK_DEL(&context, pairs)
//                (octet & 0xF8) == 0xF0 ? 4 : 0;
//    yaml_free(value_copy)
//                goto error
//    if (!STACK_INIT(&context, pairs, INITIAL_STACK_SIZE)) goto error
//        return 0
// Set a file input.
//        return 0
//yaml_document_add_mapping(document *yaml_document_t,
//    context.error = YAML_NO_ERROR // Eliminate a compiler warning.
//    if (index > 0 && document.nodes.start + index <= document.nodes.top) {
///*
//                (octet & 0xF8) == 0xF0 ? 4 : 0;
//        tag = (yaml_char_t *)YAML_DEFAULT_SCALAR_TAG
//        error yaml_error_type_t
//    if (!PUSH(&context, document.nodes, node)) goto error
//    {
///*
//
//YAML_DECLARE(int)
// * Add a scalar node to a document.
//YAML_DECLARE(int)
//            value.prefix = yaml_strdup(tag_directive.prefix)
//    STACK_DEL(&context, nodes)
//                goto error
//    } nodes = { NULL, NULL, NULL }
//        }
//
//            break;
//    if (!PUSH(&context, document.nodes, node)) goto error
//
// String write handler.
//{
//
//}
//
//    }
//
//    } tag_directives_copy = { NULL, NULL, NULL }
//        version_directive_copy.minor = version_directive.minor
//{
// * Get a document node.
//        case YAML_TAG_DIRECTIVE_TOKEN:
//    if (!PUSH(&context,
//    {
// the Software without restriction, including without limitation the rights to
//    value_copy[length] = '\0'
//                            // Valid mapping id is required.
//            style, mark, mark)
//    } context
//    if (!tag_copy) goto error
//
// Set if unescaped non-ASCII characters are allowed.
//        return document.nodes.start + index - 1
//    if (!tag_copy) goto error
//
// */
//        error yaml_error_type_t
//        unsigned char octet;
//
//                break
//        tag_directives_end *yaml_tag_directive_t,
