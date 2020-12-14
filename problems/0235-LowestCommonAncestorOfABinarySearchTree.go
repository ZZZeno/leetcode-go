package problems

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    cur := root
    var min, max *TreeNode
    if p.Val < q.Val {
        min = p
        max = q
    } else {
        min = q
        max = p
    }
    for {
        if cur.Val == p.Val || cur.Val == q.Val || (cur.Val > min.Val && cur.Val < max.Val) {
            break
        }
        if cur.Val < min.Val {
            cur = cur.Right
        }
        if cur.Val > max.Val {
            cur = cur.Left
        }
    }
    return cur
}
