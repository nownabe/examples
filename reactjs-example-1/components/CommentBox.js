import React from "react"
import CommentList from "./CommentList"
import CommentForm from "./CommentForm"

export default class CommentBox extends React.Component {
  constructor(props) {
    super(props)

    this.state = {
      data: []
    }
  }

  componentDidMount() {
    // Pull data from server.
  }

  handleCommentSubmit(comment) {
    var comments = this.state.data;
    var newComments = comments.concat([comment])
    this.setState({data: newComments})

    // Send data to server.
  }

  render() {
    return(
      <div className="commentBox">
        <h2>Comments</h2>
        <CommentForm onCommentSubmit={this.handleCommentSubmit.bind(this)} />
        <CommentList data={this.state.data} />
      </div>
    )
  }
}
