package models

type UserInfo struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID           int64  `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name,omitempty"`
			Username     string `json:"username,omitempty"`
			LanguageCode string `json:"language_code,omitempty"`
			IsPremium    bool   `json:"is_premium"`
		} `json:"from"`
		Chat struct {
			ID        int64  `json:"id"`
			FirstName string `json:"first_name,omitempty"`
			LastName  string `json:"last_name,omitempty"`
			Username  string `json:"username,omitempty"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date     int    `json:"date"`
		Text     string `json:"text,omitempty"`
		Entities []struct {
			Offset int    `json:"offset"`
			Length int    `json:"length"`
			Type   string `json:"type"`
		} `json:"entities,omitempty"`
	} `json:"message"`
}

type StartButton struct {
	ChatID      int64  `json:"chat_id"`
	Photo       string `json:"photo"`
	Caption     string `json:"caption"`
	ReplyMarkup struct {
		InlineKeyboard [][]struct {
			Text         string  `json:"text"`
			CallbackData string  `json:"callback_data,omitempty"`
			WebApp       *WebApp `json:"web_app,omitempty"`
		} `json:"inline_keyboard"`
	} `json:"reply_markup"`
}

type WebApp struct {
	URL string `json:"url"`
}

type ButtonStandard struct {
	UpdateID      int `json:"update_id"`
	CallbackQuery struct {
		ID   string `json:"id"`
		From struct {
			ID           int64  `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name,omitempty"`
			Username     string `json:"username,omitempty"`
			LanguageCode string `json:"language_code,omitempty"`
			IsPremium    bool   `json:"is_premium"`
		} `json:"from"`
		Message struct {
			MessageID int `json:"message_id"`
			From      struct {
				ID        int64  `json:"id"`
				IsBot     bool   `json:"is_bot"`
				FirstName string `json:"first_name"`
				Username  string `json:"username"`
			} `json:"from"`
			Chat struct {
				ID        int64  `json:"id"`
				FirstName string `json:"first_name,omitempty"`
				LastName  string `json:"last_name,omitempty"`
				Username  string `json:"username,omitempty"`
				Type      string `json:"type"`
			} `json:"chat"`
			Date  int `json:"date"`
			Photo []struct {
				FileID       string `json:"file_id"`
				FileUniqueID string `json:"file_unique_id"`
				FileSize     int    `json:"file_size"`
				Width        int    `json:"width"`
				Height       int    `json:"height"`
			} `json:"photo,omitempty"`
			Caption     string `json:"caption,omitempty"`
			ReplyMarkup struct {
				InlineKeyboard [][]struct {
					Text         string `json:"text"`
					CallbackData string `json:"callback_data"`
				} `json:"inline_keyboard"`
			} `json:"reply_markup,omitempty"`
		} `json:"message"`
		ChatInstance string `json:"chat_instance"`
		Data         string `json:"data"`
	} `json:"callback_query"`
}

type SubscribeModel struct {
	ChatID      int    `json:"chat_id"`
	Text        string `json:"text"`
	ReplyMarkup struct {
		InlineKeyboard [][]struct {
			Text         string `json:"text"`
			CallbackData string `json:"callback_data"`
		} `json:"inline_keyboard"`
	} `json:"reply_markup"`
}
