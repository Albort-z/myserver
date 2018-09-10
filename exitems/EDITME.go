/*
 * 配置文件中的Path为模块内部初始化的Name
 * 要加载的模块必须在此文件中import
 * ex: import ( _ "Hello" )
 *
 *
 */

package exitems

import (
	_ "github.com/Albort-z/myserver/exitems/hello"
)
