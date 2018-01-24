package actor

// [DialogFlow Parameters]
type DialogFlowParameter string

const YogaCategoryParam DialogFlowParameter = "yoga_category"
const YogaPoseParam DialogFlowParameter = "yoga_pose"

// [DialogFlow Actions]
type ActionName string

const showPosesInCategory ActionName = "show_poses_in_category"
const showOriginalPoseImg ActionName = "show_original_pose_img"
const showRightPoseImg ActionName = "show_right_pose_img"
const showWrongPoseImg ActionName = "show_wrong_pose_img"
