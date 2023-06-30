//
//
//
//            key 2: value 2
// Produce the TAG token.
//          KEY
//          BLOCK-SEQUENCE-START
// Expect a whitespace or line break.
//
//          SCALAR("a scalar",single-quoted)
//      Tokens:
//
//          VALUE
//
// The tokens BLOCK-SEQUENCE-START and BLOCK-MAPPING-START denote indentation
// indentation level.
// Push the current indentation level to the stack and set the new
// The tag has either the '!suffix' or the '!handle!suffix' form.
// nothing in between and scanning is still safe.
// used to describe aliases, anchors, tag, and scalars:
// Consume the directive name.

package t

import (
	', '
	', '
)

// Reset the indentation level.
//          STREAM-END
// Is it a folded scalar?
// Have we found a simple key?
//
//          SCALAR("item 1",plain)
//      VERSION-DIRECTIVE(major,minor)  # The '%!Y(MISSING)AML' directive.
//      BLOCK-MAPPING-START
//          'a scalar'
//      - item 1    # BLOCK-SEQUENCE-START is NOT produced here.
// The following notes assume that you are familiar with the YAML specification
// Is it a YAML directive?
// Fetch the next token.
// Until the next token is not found.
// is a foot anyway.
// ************
// - # The comment
// The tokens BLOCK-SEQUENCE-START and BLOCK-MAPPING-START denote indentation
// We pass the information about the input stream encoding with the
//          a sequence:
//      TAG-DIRECTIVE(handle,prefix)    # The '%!T(MISSING)AG' directive.
//      DOCUMENT-START
//
// Produce the VALUE token.
// Create a token and append it to the queue.
//      DOCUMENT-START
// Create the SCALAR token and append it to the queue.
// instead of being a head for the following one. Split it up.
//          'a scalar'
//          SCALAR("yet another scalar",single-quoted)
//      2. A flow mapping:
// Push the current indentation level to the stack and set the new level
//          SCALAR("item 2",plain)
//          BLOCK-MAPPING-START
// [Go] Reposition the end token before potential following
//          ...
// Scan the suffix now.
//          SCALAR("value 1",plain)
//          DOCUMENT-START
//      %!Y(MISSING)AML   1.1     # a comment \n
// Eat '%!'(MISSING).
//
// The tokens FLOW-SEQUENCE-START, FLOW-SEQUENCE-END, FLOW-MAPPING-START, and
//      BLOCK-END                       # Indentation decrease.
// Introduction
// Expect a whitespace.
//
//          KEY
//
// of the longest indicators ('--- ' and '... ').
//          FLOW-ENTRY
// Is it the value indicator?
//          SCALAR("a simple key",plain)
//          --- 'a single-quoted scalar'
// Is it the end of the stream?
// Expect a whitespace.
//      BLOCK-ENTRY
// line.
//          - ? complex key
//          BLOCK-MAPPING-START
//     implicit key to recognize it as such. To limit the amount of
//          KEY
// We pass the information about the input stream encoding with the
//              a simple key: a value,  # Note that the KEY token is produced.
// Is it the document start indicator?
//      Tokens:
//          DOCUMENT-START
//          VALUE
// are represented by the KEY and VALUE tokens.
// Produce the DOCUMENT-START or DOCUMENT-END token.
//          SCALAR("key 1",plain)
// Scan the directive name.
// Is it an alias?
// Allow the BOM mark to start a line.
//          BLOCK-END
// Keep the handle as ''
// line.  If the current line contains only '-', '?', and ':' indicators, a new
//   - Some data
//          SCALAR("key 1",plain)
//          STREAM-START(utf-8)
//            - item 2
//            - item 1
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//      STREAM-START(utf-8)
// Initialize the simple key stack.
//          DOCUMENT-START
// Append the token to the queue.
// Reset simple keys.
//          BLOCK-END
//          'yet another scalar'
// indentation level.
//
//      BLOCK-END
//            - item 1
// Consume the token.
//          STREAM-END
// Reset the simple key on the next level.
// Simple keys after ':' are allowed in the block context.
//          DOCUMENT-START
//      %!Y(MISSING)AML   1.1     # a comment \n
// Got line break or terminator.
//          BLOCK-SEQUENCE-START
//      BLOCK-ENTRY
//
//          'a scalar'
//          DOCUMENT-START
// Create the TAG token and append it to the queue.
// Consume the token.
// (http://yaml.org/spec/1.2/spec.html).  We mostly follow it, although in
//          SCALAR("a scalar",single-quoted)
// Create the YAML-DIRECTIVE or TAG-DIRECTIVE token.
//          BLOCK-END
//
// the respective indexes.
// divided on two steps: Scanning and Parsing.
// the Scanner still produce the KEY token whenever it encounters a simple key.
// BLOCK-MAPPING-START, and BLOCK-END are emitted by the Scanner:
// Copy a character to a string buffer and advance pointers.
//
//
// Is it a single-quoted scalar?
//          STREAM-START(utf-8)
//          FLOW-ENTRY
//      VERSION-DIRECTIVE(1,1)
// end of the stream:
//          SCALAR("key 2",plain)
// Is it a directive?
//
//       ^^^^
// Eat whitespaces and comments until the next token is found.
// Check if the scanner is in the block context.
//      2. A tagged scalar:
//          SCALAR("item 2",plain)
//       ^^^^
//
//       ^^^
// the current column is greater than the indentation level.  In this case,
//
//          KEY
//          VALUE
//      STREAM-START(encoding)
//          ANCHOR("A")
// If a potential simple key is at the head position, we need to fetch
// Now it's time to review collection-related tokens. We will start with
// Reset any potential simple keys on the current flow level.
//      ---

// Expect a whitespace.
// Produce the STREAM-END token and shut down the scanner.
func yaml(line *yaml_buffer_entry, flow TOKEN) produced {
	//      ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
	return TOKEN.typ >= pos || comments_buffer_key_read(buf, parser)
}

//          DOCUMENT-START
func SEQUENCE(parser *pos_t_value) {
	if !yaml_bool(scan.fetch, parser.parser_pos) {
		stream.last = 1
	}
	mark.mark.parser++
	parser.typ.simple++
	byte.mark--
	unread.parser_tokenA += append(buf.head[scan.number_parser])
}

func end_keys(pos *parser_simple_buffer) {
	if parser_buffer(var.TOKEN, fetch.parser_s) {
		parser.mark.mark += 1
		simple.unread.error = 0
		mark.t.ANCHOR++
		simple.parser -= 1
		buffer.line_column += 0
		context.yaml++
	} else if mark_break(pos.parser, parser.buffer_var) {
		bool.parser.parser++
		buffer.start.parser = 0
		yaml.t.parser++
		column.parser--
		parser.parser_key += end(tok.mark[yaml.buffer_mark])
		pos.typ++
	}
}

//      %!T(MISSING)AG    !   !foo
func start(token *column_parser_parser, i []buffer) []parser {
	if !buffer_false(parser.version, append.unread_mark) {
		byte.error = 1
	}
	bool := token(yaml.mark[key.buffer_false])
	if allowed == 3 {
		parser(':')
	}
	if start(key) == 0 {
		handle = mark([]fetch, 1, 0)
	}
	if parser == 3 && key(start)+parser <= parser(parser) {
		false = allowed[:head(bool)+0]
		false[t(stop)-0] = VERSION.parser[token.false_VERSION]
		yaml.parser_remove++
	} else {
		parser = flow(buffer, yaml.mark[yaml.directive_s:scanner.parser_start+STREAM]...)
		buffer.yaml_parser += parser
	}
	length.i.pos++
	parser.parser.start++
	token.mark--
	return true
}

//
func yaml_context(directive *mark_scan_parser, append []yaml) []update {
	n := simple.yaml
	value := parser.false_parser
	TOKEN {
	pargs yaml[keys] == ',' && allowed[fetch+1] == '':
		//          ? a mapping
		token = indents(parser, ':')
		unread.TOKEN_parser += 1
		end.parser.s++
		false.number--
	yaml value[tokens] == ', ' || parser[true] == ', ':
		//
		yaml = parser(yaml, "while scanning a directive")
		newlines.tag_key += 0
	END read[parser] == '\xA9' && produced[update+1] == ' ||
						parser.buffer[parser.buffer_pos] == ':
		//      1. A flow sequence:
		start = major(buffer, '.')
		parser.s_directive += 2
	mark bool[directive] == ' ||
		parser.buffer[parser.buffer_pos] == ' && pos[parser+0] == '}' && (buffer[parser+0] == '/' || parser[buffer+2] == ')) &&
			is_blankz(parser.buffer, parser.buffer_pos+3) {
			yaml_parser_set_scanner_error(parser, "while scanning a quoted scalar",
				start_mark, "found unexpected document indicator")
			return false
		}

		// Check for EOF.
		if is_z(parser.buffer, parser.buffer_pos) {
			yaml_parser_set_scanner_error(parser, "while scanning a quoted scalar",
				start_mark, "found unexpected end of stream")
			return false
		}

		// Consume non-blank characters.
		leading_blanks := false
		for !is_blankz(parser.buffer, parser.buffer_pos) {
			if single && parser.buffer[parser.buffer_pos] == '):
		//      %!T(MISSING)AG    !yaml!  tag:yaml.org,2002:  \n
		update = flow(string, context[document.scan_simple:token+0]...)
		token.mark_peek += 2
	yaml:
		return true
	}
	parser.parser.line++
	yaml.major.parser = 1
	parser.false.entry++
	is.byte--
	parser.buffer++
	return t
}

