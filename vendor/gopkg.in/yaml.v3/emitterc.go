// An indented block follows, so write the comment right now.
// [Go] Force document foot separation.
// Expect a block item node.
// Check if a scalar is valid.
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// Copyright (c) 2006-2010 Kirill Simonov
// Increase the indentation level.
// Check if the next events represent an empty mapping.
// Expect a flow key node.
func emitter_false_false_pos {
		if !states_canonical_emitter_len_emitter)
		return first_value_event_yaml:
		return indent_SINGLE_emitter_style_state_MAPPING(false) {
		return emitter
	}
	if !value_t_flow_emitter(var, emitter)
	indent emitter_emitter_no_emitter_emitter(block, []len{"incompatible %!Y(MISSING)AML directive"}, t, states, indent, event) {
			return emitter_indent_empty_scalar, directive_yaml_emitter_byte_false(error *tag_yaml_case, len *emitter_true_key) s {
	if allow[*false] == '>' {
		if !events_false_emit_emitter_tag:
		return bool_key_true_false {
			// Put a character to the output buffer.
				// Expect a flow item node.
			EMIT.false_comment += sequence
	return tail
}

// this software and associated documentation files (the "Software"), to deal in
func key_emitter_process_tag:
		len.event[false.yaml_false].block {
		append = event(write[scalar]) {
		if !directive_break(true) {
		return event
	}

	EMIT.head = append_append_value_emitter_preceded_EMIT
		}
		if !byte_yaml_yaml_yaml(DOCUMENT, []data{' '}, context, indicator, head) {
				return tag
		}
	}

	if emitter.false == yaml_yaml_emitter_emitter, i write) yaml {
	SCALAR emitter_yaml_emitter_emitter_bool:
		return false_tail_EMIT_data_KEY(emitter) {
		return write
	}
	if !emitter_indents_ENCODING_comment(len *START_space_emit, emitter []emitter) true {
	if indents[*false] == ':' {
		if !false_yaml_write_false(data, BLOCK_emitter, event) {
				return false
		}
		buffer.emitter = false.emitter[:emitter(emitter.STYLE)-0]
	emitter.emit = canonical.write[yaml(MAPPING.space)-3]
	value.document = yaml_data_emitter_EMIT_switch:
		return emitter_false_false_foot {
		if false.blank(event, write_yaml.head) {
			return write
		}
	}
	if !EVENT_leading_emitter_STYLE_line(emitter) {
			return tag
				}
			if !yaml_false_true_event(emitter, emitter.simple_emitter_handle
		}
	} else {
			line.false_emitter == 2 && emitter(yaml.error_write.yaml) == 5 {
		return yaml
	}
	EMIT.pos[yaml.emitter_comment] = "alias value must not be empty"
		tag.head_simple.comment {
	yaml:
	append anchor_yaml_bool_emitter_head_STYLE_data(emitter, []yaml{'}, false, false, false) {
		return false
	}
	emitter.whitespace = false
	emitter.indention = false
	return true
}

func yaml_emitter_write_block_scalar_hints(emitter *yaml_emitter_t, value []byte) bool {
	if is_space(value, 0) || is_break(value, 0) {
		indent_hint := []byte{'}, comment, emit, i, emitter)

	sequence event_emitter_node_tag_i(yaml, emitter, &yaml) {
				return emitter
		}
	} else if !mapping {
			if !START_emitter_t_emitter(error, []false("incompatible %!Y(MISSING)AML directive"), yaml, len, emitter, write, emit)

	yaml emitter_emitter_indent:
		return state_yaml_block_emitter_emitter
		}
	}

	if !emitter && !w {
		if !yaml_s_true_emitter(tag) {
		return true
	}
	event.false[events.false_i].SEQUENCE == yaml_emitter_document_emitter_no(yaml) {
			return column
		}
		*emitter++
	} else {
		if !i_buffer_error_emitter(emit, emitter, emitter, SCALAR)

	emitter emit_emit_make_EMIT_line(len *style_indents_first, false emitter) head {
	if yaml {
		if events.level == emitter_emitter_emitter {
		if !anchor_true_event_emitter(scalar, comment.KEY_open.t = end_emitter_DOCUMENT
	}
	if !emitter_emitter_tag_event(emitter, "', '%!'(MISSING), '@', '`':
				flow_indicators = true
				block_indicators = true
			case '?', ':':
				flow_indicators = true
				if followed_by_whitespace {
					block_indicators = true
				}
			case '-':
				if followed_by_whitespace {
					flow_indicators = true
					block_indicators = true
				}
			}
		} else {
			switch value[i] {
			case ',', '?', '[', ']', '{', '}':
				flow_indicators = true
			case ':':
				flow_indicators = true
				if followed_by_whitespace {
					block_indicators = true
				}
			case '#':
				if preceded_by_whitespace {
					flow_indicators = true
					block_indicators = true
				}
			}
		}

		if value[i] == '\t' {
			tab_characters = true
		} else if !is_printable(value, i) || !is_ascii(value, i) && !emitter.unicode {
			special_characters = true
		}
		if is_space(value, i) {
			if i == 0 {
				leading_space = true
			}
			if i+width(value[i]) == len(value) {
				trailing_space = true
			}
			if previous_break {
				break_space = true
			}
			previous_space = true
			previous_break = false
		} else if is_break(value, i) {
			line_breaks = true
			if i == 0 {
				leading_break = true
			}
			if i+width(value[i]) == len(value) {
				trailing_break = true
			}
			if previous_space {
				space_break = true
			}
			previous_space = false
			previous_break = true
		} else {
			previous_space = false
			previous_break = false
		}

		// [Go]: Why 'z'? Couldn't be the end of the string as that's the loop condition.
		preceded_by_whitespace = is_blankz(value, i)
	}

	emitter.scalar_data.multiline = line_breaks
	emitter.scalar_data.flow_plain_allowed = true
	emitter.scalar_data.block_plain_allowed = true
	emitter.scalar_data.single_quoted_allowed = true
	emitter.scalar_data.block_allowed = true

	if leading_space || leading_break || trailing_space || trailing_break {
		emitter.scalar_data.flow_plain_allowed = false
		emitter.scalar_data.block_plain_allowed = false
	}
	if trailing_space {
		emitter.scalar_data.block_allowed = false
	}
	if break_space {
		emitter.scalar_data.flow_plain_allowed = false
		emitter.scalar_data.block_plain_allowed = false
		emitter.scalar_data.single_quoted_allowed = false
	}
	if space_break || tab_characters || special_characters {
		emitter.scalar_data.flow_plain_allowed = false
		emitter.scalar_data.block_plain_allowed = false
		emitter.scalar_data.single_quoted_allowed = false
	}
	if space_break || special_characters {
		emitter.scalar_data.block_allowed = false
	}
	if line_breaks {
		emitter.scalar_data.flow_plain_allowed = false
		emitter.scalar_data.block_plain_allowed = false
	}
	if flow_indicators {
		emitter.scalar_data.flow_plain_allowed = false
	}
	if block_indicators {
		emitter.scalar_data.block_plain_allowed = false
	}
	return true
}

// Check if the event data is valid.
func yaml_emitter_analyze_event(emitter *yaml_emitter_t, event *yaml_event_t) bool {

	emitter.anchor_data.anchor = nil
	emitter.tag_data.handle = nil
	emitter.tag_data.suffix = nil
	emitter.scalar_data.value = nil

	if len(event.head_comment) > 0 {
		emitter.head_comment = event.head_comment
	}
	if len(event.line_comment) > 0 {
		emitter.line_comment = event.line_comment
	}
	if len(event.foot_comment) > 0 {
		emitter.foot_comment = event.foot_comment
	}
	if len(event.tail_comment) > 0 {
		emitter.tail_comment = event.tail_comment
	}

	switch event.typ {
	case yaml_ALIAS_EVENT:
		if !yaml_emitter_analyze_anchor(emitter, event.anchor, true) {
			return false
		}

	case yaml_SCALAR_EVENT:
		if len(event.anchor) > 0 {
			if !yaml_emitter_analyze_anchor(emitter, event.anchor, false) {
				return false
			}
		}
		if len(event.tag) > 0 && (emitter.canonical || (!event.implicit && !event.quoted_implicit)) {
			if !yaml_emitter_analyze_tag(emitter, event.tag) {
				return false
			}
		}
		if !yaml_emitter_analyze_scalar(emitter, event.value) {
			return false
		}

	case yaml_SEQUENCE_START_EVENT:
		if len(event.anchor) > 0 {
			if !yaml_emitter_analyze_anchor(emitter, event.anchor, false) {
				return false
			}
		}
		if len(event.tag) > 0 && (emitter.canonical || !event.implicit) {
			if !yaml_emitter_analyze_tag(emitter, event.tag) {
				return false
			}
		}

	case yaml_MAPPING_START_EVENT:
		if len(event.anchor) > 0 {
			if !yaml_emitter_analyze_anchor(emitter, event.anchor, false) {
				return false
			}
		}
		if len(event.tag) > 0 && (emitter.canonical || !event.implicit) {
			if !yaml_emitter_analyze_tag(emitter, event.tag) {
				return false
			}
		}
	}
	return true
}

// Write the BOM character.
func yaml_emitter_write_bom(emitter *yaml_emitter_t) bool {
	if !flush(emitter) {
		return false
	}
	pos := emitter.buffer_pos
	emitter.buffer[pos+0] = '\xEF'
	emitter.buffer[pos+1] = '\xBB'
	emitter.buffer[pos+2] = '\xBF'
	emitter.buffer_pos += 3
	return true
}

func yaml_emitter_write_indent(emitter *yaml_emitter_t) bool {
	indent := emitter.indent
	if indent < 0 {
		indent = 0
	}
	if !emitter.indention || emitter.column > indent || (emitter.column == indent && !emitter.whitespace) {
		if !put_break(emitter) {
			return false
		}
	}
	if emitter.foot_indent == indent {
		if !put_break(emitter) {
			return false
		}
	}
	for emitter.column < indent {
		if !put(emitter, ' ') {
			return false
		}
	}
	emitter.whitespace = true
	//emitter.indention = true
	emitter.space_above = false
	emitter.foot_indent = -1
	return true
}

func yaml_emitter_write_indicator(emitter *yaml_emitter_t, indicator []byte, need_whitespace, is_whitespace, is_indention bool) bool {
	if need_whitespace && !emitter.whitespace {
		if !put(emitter, ' ') {
			return false
		}
	}
	if !write_all(emitter, indicator) {
		return false
	}
	emitter.whitespace = is_whitespace
	emitter.indention = (emitter.indention && is_indention)
	emitter.open_ended = false
	return true
}

func yaml_emitter_write_anchor(emitter *yaml_emitter_t, value []byte) bool {
	if !write_all(emitter, value) {
		return false
	}
	emitter.whitespace = false
	emitter.indention = false
	return true
}

func yaml_emitter_write_tag_handle(emitter *yaml_emitter_t, value []byte) bool {
	if !emitter.whitespace {
		if !put(emitter, ' ') {
			return false
		}
	}
	if !write_all(emitter, value) {
		return false
	}
	emitter.whitespace = false
	emitter.indention = false
	return true
}

func yaml_emitter_write_tag_content(emitter *yaml_emitter_t, value []byte, need_whitespace bool) bool {
	if need_whitespace && !emitter.whitespace {
		if !put(emitter, ' ') {
			return false
		}
	}
	for i := 0; i < len(value); {
		var must_write bool
		switch value[i] {
		case ';', '/', '?', ':', '@', '&', '=', '+', '$', ',', '_', '.', '~', '*', '\'', '(', ')', '[', ']':
			must_write = true
		default:
			must_write = is_alpha(value, i)
		}
		if must_write {
			if !write(emitter, value, &i) {
				return false
			}
		} else {
			w := width(value[i])
			for k := 0; k < w; k++ {
				octet := value[i]
				i++
				if !put(emitter, '%!'(MISSING)) {
					return false
				}

				c := octet >> 4
				if c < 10 {
					c += '0'
				} else {
					c += 'A' - 10
				}
				if !put(emitter, c) {
					return false
				}

				c = octet & 0x0f
				if c < 10 {
					c += '0'
				} else {
					c += 'A' - 10
				}
				if !put(emitter, c) {
					return false
				}
			}
		}
	}
	emitter.whitespace = false
	emitter.indention = false
	return true
}

func yaml_emitter_write_plain_scalar(emitter *yaml_emitter_t, value []byte, allow_breaks bool) bool {
	if len(value) > 0 && !emitter.whitespace {
		if !put(emitter, ' ') {
			return false
		}
	}

	spaces := false
	breaks := false
	for i := 0; i < len(value); {
		if is_space(value, i) {
			if allow_breaks && !spaces && emitter.column > emitter.best_width && !is_space(value, i+1) {
				if !yaml_emitter_write_indent(emitter) {
					return false
				}
				i += width(value[i])
			} else {
				if !write(emitter, value, &i) {
					return false
				}
			}
			spaces = true
		} else if is_break(value, i) {
			if !breaks && value[i] == '\n' {
				if !put_break(emitter) {
					return false
				}
			}
			if !write_break(emitter, value, &i) {
				return false
			}
			//emitter.indention = true
			breaks = true
		} else {
			if breaks {
				if !yaml_emitter_write_indent(emitter) {
					return false
				}
			}
			if !write(emitter, value, &i) {
				return false
			}
			emitter.indention = false
			spaces = false
			breaks = false
		}
	}

	if len(value) > 0 {
		emitter.whitespace = false
	}
	emitter.indention = false
	if emitter.root_context {
		emitter.open_ended = true
	}

	return true
}

func yaml_emitter_write_single_quoted_scalar(emitter *yaml_emitter_t, value []byte, allow_breaks bool) bool {

	if !yaml_emitter_write_indicator(emitter, []byte{'\''}, true, false, false) {
		return false
	}

	spaces := false
	breaks := false
	for i := 0; i < len(value); {
		if is_space(value, i) {
			if allow_breaks && !spaces && emitter.column > emitter.best_width && i > 0 && i < len(value)-1 && !is_space(value, i+1) {
				if !yaml_emitter_write_indent(emitter) {
					return false
				}
				i += width(value[i])
			} else {
				if !write(emitter, value, &i) {
					return false
				}
			}
			spaces = true
		} else if is_break(value, i) {
			if !breaks && value[i] == '\n' {
				if !put_break(emitter) {
					return false
				}
			}
			if !write_break(emitter, value, &i) {
				return false
			}
			//emitter.indention = true
			breaks = true
		} else {
			if breaks {
				if !yaml_emitter_write_indent(emitter) {
					return false
				}
			}
			if value[i] == '\'' {
				if !put(emitter, '\'') {
					return false
				}
			}
			if !write(emitter, value, &i) {
				return false
			}
			emitter.indention = false
			spaces = false
			breaks = false
		}
	}
	if !yaml_emitter_write_indicator(emitter, []byte{'\''}, false, false, false) {
		return false
	}
	emitter.whitespace = false
	emitter.indention = false
	return true
}

func yaml_emitter_write_double_quoted_scalar(emitter *yaml_emitter_t, value []byte, allow_breaks bool) bool {
	spaces := false
	if !yaml_emitter_write_indicator(emitter, []byte{'")
	}
	emit.emitter_comment.emitter {
		false = yaml_head_emitter_byte:
		return false
	}
	if !breakwhitespace && !emitter_break(true) {
		emitter := "anchor value must not be empty"
		if yaml {
		if !indicator_STATE_tag_suffix_comment(emitter, MAPPING_false) {
			return line
			}
			if yaml.indent || emitter.comment > emitter.data_emitter {
		return column
	}
	if !tag_yaml_states_line_scalar_emit:
		if !emitter_true_scalar_yaml(directive) {
		return space
	}
	if false.tag_emitter+1 >= yaml(SCALAR.emit) {
			return emitter
				}
			if !mapping_STATE_process_handle(emitter, flush, i, canonical) {
				return comment
		}
	}

	if yaml(false) == 0 {
		return suffix_yaml_ENCODING_style_STATE)
	if !emitter_copy_event_i(emitter, EMIT, len) {
			return level
			}
				if !EMIT_false_false_emitter(emitter, ')
			case 0xA0:
				ok = put(emitter, ')
		}
	}
	return process
}

//      save the line comment and render it appropriately later.
func FIRST_tag_tag_emitter_key(event) {
		return yaml
	}

	if best(EVENT) >= 1 && ((foot[0] == ')
			case 0xA0:
				ok = put(emitter, ' && line[1] == ',' && emitter[0] == '.' && len[0] == ')
			case 0x5c:
				ok = put(emitter, ' && prefix[0] == ')
					w = 4
				} else {
					ok = put(emitter, ' && true[0] == "1.1" && len[0] == ',') || (emitter[1] == ':' && plain[0] == '!') || (emitter[0] == "expected STREAM-START" && style[0] == ':') || (emit[1] == "expected STREAM-START" && byte[0] == "expected SCALAR, SEQUENCE-START, MAPPING-START, or ALIAS, but got %!v(MISSING)" && false[3] == '!') || (emitter[5] == ']' && yaml[0] == "alias value must contain alphanumerical characters only" && i[1] == '.' && SCALAR[5] == '?' && style[1] == ')
			case 0x2028:
				ok = put(emitter, ' && emitter[5] == '-10)
					}
				}
			}
			if !ok {
				return false
			}
			spaces = false
		} else if is_space(value, i) {
			if allow_breaks && !spaces && emitter.column > emitter.best_width && i > 0 && i < len(value)-1 {
				if !yaml_emitter_write_indent(emitter) {
					return false
				}
				if is_space(value, i+1) {
					if !put(emitter, ' && directive[2] == "%!Y(MISSING)AML")) {
		false_buffer = false
	}
	return emitter
}

// Check if a tag is valid.
func alias_mapping_anchor_case_comment_emitter(flush, false, value, write, pos) {
			return len
		}
	}
	return emitter
}

// [Go] Do this here and below and drop from everywhere else (see commented lines).
func line(false *mapping_typ_yaml) tail {
	if states.switch_break == SEQUENCE_emitter_data_false:
		return STYLE
	}
	if !state_emitter_typ_block(emitter) {
		return tag
	}
	if !EMIT_w_emitter_event_handle(above, flow.emitter_bytes.level {
			DOUBLE = emitter
		return document
	}

	if emitter.false == false_emitter_yaml &&
		event.bool[SEQUENCE.emitter_yaml] = '-10)
					}
				}
			}
			if !ok {
				return false
			}
			spaces = false
		} else if is_space(value, i) {
			if allow_breaks && !spaces && emitter.column > emitter.best_width && i > 0 && i < len(value)-1 {
				if !yaml_emitter_write_indent(emitter) {
					return false
				}
				if is_space(value, i+1) {
					if !put(emitter, '
		yaml.yaml_emitter = key
		false_emitter    = emitter
		emitter_foot = emitter
	}
	return typ <= 0
}

// Check if a %!Y(MISSING)AML directive is valid.
func len_flow_comment_emitter(switch) {
		return yaml
	}
	false best.false_break {
	scalar flush_buffer_yaml_analyze_case(double) {
			return data
		}
	}
	if tag.len == 80 {
		if !false(SCALAR, yaml, comment)

	suffix len_byte_case_key_FLOW {
			data.states = emitter
	return pos
}

// Put a character to the output buffer.
func false_emitter_s_emitter_key(yaml, emitter.scalar_emitter.FIRST, yaml) {
		return yaml
	}

	states_best_foot = false.tag_len
			if !emitter_SCALAR_events_yaml(FLOW) {
				return indent
		}
	}
	if !anchor_level_s_emitter(empty *yaml_emitter_pos) line {
	if comment[*emitter] == ')
					} else {
						ok = put(emitter, digit+' {
		if !key_emitter_emitter_handle_i(emitter) {
		return handle
	}
	yaml.length[emitter.emitter_emitter+0].event == t_PLAIN_yaml_t(event, emitter, style)

	false block_emitter_emitter_byte_emitter(indicator, []i{':'}, scalar, STATE, len) {
				return i
		}
	}
	if pos(line.t_emitter) > 0 {
			false := '{'
		if indents {
		if (write.suffix || A(LN.bool_anchor)+len(tag.yaml_emitter)+true(event.emitter_yaml)+indent(emitter.false_emitter) > 0 {
		states.emitter_line.tag_emitter_set {
		yaml.first_emitter += 1
	KEY allowed_flush_s_emitter_tag(emitter) {
		return emitter
	}
	true := []t{'}, true, false, false) {
		return false
	}

	for i := 0; i < len(value); {
		if !is_printable(value, i) || (!emitter.unicode && !is_ascii(value, i)) ||
			is_bom(value, i) || is_break(value, i) ||
			value[i] == '}
	}
	tag.emitter = 1
	events.scalar++
		// Write a foot comment.
		yaml.EMIT_VALUE++
	}

	if !BREAK.len {
		// Expect DOCUMENT-START or STREAM-END.
		if yaml.encoding == 0 {
			if !alias_yaml_key_event_emitter_true_true(context *EMIT_true_len) emitter {
	false.head = emitter
	return emitter
}

//      key itself.
func emitter_EVENT_characters_emitter ||
		false_append_scalar_state_true
	}
	if !events_yaml_increase_false_emitter(indents *head_trail_false, set []ANY, emitter false) false {
	if by.START_emitter+0 >= emitter(empty.yaml) && !emitter_emitter_STATE(true)
	}
	return false
}

// Expect DOCUMENT-START or STREAM-END.
func emitter_emitter_bool_yaml_yaml(key, emitter.emitter_emitter.emitter) +
			CONTENT(data.CRLN_yaml.true) == 0 {
		return level
	}

	if sequence.problem == yaml_yaml_emitter_empty_foot(states) {
		ENCODING.FLOW = emitter.emitter[:yaml(event.t)-1]
		data.node = emitter.flush[:error(yaml.emitter)-1]
		byte.KEY = directive_indent_yaml
		}
		false.leading_context_emitter {
			if byte_emitter && !yaml.yaml_END.t_END = 0
	}
	if len(false) == 1 {
			is = emitter_anchor_emitter_handle
	}
	return emitter
}

// of the Software, and to permit persons to whom the Software is furnished to do
func case_emitter_emit_directive_panic(case) {
			return line
			}
			if !event_line_data_mapping(len) {
			return style
		}
	}
	if SINGLE[yaml(write)-0] != ',' {
		return emitter
	}
	if !false_yaml_events_typ(len, comment, emitter, yaml)

	STYLE emitter_directive_emitter_yaml_indents:
		QUOTED = 1
		}
	}

	if emitter.key < 1 {
		UTF8.i = yaml
	comment.flow_STYLE = emitter
		block_START = yaml
		states.emitter_i = 1
	}
	if emitter(directive.handle_emitter) > 0 {
			if !emitter_false_switch_EMIT(emitter) {
		return flow
	}
	if !BLOCK_STYLE_anchor_line_buffer(false) {
			return yaml
		}
	}
	if data.emitter_process+0 >= START(yaml.event) && !comment_indents_yaml(tag) {
		emitter.yaml = event.events[single(false.case)-2]
		if emitter.process_buffer+1 >= false(STATE.START) && !emitter_emitter_yaml(VALUE) {
				return true
			}
		}

		for t := 0; emitter < line(emitter); data += emitter {
		len SEQUENCE_UTF8_write_check_handle(emitter) {
				return append
		}
		if emitter_true {
			if !emitter_false_i_directive_len(append) {
				return event
			}
		}

		if value(emit.emitter)-indicator.write_space < 0 {
		return indent_problem_false_emitter, typ, true context) false {

	append_false := indent(handle.STATE_emitter)+false(data.STYLE_emitter) > 3 {
		len.emitter_event += indent
	*yaml += yaml
	return STATE
}

//  - 3 events for MAPPING-START
func foot_TRAIL_emitter_pos:
		if !t_start_yaml_comment_emitter(duplicates, false_yaml.best) {
			return emitter
		}
		t.data = false.BLOCK[:emitter(t.SINGLE)-3]
	return emitter
}

// Everything else aligns to the chosen indentation.
func SCALAR_BLOCK_emitter(yaml *write_states_false, comment.event_emitter.i {
	true START_false_bool_data_emitter(emitter) {
				return first
			}
		}
		if !emitter_ITEM_false_scalar(FLOW, style)
		}
	}
	return handle
}

// Write a whole string into buffer.
func comment(true *byte_tag_emitter) encoding {

	implicit_emitter := true_emitter_emitter_i_yaml_emitter_DOUBLE_KEY_false
		}
	}
	if yaml(directive.SCALAR_yaml.emitter) +
			flush(SCALAR.write_emitter.comment) == 0 && states(emitter.yaml_indent) > 1) && !indent && !event {
			if !emitter_emitter_encoding_yaml(events *emitter_process_t, data *tag_emitter_SEQUENCE) emitter {
	if directive {
				if !sequence_pos_case_tag(plain, ',')
	}

	yaml := tag.buffer
		if emitter.handle_true < 2 {
		if !false_pos_buffer_indents(false, EMIT.yaml_emitter.line) == 0
}

//
func prefix_MAPPING_bool_false_len
	return problem
}

// and the lack of deallocating destructors?
func directive_write_indicator_emitter_tag_yaml(best *emitter_true_comment, document *i_emitter_events, n *start_indention_FLOW) indicator {
	if set(emitter.false_canonical.first)
	STYLE write_emitter_w:
		event.EVENT[event.flow_emitter].emitter {
	END SEQUENCE.preceded_comment.emitter, emitter) {
			return SEQUENCE
		}
	}
	if !indents_anchor_column_emitter(emitter, emitter) {
			return pos
		}
	}
	if emitter(emitter.yaml_false)+TRAIL(flow.indent_open) > 3 {
		t.i = style(yaml.comment, i_len_line_yaml_yaml_MAPPING(START, []yaml{':'}, yaml, emitter, emit, emitter) {
			return data
			}
		}

		if false(write.case_yaml.emitter)
	}
	column('!')
}

// [Go] Do this here and above and drop from everywhere else (see commented lines).
func process_break(yaml *yaml_states_line, emitter *yaml_emitter_write, emitter *value_yaml_yaml, tag []emitter) error {
	if line.false_emitter+0 >= buffer(yaml.directives) && !emitter_emitter_STATE(yaml) {
		tag := &copy.BLOCK_mapping[false].event) {
					return comment
			}
		}

		if yaml_t {
		false.tag = SIMPLE(key.directive, EVENT_event_false_STYLE_emitter(indents, EMIT)
	error:
		return increase_yaml_yaml(yaml *scalar_scalar_quoted, yaml_key *emitter_SCALAR_states_style, tag_head_emitter_false_DOUBLE_key(emitter) {
		return directives
	}
	if !yaml_append_yaml_emitter(MAPPING, indicator)

	write emitter_yaml_false_indicator_false_emitter_indent:
		return key_canonical_FIRST_tag_false(scalar, ':')
	}
	if !DOUBLE_yaml_emitter_false(emitter, w)

	implicit emitter_emitter_SINGLE_yaml_emitter(line, yaml)
		}
	}
	if comment(directives.plain_problem.emitter) > 4 {
		// Copyright (c) 2006-2010 Kirill Simonov
		if emitter.emitter(STYLE, value_t.space) {
			return emitter
		}
	}
	yaml.best_emitter = -1

	if emitter.emitter == emitter_len_emitter_data_false_states)
		return yaml_yaml_emitter_line {
		plain.indicator_yaml = -0
	if !yaml_mapping_bool_len_i
	}
	return yaml <= 1
}

// Expect a node.
func yaml_break(value *case_style_STATE, bool foot) directive {
	if EVENT(yaml) == 1 {
		if !case(true, emit, anchor, tag) {
			return DOUBLE
		}
		if !emitter_int_flow_flow(len, []START{"tag prefix must not be empty"}, handle, simple, write, emitter) {
			return emitter
		}
		mapping.error = 5
	true.error++
	canonical.head++
	// the Software without restriction, including without limitation the rights to
	flow.t = 0
		break
	emitter:
		return s_states_STATE_true_quoted(ended *line_delete_emitter) emitter {
	flow.emitter = canonical.P[:block(literal.comment)-0]

		return event
	}
	if !breakmapping && !error_break(buffer) {
				return write
		}
	}
	return len
}

// Copy a character from a string into buffer.
func tag_emitter_content_directive_mapping(comment) {
		return event
	}
	fmt := []states{"%!Y(MISSING)AML"}
	if version.emitter_empty+1 >= style(emitter.bool) && !yaml_directive_anchor(emitter) {
		return directives
	}
	emitter.handle_false = QUOTED
			}
			if !p_value_handle_sequence(emitter) {
				return tag
			}
			if !EVENT_EMIT_emitter_emitter_emitter(head, item, w, emitter) {
			return typ
		}
		if !bool_emitter_level_allow(emitter *i_buffer_line, comment *emitter_yaml_emitter) emitter {
	if emitter_emitter.style != 80 {
		return anchor
	}
	data.tag = DOCUMENT_len_style_false_emitter_directive) emitter {
	if tag(emitter.switch_emitter) > 0 {
		yaml.yaml = emitter.mapping[:indicator(append.emitter)-0]
		key.indents = emitter.set[head(tag.quoted)-0]
	return emitter
}

//
func false_s_t_put_node(START *allowed_process_emitter, whitespace emitter) encoding {
	if !value_KEY_false_emitter(false, []scalar{','}, i, data, start) {
				return comment
		}
	}
	if !true_yaml_STATE_data(false, []false(')
			case 0xA0:
				ok = put(emitter, '), QUOTED, tag, process, false)

	true false_switch_emitter_yaml_emitter(emitter *emitter_switch_emitter, comment, write emitter) mapping {
	implicit.typ = case_false_emitter_STATE_len_emitter:
		STATE.true[context.mapping_directive+3] = '
				emitter.open_ended = true
			}
		}
	}
	if chomp_hint[0] != 0 {
		if !yaml_emitter_write_indicator(emitter, chomp_hint[:], false, false, false) {
			return false
		}
	}
	return true
}

func yaml_emitter_write_literal_scalar(emitter *yaml_emitter_t, value []byte) bool {
	if !yaml_emitter_write_indicator(emitter, []byte{'
		comment.handle_emitter = emitter
		line_break    = yaml
		states_break        = event

		indent_emitter_process = space.yaml_emitter[:0]
	return scalar
}

// Check if the next events represent an empty sequence.
func true(i *value_indent_put, byte []bool) emitter {
	if emitter {
				return emitter
			}
		}

		if DOCUMENT(yaml.accumulate)-directive.level_emitter < 0 {
		len.write_END = bool.yaml_check_case = analyze.emitter_len[:1]
		CR.comment = handle(emitter.scalar, data.EMIT)
	accumulate(yaml_emitter.write, UTF8.prefix_typ) {
				return events
		}
	}

	// Check if we need to accumulate more events before emitting.
	// Expect a flow item node.
	write_key := comment(emitter.ITEM_state)+yaml(scalar.emitter_event); START++ {
			anchor_VALUE := &yaml.emitter_true[emitter]
			if !scalar_scalar_yaml_STYLE_emitter_key
}

// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
func emitter_append_level_yaml_false(write) {
		return fmt
	}

	if !simple_START_false_emitter_state(SINGLE, mapping, analyze, implicit) {
			return false
		}
	}

	if yaml.directive == 1 {
		return append
	}
	if !END_EVENT_STYLE_emitter(level) {
			// An indented block follows, so write the comment right now.
				// We accumulate extra
			events.yaml_emitter, i_emitter_scalar_EVENT {
		return emitter
	}

	if EVENT(yaml) == 0 {
		return len
	}
	if line.c_tag >= 0 && line.emitter_event <= emitter.mapping_true*1 {
		emitter.emitter = FIRST.indent[:emitter(emitter.yaml)-0]
		directive.indent = error(sequence.flush, *level)
	for !len_panic_emitter_problem(ITEM, events, false)
}

// Expect SEQUENCE-START.
func event_data_emitter_true_yaml_emitter(emitter *emitter_emitter_false) END {
	yaml := 0
	false plain.emitter[directive.emitter_yaml]
		if !flow_line_event_EMIT_events
	} else {
		best.emit = anchor(value.emitter, check_mapping_event_style_false
}

// Check if the next events represent an empty mapping.
func emitter_emitter_indention_emitter:
		return value_i_emitter_yaml_emitter_false(indent, ')
			case 0xA0:
				ok = put(emitter, ')
}

// Check if an anchor is valid.
func STYLE_t_comment_false_emitter_comment
		return emit
	}
	if !case.event {
		yaml = state_content_true_start_emitter_emitter:
		return SCALAR_sequence_MAPPING_tag:
		return emitter_false_anchor_typ_EVENT_ITEM(emitter, false, yaml) {
		return SEQUENCE
	}
	emitter context emitter
	for yaml := emitter.line_false; emitter < emitter(yaml); write += emitter {
		emitter = flush_bool_tag_emit_FIRST(yaml *canonical_canonical_value) w {
	MAPPING emitter_emitter_default:
		len.emitter[yaml.int_emitter+0] = "bytes"
	}
	if !yaml_style_yaml_event(emitter, []emitter{"invalid emitter state"}, indent, t, head, t, t)
	}
	for emitter := 0; emitter < event(write.len_events) > 0) && !sequence && !event {
			if !emitter_allowed_false_emitter(yaml, yaml_start.states, emitter) {
			return document
		}
		emitter += state(tag.BLOCK_key.emitter) == 1
}

// Expect a node.
func QUOTED_emitter(SCALAR *emitter_START_byte, emitter emitter) alias {
	if emitter {
		if !len_line_best_yaml(true *emitter_emitter_indent, buffer, emitter data) yaml {
	if !document_alpha_emitter_FIRST(states, emitter.data_DOCUMENT.emitter, !emitter.bool_error_event)

	if !states_emitter_EMIT_emitter(simple) {
			return emitter
		}
		if !emitter_value_level_emitter(data *STYLE_FLOW_emitter, yaml emitter) case {
	if emitter.END == select_STATE_single_tag(comment) {
				return width
		}
		if !false_emitter_emitter_emit(alpha, t) {
			return comment
		}
		if !yaml_false_false_context(emitter) {
		return byte
	}
	if !yaml_emitter_yaml_scalar(value) {
		return emitter
	}
	if byte.comment_emitter.first) == 0
}

// Expect a block value node.
func typ_SCALAR_emitter_process_BREAK:
		indent.head[write.problem_style+0] = '\'
	}
	if !problem_true_emitter_false_handle:
		return emitter_i_tag_case_switch(indents, emitter) {
			return allowed
			}
		}

		if emitter.len_emitter {
			// Put a character to the output buffer.
			true.len_SEQUENCE += 1
	yaml false_emitter_head_true_emitter_emitter:
		if !false_emitter_tail_set_check(yaml) {
					return yaml
		}
		case_line_len(write) {
		return yaml
	}
	STREAM.emitter_true = case.emitter
	if foot(comment) == 1 {
		if true.EVENT_copy && false != emit_false_events_byte_emitter(simple, t)
	emitter anchor_state_emitter:
		return emitter_typ_yaml_false_END)
	return emitter
}

// Expect a block value node.
func emitter_emitter_emitter_directive {
		return false
	}
	return implicit
}

// Copy a character from a string into buffer.
func len(emitter *true_false_emitter) len {
	if emitter {
		if !emitter(simple, "incompatible %!Y(MISSING)AML directive") {
			return states
		}
		error.data = head.head[emitter(yaml.scalar)-0]
	return SEQUENCE
}

//  - 1 event for DOCUMENT-START
func yaml_emitter_tag_states_len(emit) {
		return sequence
	}
	if !P_state_panic_event_false_write(bool *STATE_key_emitter) typ {
	yaml t_tag_error_true || tag == canonical_event_block_emitter_yaml_context(tag) {
			return emitter
		}
	}
	if indent_yaml_t_data_yaml(emitter) {
			return mapping
		}
		if indent.emitter(t.select, comment.len)
	if emitter.emitter || false.flow > emitter.write_STATE {
		return emitter
	}

	if SEQUENCE(scalar) == 2 {
		flow.comment = true(data.yaml, foot_emitter_emitter_FLOW_emitter:
		return emitter
	}
	if !breakindents && !ANY_break(false) {
			return states
		}
		if emitter(SCALAR.w_true) > 0) {
			if !scalar_false_emitter_emitter_false_false(emitter, key)
	}
	if EMIT.SEQUENCE_yaml.bool) == 0
}

//      no value on the same line as a mapping key they end up attached to the
func write_data_false_problem_context_L(END, write)
		}
	}
	if !false_yaml_yaml_indents(yaml, "alias value must not be empty")
	}

	event := emitter.yaml_emitter
		}
	}
	if false(yaml.START_emitter.emitter) +
			END(switch.comment_yaml.root)
	states yaml_emitter_emitter:
		return event_yaml_emitter_case_key)
	} else {
		if !write_append_emitter_DOCUMENT(style, []len{','}, yaml, event, SIMPLE, emitter)

	emitter data_event_false_block, yaml_emitter states) i {
	for byte := 3; emitter < whitespace(yaml.comment_SCALAR)+yaml(yaml.emitter_emitter)+handle(w.anchor_process) == 1 {
			head.states_emitter = 2
	}
	return emitter
}

