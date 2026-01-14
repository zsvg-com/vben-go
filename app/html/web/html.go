package web

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 常量定义
const (
	ProtocolTitle   = "用户服务条款"
	ProtocolContent = `<h1 class="title">用户服务条款</h1>
        <div class="content">
            <h2 class="content-title">服务条款确认与接纳</h2>
            <p class="desc">
                RuoYi 拥有ruoyi.vip其涉及到的产品、相关软件的所有权和运作权， RuoYi
                享有对其产品的上一切活动的监督、提示、检查、纠正及处罚等权利。用户通过注册程序阅读本服务条款并点击"注册"按钮完成注册，即表示用户与 RuoYi
                已达成协议，自愿接受本服务条款的所有内容。如果用户不同意服务条款的条件，则不能获得使用 RuoYi 服务以及注册成为 RuoYi 用户的权利。
            </p>
        </div>
        <div class="content">
            <h2 class="content-title">使用规则</h2>
            <div class="orderly-list">
                <div class="orderly-list-item">
                    <span class="index">1.</span>
                    <p class="orderly-list-item-desc">用户注册成功后，RuoYi
                        将给予每个用户一个用户账号及相应的密码，该用户账号和密码由用户负责保管；用户应当对以其用户账号进行的所有活动和事件负法律责任。</p>
                </div>
                <div class="orderly-list-item">
                    <span class="index">2.</span>
                    <p class="orderly-list-item-desc">用户须对在 RuoYi
                        的注册信息的真实性、合法性、有效性承担全部责任，用户不得冒充他人；不得利用他人的名义发布任何信息；不得恶意使用注册帐户导致其他用户误认；否则 RuoYi
                        有权立即停止提供服务，收回其账号并由用户独自承担由此而产生的一切法律责任。</p>
                </div>
                <div class="orderly-list-item">
                    <span class="index">3.</span>
                    <p class="orderly-list-item-desc">用户不得使用 RuoYi 服务发送或传播敏感信息和违反国家法律制度的信息，包括但不限于下列信息:</p>
                    <div>
                        <ul type="circle">
                            <li>反对宪法所确定的基本原则的；</li>
                            <li>危害国家安全，泄露国家秘密，颠覆国家政权，破坏国家统一的；</li>
                            <li>损害国家荣誉和利益的；</li>
                            <li>煽动民族仇恨、民族歧视，破坏民族团结的；</li>
                            <li>破坏国家宗教政策，宣扬邪教和封建迷信的；</li>
                            <li>散布谣言，扰乱社会秩序，破坏社会稳定的；</li>
                            <li>散布淫秽、色情、赌博、暴力、凶杀、恐怖或者教唆犯罪的；</li>
                            <li>侮辱或者诽谤他人，侵害他人合法权益的；</li>
                            <li>含有法律、行政法规禁止的其他内容的。</li>
                        </ul>
                    </div>
                </div>
                <div class="orderly-list-item">
                    <span class="index">4.</span>
                    <p class="orderly-list-item-desc">RuoYi 有权对用户使用 RuoYi 产品的情况进行审查和监督，如用户在使用 RuoYi
                        产品时违反任何上述规定，RuoYi 或其授权的人有权要求用户改正或直接采取一切必要的措施以减轻用户不当行为造成的影响。</p>
                </div>
                <div class="orderly-list-item">
                    <span class="index">5.</span>
                    <p class="orderly-list-item-desc">盗取他人用户账号或利用网络通讯骚扰他人，均属于非法行为。用户不得采用测试、欺骗等任何非法手段，盗取其他用户的账号和对他人进行骚扰。</p>
                </div>
            </div>
        </div>
        <div class="content">
            <h2 class="content-title">免责声明</h2>
            <div class="orderly-list">
                <div class="orderly-list-item">
                    <span class="index">1.</span>
                    <p class="orderly-list-item-desc">RuoYi 不能对用户在本社区回答问题的答案或评论的准确性及合理性进行保证。</p>
                </div>
                <div class="orderly-list-item">
                    <span class="index">2.</span>
                    <p class="orderly-list-item-desc">若 RuoYi
                        产品已经明示其网络服务提供方式发生变更并提醒用户应当注意事项，用户未按要求操作所产生的一切后果由用户自行承担。</p>
                </div>
                <div class="orderly-list-item">
                    <span class="index">3.</span>
                    <p class="orderly-list-item-desc">用户明确同意其使用 RuoYi 产品网络服务所存在的风险将完全由其自己承担；因其使用 RuoYi
                        服务而产生的一切后果也由其自己承担，RuoYi 对用户不承担任何责任。</p>
                </div>
                <div class="orderly-list-item">
                    <span class="index">4.</span>
                    <p class="orderly-list-item-desc">RuoYi
                        不保证网络服务一定能满足用户的要求，也不保证网络服务不会中断，对网络服务的及时性、安全性、准确性也都不作保证。</p>
                </div>
                <div class="orderly-list-item">
                    <span class="index">5.</span>
                    <p class="orderly-list-item-desc">对于因不可抗力或 RuoYi 不能控制的原因造成的网络服务中断或其它缺陷，RuoYi
                        不承担任何责任，但将尽力减少因此而给用户造成的损失和影响。</p>
                </div>
                <div class="orderly-list-item">
                    <span class="index">6.</span>
                    <p class="orderly-list-item-desc">用户同意保障和维护 RuoYi 及其他用户的利益，用户在 RuoYi
                        发表的内容仅表明其个人的立场和观点，并不代表 RuoYi 的立场或观点。由于用户发表内容违法、不真实、不正当、侵犯第三方合法权益，或用户违反本协议项下的任何条款而给
                        RuoYi 或任何其他第三人造成损失，用户同意承担由此造成的损害赔偿责任。</p>
                </div>
            </div>
        </div>
        <div class="content">
            <h2 class="content-title">服务条款的修改</h2>
            <p class="desc">
                RuoYi 会在必要时修改服务条款，服务条款一旦发生变动，RuoYi
                将会在用户进入下一步使用前的页面提示修改内容。如果您同意改动，请再一次激活"我同意"按钮。如果您不接受，请及时取消您的帐户。 用户要继续使用 RuoYi 各项服务需要两方面的确认:
            </p>
            <div class="orderly-list">
                <div class="orderly-list-item">
                    <span class="index">1.</span>
                    <p class="orderly-list-item-desc">首先确认 RuoYi 服务条款及其变动。</p>
                </div>
                <div class="orderly-list-item">
                    <span class="index">2.</span>
                    <p class="orderly-list-item-desc">同意接受所有的服务条款限制。</p>
                </div>
            </div>
        </div>
        <div class="content">
            <h2 class="content-title">隐私政策</h2>
            <p class="desc">
                RuoYi 非常重视对用户隐私权的保护，承诺不会在未获得用户许可的情况下擅自将用户的个人资料信息出租或出售给任何第三方，但以下情况除外:
            </p>
            <ul type="circle">
                <li>您同意让第三方共享资料；</li>
                <li>您同意公开你的个人资料，享受为您提供的产品和服务；</li>
                <li>本站需要听从法庭传票、法律命令或遵循法律程序；</li>
                <li>本站发现您违反了本站服务条款或本站其它使用规定。</li>
            </ul>
        </div>
        <div class="msg">如果您对此条款有任何疑问或建议，欢迎通过邮件联系我们：908085643@qq.com</div>`
)

// HtmlHandler HTML处理器结构体
type HtmlHandler struct{}

// IndexHandler 手机首页
func (handler *HtmlHandler) IndexHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "mobile", nil)
}

// IndexOldHandler 手机首页(旧版)
func (handler *HtmlHandler) IndexOldHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "mobile_old", nil)
}

// ProtocolHandler 用户协议
func (handler *HtmlHandler) ProtocolHandler(context *gin.Context) {
	context.HTML(http.StatusOK, "protocol", gin.H{
		"title":   ProtocolTitle,
		"content": template.HTML(ProtocolContent),
	})
}