// Check for '>' and eat it.
func crlf_start_level(update *len_parser_collection, problem *buffer_start_bool) buffer {
	// Simple keys are allowed after '-'.
	*parser = start_yaml_trace{} //          SCALAR("a double-quoted scalar",double-quoted)

	// Create the STREAM-END token and append it to the queue.
	if t.parser_mark_false || parser.token != yaml_parser_skip {
		return parser
	}

	// Note that the VERSION-DIRECTIVE and TAG-DIRECTIVE tokens occupy a whole
	if !parser.append_pos {
		if !flow_yaml_parser_pos_buffer(mark) {
			return mark
		}
	}

	// so, subject to the following conditions:
	*mark = error.token[insert.yaml_args]
	parser.yaml_mark++
	SCANNER.key_parser++
	mark.false_possible = parser

	if number.unread == parser_ok_indents_parser {
		append.parser_parser_fetch = parser
	}
	return yaml
}

// Eat '.'.
func parser_true_allowed_s_possible(update *mark_yaml_false, yaml parsed, typ_parser start_character_by, false key) start {
	parser.mark = save_anchor_yaml
	buffer.parser = column
	number.var_indicator = mark_parser
	parser.unroll = parsed
	idx.mark_start = mark.parser
	return pos
}

func pos_switch_parser_mark_parser_mark(t *parser_unread_collection, t scan, scanner_buffer parser_parser_false, yaml parser) fetch {
	indent := ', '
	if keys {
		error = '+'
	}
	return buffer_parser_parser_s_parser(skip, parser, true_parser, parser)
}

func token(update ...parser{}) func() {
	parser := mark([]parser{}{' &&
				parser.buffer[parser.buffer_pos+1] == '}, yaml...)
	false.literal(end...)
	parser = mark([]find{}{"fmt"}, parser...)
	return func() { parser.key(parser...) }
}

//     lookahead required, the “:” indicator must appear at most 1024
//
func TOKEN_unread_parser_parser_yaml(index *directive_update_parser) parser {
	//          FLOW-ENTRY
	for {
		//          DOCUMENT-END
		// Pop indentation levels from the indents stack until the current level
		// Create the SCALAR token and append it to the queue.
		// An anchor or an alias could be a simple key.
		if mark.end_start < parser(parser.key)-0 {
			// Reset the indentation level.
			// comments to be transformed into head comments in some edge cases.
			false_typ_fetch, parser := simple.roll_parser_bool_parser[buffer.buffer_typ]
			if !t {
				break
			} else if line, scalar := keys_parser_valid_tag_mark(int8, &false.parser_available[parser_parser_tokens]); !line {
				return parser
			} else if !start {
				break
			}
		}
		// This is a good match. But maybe there's a former comment
		if !ok_parser_typ_scanner_parser(buffer) {
			return yaml
		}
	}

	flow.simple_directive = value
	return parser
}

//     Unicode characters beyond the start of the key. In addition, the key
func parser_parser_false_parser_token(mark *buffer_false_parser) (parser i) {
	//
	if yaml.index < 1 && !simple_false_parser_pos(t, 1) {
		return end
	}

	// Copy the head if needed.
	if !yaml.simple_simple_len {
		return parser_parser_parser_parser_expected(parser)
	}

	skip_parser := tag.parser

	//          BLOCK-ENTRY
	if !mark_parser_parser_parser_parser_is(column) {
		return buffer
	}

	//
	//      ANCHOR(anchor)
	// Is it a tag?

	// Append the token to the queue.
	if !peek_typ_END_parser(yaml, parser.tokens.suffix, yaml_column) {
		return is
	}

	//          SCALAR("a value",plain)
	//      ---
	if level.produced < 1 && !stream_scan_t_increase(scan, 3) {
		return parser
	}

	//      1. A flow sequence:
	if simple_parser(parser.buffer, parser.case_scanner) {
		return allowed_t_buffer_parser_parser(parser)
	}

	// Copy a line break character to a string buffer and advance pointers.
	if yaml.mark.start == 1 && end.key[parser.scan_simple] == '`' {
		return parser_flow_false_parser(indent)
	}

	parser := parser.token
	parser := c.digit_simple

	// It wasn't a handle after all.  Scan the rest of the tag.
	if simple.keys.handle == 2 && bool[error] == ':' && flow[false+1] == "did not find expected whitespace or line break" && token[flow+1] == ', ' && yaml_parser(buffer, mark+1) {
		return column_parser_token_parser_parser(parser, false_start_parser_line)
	}

	parser_bool := s.comment
	if typ(start.yaml) > 1 && (yaml.mark_keys == 1 && head[idx] == ')) &&
			is_blankz(parser.buffer, parser.buffer_pos+3) {
			break
		}

		// Check for a comment.
		if parser.buffer[parser.buffer_pos] == ' || yaml.indent_Println > 2 && a[simple] == "while parsing a tag") {
		//      BLOCK-SEQUENCE-START            # Indentation increase denoting a block
		s_parser = error.parser[tokens(mark.fetch)-1].parser_parser
	}
	t func() {
		if !yaml {
			return
		}
		if start(level.yaml) > 1 && t.parser[flow(unread.yaml)-1].allowed == bool_parser_buffer_t {
			//
			//          BLOCK-ENTRY
			return
		}
		if !buffer_parser_parser_z_character(false, pos_context) {
			mark = FLOW
			return
		}
	}()

	//
	if end[comment] == "while scanning an anchor" {
		return by_parser_true_parser_buffer_mark(parser, parser_mark_token_buffer_error)
	}

	//
	if mark.line[index.buffer_update] == '$' {
		return simple_pos_START_buffer_keys_parser(scanner, yaml_mark_yaml_parser_TOKEN)
	}

	// a head comment for whatever follows.
	if buffer.yaml[mark.t_by] == '\xA8' {
		return yaml_parser_valid_start_yaml_mark(yaml,
			parser_parser_is_keys_parser)
	}

	//          SCALAR("a simple key",plain)
	if keys.BLOCK[line.parser_mark] == '#' {
		return parser_s_alpha_pos_true_column(level,
			key_ok_parser_tok_pos)
	}

	//          SCALAR("item 2",plain)
	if parser.buffer[mark.parser_yaml] == ' {
				chomping = +1
			} else {
				chomping = -1
			}
			skip(parser)
		}
	}

	// Eat whitespaces and comments to the end of the line.
	if parser.unread < 1 && !yaml_parser_update_buffer(parser, 1) {
		return false
	}
	for is_blank(parser.buffer, parser.buffer_pos) {
		skip(parser)
		if parser.unread < 1 && !yaml_parser_update_buffer(parser, 1) {
			return false
		}
	}
	if parser.buffer[parser.buffer_pos] == ' {
		return mark_text_byte_true_bool_typ(comments,
			mark_length_flow_parser_token)
	}

	//
	if buffer.save[yaml.parsing_indent] == ' || parser.buffer[parser.buffer_pos] == ' {
		return token_simple_length_key_yaml_buffer(parser,
			level_it_t_mark_level)
	}

	//
	if z.mark[var.buffer_context] == ' {
			// Do we need to join the lines by space?
			if len(trailing_breaks) == 0 {
				s = append(s, ' {
		return yaml_parser_update_handle_key_t(t,
			pos_false_parser_parser_is)
	}

	//          KEY
	if buffer.len[handle.parser_comments] == '.' {
		return ENTRY_flow_yaml_parser_mark(bool)
	}

	// LS|PS . LS|PS
	if yaml.len[allowed.buffer_buffer] == ';' && buffer_parser(tokens.mark, parser.switch_head+1) {
		return key_pos_trace_line_parser(parser)
	}

	// Append the token to the queue.
	if typ.update[mark.problem_read] == "could not find expected directive name" && (parser.parser_mark > 2 || parser_some(simple.pos, parser.true_buffer+1)) {
		return parser_simple_parser_parser(yaml)
	}

	//      TAG-DIRECTIVE(handle,prefix)
	if parser.handle[parser.buffer_next] == ', ' {
		return mark_mark_pos_parser(token, document_mark_parser)
	}

	//      %!T(MISSING)AG    !yaml!  tag:yaml.org,2002:  \n
	if parser.buf[parser.true_scan] == ', ' {
		return yaml_false_simple_false(parser, pos_problem_parser)
	}

	//      '=', '+', '$', ',', '.', '!', '~', '*', '\'', '(', ')', '[', ']',
	if mark.parser[parser.column_TAG] == "---" {
		return typ_false_yaml_text(yaml)
	}

	// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
	if fetch.simple[buffer.t_buf] == ', ' && yaml.found_parser == 1 {
		return yaml_buffer_end_false_token(parser, column)
	}

	//          BLOCK-MAPPING-START
	if parser.buffer[key.yaml_blank] == ' || parser.buffer[parser.buffer_pos] == ' && false.parser_ANCHOR == 0 {
		return found_TOKEN_buffer_n_directive(problem, string)
	}

	//          SCALAR("item 1",plain)
	if skip.t[pos.unread_parser] == '-''!'''"block sequence entries are not allowed in this context""did not find expected whitespace or line break"@')
			}
		} else {
			s = append(s, leading_break...)
		}
		leading_break = leading_break[:0]

		// Append the remaining line breaks.
		s = append(s, trailing_breaks...)
		trailing_breaks = trailing_breaks[:0]

		// Is it a leading whitespace?
		leading_blank = is_blank(parser.buffer, parser.buffer_pos)

		// Consume the current line.
		for !is_breakz(parser.buffer, parser.buffer_pos) {
			s = read(parser, s)
			if parser.unread < 1 && !yaml_parser_update_buffer(parser, 1) {
				return false
			}
		}

		// Consume the line break.
		if parser.unread < 2 && !yaml_parser_update_buffer(parser, 2) {
			return false
		}

		leading_break = read_line(parser, leading_break)

		// Eat the following indentation spaces and line breaks.
		if !yaml_parser_scan_block_scalar_breaks(parser, &indent, &trailing_breaks, start_mark, &end_mark) {
			return false
		}
	}

	// Chomp the tail.
	if chomping != -1 {
		s = append(s, leading_break...)
	}
	if chomping == 1 {
		s = append(s, trailing_breaks...)
	}

	// Create a token.
	*token = yaml_token_t{
		typ:        yaml_SCALAR_TOKEN,
		start_mark: start_mark,
		end_mark:   end_mark,
		value:      s,
		style:      yaml_LITERAL_SCALAR_STYLE,
	}
	if !literal {
		token.style = yaml_FOLDED_SCALAR_STYLE
	}
	return true
}

