pacchetto principale

//
// import (
// "riflettere"
// "test"
//)
//
// func TestCheckUserReg (t * testing.T) {
// new_cases: = [] int {3, 5, 7, 9, 10}
// test_cases: = [] int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
//
// previsto: = [] bool {vero, falso, vero, falso, vero, falso, vero, falso, vero, vero}
// var got [] bool
//
// InsertNewCases (new_cases)
//
// for _, one_case: = range test_cases {
// got = append (got, CheckUserReg (one_case))
//}
//
// if! reflection.DeepEqual (got, expected) {
// t.Errorf ("Previsto% v, ottenuto% v", previsto, ottenuto)
//}
//}
//
// func TestIsUserChatting (t * testing.T) {
// new_cases: = [] int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// chatting_cases: = [] int {3, 5, 7, 9, 10}
//
// previsto: = [] bool {false, false, true, false, true, false, true, false, true, true}
// var got [] bool
//
// InsertNewCases (new_cases)
//
// for _, one_case: = range chatting_cases {
// ChangeUserChattingState (one_case, true)
//}
//
// for _, one_case: = range new_cases {
// got = append (got, IsUserChatting (one_case))
//}
//
// if! reflection.DeepEqual (got, expected) {
// t.Errorf ("Previsto% v, ottenuto% v", previsto, ottenuto)
//}
//}
//
// func TestIsUserSearching (t * testing.T) {
// new_cases: = [] int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// search_cases: = [] int {3, 5, 7, 9, 10}
//
// previsto: = [] bool {false, false, true, false, true, false, true, false, true, true}
// var got [] bool
//
// InsertNewCases (new_cases)
//
// for _, one_case: = range research_cases {
// ChangeUserSearchingState (one_case, true)
//}
//
// for _, one_case: = range new_cases {
// got = append (got, IsUserSearching (one_case))
//}
//
// if! reflection.DeepEqual (got, expected) {
// t.Errorf ("Previsto% v, ottenuto% v", previsto, ottenuto)
//}
//}
//
// func TestFindFreeUsers (t * testing.T) {
// new_cases: = [] int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// search_cases: = [] int {3, 5, 7, 9, 10}
//
// previsto: = [] int {3, 5, 7, 9, 10}
// var ha [] int
//
// InsertNewCases (new_cases)
//
// for _, one_case: = range research_cases {
// ChangeUserSearchingState (one_case, true)
//}
//
// got = FindFreeUsers ()
//
// if! reflection.DeepEqual (got, expected) {
// t.Errorf ("Previsto% v, ottenuto% v", previsto, ottenuto)
//}
//}
//
// func TestFindSecondUserFromChat (t * testing.T) {
// new_cases: = [] [] int {{1, 2}, {3, 4}, {5, 7}}
//
// previsto: = [] int {2, 4, 7}
// var ha [] int
//
// for _, one_pair: = range new_cases {
// AddNewChat (one_pair [0], one_pair [1])
//}
//
// for _, one_pair: = range new_cases {
// got = append (got, FindSecondUserFromChat (one_pair [0]))
//}
//
// if! reflection.DeepEqual (got, expected) {
// t.Errorf ("Previsto% v, ottenuto% v", previsto, ottenuto)
//}
//}
//
// func TestDeleteChat (t * testing.T) {
// DeleteChat (1)
// DeleteChat (3)
//
// if FindSecondUserFromChat (5)! = 7 {
// t.Errorf ("Previsto% v, ottenuto% v", 7, 0)
//}
//}
//
// func InsertNewCases (new_cases [] int) {
// for _, one_case: = range new_cases {
// UserFirstStart (one_case)
//}
//}
