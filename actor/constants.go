package actor

// [DialogFlow Parameters]
type DialogFlowParameter string

const YogaCategoryParam DialogFlowParameter = "yoga_category"
const VideoCourseNameParam DialogFlowParameter = "course_name"
const VideoNameParam DialogFlowParameter = "video_name"
const YogaPoseParam DialogFlowParameter = "yoga_pose"

// [DialogFlow Actions]
type ActionName string

const showOptionsInCategory ActionName = "show_options_in_category" // L1

const showVideoCoursesInCategory ActionName = "show_video_courses_in_category"            // L2a
const showVideosByCourseNameAndCategory ActionName = "show_videos_in_course_in_category" // L2a
const showAVideoInCourseByCategory ActionName = "show_a_video_in_course_in_category"      // L2a

const showPosesInCategory ActionName = "show_poses_in_category" // L2b
const showOriginalPoseImg ActionName = "show_original_pose_img" // L3b1
const showRightPoseImg ActionName = "show_right_pose_img"       // L3b2
const showWrongPoseImg ActionName = "show_wrong_pose_img"       // L3b3

const showArticlesInCategory ActionName = "show_articles_in_category"          // L2c
const showRandomOptionInCategory ActionName = "show_random_option_in_category" // L2d