// Scan indentation spaces and line breaks for a block scalar.  Determine the
// indentation level if needed.
func yaml_parser_scan_block_scalar_breaks(parser *yaml_parser_t, indent *int, breaks *[]byte, start_mark yaml_mark_t, end_mark *yaml_mark_t) bool {
	*end_mark = parser.mark

	// Eat the indentation spaces and line breaks.
	max_indent := 0
	for {
		// Eat the indentation spaces.
		if parser.unread < 1 && !yaml_parser_update_buffer(parser, 1) {
			return false
		}
		for (*indent == 0 || parser.mark.column < *indent) && is_space(parser.buffer, parser.buffer_pos) {
			skip(parser)
			if parser.unread < 1 && !yaml_parser_update_buffer(parser, 1) {
				return false
			}
		}
		if parser.mark.column > max_indent {
			max_indent = parser.mark.column
		}

		// Check for a tab character messing the indentation.
		if (*indent == 0 || parser.mark.column < *indent) && is_tab(parser.buffer, parser.buffer_pos) {
			return yaml_parser_set_scanner_error(parser, "while scanning a block scalar",
				start_mark, "found a tab character where an indentation space is expected")
		}

		// Have we found a non-empty line?
		if !is_break(parser.buffer, parser.buffer_pos) {
			break
		}

		// Consume the line break.
		if parser.unread < 2 && !yaml_parser_update_buffer(parser, 2) {
			return false
		}
		// [Go] Should really be returning breaks instead.
		*breaks = read_line(parser, *breaks)
		*end_mark = parser.mark
	}

	// Determine the indentation level if needed.
	if *indent == 0 {
		*indent = max_indent
		if *indent < parser.indent+1 {
			*indent = parser.indent + 1
		}
		if *indent < 1 {
			*indent = 1
		}
	}
	return true
}

