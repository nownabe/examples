import React from "react"

export default class CommentForm extends React.Component {
  handleSubmit(e) {
    e.preventDefault()
    var authorNode = React.findDOMNode(this.refs.author)
    var textNode = React.findDOMNode(this.refs.text)
    var author = authorNode.value.trim()
    var text = textNode.value.trim()

    if (!text || !author) return

    this.props.onCommentSubmit({author: author, text: text})
    authorNode.value = ""
    textNode.value = ""
  }

  render() {
    return(
      <form className="commentForm" onSubmit={this.handleSubmit.bind(this)}>
        <p><input type="text" placeholder="Your name" ref="author" /></p>
        <p><input type="text" placeholder="Say something..." ref="text" /></p>
        <p><input type="submit" value="Post" /></p>
      </form>
    )
  }
}
