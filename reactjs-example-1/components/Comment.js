import React from "react"

export default class Comment extends React.Component {
  render() {
    return(
      <div className="comment">
        <h3 className="commentAuthor">
          {this.props.author}
        </h3>
        <p>{this.props.children.toString()}</p>
      </div>
    )
  }
}