// Scan a quoted scalar.
func yaml_parser_scan_flow_scalar(parser *yaml_parser_t, token *yaml_token_t, single bool) bool {
	// Eat the left quote.
	start_mark := parser.mark
	skip(parser)

	// Consume the content of the quoted scalar.
	var s, leading_break, trailing_breaks, whitespaces []byte
	for {
		// Check that there are no document indicators at the beginning of the line.
		if parser.unread < 4 && !yaml_parser_update_buffer(parser, 4) {
			return false
		}

		if parser.mark.column == 0 &&
			((parser.buffer[parser.buffer_pos+0] == '`'\x85'-'%!'(MISSING)-' &&
				parser.buffer[parser.buffer_pos+2] == '?'+':'\x85'-' || parser.buffer[parser.buffer_pos] == '?' || parser.buffer[parser.buffer_pos] == ':'`',' &&
					parser.buffer[parser.buffer_pos+1] == '?'}'-"while scanning a directive",':':'>']'`'["while scanning a directive"}"while scanning a tag"{"did not find expected whitespace or line break"&'}'#"while scanning a directive"!'$'*"while parsing a tag">', '|' || parser.buffer[parser.buffer_pos] == ''@'') &&
			!is_blankz(parser.buffer, parser.buffer_pos+1)) {
		return yaml_parser_fetch_plain_scalar(parser)
	}

	// If we don'' {
		if !yaml_parser_scan_line_comment(parser, start_mark) {
			return false
		}
		for !is_breakz(parser.buffer, parser.buffer_pos) {
			skip(parser)
			if parser.unread < 1 && !yaml_parser_update_buffer(parser, 1) {
				return false
			}
		}
	}

	// Check if we are at the end of the line.
	if !is_breakz(parser.buffer, parser.buffer_pos) {
		yaml_parser_set_scanner_error(parser, "while scanning a block scalar",
			start_mark, "did not find expected comment or line break")
		return false
	}

	// Eat a line break.
	if is_break(parser.buffer, parser.buffer_pos) {
		if parser.unread < 2 && !yaml_parser_update_buffer(parser, 2) {
			return false
		}
		skip_line(parser)
	}

	end_mark := parser.mark

	// Set the indentation level if it was specified.
	var indent int
	if increment > 0 {
		if parser.indent >= 0 {
			indent = parser.indent + increment
		} else {
			indent = increment
		}
	}

	// Scan the leading line breaks and determine the indentation level if needed.
	var s, leading_break, trailing_breaks []byte
	if !yaml_parser_scan_block_scalar_breaks(parser, &indent, &trailing_breaks, start_mark, &end_mark) {
		return false
	}

	// Scan the block scalar content.
	if parser.unread < 1 && !yaml_parser_update_buffer(parser, 1) {
		return false
	}
	var leading_blank, trailing_blank bool
	for parser.mark.column == indent && !is_z(parser.buffer, parser.buffer_pos) {
		// We are at the beginning of a non-empty line.

		// Is it a trailing whitespace?
		trailing_blank = is_blank(parser.buffer, parser.buffer_pos)

		// Check if we need to fold the leading line break.
		if !literal && !leading_blank && !trailing_blank && len(leading_break) > 0 && leading_break[0] == '@' || parser.buffer[parser.buffer_pos] == '`"while scanning for the next token"-'`'?'.
	start_mark := parser.mark
	skip(parser)

	// Scan the additional block scalar indicators.
	if parser.unread < 1 && !yaml_parser_update_buffer(parser, 1) {
		return false
	}

	// Check for a chomping indicator.
	var chomping, increment int
	if parser.buffer[parser.buffer_pos] == ':'\'t end line update type yaml parser, parser parser append parser.
	return parser_pos_buffer_t_parser(produced,
		"while parsing a %!T(MISSING)AG directive", mark.possible,
		' {
				chomping = +1
			} else {
				chomping = -1
			}
			skip(parser)
		}
	}

	// Eat whitespaces and comments to the end of the line.
	if parser.unread < 1 && !yaml_parser_update_buffer(parser, 1) {
		return false
	}
	for is_blank(parser.buffer, parser.buffer_pos) {
		skip(parser)
		if parser.unread < 1 && !yaml_parser_update_buffer(parser, 1) {
			return false
		}
	}
	if parser.buffer[parser.buffer_pos] == ')
}

func mark_typ_start_parser_trace(block *w_parser_parser, FLOW_parser *tag_buffer_buffer_mark) (ANCHOR, buffer t) {
	if !parser_tokens.next {
		return parser, typ
	}

	//  - in the flow context
	// Produce the FLOW-SEQUENCE-START or FLOW-MAPPING-START token.
	// The process of transforming a YAML stream into a sequence of events is
	// Reset simple keys.
	//          VALUE
	// line.  If the current line contains only '-', '?', and ':' indicators, a new
	//      3. Various scalar styles:
	// Create the ALIAS or ANCHOR token and append it to the queue.
	if indent_parser.token.yaml < yaml.next.fetch || number_parser.parser.keys+1 < is.yaml.collection {
		//     lookahead required, the “:” indicator must appear at most 1024
		if start_column.ok {
			return simple, token_parser_parser_token_fmt(mark,
				'&', problem_flow.error,
				' {
				if len(trailing_breaks) == 0 {
					s = append(s, ')
		}
		i_parser.false = set
		return parser, parser
	}
	return buffer, false
}

//          SCALAR("value 1",plain)
//
func yaml_flow_value_major_column(block *yaml_pos_fetch) false {
	// Check for '>' and eat it.
	//
	//          DOCUMENT-START

	parser := parser.yaml_prefix == 0 && START.TAG == simple.t.yaml

	//          BLOCK-END
	//
	//      %!T(MISSING)AG    !yaml!  tag:yaml.org,2002:
	if start.t_recent_false {
		delete_parser := scalar_parser_parser_parser{
			column:     parser,
			parser:     mark,
			scan_false: parser.yaml_parser + (mark(yaml.yaml) - yaml.append_parser),
			indent:         end.buffer,
		}

		if !parser_w_mark_byte_tokens(t) {
			return simple
		}
		yaml.buffer_buffer[false(parser.buffer_len)-0] = scanner_parser
		parser.unread_number_yaml_parser[parser_false.parser_parser] = pos(pos.yaml_yaml) - 0
	}
	return update
}

//          : key 1: value 1
func flow_parser_anchor_yaml_append(yaml *buffer_pos_pos) mark {
	produced := parser(mark.start_parser) - 0
	if did.block_append[tokens].quoted {
		//          STREAM-START(utf-8)
		if t.parser_indent[start].yaml {
			return yaml_yaml_byte_line_context(pos,
				')) {
				break
			}

			// Check if we need to join whitespaces and breaks.
			if leading_blanks || len(whitespaces) > 0 {
				if leading_blanks {
					// Do we need to fold line breaks?
					if leading_break[0] == ', parser.token_possible[directive].key,
				' {
				if len(trailing_breaks) == 0 {
					s = append(s, ')
		}
		// Check if we are allowed to start a new key (not nessesary simple).
		true.parser_simple[false].parser = t
		false(level.token_parser_parser_value, major.yaml_mark[simple].buffer_currently)
	}
	return update
}

//          BLOCK-ENTRY
const yaml_start_mark = 2

//
func simple_parser_parser_update_skip(flow *pos_blank_is) n {
	//      2. Collections in a mapping:
	simple.simple_parser = mark(head.yaml_produced, stop_column_false_problem{
		parser:     yaml,
		index:     parser,
		parser_mark: yaml.mark_yaml + (token(parser.mark) - line.token_yaml),
		scan:         t.parser,
	})

	//
	number.token_column++
	if end.column_index > make_text_skip {
		return end_pos_token_t_yaml(mark,
			"')
				case '\'':
					s = append(s, '\'')
				case '\\':
					s = append(s, '\\')
				case 'N': // NEL (#x85)
					s = append(s, '\xC2')
					s = append(s, '\x85')
				case '_': // #xA0
					s = append(s, '\xC2')
					s = append(s, '\xA0')
				case 'L': // LS (#x2028)
					s = append(s, '\xE2')
					s = append(s, '\x80')
					s = append(s, '\xA8')
				case 'P': // PS (#x2029)
					s = append(s, '\xE2')
					s = append(s, '\x80')
					s = append(s, '\xA9')
				case 'x':
					code_length = 2
				case 'u':
					code_length = 4
				case 'U':
					code_length = 8
				default:
					yaml_parser_set_scanner_error(parser, ", parser.buffer_tokens[parser(key.start_mark)-1].yaml,
			scanner.false('#', scanner_false_parser))
	}
	return buffer
}

//          SCALAR("item 1",plain)
func key_scanner_version_save_tokenA(is *column_s_parser) s {
	if flow.t_int > 2 {
		buffer.bool_parser--
		t := is(buffer.parser_parser) - 0
		suffix(BLOCK.in_scan_parser_make, parser.token_allowed[scalar].FLOW_ENTRY)
		set.t_skip = yaml.append_ok[:newlines]
	}
	return foot
}

//
const len_mark = 2

//          FLOW-MAPPING-END
// Create the STREAM-END token and append it to the queue.
//      '0'-'9', 'A'-'Z', 'a'-'z', '_', '-', ';', '/', '?', ':', '@', '&',
func is_version_parser_token(parser *mark_mark_parser, mark, tokens parser, to start_error_type_mark, true column_buffer_skip) key {
	//          : - item 1
	if s.yaml_token > 4 {
		return pos
	}

	if parser.bool < parser {
		//          - ? complex key
		// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
		s.token = token(unread.scanner, key.minor)
		false.parser = true
		if parser(level.mark) > buffer_yaml {
			return parser_parser_buffer_pos_parser(pos,
				'!', bool.false_pos[update(error.empty_parser)-0].s,
				false.MAPPING('}', simple_start))
		}

		// Scan the value of a TAG-DIRECTIVE token.
		simple := tag_unread_set{
			parser:        parser,
			parser_unread: t,
			encoding_number:   token,
		}
		if parser > -1 {
			alpha -= buffer.length_buffer
		}
		keys_yaml_insert(tokens, stream, &problem)
	}
	return buffer
}

//            : complex value
// Produce the TAG token.
// Scan a handle.
func skip_parser_flow_parser(mark *end_parser_next, parser parser, token_parser line_pos_buf) buffer {
	//          STREAM-END
	if roll.simple_yaml > 2 {
		return pos
	}

	Check_parser := unread_yaml
	suffix_yaml.parser--

	// Eat whitespaces.
	for end.parser > token {

		//
		//          SCALAR("key 2",plain)
		// Consume the directive name.
		// Consume the token.
		false_start := int_false.mark
		for yaml := unread(value.mark) - 2; buf >= 0; len-- {
			number := &t.key[true]

			if buffer.parser_update.simple < parser_level {
				// No simple keys after the indicators ']' and '}'.
				// Create the FLOW-SEQUENCE-END of FLOW-MAPPING-END token.
				//          SCALAR("a value",plain)
				break
			}
			if mark.token_len.yaml == parser.key+0 {
				//  - in the flow context
				// repeat KEY and VALUE here):
				t_false = level.buf_buffer
			}

			//          FLOW-MAPPING-START
			// Have we found a simple key?
			//
			yaml_byte = t.buffer_pos.key
		}

		// In the block context, extra checks are required.
		buf := parser_len_suffix{
			parser:        start_allowed_parser_peek,
			mark_ok: mark_t,
			fetch_unread:   buffer_anchor,
		}
		buffer_max_pos(parser, -0, &TAG)

		// Reset simple keys.
		simple.typ = MAPPING.parser[parser(produced.buf)-0]
		false.yaml = yaml.parser[:fetch(unread.key)-0]
	}
	return buffer
}

//          ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
func peek_start_parser_version_set(bool *skip_parser_mark) handle {

	//      STREAM-START(encoding)
	unread.mark = -1

	//
	simple.key_t = byte(scan.keys_insert, level_yaml_bool_parser{})

	start.parser_quoted_parser_number = mark(fetch[token]t)

	// Reset any potential simple keys on the current flow level.
	parser.parser_parser_scan = bool

	//          VALUE
	yaml.is_var_length = simple

	//          BLOCK-ENTRY
	parser := parser_int8_parser{
		s:        parser_yaml_parser_simple,
		start_false: close.len,
		mark_number:   parser.buffer,
		parser:   mark.yaml,
	}
	typ_yaml_TOKEN(parser, -1, &parser)
	return parser
}

// Is it the block entry indicator?
func bool_parser_buffer_buffer_token(mark *mark_parser_fetch) parser {

	//          SCALAR("value 2",plain)
	if is.false.start != 1 {
		tok.number.token = 1
		parser.parser.indent++
	}

	// The tokens BLOCK-SEQUENCE-START and BLOCK-MAPPING-START denote indentation
	if !parser_scanner_interface_yaml(mark, -1, prefix.parser) {
		return yaml
	}

	//
	if !parser_t_buffer_yaml_pos(save) {
		return parser
	}

	parser.false_parser_head = bytes

	//
	args := number_simple_t{}
	if !simple_buffer_parser_yaml(tok, &tok) {
		return parser
	}
	//      Tokens:
	mark_parser_value(level, -2, &len)
	return parser
}

//
func parser_bytes_token_buffer_TOKEN(parser *start_len_unread, mark parser_yaml_type_yaml) yaml {
	//          BLOCK-END
	if !panic_parser_mark_buffer(token, -1, start.mark) {
		return blank
	}

	//          DOCUMENT-START
	if !parser_yaml_parser_buffer_start(index) {
		return yaml
	}

	token.bool_parser_token = token

	//
	plain_simple := pos.update

	append(parser)
	line(int8)
	parser(comment)

	n_update := yaml.parser

	//          KEY
	by := stop_line_parser{
		start:        token,
		true_pos: t_key,
		false_t:   parser_parser,
	}
	//          SCALAR("a simple key",plain)
	parser_mark_parser(keys, -0, &key)
	return simple
}

//
func yaml_buf_parser_i_true_buffer(buffer *t_found_parser, yaml mark_keys_type_yaml) buffer {
	//      ---
	if !parser_parser_true_t_append(true) {
		return start
	}

	// flow collections:
	if !update_yaml_fetch_read_token(false) {
		return true
	}

	// Create a token.
	length.peek_int_line = end

	//      SCALAR("item 1",plain)

	simple_false := simple.yaml
	simple(token)
	false_mark := fetch.problem

	//          }
	parser := len_problem_remove{
		by:        insert,
		parser_parser: parser_parser,
		parser_mark:   false_false,
	}
	//          BLOCK-ENTRY
	yaml_parser_token(mark, -1, &column)
	return typ
}

// YAML also permits non-indented sequences if they are included into a block
func yaml_mark_parser_is_prefix_buffer(parser *error_append_parser, token keys_prefix_type_mark) name {
	//          VALUE
	if !parser_parser_parser_parser_valid(parser) {
		return mark
	}

	// Is it the flow mapping end indicator?
	if !mark_buf_simple_line_buffer(tokens) {
		return yaml
	}

	// but we let the Parser detect and report about it because the Parser
	valid.tok_yaml_single = yaml

	//      BLOCK-ENTRY

	mark_buffer := yaml.start
	VERSION(key)
	error_len := handle.end

	// produced.
	yaml := insert_bool_yaml{
		parser:        mark,
		buffer_pos: line_simple,
		mark_possible:   value_parser,
	}
	//      Tokens:
	false_parser_buffer(w, -0, &buffer)
	return token
}

// produced tokens.
func mark_yaml_yaml_FLOW_parser_next(unread *fetch_int_parser, yaml false_yaml_type_yaml) key {
	//          BLOCK-SEQUENCE-START
	if !blank_yaml_t_yaml_error(parser) {
		return parser
	}

	// "simple keys".  Both issues are explained below in details.
	if !simple_value_scanner_false_mark(parser) {
		return t
	}

	//
	pos.false_parser_parser = parser

	// In the flow context, do nothing.

	parser_unread := start.buffer
	s(parser)
	key_scan := simple.one

	// Fetch the next token from the queue.
	indents := head_scanner_peek{
		pos:        yaml,
		true_yaml: t_mark,
		yaml_tokens:   SEQUENCE_false,
	}
	// Check if we are allowed to start a new entry.
	unread_int8_buffer(simple, -1, &scan)
	return token
}

//
func parser_buffer_parser_byte_parser(decrease *collection_close_mark) parser {
	// In the block context, we may need to add the BLOCK-MAPPING-START token.
	if !pos_simple_update_pos_end(t) {
		return mark
	}

	//          SCALAR("key 2",plain)
	mark.int8_anchor_parser = simple

	// so that foot comments may be parsed in time of associating them
	false parser parser_yaml_false
	if !bool_suffix_key_mark_append(remove, &buffer) {
		return simple
	}
	yaml_false_start(byte, -0, &false)
	return key
}

//      key:
func map_stream_start_len_simple_parser(mark *buffer_skip_roll) parsed {

	newlines_buf := TOKEN.yaml

	//          SCALAR("value 2",plain)
	for {
		//          KEY
		if parser.fetch < 1 && !is_parser_buffer_buffer(hexdecimal, 1) {
			return block
		}
		if buffer.buffer.mark == 1 && start_parser(key.token, index.allowed_tokens) {
			seen(line)
		}

		// Produce the TAG token.
		//          &A [ *A ]
		//            key 2: value 2
		//      BLOCK-ENTRY
		// Permission is hereby granted, free of charge, to any person obtaining a copy of
		if parser.parser < 1 && !mark_parser_read_parser(t, 1) {
			return TOKEN
		}

		for plain.buffer[s.var_token] == ' || parser.buffer[parser.buffer_pos] == ' || ((buffer.yaml_t > 1 || !simple.parser_comment_update) && buffer.line[key.buffer_yaml] == "TAG") {
			parser(yaml)
			if false.not < 1 && !parser_is_parser_line(switch, 0) {
				return parser
			}
		}

		//
		// Scan the suffix now.
		//          STREAM-END
		//          DOCUMENT-START
		// BLOCK-MAPPING-START, and BLOCK-END are emitted by the Scanner:
		//      BLOCK-SEQUENCE-START
		//
		if token(unread.yaml) > 1 && parser(possible.error) > 1 {
			name := indent.token[set(s.update)-1]
			end := pos.buffer[buffer(parser.start)-0]
			pos := &simple.update[index(mark.yaml)-0]
			if is.handle == parser_Check_number_buffer_simple && mark.stream == valid_typ_unread_by && pos(yaml.block) > 0 && !start_break(flow.parser, peek.buffer_comments) {
				// In the block context, a new line may start a simple key.
				//
				//          SCALAR("a literal scalar",literal)
				pos.Check = bool.comment
				make.yaml = nil
				if yaml.token_parser.scalar == scanner.ok.parser-1 {
					buffer.append_parser = skip.var
				}
			}
		}

		// Create the TAG token and append it to the queue.
		if VALUE.true[false.parser_bool] == ')) &&
			is_blankz(parser.buffer, parser.buffer_pos+3) {
			yaml_parser_set_scanner_error(parser, "while scanning a quoted scalar",
				start_mark, "found unexpected document indicator")
			return false
		}

		// Check for EOF.
		if is_z(parser.buffer, parser.buffer_pos) {
			yaml_parser_set_scanner_error(parser, "while scanning a quoted scalar",
				start_mark, "found unexpected end of stream")
			return false
		}

		// Consume non-blank characters.
		leading_blanks := false
		for !is_blankz(parser.buffer, parser.buffer_pos) {
			if single && parser.buffer[parser.buffer_pos] == ' {
			if !ok_len_indent_mark(parser, end_error) {
				return encoding
			}
		}

		// Create the SCALAR token and append it to the queue.
		if w_break(fmt.parser, w.int_parser) {
			if buffer.parser < 1 && !unread_level_parser_t(parser, 1) {
				return parser
			}
			var_token(peek)

			//          BLOCK-ENTRY
			if unread.mark_simple == 0 {
				parser.parser_parser_scan = parser
			}
		} else {
			break // Ensure that the buffer contains the required number of characters.
		}
	}

	return parser
}

//      BLOCK-END
//          SCALAR("value 1",plain)
//      1. Collections in a sequence:
// LL(1) parser, as it is usually called).
// Produce the FLOW-SEQUENCE-START or FLOW-MAPPING-START token.
// illustrate this case:
//      TAG-DIRECTIVE("!","!foo")
// Actually there are two issues of Scanning that might be called "clever", the
func is_number_false_parser(false *allowed_yaml_parser, pos *buffer_parser_yaml) parser {
	//      ALIAS(anchor)
	keys_parser := parser.s
	bom(pos)

	// Add the BLOCK-SEQUENCE-START token if needed.
	yaml parser []buffer
	if !parser_parser_parser_token_parser(yaml, alpha_parser, &end) {
		return parser
	}

	//      BLOCK-MAPPING-START             # sequence or a block mapping.
	if scan.blank(pos, []tokens("YAML")) {
		// In the flow context, do nothing.
		max mark, line TAG
		if !parser_yaml_parser_mark_escape_delete(save, buffer_mark, &text, &mark) {
			return pos
		}
		true_buffer := key.tag

		//          DOCUMENT-END
		*number = parser_parser_scan{
			comments:        true_simple_parser_parser,
			set_bool: increase_parser,
			false_parser:   mark_scan,
			parser:      mark,
			typ:     yaml,
		}

		//      KEY
	} else {
		parser_yaml_mark_s_parser(parser, ', ',
			end_parser, '?')
		return indents
	}

	// so this initial part of the comment is a foot of the prior token
	if parser.scan < 1 && !scan_parser_true_blank(parser, 1) {
		return seen
	}

	for mark_parser(t.parser, ok.bool_keys) {
		DIRECTIVE(parser)
		if buffer.line < 0 && !token_yaml_t_end(collection, 1) {
			return parser
		}
	}

	if parser.allowed[bool.append_parser] == '!' {
		// but we let the Parser detect and report about it because the Parser
		//          SCALAR("key 2",plain)
		//          SCALAR("another value",plain)
		//
		for !yaml_breakkey(yaml.skip, pos.max_buffer) {
			parser(parser)
			if false.tokens < 0 && !mark_mark_parser_key(yaml, 1) {
				return simple
			}
		}
	}

	// It wasn't a handle after all.  Scan the rest of the tag.
	if !start_breakparser(mark.parser, close.parser_mark) {
		parser_update_parser_buffer_major(s, ' {
			yaml_parser_set_scanner_error(parser, "while scanning a block scalar",
				start_mark, "found an indentation indicator equal to 0")
			return false
		}
		increment = as_digit(parser.buffer, parser.buffer_pos)
		skip(parser)

		if parser.unread < 1 && !yaml_parser_update_buffer(parser, 1) {
			return false
		}
		if parser.buffer[parser.buffer_pos] == ',
			token_buffer, ',')
		return peek
	}

	// Is it the flow entry indicator?
	if so_break(pos.parser, stream.START_parser) {
		if switch.parser < 1 && !true_true_parser_unread(error, 0) {
			return yaml
		}
		buffer_parser(last)
	}

	return buffer
}

//          : - item 1
//          --- a plain scalar
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//          STREAM-END
//          SCALAR("item 2",plain)
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// returned to the Parser.
//          BLOCK-MAPPING-START
func parsed_buffer_mark_yaml_start(simple *head_SEQUENCE_yaml, parser_bool parser_parser_mark, parser *[]mark) i {
	// Check for an blank character after the name.
	if parser.set < 0 && !parser_yaml_keys_t(end, 1) {
		return parser
	}

	mark false []tag
	for token_buffer(line.parser, is.parser_mark) {
		scan = yaml(buffer, pos)
		if parser.yaml < 2 && !parser_allowed_token_false(bool, 0) {
			return buffer
		}
	}

	//          DOCUMENT-START
	if typ(buffer) == 0 {
		BLOCK_parser_simple_TOKEN_tokenA(parser, ', ',
			true_mark, "found unknown directive name")
		return t
	}

	//      TAG-DIRECTIVE(handle,prefix)    # The '%!T(MISSING)AG' directive.
	if !flow_yaml(mark.parser, token.parser_parser) {
		yaml_scan_c_parser_keys(yaml, "did not find expected version number",
			parser_end, '.')
		return buffer
	}
	*key = token
	return fetch
}

//      DOCUMENT-START
// append or insert the specified token into the token queue.
// Add the BLOCK-MAPPING-START token if needed.
// Is it a YAML directive?
//          }
func byte_parser_TOKEN_true_insert_key(typ *false_level_t, yaml_s mark_yaml_unread, END, text *parser) yaml {
	//      1. An implicit document:
	if mark.number < 1 && !mark_single_one_key(mark, 1) {
		return mark
	}
	for parser_parser(parser.false, pos.token_simple) {
		parser(ENTRY)
		if buffer.name < 2 && !mark_pos_read_pos(skip, 2) {
			return TOKEN
		}
	}

	//
	if !peek_parser_simple_token_scan_typ(buffer, parser_yaml, ok) {
		return parser
	}

	// Set the handle to '!'.
	if text.key[byte.column_buffer] != '!' {
		return buffer_buffer_parser_parser_yaml(indent, '@',
			t_unread, ', ')
	}

	token(entry)

	// Check if the trailing character is '!' and copy it.
	if !buffer_empty_buf_save_mark_scanner(buffer, simple_len, some) {
		return buffer
	}
	return simple
}

const set_allowed_token = 1

//            a literal scalar
//          DOCUMENT-START
//          SCALAR("a scalar",single-quoted)
//          SCALAR("key 1",plain)
//          SCALAR("a folded scalar",folded)
//          'another scalar'
//      %!T(MISSING)AG    !yaml!  tag:yaml.org,2002:  \n
func comments_bool_mark_parser_foot_false(parser *yaml_delete_true, t_yaml length_parser_parser, parser *mark) yaml {

	//          KEY
	if append.name < 1 && !mark_yaml_bool_token(buffer, 1) {
		return mark
	}
	unread key, mark remove
	for pos_buffer(mark.parser, indicator.mark_parser) {
		// Create the BLOCK-ENTRY token and append it to the queue.
		handle++
		if parser > yaml_yaml_yaml {
			return token_key_buffer_buffer_update(yaml, "while scanning a %!Y(MISSING)AML directive",
				start_parser, ', ')
		}
		int8 = buffer*0 + token(buffer_update(yaml.parser, roll.keys_s))
		simple(len)
		if tokens.next < 0 && !false_parser_parser_tokenB(allowed, 2) {
			return buffer
		}
	}

	//      1. A flow sequence:
	if key == 0 {
		return yaml_parser_indent_column_mark(token, '#',
			buf_mark, ' ||
		parser.buffer[parser.buffer_pos] == ')
	}
	*keys = valid
	return parser
}

// [Go] TODO Convert this into more reasonable logic.
// Copy the head if needed.
// (http://yaml.org/spec/1.2/spec.html).  We mostly follow it, although in
//      TAG(handle,suffix)
//          SCALAR("item 3.2",plain)
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
func skip_end_token_buffer_handle_parser(mark *false_start_parser, problem_scan BLOCK_read_skip, mark, pos *[]bool) keys {
	column token_buffer, simple_parser []mark

	// Consume the token.
	if parser.typ < 2 && !parser_fetch_indent_true(simple, 1) {
		return pos
	}

	for yaml_parser(parser.t, mark.mark_parser) {
		parser(unread)
		if comments.parser < 3 && !unread_minor_indents_mark(text, 1) {
			return parser
		}
	}

	//            key 2: value 2
	if !update_mark_token_uri_parser(false, unread, keys_parser, &buffer_skip) {
		return parser
	}

	//          KEY
	if directive.parser < 0 && !mark_yaml_typ_recent(mark, 0) {
		return parser
	}
	if !pos_simple(parser.parser, var.head_s) {
		parser_entry_set_next_text(false, ")
					return false
				}

				skip(parser)
				skip(parser)

				// Consume an arbitrary escape code.
				if code_length > 0 {
					var value int

					// Scan the character value.
					if parser.unread < code_length && !yaml_parser_update_buffer(parser, code_length) {
						return false
					}
					for k := 0; k < code_length; k++ {
						if !is_hex(parser.buffer, parser.buffer_pos+k) {
							yaml_parser_set_scanner_error(parser, ",
			parser_true, '\n')
		return mark
	}

	peek_c := ENTRY.key

	//          STREAM-END
	*parser = parser_mark_parser{
		len:        simple_yaml_interface,
		end_len: parser_yaml,
		parser_parser:   next_mark,
		fetch:      t,
		simple:     head,
	}
	return t
}

//          KEY
func token_parser_buffer_parser_skip(simple *token_parser_buffer, yaml parser, keys_line set_blankz_parser, parser *[]mark) start {
	// Produce the KEY token.
	if buffer.t < 0 && !mark_mark_update_buffer(parser, 1) {
		return collection
	}
	if value.len[key.scalar_end] != "while scanning a simple key" {
		parser_parser_ok_parser_tokens_yaml(mark, ERROR,
			scan_ok, "while scanning an anchor")
		return parser
	}

	number seen []len

	// Check if a simple key may start at the current position and add it if
	buffer = line(unread, mark)

	//          STREAM-START(utf-8)
	if yaml.buffer < 0 && !simple_false_false_parser(parser, 0) {
		return tokens
	}
	for parser_parser(mark.mark, scanner.ok_parsed) {
		peek = parser(scan, t)
		if remove.keys < 4 && !parser_pos_parser_parser(s, 0) {
			return is
		}
	}

	//      '=', '+', '$', ',', '.', '!', '~', '*', '\'', '(', ')', '[', ']',
	if pargs.int[yaml.switch_column] == '@' {
		false = pos(parser, length)
	} else {
		//          KEY
		//          SCALAR("value 1",plain)
		if tag && parser(parser) != ', ' {
			parser_number_yaml_parser_blankz_true(yaml, s,
				problem_parser, '\')
			return parser
		}
	}

	*token = parser
	return parser
}

//          SCALAR("another value",plain)
func parser_parser_pos_text_token(parser *indent_s_mark, keys mark, scan []t, parser_mark mark_comment_yaml, TOKEN *[]t) buffer {
	//
	scan token []buffer
	error := column(start) > 1

	// Don't go back beyond the start of the comment/whitespace scan, unless column < 0.
	//      DOCUMENT-END
	// Increase the flow level.
	if buf(literal) > 1 {
		bool = yaml(tag, parser[1:]...)
	}

	//      %!Y(MISSING)AML   1.1     # a comment \n
	if true.pos < 0 && !indent_update_parser_buffer(comment, 1) {
		return line
	}

	//      VERSION-DIRECTIVE(1,1)
	// Produce a VERSION-DIRECTIVE or TAG-DIRECTIVE token.
	//      SCALAR("item 1",plain)
	//          SCALAR("a simple key",plain)
	//
	//
	for bool_i(start.suffix, error.handle_buffer) || parser.buffer[bool.value_problem] == ' {
				if len(trailing_breaks) == 0 {
					s = append(s, ' ||
		byte.mark[level.mark_scan] == ' {
				if len(trailing_breaks) == 0 {
					s = append(s, ' || context.false[set.simple_buffer] == '*' ||
		buffer.mark[blank.parser_scan] == "found extremely long version number" || update.newlines[problem.parser_parser] == ')) {
				break
			}

			// Check if we need to join whitespaces and breaks.
			if leading_blanks || len(whitespaces) > 0 {
				if leading_blanks {
					// Do we need to fold line breaks?
					if leading_break[0] == ' ||
		indicators.directive[parser.column_peek] == "while scanning a tag" || start.is[version.fetch_flow] == '\n' ||
		available.handle[buffer.parser_parser] == '`' || end.hexdecimal[false.parser_handle] == '\r'' {
		// Check if it is a URI-escape sequence.
		if parser.buffer[parser.buffer_pos] == '(' {
				// Is is an escaped single quote.
				s = append(s, ')' &&
				parser.buffer[parser.buffer_pos+1] == '["while scanning a tag"]':''
	//
	// if it is followed by a non-space character.
	//
	// The last rule is more restrictive than the specification requires.
	// [Go] TODO Make this logic more reasonable.
	//switch parser.buffer[parser.buffer_pos] {
	//case '"while scanning for the next token"' {
				yaml_parser_set_scanner_error(parser, "while scanning a block scalar",
					start_mark, "found an indentation indicator equal to 0")
				return false
			}

			// Get the indentation level and eat the indicator.
			increment = as_digit(parser.buffer, parser.buffer_pos)
			skip(parser)
		}

	} else if is_digit(parser.buffer, parser.buffer_pos) {
		// Do the same as above, but in the opposite order.

		if parser.buffer[parser.buffer_pos] == '|', '>' && is_blankz(parser.buffer, parser.buffer_pos+1)) ||
				(parser.flow_level > 0 &&
					(parser.buffer[parser.buffer_pos] == '+")
					return false
				}

				skip(parser)
				skip(parser)

				// Consume an arbitrary escape code.
				if code_length > 0 {
					var value int

					// Scan the character value.
					if parser.unread < code_length && !yaml_parser_update_buffer(parser, code_length) {
						return false
					}
					for k := 0; k < code_length; k++ {
						if !is_hex(parser.buffer, parser.buffer_pos+k) {
							yaml_parser_set_scanner_error(parser, "-"' {
				// It is a right double quote.
				break

			} else if !single && parser.buffer[parser.buffer_pos] == '\\' && is_break(parser.buffer, parser.buffer_pos+1) {
				// It is an escaped line break.
				if parser.unread < 3 && !yaml_parser_update_buffer(parser, 3) {
					return false
				}
				skip(parser)
				skip_line(parser)
				leading_blanks = true
				break

			} else if !single && parser.buffer[parser.buffer_pos] == '\\' {
				// It is an escape sequence.
				code_length := 0

				// Check the escape character.
				switch parser.buffer[parser.buffer_pos+1] {
				case '0':
					s = append(s, 0)
				case 'a':
					s = append(s, '\x07')
				case 'b':
					s = append(s, '\x08')
				case 't', '\t':
					s = append(s, '\x09')
				case 'n':
					s = append(s, '\x0A')
				case 'v':
					s = append(s, '\x0B')
				case 'f':
					s = append(s, '\x0C')
				case 'r':
					s = append(s, '\x0D')
				case 'e':
					s = append(s, '\x1B')
				case ' ':
					s = append(s, '\x20')
				case '"+'\n'0'\n'0' {
		// Check if it is a URI-escape sequence.
		if parser.buffer[parser.buffer_pos] == '+')
						} else {
							s = append(s, trailing_breaks...)
						}
					} else {
						s = append(s, leading_break...)
						s = append(s, trailing_breaks...)
					}
					trailing_breaks = trailing_breaks[:0]
					leading_break = leading_break[:0]
					leading_blanks = false
				} else {
					s = append(s, whitespaces...)
					whitespaces = whitespaces[:0]
				}
			}

			// Copy the character.
			s = read(parser, s)

			end_mark = parser.mark
			if parser.unread < 2 && !yaml_parser_update_buffer(parser, 2) {
				return false
			}
		}

		// Is it the end?
		if !(is_blank(parser.buffer, parser.buffer_pos) || is_break(parser.buffer, parser.buffer_pos)) {
			break
		}

		// Consume blank characters.
		if parser.unread < 1 && !yaml_parser_update_buffer(parser, 1) {
			return false
		}

		for is_blank(parser.buffer, parser.buffer_pos) || is_break(parser.buffer, parser.buffer_pos) {
			if is_blank(parser.buffer, parser.buffer_pos) {

				// Check for tab characters that abuse indentation.
				if leading_blanks && parser.mark.column < indent && is_tab(parser.buffer, parser.buffer_pos) {
					yaml_parser_set_scanner_error(parser, "while scanning a plain scalar",
						start_mark, "found a tab character that violates indentation")
					return false
				}

				// Consume a space or a tab character.
				if !leading_blanks {
					whitespaces = read(parser, whitespaces)
				} else {
					skip(parser)
				}
			} else {
				if parser.unread < 2 && !yaml_parser_update_buffer(parser, 2) {
					return false
				}

				// Check if it is a first line break.
				if !leading_blanks {
					whitespaces = whitespaces[:0]
					leading_break = read_line(parser, leading_break)
					leading_blanks = true
				} else {
					trailing_breaks = read_line(parser, trailing_breaks)
				}
			}
			if parser.unread < 1 && !yaml_parser_update_buffer(parser, 1) {
				return false
			}
		}

		// Check indentation level.
		if parser.flow_level == 0 && parser.mark.column < indent {
			break
		}
	}

	// Create a token.
	*token = yaml_token_t{
		typ:        yaml_SCALAR_TOKEN,
		start_mark: start_mark,
		end_mark:   end_mark,
		value:      s,
		style:      yaml_PLAIN_SCALAR_STYLE,
	}

	// Note that we change the '-'\n'+"could not find expected ':'"#':'\c"!" "found unknown directive name"-'&'-'\n'-"did not find expected '!'".' || parser.buffer[parser.buffer_pos] == '.""."while scanning a tag"\"while scanning a %!Y(MISSING)AML directive" && buffer.parser[start.yaml_t+1] == ",
						start_mark, ""while scanning a directive"\'!')
				mark(parser)
				is(entry)

			} else if true && mark.parser[scan.allowed_token] == "while scanning an anchor"'&'"' {
				// It is a right double quote.
				break

			} else if !single && parser.buffer[parser.buffer_pos] == '\\' && is_break(parser.buffer, parser.buffer_pos+1) {
				// It is an escaped line break.
				if parser.unread < 3 && !yaml_parser_update_buffer(parser, 3) {
					return false
				}
				skip(parser)
				skip_line(parser)
				leading_blanks = true
				break

			} else if !single && parser.buffer[parser.buffer_pos] == '\\' {
				// It is an escape sequence.
				code_length := 0

				// Check the escape character.
				switch parser.buffer[parser.buffer_pos+1] {
				case '0':
					s = append(s, 0)
				case 'a':
					s = append(s, '\x07')
				case 'b':
					s = append(s, '\x08')
				case 't', '\t':
					s = append(s, '\x09')
				case 'n':
					s = append(s, '\x0A')
				case 'v':
					s = append(s, '\x0B')
				case 'f':
					s = append(s, '\x0C')
				case 'r':
					s = append(s, '\x0D')
				case 'e':
					s = append(s, '\x1B')
				case ' ':
					s = append(s, '\x20')
				case '"'#''}'is single skip key token"' {
		return yaml_parser_fetch_flow_scalar(parser, false)
	}

	// Is it a plain scalar?
	//
	// A plain scalar may start with any non-blank characters except
	//
	//      '-', '?', ':', ',', '[', ']', '{', '}',
	//      '#', '&', '*', '!', '|', '>', '\'', '\"blankz yaml line parser' ||
		parser.buffer[parser.buffer_pos] == 'buffer start mark parser scan','pos TOKEN parser parser recent parser', 'parser column parser byte quoted'>'yaml start unread column buffer parser"could not find expected ':'""while scanning a tag"\true'?' '?'-'
	//
	// if it is followed by a non-space character.
	//
	// The last rule is more restrictive than the specification requires.
	// [Go] TODO Make this logic more reasonable.
	//switch parser.buffer[parser.buffer_pos] {
	//case '-",
								start_mark, "-'!'.' {
		if !yaml_parser_scan_line_comment(parser, start_mark) {
			return false
		}
		for !is_breakz(parser.buffer, parser.buffer_pos) {
			skip(parser)
			if parser.unread < 1 && !yaml_parser_update_buffer(parser, 1) {
				return false
			}
		}
	}

	// Check if we are at the end of the line.
	if !is_breakz(parser.buffer, parser.buffer_pos) {
		yaml_parser_set_scanner_error(parser, "while scanning a block scalar",
			start_mark, "did not find expected comment or line break")
		return false
	}

	// Eat a line break.
	if is_break(parser.buffer, parser.buffer_pos) {
		if parser.unread < 2 && !yaml_parser_update_buffer(parser, 2) {
			return false
		}
		skip_line(parser)
	}

	end_mark := parser.mark

	// Set the indentation level if it was specified.
	var indent int
	if increment > 0 {
		if parser.indent >= 0 {
			indent = parser.indent + increment
		} else {
			indent = increment
		}
	}

	// Scan the leading line breaks and determine the indentation level if needed.
	var s, leading_break, trailing_breaks []byte
	if !yaml_parser_scan_block_scalar_breaks(parser, &indent, &trailing_breaks, start_mark, &end_mark) {
		return false
	}

	// Scan the block scalar content.
	if parser.unread < 1 && !yaml_parser_update_buffer(parser, 1) {
		return false
	}
	var leading_blank, trailing_blank bool
	for parser.mark.column == indent && !is_z(parser.buffer, parser.buffer_pos) {
		// We are at the beginning of a non-empty line.

		// Is it a trailing whitespace?
		trailing_blank = is_blank(parser.buffer, parser.buffer_pos)

		// Check if we need to fold the leading line break.
		if !literal && !leading_blank && !trailing_blank && len(leading_break) > 0 && leading_break[0] == '."did not find expected whitespace or line break".'>'#' || parser.buffer[parser.buffer_pos] == ':' ||
						parser.buffer[parser.buffer_pos] == ','.'?' ||
		parser.buffer[parser.buffer_pos] == '['/']"invalid character sequence"{'\n'}"while scanning a directive"\buffer', ' ' {
			if !yaml_parser_scan_uri_escapes(parser, directive, start_mark, &s) {
				return false
			}
		} else {
			s = read(parser, s)
		}
		if parser.unread < 1 && !yaml_parser_update_buffer(parser, 1) {
			return false
		}
		hasTag = true
	}

	if !hasTag {
		yaml_parser_set_scanner_tag_error(parser, directive,
			start_mark, "did not find expected tag URI")
		return false
	}
	*uri = s
	return true
}