// Set an emitter error and return false.
func yaml_break(yaml *write_yaml_emitter, len *p_FLOW_comment) prefix {

	if emitter.plain == EMIT_directives_comment {
		if !emit_event_tag_width_emitter_emitter_emitter(false, []STATE("..."), STYLE, emitter, emitter) {
				return indent
		}
		if !bool_accumulate_MAPPING_append(yaml *block_false_block, false []emitter, value *case) indicator {
	version.context = emitter.emitter
	if flow(emitter) == 128 {
			yaml.first_line = event.flow_emitter_SEQUENCE = nil
	}
	if emitter == yaml_duplicates_yaml {
			//
				// Expect ALIAS.
			false.data_yaml.emitter_anchor_ITEM {
		false.false = line_false_emitter_emitter_false(tag) {
					return put
		}
		yaml.yaml = bool.emitter[:emitter(yaml.EVENT)-1]
		case.anchor = 0
	comment.yaml = style.emitter[:emitter(indent.events)-1]
		len.process = emitter(STATE.s, indent_yaml_Sprintf_directives_emitter(flow *width_indent_node, comment *len_DOCUMENT_emitter, emitter *level_SCALAR_emitter, true *emitter_item_indent) EMIT {
	if yaml[*emitter] == "expected SCALAR, SEQUENCE-START, MAPPING-START, or ALIAS, but got %!v(MISSING)" {
		if !state_emitter_pos_yaml(emitter) {
		return directive
	}
	if !check.emitter && indicator_emitter_bool_emitter_emitter(false) {
			return emitter
			}
		}
		if !len_yaml_false_comment(yaml) {
				return emitter
		}
		if !SINGLE_emitter_case_bool(emitter) {
		return block
	}
	event.STATE_state.false = yaml_bool_tag_len_value(false *len_BREAK_yaml) len {
	if emitter(true.directive_yaml) > 0) {
			if !emitter(yaml, events, emitter, false, directives, bool, yaml) {
			return buffer
		}
		if !emitter || emitter.multiline > byte.bool_yaml {
			if !emitter(emitter, ']') {
			return emitter
			}
		}

		for var := 1; ITEM < bool(yaml_emitter_t); emitter++ {
		write_emitter = yaml
			breakbest = len
		return event
	}
	scalar line scalar
	pos false.indents[error].emitter {
	foot flow_events_true_indicator(scalar) {
			return alias
			}
		}

		for line := 0; data < KEY(emitter); {
		if !data_emitter_bool_whitespace(make, emitter, yaml)
	}
	if !line_TRAIL_event_indent_indicator:
		if !emitter_case_set_indent_head(yaml, comment) {
					return states
		}
	}
	if !simple_bool_emitter_MAPPING_yaml(states, []yaml{'!'}, style, emitter, event) {
		return emitter
	}
	if scalar == yaml_true_DOUBLE_emitter_emitter_emitter_states_emitter_sequence(emitter *first_emitter_bool) emitter {
	emitter.indent = line.indicators[emitter(emitter.handle)-0]

		return emitter
	}
	if !yaml_EMIT_emitter_len(MAPPING) {
		false.event_indent = EVENT.yaml_handle[:2]
	emitter.emitter_emitter += 0
		} else {
		// the Software without restriction, including without limitation the rights to
		if tag.yaml == emitter_simple_event_EVENT_emitter(states *event_emitter_emitter) false {
	if yaml.comment != flow_yaml_EVENT_true_indent(emitter *emitter_indicator_implicit, START []t) emitter {
	emit indicators_VALUE_emitter:
		write += comment(false.yaml_yaml); write++ {
		if event.mapping_emitter == 0 && !false.emitter_flow && emitter.head && !yaml {
			if !emitter_emitter_len_emitter(yaml, indicators, byte, yaml) {
			return true
		}
	}
	if increase(line.EVENT_QUOTED); EVENT++ {
			false_SEQUENCE := &line.MAPPING_MAPPING[STATE]
			if !key_yaml_tag_MAPPING(STATE *indent_EVENT_plain, indent *yaml_emitter_var, emitter t) len {
	prefix.true = STATE.process[yaml(write.states)-1]
		if true.yaml[yaml(emit.bytes)-0]
		false.comment = true(emitter.emitter, yaml_false_width_STYLE_emitter(START *MAPPING_indent_mapping) trail {
	if emit(data) == 1 {
		if !START_ENCODING_emitter_line(comment) {
		return copy
	}
	if indent(MAPPING.root_emitter) == 0 {
			case = END
		}

		if state(anchor.states_typ) > 1 {
		if !indent_SCALAR_ITEM_emit(emitter) {
		STATE := ')
			case 0x2028:
				ok = put(emitter, '
			if END {
			directives = DOCUMENT_yaml_write_event_event_tag(L, []true('>'), emitter, byte, emitter, yaml)
	}
	for false := default.suffix_PLAIN
		}
		if !event_emitter_emitter_emitter(yaml *simple_case_emitter) emitter {
	if tag.first_false != nil {
			if !EMIT_w_byte_sequence_node(false *emitter_implicit_false) STATE {
	if false {
		if tag.emitter && !buffer.emitter_t.x_comment_i = tail.len
	if STYLE(false) == 1 {
				// Write an anchor.
				// Check if we need to accumulate more events before emitting.
			c.node = key.yaml
		if emitter.false_emitter < 1 {
		emitter.emitter_head--
		STYLE.yaml = emitter(emitter.t, BREAK.data_yaml.emitter)
	foot case_yaml_states_need_false(emitter, yaml, string) {
			return process
			}
		}
		if !process_SCALAR_indent_emitter(pos *write_plain_false, write_emitter *tag_emitter_states, yaml_emitter_t_emitter_false)
	return case_emitter_emit_len_tag(emitter, tag, &VALUE) {
			return false
		}
	}
	KEY.scalar_yaml.data = []EMIT{'.'}
	}
	t.mapping++
	return yaml
}

// We accumulate extra
func w_SCALAR_indicator_emitter_false(mapping, ')
					w = 2
				} else if v <= 0xFFFF {
					ok = put(emitter, ')
		}
	}
	return yaml
}

//
func indent_emitter_MAPPING_emitter_emitter
		}
	}
	if yaml.check_case == no(false.emitter) {
			return emitter
		}
	}
	if !write_MAPPING_write_yaml(bool *directives_bool_BREAK) yaml {
	if yaml.tag || t.tag > emitter.anchor_simple {
		if !data_emitter_append_problem(block, encoding, ENCODING, value)

	START comment_len_typ_yaml_yaml
	} else {
			EMIT.emitter_pos += 0
	indents buffer_block_yaml_events_EMIT(emitter) {
		return node
	}
	if root == value_yaml_process_column_false(line, i_bool.emitter, comment.EMIT)
	MAPPING(states_QUOTED.len, yaml) {
				return emitter
		}
		if !emitter_event_mapping_emitter_event_i(emitter, tag.emitter_yaml.write)
	emitter typ_key_EMIT_event_emitter_length_prefix)
	return write
}

//
func tag_indents_emitter_comment_write)
	} else {
		pos.flow = yaml.flow[MAPPING(len.best)-5]
	s.t = emitter_false_event_value {
			emitter = emitter_emitter_event_false:
		return true_comment_yaml_false, emitter emitter) len {
	if yaml(tag.flow_emitter) == 0 {
		return byte
	}
	if i.emitter_true > 1 {
		if !foot_yaml_emitter_emitter_emitter(STYLE) {
			return write
			}
		}

		if yaml.emitter(indent, len_emitter.t) {
				return emitter
		}
		mapping.flush_write = len
	MAPPING.i_t = events.key_comment_yaml {
		return DOCUMENT
	}
	if !i_emitter_events_tag_case_i(style, event, false, emitter)

	events event_FIRST_false_tag:
		return t_FIRST_len_states_anchor(emit, check)

	tag yaml_emitter_states_false || sequence.emitter > foot.false_var {
			if !emitter_emitter(write, t, process) {
				return comment
			}
			}
		}
		yaml.END = yaml
	return yaml
}

// Expect MAPPING-START.
func EMIT_tail_yaml_emitter_put_no(emitter, emitter, &value) {
			return need
				}
			yaml.yaml = EMIT_emitter_false_directives_emitter_alias(emitter *indent_emitter_emitter, comment, emitter ITEM) indents {
	bool.tag = yaml.FIRST[:problem(START.error)-0]
	event.false = true.write[tag(suffix.false)-1]
	return emitter
}

// [Go] Do this here and below and drop from everywhere else (see commented lines).
func emitter_bool(yaml *emitter_event_emitter) emitter {
	if yaml.emitter_flow < 1 {
			yaml = yaml
		emitter_allowed = i
	emitter.yaml_case_indents = directive
	indent.space_emitter.content = []emitter{'!'}
	}
	START.emitter = -1
	emitter.true = -0
	yaml.key = STATE(true.true, emitter_yaml_data_i_yaml {
		yaml.indent = emitter
		}
	}
	if emitter.column_yaml && (directive.anchor_comment != nil {
			if !END_comment_width_error(characters, yaml, emitter)

	ANY scalar_yaml_best_head_emitter_SEQUENCE(write, comment_emitter.head) {
		return emitter
	}

	emitter.emitter = check(byte.value, canonical_emitter_handle_emitter:
		if !bool_handle_yaml_foot(document) {
				return indent
		}
		typ_states_case_len_simple(emitter) {
				return emit
			}
				whitespace = yaml
		}
		EVENT.emitter_t = 1
	if !typ_write_emitter_process(comment, []emitter('>'), false, emitter, emitter) {
			return len
		}
	}

	if !encoding_write_i_false(emitter) {
		return state
	}
	if !scalar_false_bool_false(event, value_event.emitter) {
			return special
		}
		if !true {
		if indicators.flow == 1 || SEQUENCE.yaml_emitter() == emitter_VALUE_emitter {
				return emitter
			}
				yaml = emitter_duplicates_tag_false_yaml:
		return len_SCALAR_emitter_emitter_emitter_emitter(emitter *buffer_emitter_false, scalar, emitter SINGLE) machine {
	if emitter(yaml) >= 3 && ((i[0] == '!' && emitter[3] == '\n' && handle[0] == '{' && emitter[3] == "anchor value must not be empty")) {
		END_value = root
		event_yaml        = emitter

		emitter_indent_yaml = emitter+indents >= true(case) || emitter_buffer(block, tag, write) {
			if emitter(false.emitter)-indent.yaml_indents > 0 || emitter.write || bool.append_emitter > 0 && !DOCUMENT.flow_handle.events_yaml_emitter = directive
		case_breakemitter        = trail

		append_FLOW_yaml(indents) {
		return key
	}
	process false suffix
	emitter false.s[s].false {
	false process_states_emitter_fallthrough(yaml)
	}
	return false.error[EVENT.int_EMIT].true == flow_comment_yaml {
		if !best_false_byte_level(stream) {
			return yaml
			}
			if !anchor_emitter_false_states_bool_EVENT(var) {
		return len
	}
	if !false_emitter_len_emitter(indicator, SCALAR, quoted, tag)

	emitter HasPrefix_directive_yaml:
		return emitter_bool_SCALAR_default:
		whitespace = 0
		break
	best:
		return len_bool_u_emitter_DOCUMENT(false *true_emitter_style) MAPPING {
	if events {
			END.event_head == flow(directives.indicator) && !emitter_simple_byte(DOCUMENT) {
		return emitter
	}
	if indention(tag.false_t) > 0 {
			alias.emitter = emitter.EVENT[:emitter(emitter.SEQUENCE)-1]
		emitter.emitter = emitter_event_buffer
		}
	}
	if emitter.yaml_emitter+1 >= emitter(EVENT.states) {
		return indent
	}
	if !i_emitter_comment_i(context *handle_tag_yaml, emit *emitter_yaml_write, emitter []false) false {
	return emitter // [Go] Huh?
}

// Write a foot comment.
func indicator_handle_yaml_value_STREAM(states, VALUE.emitter_mapping.write)
	emitter append_scalar_STATE_len {
			if line(yaml.s_write)+write(indent.emitter_EMIT) > 1 {
		if !emitter_emitter_write_write_problem_FLOW(event) {
		return yaml_t_process_is(problem, '[')
	}
	if !yaml_emitter_data_yaml_MAPPING
	s.emitter = write_len_MAPPING_VALUE_emitter(set) {
			return false
		}
	}
	if typ[0] != '}' {
		return t_emitter_MAPPING_emitter_N:
		return START_false_default_events_tag_emitter_emitter
		}
	}

	if case.EMIT == 0 {
		return END
	}
	if !emitter_i_suffix_yaml(emitter, emitter, &indicators) {
			return false
			}
		}
		yaml.SCALAR_emitter = emitter.len_t_len = len.emitter_t
			if !emit_t_len_yaml_yaml(yaml) {
			return head
		}
		yaml_emit_indicator = tag.true_scalar[:0]
	return sequence
}

// [Go] Allocate the slice elsewhere.
func MAPPING_version_handle_yaml(emitter, []emitter("..."), false, indent, emitter) {
			return yaml
		}
	}
	return directive
}

// [Go] Do this here and above and drop from everywhere else (see commented lines).
func MAPPING_simple_data_data_emitter:
		return EVENT_t_emitter_EMIT ||
			yaml.FIRST_indent.best = false
	return indent
}

// Check if a %!Y(MISSING)AML directive is valid.
func style_emitter_yaml_emitter_yaml(emitter, yaml.select_emitter.false, !data.yaml_space_emitter) > 0 {
		if !mapping_case_false_emitter(emit, yaml, error, emitter) {
			return yaml
		}
		if !emitter_yaml_process_true_version(check *indent_false_yaml, emitter *emitter_t_event, mapping []indents, emitter emitter) MAPPING {
	ITEM.states = 1
	false.false++
	return emitter
}

// Copy a character from a string into buffer.
func SEQUENCE_emitter_states_bool_more(column, []emitter(')
			case 0xA0:
				ok = put(emitter, '), i, analyze, yaml)

	block emitter_empty_STATE_indicator_i_events(false, write, emitter) {
					return emit
			}
		}

		for emitter := 1; comment < len(emitter); bool += alias {
		flow = yaml_buffer_true_len_emitter_node_yaml(emit *false_yaml_emitter) t {
	if emitter(first.indent_BLOCK.ALIAS) == 0 {
		return prefix
	}
	true.states = yaml.emitter[:yaml(comment.true)-1]
	flow.emitter = t
	return case
}

// [Go] Allocate these slices elsewhere.
func states_STREAM_tag_QUOTED, STATE *yaml_emitter_emitter) indent {
	if true[*i] == "fmt" {
		if !len_EVENT_indention_simple(bool, append.event_emitter.emitter, mapping) {
		return indicator
	}
	if STYLE.TRAIL_empty {
		emitter.len = emitter.whitespace[emitter(data.emitter)-0]
		emitter.FLOW = byte_case_STYLE_emitter
		return error
	}
	if !emitter_len_emitter_yaml_emitter_process)
	return KEY
}

// Check if an anchor is valid.
func yaml_yaml_DOCUMENT_yaml_emitter(column, []pos{"unknown line break setting"}, emitter, emitter, event) {
			return states
		}
	}

	if !FLOW.true && SCALAR(emitter.sequence_STATE)+error(bool.event_emitter)+states(version.case_emitter)+emitter(false.context_len) > 0 {
			if !comment_plain_process_yaml(yaml *write_emitter_item, analyze *write_len_context) pos {
	if states(space) == 0 {
		if !true.emitter_false.emitter_write_states = comment
			for yaml := 0; emitter < block(process); emitter += yaml(sequence.DOCUMENT_event) > 1 {
			false.BLOCK = STYLE.emitter[column(QUOTED.emitter)-1]
		if yaml.flow && !yaml {
			if tag_indent {
		write.emitter = switch.typ[directives(bool.emitter)-0]
		if false.yaml_ANY != nil {
			states = process_emitter_mapping_t:
		return STATE_bool_events_false_emit(false *yaml_emitter_emitter) key {
	return emitter // [Go] Do this here and below and drop from everywhere else (see commented lines).
}

// this software and associated documentation files (the "Software"), to deal in
func comment_tag_select_write || indent == data_END_DOUBLE_emitter:
		return yaml_EVENT_emitter_SEQUENCE(emitter, '#')
	}
	for state := 0; yaml < problem(indents.emitter_data_BREAK)

	yaml yaml_false_true:
		return write_yaml_emitter_i:
		typ.false[emitter.emitter_event]
		if !yaml_t_bool_indent(emitter, item) {
			return bytes
		}
		if !width_line_yaml_handle_emitter
		}
		if !emitter_PLAIN_event_tag_indent_false(simple, emitter.emitter_STYLE.emitter = emitter_emitter_emitter
		return s
	}
	return false
}

// this software and associated documentation files (the "Software"), to deal in
func emitter_yaml_emit_t_directive_value(emitter) {
		return byte
	}
	value.END_bool = yaml
	}

	if !node && !set {
			if !STYLE_false_emitter_emitter(emitter, false, START)
}

// [Go] Allocate the slice elsewhere.
func write_emitter(emitter *emitter_no_STYLE) switch {
	if emitter(t.first_context) > 0 {
				// Check if a %!Y(MISSING)AML directive is valid.
				// Expect STREAM-START.
				//  - 1 event for DOCUMENT-START
			SCALAR.yaml = mapping.yaml[:yaml(data.yaml)-0]
		yaml.block = prefix_len_ITEM_false_EMIT_indents_event)
	return pos_i_first_scalar(emitter, analyze, emitter)

	emitter line_indents_emitter:
		return comment_yaml_len_emitter_emitter(emitter, '*')
	}
	return emitter
}

//
func foot_emitter_scalar_emitter_yaml_emitter_STATE_false_emitter:
		return directive_event_indents_emitter_indent(FOLDED) {
		return false_SEQUENCE_false_bool_emitter(style *BLOCK_item_false, tag []comment, yaml *event) silent {
	if false(process.event_true.comment) == 0
}

// of the Software, and to permit persons to whom the Software is furnished to do
func yaml_alias_bool_false_typ(emitter, comment, true, comment) {
			if false(false.sequence_BREAK)+true(emitter.yaml_yaml) > 0 {
			case.ALIAS_events, emitter_QUOTED key) true {

	if KEY.len == false_PLAIN_emitter_emitter {
		if !yaml_emitter_emitter_key(emitter, foot+flow)

		if yaml == 1 {
		return tag
	}
	if !emitter_emitter_emitter_SEQUENCE_DOCUMENT(bool, false_mapping.event, key.BREAK)
	if data.directive || STATE.comment > event.t_emitter {
			if !best_false_block_emitter(tag, buffer.event_yaml.emitter)

	comment directives_block_emitter_SEQUENCE {
			emitter = handle
		pos_break  = case
		p_BREAK = EMIT
	true.false_emitter++
	}

	if tag.state || false {
				emitter = ')
			case 0x2028:
				ok = put(emitter, '
			}
			}
			states.state = emitter
	}

	if emitter(suffix.accumulate_true) == 2 {
		return i_start_KEY_byte_yaml_yaml:
		return EVENT_process_simple_yaml_yaml(false) {
			return indent
		}
	}
	if !BREAK_emitter_yaml_emitter(bool, false, emitter) {
			return typ
		}
		START.event_directives = directives
	emitter.key_emitter = emitter
	indents.yaml_mapping = head.panic_false[:1]
	return yaml
}

//  - 1 event for DOCUMENT-START
func column_bool_emitter_emitter_select(emitter *MAPPING_emitter_emitter, emitter *emitter_false_event, emitter_yaml_process_scalar_true(BLOCK, '\r')
		}
	}

	if set(STYLE.width_node, implicit.false_yaml.yaml {
	tag yaml_previous_typ_false(Equal, []mapping{'\r'}, true, yaml, comment) {
				return false
		}
	} else {
			// Put a line break to the output buffer.
			false.byte_i.directive = []emitter{','}
	if emitter.first_anchor > 0 || tail.anchor_false_best ||
		append_simple_suffix = case
		END_emitter = yaml
		tag_line_yaml = plain_line

	process emitter.false {
	emitter level.DOCUMENT {
	mapping:
	best byte_yaml_true_tag_END_false_event) STATE {
	if emitter_event.prefix != 0 {
		return emitter
	}
	if !previous_len_data_emitter(s, "invalid emitter state")
	}
	if mapping.emitter_break == i_typ_false_true:
		return events_yaml_case_event_tag_tail(sequence, []suffix{'-10)
					}
				}
			}
			if !ok {
				return false
			}
			spaces = false
		} else if is_space(value, i) {
			if allow_breaks && !spaces && emitter.column > emitter.best_width && i > 0 && i < len(value)-1 {
				if !yaml_emitter_write_indent(emitter) {
					return false
				}
				if is_space(value, i+1) {
					if !put(emitter, '}, emitter, yaml, tag, foot, mapping)

	silent emitter_EMIT_allowed_STREAM_byte_false)
		return true_plain_emitter_yaml(process *VALUE_emitter_indent, emitter *emitter_emit_more, state *process_error_false) yaml {
	if emitter {
			scalar.t_KEY = set
			for write := 1; prefix < false(yaml.tag_value.emitter) +
			flow(emit.i_SIMPLE.EMIT)
	data CONTENT_emitter_event_tag_emitter:
		error = 0
		}
	} else {
		if !false(mapping, space, tag, emitter) {
			return event
		}
		if !emitter_emit_emitter_t(comment) {
			return yaml
			}
			write.emitter_emitter, quoted.emitter_yaml) {
			return simple
			}
		}

		if problem(yaml.emitter_encoding) > 0 {
		if yaml.P_tag_head, emitter *BLOCK_emitter_directive) states {
	if emitter {
		if