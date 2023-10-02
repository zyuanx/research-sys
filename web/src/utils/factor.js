/**
 * 此文件存放问题类型和校验规则
 */
export const factorItems = {
  // 短文本
  input: {
    factor: 'input',
    label: '短文本',
    value: '',
    placeholder: '请输入内容',
    required: true,
    showWordLimit: true,
    minLength: 0,
    maxLength: 10,
    rules: []
  },
  // 长文本
  textarea: {
    factor: 'textarea',
    label: '长文本',
    value: '',
    placeholder: '请输入内容',
    required: true,
    showWordLimit: true,
    minLength: 0,
    maxLength: 20,
    rows: 3,
    rules: []
  },
  // 单选
  radio: {
    factor: 'radio',
    label: '单选',
    value: '',
    options: [
      { label: '选项1', value: '1' },
      { label: '选项2', value: '2' }
    ],
    required: true,
    requiredMsg: '请选择选项',
    rules: []
  },
  // 多选
  checkbox: {
    factor: 'checkbox',
    label: '多选',
    value: [],
    min: 0,
    max: 2,
    options: [
      { label: '选项1', value: '值1' },
      { label: '选项2', value: '值2' }
    ],
    requiredMsg: '请选择选项',
    required: true,
    rules: []
  },
  // 下拉选择
  select: {
    factor: 'select',
    label: '下拉选择',
    value: '',
    options: [
      { label: '选项1', value: '值1' },
      { label: '选项2', value: '值2' }
    ],
    required: true,
    requiredMsg: '请选择选项',
    rules: []
  },
  // 时间选择
  timePicker: {
    factor: 'timePicker',
    label: '时间选择',
    value: '',
    required: true,
    requiredMsg: '请选择时间'
  },
  // 日期选择
  datePicker: {
    factor: 'datePicker',
    label: '日期选择',
    type: 'date',
    value: '',
    required: true,
    requiredMsg: '请选择日期'
  }
}

export const factorRules = [
  {
    id: 'number',
    label: '数字',
    reg: /^\d+$/,
    tip: '请输入数字'
  },
  {
    id: 'phone',
    label: '手机号',
    reg: /^(?:(?:\+|00)86)?1(?:(?:3[\d])|(?:4[5-79])|(?:5[0-35-9])|(?:6[5-7])|(?:7[0-8])|(?:8[\d])|(?:9[1589]))\d{8}$/,
    tip: '请输入正确的手机号'
  },
  {
    id: 'email',
    label: '邮箱',
    reg: /^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/,
    tip: '请输入正确的邮箱'
  },
  {
    id: 'idCard',
    label: '身份证',
    reg: /^[1-9]\d{5}(?:18|19|20)\d{2}(?:0[1-9]|10|11|12)(?:0[1-9]|[1-2]\d|30|31)\d{3}[\dXx]$/,
    tip: '请输入正确的身份证号'
  },
  {
    id: 'name',
    label: '姓名',
    reg: /^(?:[\u4e00-\u9fa5·]{2,16})$/,
    tip: '请输入正确的姓名'
  },
  {
    id: 'custom',
    label: '自定义',
    reg: '',
    tip: '请输入正确的内容'
  }
]

export const factorRulesMap = factorRules.reduce((map, item) => {
  map[item.id] = item
  return map
}, {})

export const research = {
  pattern: {
    size: 'default',
    labelPosition: 'top',
    labelWidth: '100px'
  },
  config: {
    title: '问卷标题',
    description:
      '同学你好，看到这个问卷说明你收到了我们的低电量推送提醒，目前为测试阶段，您的宝贵意见将是我们进步的动力，感谢对微生活的关注和支持，相信在我们共同努力下能够一起度过美好的时光，祝学业有成，加油！',
    once: 0,
    open: 0,
    startAt: '2023-01-01 00:00:00',
    endAt: '2023-01-31 23:59:59',
    access: []
  },
  items: [
    {
      fieldID: '7762c89-2d78-4447-b1b4-19f902149cbb',
      factor: 'radio',
      label: '单选',
      value: '',
      options: [
        { label: '选项1', value: '1' },
        { label: '选项2', value: '2' }
      ],
      required: true,
      requiredMsg: '请选择选项',
      rules: []
    },
    {
      fieldID: '7cb5873f-697e-4c46-87b6-156d0a368553',
      factor: 'input',
      label: '短文本',
      value: '',
      placeholder: '请输入内容',
      required: true,
      requiredMsg: '请输入内容',
      showWordLimit: true,
      minLength: 0,
      maxLength: 10,
      rules: []
    },
    {
      fieldID: 'c0a35eac-ccee-4705-bb3e-729bc6afb657',
      factor: 'checkbox',
      label: '多选',
      value: [],
      min: 0,
      max: 2,
      required: false,
      requiredMsg: '请选择选项',
      options: [
        { label: '选项1', value: '1' },
        { label: '选项2', value: '2' }
      ],
      rules: []
    },
    {
      fieldID: '7cb5873f-697e-4c46-87b6-156d0a368553',
      factor: 'textarea',
      label: '长文本',
      value: '',
      placeholder: '请输入内容',
      required: true,
      requiredMsg: '请输入内容',
      showWordLimit: true,
      minLength: 0,
      maxLength: 20,
      rows: 3,
      rules: []
    }
  ]
}