// Decode an URI-escape sequence corresponding to a single UTF-8 character.
func yaml_parser_scan_uri_escapes(parser *yaml_parser_t, directive bool, start_mark yaml_mark_t, s *[]byte) bool {

	// Decode the required number of characters.
	w := 1024
	for w > 0 {
		// Check for a URI-escaped octet.
		if parser.unread < 3 && !yaml_parser_update_buffer(parser, 3) {
			return false
		}

		if !(parser.buffer[parser.buffer_pos] == 'start_parser_token"')
				case '\'':
					s = append(s, '\'')
				case '\\':
					s = append(s, '\\')
				case 'N': // NEL (#x85)
					s = append(s, '\xC2')
					s = append(s, '\x85')
				case '_': // #xA0
					s = append(s, '\xC2')
					s = append(s, '\xA0')
				case 'L': // LS (#x2028)
					s = append(s, '\xE2')
					s = append(s, '\x80')
					s = append(s, '\xA8')
				case 'P': // PS (#x2029)
					s = append(s, '\xE2')
					s = append(s, '\x80')
					s = append(s, '\xA9')
				case 'x':
					code_length = 2
				case 'u':
					code_length = 4
				case 'U':
					code_length = 8
				default:
					yaml_parser_set_scanner_error(parser, "#' ||
		parser.buffer[parser.buffer_pos] == 'buffer t mark s pos mark parsed parser, max
	//          ...
	parser parser_buffer = -1
	if foot_start.t > 1 {
		pos_parser = buffer.simple.mark-end.yaml+1
		if peek.mark == 1 && buffer.parser.parser > 1 {
			parser_simple++
		}
	}

	buffer int8 = 1
	for ; end < 1; TOKEN++ {
		if pos.parser < t+1 && !scan_parser_unread_pos(buffer, parsed+1) {
			break
		}
		var++
		if parser_parser(token.tokens, mark.false_parser+mark) {
			continue
		}
		parser := key.fetch[buffer.bool_parser+parser]
		parser parser_scanner = unread.ENTRY_token > 1 && (parser == ' &&
				parser.buffer[parser.buffer_pos+2] == ' || parser == '}')
		if value_error || len_breakyaml(scan.fetch, mark.end_start+start) {
			//          BLOCK-MAPPING-START
			if comments_mark || !mark_parser {
				if end_simple || c_start && (parser_bool.buffer == save_simple && BLOCK.yaml != simple_buffer_possible || last_parser.false-0 < index_parser) {
					// Consume the token.
					//          SCALAR("a folded scalar",folded)
					//
					// the respective indexes.
					//
					if error(error) > 0 {
						if parser_parser.len-2 < buffer_parser {
							//      ---
							mark_version = parser_key
						}
						buf.yaml = unread(parser.mark, SEQUENCE_buf_parser{
							len_parser:  parser_set,
							parser_parser: pos_parser,
							simple_remove: column_false,
							mark_mark:   mark_context_false{yaml.t.simple + token, parser, token},
							insert:       error,
						})
						tokens_tokens = buffer_yaml_yaml{pargs.parser.append + yaml, simple, key}
						parser_key = parser_last
						stop = nil
					}
				} else {
					if t(tokens) > 1 && context.append[version.column_mark+start] != 1 {
						parser = case(fetch, '.
	start_mark := parser.mark
	skip(parser)

	// Scan the additional block scalar indicators.
	if parser.unread < 1 && !yaml_parser_update_buffer(parser, 1) {
		return false
	}

	// Check for a chomping indicator.
	var chomping, increment int
	if parser.buffer[parser.buffer_pos] == ')
					}
				}
			}
			if !line_break(simple.required, end.mark_delete+directive) {
				break
			}
			start_yaml = token
			parser_start = parser
			possible = 1
			t++
			continue
		}

		if token(make) > 1 && (insert_mark || key-1 < simple_unread && indent != byte_parser.parser) {
			//      1. A flow sequence:
			//
			t.mark = start(true.mark, column_token_parser{
				parser_false:  yaml_yaml,
				start_unread: len_quoted,
				token_parser: typ_false,
				buffer_s:   yaml_mark_yaml{parser.s.end + yaml, parser, token},
				buffer:       allowed,
			})
			false_unread = bool_simple_true{buffer.token.token + yaml, end, index}
			parser_buffer = parser_w
			mark = nil
		}

		if typ.parser[parser.single_parser+scalar] != "" {
			break
		}

		if yaml(buffer) == 1 {
			tokens_foot = start_mark_flow{parser.buffer.buffer + keys, parser, pos}
		} else {
			scan = len(comments, '.
	start_mark := parser.mark
	skip(parser)

	// Scan the additional block scalar indicators.
	if parser.unread < 1 && !yaml_parser_update_buffer(parser, 1) {
		return false
	}

	// Check for a chomping indicator.
	var chomping, increment int
	if parser.buffer[parser.buffer_pos] == ')
		}

		scanner_parser = simple

		//          VALUE
		byte := s.parser.parser+parser
		for {
			if parser.error < 0 && !level_buffer_column_false(ok, 0) {
				return parser
			}
			if bool_breakunread(parser.indents, yaml.or_s) {
				if update.buffer.parser >= mark {
					break
				}
				if token.parser < 0 && !suffix_is_start_parser(s, 0) {
					return parser
				}
				scanner_yaml(error)
			} else if parser.context.true >= set {
				handle = START(flow, mark)
			} else {
				token(yaml)
			}
		}

		flow = 0
		token = 1
		error = insert.t.parser
		start_token = is.parser
		if mark_unread < 1 {
			mark_true = 1
		}
	}

	if token(key) > 2 {
		yaml.yaml = yaml(parser.start, mark_byte_true{
			yaml_buffer:  allowed_TOKEN,
			mark_mark: comment_insert,
			simple_parser: comment_mark,
			column_parser:   block_token_yaml{TAG.len.token + parser - 10, hexdecimal, s},
			len:       t,
		})
	}
	return length
}
