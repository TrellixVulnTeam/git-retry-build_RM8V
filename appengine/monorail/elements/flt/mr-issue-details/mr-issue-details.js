'use strict';

/**
 * `<mr-issue-details>`
 *
 * This is the main details section for a given issue.
 *
 */
class MrIssueDetails extends ReduxMixin(Polymer.Element) {
  static get is() {
    return 'mr-issue-details';
  }

  static get properties() {
    return {
      comments: {
        type: Array,
        statePath: 'comments',
      },
      issueId: {
        type: Number,
        statePath: 'issueId',
      },
      projectName: {
        type: String,
        statePath: 'projectName',
      },
      token: {
        type: String,
        statePath: 'token',
      },
      _description: {
        type: String,
        computed: '_computeDescription(comments)',
      },
      _comments: {
        type: Array,
        computed: '_filterComments(comments)',
      },
      _onSubmitComment: {
        type: Function,
        value: function() {
          return this._submitCommentHandler.bind(this);
        },
      },
      _updateDescription: {
        type: Function,
        value: function() {
          return this._updateDescriptionHandler.bind(this);
        },
      },
    };
  }

  _submitCommentHandler(comment) {
    this._updateIssue(comment);
  }

  _updateDescriptionHandler(content) {
    this._updateIssue(content, null, true);
  }

  _updateIssue(commentData, delta, isDescription) {
    const message = {
      trace: {token: this.token},
      issue_ref: {
        project_name: this.projectName,
        local_id: this.issueId,
      },
      comment_content: commentData || '',
    };

    if (delta) {
      message.delta = delta;
    }

    if (isDescription) {
      message.is_description = true;
    }

    actionCreator.updateIssue(this.dispatch.bind(this), message);
  }

  _filterComments(comments) {
    return comments.filter((c) => (!c.approvalRef && c.sequenceNum));
  }

  _computeDescription(comments) {
    for (let i = comments.length - 1; i >= 0; i--) {
      if (!comments[i].approvalRef && comments[i].descriptionNum) {
        return comments[i];
      }
    }
    return {};
  }
}
customElements.define(MrIssueDetails.is, MrIssueDetails);
